package services

import (
	"github.com/couchbase-examples/golang-quickstart/models"

	"github.com/couchbase/gocb/v2"
)

type IAirlineService interface {
	CreateAirline(string, *models.Airline) error
	GetAirline(string) (*models.Airline, error)
	UpdateAirline(string, *models.Airline) error
	DeleteAirline(string) error
	QueryAirline(string) ([]models.Airline, error)
}

type AirlineService struct {
	collectionName string
	scope          *gocb.Scope
}

func NewAirlineService(scope *gocb.Scope) *AirlineService {
	return &AirlineService{
		collectionName: "airline",
		scope:          scope,
	}
}

func (s *AirlineService) CreateAirline(docKey string, data *models.Airline) error {
	_, err := s.scope.Collection(s.collectionName).Insert(docKey, data, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *AirlineService) GetAirline(docKey string) (*models.Airline, error) {
	getResult, err := s.scope.Collection(s.collectionName).Get(docKey, nil)
	if err != nil {
		return nil, err
	}

	var airlineData models.Airline

	if err := getResult.Content(&airlineData); err != nil {
		return nil, err
	}
	return &airlineData, nil
}

func (s *AirlineService) UpdateAirline(docKey string, data *models.Airline) error {
	_, err := s.scope.Collection(s.collectionName).Upsert(docKey, data, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *AirlineService) DeleteAirline(docKey string) error {
	_, err := s.scope.Collection(s.collectionName).Remove(docKey, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *AirlineService) QueryAirline(query string) ([]models.Airline, error) {
	queryResult, err := s.scope.Query(query, nil)
	if err != nil {
		return nil, err
	}
	var document models.Airline
	var documents []models.Airline

	if queryResult == nil {
		return nil, err
	}

	for queryResult.Next() {
		err := queryResult.Row(&document)
		if err != nil {
			return nil, err
		}
		documents = append(documents, document)
	}
	return documents, nil
}

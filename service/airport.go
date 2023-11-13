package services

import (
	"github.com/couchbase-examples/golang-quickstart/models"

	"github.com/couchbase/gocb/v2"
)

type IAirportService interface {
	CreateAirport(string, *models.Airport) error
	GetAirport(string) (*models.Airport, error)
	UpdateAirport(string, *models.Airport) error
	DeleteAirport(string) error
	QueryAirport(string) ([]models.Airport, error)
	QueryDirectConnectionAirport(string) ([]models.Destination, error)
}

type AirportService struct {
	collectionName string
	scope          *gocb.Scope
}

func NewAirportService(scope *gocb.Scope) *AirportService {
	return &AirportService{
		collectionName: "airport",
		scope:          scope,
	}
}

func (s *AirportService) CreateAirport(docKey string, data *models.Airport) error {
	_, err := s.scope.Collection(s.collectionName).Insert(docKey, data, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *AirportService) GetAirport(docKey string) (*models.Airport, error) {
	getResult, err := s.scope.Collection(s.collectionName).Get(docKey, nil)
	if err != nil {
		return nil, err
	}

	var airportData models.Airport

	if err := getResult.Content(&airportData); err != nil {
		return nil, err
	}

	return &airportData, nil
}

func (s *AirportService) UpdateAirport(docKey string, data *models.Airport) error {
	_, err := s.scope.Collection(s.collectionName).Upsert(docKey, data, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *AirportService) DeleteAirport(docKey string) error {
	_, err := s.scope.Collection(s.collectionName).Remove(docKey, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *AirportService) QueryAirport(query string) ([]models.Airport, error) {
	queryResult, err := s.scope.Query(query, nil)
	if err != nil {
		return nil, err
	}
	var document models.Airport
	var documents []models.Airport

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

func (s *AirportService) QueryDirectConnectionAirport(query string) ([]models.Destination, error) {
	queryResult, err := s.scope.Query(query, nil)
	if err != nil {
		return nil, err
	}
	var document models.Destination
	var documents []models.Destination

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

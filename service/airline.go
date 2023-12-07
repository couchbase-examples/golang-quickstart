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
	ListAirlines(string, int, int) ([]models.Airline, error)
	ListAirlinesToAirport(string, int,int) ([]models.Airline, error)
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

func (s *AirlineService) ListAirlines(country string, limit, offset int) ([]models.Airline, error) {
	var query string
	var params map[string]interface{}

	if country != "" {
		query = `
			SELECT airline.callsign,
				airline.country,
				airline.iata,
				airline.icao,
				airline.name
			FROM airline AS airline
			WHERE airline.country=$country
			ORDER BY airline.name
			LIMIT $limit
			OFFSET $offset;
		`
		params = map[string]interface{}{
			"country": country,
			"limit":   limit,
			"offset":  offset,
		}
	} else {
		query = `
			SELECT airline.callsign,
				airline.country,
				airline.iata,
				airline.icao,
				airline.name
			FROM airline AS airline
			ORDER BY airline.name
			LIMIT $limit
			OFFSET $offset;
		`
		params = map[string]interface{}{
			"limit":  limit,
			"offset": offset,
		}
	}
	queryResult, err := s.scope.Query(query, &gocb.QueryOptions{NamedParameters: params})
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

func (s *AirlineService) ListAirlinesToAirport(airport string, limit, offset int) ([]models.Airline, error) {
	// Query for airlines flying to the airport
	query := `
		SELECT air.callsign,
			air.country,
			air.iata,
			air.icao,
			air.name
		FROM (
			SELECT DISTINCT META(airline).id AS airlineId
			FROM route
			JOIN airline ON route.airlineid = META(airline).id
			WHERE route.destinationairport = $airport
		) AS subquery
		JOIN airline AS air ON META(air).id = subquery.airlineId
		ORDER BY air.name
		LIMIT $limit
		OFFSET $offset;
	`

	params := map[string]interface{}{
		"airport": airport,
		"limit":   limit,
		"offset":  offset,
	}
	queryResult, err := s.scope.Query(query, &gocb.QueryOptions{NamedParameters: params})
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

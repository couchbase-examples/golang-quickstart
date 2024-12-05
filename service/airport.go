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
	ListAirport(string, int, int) ([]models.Airport, error)
	ListDirectConnection(string, int, int) ([]models.Destination, error)
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

func (s *AirportService) ListAirport(country string, limit, offset int) ([]models.Airport, error) {
	var query string
	var params map[string]interface{}

	if country != "" {
		query = `
                SELECT airport.airportname,
                    airport.city,
                    airport.country,
                    airport.faa,
                    airport.geo,
                    airport.icao,
                    airport.tz
                FROM airport AS airport
                WHERE airport.country = $country
                ORDER BY airport.airportname
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
                SELECT airport.airportname,
                    airport.city,
                    airport.country,
                    airport.faa,
                    airport.geo,
                    airport.icao,
                    airport.tz
                FROM airport AS airport
                ORDER BY airport.airportname
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

func (s *AirportService) ListDirectConnection(airport string, limit, offset int) ([]models.Destination, error) {
	query := `
	SELECT DISTINCT route.destinationairport
	FROM airport AS airport
	JOIN route AS route ON route.sourceairport = airport.faa
	WHERE airport.faa = $airport AND route.stops = 0
	ORDER BY route.destinationairport
	LIMIT $limit
	OFFSET $offset
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

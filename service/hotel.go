package services

import (
	"github.com/couchbase-examples/golang-quickstart/models"
	"github.com/couchbase/gocb/v2"
	"github.com/couchbase/gocb/v2/search"
)

const (
	nameField        = "name"
	nameKeywordField = "name_keyword"
	titleField       = "title"
	descriptionField = "description"
	countryField     = "country"
	stateField       = "state"
	cityField        = "city"
)

type IHotelService interface {
	SearchByName(string) ([]string, error)
	Filter(filter *models.HotelSearchRequest) ([]models.HotelSearch, error)
}

type HotelService struct {
	IndexName string
	scope     *gocb.Scope
	cluster   *gocb.Cluster
}

func NewHotelService(scope *gocb.Scope, indexName string) *HotelService {
	return &HotelService{
		IndexName: indexName,
		scope:     scope,
	}
}

func (bs *HotelService) SearchByName(name string) ([]string, error) {
	names := []string{}
	request := gocb.SearchRequest{
		SearchQuery: search.NewMatchQuery(name).Field(nameField),
	}
	result, err := bs.scope.Search(bs.IndexName, request, &gocb.SearchOptions{
		Limit:  50,
		Fields: []string{nameField},
		Sort: []search.Sort{
			"-_score", "name_keyword",
		},
	})
	if err != nil {
		return nil, err
	}
	defer result.Close()
	for result.Next() {
		row := result.Row()
		var hotel models.HotelSearch
		err := row.Fields(&hotel)
		if err != nil {
			panic(err)
		}
		names = append(names, hotel.Name)
	}
	err = result.Err()
	if err != nil {
		return nil, err
	}
	return names, err
}

func (bs *HotelService) Filter(filter *models.HotelSearchRequest) ([]models.HotelSearch, error) {
	hotels := []models.HotelSearch{}
	query := search.NewConjunctionQuery()
	if filter.Name != "" {
		query.And(search.NewMatchQuery(filter.Name).Field(nameKeywordField))
	}
	if filter.Title != "" {
		query.And(search.NewMatchQuery(filter.Title).Field(titleField))
	}
	if filter.Description != "" {
		query.And(search.NewMatchQuery(filter.Description).Field(descriptionField))
	}
	if filter.City != "" {
		query.And(search.NewTermQuery(filter.City).Field(cityField))
	}
	if filter.State != "" {
		query.And(search.NewTermQuery(filter.State).Field(stateField))
	}
	if filter.Country != "" {
		query.And(search.NewTermQuery(filter.Country).Field(countryField))
	}
	request := gocb.SearchRequest{
		SearchQuery: query,
	}
	searchOptions := &gocb.SearchOptions{
		Fields: []string{"*"},
		Limit:  50,
	}
	if filter.Offset > 0 {
		searchOptions.Skip = filter.Offset
	}
	if filter.Limit > 0 {
		searchOptions.Limit = filter.Limit
	}
	result, err := bs.scope.Search(bs.IndexName, request, searchOptions)
	if err != nil {
		return nil, err
	}
	defer result.Close()
	for result.Next() {
		row := result.Row()
		var hotel models.HotelSearch
		err := row.Fields(&hotel)
		if err != nil {
			panic(err)
		}
		hotels = append(hotels, hotel)
	}
	err = result.Err()
	if err != nil {
		return nil, err
	}
	return hotels, err
}

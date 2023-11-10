package services

import (
	"src/models"

	"github.com/couchbase/gocb/v2"
)

type IRouteService interface {
	CreateRoute(string, *models.Route) error
	GetRoute(string) (*models.Route, error)
	UpdateRoute(string, *models.Route) error
	DeleteRoute(string) error
}

type RouteService struct {
	collectionName string
	sharedScope    *gocb.Scope
}

func NewRouteService(sharedScope *gocb.Scope) *RouteService {
	return &RouteService{
		collectionName: "route",
		sharedScope:    sharedScope,
	}
}

func (s *RouteService) CreateRoute(docKey string, data *models.Route) error {
	_, err := s.sharedScope.Collection(s.collectionName).Insert(docKey, data, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *RouteService) GetRoute(docKey string) (*models.Route, error) {
	getResult, err := s.sharedScope.Collection(s.collectionName).Get(docKey, nil)
	if err != nil {
		return nil, err
	}

	var routeData models.Route

	if err := getResult.Content(&routeData); err != nil {
		return nil, err
	}

	return &routeData, nil
}

func (s *RouteService) UpdateRoute(docKey string, data *models.Route) error {
	_, err := s.sharedScope.Collection(s.collectionName).Upsert(docKey, data, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *RouteService) DeleteRoute(docKey string) error {
	_, err := s.sharedScope.Collection(s.collectionName).Remove(docKey, nil)
	if err != nil {
		return err
	}
	return nil
}

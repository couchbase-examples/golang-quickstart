package test

import (
	"testing"

	"os"

	"github.com/couchbase-examples/golang-quickstart/db"
	service "github.com/couchbase-examples/golang-quickstart/service"
	"github.com/couchbase/gocb/v2"
)

var (
	cluster        *gocb.Cluster
	scope          *gocb.Scope
	airlineService *service.AirlineService
	airportService *service.AirportService
	routeService   *service.RouteService
)

func TestMain(m *testing.M) {
	cluster = db.InitializeCluster()

	// Initialize the scope
	scope = db.GetScope(cluster)

	// Create service instances
	airlineService = service.NewAirlineService(scope)
	airportService = service.NewAirportService(scope)
	routeService = service.NewRouteService(scope)
	exitVal := m.Run()

	os.Exit(exitVal)
}

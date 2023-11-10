package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"src/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

var collectionBaseForRoute = "http://127.0.0.1:8080"

func TestAddRoute(t *testing.T) {
	documentID := "route_test_add"
	url := collectionBaseForRoute + "/api/v1/route/" + documentID

	// Define the route data
	routeData := models.Route{
		Airline:            "AF",
		Airline_id:         "airline_10",
		SourceAirport:      "SFO",
		DestinationAirport: "JFK",
		Stops:              0,
		Equipment:          "CRJ",
		Distance:           4151.79,
	}

	// Convert the data to JSON
	requestData, err := json.Marshal(routeData)
	if err != nil {
		t.Fatal(err)
	}

	// Send a POST request to add the route
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestData))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	// Ensure that the POST request was successful (HTTP status 201)
	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected status code %d, got %d", http.StatusCreated, resp.StatusCode)
	}

	// Fetch the document to validate it was stored correctly
	getResp, err := http.Get(url)
	if err != nil {
		t.Fatal(err)
	}
	defer getResp.Body.Close()

	// Deserialize the response JSON
	var retrievedData models.Route
	decoder := json.NewDecoder(getResp.Body)
	err = decoder.Decode(&retrievedData)
	if err != nil {
		t.Fatal(err)
	}

	// Validate the retrieved document
	assert.Equal(t, routeData, retrievedData)

	// Clean up (delete the document)
	deleteReq, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		t.Fatal(err)
	}
	deleteResp, err := http.DefaultClient.Do(deleteReq)
	if err != nil {
		t.Fatal(err)
	}
	defer deleteResp.Body.Close()

	// Ensure that the DELETE request was successful (HTTP status 204)
	if deleteResp.StatusCode != http.StatusNoContent {
		t.Errorf("Expected status code %d, got %d", http.StatusNoContent, deleteResp.StatusCode)
	}
}

func TestAddDuplicateRoute(t *testing.T) {
	documentID := "route_test_duplicate"
	url := collectionBaseForRoute + "/api/v1/route/" + documentID

	// Define the route data
	routeData := models.Route{
		Airline:            "AF",
		Airline_id:         "airline_10",
		SourceAirport:      "SFO",
		DestinationAirport: "JFK",
		Stops:              0,
		Equipment:          "CRJ",
		Distance:           4151.79,
	}

	// Convert the data to JSON
	requestData, err := json.Marshal(routeData)
	if err != nil {
		t.Fatal(err)
	}

	// Create the initial route (HTTP POST request)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestData))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	// Ensure that the POST request was successful (HTTP status 201)
	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected status code %d, got %d", http.StatusCreated, resp.StatusCode)
	}

	// Attempt to create a duplicate (HTTP POST request)
	resp, err = http.Post(url, "application/json", bytes.NewBuffer(requestData))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	// Ensure that the POST request returned an HTTP status 409 (Conflict)
	if resp.StatusCode != http.StatusConflict {
		t.Errorf("Expected status code %d, got %d", http.StatusConflict, resp.StatusCode)
	}

	// Clean up (delete the document)
	deleteReq, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		t.Fatal(err)
	}
	deleteResp, err := http.DefaultClient.Do(deleteReq)
	if err != nil {
		t.Fatal(err)
	}
	defer deleteResp.Body.Close()

	// Ensure that the DELETE request was successful (HTTP status 204)
	if deleteResp.StatusCode != http.StatusNoContent {
		t.Errorf("Expected status code %d, got %d", http.StatusNoContent, deleteResp.StatusCode)
	}
}

func TestAddRouteWithoutRequiredFields(t *testing.T) {
	documentID := "route_test_invalid_payload"
	url := collectionBaseForRoute + "/api/v1/route/" + documentID

	// Define the route data without required fields
	routeData := models.Route{
		// Missing required fields
	}

	// Convert the data to JSON
	requestData, err := json.Marshal(routeData)
	if err != nil {
		t.Fatal(err)
	}

	// Create a route without required fields (HTTP POST request)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestData))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	// Ensure that the POST request returned an HTTP status 400 (Bad Request)
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, resp.StatusCode)
	}
}

func TestReadRoute(t *testing.T) {
	documentID := "route_test_read"
	url := collectionBaseForRoute + "/api/v1/route/" + documentID

	// Define the route data
	routeData := models.Route{
		Airline:            "AF",
		Airline_id:         "airline_10",
		SourceAirport:      "SFO",
		DestinationAirport: "JFK",
		Stops:              0,
		Equipment:          "CRJ",
		Distance:           4151.79,
	}

	// Convert the data to JSON
	requestData, err := json.Marshal(routeData)
	if err != nil {
		t.Fatal(err)
	}

	// Create the route (HTTP POST request)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestData))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	// Fetch the route (HTTP GET request)
	getResp, err := http.Get(url)
	if err != nil {
		t.Fatal(err)
	}
	defer getResp.Body.Close()

	// Ensure that the GET request was successful (HTTP status 200)
	if getResp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, getResp.StatusCode)
	}

	// Validate the retrieved data
	var retrievedData models.Route
	err = json.NewDecoder(getResp.Body).Decode(&retrievedData)
	if err != nil {
		t.Fatal(err)
	}

	// Now, compare the retrieved data with the expected data
	assert.Equal(t, routeData, retrievedData)

	// Clean up (delete the document)
	deleteReq, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		t.Fatal(err)
	}
	deleteResp, err := http.DefaultClient.Do(deleteReq)
	if err != nil {
		t.Fatal(err)
	}
	defer deleteResp.Body.Close()

	// Ensure that the DELETE request was successful (HTTP status 204)
	if deleteResp.StatusCode != http.StatusNoContent {
		t.Errorf("Expected status code %d, got %d", http.StatusNoContent, deleteResp.StatusCode)
	}
}

func TestReadInvalidRoute(t *testing.T) {
	documentID := "route_test_invalid_id"
	url := collectionBaseForRoute + "/api/v1/route/" + documentID

	// Fetch an invalid route (HTTP GET request)
	resp, err := http.Get(url)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	// Ensure that the GET request returned an HTTP status 404 (Not Found)
	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Expected status code %d, got %d", http.StatusNotFound, resp.StatusCode)
	}
}
func TestUpdateRoute(t *testing.T) {

	documentID := "route_test_update"
	url := collectionBaseForRoute + "/api/v1/route/" + documentID

	routeData := models.Route{
		Airline:            "AF",
		Airline_id:         "airline_10",
		SourceAirport:      "SFO",
		DestinationAirport: "JFK",
		Stops:              0,
		Equipment:          "CRJ",
		Distance:           4151.79,
	}

	requestData, err := json.Marshal(routeData)
	if err != nil {
		t.Fatal(err)
	}

	// Create the route (HTTP POST request)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestData))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	// Update the route (HTTP PUT request)
	updatedRouteData := models.Route{
		Airline:            "Updated Airline",
		Airline_id:         "updated_airline_10",
		SourceAirport:      "SFO",
		DestinationAirport: "LAX",
		Stops:              1,
		Equipment:          "B737",
		Distance:           2000.0,
	}

	updatedData, err := json.Marshal(updatedRouteData)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(updatedData))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	updateResp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer updateResp.Body.Close()

	// Ensure that the PUT request was successful (HTTP status 200)
	if updateResp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, updateResp.StatusCode)
	}

	// Fetch the updated route (HTTP GET request)
	resp, err = http.Get(url)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	// Ensure that the GET request was successful (HTTP status 200)
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}

	// Validate the retrieved data
	var retrievedData models.Route
	err = json.NewDecoder(resp.Body).Decode(&retrievedData)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, updatedRouteData, retrievedData)
}

func TestUpdateRouteWithInvalidData(t *testing.T) {

	documentID := "route_test_update_invalid_doc"
	url := collectionBaseForRoute + "/api/v1/route/" + documentID

	// Create the route with invalid data (HTTP POST request)
	initialRouteData := models.Route{
		Airline:            "Invalid Airline",
		SourceAirport:      "Invalid Airport",
		DestinationAirport: "Invalid Airport",
		Stops:              -1,
		Equipment:          "Invalid Equipment",
		Distance:           -1000.0,
	}

	requestData, err := json.Marshal(initialRouteData)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestData))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	// Ensure that the POST request was not successful (HTTP status other than 200)
	if resp.StatusCode != http.StatusBadRequest {
		t.Fatalf("Expected status code %d, got %d", http.StatusBadRequest, resp.StatusCode)
	}
}

func TestDeleteRoute(t *testing.T) {

	routeData := models.Route{
		Airline:            "AF",
		Airline_id:         "airline_10",
		SourceAirport:      "SFO",
		DestinationAirport: "JFK",
		Stops:              0,
		Equipment:          "CRJ",
		Distance:           4151.79,
	}
	documentID := "route_test_delete"

	// Create the document (HTTP POST request)
	url := collectionBaseForRoute + "/api/v1/route/" + documentID
	requestData, err := json.Marshal(routeData)
	if err != nil {
		t.Fatal(err)
	}

	postResp, err := http.Post(url, "application/json", bytes.NewBuffer(requestData))
	if err != nil {
		t.Fatal(err)
	}
	defer postResp.Body.Close()

	// Ensure that the POST request was successful (HTTP status 201)
	if postResp.StatusCode != http.StatusCreated {
		t.Fatalf("Expected status code %d, got %d", http.StatusCreated, postResp.StatusCode)
	}

	// Delete the created document (HTTP DELETE request)
	deleteURL := collectionBaseForRoute + "/api/v1/route/" + documentID
	req, err := http.NewRequest("DELETE", deleteURL, nil)
	if err != nil {
		t.Fatal(err)
	}

	client := &http.Client{}
	deleteResp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer deleteResp.Body.Close()

	// Ensure that the DELETE request was successful (HTTP status 204)
	if deleteResp.StatusCode != http.StatusNoContent {
		t.Fatalf("Expected status code %d, got %d", http.StatusNoContent, deleteResp.StatusCode)
	}
}

func TestDeleteInvalidDocumentRoute(t *testing.T) {

	invalidDocumentID := "non_existent_document"

	// Attempt to delete a non-existing document (HTTP DELETE request)
	url := collectionBaseForRoute + "/api/v1/route/" + invalidDocumentID

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		t.Fatal(err)
	}

	client := &http.Client{}
	deleteResp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer deleteResp.Body.Close()

	// Ensure that the DELETE request was not successful (HTTP status other than 204)
	if deleteResp.StatusCode != http.StatusNotFound {
		t.Errorf("Expected a non-204 status code for the invalid delete, got %d", deleteResp.StatusCode)
	}
}

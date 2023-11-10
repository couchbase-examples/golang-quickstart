package test

import (
	"bytes"
	"encoding/json"

	"github.com/couchbase-examples/golang-quickstart/models"

	"fmt"
	"net/http"

	"testing"

	"github.com/stretchr/testify/assert"
)

var collectionBaseForAirline = "http://127.0.0.1:8080"

func TestAddairline(t *testing.T) {
	documentID := "airline_test_add"
	url := collectionBaseForAirline + "/api/v1/airline/" + documentID

	// Define the airline data
	airlineData := models.Airline{
		Name:     "Sample Airline",
		ICAO:     "SALL",
		Callsign: "SAM",
		Country:  "Sample Country",
	}

	// Convert the data to JSON
	requestData, err := json.Marshal(airlineData)
	if err != nil {
		t.Fatal(err)
	}

	// Send a POST request to add the airport
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
	var retrievedData models.Airline
	decoder := json.NewDecoder(getResp.Body)
	err = decoder.Decode(&retrievedData)
	if err != nil {
		t.Fatal(err)
	}

	// Validate the retrieved document
	assert.Equal(t, airlineData, retrievedData)

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

func TestAddDuplicateAirline(t *testing.T) {

	documentID := "airline_test_duplicate"
	url := collectionBaseForAirline + "/api/v1/airline/" + documentID

	airlineData := models.Airline{
		Name:     "Sample Airline",
		ICAO:     "SALL",
		Callsign: "SAM",
		Country:  "Sample Country",
	}

	// Convert the data to JSON
	requestData, err := json.Marshal(airlineData)
	if err != nil {
		t.Fatal(err)
	}

	// Create the initial airline (HTTP POST request)
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

func TestAddAirlineWithoutRequiredFields(t *testing.T) {

	documentID := "airline_test_invalid_payload"
	url := collectionBaseForAirline + "/api/v1/airline/" + documentID

	airlineData := models.Airline{
		ICAO:    "SALL",
		Country: "Sample Country",
	}

	// Convert the data to JSON
	requestData, err := json.Marshal(airlineData)
	if err != nil {
		t.Fatal(err)
	}

	// Create an airline without required fields (HTTP POST request)
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

func TestReadAirline(t *testing.T) {

	documentID := "airline_test_read"
	url := collectionBaseForAirline + "/api/v1/airline/" + documentID

	airlineData := models.Airline{
		Name:     "Sample Airline",
		ICAO:     "SALL",
		Callsign: "SAM",
		Country:  "Sample Country",
	}

	requestData, err := json.Marshal(airlineData)
	if err != nil {
		t.Fatal(err)
	}

	// Create the airline (HTTP POST request)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestData))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	// Fetch the airline (HTTP GET request)
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
	var retrievedData models.Airline
	err = json.NewDecoder(getResp.Body).Decode(&retrievedData)
	if err != nil {
		t.Fatal(err)
	}

	// Now, compare the retrieved data with the expected data
	if retrievedData != airlineData {
		t.Errorf("Retrieved data does not match expected data. Expected: %v, Actual: %v", airlineData, retrievedData)
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

func TestReadInvalidAirline(t *testing.T) {

	documentID := "airline_test_invalid_id"
	url := collectionBaseForAirline + "/api/v1/airline/" + documentID

	// Fetch an invalid airline (HTTP GET request)
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

func TestUpdateAirline(t *testing.T) {

	documentID := "airline_test_update"
	url := collectionBaseForAirline + "/api/v1/airline/" + documentID

	airlineData := models.Airline{
		Name:     "Sample Airline",
		ICAO:     "SALL",
		Callsign: "SAM",
		Country:  "Sample Country",
	}

	requestData, err := json.Marshal(airlineData)
	if err != nil {
		t.Fatal(err)
	}

	// Create the airline (HTTP POST request)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestData))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	// Update the airline (HTTP PUT request)
	updatedAirlineData := models.Airline{
		Name:     "Updated Airline",
		IATA:     "SAL",
		ICAO:     "SALL",
		Callsign: "SAM",
		Country:  "Updated Country",
	}

	updatedData, err := json.Marshal(updatedAirlineData)
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

	// Fetch the updated airline (HTTP GET request)
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
	var retrievedData models.Airline
	err = json.NewDecoder(resp.Body).Decode(&retrievedData)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, updatedAirlineData, retrievedData)

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

func TestUpdateAirlineWithInvalidData(t *testing.T) {
	documentID := "airline_test_update_invalid_doc"
	url := collectionBaseForAirline + "/api/v1/airline/" + documentID

	// Create the airline with invalid data (HTTP POST request)
	initialAirlineData := models.Airline{
		ICAO:     "SALL",
		Callsign: "SAM",
		Country:  "Sample Country",
	}

	requestData, err := json.Marshal(initialAirlineData)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestData))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	// Ensure that the PUT request was not successful (HTTP status other than 200)
	if resp.StatusCode != http.StatusBadRequest {
		t.Fatalf("Expected status code %d, got %d", http.StatusCreated, resp.StatusCode)
	}
}

func TestDeletAirline(t *testing.T) {
	airlineData := models.Airline{
		Name:     "Sample Airline",
		ICAO:     "SALL",
		Callsign: "SAM",
		Country:  "Sample Country",
	}
	documentID := "airline_test_delete"

	// Create the document (HTTP POST request)
	url := collectionBaseForAirline + "/api/v1/airline/" + documentID
	requestData, err := json.Marshal(airlineData)
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
	deleteURL := collectionBaseForAirline + "/api/v1/airline/" + documentID
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

func TestDeleteInvalidDocumentAirline(t *testing.T) {
	invalidDocumentID := "non_existent_document"

	// Attempt to delete an non existing document (HTTP DELETE request)
	url := collectionBaseForAirline + "/api/v1/airline/" + invalidDocumentID

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

func TestListAirlinesInCountry(t *testing.T) {
	country := "France"

	url := "http://127.0.0.1:8080/api/v1/airline/list?country=" + country
	response, err := http.Get(url)
	if err != nil {
		t.Fatal(err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, response.StatusCode)
	}

	var result []models.Airline
	decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(&result)
	if err != nil {
		t.Fatal(err)
	}

	// Access and validate the retrieved data
	for _, item := range result {
		if item.Country != country {
			t.Errorf("Expected country %s, got %s", country, item.Country)
		}
	}
}

func TestListAirlinesInCountryWithPagination(t *testing.T) {
	country := "France"
	pageSize := 1
	iterations := 2
	airlinesList := make(map[string]bool)

	for i := 0; i < iterations; i++ {
		url := fmt.Sprintf("http://127.0.0.1:8080/api/v1/airline/list?country=%s&limit=%d&offset=%d", country, pageSize, pageSize*i)

		response, err := http.Get(url)
		if err != nil {
			t.Fatal(err)
		}
		defer response.Body.Close()

		if response.StatusCode != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, response.StatusCode)
		}

		var result []models.Airline
		decoder := json.NewDecoder(response.Body)
		err = decoder.Decode(&result)
		if err != nil {
			t.Fatal(err)
		}

		if len(result) != pageSize {
			t.Errorf("Expected %d items in the response, got %d", pageSize, len(result))
		}

		for _, item := range result {
			airlinesList[item.Name] = true

			if item.Country != country {
				t.Errorf("Expected country %s, got %s", country, item.Country)
			}
		}
	}

	if len(airlinesList) != pageSize*iterations {
		t.Errorf("Expected %d unique airline names in the response, got %d", pageSize*iterations, len(airlinesList))
	}
}

func TestListAirlinesInInvalidCountry(t *testing.T) {
	url := "http://127.0.0.1:8080/api/v1/airline/list?country=invalid"

	response, err := http.Get(url)
	if err != nil {
		t.Fatal(err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusInternalServerError {
		t.Errorf("Expected status code %d, got %d", http.StatusInternalServerError, response.StatusCode)
	}

}

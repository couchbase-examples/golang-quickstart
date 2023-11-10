package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/couchbase-examples/golang-quickstart/models"

	"github.com/stretchr/testify/assert"
)

var collectionBaseForAirport = "http://127.0.0.1:8080"

func TestAddairport(t *testing.T) {

	documentID := "airport_test_add"
	url := collectionBaseForAirport + "/api/v1/airport/" + documentID

	// Define the airport data
	airportData := models.Airport{
		AirportName: "SampleAirport",
		City:        "SampleCity",
		Country:     "United Kingdom",
		FAA:         "SAA",
		ICAO:        "SAAA",
		TZ:          "Europe/Paris",
	}

	// Convert the data to JSON
	requestData, err := json.Marshal(airportData)
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
	var retrievedData models.Airport
	decoder := json.NewDecoder(getResp.Body)
	err = decoder.Decode(&retrievedData)
	if err != nil {
		t.Fatal(err)
	}

	// Validate the retrieved document
	assert.Equal(t, airportData, retrievedData)

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
func TestAddDuplicateairport(t *testing.T) {
	documentID := "airport_test_duplicate"
	url := collectionBaseForAirport + "/api/v1/airport/" + documentID

	airportData := models.Airport{
		AirportName: "Test Airport",
		City:        "Test City",
		Country:     "Test Country",
		FAA:         "TAA",
		ICAO:        "TAAS",
		TZ:          "Europe/Berlin",
	}

	requestData, err := json.Marshal(airportData)
	if err != nil {
		t.Fatal(err)
	}

	// Create the initial airport (HTTP POST request)
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

func TestAddairportWithoutRequiredFields(t *testing.T) {
	documentID := "airport_test_invalid"
	url := collectionBaseForAirport + "/api/v1/airport/" + documentID
	// Missing Required Airport Name field
	airportData := models.Airport{
		City:    "Test City",
		Country: "Test Country",
		FAA:     "TAA",
		ICAO:    "TAAS",
		TZ:      "Europe/Berlin",
	}
	requestData, err := json.Marshal(airportData)
	if err != nil {
		t.Fatal(err)
	}

	// Create an airport without required fields (HTTP POST request)
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

func TestReadairport(t *testing.T) {

	documentID := "airport_test_read"
	url := collectionBaseForAirport + "/api/v1/airport/" + documentID

	airportData := models.Airport{
		AirportName: "Test Airport",
		City:        "Test City",
		Country:     "Test Country",
		FAA:         "TAA",
		ICAO:        "TAAS",
		TZ:          "Europe/Berlin",
	}

	requestData, err := json.Marshal(airportData)
	if err != nil {
		t.Fatal(err)
	}

	// Create the airport (HTTP POST request)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestData))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	// Fetch the airport (HTTP GET request)
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
	var retrievedData models.Airport
	err = json.NewDecoder(resp.Body).Decode(&retrievedData)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, airportData, retrievedData)

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

func TestReadInvalidairport(t *testing.T) {
	documentID := "airport_test_invalid_id"
	url := collectionBaseForAirport + "/api/v1/airport/" + documentID

	// Fetch an invalid airport (HTTP GET request)
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

func TestUpdateairport(t *testing.T) {
	documentID := "airport_test_update"
	url := collectionBaseForAirport + "/api/v1/airport/" + documentID

	initialAirportData := models.Airport{
		AirportName: "Test Airport",
		City:        "Test City",
		Country:     "Test Country",
		FAA:         "TAA",
		ICAO:        "TAAS",
		TZ:          "Europe/Berlin",
	}

	requestData, err := json.Marshal(initialAirportData)
	if err != nil {
		t.Fatal(err)
	}

	// Create the airport (HTTP POST request)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestData))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	// Update the airport (HTTP PUT request)
	updatedAirportData := models.Airport{
		AirportName: "Updated Airport",
		City:        "Updated City",
		Country:     "Updated Country",
		FAA:         "TAA",
		ICAO:        "TAAS",
		TZ:          "USA",
	}

	updatedData, err := json.Marshal(updatedAirportData)
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

	// Fetch the updated airport (HTTP GET request)
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
	var retrievedData models.Airport
	err = json.NewDecoder(resp.Body).Decode(&retrievedData)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, updatedAirportData, retrievedData)

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

func TestUpdateAirportWithInvalidData(t *testing.T) {

	documentID := "airport_test_update_invalid_doc"
	url := collectionBaseForAirport + "/api/v1/airport/" + documentID

	// Create the airline with invalid data (HTTP POST request)
	airportData := models.Airport{
		City:    "Test City",
		Country: "Test Country",
		FAA:     "TAA",
		ICAO:    "TAAS",
		TZ:      "Europe/Berlin",
	}

	requestData, err := json.Marshal(airportData)
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

func TestDeletAirport(t *testing.T) {
	airportData := models.Airport{
		AirportName: "Test Airport",
		City:        "Test City",
		Country:     "Test Country",
		FAA:         "TAA",
		ICAO:        "TAAS",
		TZ:          "Europe/Berlin",
	}
	documentID := "airport_test_delete"

	// Create the document (HTTP POST request)
	url := collectionBaseForAirport + "/api/v1/airport/" + documentID
	requestData, err := json.Marshal(airportData)
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
	deleteURL := collectionBaseForAirport + "/api/v1/airport/" + documentID
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

func TestDeleteAirportInvalidDocument(t *testing.T) {
	invalidDocumentID := "non_existent_document"

	// Attempt to delete an non existing document (HTTP DELETE request)
	url := collectionBaseForAirport + "/api/v1/airport/" + invalidDocumentID

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

func TestListAirportsInCountryWithPagination(t *testing.T) {
	country := "France"
	pageSize := 1
	iterations := 2
	airportsList := make(map[string]bool)

	for i := 0; i < iterations; i++ {
		url := fmt.Sprintf("http://127.0.0.1:8080/api/v1/airport/list?country=%s&limit=%d&offset=%d", country, pageSize, pageSize*i)

		response, err := http.Get(url)
		if err != nil {
			t.Fatal(err)
		}
		defer response.Body.Close()

		if response.StatusCode != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, response.StatusCode)
		}

		var result []models.Airport
		decoder := json.NewDecoder(response.Body)
		err = decoder.Decode(&result)
		if err != nil {
			t.Fatal(err)
		}

		if len(result) != pageSize {
			t.Errorf("Expected %d items in the response, got %d", pageSize, len(result))
		}

		for _, item := range result {
			airportsList[item.AirportName] = true

			if item.Country != country {
				t.Errorf("Expected country %s, got %s", country, item.Country)
			}
		}
	}

	if len(airportsList) != pageSize*iterations {
		t.Errorf("Expected %d unique airport names in the response, got %d", pageSize*iterations, len(airportsList))
	}
}

func TestListAirportsInInvalidCountry(t *testing.T) {
	url := "http://127.0.0.1:8080/api/v1/airport/list?country=invalid"

	response, err := http.Get(url)
	if err != nil {
		t.Fatal(err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusInternalServerError {
		t.Errorf("Expected status code %d, got %d", http.StatusInternalServerError, response.StatusCode)
	}

}

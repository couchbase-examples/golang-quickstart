package test

import (
	"bytes"
	"encoding/json"
	"github.com/couchbase-examples/golang-quickstart/models"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHotelAutoComplete(t *testing.T) {
	url := collectionBaseForRoute + "/api/v1/hotel/autocomplete?name=sea"
	resp, err := http.Get(url)
	assert.Nil(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode, "StatusCode")

	resultByte, err := io.ReadAll(resp.Body)
	assert.Nil(t, err)

	var result []string
	err = json.Unmarshal(resultByte, &result)
	assert.Nil(t, err)

	assert.Equal(t, result, []string{
		"Seal View", "Seal Rock Inn", "Seal Cove Inn", "Seaview Hotel", "Sea Ranch Lodge", "Sea-Change Living",
		"Four Seasons Hotel", "Four Seasons Hotel", "Four Seasons Hotel George V", "Best Western Blue Sea Lodge",
		"Carlsbad Seapointe Resort", "Beautiful Wales Sea View Luxury Holiday Barns", "Hotel Sergul", "Serrano Hotel",
		"Selby Campground", "Camping du Mont Serein", "Gower Edge Self Catering", "Beach House Self Catering",
		"Grand Hotel Serre Chevalier", "Lamington Serviced Apartments", "Forest Cottage Self-Catering",
		"Stangwrach Self-catering Cottage", "Sequoia River Dance Bed & Breakfast", "Nant Hafod, Self-catering Bungalow",
		"Coedmor Self Catering Holiday Cottages",
	})
}

func TestHotelFetchCase1(t *testing.T) {
	url := collectionBaseForRoute + "/api/v1/hotel/filter"

	hotelFilter := models.HotelSearchRequest{
		HotelSearch: models.HotelSearch{
			Title:       "Carrizo Plain National Monument",
			Name:        "KCL Campground",
			Country:     "United States",
			City:        "Santa Margarita",
			State:       "California",
			Description: "newly renovated",
		},
	}

	expectedHotels := []models.HotelSearch{{
		Title:   "Carrizo Plain National Monument",
		Name:    "KCL Campground",
		Country: "United States",
		City:    "Santa Margarita",
		State:   "California",
		Description: "The campground has a gravel road, pit toilets, corrals and water for livestock. " +
			"There are some well established shade trees and the facilities have just been renovated to include new " +
			"fire rings with BBQ grates, lantern poles, and gravel roads and tent platforms.  " +
			"Tenters, and small to medium sized campers will find the KCL a good fit.",
	}}

	// Convert the data to JSON
	hotelFilterData, err := json.Marshal(hotelFilter)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(hotelFilterData))
	assert.Nil(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode, "StatusCode")

	resultByte, err := io.ReadAll(resp.Body)
	assert.Nil(t, err)

	var result []models.HotelSearch
	err = json.Unmarshal(resultByte, &result)
	assert.Nil(t, err)

	assert.Equal(t, result, expectedHotels)
}

func TestHotelFetchCase2(t *testing.T) {
	url := collectionBaseForRoute + "/api/v1/hotel/filter"

	hotelFilter := models.HotelSearchRequest{
		HotelSearch: models.HotelSearch{
			Description: "newly renovated",
		},
	}

	// Convert the data to JSON
	hotelFilterData, err := json.Marshal(hotelFilter)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(hotelFilterData))
	assert.Nil(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode, "StatusCode")

	resultByte, err := io.ReadAll(resp.Body)
	assert.Nil(t, err)

	var result []models.HotelSearch
	err = json.Unmarshal(resultByte, &result)
	assert.Nil(t, err)
	assert.Greater(t, len(result), 2)
}

func TestHotelFetchCase3(t *testing.T) {
	url := collectionBaseForRoute + "/api/v1/hotel/filter"

	hotelFilter := models.HotelSearchRequest{
		HotelSearch: models.HotelSearch{
			Description: "newly renovated",
		},
		Offset: 5,
		Limit:  2,
	}

	// Convert the data to JSON
	hotelFilterData, err := json.Marshal(hotelFilter)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(hotelFilterData))
	assert.Nil(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode, "StatusCode")

	resultByte, err := io.ReadAll(resp.Body)
	assert.Nil(t, err)

	var result []models.HotelSearch
	err = json.Unmarshal(resultByte, &result)
	assert.Nil(t, err)
	assert.Equal(t, len(result), 2)
}

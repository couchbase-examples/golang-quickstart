package controllers

import (
	"fmt"
	"net/http"
	"src/models"
	"src/responses"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary      Insert Document
// @Description  Inserts a document into the "airline" collection.
// @Tags         Airline collection
// @Produce      json
// @Param        id path string true "Create document by specifying ID.    Example: airline_123"
// @Param        data body models.RequestBodyForAirline true "Data to create a new document"
// @Success      201 {object} responses.TravelSampleResponseForAirline
// @Failure      400 "Bad Request"
// @Failure      409 "Airline document already exists"
// @Failure      500 "Internal Server Error"
// @Router       /api/v1/airline/{id} [post]
func InsertDocumentForAirline() gin.HandlerFunc {
	return func(context *gin.Context) {
		var data models.RequestBodyForAirline
		docKey := context.Param("id")
		// Bind the JSON data to the "data" variable
		if err := context.BindJSON(&data); err != nil {
			context.JSON(http.StatusBadRequest, responses.TravelSampleResponse{
				Status:         http.StatusBadRequest,
				Message:        "Bad Request",
				CollectionData: "Error, Invalid request data: " + err.Error(),
			})
			return
		}
		insertDocument(context, "airline", docKey, data)
	}
}

// @Summary      Get Airline Document
// @Description  Gets a document from the "airline" collection based on the provided key.
// @Tags         Airline collection
// @Produce      json
// @Param        id path string true "Get document by ID.    Example: 'airline_123'"
// @Success      200 {object} responses.TravelSampleResponse
// @Failure      404 "Airline Document ID not found"
// @Failure      500 "Internal Server Error"
// @Router       /api/v1/airline/{id} [get]
func GetDocumentForAirline() gin.HandlerFunc {
	return func(context *gin.Context) {
		documentID := context.Param("id")
		var getDoc models.RequestBodyForAirline
		getDocument(context, "airline", documentID, &getDoc)
	}
}

// @Summary      Update Document
// @Description  Updates a document in the "airline" collection based on the provided key.
// @Tags         Airline collection
// @Produce      json
// @Param       id path string  true  "Update document by ID.    Example: 'airline_123'"
// @Param		 data body models.RequestBodyForAirline true  "Updates document"
// @Success      200  {array}  responses.TravelSampleResponse
// @Failure      400 "Bad Request"
// @Failure      500 "Internal Server Error"
// @Router       /api/v1/airline/{id} [put]
func UpdateDocumentForAirline() gin.HandlerFunc {
	return func(context *gin.Context) {
		var data models.RequestBodyForAirline
		docKey := context.Param("id")
		if err := context.BindJSON(&data); err != nil {
			context.JSON(http.StatusBadRequest, responses.TravelSampleResponse{Status: http.StatusBadRequest, Message: "Error while getting the request", CollectionData: err.Error()})
			return

		}
		updateDocument(context, "airline", docKey, data)

	}

}

// @Summary      Delete Document
// @Description  Deletes a document in the "airline" collection based on the provided key.
// @Tags         Airline collection
// @Produce      json
// @Param        id path string true "Deletes a document with the specified key. Example: 'sample_id'"
// @Success      204 {object} responses.TravelSampleResponse
// @Failure      404 "Airline Document ID Not Found"
// @Failure      500 "Internal Server Error"
// @Router       /api/v1/airline/{id} [delete]
func DeleteDocumentForAirline() gin.HandlerFunc {
	return func(context *gin.Context) {
		docKey := context.Param("id")
		deleteDocument(context, "airline", docKey)

	}
}

// @Summary      Get Airlines by Country
// @Description  Get a list of airlines filtered by country
// @Tags         Airline collection
// @Produce      json
// @Param        country query string true "Filter by country. Example: 'France'"
// @Param        limit query int false "Number of airlines to return (page size). Example: 10"
// @Param        offset query int false "Number of airlines to skip (for pagination). Example: 0"
// @Success      200 {object} responses.TravelSampleResponse
// @Failure      500 "Internal Server Error"
// @Router       /api/v1/airline/list [get]
func GetAirlines() gin.HandlerFunc {
	return func(context *gin.Context) {
		// Parse query parameters
		country := context.DefaultQuery("country", "")
		limit, err := strconv.Atoi(context.DefaultQuery("limit", "10"))
		if err != nil {
			limit = 10
		}
		offset, err := strconv.Atoi(context.DefaultQuery("offset", "0"))
		if err != nil {
			offset = 0
		}

		// Query to get a list of airlines filtered by country
		query := fmt.Sprintf(`
		SELECT airline.callsign,
			airline.country,
			airline.iata,
			airline.icao,
			airline.name
		FROM airline as airline
		WHERE airline.country='%s'
		ORDER BY airline.name
		LIMIT %d
		OFFSET %d;
	`, country, limit, offset)
		var airlines models.RequestBodyForAirline
		// Execute the query, retrieve and return the result
		GetDocumentsFromQuery(context, "airline", query, &airlines)
	}
}

// @Summary      Get Airlines Flying to Airport
// @Description  Get a list of airlines flying to a specific airport
// @Tags         Airline collection
// @Produce      json
// @Param        airport query string true "Destination airport. Example: 'JFK'"
// @Param        limit query int false "Number of airlines to return (page size). Example: 10"
// @Param        offset query int false "Number of airlines to skip (for pagination). Example: 0"
// @Success      200 {object} responses.TravelSampleResponse
// @Failure      500 "Internal Server Error"
// @Router       /api/v1/airline/to-airport [get]
func GetAirlinesToAirport() gin.HandlerFunc {
	return func(context *gin.Context) {
		// Parse query parameters
		airport := context.Query("airport")
		limit, err := strconv.Atoi(context.DefaultQuery("limit", "10"))
		if err != nil {
			limit = 10
		}
		offset, err := strconv.Atoi(context.DefaultQuery("offset", "0"))
		if err != nil {
			offset = 0
		}

		// Query for airlines flying to the airport
		query := fmt.Sprintf(`
            SELECT air.callsign,
                air.country,
                air.iata,
                air.icao,
                air.name
            FROM (
                SELECT DISTINCT META(airline).id AS airlineId
                FROM route
                JOIN airline ON route.airlineid = META(airline).id
                WHERE route.destinationairport = "%s"
            ) AS subquery
            JOIN airline AS air ON META(air).id = subquery.airlineId
            ORDER BY air.name
            LIMIT %d
            OFFSET %d;
        `, airport, limit, offset)

		// Execute the query, retrieve and return the result
		var airlines models.RequestBodyForAirline
		GetDocumentsFromQuery(context, "airline", query, &airlines)
	}
}

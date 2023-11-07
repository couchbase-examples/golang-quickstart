package controllers

import (
	"fmt"
	"net/http"
	"src/models"
	"src/responses"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary      Insert Airport Document
// @Description  Create Airport with specified ID
// @Tags         Airport collection
// @Produce      json
// @Param        id path string true "Create document by specifying ID      Example: airport_1273"
// @Param        data body models.RequestBodyForAirport true "Data to create a document"
// @Success      201 {object} responses.TravelSampleResponse
// @Failure      400 "Bad Request"
// @Failure      409 "Airport Document already exists"
// @Failure      500 "Internal Server Error"
// @Router       /api/v1/airport/{id} [post]
func InsertDocumentForAirport() gin.HandlerFunc {
	return func(context *gin.Context) {
		var data models.RequestBodyForAirport
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
		insertDocument(context, "airport", docKey, data)
	}
}

// @Summary      Get Airport Document
// @Description  Get Airport with specified ID
// @Tags         Airport collection
// @Produce      json
// @Param        id path string true "Search document by ID    Example: airport_1273"
// @Success      200 {object} responses.TravelSampleResponse
// @Failure      404 "Airport Document ID Not Found"
// @Failure      500 "Internal Server Error"
// @Router       /api/v1/airport/{id} [get]
func GetDocumentForAirport() gin.HandlerFunc {
	return func(context *gin.Context) {
		docKey := context.Param("id")
		var getDoc models.RequestBodyForAirport
		getDocument(context, "airport", docKey, &getDoc)
	}
}

// @Summary      Update Airport Document
// @Description  Update Airport with specified ID
// @Tags         Airport collection
// @Produce      json
// @Param 		 id path string  true  "Update document by id         Example: airport_1273"
// @Param		 data body models.RequestBodyForAirport true  "Updates document"
// @Success      200  {array}  responses.TravelSampleResponse
// @Failure      400 "Bad Request"
// @Failure      500			"Internal Server Error"
// @Router       /api/v1/airport/{id} [put]
func UpdateDocumentForAirport() gin.HandlerFunc {
	return func(context *gin.Context) {
		var data models.RequestBodyForAirport
		docKey := context.Param("id")
		if err := context.BindJSON(&data); err != nil {
			context.JSON(http.StatusBadRequest, responses.TravelSampleResponse{Status: http.StatusBadRequest, Message: "Error while getting the request", CollectionData: err.Error()})
			return

		}
		updateDocument(context, "airport", docKey, data)
	}

}

// @Summary      Deletes Airport Document
// @Description  Delete Airport with specified ID
// @Tags         Airport collection
// @Produce      json
// @Param 		 id  path string true  "Deletes a document with key specified      Example: airport_1273"
// @Success      204  {array}  responses.TravelSampleResponse
// @Failure 	 404			"Airport Document ID Not Found"
// @Failure      500			"Internal Server Error"
// @Router       /api/v1/airport/{id} [delete]
func DeleteDocumentForAirport() gin.HandlerFunc {
	return func(context *gin.Context) {
		docKey := context.Param("id")
		deleteDocument(context, "airport", docKey)
	}

}

// @Summary      List Airport Document
// @Description  Get list of Airports. Optionally, you can filter the list by Country
// @Tags         Airport collection
// @Produce      json
// @Param        country query string true "Country     Example: United Kingdom, France, United States"
// @Param        limit query int false "Number of airports to return (page size)     Default value : 10"
// @Param        offset query int false "Number of airports to skip (for pagination)     Default value : 0"
// @Success      200 {array} responses.TravelSampleResponse
// @Failure      500 "Internal Server Error"
// @Router       /api/v1/airport/list [get]
func GetAirports() gin.HandlerFunc {
	return func(context *gin.Context) {
		country := context.Query("country")
		limit, err := strconv.Atoi(context.DefaultQuery("limit", "10"))
		if err != nil {
			limit = 10
		}
		offset, err := strconv.Atoi(context.DefaultQuery("offset", "0"))
		if err != nil {
			offset = 0
		}

		query := fmt.Sprintf(`
            SELECT airport.airportname,
                airport.city,
                airport.country,
                airport.faa,
                airport.geo,
                airport.icao,
                airport.tz
            FROM airport AS airport
            WHERE airport.country="%s"
            ORDER BY airport.airportname
            LIMIT %d
            OFFSET %d;
        `, country, limit, offset)

		// Use the common method to execute the query and return the results
		var airports models.RequestBodyForAirport
		GetDocumentsFromQuery(context, "airports", query, &airports)
	}
}

// @Summary      Get Direct Connections from Airport
// @Description  Get Direct Connections from specified Airport
// @Tags         Airport collection
// @Produce      json
// @Param        airport query string true "Source airport       Example: SFO, LHR, CDG"
// @Param        limit query int false "Number of direct connections to return (page size)      Default value : 10"
// @Param        offset query int false "Number of direct connections to skip (for pagination)  Default value : 0"
// @Success      200 {array} responses.TravelSampleResponse
// @Failure      500 "Internal Server Error"
// @Router       /api/v1/airport/direct-connections [get]
func GetDirectConnections() gin.HandlerFunc {
	return func(context *gin.Context) {
		airport := context.Query("airport")
		limit, err := strconv.Atoi(context.DefaultQuery("limit", "10"))
		if err != nil {
			limit = 10
		}
		offset, err := strconv.Atoi(context.DefaultQuery("offset", "0"))
		if err != nil {
			offset = 0
		}

		query := fmt.Sprintf(`
            SELECT distinct (route.destinationairport)
            FROM airport as airport
            JOIN route as route on route.sourceairport = airport.faa
            WHERE airport.faa="%s" and route.stops = 0
            ORDER BY route.destinationairport
            LIMIT %d
            OFFSET %d
        `, airport, limit, offset)

		// Use the common method to execute the query and return the results
		var airports models.RequestBodyForAirport
		GetDocumentsFromQuery(context, "airports", query, &airports)
	}
}

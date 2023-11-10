package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	cError "github.com/couchbase-examples/golang-quickstart/errors"
	"github.com/couchbase-examples/golang-quickstart/models"
	services "github.com/couchbase-examples/golang-quickstart/service"

	"github.com/couchbase/gocb/v2"
	"github.com/gin-gonic/gin"
)

type AirportController struct {
	AirportService services.IAirportService
}

func NewAirportController(AirportService services.IAirportService) *AirportController {
	return &AirportController{
		AirportService: AirportService,
	}
}

// @Summary      Insert Airport Document
// @Description  Create Airport with specified ID
// @Tags         Airport collection
// @Produce      json
// @Param        id path string true "Airport ID like airport_1273"
// @Param        data body models.Airport true "Data to create a document"
// @Success      201 {object} models.Airport
// @Failure      400 "Bad Request"
// @Failure      409 "Airport Document already exists"
// @Failure      500 "Internal Server Error"
// @Router       /api/v1/airport/{id} [post]
func (ac *AirportController) InsertDocumentForAirport() gin.HandlerFunc {
	return func(context *gin.Context) {
		docKey := context.Param("id")
		data := models.Airport{}
		if err := context.ShouldBindJSON(&data); err != nil {
			context.JSON(http.StatusBadRequest, cError.Errors{
				Error: "Error, Invalid request data: " + err.Error(),
			})
			return
		}

		err := ac.AirportService.CreateAirport(docKey, &data)
		if err != nil {
			if errors.Is(err, gocb.ErrDocumentExists) {
				context.JSON(http.StatusConflict, cError.Errors{
					Error: "Error, Airport Document already exists: " + err.Error(),
				})
			} else {
				context.JSON(http.StatusInternalServerError, cError.Errors{
					Error: "Error, Airport Document could not be inserted: " + err.Error(),
				})
			}
			return
		}
		context.JSON(http.StatusCreated, data)
	}
}

// @Summary      Get Airport Document
// @Description  Get Airport with specified ID
// @Tags         Airport collection
// @Produce      json
// @Param        id path string true "Airport ID like airport_1273"
// @Success      200 {object} models.Airport
// @Failure      404 "Airport Document ID Not Found"
// @Failure      500 "Internal Server Error"
// @Router       /api/v1/airport/{id} [get]
func (ac *AirportController) GetDocumentForAirport() gin.HandlerFunc {
	return func(context *gin.Context) {
		docKey := context.Param("id")
		airportData, err := ac.AirportService.GetAirport(docKey)
		if err != nil {
			if errors.Is(err, gocb.ErrDocumentNotFound) {
				context.JSON(http.StatusNotFound, cError.Errors{
					Error: "Error, Airport Document not found",
				})
			} else {
				context.JSON(http.StatusInternalServerError, cError.Errors{
					Error: "Error, Document could not be fetched: " + err.Error(),
				})
			}
		} else {
			context.JSON(http.StatusOK, *airportData)
		}
	}
}

// @Summary      Update Airport Document
// @Description  Update Airport with specified ID
// @Tags         Airport collection
// @Produce      json
// @Param       id path string true "Airport ID like airport_1273"
// @Param       data body models.Airport true "Updates document"
// @Success      200 {object} models.Airport
// @Failure      400 "Bad Request"
// @Failure      500 "Internal Server Error"
// @Router       /api/v1/airport/{id} [put]
func (ac *AirportController) UpdateDocumentForAirport() gin.HandlerFunc {
	return func(context *gin.Context) {
		docKey := context.Param("id")
		data := models.Airport{}
		if err := context.ShouldBindJSON(&data); err != nil {
			context.JSON(http.StatusBadRequest, cError.Errors{
				Error: "Error while getting the request: " + err.Error(),
			})
			return
		}
		err := ac.AirportService.UpdateAirport(docKey, &data)
		if err != nil {
			context.JSON(http.StatusInternalServerError, cError.Errors{
				Error: "Error, Airport Document could not be updated: " + err.Error(),
			})
			return
		}
		context.JSON(http.StatusOK, data)
	}
}

// @Summary      Deletes Airport Document
// @Description  Delete Airport with specified ID
// @Tags         Airport collection
// @Produce      json
// @Param 		 id  path string true  "Airport ID like airport_1273"
// @Success      204    "Airport deleted"
// @Failure 	 404			"Airport Document ID Not Found"
// @Failure      500			"Internal Server Error"
// @Router       /api/v1/airport/{id} [delete]
func (ac *AirportController) DeleteDocumentForAirport() gin.HandlerFunc {
	return func(context *gin.Context) {
		docKey := context.Param("id")
		err := ac.AirportService.DeleteAirport(docKey)
		if err != nil {
			if errors.Is(err, gocb.ErrDocumentNotFound) {
				context.JSON(http.StatusNotFound, cError.Errors{
					Error: "Error, Airport Document not found",
				})
			} else {
				context.JSON(http.StatusInternalServerError, cError.Errors{
					Error: "Error, Internal Server Error: " + err.Error(),
				})
			}
			return
		}
		context.JSON(http.StatusNoContent, nil)
	}
}

// @Summary      List Airport Document
// @Description  Get list of Airports. Optionally, you can filter the list by Country
// @Tags         Airport collection
// @Produce      json
// @Param        country query string true "Country<br>Example: United Kingdom, France, United States"
// @Param        limit query int false "Number of airports to return (page size)<br>Default value : 10"
// @Param        offset query int false "Number of airports to skip (for pagination)<br>Default value : 0"
// @Success      200 {object} []models.Airport
// @Failure      500 "Internal Server Error"
// @Router       /api/v1/airport/list [get]
func (ac *AirportController) GetAirports() gin.HandlerFunc {
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
		queryResult, err := ac.AirportService.QueryAirport(query)
		if err != nil {
			context.JSON(http.StatusInternalServerError, cError.Errors{
				Error: "Error, Query execution: " + err.Error(),
			})
		}
		if queryResult != nil {
			context.JSON(http.StatusOK, queryResult)
		} else {
			context.JSON(http.StatusInternalServerError, cError.Errors{
				Error: "Error, Document not found with the search query specified",
			})
		}
	}
}

// @Summary      Get Direct Connections from Airport
// @Description  Get Direct Connections from specified Airport
// @Tags         Airport collection
// @Produce      json
// @Param        airport query string true "Source airport<br>Example: SFO, LHR, CDG"
// @Param        limit query int false "Number of direct connections to return (page size)<br>Default value : 10"
// @Param        offset query int false "Number of direct connections to skip (for pagination)<br>Default value : 0"
// @Success      200 {object} []models.Destination
// @Failure      500 "Internal Server Error"
// @Router       /api/v1/airport/direct-connections [get]
func (ac *AirportController) GetDirectConnections() gin.HandlerFunc {
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
		queryResult, err := ac.AirportService.QueryDirectConnectionAirport(query)
		if err != nil {
			context.JSON(http.StatusInternalServerError, cError.Errors{
				Error: "Error, Query execution: " + err.Error(),
			})
		}
		if queryResult != nil {
			context.JSON(http.StatusOK, queryResult)
		} else {
			context.JSON(http.StatusInternalServerError, cError.Errors{
				Error: "Error, Document not found with the search query specified",
			})
		}
	}
}

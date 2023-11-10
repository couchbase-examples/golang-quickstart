package controllers

import (
	"errors"
	"fmt"
	"net/http"
	cError "src/errors"
	"src/models"
	services "src/service"
	"strconv"

	"github.com/couchbase/gocb/v2"
	"github.com/gin-gonic/gin"
)

type AirlineController struct {
	AirlineService services.IAirlineService
}

func NewAirlineController(airlineService services.IAirlineService) *AirlineController {
	return &AirlineController{
		AirlineService: airlineService,
	}
}

// @Summary      Insert Document
// @Description  Create Airline with specified ID
// @Tags         Airline collection
// @Produce      json
// @Param        id path string true "Airline ID like airline_10"
// @Param        data body models.Airline true "Data to create a new document"
// @Success      201 {object} models.Airline
// @Failure      400 "Bad Request"
// @Failure      409 "Airline document already exists"
// @Failure      500 "Internal Server Error"
// @Router       /api/v1/airline/{id} [post]
func (ac *AirlineController) InsertDocumentForAirline() gin.HandlerFunc {
	return func(context *gin.Context) {
		docKey := context.Param("id")
		data := models.Airline{}
		if err := context.ShouldBindJSON(&data); err != nil {
			context.JSON(http.StatusBadRequest, cError.Errors{
				Error: "Error, Invalid request data: " + err.Error(),
			})
			return
		}

		err := ac.AirlineService.CreateAirline(docKey, &data)
		if err != nil {
			if errors.Is(err, gocb.ErrDocumentExists) {
				context.JSON(http.StatusConflict, cError.Errors{
					Error: "Error, Airline Document already exists: " + err.Error(),
				})
			} else {
				context.JSON(http.StatusInternalServerError, cError.Errors{
					Error: "Error, Airline Document could not be inserted: " + err.Error(),
				})
			}
			return
		}
		context.JSON(http.StatusCreated, data)
	}
}

// @Summary      Get Airline Document
// @Description  Get Airline with specified ID
// @Tags         Airline collection
// @Produce      json
// @Param        id path string true "Airline ID like airline_10"
// @Success      200 {object} models.Airline
// @Failure      404 "Airline Document ID not found"
// @Failure      500 "Internal Server Error"
// @Router       /api/v1/airline/{id} [get]
func (ac *AirlineController) GetDocumentForAirline() gin.HandlerFunc {
	return func(context *gin.Context) {
		docKey := context.Param("id")
		airlineData, err := ac.AirlineService.GetAirline(docKey)
		if err != nil {
			if errors.Is(err, gocb.ErrDocumentNotFound) {
				context.JSON(http.StatusNotFound, cError.Errors{
					Error: "Error, Airline Document not found",
				})
			} else {
				context.JSON(http.StatusInternalServerError, cError.Errors{
					Error: "Error, Document could not be fetched: " + err.Error(),
				})
			}
		} else {
			context.JSON(http.StatusOK, *airlineData)
		}
	}
}

// @Summary      Update Document
// @Description  Update Airline with specified ID
// @Tags         Airline collection
// @Produce      json
// @Param       id path string true "Airline ID like airline_10"
// @Param       data body models.Airline true "Updates document"
// @Success      200 {object} models.Airline
// @Failure      400 "Bad Request"
// @Failure      500 "Internal Server Error"
// @Router       /api/v1/airline/{id} [put]
func (ac *AirlineController) UpdateDocumentForAirline() gin.HandlerFunc {
	return func(context *gin.Context) {
		docKey := context.Param("id")
		data := models.Airline{}
		if err := context.ShouldBindJSON(&data); err != nil {
			context.JSON(http.StatusBadRequest, cError.Errors{
				Error: "Error while getting the request: " + err.Error(),
			})
			return
		}
		err := ac.AirlineService.UpdateAirline(docKey, &data)
		if err != nil {
			context.JSON(http.StatusInternalServerError, cError.Errors{
				Error: "Error, Airline Document could not be updated: " + err.Error(),
			})
			return
		}
		context.JSON(http.StatusOK, data)
	}
}

// @Summary      Delete Document
// @Description  Delete Airline with specified ID
// @Tags         Airline collection
// @Produce      json
// @Param       id path string true "Airline ID like airline_10"
// @Success      204  "Airline deleted"
// @Failure      404 "Airline Document ID Not Found"
// @Failure      500 "Internal Server Error"
// @Router       /api/v1/airline/{id} [delete]
func (ac *AirlineController) DeleteDocumentForAirline() gin.HandlerFunc {
	return func(context *gin.Context) {
		docKey := context.Param("id")
		err := ac.AirlineService.DeleteAirline(docKey)
		if err != nil {
			if errors.Is(err, gocb.ErrDocumentNotFound) {
				context.JSON(http.StatusNotFound, cError.Errors{
					Error: "Error, Airline Document not found",
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

// @Summary      Get Airlines by Country
// @Description  Get list of Airlines. Optionally, you can filter the list by Country
// @Tags         Airline collection
// @Produce      json
// @Param        country query string false "Filter by country<br>Example: France, United Kingdom, United States"
// @Param        limit query int false "Number of airlines to return (page size).<br>Example: 10"
// @Param        offset query int false "Number of airlines to skip (for pagination).<br>Example: 0"
// @Success      200 {object} []models.Airline
// @Failure      500 "Internal Server Error"
// @Router       /api/v1/airline/list [get]
func (ac *AirlineController) GetAirlines() gin.HandlerFunc {
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
		var query string
		// Query to get a list of airlines filtered by country
		if country != "" {
			query = fmt.Sprintf(`
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
		} else {
			query = fmt.Sprintf(`
			SELECT airline.callsign,
				airline.country,
				airline.iata,
				airline.icao,
				airline.name
			FROM airline as airline
			ORDER BY airline.name
			LIMIT %d
			OFFSET %d;
		`, limit, offset)
		}
		queryResult, err := ac.AirlineService.QueryAirline(query)
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

// @Summary      Get Airlines Flying to Airport
// @Description  Get Airlines flying to specified destination Airport
// @Tags         Airline collection
// @Produce      json
// @Param        airport query string true "Destination airport<br>Example : SFO, JFK, LAX"
// @Param        limit query int false "Number of airlines to return (page size)<br>Default value : 10"
// @Param        offset query int false "Number of airlines to skip (for pagination)<br>Default value : 0"
// @Success      200 {object} []models.Airline
// @Failure      500 "Internal Server Error"
// @Router       /api/v1/airline/to-airport [get]
func (ac *AirlineController) GetAirlinesToAirport() gin.HandlerFunc {
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

		queryResult, err := ac.AirlineService.QueryAirline(query)
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

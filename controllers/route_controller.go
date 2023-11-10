package controllers

import (
	"errors"
	"net/http"

	cError "github.com/couchbase-examples/golang-quickstart/errors"
	"github.com/couchbase-examples/golang-quickstart/models"
	services "github.com/couchbase-examples/golang-quickstart/service"

	"github.com/couchbase/gocb/v2"
	"github.com/gin-gonic/gin"
)

type RouteController struct {
	RouteService services.IRouteService
}

func NewRouteController(routeService services.IRouteService) *RouteController {
	return &RouteController{
		RouteService: routeService,
	}
}

// @Summary      Insert Route Document
// @Description  Create Route with specified ID
// @Tags         Route collection
// @Produce      json
// @Param        id path string true "Route ID like route_10000"
// @Param        data body models.Route true "Data to create a document"
// @Success      201 {object} models.Route
// @Failure      400 "Bad Request"
// @Failure      409 "Route Document already exists"
// @Failure      500 "Internal Server Error"
// @Router       /api/v1/route/{id} [post]
func (ac *RouteController) InsertDocumentForRoute() gin.HandlerFunc {
	return func(context *gin.Context) {
		docKey := context.Param("id")
		data := models.Route{}
		if err := context.ShouldBindJSON(&data); err != nil {
			context.JSON(http.StatusBadRequest, cError.Errors{
				Error: "Error, Invalid request data: " + err.Error(),
			})
			return
		}

		err := ac.RouteService.CreateRoute(docKey, &data)
		if err != nil {
			if errors.Is(err, gocb.ErrDocumentExists) {
				context.JSON(http.StatusConflict, cError.Errors{
					Error: "Error, Route Document already exists: " + err.Error(),
				})
			} else {
				context.JSON(http.StatusInternalServerError, cError.Errors{
					Error: "Error, Route Document could not be inserted: " + err.Error(),
				})
			}
			return
		}
		context.JSON(http.StatusCreated, data)
	}
}

// @Summary      Get Route Document
// @Description  Get Route with specified ID
// @Tags         Route collection
// @Produce      json
// @Param        id path string true "Route ID like route_10000"
// @Success      200 {object} models.Route
// @Failure      404 "Route Document ID Not Found"
// @Failure      500 "Internal Server Error"
// @Router       /api/v1/route/{id} [get]
func (ac *RouteController) GetDocumentForRoute() gin.HandlerFunc {
	return func(context *gin.Context) {
		docKey := context.Param("id")
		routeData, err := ac.RouteService.GetRoute(docKey)
		if err != nil {
			if errors.Is(err, gocb.ErrDocumentNotFound) {
				context.JSON(http.StatusNotFound, cError.Errors{
					Error: "Error, Route Document not found",
				})
			} else {
				context.JSON(http.StatusInternalServerError, cError.Errors{
					Error: "Error, Document could not be fetched: " + err.Error(),
				})
			}
		} else {
			context.JSON(http.StatusOK, &routeData)
		}
	}
}

// @Summary      Update Route Document
// @Description  Update Route with specified ID
// @Tags         Route collection
// @Produce      json
// @Param        id path string true "Route ID like route_10000"
// @Param        data body models.Route true "Updates document"
// @Success      200 {object} models.Route
// @Failure      400 "Bad Request"
// @Failure      500 "Internal Server Error"
// @Router       /api/v1/route/{id} [put]
func (ac *RouteController) UpdateDocumentForRoute() gin.HandlerFunc {
	return func(context *gin.Context) {
		docKey := context.Param("id")
		data := models.Route{}
		if err := context.ShouldBindJSON(&data); err != nil {
			context.JSON(http.StatusBadRequest, cError.Errors{
				Error: "Error while getting the request: " + err.Error(),
			})
			return
		}
		err := ac.RouteService.UpdateRoute(docKey, &data)
		if err != nil {
			context.JSON(http.StatusInternalServerError, cError.Errors{
				Error: "Error, Route Document could not be updated: " + err.Error(),
			})
			return
		}
		context.JSON(http.StatusOK, data)
	}
}

// @Summary      Delete Route Document
// @Description  Delete Route with specified ID
// @Tags         Route collection
// @Produce      json
// @Param        id path string true "Route ID like route_10000"
// @Success      204 "Route Deleted"
// @Failure      404 "Route Document ID Not Found"
// @Failure      500 "Internal Server Error"
// @Router       /api/v1/route/{id} [delete]
func (ac *RouteController) DeleteDocumentForRoute() gin.HandlerFunc {
	return func(context *gin.Context) {
		docKey := context.Param("id")
		err := ac.RouteService.DeleteRoute(docKey)
		if err != nil {
			if errors.Is(err, gocb.ErrDocumentNotFound) {
				context.JSON(http.StatusNotFound, cError.Errors{
					Error: "Error, Route Document not found",
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

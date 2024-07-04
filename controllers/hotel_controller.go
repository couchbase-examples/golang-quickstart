package controllers

import (
	cError "github.com/couchbase-examples/golang-quickstart/errors"
	"github.com/couchbase-examples/golang-quickstart/models"
	services "github.com/couchbase-examples/golang-quickstart/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HotelController struct {
	HotelService services.IHotelService
}

func NewHotelController(airlineService services.IHotelService) *HotelController {
	return &HotelController{
		HotelService: airlineService,
	}
}

// SearchByName
//
// @Summary Search by hotel name
// @Description Search for hotels based on their name.
// @Tags Hotel
// @Produce json
// @Param name query string true "name search"
// @Success 200 {object} []string
// @Failure 400 "Bad Request"
// @Failure 500 "Internal Server Error"
// @Router /api/v1/hotel/autocomplete [get]
func (h *HotelController) SearchByName() gin.HandlerFunc {
	return func(context *gin.Context) {
		name := context.Query("name")
		if name == "" {
			context.JSON(http.StatusBadRequest, cError.Errors{Error: "name query parameter is required"})
			return
		}
		hotels, err := h.HotelService.SearchByName(name)
		if err != nil {
			context.JSON(http.StatusInternalServerError, cError.Errors{
				Error: err.Error(),
			})
			return
		}
		context.JSON(http.StatusOK, hotels)
	}
}

// Filter controller
//
// @Summary Fetch hotels with multiple filters
// @Description Fetch hotels using various filters such as location, rating, and price.
// @Tags Hotel
// @Produce json
// @Param data body models.HotelFilter true "Filters document"
// @Success 200 {object} []models.Hotel
// @Failure 400 "Bad Request"
// @Failure 500 "Internal Server Error"
// @Router /api/v1/hotel/fetch [post]
func (h *HotelController) Filter() gin.HandlerFunc {
	return func(context *gin.Context) {
		isNullFilter := models.HotelSearchRequest{}
		data := models.HotelSearchRequest{}
		if err := context.ShouldBindJSON(&data); err != nil {
			context.JSON(http.StatusBadRequest, cError.Errors{
				Error: "Error, Invalid request data: " + err.Error(),
			})
			return
		}
		if data == isNullFilter {
			context.JSON(http.StatusBadRequest, cError.Errors{
				Error: "Error, Invalid request data",
			})
		}
		hotels, err := h.HotelService.Filter(&data)
		if err != nil {
			context.JSON(http.StatusInternalServerError, cError.Errors{
				Error: err.Error(),
			})
			return
		}
		context.JSON(http.StatusOK, hotels)
	}
}

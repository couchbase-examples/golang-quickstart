package controllers

import (
    "net/http"
    "src/models"
    "src/responses"
    "github.com/gin-gonic/gin"
)


// @Summary      Insert Document
// @Description  Inserts a document to the "route" collection.
// @Tags         Route collection
// @Produce      json
// @Param        id path string true "Create document by specifying ID"
// @Param        data body models.RequestBodyForRoute true "Data to create a document"
// @Success      201 {object} responses.TravelSampleResponse
// @Failure      400 "Bad Request"
// @Failure      409 "Route Document already exists"
// @Failure      500 "Internal Server Error"
// @Router       /api/v1/route/{id} [post]
func InsertDocumentForRoute() gin.HandlerFunc {
	return func(context *gin.Context) {
		var data models.RequestBodyForRoute
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
		insertDocument(context, "route", docKey, data)
	}
}

// @Summary      Get Document
// @Description  Gets a document from the "route" collection based on the provided key.
// @Tags         Route collection
// @Produce      json
// @Param 		 id path string  true  "search document by id"
// @Success      200  {array}  responses.TravelSampleResponse
// @Failure 	 404			"Route Document ID Not Found"
// @Failure      500			"Internal Server Error"
// @Router       /api/v1/route/{id} [get]
func GetDocumentForRoute() gin.HandlerFunc {
	return func(context *gin.Context) {
		docKey := context.Param("id")
		var getDoc models.RequestBodyForRoute
		getDocument(context, "route", docKey, &getDoc)
	}

}

// @Summary      Update Document
// @Description  Updates a document in the "route" collection based on the provided key.
// @Tags         Route collection
// @Produce      json
// @Param 		 id path string  true  "Update document by id"
// @Param		 data body models.RequestBodyForRoute true  "Updates document"
// @Success      200  {array}  responses.TravelSampleResponse
// @Failure      400 "Bad Request"
// @Failure      500			"Internal Server Error"
// @Router       /api/v1/route/{id} [put]
func UpdateDocumentForRoute() gin.HandlerFunc {
	return func(context *gin.Context) {
		var data models.RequestBodyForRoute
		docKey := context.Param("id")
		if err := context.BindJSON(&data); err != nil {
			context.JSON(http.StatusBadRequest, responses.TravelSampleResponse{Status: http.StatusBadRequest, Message: "Error while getting the request", CollectionData: err.Error()})
			return

		}
		updateDocument(context, "route", docKey, data)
	}

}

// @Summary      Deletes Document
// @Description  Deletes a document in the "route" collection based on the provided key.
// @Tags         Route collection
// @Produce      json
// @Param 		 id  path string true  "Deletes a document with key specified"
// @Success      204  {array}  responses.TravelSampleResponse
// @Failure 	 404			"Route Document ID Not Found"
// @Failure      500			"Internal Server Error"
// @Router       /api/v1/route/{id} [delete]
func DeleteDocumentForRoute() gin.HandlerFunc {
	return func(context *gin.Context) {
		docKey := context.Param("id")
		deleteDocument(context, "route", docKey)
	}

}

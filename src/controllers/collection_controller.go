package controllers

// import (
// 	"errors"
// 	"net/http"
// 	"src/models"
// 	"src/responses"
// 	"time"
// 	"src/config"
// 	"github.com/couchbase/gocb/v2"
// 	"github.com/gin-gonic/gin"
// )


// // GetBooks		 Service Health
// // @Summary      Checks for service
// // @Description  Checks if service is running
// // @Tags         Health Check Controller
// // @Produce      json
// // @Success      200  {}  time
// // @Router       /api/v1/health [get]
// func Healthcheck() gin.HandlerFunc {
// 	return func(context *gin.Context) {

// 		context.JSON(http.StatusOK,
// 			time.Now())
// 		return
// 	}

// }

// // Insert document using KV operation
// func insertDocument(context *gin.Context, collectionName string, docKey string, data interface{}) {
// 	var getDoc interface{}

// 	_, err := config.SharedScope.Collection(collectionName).Insert(docKey, data, nil)
// 	if err != nil {
// 		if errors.Is(err, gocb.ErrDocumentExists) {
// 			context.JSON(http.StatusConflict, responses.TravelSampleResponse{
// 				Status:         http.StatusConflict,
// 				Message:        "Conflict - Document already exists",
// 				CollectionData: "Error, Document already exists: " + err.Error(),
// 			})
// 		} else {
// 			context.JSON(http.StatusInternalServerError, responses.TravelSampleResponse{
// 				Status:         http.StatusInternalServerError,
// 				Message:        "Internal Server Error",
// 				CollectionData: "Error, Document could not be inserted: " + err.Error(),
// 			})
// 		}
// 		return
// 	}

// 	// Retrieve the inserted document
// 	getResult, err := config.SharedScope.Collection(collectionName).Get(docKey, nil)
// 	if err != nil {
// 		context.JSON(http.StatusInternalServerError, responses.TravelSampleResponse{
// 			Status:         http.StatusInternalServerError,
// 			Message:        "Internal Server Error",
// 			CollectionData: "Error, Inserted Document could not be fetched: " + err.Error(),
// 		})
// 		return
// 	}

// 	err = getResult.Content(&getDoc)
// 	if err != nil {
// 		context.JSON(http.StatusInternalServerError, responses.TravelSampleResponse{
// 			Status:         http.StatusInternalServerError,
// 			Message:        "Internal Server Error",
// 			CollectionData: "Error, Failed to retrieve document content: " + err.Error(),
// 		})
// 		return
// 	}

// 	// Return a 201 response with the inserted document
// 	context.JSON(http.StatusCreated, responses.TravelSampleResponse{
// 		Status:         http.StatusCreated,
// 		Message:        "Document successfully inserted into the " + collectionName + " collection",
// 		CollectionData: getDoc,
// 	})
// }

// // Get document by key using KV operation
// func getDocument(context *gin.Context, collectionName string, docKey string, result interface{}) {
// 	getResult, err := config.SharedScope.Collection(collectionName).Get(docKey, nil)
// 	if err != nil {
// 		var statusCode int
// 		var errorMessage string

// 		if errors.Is(err, gocb.ErrDocumentNotFound) {
// 			statusCode = http.StatusNotFound
// 			errorMessage = "Document Not Found"
// 		} else {
// 			statusCode = http.StatusInternalServerError
// 			errorMessage = "Internal Server Error"
// 		}

// 		context.JSON(statusCode, responses.TravelSampleResponse{
// 			Status:         statusCode,
// 			Message:        errorMessage,
// 			CollectionData: "Error, Document could not be fetched: " + err.Error(),
// 		})
// 		return
// 	}

// 	if err := getResult.Content(result); err != nil {
// 		context.JSON(http.StatusInternalServerError, responses.TravelSampleResponse{
// 			Status:         http.StatusInternalServerError,
// 			Message:        "Internal Server Error",
// 			CollectionData: "Error, Failed to retrieve document: " + err.Error(),
// 		})
// 		return
// 	}

// 	context.JSON(http.StatusOK, responses.TravelSampleResponse{
// 		Status:         http.StatusOK,
// 		Message:        "Successfully fetched Document",
// 		CollectionData: result,
// 	})
// }

// // Upsert document using KV operation
// func updateDocument(context *gin.Context, collectionName string, docKey string, data interface{}) {
// 	var getDoc interface{}
// 	_, err := config.SharedScope.Collection(collectionName).Upsert(docKey, data, nil)
// 	if err != nil {
// 		context.JSON(http.StatusInternalServerError, responses.TravelSampleResponse{
// 			Status:         http.StatusInternalServerError,
// 			Message:        "Internal Server Error",
// 			CollectionData: "Error, Document could not be updated: " + err.Error(),
// 		})
// 		return
// 	}

// 	// Retrieve the updated document
// 	getResult, err := config.SharedScope.Collection(collectionName).Get(docKey, nil)
// 	if err != nil {
// 		context.JSON(http.StatusInternalServerError, responses.TravelSampleResponse{
// 			Status:         http.StatusInternalServerError,
// 			Message:        "Internal Server Error",
// 			CollectionData: "Error, Updated Document could not be fetched: " + err.Error(),
// 		})
// 		return
// 	}

// 	err = getResult.Content(&getDoc)
// 	context.JSON(http.StatusOK, responses.TravelSampleResponse{
// 		Status:         http.StatusOK,
// 		Message:        "Successfully Updated the document",
// 		CollectionData: getDoc,
// 	})
// }

// // Delete document using KV operation
// func deleteDocument(context *gin.Context, collectionName string, docKey string) {
// 	_, err := config.SharedScope.Collection(collectionName).Remove(docKey, nil)
// 	if err != nil {
// 		if errors.Is(err, gocb.ErrDocumentNotFound) {
// 			context.JSON(http.StatusNotFound, responses.TravelSampleResponse{
// 				Status:         http.StatusNotFound,
// 				Message:        "Document Not Found",
// 				CollectionData: err.Error(),
// 			})
// 		} else {
// 			context.JSON(http.StatusInternalServerError, responses.TravelSampleResponse{
// 				Status:         http.StatusInternalServerError,
// 				Message:        "Internal Server Error",
// 				CollectionData: "Error, Document could not be deleted: " + err.Error(),
// 			})
// 		}
// 		return
// 	}

// 	context.JSON(http.StatusNoContent, responses.TravelSampleResponse{
// 		Status:         http.StatusNoContent,
// 		Message:        "Successfully deleted the document",
// 		CollectionData: docKey,
// 	})
// }

// // @Summary      Insert Document
// // @Description  Inserts a document into the "airline" collection.
// // @Tags         Airline collection
// // @Produce      json
// // @Param        id path string true "Create document by specifying ID"
// // @Param        data body models.RequestBodyForAirline true "Data to create a new document"
// // @Success      201 {object} responses.TravelSampleResponse
// // @Failure      400 "Bad Request"
// // @Failure      409 "Airline document already exists"
// // @Failure      500 "Internal Server Error"
// // @Router       /api/v1/airline/{id} [post]
// func InsertDocumentForAirline() gin.HandlerFunc {
// 	return func(context *gin.Context) {
// 		var data models.RequestBodyForAirline
// 		docKey := context.Param("id")
// 		// Bind the JSON data to the "data" variable
// 		if err := context.BindJSON(&data); err != nil {
// 			context.JSON(http.StatusBadRequest, responses.TravelSampleResponse{
// 				Status:         http.StatusBadRequest,
// 				Message:        "Bad Request",
// 				CollectionData: "Error, Invalid request data: " + err.Error(),
// 			})
// 			return
// 		}
// 		insertDocument(context, "airline", docKey, data)
// 	}
// }

// // @Summary      Get Airline Document
// // @Description  Gets a document from the "airline" collection based on the provided key.
// // @Tags         Airline collection
// // @Produce      json
// // @Param        id path string true "Search document by ID"
// // @Success      200 {object} responses.TravelSampleResponse
// // @Failure      404 "Airline Document ID not found"
// // @Failure      500 "Internal Server Error"
// // @Router       /api/v1/airline/{id} [get]
// func GetDocumentForAirline() gin.HandlerFunc {
// 	return func(context *gin.Context) {
// 		documentID := context.Param("id")
// 		var getDoc models.RequestBodyForAirline
// 		getDocument(context, "airline", documentID, &getDoc)
// 	}
// }

// // @Summary      Update Document
// // @Description  Updates a document in the "airline" collection based on the provided key.
// // @Tags         Airline collection
// // @Produce      json
// // @Param 		 id path string  true  "Update document by ID"
// // @Param		 data body models.RequestBodyForAirline true  "Updates document"
// // @Success      200  {array}  responses.TravelSampleResponse
// // @Failure      400 "Bad Request"
// // @Failure      500			"Internal Server Error"
// // @Router       /api/v1/airline/{id} [put]
// func UpdateDocumentForAirline() gin.HandlerFunc {
// 	return func(context *gin.Context) {
// 		var data models.RequestBodyForAirline
// 		docKey := context.Param("id")
// 		if err := context.BindJSON(&data); err != nil {
// 			context.JSON(http.StatusBadRequest, responses.TravelSampleResponse{Status: http.StatusBadRequest, Message: "Error while getting the request", CollectionData: err.Error()})
// 			return

// 		}
// 		updateDocument(context, "airline", docKey, data)

// 	}

// }

// // @Summary      Delete Document
// // @Description  Deletes a document in the "airline" collection based on the provided key.
// // @Tags         Airline collection
// // @Produce      json
// // @Param        id path string true "Deletes a document with the specified key"
// // @Success      204 {object} responses.TravelSampleResponse
// // @Failure      404 "Airline Document ID Not Found"
// // @Failure      500 "Internal Server Error"
// // @Router       /api/v1/airline/{id} [delete]
// func DeleteDocumentForAirline() gin.HandlerFunc {
// 	return func(context *gin.Context) {
// 		docKey := context.Param("id")
// 		deleteDocument(context, "airline", docKey)

// 	}
// }

// // @Summary      Insert Document
// // @Description  Inserts a document to the "airport" collection.
// // @Tags         Airport collection
// // @Produce      json
// // @Param        id path string true "Create document by specifying ID"
// // @Param        data body models.RequestBodyForAirport true "Data to create a document"
// // @Success      201 {object} responses.TravelSampleResponse
// // @Failure      400 "Bad Request"
// // @Failure      409 "Airport Document already exists"
// // @Failure      500 "Internal Server Error"
// // @Router       /api/v1/airport/{id} [post]
// func InsertDocumentForAirport() gin.HandlerFunc {
// 	return func(context *gin.Context) {
// 		var data models.RequestBodyForAirport
// 		docKey := context.Param("id")
// 		// Bind the JSON data to the "data" variable
// 		if err := context.BindJSON(&data); err != nil {
// 			context.JSON(http.StatusBadRequest, responses.TravelSampleResponse{
// 				Status:         http.StatusBadRequest,
// 				Message:        "Bad Request",
// 				CollectionData: "Error, Invalid request data: " + err.Error(),
// 			})
// 			return
// 		}
// 		insertDocument(context, "airport", docKey, data)
// 	}
// }

// // @Summary      Get Document
// // @Description  Gets a document from the "airport" collection based on the provided key.
// // @Tags         Airport collection
// // @Produce      json
// // @Param        id path string true "Search document by ID"
// // @Success      200 {object} responses.TravelSampleResponse
// // @Failure      404 "Airport Document ID Not Found"
// // @Failure      500 "Internal Server Error"
// // @Router       /api/v1/airport/{id} [get]
// func GetDocumentForAirport() gin.HandlerFunc {
// 	return func(context *gin.Context) {
// 		docKey := context.Param("id")
// 		var getDoc models.RequestBodyForAirport
// 		getDocument(context, "airport", docKey, &getDoc)
// 	}
// }

// // @Summary      Update Document
// // @Description  Updates a document in the "airport" collection based on the provided key.
// // @Tags         Airport collection
// // @Produce      json
// // @Param 		 id path string  true  "Update document by id"
// // @Param		 data body models.RequestBodyForAirport true  "Updates document"
// // @Success      200  {array}  responses.TravelSampleResponse
// // @Failure      400 "Bad Request"
// // @Failure      500			"Internal Server Error"
// // @Router       /api/v1/airport/{id} [put]
// func UpdateDocumentForAirport() gin.HandlerFunc {
// 	return func(context *gin.Context) {
// 		var data models.RequestBodyForAirport
// 		docKey := context.Param("id")
// 		if err := context.BindJSON(&data); err != nil {
// 			context.JSON(http.StatusBadRequest, responses.TravelSampleResponse{Status: http.StatusBadRequest, Message: "Error while getting the request", CollectionData: err.Error()})
// 			return

// 		}
// 		updateDocument(context, "airport", docKey, data)
// 	}

// }

// // @Summary      Deletes Document
// // @Description  Deletes a document in the "airport" collection based on the provided key.
// // @Tags         Airport collection
// // @Produce      json
// // @Param 		 id  path string true  "Deletes a document with key specified"
// // @Success      204  {array}  responses.TravelSampleResponse
// // @Failure 	 404			"Airport Document ID Not Found"
// // @Failure      500			"Internal Server Error"
// // @Router       /api/v1/airport/{id} [delete]
// func DeleteDocumentForAirport() gin.HandlerFunc {
// 	return func(context *gin.Context) {
// 		docKey := context.Param("id")
// 		deleteDocument(context, "airport", docKey)
// 	}

// }

// // @Summary      Insert Document
// // @Description  Inserts a document to the "route" collection.
// // @Tags         Route collection
// // @Produce      json
// // @Param        id path string true "Create document by specifying ID"
// // @Param        data body models.RequestBodyForRoute true "Data to create a document"
// // @Success      201 {object} responses.TravelSampleResponse
// // @Failure      400 "Bad Request"
// // @Failure      409 "Route Document already exists"
// // @Failure      500 "Internal Server Error"
// // @Router       /api/v1/route/{id} [post]
// func InsertDocumentForRoute() gin.HandlerFunc {
// 	return func(context *gin.Context) {
// 		var data models.RequestBodyForRoute
// 		docKey := context.Param("id")
// 		// Bind the JSON data to the "data" variable
// 		if err := context.BindJSON(&data); err != nil {
// 			context.JSON(http.StatusBadRequest, responses.TravelSampleResponse{
// 				Status:         http.StatusBadRequest,
// 				Message:        "Bad Request",
// 				CollectionData: "Error, Invalid request data: " + err.Error(),
// 			})
// 			return
// 		}
// 		insertDocument(context, "route", docKey, data)
// 	}
// }

// // @Summary      Get Document
// // @Description  Gets a document from the "route" collection based on the provided key.
// // @Tags         Route collection
// // @Produce      json
// // @Param 		 id path string  true  "search document by id"
// // @Success      200  {array}  responses.TravelSampleResponse
// // @Failure 	 404			"Route Document ID Not Found"
// // @Failure      500			"Internal Server Error"
// // @Router       /api/v1/route/{id} [get]
// func GetDocumentForRoute() gin.HandlerFunc {
// 	return func(context *gin.Context) {
// 		docKey := context.Param("id")
// 		var getDoc models.RequestBodyForRoute
// 		getDocument(context, "route", docKey, &getDoc)
// 	}

// }

// // @Summary      Update Document
// // @Description  Updates a document in the "route" collection based on the provided key.
// // @Tags         Route collection
// // @Produce      json
// // @Param 		 id path string  true  "Update document by id"
// // @Param		 data body models.RequestBodyForRoute true  "Updates document"
// // @Success      200  {array}  responses.TravelSampleResponse
// // @Failure      400 "Bad Request"
// // @Failure      500			"Internal Server Error"
// // @Router       /api/v1/route/{id} [put]
// func UpdateDocumentForRoute() gin.HandlerFunc {
// 	return func(context *gin.Context) {
// 		var data models.RequestBodyForRoute
// 		docKey := context.Param("id")
// 		if err := context.BindJSON(&data); err != nil {
// 			context.JSON(http.StatusBadRequest, responses.TravelSampleResponse{Status: http.StatusBadRequest, Message: "Error while getting the request", CollectionData: err.Error()})
// 			return

// 		}
// 		updateDocument(context, "route", docKey, data)
// 	}

// }

// // @Summary      Deletes Document
// // @Description  Deletes a document in the "route" collection based on the provided key.
// // @Tags         Route collection
// // @Produce      json
// // @Param 		 id  path string true  "Deletes a document with key specified"
// // @Success      204  {array}  responses.TravelSampleResponse
// // @Failure 	 404			"Route Document ID Not Found"
// // @Failure      500			"Internal Server Error"
// // @Router       /api/v1/route/{id} [delete]
// func DeleteDocumentForRoute() gin.HandlerFunc {
// 	return func(context *gin.Context) {
// 		docKey := context.Param("id")
// 		deleteDocument(context, "route", docKey)
// 	}

// }

// // // @Summary      Searches the Document by word
// // // @Description  Searches the Document by word
// // // @Tags         Profile Controller
// // // @Produce      json
// // // @Param 		 search  query string  true  "search document by word"
// // // @Param 		 limit  query string  false  "specify limit"
// // // @Param 		 skip  query string  false  "skip document"
// // // @Success      200  {array}  responses.ProfileResponse
// // // @Failure      500			"Error while getting examples"
// // // @Failure      403			"Forbidden"
// // // @Failure 	 404			"Page Not found"
// // // @Router       /api/v1/profile/profiles [get]
// // func SearchProfile() gin.HandlerFunc {
// // 	return func(context *gin.Context) {
// // 		search := context.Query("search")
// // 		limit := context.Query("limit")
// // 		skip := context.Query("skip")
// // 		if limit == "" {
// // 			limit = "5"
// // 		}
// // 		if skip == "" {
// // 			skip = "0"
// // 		}
// 		var search_query = search
// 		//var query2 = "SELECT p.* FROM " + bucket_name + "." + scope_name + "." + collection_name + " p WHERE lower(p.FirstName) LIKE " + "'" + search_query + "%'" + " OR lower(p.LastName) LIKE " + "'" + search_query + "%'" + " LIMIT " + limit + " OFFSET " + skip + ";"
// 		query := fmt.Sprintf("SELECT p.* FROM %s.%s.%s p WHERE lower(p.FirstName) LIKE '%s' OR lower(p.LastName) LIKE '%s' LIMIT %s OFFSET %s%s", bucket_name, scope_name, collection_name, search_query, search_query, limit, skip, ";")
// 		results, err := cluster.Query(query, nil)
// 		var s interface{}
// 		var profiles []interface{}
// 		if results == nil {
// 			panic(err)
// 		}
// 		for results.Next() {
// 			err := results.Row(&s)
// 			if err != nil {
// 				panic(err)
// 			}
// 			profiles = append(profiles, s)

// 		}
// 		if s != nil {
// 			context.JSON(http.StatusOK, responses.ProfileResponse{Status: http.StatusOK, Message: "Successfully found the  the document", Profile: profiles})
// 		} else {
// 			context.JSON(http.StatusInternalServerError, responses.ProfileResponse{Status: http.StatusInternalServerError, Message: "Error, Document Not found with the search query specified", Profile: profiles})
// 		}

// 	}

// // }

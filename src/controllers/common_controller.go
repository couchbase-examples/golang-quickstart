package controllers

import (
	"errors"
	"net/http"
	"src/config"
	"src/responses"

	"github.com/couchbase/gocb/v2"
	"github.com/gin-gonic/gin"
)

// Insert document using KV operation
func insertDocument(context *gin.Context, collectionName string, docKey string, data interface{}) {
	var getDoc interface{}

	_, err := config.SharedScope.Collection(collectionName).Insert(docKey, data, nil)
	if err != nil {
		if errors.Is(err, gocb.ErrDocumentExists) {
			context.JSON(http.StatusConflict, responses.TravelSampleResponse{
				Status:         http.StatusConflict,
				Message:        "Conflict - Document already exists",
				CollectionData: "Error, Document already exists: " + err.Error(),
			})
		} else {
			context.JSON(http.StatusInternalServerError, responses.TravelSampleResponse{
				Status:         http.StatusInternalServerError,
				Message:        "Internal Server Error",
				CollectionData: "Error, Document could not be inserted: " + err.Error(),
			})
		}
		return
	}

	// Retrieve the inserted document
	getResult, err := config.SharedScope.Collection(collectionName).Get(docKey, nil)
	if err != nil {
		context.JSON(http.StatusInternalServerError, responses.TravelSampleResponse{
			Status:         http.StatusInternalServerError,
			Message:        "Internal Server Error",
			CollectionData: "Error, Inserted Document could not be fetched: " + err.Error(),
		})
		return
	}

	err = getResult.Content(&getDoc)
	if err != nil {
		context.JSON(http.StatusInternalServerError, responses.TravelSampleResponse{
			Status:         http.StatusInternalServerError,
			Message:        "Internal Server Error",
			CollectionData: "Error, Failed to retrieve document content: " + err.Error(),
		})
		return
	}

	// Return a 201 response with the inserted document
	context.JSON(http.StatusCreated, responses.TravelSampleResponse{
		Status:         http.StatusCreated,
		Message:        "Document successfully inserted into the " + collectionName + " collection",
		CollectionData: getDoc,
	})
}

// Get document by key using KV operation
func getDocument(context *gin.Context, collectionName string, docKey string, result interface{}) {
	getResult, err := config.SharedScope.Collection(collectionName).Get(docKey, nil)
	if err != nil {
		var statusCode int
		var errorMessage string

		if errors.Is(err, gocb.ErrDocumentNotFound) {
			statusCode = http.StatusNotFound
			errorMessage = "Document Not Found"
		} else {
			statusCode = http.StatusInternalServerError
			errorMessage = "Internal Server Error"
		}

		context.JSON(statusCode, responses.TravelSampleResponse{
			Status:         statusCode,
			Message:        errorMessage,
			CollectionData: "Error, Document could not be fetched: " + err.Error(),
		})
		return
	}

	if err := getResult.Content(result); err != nil {
		context.JSON(http.StatusInternalServerError, responses.TravelSampleResponse{
			Status:         http.StatusInternalServerError,
			Message:        "Internal Server Error",
			CollectionData: "Error, Failed to retrieve document: " + err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, responses.TravelSampleResponse{
		Status:         http.StatusOK,
		Message:        "Successfully fetched Document",
		CollectionData: result,
	})
}

// Upsert document using KV operation
func updateDocument(context *gin.Context, collectionName string, docKey string, data interface{}) {
	var getDoc interface{}
	_, err := config.SharedScope.Collection(collectionName).Upsert(docKey, data, nil)
	if err != nil {
		context.JSON(http.StatusInternalServerError, responses.TravelSampleResponse{
			Status:         http.StatusInternalServerError,
			Message:        "Internal Server Error",
			CollectionData: "Error, Document could not be updated: " + err.Error(),
		})
		return
	}

	// Retrieve the updated document
	getResult, err := config.SharedScope.Collection(collectionName).Get(docKey, nil)
	if err != nil {
		context.JSON(http.StatusInternalServerError, responses.TravelSampleResponse{
			Status:         http.StatusInternalServerError,
			Message:        "Internal Server Error",
			CollectionData: "Error, Updated Document could not be fetched: " + err.Error(),
		})
		return
	}

	err = getResult.Content(&getDoc)
	context.JSON(http.StatusOK, responses.TravelSampleResponse{
		Status:         http.StatusOK,
		Message:        "Successfully Updated the document",
		CollectionData: getDoc,
	})
}

// Delete document using KV operation
func deleteDocument(context *gin.Context, collectionName string, docKey string) {
	_, err := config.SharedScope.Collection(collectionName).Remove(docKey, nil)
	if err != nil {
		if errors.Is(err, gocb.ErrDocumentNotFound) {
			context.JSON(http.StatusNotFound, responses.TravelSampleResponse{
				Status:         http.StatusNotFound,
				Message:        "Document Not Found",
				CollectionData: err.Error(),
			})
		} else {
			context.JSON(http.StatusInternalServerError, responses.TravelSampleResponse{
				Status:         http.StatusInternalServerError,
				Message:        "Internal Server Error",
				CollectionData: "Error, Document could not be deleted: " + err.Error(),
			})
		}
		return
	}

	context.JSON(http.StatusNoContent, responses.TravelSampleResponse{
		Status:         http.StatusNoContent,
		Message:        "Successfully deleted the document",
		CollectionData: docKey,
	})
}

// GetDocumentsFromQuery executes a query for a given collection and returns the results.
func GetDocumentsFromQuery(context *gin.Context, collectionName string, query string, result interface{}) {
	queryResult, err := config.SharedScope.Query(query, nil)
	if err != nil {
		context.JSON(http.StatusInternalServerError, responses.TravelSampleResponse{
			Status:         http.StatusInternalServerError,
			Message:        "Internal Server Error",
			CollectionData: "Error, Query execution: " + err.Error(),
		})
		return
	}
	var document interface{}
	var documents []interface{}

	if queryResult == nil {
		panic(err)
	}

	for queryResult.Next() {
		err := queryResult.Row(&document)
		if err != nil {
			panic(err)
		}
		documents = append(documents, document)
	}

	if document != nil {
		context.JSON(http.StatusOK, responses.TravelSampleResponse{
			Status:         http.StatusOK,
			Message:        "Successfully found the document",
			CollectionData: documents,
		})
	} else {
		context.JSON(http.StatusInternalServerError, responses.TravelSampleResponse{
			Status:         http.StatusInternalServerError,
			Message:        "Error, Document not found with the search query specified",
			CollectionData: documents,
		})
	}

}

package controllers

import (
	"net/http"
	"src/db"
	"src/models"
	"src/responses"
	"errors"
	"github.com/couchbase/gocb/v2"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
	"fmt"
)

var cluster *gocb.Cluster = db.Initialize_db()
var col *gocb.Collection = db.GetCollection(cluster)
var bucket_name = db.Bucket_name
var scope_name = db.Scope_name
var collection_name = db.Collection_name

// GetBooks		 Service Health
// @Summary      Checks for service
// @Description  Checks if service is running
// @Tags         Health Check Controller
// @Produce      json
// @Success      200  {}  time
// @Router       /api/v1/health [get]
func Healthcheck() gin.HandlerFunc {
	return func(context *gin.Context) {

		context.JSON(http.StatusOK,
			time.Now())
		return
	}

}

// @Summary      Create Document
// @Description  Creates the Document with key
// @Tags         Profile Controller
// @Produce      json
// @Param 		 data body models.Profile true  "Creates a document"
// @Success      200  {array} responses.ProfileResponse
// @Failure      500			"Error while getting examples"
// @Failure      403			"Forbidden"
// @Failure 	 404			"Page Not found"
// @Router       /api/v1/profile [post]
func InsertProfile() gin.HandlerFunc {
	return func(context *gin.Context) {
		var data models.RequestBody
		var getDoc models.RequestBody
		if err := context.BindJSON(&data); err != nil {
			// DO SOMETHING WITH THE ERROR
			//Error while getting request
			context.JSON(http.StatusBadRequest, responses.ProfileResponse{Status: http.StatusBadRequest, Message: "Error while getting the request", Profile: err.Error()})
			return
		}

		key := uuid.New().String()
		//Hash the password here
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
		if err != nil {
			panic(err)
		}
		new_profile := models.RequestBody{
			Pid:       key,
			FirstName: data.FirstName,
			LastName:  data.LastName,
			Email:     data.Email,
			Password:  string(hashedPassword),
		}
		profile_key := new_profile.Pid
		//perform insert operation
		result, err := col.Insert(profile_key, new_profile, nil)
		_ =result
		getResult, err := col.Get(profile_key, nil)
		err = getResult.Content(&getDoc)
		context.JSON(http.StatusOK, responses.ProfileResponse{Status: http.StatusOK, Message: "Document successfully inserted", Profile: getDoc})

	}
}

// @Summary      Get Document
// @Description  Gets the Document with key
// @Tags         Profile Controller
// @Produce      json
// @Param 		 id path string  true  "search document by id"
// @Success      200  {array}  responses.ProfileResponse
// @Failure      500			"Error while getting examples"
// @Failure      403			"Forbidden"
// @Failure 	 404			"Page Not found"
// @Router       /api/v1/profile/{id} [get]
func GetProfile() gin.HandlerFunc {
	return func(context *gin.Context) {
		Pid := context.Param("id")
		var getDoc models.RequestBody
		getResult, err := col.Get(Pid, nil)
		if errors.Is(err, gocb.ErrDocumentNotFound) {
			context.JSON(http.StatusNotFound, responses.ProfileResponse{Status: http.StatusNotFound, Message: "Error, Document does not exist", Profile: err})
			return
		}
		err = getResult.Content(&getDoc)
		if err != nil {
			panic(err)
		}

		context.JSON(http.StatusOK, responses.ProfileResponse{Status: http.StatusOK, Message: "Successfully fetched Document", Profile: getDoc})

	}

}

// @Summary      Update Document
// @Description  Updates the Document with key
// @Tags         Profile Controller
// @Produce      json
// @Param 		 id path string  true  "Update document by id"
// @Param		 data body models.Profile true  "Updates document"
// @Success      200  {array}  responses.ProfileResponse
// @Failure      500			"Error while getting examples"
// @Failure      403			"Forbidden"
// @Failure 	 404			"Page Not found"
// @Router       /api/v1/profile/{id} [put]
func UpdateProfile() gin.HandlerFunc {
	return func(context *gin.Context) {
		var data models.RequestBody
		profile_id := context.Param("id")
		if err := context.BindJSON(&data); err != nil {
			context.JSON(http.StatusBadRequest, responses.ProfileResponse{Status: http.StatusBadRequest, Message: "Error while getting the request", Profile: err.Error()})
			return

		}
		var getDoc models.RequestBody
		data.Pid = profile_id
		hashedPassword, err_password := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
		if err_password != nil {
			panic(err_password)
		}
		data.Password = string(hashedPassword)

		_, err := col.Upsert(profile_id, data, nil)
		if errors.Is(err, gocb.ErrDocumentNotFound) {
			context.JSON(http.StatusNotFound, responses.ProfileResponse{Status: http.StatusNotFound, Message: "Error, Document with the key does not exist", Profile: err.Error()})
			return

		}
		getResult, err := col.Get(profile_id, nil)
		err = getResult.Content(&getDoc)
		context.JSON(http.StatusOK, responses.ProfileResponse{Status: http.StatusOK, Message: "Successfully Updated the document", Profile: getDoc})
	}

}

// @Summary      Deletes Document
// @Description  Deletes the Document with key
// @Tags         Profile Controller
// @Produce      json
// @Param 		 id  path string true  "Deletes a document with key specified"
// @Success      200  {array}  responses.ProfileResponse
// @Failure      500			"Error while getting examples"
// @Failure      403			"Forbidden"
// @Failure 	 404			"Page Not found"
// @Router       /api/v1/profile/{id} [delete]
func DeleteProfile() gin.HandlerFunc {
	return func(context *gin.Context) {
		profile_id := context.Param("id")
		result, err := col.Remove(profile_id, nil)
		if errors.Is(err, gocb.ErrDocumentNotFound) {
			context.JSON(http.StatusInternalServerError, responses.ProfileResponse{Status: http.StatusInternalServerError, Message: "Error, Document Not found with the specified key", Profile: err.Error()})
			return
		}

		context.JSON(http.StatusOK, responses.ProfileResponse{Status: http.StatusOK, Message: "Successfully deleted the document", Profile: result})

	}

}

// @Summary      Searches the Document by word
// @Description  Searches the Document by word
// @Tags         Profile Controller
// @Produce      json
// @Param 		 search  query string  true  "search document by word"
// @Param 		 limit  query string  false  "specify limit"
// @Param 		 skip  query string  false  "skip document"
// @Success      200  {array}  responses.ProfileResponse
// @Failure      500			"Error while getting examples"
// @Failure      403			"Forbidden"
// @Failure 	 404			"Page Not found"
// @Router       /api/v1/profile/profiles [get]
func SearchProfile() gin.HandlerFunc {
	return func(context *gin.Context) {
		search := context.Query("search")
		limit := context.Query("limit")
		skip := context.Query("skip")
		if limit == "" {
			limit = "5"
		}
		if skip == "" {
			skip = "0"
		}
		var search_query = search
		//var query2 = "SELECT p.* FROM " + bucket_name + "." + scope_name + "." + collection_name + " p WHERE lower(p.FirstName) LIKE " + "'" + search_query + "%'" + " OR lower(p.LastName) LIKE " + "'" + search_query + "%'" + " LIMIT " + limit + " OFFSET " + skip + ";"
		query:= fmt.Sprintf("SELECT p.* FROM %s.%s.%s p WHERE lower(p.FirstName) LIKE '%s' OR lower(p.LastName) LIKE '%s' LIMIT %s OFFSET %s%s",bucket_name,scope_name,collection_name,search_query,search_query,limit,skip,";")
		results, err := cluster.Query(query, nil)
		var s interface{}
		var profiles []interface{}
		if results==nil{
			panic(err)
		}
		for results.Next() {
			err := results.Row(&s)
			if err != nil {
				panic(err)
			}
			profiles = append(profiles, s)

		}
		if s != nil {
			context.JSON(http.StatusOK, responses.ProfileResponse{Status: http.StatusOK, Message: "Successfully found the  the document", Profile: profiles})
		} else {
			context.JSON(http.StatusInternalServerError, responses.ProfileResponse{Status: http.StatusInternalServerError, Message: "Error, Document Not found with the search query specified", Profile: profiles})
		}

	}

}

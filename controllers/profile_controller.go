package controllers

import (
	"net/http"
	"app/db"
	"app/models"
	"app/responses"
	"errors"
	"fmt"
	"github.com/couchbase/gocb/v2"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
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
// @Success      200  {array}  time
// @Router       /api/v1/health [get]
func Healthcheck() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.JSON(http.StatusOK,
			time.Now())
		return
	}

}

// @Summary      Create Document
// @Description  Creates the Document with key
// @Tags         Profile Controller
// @Produce      json
// @Param 		 data body models.Profile true  "Creates a document"
// @Success      200  {array}  responses.ProfileResponse
// @Failure      500			"Error while getting examples"
// @Failure      403			"Forbidden"
// @Failure 	 404			"Page Not found"
// @Router       /api/v1/profile [post]
func InsertProfile() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data models.RequestBody
		var getDoc models.RequestBody
		if err := c.BindJSON(&data); err != nil {
			// DO SOMETHING WITH THE ERROR
			//Error while getting request
			c.JSON(http.StatusBadRequest, responses.ProfileResponse{Status: http.StatusBadRequest, Message: "Error while getting the request", Data: map[string]interface{}{"profile": err.Error()}})
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
		fmt.Println(result)
		//fmt.Println(new_profile.LastName)
		getResult, err := col.Get(profile_key, nil)
		err = getResult.Content(&getDoc)
		c.JSON(http.StatusOK, responses.ProfileResponse{Status: http.StatusOK, Message: "Document successfully inserted", Data: map[string]interface{}{"profile": getDoc}})

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
	return func(c *gin.Context) {
		Pid := c.Param("id")
		var getDoc models.RequestBody
		getResult, err := col.Get(Pid, nil)
		if errors.Is(err, gocb.ErrDocumentNotFound) {
			c.JSON(http.StatusInternalServerError, responses.ProfileResponse{Status: http.StatusInternalServerError, Message: "Error, Document does not exist", Data: map[string]interface{}{"profile": err}})
			return
		}
		err = getResult.Content(&getDoc)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, responses.ProfileResponse{Status: http.StatusOK, Message: "Successfully fetched Document", Data: map[string]interface{}{"profile": getDoc}})

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
	return func(c *gin.Context) {
		var data models.RequestBody
		profile_id := c.Param("id")
		if err := c.BindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, responses.ProfileResponse{Status: http.StatusBadRequest, Message: "Error while getting the request", Data: map[string]interface{}{"profile": err.Error()}})
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
		fmt.Println()
		if errors.Is(err, gocb.ErrDocumentNotFound) {
			c.JSON(http.StatusInternalServerError, responses.ProfileResponse{Status: http.StatusInternalServerError, Message: "Error, Document with the key does not exist", Data: map[string]interface{}{"profile": err.Error()}})
			return

		}
		getResult, err := col.Get(profile_id, nil)
		err = getResult.Content(&getDoc)
		c.JSON(http.StatusOK, responses.ProfileResponse{Status: http.StatusOK, Message: "Successfully Updated the document", Data: map[string]interface{}{"profile": getDoc}})
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
	return func(c *gin.Context) {
		profile_id := c.Param("id")
		result, err := col.Remove(profile_id, nil)
		if errors.Is(err, gocb.ErrDocumentNotFound) {
			c.JSON(http.StatusInternalServerError, responses.ProfileResponse{Status: http.StatusInternalServerError, Message: "Error, Document Not found with the specified key", Data: map[string]interface{}{"profile": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.ProfileResponse{Status: http.StatusOK, Message: "Successfully deleted the document", Data: map[string]interface{}{"profile": result}})

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
	return func(c *gin.Context) {
		search := c.Query("search")
		limit := c.Query("limit")
		skip := c.Query("skip")
		//fmt.Println(search)
		if limit == "" {
			limit = "5"
		}
		if skip == "" {
			skip = "0"
		}
		var search_query = search
		var query = "SELECT p.* FROM " + bucket_name + "." + scope_name + "." + collection_name + " p WHERE lower(p.FirstName) LIKE " + "'" + search_query + "%'" + " OR lower(p.LastName) LIKE " + "'" + search_query + "%'" + " LIMIT " + limit + " OFFSET " + skip + ";"
		fmt.Println(query)
		//var query ="SELECT p.* FROM user_profile.default.default p WHERE lower(p.FirstName) LIKE 'qw' OR lower(p.LastName) LIKE 'er'"
		//fmt.Println(query)
		results, _ := cluster.Query(query, nil)
		//fmt.Println(results)
		var profile interface{}
		var s []interface{}
		for results.Next() {
			err := results.Row(&profile)
			if err != nil {
				panic(err)
			}
			s = append(s, profile)
			//fmt.Println(s)

		}
		if profile != nil {
			c.JSON(http.StatusOK, responses.ProfileResponse{Status: http.StatusOK, Message: "Successfully found the  the document", Data: map[string]interface{}{"data": s}})
		} else {
			c.JSON(http.StatusInternalServerError, responses.ProfileResponse{Status: http.StatusInternalServerError, Message: "Error, Document Not found with the search query specified", Data: map[string]interface{}{"profile": s}})
		}

	}

}
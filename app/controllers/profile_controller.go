package controllers


import(
	"net/http"

	"app/models"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/couchbase/gocb/v2"
	"errors"
	"github.com/google/uuid"
	"app/db"
	"app/responses"
)
var cluster *gocb.Cluster=db.Initialize_db()

var col *gocb.Collection=db.GetCollection(cluster)


func Healthcheck() gin.HandlerFunc {
	return func(c *gin.Context){

		c.JSON(http.StatusOK,
			time.Now())
		return 
	}





}


func InsertProfile() gin.HandlerFunc{
	return func(c *gin.Context){
		var data models.RequestBody
		if err := c.BindJSON(&data); err != nil {
			// DO SOMETHING WITH THE ERROR
			//Error while getting request
			c.JSON(http.StatusBadRequest,responses.ProfileResponse{Status: http.StatusBadRequest, Message: "Error while getting the request", Data: map[string] interface{}{"data":err.Error()}})
			return
			}
			check_key:=data.Pid
			_,err:=col.Get(check_key,nil)

			if err==nil{
				c.JSON(http.StatusInternalServerError,responses.ProfileResponse{Status: http.StatusInternalServerError, Message: "Error", Data: map[string] interface{}{"error":"Document Already exists"}})
				return
			}
			key:=uuid.New().String()
			new_profile:=models.RequestBody{
				Pid: key,
				FirstName: data.FirstName,
				LastName: data.LastName,
				Email: data.Email,
				Password: data.Password,

			}
			//Hash the password here
			profile_key:=new_profile.Pid
			//perform insert operation
			res, err:=col.Insert(profile_key,new_profile,nil)
			_ =res
			c.JSON(http.StatusOK,responses.ProfileResponse{Status: http.StatusOK, Message: "Document successfully inserted", Data: map[string] interface{}{"pid":profile_key}})


	}
}


func GetProfile() gin.HandlerFunc{
	return func(c *gin.Context){
		profile_id:=c.Param("id")
		var getDoc models.RequestBody
		getResult,err:=col.Get(profile_id,nil)
		if errors.Is(err, gocb.ErrDocumentNotFound) {
			c.JSON(http.StatusInternalServerError,responses.ProfileResponse{Status: http.StatusInternalServerError, Message: "Error, Document does not exist", Data: map[string] interface{}{"data":err}})
			return 
		}
		err = getResult.Content(&getDoc)
		if err!=nil{
			panic(err)
		}
		

		c.JSON(http.StatusOK,responses.ProfileResponse{Status: http.StatusOK, Message: "Successfully fetched Document", Data: map[string] interface{}{"data":getDoc}})





	}


}


func UpdateProfile() gin.HandlerFunc{
	return func(c *gin.Context){
		var data models.RequestBody
		profile_id:=c.Param("id")
		if err := c.BindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest,responses.ProfileResponse{Status: http.StatusBadRequest, Message: "Error while getting the request", Data: map[string] interface{}{"data":err.Error()}})
			return


	}
	var getDoc models.RequestBody
	data.Pid=profile_id
	_,err:=col.Upsert(profile_id,data,nil)
	if errors.Is(err, gocb.ErrDocumentNotFound) {
		c.JSON(http.StatusInternalServerError,responses.ProfileResponse{Status: http.StatusInternalServerError, Message: "Error, Document with the key does not exist", Data: map[string] interface{}{"data":err.Error()}})
			return

	}
	getResult,err:=col.Get(profile_id,nil)
	err = getResult.Content(&getDoc)
	c.JSON(http.StatusOK,responses.ProfileResponse{Status: http.StatusOK, Message: "Successfully Updated the document", Data: map[string] interface{}{"data":getDoc}})
	}

}


func DeleteProfile() gin.HandlerFunc{
	return func(c *gin.Context){
		profile_id:=c.Param("id")
		result,err:=col.Remove(profile_id,nil)
		if errors.Is(err, gocb.ErrDocumentNotFound) {
			c.JSON(http.StatusInternalServerError,responses.ProfileResponse{Status: http.StatusInternalServerError, Message: "Error, Document Not found with the specified key", Data: map[string] interface{}{"data":err.Error()}})
			return 
		}

		c.JSON(http.StatusOK,responses.ProfileResponse{Status: http.StatusOK, Message: "Successfully deleted the document", Data: map[string] interface{}{"data":result}})


	}




}
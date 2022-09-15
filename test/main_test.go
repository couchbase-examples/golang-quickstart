package main

import (
    "encoding/json"
    "testing"
    "net/http"
    //"fmt"
    //"strconv"
    "bytes"
    "io/ioutil"
    "app/responses"
    
   
    
)

type Profile struct{
    Pid string 
	FirstName string 
	LastName string 
	Email string 
	Password string 
}


var base="http://127.0.0.1:8080"

func Test_health_check(t *testing.T){
    //Test the health check end point
    var url=base+"/api/v1/health"
    resp, err:=http.Get(url)
    //fmt.Println(resp.StatusCode)
    if resp.StatusCode!=200{
        t.Errorf("Error %d", err)
    }

    //fmt.Println(resp)
}

func Test_add_profile(t *testing.T){
    
    var url=base+"/api/v1/profile"
    
    postBody, _ := json.Marshal(map[string]string{
        "FirstName": "Mike",
        "LastName": "doe",
        "Email": "john.doe@couchbase.com",
        "Password": "password",
     })
     responseBody := bytes.NewBuffer(postBody)
     //fmt.Println(responseBody)
     resp, err := http.Post(url, "application/json", responseBody)
     
    
     
     //fmt.Println(sb)
    if resp.StatusCode!=200{
        t.Errorf("Error %d",err)
    }

}

func Test_add_profile_without_email(t *testing.T){

    var url=base+"/api/v1/profile"
    
    postBody, _ := json.Marshal(map[string]string{
        "FirstName": "Mike",
        "LastName": "doe",
        "Password": "password",
     })

     responseBody := bytes.NewBuffer(postBody)
     //fmt.Println(responseBody)
     resp, err := http.Post(url, "application/json", responseBody)
     
     
     
     //fmt.Println(sb)
    if resp.StatusCode==200{
        t.Errorf("Error %d",err)
    }

}


func Test_add_profile_without_email_and_password(t *testing.T){
    var url=base+"/api/v1/profile"
    
    postBody, _ := json.Marshal(map[string]string{
        "FirstName": "Mike",
        "LastName": "doe",
     })

     responseBody := bytes.NewBuffer(postBody)
     //fmt.Println(responseBody)
     resp, err := http.Post(url, "application/json", responseBody)
     
   
     //fmt.Println(sb)
    if resp.StatusCode==200{
        t.Errorf("Error %d",err)
    }
}


func Test_get_user_profile__invalid_id(t *testing.T){
    var id="1234"
    var url=base+"/api/v1/profile/"+id
    //fmt.Println(url)
    resp, err:=http.Get(url)
    //fmt.Println(resp)
    //fmt.Println(resp.StatusCode)
    if resp.StatusCode==200{
        //Error:Document not found
        t.Errorf("Error %d", err)
    }  

}

func Test_update_user_profile(t *testing.T){
    var url=base+"/api/v1/profile/"
    
    postBody, _ := json.Marshal(map[string]string{
        "FirstName": "Mike",
        "LastName": "doe",
        "Email": "john.doe@couchbase.com",
        "Password": "password",
     })
     responseBody := bytes.NewBuffer(postBody)
     resp, err := http.Post(url, "application/json", responseBody)
     //fmt.Println(resp)
     body, _ := ioutil.ReadAll(resp.Body)
     //sb := string(body)
     //fmt.Println(sb)
    if resp.StatusCode!=200{
        t.Errorf("Error %d",err)
    }
    m := make(map[string]interface{})
    json.Unmarshal(body,&m)
    id:=m["data"].(map[string] interface{})["pid"].(string)
    //fmt.Println(id)
    updated_postBody, _ := json.Marshal(map[string]string{
        "FirstName": "Mike",
        "LastName": "John",
        "Email": "mike.john@couchbase.com",
        "Password": "password",
     })
    updated_responseBody := bytes.NewBuffer(updated_postBody)
    
    req_update, _ := http.NewRequest(http.MethodPut, url+id, updated_responseBody)
    client := &http.Client{}
    resp_update, _ := client.Do(req_update)
    req_update.Header.Set("Content-Type", "application/json; charset=utf-8")
    resp_body, _ := ioutil.ReadAll(resp_update.Body)
    //fmt.Println(string(resp_body))
     var q responses.ProfileResponse
     json.Unmarshal(resp_body,&q)
     if q.Status!=200{
        t.Errorf("Error")
     }


}


func Test_delete_user_profile(t *testing.T){
    var url=base+"/api/v1/profile/"
    
    postBody, _ := json.Marshal(map[string]string{
        "FirstName": "Mike",
        "LastName": "doe",
        "Email": "john.doe@couchbase.com",
        "Password": "password",
     })
     responseBody := bytes.NewBuffer(postBody)
     resp, err := http.Post(url, "application/json", responseBody)
     body, _ := ioutil.ReadAll(resp.Body)
    if resp.StatusCode!=200{
        t.Errorf("Error %d",err)
    }
    m := make(map[string]interface{})
    json.Unmarshal(body,&m)
    id:=m["data"].(map[string] interface{})["pid"].(string)
    //fmt.Println(id)
    var url_delete=url+id
    req_delete, _ := http.NewRequest(http.MethodDelete, url_delete,nil)
    client := &http.Client{}
    resp_delete, _ := client.Do(req_delete)
    req_delete.Header.Set("Content-Type", "application/json")
    resp_body, _ := ioutil.ReadAll(resp_delete.Body)
    var q responses.ProfileResponse
    json.Unmarshal(resp_body,&q)
    if q.Status!=200{
       t.Errorf("Error")
    }
}


func Test_delete_invalid_user_profile(t *testing.T){
    id:="1234"
    var url=base+"/api/v1/profile/"
    var url_delete=url+id
    req_delete, _ := http.NewRequest(http.MethodDelete, url_delete,nil)
    client := &http.Client{}
    resp_delete, _ := client.Do(req_delete)
    req_delete.Header.Set("Content-Type", "application/json")
    resp_body, _ := ioutil.ReadAll(resp_delete.Body)
    //fmt.Println(resp_body)
    var q responses.ProfileResponse
    json.Unmarshal(resp_body,&q)
    if q.Status==200{
        t.Errorf("Error")
    }

}
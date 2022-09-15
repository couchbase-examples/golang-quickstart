package configs

import (
"github.com/couchbase/gocb/v2"
)

type SampleApp struct {
	cluster *gocb.Cluster
	bucket  *gocb.Bucket
	
}

func ConnectDB() *SampleApp{

username:="Administrator"
password:="Password"
bucket_name:="travel-sample"


clusterOpts := gocb.ClusterOptions{
	Authenticator: gocb.PasswordAuthenticator{
		Username: username,
		Password: password,
	},

	
}
cluster, err := gocb.Connect("couchbase://127.0.0.1", clusterOpts)

if err != nil {
	panic(err)
}

bucket:=cluster.Bucket(bucket_name)
app:= &SampleApp{
	cluster: cluster,
	bucket: bucket,

}
//scope:=app.bucket.Scope("_default")
//col:=scope.Collection("_default")
return app
}

var Cb *SampleApp=ConnectDB()

func GetCollection(app *SampleApp,CollectionName string) *gocb.Collection{
scope:=app.bucket.Scope("_default")
col:=scope.Collection("_default")

return col

}

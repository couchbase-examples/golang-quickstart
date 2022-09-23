package db


import(

	"github.com/couchbase/gocb/v2"
	"errors"
	"fmt"
	"time"
)

var username = EnvCouchBase("USERNAME")
var password=EnvCouchBase("PASSWORD")
var bucket_name=EnvCouchBase("BUCKET")
var scope_name=EnvCouchBase("SCOPE")
var collection_name=EnvCouchBase("COLLECTION")

func Create_bucket(cluster *gocb.Cluster) {
	fmt.Println("Creating bucket")
	//bucket_manager:=gocb.Buckets(cluster)
	bucket_manager:=cluster.Buckets()
	err:=bucket_manager.CreateBucket(gocb.CreateBucketSettings{
		BucketSettings:gocb.BucketSettings{
		Name: bucket_name,
		FlushEnabled: false,
		ReplicaIndexDisabled: true,
		RAMQuotaMB: 256,
		NumReplicas: 1,
		BucketType: gocb.CouchbaseBucketType,
		},
		ConflictResolutionType: gocb.ConflictResolutionTypeSequenceNumber,
	},nil)
	
	if(err!=nil){
		fmt.Println("Bucket already exists..")
		//panic(err)
	}
	fmt.Println("Created bucket..")

} 

func Create_scope(cluster *gocb.Cluster){
	bucket:=cluster.Bucket(bucket_name)
	collections:=bucket.Collections()
	err := collections.CreateScope(scope_name, nil)
	if(err!=nil){
		if errors.Is(err, gocb.ErrScopeExists) {
			fmt.Println("Scope already exists")
		} else {
			panic(err)
		}
	}



}

func Create_collection(cluster *gocb.Cluster){
	bucket:=cluster.Bucket(bucket_name)
	collections:=bucket.Collections()
	collection:=gocb.CollectionSpec{
		Name: collection_name,
		ScopeName: scope_name,
	}
	err:=collections.CreateCollection(collection,nil)
	if err != nil {
		if errors.Is(err, gocb.ErrCollectionExists) {
			fmt.Println("Collection already exists")
		} else {
			panic(err)
		}
	}



}

func Initialize_db() *gocb.Cluster{
	fmt.Println("Getting Environment Variables")
	clusterOpts := gocb.ClusterOptions{
		Authenticator: gocb.PasswordAuthenticator{
			Username: username,
			Password: password,
		},
	
		
	}
	cluster, err := gocb.Connect("couchbase://127.0.0.1", clusterOpts)
	//fmt.Println("Connection done")
	if err != nil {
		panic(err)
	}
	Create_bucket(cluster)
	time.Sleep(5*time.Second)
	Create_scope(cluster)
	Create_collection(cluster)
	time.Sleep(5*time.Second)
	fmt.Println("Initializing Database complete")
	return cluster

}

func GetCollection(cluster *gocb.Cluster) *gocb.Collection{
	bucket:=cluster.Bucket(bucket_name)
	scope:=bucket.Scope(scope_name)
	col:=scope.Collection(collection_name)
	return col
 


}
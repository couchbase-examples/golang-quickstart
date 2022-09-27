package db

import (
	"errors"
	"fmt"
	"github.com/couchbase/gocb/v2"
	"time"
)

var Host_name = EnvCouchBase("DB_HOST")
var Username = EnvCouchBase("USERNAME")
var Password = EnvCouchBase("PASSWORD")
var Bucket_name = EnvCouchBase("BUCKET")
var Scope_name = EnvCouchBase("SCOPE")
var Collection_name = EnvCouchBase("COLLECTION")

func Create_bucket(cluster *gocb.Cluster) {
	fmt.Println("Creating bucket")
	bucket_manager := cluster.Buckets()
	err := bucket_manager.CreateBucket(gocb.CreateBucketSettings{
		BucketSettings: gocb.BucketSettings{
			Name:                 Bucket_name,
			FlushEnabled:         false,
			ReplicaIndexDisabled: true,
			RAMQuotaMB:           256,
			NumReplicas:          1,
			BucketType:           gocb.CouchbaseBucketType,
		},
		ConflictResolutionType: gocb.ConflictResolutionTypeSequenceNumber,
	}, nil)

	if err != nil {
		fmt.Println("Bucket already exists..")
	}else{
	fmt.Println("Created bucket..")
	}

}

func Create_scope(cluster *gocb.Cluster) {
	bucket := cluster.Bucket(Bucket_name)
	collections := bucket.Collections()
	err := collections.CreateScope(Scope_name, nil)
	if err != nil {
		if errors.Is(err, gocb.ErrScopeExists) {
			fmt.Println("Scope already exists")
		} else {
			panic(err)
		}
	}

}

func Create_collection(cluster *gocb.Cluster) {
	bucket := cluster.Bucket(Bucket_name)
	collections := bucket.Collections()
	collection := gocb.CollectionSpec{
		Name:      Collection_name,
		ScopeName: Scope_name,
	}
	err := collections.CreateCollection(collection, nil)
	if err != nil {
		if errors.Is(err, gocb.ErrCollectionExists) {
			fmt.Println("Collection already exists")
		} else {
			panic(err)
		}
	}

}

func Create_primary_index(cluster *gocb.Cluster) {
	mgr := cluster.QueryIndexes()
	var a gocb.CreatePrimaryQueryIndexOptions
	a.ScopeName = Scope_name
	a.CollectionName = Collection_name
	if err := mgr.CreatePrimaryIndex(Bucket_name, &a); err != nil {
		if errors.Is(err, gocb.ErrIndexExists) {
			fmt.Println("Index already exists")
		} else {
			panic(err)
		}
	}
}

func Initialize_db() *gocb.Cluster {
	fmt.Println("Getting Environment Variables")
	clusterOpts := gocb.ClusterOptions{
		Authenticator: gocb.PasswordAuthenticator{
			Username: Username,
			Password: Password,
		},
	}
	//Connecting to Couchbase Server locally
	cluster, err := gocb.Connect("couchbase://"+Host_name, clusterOpts)
	//Connecting to Couchbase Capella
	//cluster, err := gocb.Connect("couchbases://"+Host_name, clusterOpts)
	if err != nil {
		panic(err)
	}
	Create_bucket(cluster)
	time.Sleep(5 * time.Second)
	Create_scope(cluster)
	Create_collection(cluster)
	time.Sleep(5 * time.Second)
	Create_primary_index(cluster)
	fmt.Println("Initializing Database complete")
	return cluster

}

func GetCollection(cluster *gocb.Cluster) *gocb.Collection {
	bucket := cluster.Bucket(Bucket_name)
	scope := bucket.Scope(Scope_name)
	col := scope.Collection(Collection_name)
	return col

}

package db

import (
	"fmt"
	"os"

	"github.com/couchbase/gocb/v2"
	"github.com/joho/godotenv"
)

// Define constants for environment variable keys
const (
	ConnectionString = "CONNECTION_STRING"
	UsernameKey      = "USERNAME"
	PasswordKey      = "PASSWORD"
	BucketName       = "travel-sample"
	ScopeName        = "inventory"
)

// InitializeCluster initializes the Couchbase cluster and returns it.
func InitializeCluster() *gocb.Cluster {
	fmt.Println("Initializing Database")

	// Load environment variables from a .env file
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}

	// Retrieve environment variables
	connectionString := getEnvVar(ConnectionString)
	username := getEnvVar(UsernameKey)
	password := getEnvVar(PasswordKey)

	// Configure cluster options
	clusterOpts := gocb.ClusterOptions{
		Authenticator: gocb.PasswordAuthenticator{
			Username: username,
			Password: password,
		},
	}

	// Connect to the Couchbase cluster
	cluster, err := gocb.Connect(connectionString, clusterOpts)
	if err != nil {
		panic(err)
	}

	// Check if the specified scope exists
	if !checkScopeExists(cluster, BucketName, ScopeName) {
		fmt.Println("The specified scope does not exist in the bucket.")
		os.Exit(1)
	}

	fmt.Println("Database initialization complete")
	return cluster
}

// GetScope returns a scope for the specified cluster, bucket, and scope name.
func GetScope(cluster *gocb.Cluster) *gocb.Scope {
	bucket := cluster.Bucket(BucketName)
	scope := bucket.Scope(ScopeName)
	return scope
}

// Helper function to retrieve an environment variable
func getEnvVar(key string) string {
	value := os.Getenv(key)
	if value == "" && (key == "USERNAME" || key == "PASSWORD" || key == "CONNECTION_STRING") {
		fmt.Printf("Environment variable %s is empty.\n", key)
		os.Exit(1)
	}
	return value
}

// Function to check if a scope exists in a bucket
func checkScopeExists(cluster *gocb.Cluster, bucketName, scopeName string) bool {
	bucket := cluster.Bucket(bucketName)
	// Fetch all scopes in the bucket
	scopes, err := bucket.Collections().GetAllScopes(nil)
	if err != nil {
		fmt.Println("Error fetching scopes in the cluster. Ensure that the bucket exists.")
		return false
	}

	// Check if the specified scope exists in the list of scopes
	for _, s := range scopes {
		if s.Name == scopeName {
			return true
		}
	}

	// Return false if the scope doesn't exist
	return false
}

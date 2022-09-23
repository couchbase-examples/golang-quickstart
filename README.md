# Quickstart in Couchbase with Golang  and Gin Gonic

#### Build a REST API with Couchbase's Golang SDK 2.5 and Gin Gonic

> This repo is designed to teach you how to connect to a Couchbase cluster to create, read, update, and delete documents and how to write simple parametrized N1QL queries.


Full documentation can be found on the [Couchbase Developer Portal]().

## Prerequisites

To run this prebuilt project, you will need:

- Follow [Couchbase Installation Options](https://developer.couchbase.com/tutorial-couchbase-installation-options) for installing the lastest Couchbase Database Server Instance
- [Golang v1.19.x](https://go.dev/dl/) installed
- Code Editor installed

## Install Dependencies

Any dependencies will be installed by running the go run command which installs the dependencies required from the go.mod file.

## Database Server Configuration

All configuration for communication with the database is stored in the .env file.The default username is assumed to be `Administrator` and the default password is assumed to be `Password`. If these are different in your environment, you will need to change them before running the application in the .env file.

Note: If you are running with [Couchbase Capella](https://cloud.couchbase.com/), the application requires the bucket and the database user to be setup from Capella Console. Also ensure that the bucket exists on the cluster and your [IP address is whitelisted](https://docs.couchbase.com/cloud/get-started/cluster-and-data.html#allowed) on the Cluster. In order to inialize the scope and collection in the Capella cluster, enable the commented out code in the Initialize_db() in the file db/db_init.go that is used to connect to Capella. Also Comment out the code to connect to Couchbase server locally.

### Running The Application

At this point the application is ready. Make sure you are in the **app** directory. You can run it with the following commands from the terminal/command prompt:

The bucket along with the scope and collection will be created on the cluster. For Capella, you need to ensure that the bucket is created before running the application.
```shell
go run .
```
### Running The Tests

To run the standard integration tests, use the following commands:

```bash
cd test
go test -v
```

## Conclusion

Setting up a basic REST API in Golang and Gin Gonic with Couchbase is fairly simple. In this project when ran with Couchbase Server 7 installed, it will create a bucket in Couchbase, an index for our parameterized [N1QL query](https://docs.couchbase.com/go-sdk/current/howtos/n1ql-queries-with-sdk.html), and showcases basic CRUD operations needed in most applications.
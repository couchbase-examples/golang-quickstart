# Quickstart in Couchbase with Golang  and Gin Gonic

#### Build a REST API with Couchbase's Golang SDK 2.5 and Gin Gonic

> This repo is designed to teach you how to connect to a Couchbase cluster to create, read, update, and delete documents and how to write simple parametrized N1QL queries.


Full documentation can be found on the [Couchbase Developer Portal]().

## Prerequisites

To run this prebuilt project, you will need:

- Follow [Get Started with Couchbase Capella](https://docs.couchbase.com/cloud/get-started/get-started.html) for more information about Couchbase Capella.
- Follow [Couchbase Installation Options](https://developer.couchbase.com/tutorial-couchbase-installation-options) for installing the latest Couchbase Database Server Instance
- Basic knowledge of [Golang](https://go.dev/tour/welcome/1) and [Gin Gonic](https://gin-gonic.com/docs/)
- [Golang v1.19.x](https://go.dev/dl/) installed
- Code Editor installed
- Note that this tutorial is designed to work with the latest Golang SDK (2.x) for Couchbase. It will not work with the older Golang SDK for Couchbase without adapting the code.

## Install Dependencies

Any dependencies will be installed by running the go run command, which installs any dependencies required from the go.mod file.

## Database Server Configuration

All configuration for communication with the database is stored in the `.env` file. This includes the Connection string, username, password, bucket name, collection name and scope name. The default username is assumed to be `Administrator` and the default password is assumed to be `Password1$`. If these are different in your environment you will need to change them before running the application.

### Running Couchbase Capella

When running Couchbase using Capella, the application requires the bucket and the database user to be setup from Capella Console. The directions for creating a bucket called `user_profile` can be found on the [documentation website](https://docs.couchbase.com/cloud/clusters/data-service/manage-buckets.html#add-bucket). Next, follow the directions for [Configure Database Credentials](https://docs.couchbase.com/cloud/clusters/manage-database-users.html) and name the username `Administrator` and password `Password1$`.

Next, open the `.env` file. Locate CONNECTION_STRING and update it to match your Wide Area Network name found in the [Capella Portal UI Connect Tab](https://docs.couchbase.com/cloud/get-started/connect-to-cluster.html#connect-to-your-cluster-using-the-built-in-sdk-examples). Note that Capella uses TLS so the Connection string must start with couchbases://. Note that this configuration is designed for development environments only.

```
CONNECTION_STRING=couchbases://yourhostname.cloud.couchbase.com
BUCKET=user_profile
COLLECTION=default
SCOPE=default
USERNAME=Administrator
PASSWORD=Password1$
```

### Running Couchbase Locally

For local installation and Docker users, follow the directions found on the [documentation website](https://docs.couchbase.com/server/current/manage/manage-buckets/create-bucket.html) for creating a bucket called `user_profile`. Next, follow the directions for [Creating a user](); name the username `Administrator` and password `Password1$`. For this tutorial, make sure it has `Full Admin` rights so that the application can create collections and indexes.

Next, open the `.env` file and validate that the configuration information matches your setup.

> **NOTE:** For docker and local Couchbase installations, Couchbase must be installed and running on localhost (<http://127.0.0.1:8091>) prior to running the the Golang application.

### Running The Application

At this point the application is ready. Make sure you are in the src directory. You can run it with the following command from the terminal/command prompt:

```shell
go run .
```

Once the site is up and running, you can launch your browser and go to the [Swagger start page](http://127.0.0.1:8080/docs/index.html) to test the APIs.

### Running The Tests

To run the standard integration tests, use the following commands from the src directory:

```bash
cd test
go test -v
```

## Conclusion

Setting up a basic REST API in Golang and Gin Gonic with Couchbase is fairly simple,this project when run will showcase basic CRUD operations along with creating an index for our parameterized  [N1QL query](https://docs.couchbase.com/go-sdk/current/howtos/n1ql-queries-with-sdk.html) which is used in most applications.
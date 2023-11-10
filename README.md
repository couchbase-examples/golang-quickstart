# Quickstart in Couchbase with Golang  and Gin Gonic

#### Build a REST API with Couchbase's Golang SDK 2.5 and Gin Gonic

> This repo is designed to teach you how to connect to a Couchbase Capella cluster to create, read, update, and delete documents and how to write simple parametrized SQL++ queries using the built-in travel-sample bucket. If you want to run this tutorial using a self managed Couchbase cluster, please refer to the [appendix](#appendix-running-self-managed-couchbase-cluster).


Full documentation can be found on the [Couchbase Developer Portal](https://developer.couchbase.com/tutorial-quickstart-golang-gin-gonic).

## Prerequisites

To run this prebuilt project, you will need:

- Couchbase Server (7 or higher) with [travel-sample](https://docs.couchbase.com/go-sdk/current/ref/travel-app-data-model.html) bucket loaded.
  - [Couchbase Capella](https://www.couchbase.com/products/capella/) is the easiest way to get started.
- Basic knowledge of [Golang](https://go.dev/tour/welcome/1) and [Gin Gonic](https://gin-gonic.com/docs/)
- [Golang v1.21.x](https://go.dev/dl/) installed

### Loading Travel Sample Bucket

If travel-sample is not loaded in your Capella cluster, you can load it by following the instructions for your Capella Cluster:

- [Load travel-sample bucket in Couchbase Capella](https://docs.couchbase.com/cloud/clusters/data-service/import-data-documents.html#import-sample-data)
## Install Dependencies

Any dependencies will be installed by running the go run command, which installs any dependencies required from the go.mod file.


### Database Server Configuration

All configuration for communication with the database is read from the environment variables. We have provided a convenience feature in this quickstart to read the environment variables from a local file, `.env` in the source folder.

Create a copy of .env.example & rename it to .env & add the values for the Couchbase connection.

To know more about connecting to your Capella cluster, please follow the [instructions](https://docs.couchbase.com/cloud/get-started/connect.html).

```sh
CONNECTION_STRING=<connection_string>
USERNAME=<user_with_read_write_permission_to_travel-sample_bucket>
PASSWORD=<password_for_user>
```

> Note: The connection string expects the `couchbases://` or `couchbase://` part.

## Running The Application

### Running directly on machine

At this point, we have installed the dependencies, loaded the travel-sample data and configured the application with the credentials. The application is now ready and you can run it.

The application will run on port 8080 of your local machine (http://localhost:8080). You will find the Swagger documentation of the API.

```sh
cd src
go run .
```

### Running using Docker

- Build the Docker image

```sh
cd src
docker build -t couchbase-gin-gonic-quickstart .
```

- Run the Docker image

```sh
docker run -it --env-file .env -p 8080:8080 couchbase-gin-gonic-quickstart
```

> Note: The `.env` file has the connection information to connect to your Capella cluster. The application can now be reached on port 8080 of your local machine.

## Running The Tests

To run the standard unit tests, use the following commands:

```sh
cd test
go test -v
```

## Appendix: Data Model

For this quickstart, we use three collections, `airport`, `airline` and `routes` that contain sample airports, airlines and airline routes respectively. The routes collection connects the airports and airlines as seen in the figure below. We use these connections in the quickstart to generate airports that are directly connected and airlines connecting to a destination airport. Note that these are just examples to highlight how you can use SQL++ queries to join the collections.

![travel sample data model](travel_sample_data_model.png)

## Appendix: Running Self Managed Couchbase Cluster

If you are running this quickstart with a self managed Couchbase cluster, you need to [load](https://docs.couchbase.com/server/current/manage/manage-settings/install-sample-buckets.html) the travel-sample data bucket in your cluster and generate the credentials for the bucket.

You need to update the connection string and the credentials in the `.env` file in the source folder.

> Note: Couchbase Server must be installed and running prior to running this app.


## Conclusion

Setting up a basic REST API in Golang and Gin Gonic with Couchbase is fairly simple,this project when run will showcase basic CRUD operations along with executing [SQL++ query](https://docs.couchbase.com/go-sdk/current/howtos/n1ql-queries-with-sdk.html) which is used in most applications.

# Use the official Golang image as the base image
FROM golang:1.21.3

# Set the working directory within the container
WORKDIR /app

# Copy the source code into the container
COPY . .

# Build the Golang application inside the container
RUN go build -o main

# Expose any necessary ports (if applicable)
EXPOSE 8080

# Set the entry point to run your Golang application
CMD ["./main"]
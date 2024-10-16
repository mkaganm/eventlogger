# Step 1: Build the Go application
FROM golang:1.23-alpine as builder

# Set the current working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all Go dependencies
RUN go mod download

# Copy the source code to the container
COPY . .

# Build the Go application
RUN go build -o /app/main .

# Step 2: Create a lightweight container for running the Go application
FROM alpine:latest

# Install bash to use it in CMD for .env sourcing
RUN apk add --no-cache bash

# Set the working directory inside the container
WORKDIR /app

# Copy the built binary from the builder
COPY --from=builder /app/main .

# Copy the environment file
COPY .env .env

# Expose port (optional, if your app exposes a port)
EXPOSE 8080

# Command to run the executable, sourcing the .env file
CMD ["bash", "-c", "source .env && ./main"]

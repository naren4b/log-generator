# Use the official Golang image
FROM golang:1.17-alpine AS build

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod .
COPY go.sum .

# Download dependencies
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go application
RUN go build -o loggenerator .

# Use a minimal Alpine image to create the final runtime image
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the built executable from the previous stage
COPY --from=build /app/loggenerator .

# Run the Go application
CMD ["./loggenerator"]

# Optionally, specify a default duration (5 minutes)
# CMD ["./loggenerator", "5"]

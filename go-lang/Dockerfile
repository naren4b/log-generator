# Use an official Go runtime as a parent image
FROM golang:1.17

# Set the working directory inside the container
WORKDIR /app

# Copy the Go application source code and go.mod/go.sum files into the container
COPY go.mod ./

# Install dependencies
RUN go mod download

# Copy the rest of the application source code into the container
COPY . .

# Build the Go application
RUN go build -o logger .

# Command to run the Go application
CMD ["./logger"]

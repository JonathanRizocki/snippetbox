# Start from the latest golang base image
FROM golang:latest

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files first (for dependency caching)
COPY go.mod ./

# Download dependencies
RUN go mod download

# Copy the rest of the source code into the container
COPY . .

# Build the Go app
RUN go build -o /app/web ./cmd/web/.

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["/app/web"]

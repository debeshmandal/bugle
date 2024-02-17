# Stage 1: Build the Go application
FROM golang:1.22-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o bugle .

# Stage 2: Build a small image
FROM alpine:latest

WORKDIR /opt/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/bugle .

# Command to run the executable
ENTRYPOINT [ "/opt/bugle" ]

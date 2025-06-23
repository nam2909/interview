# Start with a golang base image
FROM golang:1.22.1 as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o counter ./

# Start a new stage from scratch
FROM alpine:latest

# Add Maintainer Info
LABEL maintainer="namnguyen <nguyenducnam2509@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR /app/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/counter .

# Command to run the executable
CMD ["./counter"]
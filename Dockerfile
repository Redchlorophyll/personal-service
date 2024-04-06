# Build stage
FROM golang:1.21.6-alpine3.19 as builder

# Set the working directory in the container
WORKDIR /app

# Copy the application files into the working directory
COPY . /app

# Download dependencies
RUN go mod download

# Build the application
RUN go build -o build/monolith ./cmd/httpservice/personal_service

# Runtime stage
FROM alpine:latest

# Set the working directory in the container
WORKDIR /app

# Copy only the build binary from the builder image
COPY --from=builder /app/build .

# Expose port 8080
EXPOSE 8080

# Define the entry point for the container
CMD ["./monolith"]

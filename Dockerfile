# Build stage
FROM golang:1.21.6-alpine3.19 as builder

# Install git
RUN apk --no-cache add git

# Set environment variable for builder
ARG GIT_PERSONAL_USERNAME
ARG GIT_PERSONAL_EMAIL
ARG PAT_TOKEN
ARG ENV_PATH
ARG SECRET_REPO
ENV GITHUB_TOKEN=$PAT_TOKEN
ENV ENV_PATH=$ENV_PATH
ENV GIT_PERSONAL_USERNAME=$GIT_PERSONAL_USERNAME
ENV GIT_PERSONAL_EMAIL=$GIT_PERSONAL_EMAIL
ENV SECRET_REPO=$SECRET_REPO

RUN git config --global user.name "$GIT_PERSONAL_USERNAME" && \
    git config --global user.email "$GIT_PERSONAL_EMAIL"

# Set the working directory in the container
WORKDIR /app

# Clone the private repository
RUN git clone https://$GIT_PERSONAL_USERNAME:$PAT_TOKEN@github.com/$GIT_PERSONAL_USERNAME/$SECRET_REPO

# Copy the application files into the working directory
COPY . /app

# Download dependencies
RUN go mod download

# Build the application
RUN go build -o build/monolith ./cmd/httpservice/personal_service

# Runtime stage
FROM alpine:latest

# Set environment variable for image
ARG ENV_PATH
ENV ENV_PATH=$ENV_PATH

# Set the working directory in the container
WORKDIR /app

# Copy only the build binary from the builder image
COPY --from=builder /app/build .
COPY --from=builder /app/$ENV_PATH .

# Expose port 8080
EXPOSE 8080

# Define the entry point for the container
CMD ["./monolith"]

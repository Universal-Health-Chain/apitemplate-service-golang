
# build it:
# docker build -t go-docker .
# run it:
# docker run -d -p 8000:8000 go-docker
# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Using the last version
FROM golang:1.21 AS builder

# Add Maintainer Info

# Set the Current Working Directory inside the container
WORKDIR /app

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
#RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o docs-service

# Use a Docker multi-stage build to create a lean production image.
# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
FROM debian:12-slim

# Copy the binary to the production image from the builder stage.
COPY --from=builder /app/docs-service /docs-service

# Expose port 8024 to the outside world
EXPOSE 8024

VOLUME [ "/logs" ]

# Command to run the executable
CMD ["/docs-service"]
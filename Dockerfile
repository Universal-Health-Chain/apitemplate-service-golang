
# build it:
# docker build -t go-docker .
# run it:
# docker run -d -p 8000:8000 go-docker
# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Using the last version
FROM golang:alpine

# Add Maintainer Info

# Set the Current Working Directory inside the container
WORKDIR /app

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
#RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main .

# Expose port 8024 to the outside world
EXPOSE 8024

# Command to run the executable
CMD ["./main"]
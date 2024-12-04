# Start from the latest golang base image
FROM golang:latest

# Add Maintainer Info
LABEL maintainer="minhnguyenvu4@gmail.com"

# Set the Current Working Directory inside the container
WORKDIR /

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Download and install any required dependencies
RUN go mod download

# Build the Go app
RUN go build -o main .

# Expose port 8081 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
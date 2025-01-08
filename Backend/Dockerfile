# Start from the official Golang image (latest version)
FROM golang:1.23

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Install Air for live reload
RUN go install github.com/air-verse/air@latest

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the Go app using Air directly (without Docker inside)
CMD ["air", "-c", "/app/.air.toml"]

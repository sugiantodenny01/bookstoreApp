# Use the official Go image
FROM golang:1.20

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum, then download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the Go app
RUN go build -o main .

# Expose the app's port
EXPOSE 8080

# Run the binary
CMD ["./main"]

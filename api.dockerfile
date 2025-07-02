# Use official Go image
FROM golang:1.22-alpine

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum files first
COPY go.mod ./
COPY go.sum ./

# Download dependencies
RUN go mod download

# Copy rest of the source code
COPY . .

# Build the binary
RUN go build -o main ./cmd/api

# Run the app
CMD ["./api"]

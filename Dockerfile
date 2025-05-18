FROM golang:1.21

# Set working directory inside container
WORKDIR /app

# Explicitly enable Go modules
ENV GO111MODULE=on

# Copy go.mod and go.sum first (for layer caching)
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire project
COPY . .

# Build the binary
RUN go build -o main ./cmd/main.go

# Expose port used by the app
EXPOSE 8080

# Run the app
CMD ["./main"]

FROM golang:1.24 AS builder

WORKDIR /app

# Copy go.mod file
COPY ./src/go.mod ./

# Download dependencies
RUN go mod download

# Copy source code
COPY ./src/ ./

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /assets ./main.go

# Print the contents to verify the file exists
RUN ls -la /assets

# Use a smaller image for the final container
FROM alpine:latest

WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /assets .

COPY ./src/config ./config

# Print directory contents for debugging
RUN ls -la /root

# Make sure the file is executable
RUN chmod +x /root/assets

# Add a sleep command to keep the container running for debugging
CMD ["./assets"]

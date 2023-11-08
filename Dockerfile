# Base image
FROM golang:1.21-alpine as builder

# Set the working directory
WORKDIR /app

# Copy the source files
COPY go.mod .
COPY go.sum .
COPY cmd/server/* .

# Download and install dependencies
RUN go mod download

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

# Final image
FROM alpine:latest

# Install necessary packages
RUN apk add --no-cache ca-certificates

# Copy over the built binary from the builder stage
COPY --from=builder /app/main .

# Define the entrypoint
ENTRYPOINT ["./cmd/server/main"]

EXPOSE 50051
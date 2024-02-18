#
# Stage 1: Build the application
#
FROM golang:1.22.0-alpine3.19 as builder

# Copy .env file into the Docker context
COPY .env .

WORKDIR /app

COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

#
# Stage 2: Create a minimal runtime container
#
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/main .

# Expose port if your Go application listens on a specific port
EXPOSE 3000

# Command to run the executable with environment variable
CMD ["./main"]
#
# Stage 1: Build the application
#
FROM golang:1.22.0-alpine3.19 as builder

WORKDIR /app

COPY . .

# Build the Go application
RUN go build -o main main.go

#
# Stage 2: Create a minimal runtime container
#
FROM alpine:latest

WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/main .

COPY app.env .

# Expose port if your Go application listens on a specific port
EXPOSE 3000

# Command to run the executable with environment variable
CMD ["/app/main"]
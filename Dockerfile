#
# Stage 1: Build the application
#
FROM golang:1.22.0-alpine3.19 as builder

WORKDIR /app

COPY . .

# Install softwares
RUN apk add --no-cache curl

# Build the Go application
RUN go build -o main main.go
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.0/migrate.linux-amd64.tar.gz | tar xvz

#
# Stage 2: Create a minimal runtime container
#
FROM alpine:latest

WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/main .
COPY --from=builder /app/migrate ./migrate

COPY app.env .
COPY db/migrations ./db/migrations
COPY start.sh .

# Expose port if your Go application listens on a specific port
EXPOSE 3000

# Command to run the executable with environment variable
CMD ["/app/main"]
ENTRYPOINT [ "/app/start.sh" ]
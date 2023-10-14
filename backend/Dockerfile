# Use the golang:1.21.3-alpine image as a base image
FROM golang:1.21.3-alpine as builder

# Set the working directory to /app
WORKDIR /app

COPY . .
RUN go mod download

# Copy the current directory contents into the container at /app
#COPY . .

# Build the application
#RUN go build
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/app

# Use the alpine image as a base image
FROM alpine:latest as final

WORKDIR /app

# Copy the binary from the first stage
COPY --from=builder /app/app .

# Copy env.local and config.yml from the first stage
COPY --from=builder /app/.env.local .
COPY --from=builder /app/config.yml .

# Use port from config.yml or docker-compose.yml
EXPOSE 8080

# Run the executable
CMD ["./app"]

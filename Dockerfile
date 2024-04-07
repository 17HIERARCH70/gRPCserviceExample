# Use an official Go runtime as a base image
FROM golang:1.22.1-alpine AS builder

RUN apk add --no-cache \
    gcc \
    musl-dev

WORKDIR /app

# Copy go mod and sum files
COPY ./sso/go.mod ./sso/go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the rest of the application code
COPY ./sso .

# Build the executable
RUN CGO_ENABLED=1 GOOS=linux go build -a -ldflags '-linkmode external -extldflags "-static"' -o grpc-sso ./cmd/sso/main.go

# Use a newer base image to run the application
FROM scratch
WORKDIR /app
COPY --from=builder /app/grpc-sso .
COPY ./sso/config/prod.yaml /app/config/prod.yaml
EXPOSE 44044
CMD ["./grpc-sso", "--config=config/prod.yaml"]

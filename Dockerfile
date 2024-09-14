# Use the official Golang image as the base image
FROM golang:1.23-alpine AS build
RUN apk add --no-cache gcc musl-dev

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=1 GOOS=linux go build -ldflags="-w -s" -o main main.go

# Production
FROM alpine:latest
WORKDIR /app
COPY --from=build ./app/main .
COPY .env .env

EXPOSE 30001

CMD ["./main"]
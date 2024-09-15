# Use the official Golang image as the base image
FROM golang:1.23-alpine AS build

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags="-w -s" -o main main.go

# Production
FROM alpine:latest
WORKDIR /app
COPY --from=build ./app/main .
COPY .env .env

EXPOSE 30001

CMD ["./main"]
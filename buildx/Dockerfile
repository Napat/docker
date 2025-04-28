# syntax=docker/dockerfile:1
FROM golang:1.24-alpine AS build
WORKDIR /app

COPY go.mod ./
COPY main.go ./

# Build binary without CGO
RUN CGO_ENABLED=0 go build -o /app/server

# Create a minimal image
FROM alpine:3.21
COPY --from=build /app/server /server
EXPOSE 8080

CMD ["/server"]
# syntax=docker/dockerfile:1

# Build the application from source
FROM golang:alpine AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /gotiles

# Deploy the application binary into a lean image
FROM alpine AS build-release-stage

WORKDIR /app

COPY --from=build-stage /gotiles ./gotiles

EXPOSE 3003

USER nonroot:nonroot

ENTRYPOINT ["./gotiles"]
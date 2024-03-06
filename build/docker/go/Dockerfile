FROM golang:1.22-alpine AS build

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o pismo ./cmd/api

FROM alpine:3.18

COPY --from=build /app/pismo /usr/local/bin
COPY --from=build /app/.env .

EXPOSE 8080
FROM golang:1.22.4-alpine AS build-stage

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/server

FROM alpine:latest AS build-release-stage

WORKDIR /

COPY --from=build-stage /app/main /echelon-app

EXPOSE 54320

ENTRYPOINT ["./echelon-app"]
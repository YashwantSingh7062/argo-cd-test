FROM golang:1.24-alpine AS base

WORKDIR /app

COPY lib ./
COPY main.go ./
COPY go.mod go.sum ./

RUN go mod download && go mod verify
RUN GOOS=linux go build -o ./tmp/release .

FROM alpine:latest

WORKDIR /app

COPY --from=base /app/tmp/release ./

EXPOSE "8000"

ENTRYPOINT ["./release"]
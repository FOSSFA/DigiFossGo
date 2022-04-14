# syntax=docker/dockerfile:1
FROM golang:1.18-alpine as builder

ENV GOOS=linux
ENV CGO_ENABLED=0

WORKDIR /app

# COPY go.mod ./
# COPY go.sum ./

COPY . ./

RUN go mod download

RUN go build -o bot ./cmd/DigiFoss/main.go


FROM alpine

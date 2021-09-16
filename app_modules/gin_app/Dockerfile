FROM docker.io/golang:1.17-alpine AS builder
ENV GOOS linux
ENV CGO_ENABLED 0
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o app

FROM docker.io/alpine:latest AS production
RUN apk add --no-cache ca-certificates
COPY --from=builder app .
EXPOSE 9000
CMD ./app
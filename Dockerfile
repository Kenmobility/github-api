# Build stage
FROM golang:1.20-alpine3.16 AS builder

RUN mkdir /app
ADD . /app
WORKDIR /app

COPY . .

RUN go build -o main cmd/main.go

# Run stage
FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/main .
COPY .env .

EXPOSE 8080
CMD [ "/app/main" ]
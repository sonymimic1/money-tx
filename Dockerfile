# Build stage
FROM golang:1.19-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go


# Run stage
FROM alpine:3.16.3
WORKDIR /app
COPY --from=builder /app/main .
COPY ./.env .
EXPOSE 8080
CMD ["/app/main"]

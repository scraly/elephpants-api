FROM golang:1.18-alpine AS builder
WORKDIR /build
COPY . .
RUN go mod download
RUN go build -o ./elephpants-api internal/main.go

FROM alpine:latest AS final
COPY --from=builder /build/elephpants-api .
EXPOSE 8080
CMD ["./elephpants-api"]

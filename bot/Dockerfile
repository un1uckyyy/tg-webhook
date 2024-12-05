FROM golang:1.23.2-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY .. .
ENV GOCACHE=/root/.cache/go-build
RUN --mount=type=cache,target=/root/.cache/go-build go build -o /telegram-bot ./main.go

FROM alpine:latest
COPY --from=builder /telegram-bot /telegram-bot
RUN chmod +x /telegram-bot
EXPOSE 8080
CMD ["/telegram-bot"]

FROM golang:1.25.7-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
COPY . .

RUN go build -o kino-server ./cmd/main.go

FROM alpine:3.18

WORKDIR /app

COPY --from=builder /app/kino-server .
COPY --from=builder /app/web ./web

EXPOSE 1111

ENV DB_HOST=db
ENV DB_PORT=5432
ENV DB_USER=postgres
ENV DB_PASSWORD=1234
ENV DB_NAME=kino

CMD ["./kino-server"]
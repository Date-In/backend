FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main ./cmd/main.go

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/migrate ./cmd/migrate/main.go


FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/migrate .

EXPOSE 8081

CMD ["/app/main"]
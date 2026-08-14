# Etapa 1: Build da aplicacao
FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o http-server-projeto-korp .

# Etapa 2: Imagem final
FROM alpine:3.20

WORKDIR /app

COPY --from=builder /app/http-server-projeto-korp .

EXPOSE 8080

CMD ["./http-server-projeto-korp"]

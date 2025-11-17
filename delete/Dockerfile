FROM golang:1.25 AS builder
WORKDIR /app
COPY . .

# Instalar dependencias
RUN go mod tidy

# Compilar binario
RUN go build -o main .

# Imagen final minimalista
FROM gcr.io/distroless/base-debian12
COPY --from=builder /app/main .

CMD ["./main"]

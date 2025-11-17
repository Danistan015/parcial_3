FROM golang:1.25 AS builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o main .

FROM gcr.io/distroless/base-debian12
COPY --from=builder /app/main .
CMD ["./main"]

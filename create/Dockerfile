FROM golang:1.25 AS builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o main .

FROM gcr.io/distroless/base-debian12
WORKDIR /
COPY --from=builder /app/main .
EXPOSE 8001
CMD ["./main"]

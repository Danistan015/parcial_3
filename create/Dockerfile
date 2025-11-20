# esta imagen trae go ya instalado, la uso para compilar
FROM golang:1.25 AS builder

# creo una carpeta donde voy a trabajar
WORKDIR /app

# copio todo el proyecto dentro del contenedor
COPY . .

# arreglo los modulos por si falta algo
RUN go mod tidy

# compilo el proyecto y dejo el binario llamado main
RUN go build -o main .

# aca uso una imagen mas liviana solo para correr el binario
FROM gcr.io/distroless/base-debian12
    
# carpeta donde correra la app
WORKDIR /

# copio el binario ya compilado desde el stage builder
COPY --from=builder /app/main .

# puerto donde escucha mi microservicio
EXPOSE 8001

# comando para ejecutar la app
CMD ["./main"]

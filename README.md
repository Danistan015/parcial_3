Parcial 3 – Microservicios CRUD en Golang con Docker y MongoDB

Este proyecto desarrolla un sistema basado en microservicios independientes, donde cada operación del CRUD (Create, Read, Update y Delete) corresponde a un servicio separado.
Cada microservicio está implementado en Golang, utiliza MongoDB como base de datos y está contenedorizado mediante Docker, con orquestación en Docker Compose.

El trabajo busca evidenciar conocimientos en:

Arquitectura de microservicios

Go + mongo-go-driver

Docker y Docker Compose

Pruebas unitarias

Buenas prácticas de desarrollo

Backup y restauración de bases de datos MongoDB

Uso de Postman para validación

Preparación para CI/CD con GitHub Actions

1. Estructura del Proyecto
crud-albums/
│── create/               # Microservicio CREATE
│── read/                 # Microservicio READ
│── update/               # Microservicio UPDATE
│── delete/               # Microservicio DELETE
│── backup/               # Respaldos de MongoDB
│── postman/              # Colección Postman
│── docker-compose.yml
│── .env
│── README.md


Cada microservicio contiene:

controller.go
service.go
repository.go
model.go
main.go
Dockerfile
go.mod
go.sum

2. Arquitectura General

Flujo interno:

Cliente → Controller → Service → Repository → MongoDB


Servicios orquestados:

Docker Compose
├── create (8001)
├── read (8002)
├── update (8003)
├── delete (8004)
└── mongo (27017)

3. Ejecución del Proyecto con Docker Compose
Requisitos previos

Docker

Docker Compose

Archivo .env
MONGO_USER=admin
MONGO_PASS=admin123
MONGO_DB=clientsdb
MONGO_COLLECTION=clients

Levantar la plataforma completa
docker compose up --build

Detener todos los servicios
docker compose down

4. Endpoints de los Microservicios
Crear Cliente (CREATE)
POST http://localhost:8001/clients


Body:

{
  "name": "Daniela",
  "email": "daniela@example.com",
  "phone": "3210001111"
}

Obtener todos los clientes (READ)
GET http://localhost:8002/clients

Obtener cliente por ID
GET http://localhost:8002/clients/{id}

Actualizar Cliente (UPDATE)
PUT http://localhost:8003/clients/{id}


Body:

{
  "name": "Nuevo Nombre",
  "email": "nuevo@example.com",
  "phone": "3001112222"
}

Eliminar Cliente (DELETE)
DELETE http://localhost:8004/clients/{id}

5. Pruebas Unitarias

Cada microservicio incluye pruebas unitarias enfocadas en su controlador.

Ejecutar pruebas:

go test ./...


Ver cobertura:

go test ./... -cover


Ejemplo de salida:

ok    create   0.31s   coverage: 90.0% of statements

6. Backup y Restore de MongoDB
Crear un backup dentro del contenedor
docker exec -it mongo-albums bash

mongodump \
  -u "admin" \
  -p "admin123" \
  --authenticationDatabase "admin" \
  --db "clientsdb" \
  --out "/backup"


Comprimir:

tar -czvf /backup-YYYYMMDD-HHMM.tar.gz /backup
exit

Exportar backup al equipo local
docker cp mongo-albums:/backup-20251117-1650.tar.gz ./backup/

7. Colección de Postman

Se incluye en:

/postman/clients-crud.postman_collection.json


Para usarla:

Abrir Postman

Importar el archivo

Ejecutar pruebas

8. Preparación para CI/CD

El proyecto está preparado para incorporar:

GitHub Actions

Ejecución automática de pruebas

Build y push de imágenes

Escaneo de seguridad con Trivy

Flujo de despliegue automático

Archivo ci.yml pendiente de integración.
# Parcial 3 – Microservicios en Golang con Docker y MongoDB

Este proyecto implementa un sistema de **microservicios CRUD**, donde cada operación (**Create, Read, Update, Delete**) es un microservicio independiente, desarrollado en **Golang**, con base de datos **MongoDB**, completamente contenedorizado con **Docker** y orquestado mediante **Docker Compose**.

El objetivo del parcial es demostrar el uso de:

- Arquitectura de microservicios  
- Golang + mongo-go-driver  
- Docker + Docker Compose  
- Pruebas unitarias  
- Buenas prácticas de desarrollo  
- Backup de base de datos  
- Postman para validación  
- Preparación para CI/CD en GitHub Actions  

---

## 📂 Estructura del Proyecto

crud-albums/
│── create/ # Microservicio CREATE
│── read/ # Microservicio READ
│── update/ # Microservicio UPDATE
│── delete/ # Microservicio DELETE
│── backup/ # Archivo de respaldo Mongo
│── postman/ # Colección de pruebas
│── docker-compose.yml
│── .env
│── README.md


Cada microservicio contiene:

controller.go
repository.go
service.go
model.go
main.go
Dockerfile
go.mod
go.sum


---

## 🧱 Arquitectura General



Cliente → Controller → Service → Repository → MongoDB


Arquitectura de servicios:



Docker Compose
├── create (8001)
├── read (8002)
├── update (8003)
├── delete (8004)
└── mongo (27017)


---

## 🐳 Ejecución del Proyecto (Docker Compose)

### Requisitos

- Docker  
- Docker Compose  
- Archivo `.env`:



MONGO_USER=admin
MONGO_PASS=admin123
MONGO_DB=clientsdb
MONGO_COLLECTION=clients


### Ejecutar todos los servicios



docker compose up --build


### Detener los servicios



docker compose down


---

## 🔌 Endpoints de los Microservicios

### ➤ Crear Cliente (CREATE)



POST http://localhost:8001/clients

{
  "name": "Daniela",
  "email": "daniela@example.com",
  "phone": "3210001111"
}

➤ Obtener Todos los Clientes (READ)
GET http://localhost:8002/clients

➤ Obtener Cliente por ID
GET http://localhost:8002/clients/{id}

➤ Actualizar Cliente (UPDATE)
PUT http://localhost:8003/clients/{id}


Body JSON:

{
  "name": "Nuevo Nombre",
  "email": "nuevo@example.com",
  "phone": "3001112222"
}

➤ Eliminar Cliente (DELETE)
DELETE http://localhost:8004/clients/{id}

🧪 Pruebas Unitarias

Cada microservicio tiene pruebas unitarias para su controlador.

Ejecutar pruebas
go test ./...

Ver cobertura
go test ./... -cover


Ejemplo esperado:

ok  	create	0.312s	coverage: 90.0% of statements

📦 Backup y Restore de MongoDB
Generar backup dentro del contenedor
docker exec -it mongo-albums bash

mongodump \
  -u "admin" \
  -p "admin123" \
  --authenticationDatabase "admin" \
  --db "clientsdb" \
  --out "/backup"

tar -czvf /backup-YYYYMMDD-HHMM.tar.gz /backup

exit

Copiar backup a tu máquina local
docker cp mongo-albums:/backup-20251117-1650.tar.gz ./backup/

🧪 Colección de Postman

Se incluye una colección exportada en:

/postman/clients-crud.postman_collection.json


Para usarla:

Abrir Postman

Importar archivo

Ejecutar pruebas

🛠 Preparación para CI/CD

El proyecto está preparado para integrar:

GitHub Actions

Ejecución automática de tests

Build y push de imágenes

Escaneo de seguridad (Trivy)

Releases automáticos

(Workflow ci.yml pendiente.)

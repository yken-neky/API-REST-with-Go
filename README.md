# API REST with Go

Este proyecto implementa una API REST en Go con las siguientes características:

- CRUD operations
- PostgreSQL como base de datos
- Arquitectura limpia y moderna
- Preparado para producción

## Estructura del Proyecto

```
.
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   ├── config/
│   ├── handlers/
│   ├── models/
│   ├── repository/
│   └── service/
├── pkg/
│   └── database/
├── .env
├── go.mod
└── README.md
```

## Requisitos

- Go 1.21 o superior
- PostgreSQL 15 o superior
- Docker (opcional)

## Configuración

1. Clonar el repositorio
2. Copiar `.env.example` a `.env` y configurar las variables de entorno
3. Ejecutar `go mod download` para instalar dependencias
4. Iniciar el servidor con `go run cmd/api/main.go`

## Endpoints

- `GET /api/v1/items` - Listar todos los items
- `GET /api/v1/items/:id` - Obtener un item
- `POST /api/v1/items` - Crear un item
- `PUT /api/v1/items/:id` - Actualizar un item
- `DELETE /api/v1/items/:id` - Eliminar un item

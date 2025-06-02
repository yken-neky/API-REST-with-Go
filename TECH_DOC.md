# Documentación Técnica - API REST con Go

## Índice
1. [Estructura del Proyecto](#estructura-del-proyecto)
2. [Arquitectura](#arquitectura)
3. [Configuración](#configuración)
4. [Base de Datos](#base-de-datos)
5. [API Endpoints](#api-endpoints)
6. [Middleware](#middleware)
7. [Monitoreo y Salud](#monitoreo-y-salud)
8. [Guía de Desarrollo](#guía-de-desarrollo)
9. [Notas de Migración](#notas-de-migración)

## Estructura del Proyecto

```
.
├── cmd/
│   └── api/
│       └── main.go           # Punto de entrada de la aplicación
├── internal/
│   ├── config/              # Configuración de la aplicación
│   │   ├── config.go        # Estructura de configuración
│   │   └── environment.go   # Carga de variables de entorno
│   ├── core/               # Núcleo de la aplicación
│   │   ├── domain/         # Modelos de dominio
│   │   ├── ports/          # Interfaces (puertos)
│   │   └── services/       # Implementación de servicios
│   └── infrastructure/     # Implementaciones concretas
│       ├── database/       # Capa de persistencia
│       └── server/         # Servidor HTTP y rutas
├── .env                    # Variables de entorno
└── go.mod                  # Dependencias del proyecto
```

### Descripción de Carpetas

- `cmd/`: Contiene los puntos de entrada de la aplicación.
- `internal/`: Código privado de la aplicación.
  - `config/`: Manejo de configuración y variables de entorno.
  - `core/`: Implementa la lógica de negocio.
    - `domain/`: Modelos y reglas de negocio.
    - `ports/`: Interfaces que definen los contratos.
    - `services/`: Implementación de servicios.
  - `infrastructure/`: Implementaciones técnicas.
    - `database/`: Acceso a datos y repositorios.
    - `server/`: Servidor HTTP, handlers y middleware.

## Descripción Detallada de Archivos

### cmd/api/
- `main.go`: Punto de entrada principal de la aplicación. Inicializa la configuración, base de datos, servicios y el servidor HTTP.

### internal/config/
- `config.go`: Define las estructuras de configuración y métodos para acceder a ella. Incluye configuraciones de la app, base de datos y servidor.
- `environment.go`: Maneja la carga de variables de entorno desde el archivo .env y proporciona valores por defecto.

### internal/core/domain/
- `item.go`: Define el modelo de dominio Item con sus campos, validaciones y métodos de ciclo de vida (BeforeCreate, BeforeUpdate).

### internal/core/ports/
- `repository.go`: Define las interfaces para los repositorios, estableciendo el contrato para operaciones CRUD.
- `service.go`: Define las interfaces para los servicios, estableciendo el contrato para la lógica de negocio.

### internal/core/services/
- `item_service.go`: Implementa la lógica de negocio para items, incluyendo validaciones y reglas de negocio.

### internal/infrastructure/database/postgres/
- `connection.go`: Gestiona la conexión a PostgreSQL usando GORM, configurando el logger y opciones de conexión.
- `item_repository.go`: Implementa las operaciones CRUD para items en PostgreSQL, incluyendo soft deletes y manejo de errores.

### internal/infrastructure/server/
- `server.go`: Configura y gestiona el servidor HTTP, incluyendo graceful shutdown y manejo de señales del sistema.

### internal/infrastructure/server/http/handlers/
- `health_handler.go`: Implementa endpoints de health check y monitoreo del sistema.
- `item_handler.go`: Maneja las peticiones HTTP relacionadas con items, incluyendo validación de entrada y respuestas HTTP.

### internal/infrastructure/server/http/middleware/
- `middleware.go`: Contiene middlewares para:
  - Logging de requests
  - Manejo de errores
  - Timeouts
  - Rate limiting

### internal/infrastructure/server/http/routes/
- `routes.go`: Configura todas las rutas de la API, agrupa endpoints y aplica middleware.

### Archivos Raíz
- `.env`: Archivo de configuración con variables de entorno (no versionado).
- `.env.example`: Ejemplo de configuración de variables de entorno.
- `go.mod`: Gestión de dependencias y versión del módulo.
- `go.sum`: Checksums de dependencias para garantizar integridad.
- `README.md`: Documentación general del proyecto.
- `TECH_DOC.md`: Documentación técnica detallada.
- `LICENSE`: Licencia del proyecto.

## Flujo de Datos

Para entender mejor cómo interactúan estos archivos, aquí un ejemplo del flujo de una petición HTTP:

1. El cliente hace una petición a `/api/v1/items`
2. `routes.go` dirige la petición al handler correspondiente
3. `middleware.go` aplica los middlewares configurados
4. `item_handler.go` procesa la petición
5. `item_service.go` aplica la lógica de negocio
6. `item_repository.go` realiza operaciones en la base de datos
7. La respuesta sigue el camino inverso hasta el cliente

## Patrones de Diseño Utilizados

### En cada capa:

1. **Handlers**:
   - Patrón Decorator (middlewares)
   - Manejo de errores HTTP
   - Validación de entrada

2. **Servicios**:
   - Patrón Strategy para lógica de negocio
   - Validación de dominio
   - Manejo de transacciones

3. **Repositorios**:
   - Patrón Repository
   - Unit of Work (transacciones)
   - Soft Delete

## Arquitectura

El proyecto implementa una Arquitectura Hexagonal (Ports and Adapters) que:

1. Separa la lógica de negocio de las implementaciones técnicas
2. Facilita el testing y el mantenimiento
3. Permite cambiar implementaciones sin afectar el núcleo

### Capas

1. **Domain**: Modelos y reglas de negocio
2. **Ports**: Interfaces que definen los contratos
3. **Services**: Implementación de la lógica de negocio
4. **Infrastructure**: Implementaciones técnicas (DB, HTTP, etc.)

## Configuración

### Variables de Entorno

```env
# App
APP_NAME=api-rest-with-go
APP_VERSION=1.0.0
APP_ENV=development

# Database
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=api_rest_go
DB_SSLMODE=disable

# Server
SERVER_PORT=8080
SERVER_READ_TIMEOUT=10
SERVER_WRITE_TIMEOUT=10
SERVER_IDLE_TIMEOUT=60
```

## Base de Datos

### Modelo Item

```go
type Item struct {
    ID          uint       `json:"id"`
    Name        string     `json:"name" validate:"required,min=3,max=100"`
    Description string     `json:"description" validate:"max=500"`
    Price       float64    `json:"price" validate:"required,gt=0"`
    CreatedAt   time.Time  `json:"created_at"`
    UpdatedAt   time.Time  `json:"updated_at"`
    DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}
```

## API Endpoints

### Health Check

#### GET /health
Verifica el estado de la aplicación.

**Respuesta**:
```json
{
    "status": "ok",
    "info": {
        "database_status": "up",
        "go_version": "go1.23",
        "go_os": "windows",
        "go_arch": "amd64",
        "cpu_cores": 8,
        "goroutines": 10
    }
}
```

#### GET /ping
Endpoint simple para verificar disponibilidad.

**Respuesta**: `pong`

### Items

#### POST /api/v1/items
Crea un nuevo item.

**Body**:
```json
{
    "name": "Nuevo Item",
    "description": "Descripción del item",
    "price": 29.99
}
```

**Respuestas**:
- 201: Item creado exitosamente
- 400: Error de validación
- 500: Error interno del servidor

#### GET /api/v1/items
Obtiene todos los items activos (no eliminados).

**Respuestas**:
- 200: Lista de items
- 500: Error interno del servidor

#### GET /api/v1/items/:id
Obtiene un item por ID.

**Respuestas**:
- 200: Item encontrado
- 404: Item no encontrado o eliminado
- 400: ID inválido
- 500: Error interno del servidor

#### PUT /api/v1/items/:id
Actualiza un item existente.

**Body**: Igual que en POST

**Respuestas**:
- 200: Item actualizado exitosamente
- 400: Error de validación o ID inválido
- 404: Item no encontrado
- 409: Item eliminado (soft deleted)
- 500: Error interno del servidor

#### DELETE /api/v1/items/:id
Realiza un soft delete del item.

**Respuestas**:
- 200: Item eliminado exitosamente
```json
{
    "message": "item deleted successfully",
    "id": 1
}
```
- 404: Item no encontrado
- 400: ID inválido
- 500: Error interno del servidor

## Manejo de Errores

### Códigos HTTP
- 200: Operación exitosa
- 201: Recurso creado exitosamente
- 400: Error de validación o solicitud incorrecta
- 404: Recurso no encontrado
- 409: Conflicto (por ejemplo, intentar actualizar un item eliminado)
- 500: Error interno del servidor

### Formato de Errores
```json
{
    "error": "mensaje descriptivo del error"
}
```

## Soft Delete

El sistema implementa soft delete para los items, lo que significa:

1. Los items no se eliminan físicamente de la base de datos
2. Se marca la fecha de eliminación en el campo `deleted_at`
3. Los items eliminados:
   - No aparecen en las consultas normales
   - No pueden ser actualizados
   - Generan un error específico al intentar acceder a ellos
4. La respuesta del DELETE incluye confirmación y el ID del item eliminado

## Arquitectura del Modelo

### Domain Layer
```go
type Item struct {
    ID          uint           `json:"id" gorm:"primaryKey"`
    Name        string         `json:"name" validate:"required,min=3,max=100" gorm:"not null"`
    Description string         `json:"description" validate:"max=500"`
    Price       float64        `json:"price" validate:"required,gt=0" gorm:"not null"`
    CreatedAt   time.Time      `json:"created_at"`
    UpdatedAt   time.Time      `json:"updated_at"`
    DeletedAt   gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}
```

### Validaciones
1. **Nombre**:
   - Requerido
   - Mínimo 3 caracteres
   - Máximo 100 caracteres

2. **Descripción**:
   - Opcional
   - Máximo 500 caracteres

3. **Precio**:
   - Requerido
   - Debe ser mayor que 0

### Manejo de Tiempo
- `CreatedAt`: Se establece automáticamente en la creación
- `UpdatedAt`: Se actualiza en cada modificación
- `DeletedAt`: Se establece al realizar soft delete

## Middleware

### 1. RequestLogger
Registra información de cada request:
- Método HTTP
- Path
- Código de estado
- Tiempo de respuesta

### 2. ErrorHandler
Manejo centralizado de errores.

### 3. TimeoutMiddleware
Establece un timeout para cada request (10 segundos por defecto).

### 4. RateLimiter
Limita el número de requests por segundo (100 por defecto).

## Monitoreo y Salud

El endpoint `/health` proporciona:
- Estado de la base de datos
- Versión de Go
- Sistema operativo y arquitectura
- Número de núcleos CPU
- Número de goroutines

## Guía de Desarrollo

### Requisitos
- Go 1.21 o superior
- PostgreSQL
- Make (opcional)

### Instalación

1. Clonar el repositorio
2. Copiar `.env.example` a `.env` y configurar
3. Instalar dependencias:
   ```bash
   go mod download
   ```
4. Ejecutar:
   ```bash
   go run cmd/api/main.go
   ```

### Convenciones de Código

1. **Nombres**:
   - Interfaces: Sufijo `er` (ej: `ItemService`)
   - Implementaciones: Minúscula (ej: `itemService`)

2. **Errores**:
   - Usar errores descriptivos
   - Manejar todos los errores
   - No ignorar errores de cierre

3. **Testing**:
   - Tests unitarios para lógica de negocio
   - Tests de integración para APIs
   - Usar mocks para dependencias

### Mejores Prácticas

1. **Base de Datos**:
   - Usar transacciones cuando sea necesario
   - Implementar soft delete
   - Índices para campos frecuentemente consultados

2. **API**:
   - Versionado en la URL
   - Respuestas consistentes
   - Validación de entrada
   - Rate limiting

3. **Seguridad**:
   - No exponer información sensible en logs
   - Validar todas las entradas
   - Usar HTTPS en producción
   - Implementar timeouts

## Notas de Migración

### Migración a Arquitectura Hexagonal

El proyecto ha evolucionado desde una estructura tradicional a una arquitectura hexagonal. Los siguientes cambios se han realizado:

1. **Modelos**:
   - Ubicación anterior: `internal/models/`
   - Nueva ubicación: `internal/core/domain/`
   - Razón: Mejor separación de responsabilidades y aislamiento del dominio

2. **Repositorios**:
   - Ubicación anterior: `internal/repository/`
   - Nueva ubicación: `internal/infrastructure/database/postgres/`
   - Razón: Clara separación entre la interfaz (ports) y la implementación (adapters)

### Beneficios de la Nueva Estructura

1. **Mejor Testabilidad**: 
   - Interfaces claramente definidas
   - Fácil implementación de mocks
   - Pruebas unitarias más limpias

2. **Mayor Flexibilidad**:
   - Facilidad para cambiar implementaciones
   - Mejor separación de responsabilidades
   - Código más mantenible

3. **Claridad en las Dependencias**:
   - Flujo de dependencias de fuera hacia dentro
   - Dominio central protegido
   - Interfaces explícitas 
# Changelog

## [v2.0.0] 

### Cambios Arquitectónicos
- Implementación de arquitectura hexagonal
- Separación clara de capas (domain, ports, infrastructure)
- Mejor organización del código siguiendo principios SOLID

### Nuevas Características
- Sistema de configuración centralizado
- Manejo mejorado de variables de entorno
- Health checks completos
- Middleware optimizado:
  - Rate limiting
  - Timeouts
  - Logging mejorado
  - Manejo de errores centralizado
- Graceful shutdown

### Mejoras Técnicas
- Soft delete implementado correctamente
- Validaciones mejoradas
- Mejor manejo de conexiones a base de datos
- Documentación técnica detallada

## [v1.0.0] 

### Características Iniciales
- Operaciones CRUD básicas
- Conexión PostgreSQL
- Estructura básica del proyecto
- Validaciones simples
- Manejo básico de errores 
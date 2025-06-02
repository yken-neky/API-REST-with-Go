package errors

import (
	"fmt"
	"net/http"
)

// APIError representa un error de la API con contexto
type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

func (e *APIError) Error() string {
	return e.Message
}

// NewNotFoundError crea un error de recurso no encontrado
func NewNotFoundError(resource string, id interface{}) *APIError {
	return &APIError{
		Code:    http.StatusNotFound,
		Message: fmt.Sprintf("%s with id %v not found", resource, id),
	}
}

// NewValidationError crea un error de validación
func NewValidationError(message string) *APIError {
	return &APIError{
		Code:    http.StatusBadRequest,
		Message: message,
	}
}

// NewConflictError crea un error de conflicto
func NewConflictError(message string) *APIError {
	return &APIError{
		Code:    http.StatusConflict,
		Message: message,
	}
}

// NewInternalError crea un error interno del servidor
func NewInternalError(err error) *APIError {
	return &APIError{
		Code:    http.StatusInternalServerError,
		Message: "internal server error",
		Details: err.Error(),
	}
}

// IsNotFound verifica si un error es del tipo NotFound
func IsNotFound(err error) bool {
	if apiErr, ok := err.(*APIError); ok {
		return apiErr.Code == http.StatusNotFound
	}
	return false
}

// IsValidationError verifica si un error es de validación
func IsValidationError(err error) bool {
	if apiErr, ok := err.(*APIError); ok {
		return apiErr.Code == http.StatusBadRequest
	}
	return false
}

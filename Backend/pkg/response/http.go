package responses

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response envelope containing meta information and data payload
type ResponseEnvelope struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data,omitempty"`
}

// Meta contains metadata about the API response
type Meta struct {
	Status  int      `json:"status"`            // HTTP status code
	Message string   `json:"message,omitempty"` // Human-readable message
	Errors  []string `json:"errors,omitempty"`  // List of error messages
	Page    int      `json:"page,omitempty"`    // Pagination page number
	Total   int      `json:"total,omitempty"`   // Total records available
	Locale  string   `json:"locale,omitempty"`  // Localization information
}

// Option type for functional options pattern
type Option func(*Meta)

// Core response handler
func respond(c *gin.Context, status int, data interface{}, opts ...Option) {
	meta := Meta{Status: status}
	for _, opt := range opts {
		opt(&meta)
	}

	c.JSON(status, ResponseEnvelope{
		Meta: meta,
		Data: data,
	})
}

// Success response helper for 200 status
func Success(c *gin.Context, data interface{}, opts ...Option) {
	opts = append([]Option{WithMessage(http.StatusText(http.StatusOK))}, opts...)
	respond(c, http.StatusOK, data, opts...)
}

// Error response helper for generic error handling
func Error(c *gin.Context, status int, err error, opts ...Option) {
	opts = append([]Option{
		WithMessage(http.StatusText(status)),
		WithErrors(err),
	}, opts...)
	respond(c, status, nil, opts...)
}

// Common error helpers
func BadRequest(c *gin.Context, err error, opts ...Option) {
	Error(c, http.StatusBadRequest, err, opts...)
}

func Unauthorized(c *gin.Context, err error, opts ...Option) {
	Error(c, http.StatusUnauthorized, err, opts...)
}

func NotFound(c *gin.Context, err error, opts ...Option) {
	Error(c, http.StatusNotFound, err, opts...)
}

func InternalServerError(c *gin.Context, err error, opts ...Option) {
	Error(c, http.StatusInternalServerError, err, opts...)
}

// Functional options
func WithMessage(message string) Option {
	return func(m *Meta) {
		m.Message = message
	}
}

func WithErrors(errs ...error) Option {
	return func(m *Meta) {
		for _, err := range errs {
			if err != nil {
				m.Errors = append(m.Errors, err.Error())
			}
		}
	}
}

func WithPage(page int) Option {
	return func(m *Meta) {
		m.Page = page
	}
}

func WithTotal(total int) Option {
	return func(m *Meta) {
		m.Total = total
	}
}

func WithLocale(locale string) Option {
	return func(m *Meta) {
		m.Locale = locale
	}
}

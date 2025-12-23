package dgvalidation

import (
	"context"
	"net/http"

	validationContract "github.com/donnigundala/dg-core/contracts/validation"
	"github.com/gin-gonic/gin"
	"github.com/gookit/validate"
)

// Validator is the concrete implementation of validationContract.Validator.
type Validator struct {
	config *Config
}

// Config holds validator configuration
type Config struct {
	DefaultLocale string
	StopOnError   bool
	SkipOnEmpty   bool
}

var _ validationContract.Validator = (*Validator)(nil)

// Option configures the Validator.
type Option func(*Validator)

// WithStopOnError sets whether validation should stop on the first error.
func WithStopOnError(stop bool) Option {
	return func(v *Validator) {
		v.config.StopOnError = stop
	}
}

// WithDefaultLocale sets the default locale.
func WithDefaultLocale(locale string) Option {
	return func(v *Validator) {
		v.config.DefaultLocale = locale
	}
}

// NewValidator creates a new validator instance.
func NewValidator(opts ...Option) validationContract.Validator {
	v := &Validator{
		config: &Config{
			DefaultLocale: "en",
			StopOnError:   false,
			SkipOnEmpty:   true,
		},
	}

	for _, opt := range opts {
		opt(v)
	}

	// Register built-in custom validators
	registerCustomValidators()
	registerDatabaseValidators()

	// Global config for gookit/validate
	validate.Config(func(opt *validate.GlobalOption) {
		opt.StopOnError = v.config.StopOnError
		opt.SkipOnEmpty = v.config.SkipOnEmpty
	})

	return v
}

// ValidateStruct performs validation on a struct.
func (v *Validator) ValidateStruct(ctx context.Context, i interface{}) error {
	val := validate.Struct(i)
	if !val.Validate() {
		return &Error{Errors: val.Errors.All()}
	}
	return nil
}

// --- Gin Integration & FormRequest Support ---

// Validate is the main helper for controllers.
// It binds the request, validates it, and returns true if valid.
// If invalid, it sends a 422 response and returns false.
func Validate(c *gin.Context, req interface{}, scene ...string) bool {
	// 1. Bind data (JSON, Query, or Form)
	if err := c.ShouldBind(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request data",
			"error":   err.Error(),
		})
		return false
	}

	// 2. Create validator from the struct
	v := validate.Struct(req)

	// 3. Apply scene if provided
	if len(scene) > 0 && scene[0] != "" {
		v.SetScene(scene[0])
	}

	// 4. Perform validation
	if !v.Validate() {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Validation failed",
			"errors":  v.Errors.All(),
		})
		return false
	}

	return true
}

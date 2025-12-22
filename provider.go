package dgvalidation

import (
	"github.com/donnigundala/dg-core/contracts/foundation"
)

// ValidationServiceProvider implements the PluginProvider interface.
type ValidationServiceProvider struct {
	// Options can be set before registration to customize the validator
	Options []Option
}

// NewValidationServiceProvider creates a new validation service provider.
func NewValidationServiceProvider(options ...Option) *ValidationServiceProvider {
	return &ValidationServiceProvider{
		Options: options,
	}
}

// Name returns the plugin name.
func (p *ValidationServiceProvider) Name() string {
	return Binding
}

// Version returns the plugin version.
func (p *ValidationServiceProvider) Version() string {
	return Version
}

// Dependencies returns the list of dependencies.
func (p *ValidationServiceProvider) Dependencies() []string {
	return []string{}
}

// Register registers the validation services into the container.
func (p *ValidationServiceProvider) Register(app foundation.Application) error {
	app.Singleton(Binding, func() (interface{}, error) {
		return NewValidator(p.Options...), nil
	})

	return nil
}

// Boot boots the validation services.
func (p *ValidationServiceProvider) Boot(app foundation.Application) error {
	return nil
}

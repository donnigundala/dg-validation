package validation

import (
	"github.com/donnigundala/dg-core/contracts/foundation"
)

// ValidationServiceProvider implements the PluginProvider interface.
type ValidationServiceProvider struct {
	// Options can be set before registration to customize the validator
	Options []Option
}

// Name returns the plugin name.
func (p *ValidationServiceProvider) Name() string {
	return "validation"
}

// Version returns the plugin version.
func (p *ValidationServiceProvider) Version() string {
	return "1.0.0"
}

// Dependencies returns the list of dependencies.
func (p *ValidationServiceProvider) Dependencies() []string {
	return []string{}
}

// Register registers the validation services into the container.
func (p *ValidationServiceProvider) Register(app foundation.Application) error {
	app.Singleton("validator", func() (interface{}, error) {
		return NewValidator(p.Options...), nil
	})

	return nil
}

// Boot boots the validation services.
func (p *ValidationServiceProvider) Boot(app foundation.Application) error {
	return nil
}

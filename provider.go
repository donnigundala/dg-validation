package dgvalidation

import (
	"github.com/donnigundala/dg-core/contracts/foundation"
	"gorm.io/gorm"
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
	return []string{"database"} // Optional dependency for database validators
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
	// Try to resolve database and inject it into validators
	if dbInstance, err := app.Make("database"); err == nil {
		// Use reflection to get the *gorm.DB from the manager if needed
		// For now, we assume the manager has a way to get the default connection.
		// In our dg-database implementation, the manager is the primary binding.

		// Let's check how to get the *gorm.DB.
		// Usually, the manager exports a Connection() method.

		// Actually, a simpler way is to just resolve "database" and check if it's the right type.
		// If it's a *gorm.DB (for single connection setups) or something else.

		// In our dg-database, app.Make("database") returns a *Manager.
		// We can get the default connection from it.

		// For now, let's try a safe approach using an interface check if possible,
		// or just documented expectation.

		// I'll use a hacky but effective way for this framework:
		type dbGetter interface {
			Connection(name ...string) any
		}

		if getter, ok := dbInstance.(dbGetter); ok {
			if conn := getter.Connection(); conn != nil {
				if gormDB, ok := conn.(*gorm.DB); ok {
					SetDB(gormDB)
				}
			}
		}
	}

	return nil
}

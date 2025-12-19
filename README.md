# DG Validation

A production-ready validation service wrapper for Go, based on the excellent `github.com/go-playground/validator/v10`. It provides a structured, extensible way to handle validation in your application with built-in custom validators and localized error messages.

## Features

-   **Standard Validators**: Full access to all `go-playground/validator` tags.
-   **Custom Validators**: Built-in support for common needs (UUID, Slug, Phone, Password strength, etc.).
-   **Localization**: Context-aware error messages with support for multiple locales.
-   **Framework Integration**: includes a `ValidationServiceProvider` for seamless integration with `dg-core` applications.
-   **Structured Errors**: Returns a structured error object containing a map of field-specific messages.

## Installation

```bash
go get github.com/donnigundala/dg-validation
```

## Quick Start

### Basic Usage

```go
package main

import (
	"context"
	"fmt"
	"github.com/donnigundala/dg-validation"
)

type User struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Age   int    `json:"age" validate:"gte=18"`
}

func main() {
	v := validation.NewValidator()
	ctx := context.Background()

	user := User{
		Name:  "John Doe",
		Email: "invalid-email",
		Age:   15,
	}

	err := v.ValidateStruct(ctx, &user)
	if err != nil {
		if vErr, ok := err.(*validation.Error); ok {
			for field, msg := range vErr.Errors {
				fmt.Printf("%s: %s\n", field, msg)
			}
		}
	}
}
```

### Registered Custom Validators

| Tag | Description |
| :--- | :--- |
| `uuid` | Validates that the string is a valid UUID. |
| `slug` | Validates a URL-friendly slug (lowercase, numbers, hyphens). |
| `phone` | Validates basic phone number formats. |
| `password` | Enforces strength: min 8 chars, 1 upper, 1 lower, 1 digit. |
| `username` | Validates username (3-20 chars, alphanumeric, _, -). |
| `alpha_space` | Allows only letters and spaces. |
| `no_sql` | Basic SQL injection pattern detection. |
| `no_xss` | Basic XSS pattern detection. |
| `color_hex` | Validates hex color codes (e.g., #FFF or #FFFFFF). |
| `timezone` | Validates basic timezone strings (e.g., UTC, America/New_York). |

## Advanced Configuration

### Using the Service Provider

In a `dg-core` application, you can register the validation service in your bootstrap process:

```go
import (
    "github.com/donnigundala/dg-validation"
)

// In your app registration
app.Register(&validation.ValidationServiceProvider{
    Options: []validation.Option{
        validation.WithDefaultLocale("en"),
        validation.WithFieldNameTag("json"),
    },
})
```

### Localization

You can provide localized messages for your validation tags:

```go
v := validation.NewValidator(
    validation.WithLocaleMessages("id", map[string]string{
        "required": "harus diisi",
        "email":    "format email tidak valid",
    }),
)

// Use the context to specify the locale
ctx := validation.ToContext(context.Background(), "id")
err := v.ValidateStruct(ctx, &model)
```

## Contributing

See our contributing guidelines for details on how to submit pull requests.

## License

This project is licensed under the MIT License.

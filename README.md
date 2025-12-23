# DG Validation

A powerful, production-ready validation service wrapper for Go, based on `github.com/gookit/validate`. It provides a Laravel-inspired validation experience with "FormRequest" patterns, database-aware rules, and zero-boilerplate integration.

## Features

-   **Standard Validators**: Full access to all `gookit/validate` tags.
-   **Pipes-based Syntax**: Clean `validate:"required|email|minLen:8"` syntax.
-   **FormRequest Pattern**: Support for class-based validation with `Messages()` and `Scene()` methods on request structs.
-   **Database Rules**: Built-in `unique` and `exists` rules that automatically resolve your DB connection.
-   **Zero-Boilerplate Gin Helper**: Single-line validation in controllers with automatic 422 error responses.
-   **Custom Validators**: Built-in support for UUID, Slug, Phone, Password strength, etc.

## Installation

```bash
go get github.com/donnigundala/dg-validation
```

## Quick Start

### Basic Struct Validation

```go
type User struct {
    Name  string `json:"name" validate:"required|minLen:3"`
    Email string `json:"email" validate:"required|email"`
}

v := dgvalidation.NewValidator()
err := v.ValidateStruct(ctx, &user)
```

### FormRequest Pattern (Recommended)

Define your request with rules and custom messages:

```go
type CreateUserRequest struct {
    Name     string `validate:"required|minLen:2" message:"name.required: Name is mandatory"`
    Email    string `validate:"required|email|unique:users,email" message:"email.unique: Email already taken"`
    Password string `validate:"required|password"`
}

func (f CreateUserRequest) Messages() map[string]string {
    return validate.MS{
        "required": "The {field} field is required.",
    }
}
```

Then use it in your Gin controller:

```go
func Store(c *gin.Context) {
    var req CreateUserRequest
    if !dgvalidation.Validate(c, &req) {
        return // 422 Unprocessable Entity sent automatically
    }
    
    // Proceed with validated data...
}
```

## Database Validation Rules

The following rules are available if the `dg-database` plugin is registered:

| Tag | Syntax | Description |
| :--- | :--- | :--- |
| `unique` | `unique:table,column[,ignoreColumn,ignoreValue]` | Checks if value is unique in the table. |
| `unique_multi` | `unique_multi:table,column[,extraColumn,extraValue...]` | Checks uniqueness with extra conditions. |
| `exists` | `exists:table,column[,extraColumn,extraValue...]` | Checks if value exists in the table. |

## Built-in Custom Validators

| Tag | Description |
| :--- | :--- |
| `uuid` | Validates that the string is a valid UUID. |
| `slug` | Validates a URL-friendly slug. |
| `phone` | Validates basic phone number formats. |
| `password` | Enforces strength (8+ chars, upper, lower, digit). |
| `username` | Validates alphanumeric username (3-20 chars). |
| `alpha_space`| Allows only letters and spaces. |
| `no_sql` | Basic SQL injection pattern detection. |
| `no_xss` | Basic XSS pattern detection. |

## Advanced Configuration

### Service Provider

In a `dg-core` application, the validator is automatically registered. You can customize it via options:

```go
app.Register(&dgvalidation.ValidationServiceProvider{
    Options: []dgvalidation.Option{
        dgvalidation.WithStopOnError(true),
        dgvalidation.WithSkipOnEmpty(true),
    },
})
```

## License

This project is licensed under the MIT License.

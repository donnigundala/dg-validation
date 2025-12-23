# DG Validation Documentation

The `dg-validation` package is a comprehensive validation wrapper for the DG Framework, built on top of `gookit/validate`. It provides clean syntax, advanced features like FormRequests, and seamless integration with the framework's database layer.

## Key Features

- **Singleton Registry**: Global access to the validation engine.
- **Advanced Rules**: Built-in support for `uuid`, `slug`, `unique`, `exists`, and more.
- **Pipe Syntax**: Declarative rules using the `|` separator.
- **FormRequest Pattern**: Encapsulate validation logic, scenes, and messages in request structs.
- **Zero-Boilerplate Gin Integration**: Simplified controller validation with automatic error handling.

## Basic Usage

### Defining a Request Struct

```go
type UserRequest struct {
    Email string `validate:"required|email"`
    Age   int    `validate:"required|min:18"`
}
```

### Manual Validation

```go
user := UserRequest{...}
validator := dgvalidation.NewValidator()
if err := validator.ValidateStruct(ctx, &user); err != nil {
    // Handle validation errors (type *dgvalidation.Error)
}
```

## Gin Integration (FormRequest Pattern)

This is the recommended way to handle validation in your controllers. The `Validate` helper automatically binds the request, applies validation rules/scenes, and sends a 422 response if validation fails.

```go
func Store(c *gin.Context) {
    var req CreateUserRequest
    
    // Automatically binds, validates, and responds on failure
    if !dgvalidation.Validate(c, &req) {
        return 
    }
    
    // Success: 'req' is now populated and validated
}
```

## Custom Validators

You can register your own validation logic globally:

```go
validate.AddValidator("is_foo", func(val any) bool {
    s, ok := val.(string)
    return ok && s == "foo"
})
```

Usage in tags:
```go
type MyStruct struct {
    Name string `validate:"required|is_foo"`
}
```

## Database Rules

The `unique` and `exists` rules are automatically enabled when the `dg-database` plugin is booted.

- **Unique**: `unique:table,column[,ignoreColumn,ignoreValue]`
- **Exists**: `exists:table,column[,extraColumn,extraValue]`

# DG Validation Documentation

The `dg-validation` package is a high-level validation wrapper for the DG Framework, built on top of `gookit/validate`. It's designed to provide a Laravel-like "FormRequest" experience, making validation declarative, reusable, and easy to maintain.

---

## 🚀 Quick Start

### 1. Define your Request DTO
Use the `validate` tag with pipe-separated rules.

```go
type CreateUserRequest struct {
    Username string `validate:"required|minLen:3|maxLen:20"`
    Email    string `validate:"required|email"`
    Password string `validate:"required|minLen:8"`
}
```

### 2. Use the Controller Helper
In your Gin controller, use the `Validate` helper for zero-boilerplate handling.

```go
func (ctrl *UserController) Store(c *gin.Context) {
    var req CreateUserRequest
    
    // 1. Binds data from JSON/Form/Query
    // 2. Runs validation
    // 3. Sends 422 Unprocessable Entity if it fails
    if !dgvalidation.Validate(c, &req) {
        return 
    }

    // Success! Data is now in 'req'
}
```

---

## 🛠 Advanced FormRequest Patterns

### Custom Error Messages
You can define field-specific messages directly in the tag or via a `Messages()` method.

```go
type RegisterRequest struct {
    Email string `validate:"required|email" message:"required: Email is mandatory|email: Format is invalid"`
}

// OR use the implementation method for cleaner code
func (f RegisterRequest) Messages() map[string]string {
    return validate.MS{
        "required": "The {field} field cannot be empty.",
        "email":    "Please provide a valid email address.",
    }
}
```

### Validation Scenes
Scenes allow you to use the same struct for different actions (e.g., Create vs Update).

```go
type ProfileRequest struct {
    ID    uint   `validate:"required_if:scene,update"`
    Name  string `validate:"required|minLen:2"`
    Email string `validate:"required|email"`
}

func (f ProfileRequest) Scene() map[string][]string {
    return map[string][]string{
        "update": {"ID", "Name", "Email"},
        "create": {"Name", "Email"},
    }
}

// In Controller:
dgvalidation.Validate(c, &req, "update")
```

---

## 🗄 Database Validation
These rules require the `dg-database` plugin to be enabled.

### 1. `unique` (Uniqueness Check)
Syntax: `unique:table,column[,ignoreColumn,ignoreValue]`

```go
// Basic check
Email string `validate:"required|unique:users,email"`

// Update check (ignore my own ID)
// SELECT count(*) FROM users WHERE email = ? AND id != ?
type UpdateRequest struct {
    ID    uint
    Email string `validate:"required|unique:users,email,id,{ID}"`
}
```

### 2. `exists` (Existence Check)
Syntax: `exists:table,column[,extraColumn,extraValue...]`

```go
// Ensure category exists
CategoryID uint `validate:"required|exists:categories,id"`

// Ensure user exists AND is active
UserID uint `validate:"required|exists:users,id,status,active"`
```

### 3. `unique_multi` (Composite Uniqueness)
Syntax: `unique_multi:table,column[,extraColumn,extraValue...]`

```go
// Ensure a user can only have one unique role in a specific project
// SELECT count(*) FROM project_members WHERE role_id = ? AND user_id = ? AND project_id = ?
type MemberRequest struct {
    ProjectID uint
    UserID    uint
    RoleID    uint `validate:"required|unique_multi:project_members,role_id,user_id,{UserID},project_id,{ProjectID}"`
}
```

---

## 🎨 Custom Validators
You can register global validators for your specific business needs.

```go
func init() {
    validate.AddValidator("is_awesome", func(val any) bool {
        s, ok := val.(string)
        return ok && s == "dg-framework"
    })
}

// Usage
type Framework struct {
    Name string `validate:"required|is_awesome"`
}
```

---

## 🔧 Manual Validation (Standalone)
If you're not using Gin, you can still use the validator manually.

```go
v := dgvalidation.NewValidator()
if err := v.ValidateStruct(ctx, &myStruct); err != nil {
    // 'err' is of type *dgvalidation.Error
    // Access detailed map via err.Errors
}
```

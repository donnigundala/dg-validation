/*
Package validation provides a wrapper around go-playground/validator/v10 with custom rules and helpers.
It simplifies struct validation and provides a singleton validator instance for the application.

# Key Features

  - Singleton validator instance
  - Custom validation rules (e.g., "phone")
  - Struct validation with tags

# Basic Usage

Define a struct with validation tags:

	type UserRequest struct {
	    Email string `validate:"required,email"`
	    Age   int    `validate:"required,min=18"`
	}

Validate an instance:

	user := UserRequest{...}
	if err := validation.ValidateStruct(&user); err != nil {
	    // Handle validation errors
	}

# Custom Validators

Register a custom validation function:

	validation.RegisterValidation("is-foo", func(fl validator.FieldLevel) bool {
	    return fl.Field().String() == "foo"
	})
*/
package dgvalidation

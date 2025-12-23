package dgvalidation

import (
	"regexp"
	"strings"
	"unicode"

	"github.com/google/uuid"
	"github.com/gookit/validate"
)

// RegisterCustomValidators registers all built-in custom validators.
func registerCustomValidators() {
	// UUID validation
	validate.AddValidator("uuid", validateUUID)

	// Slug validation (URL-friendly)
	validate.AddValidator("slug", validateSlug)

	// Phone number validation (basic)
	validate.AddValidator("phone", validatePhone)

	// Password strength validation
	validate.AddValidator("password", validatePassword)

	// Username validation
	validate.AddValidator("username", validateUsername)

	// Alpha with spaces
	validate.AddValidator("alpha_space", validateAlphaSpace)

	// SQL injection prevention (basic)
	validate.AddValidator("no_sql", validateNoSQL)

	// XSS prevention (basic)
	validate.AddValidator("no_xss", validateNoXSS)

	// Hex color validation
	validate.AddValidator("color_hex", validateColorHex)

	// Timezone validation
	validate.AddValidator("timezone", validateTimezone)
}

// validateUUID checks if the value is a valid UUID.
func validateUUID(val any) bool {
	s, ok := val.(string)
	if !ok {
		return false
	}
	_, err := uuid.Parse(s)
	return err == nil
}

// validateSlug checks if the value is a valid URL slug.
func validateSlug(val any) bool {
	slug, ok := val.(string)
	if !ok || slug == "" {
		return false
	}

	slugPattern := `^[a-z0-9]+(?:-[a-z0-9]+)*$`
	matched, _ := regexp.MatchString(slugPattern, slug)
	return matched
}

// validatePhone checks if the value is a valid phone number.
func validatePhone(val any) bool {
	phone, ok := val.(string)
	if !ok || phone == "" {
		return false
	}

	cleaned := strings.Map(func(r rune) rune {
		if unicode.IsDigit(r) || r == '+' {
			return r
		}
		return -1
	}, phone)

	length := len(cleaned)
	if strings.HasPrefix(cleaned, "+") {
		return length >= 11 && length <= 16
	}
	return length >= 10 && length <= 15
}

// validatePassword checks password strength.
func validatePassword(val any) bool {
	password, ok := val.(string)
	if !ok || len(password) < 8 {
		return false
	}

	var (
		hasUpper  = false
		hasLower  = false
		hasNumber = false
	)

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsDigit(char):
			hasNumber = true
		}
	}

	return hasUpper && hasLower && hasNumber
}

// validateUsername checks if the value is a valid username.
func validateUsername(val any) bool {
	username, ok := val.(string)
	if !ok || len(username) < 3 || len(username) > 20 {
		return false
	}

	usernamePattern := `^[a-zA-Z0-9_-]+$`
	matched, _ := regexp.MatchString(usernamePattern, username)
	return matched
}

// validateAlphaSpace allows letters and spaces only.
func validateAlphaSpace(val any) bool {
	value, ok := val.(string)
	if !ok {
		return false
	}
	for _, char := range value {
		if !unicode.IsLetter(char) && !unicode.IsSpace(char) {
			return false
		}
	}
	return true
}

// validateNoSQL performs basic SQL injection prevention.
func validateNoSQL(val any) bool {
	value, ok := val.(string)
	if !ok {
		return false
	}
	value = strings.ToLower(value)

	sqlKeywords := []string{
		"select", "insert", "update", "delete", "drop", "create",
		"alter", "exec", "execute", "union", "declare", "--", "/*", "*/",
		"xp_", "sp_", "0x", "char(", "nchar(", "varchar(", "nvarchar(",
	}

	for _, keyword := range sqlKeywords {
		if strings.Contains(value, keyword) {
			return false
		}
	}

	return true
}

// validateNoXSS performs basic XSS prevention.
func validateNoXSS(val any) bool {
	value, ok := val.(string)
	if !ok {
		return false
	}
	value = strings.ToLower(value)

	xssPatterns := []string{
		"<script", "</script", "javascript:", "onerror=", "onload=",
		"onclick=", "onmouseover=", "<iframe", "<object", "<embed",
		"eval(", "expression(", "vbscript:", "data:text/html",
	}

	for _, pattern := range xssPatterns {
		if strings.Contains(value, pattern) {
			return false
		}
	}

	return true
}

// validateColorHex checks if the value is a valid hex color code.
func validateColorHex(val any) bool {
	color, ok := val.(string)
	if !ok {
		return false
	}
	hexPattern := `^#([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$`
	matched, _ := regexp.MatchString(hexPattern, color)
	return matched
}

// validateTimezone checks if the value is a valid timezone string.
func validateTimezone(val any) bool {
	tz, ok := val.(string)
	if !ok {
		return false
	}

	if tz == "UTC" || tz == "GMT" {
		return true
	}

	tzPattern := `^[A-Z][a-z]+/[A-Z][a-z_]+$`
	matched, _ := regexp.MatchString(tzPattern, tz)
	return matched
}

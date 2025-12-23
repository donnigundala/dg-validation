package dgvalidation

import (
	"fmt"

	"github.com/gookit/validate"
	"gorm.io/gorm"
)

// db is the global gorm instance used for validation.
// It should be injected by the service provider.
var db *gorm.DB

// SetDB sets the database instance for validators.
func SetDB(instance *gorm.DB) {
	db = instance
}

// registerDatabaseValidators registers unique and exists validators.
func registerDatabaseValidators() {
	validate.AddValidator("unique", ValidateUnique)
	validate.AddValidator("exists", ValidateExists)
	validate.AddValidator("unique_multi", ValidateUniqueMulti)
}

// ValidateUnique checks if a value is unique in the database.
// Syntax: unique:table,column[,ignoreColumn,ignoreValue]
// Example: unique:users,email
// Example: unique:users,email,id,1 (ignores user with ID 1)
func ValidateUnique(val any, args ...any) bool {
	if db == nil {
		return true // Skip if DB is not configured
	}

	if len(args) < 2 {
		return true
	}

	table := args[0].(string)
	column := args[1].(string)

	query := db.Table(table).Where(fmt.Sprintf("%s = ?", column), val)

	// Handle ignore column/value
	if len(args) >= 4 {
		ignoreCol := args[2].(string)
		ignoreVal := args[3]
		query = query.Where(fmt.Sprintf("%s != ?", ignoreCol), ignoreVal)
	}

	var count int64
	query.Count(&count)

	return count == 0
}

// ValidateExists checks if a value exists in the database.
// Syntax: exists:table,column[,extraColumn,extraValue]
// Example: exists:categories,id (single condition)
// Example: exists:users,id,status,active (multiple conditions)
// Example: exists:items,id,shop_id,5,is_available,1 (multiple extra filters)
func ValidateExists(val any, args ...any) bool {
	if db == nil {
		return true
	}

	if len(args) < 2 {
		return true
	}

	table := args[0].(string)
	column := args[1].(string)

	query := db.Table(table).Where(fmt.Sprintf("%s = ?", column), val)

	// Handle extra filters (pairs of col,val)
	if len(args) > 2 {
		for i := 2; i < len(args); i += 2 {
			if i+1 < len(args) {
				col := args[i].(string)
				val := args[i+1]
				query = query.Where(fmt.Sprintf("%s = ?", col), val)
			}
		}
	}

	var count int64
	query.Count(&count)

	return count > 0
}

// --- Multi-column Unique (Laravel style) ---

// ValidateUniqueMulti supports multi-column uniqueness.
// Syntax: unique_multi:table,column[,extraColumn,extraValue...]
// Example: unique_multi:project_members,user_id,project_id,5
func ValidateUniqueMulti(val any, args ...any) bool {
	if db == nil {
		return true
	}

	if len(args) < 2 {
		return true
	}

	table := args[0].(string)
	column := args[1].(string)

	query := db.Table(table).Where(fmt.Sprintf("%s = ?", column), val)

	// Handle extra filters (pairs of col,val)
	if len(args) > 2 {
		for i := 2; i < len(args); i += 2 {
			if i+1 < len(args) {
				col := args[i].(string)
				val := args[i+1]
				query = query.Where(fmt.Sprintf("%s = ?", col), val)
			}
		}
	}

	var count int64
	query.Count(&count)

	return count == 0
}

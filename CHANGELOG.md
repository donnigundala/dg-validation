# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.0.0] - 2025-12-27

### Added
- Initial stable release of the `dg-validation` plugin.
- **Form Request Pattern**: Laravel-inspired validation with `FormRequest` interface.
- **Database Validation Rules**: `unique` and `exists` rules with multi-column support.
- **Gin Integration**: `Validate(c, &req)` helper for seamless controller validation.
- **Container Integration**: Auto-registration with Injectable pattern.
- **Advanced Features**:
  - Multi-column unique validation with ignore ID
  - Soft-delete awareness in database rules
  - Custom where clauses for exists/unique rules
  - Fluent rule builder API

### Features
- Built on `gookit/validate` for robust validation
- Type-safe validation with struct tags
- Custom error messages and field names
- Automatic HTTP 422 responses for validation failures
- Database integration via `dg-database` Injectable pattern
- Comprehensive test coverage

### Documentation
- Complete README with examples
- Form Request pattern guide
- Database validation rules documentation

---

## Development History

The following versions represent the development journey leading to v1.0.0:

### 2025-11-25
- Database validation rules (unique, exists)
- Multi-column support with soft-delete awareness

### 2025-11-24
- Form Request pattern implementation
- Gin integration helper
- Initial validation engine

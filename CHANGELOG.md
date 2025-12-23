# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.3.0] - 2025-12-24

### Added
- New `unique_multi` database validation rule for composite uniqueness.
- Comprehensive `docs.md` with real-world examples and advanced patterns.
- Support for multiple extra filters in `exists` and `unique_multi` rules.

### Removed
- Unused `helper.go` legacy functions.

## [1.2.0] - 2025-12-23

### Changed
- Migrated core validation engine from `go-playground/validator` to `gookit/validate`.
- Updated all custom validators to match `gookit/validate` signatures.
- Switched to pipe-separated validation tags (e.g., `required|email`).

### Added
- Laravel-inspired "FormRequest" pattern support in Gin controllers.
- `dgvalidation.Validate` helper for zero-boilerplate controller validation.
- Database-aware rules: `unique` and `exists`.
- Support for validation scenes and custom error messages on request structs.

## [1.0.0] - 2025-12-19

# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.0.0] - 2025-12-19

### Added
- Initial release of the `dg-validation` package.
- Integration with `go-playground/validator/v10`.
- Built-in custom validators: `uuid`, `slug`, `phone`, `password`, `username`, `alpha_space`, `no_sql`, `no_xss`, `color_hex`, `timezone`.
- Context-aware localization support for error messages.
- `ValidationServiceProvider` for `dg-core` integration.
- Structured validation errors.

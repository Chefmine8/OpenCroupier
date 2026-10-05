# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- Project README describing the target architecture, pairing flow and security design.
- Apache License 2.0.
- Contributing guide (workflow, Conventional Commits, code style, security rules).
- `.gitignore` covering Go builds, SQLite databases, TLS keys and secrets, firmware build output and JS tooling.
- Go HTTP/HTTPS server with TLS support, Chi router middleware, and graceful shutdown handling.
- SQLite database initialization with automatic creation of `customer` and `user` tables.
- API endpoints for user and customer creation (`/api/create/createAccount`, `/api/create/createCustomer`) and account value lookup (`/api/gestion/accountValue/{id}`).
- Environment variable loading (`.env`) for server port and host IP configuration.

### Changed
- Server stack is now Go (Chi on `net/http`) instead of Python/FastAPI.
- Code style: replaced the Python conventions with a single rule, one purpose per function.

### Removed
- `ruff` configuration (`pyproject.toml`).

[Unreleased]: https://github.com/Chefmine8/OpenCroupier/commits/main

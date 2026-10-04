# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- Project README describing the target architecture, pairing flow and security design.
- Apache License 2.0.
- Contributing guide (workflow, Conventional Commits, code style, security rules).

### Changed
- Server stack is now Go (Chi on `net/http`) instead of Python/FastAPI.
- Code style: replaced the Python conventions with a single rule, one purpose per function.

### Removed
- `ruff` configuration (`pyproject.toml`).

[Unreleased]: https://github.com/Chefmine8/OpenCroupier/commits/main

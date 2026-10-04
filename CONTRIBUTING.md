# Contributing to OpenCroupier

Thanks for your interest! OpenCroupier is maintained by **Chefmine**
(<ChMi8@pm.me>). Contributions are welcome through issues and pull/merge requests.

> The project is in an early design phase: opening an issue to discuss an idea
> before writing code is strongly encouraged.

## Reporting bugs and requesting features

Open an issue and include:

- what you expected and what happened,
- steps to reproduce,
- the component concerned (server, terminal firmware, dealer UI),
- versions and hardware (Go version, ESP32 board, RFID module).

Do **not** publish security vulnerabilities in public issues. Email the maintainer instead.

## Development workflow

1. Fork the repository (or create a branch if you have access).
2. Create a branch from `main`: `feat/short-description` or `fix/short-description`.
3. Make focused changes; one logical change per pull request.
4. Add or update tests and documentation when relevant.
5. Open a pull/merge request against `main` and describe what and why.

`main` is protected: changes land only through reviewed pull/merge requests.

## Commit messages

Use [Conventional Commits](https://www.conventionalcommits.org/):

```
feat(server): add table assignment endpoint
fix(firmware): ignore repeated card reads for 3 seconds
docs: clarify pairing flow
```

Common types: `feat`, `fix`, `docs`, `refactor`, `test`, `chore`.
Common scopes: `server`, `firmware`, `ui`.

## Code style

- **All components**: no formatting convention is enforced. The only rule is that every function has one precise purpose; do not write functions that do several unrelated things.
- **Firmware (C/C++, Arduino)**: no secrets or credentials in source; keep the loop non-blocking where possible.
- **Dealer UI (JavaScript)**: readable, dependency-light code; large touch-friendly controls.

## Security rules

- Never commit Wi-Fi credentials, TLS keys, tokens or real card UIDs.
- Never put card UIDs or secrets in URLs.
- Keep all terminal ↔ server traffic over HTTPS.

## License

By contributing, you agree that your contributions are licensed under the
[Apache License 2.0](LICENSE), as stated in section 5 of that license.
No CLA or sign-off is required.

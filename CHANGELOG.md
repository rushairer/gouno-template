# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/),
and this project adheres to [Semantic Versioning](https://semver.org/).

## [1.1.0] - 2026-08-20

### Changed

- Delegate `SecurityHeadersMiddleware` to `gouno/middleware.SecurityHeaders` instead of maintaining a custom header-setting implementation (`middleware/middleware.go`).

### Added

- Add unit test suite `tests/web_router_test.go` verifying root (`/`) and health check (`/test/alive`) endpoints.

### Fixed

- Remove unsupported `${VAR}` placeholders from `config/production.yaml` so the default `web` command (production env) can start; sensitive values can now be overridden via `GOUNO_` prefixed environment variables.
- Fix README to match the actual CLI flags (`--config_path`, `--env`) and the `suite`/`task` generator commands.
- Fix the `captcha_type` default key to `captcha.type` in the config manager.
- Remove the ineffective `gouno_env` pflag binding in the config manager.
- Make `NewTestDB` use the configured default database driver instead of hardcoding postgres.
- Remove dead `os.Exit(1)` in `Execute` and a stray newline in the listen error log.
- Use separate `args_bin` entries in `.air.toml`.

## [1.0.2] - 2026-07-29

### Fixed

- Require `gouno` v1.0.2 so generated projects include the response constructors used by the middleware.
- Import `fmt` in the recovery middleware so generated projects compile successfully.

## [1.0.1] - 2026-06-13

### Changed

- Include complete module requirements and checksums so rendered projects can run Go tooling immediately.
- Return configuration load and validation errors from `ConfigManager` instead of exiting inside the config package.
- Add baseline configuration validation for generated projects.
- Strengthen template verification to cover downloaded module checksums.

## [1.0.0] - 2026-05-31

### Added

- Complete DDD project scaffold: `cmd/`, `config/`, `internal/` (domain, repository, service, task), `router/`, `middleware/`, `utility/`.
- Cobra CLI with `web` and `generator` commands.
- Viper multi-environment configuration (`development.yaml`, `test.yaml`, `production.yaml`).
- `ConfigManager` thread-safe configuration singleton.
- Gin web server with graceful shutdown.
- `Makefile` with build, run, dev, test targets.
- `.air.toml` for hot-reload development.
- Code generation templates (`domain.tmpl`, `repository.tmpl`, `service.tmpl`, `controller.tmpl`, `task.tmpl`).
- Bilingual README (English / Chinese).

# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/),
and this project adheres to [Semantic Versioning](https://semver.org/).

## [Unreleased]

### Changed

- Formally name the default template's architecture profile **Flat Layered** and document its intended use for simpler applications, while distinguishing it from Capability Module project/template policies used by more complex applications.

## [1.3.0] - 2026-09-11

### Added

- Add `.gouno/codegen.yaml` using the `gouno.dev/codegen/v1` protocol.
- Ship template-owned codegen sources under `.gouno/codegen/` for controller, domain, repository, service, task, and suite workflows.

### Changed

- The generated project CLI now attaches codegen dynamically from the template manifest; templates without the manifest expose no `gen` command.
- Move concrete generator policy out of the `gouno` core library and into this default project template.
- Remove the legacy top-level `templates/*.tmpl` codegen set so the manifest and `.gouno/codegen/` are the single source of truth.
- Require `github.com/rushairer/gouno` v1.3.0 for the published Codegen v1 runtime.

## [1.2.0] - 2026-08-24

### Changed

- Generated projects now require Go 1.25.0 or newer and use `gouno` v1.2.0; validated by `scripts/verify-template.sh`.
- Template verification now uses the published framework by default. Local framework replacement is available only through explicit `GOUNO_REPLACE_DIR`.

### Security

- Add rendered-project formatting, tidy, race, vet, vulnerability, and lint gates to `scripts/verify-template.sh`.

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

# Project Engineering Contract

This file is intentionally part of the project template and is expected to be inherited by projects created from it. Keep every rule valid both in the template repository and in a generated project.

## Ownership boundary

This project/template owns its application architecture and development conventions. Gouno Core does not prescribe this structure.

The current default template implements the **Flat Layered** reference profile: application code is organized first by global implementation layer, with controller/domain/repository/service/task conventions. Treat those names and boundaries as this template's policy, not as universal Gouno concepts.

Flat Layered is intended as the default reference for simpler applications. Complex applications may instead choose a Capability Module project/template policy where `internal/<capability>/` is the primary ownership boundary and layers live inside the capability. Do not mix the two ownership models casually inside one project.

Do not move template-specific architecture policy into `github.com/rushairer/gouno` merely to make generation easier.

## Codegen source of truth

If `.gouno/codegen.yaml` exists, it is the source of truth for the project Codegen command tree. Concrete generator names, arguments, flags, composition, output paths, and referenced `.gouno/codegen/*.tmpl` files belong to this project/template.

Do not assume a generator exists because an older Gouno release once provided it. Use the manifest that ships with the project.

When changing Codegen policy, update the manifest and its referenced templates together and verify the generated output compiles and follows the project architecture.

`.gouno/codegen.yaml` and `.gouno/codegen/**` are runtime template resources. They must survive project bootstrap verbatim; their Codegen template expressions are evaluated later by Codegen, not by the initial project renderer.

## Optional capability

Codegen is optional in Gouno projects. Removing the manifest intentionally removes the dynamically attached Codegen command. Do not add fallback built-in generators to compensate for a missing manifest.

## Dependency boundary

Use Gouno runtime packages only for reusable mechanisms that the project intentionally depends on. Business/domain architecture remains in the project.

When upgrading Gouno, review runtime/API compatibility and regenerate smoke examples where appropriate. Do not switch to pseudo-versions or local `replace` directives in release-ready project sources.

## Generated-code discipline

Generated code is normal project code after creation. Review and test it like handwritten code.

Do not bypass package boundaries, security controls, error handling, or configuration conventions just because code came from a generator. If generated output repeatedly needs manual correction, fix the template policy or generator template rather than normalizing repeated edits.

## Validation

Before accepting changes that affect project structure or Codegen behavior, verify at minimum:

- the project builds and tests;
- `gouno --help` reflects the capabilities actually declared by the project;
- representative individual generators work;
- composition generators work when declared;
- generated Go files are formatted and compile;
- no template/runtime expression was consumed in the wrong rendering stage.

Repository-specific CI may enforce additional race, vet, lint, and vulnerability checks.

## Toolchain and dependency security baseline

- Generated projects from this default template require Go 1.26.0 or newer. Template CI must validate the latest patch releases of Go 1.26.x and 1.27.x.
- Keep `golang.org/x/net` at v0.60.0 or newer for the October 2026 HTTP/2 advisories. Keep the compatible dependency floors at `golang.org/x/crypto` v0.58.0, `golang.org/x/sys` v0.49.0, and `golang.org/x/text` v0.43.0. Do not add `golang.org/x/sync` solely to mirror the Core dependency list; the rendered-project tidy gate must remain authoritative for the template's actual dependency graph.
- The rendered-project verification script must continue to run tidy, race tests, vet, lint, and `govulncheck`; use `golangci-lint` v2.15.0 or newer for Go 1.27 compatibility. Never suppress reachable findings or downgrade security floors to retain an unsupported Go series.

## Core release compatibility

- The default template's released dependency must be pinned to a published Gouno Core release. Upgrade the Core version only after its immutable tag resolves through the Go module proxy and the exact `go.mod` / `go.sum` checksums are verified.
- For coordinated Core upgrades, run the freshly rendered project verification on Go 1.26.x and 1.27.x, including tidy-without-diff, generator smoke tests, race, vet, lint, and govulncheck. Keep the generated module's dependency graph authoritative; do not add transitive requirements merely to mirror Core.

# Release Checklist

- [x] Confirm the official Gouno Core v1.4.0 tag resolves to the reviewed Core commit and its module checksums match the verified public module.
- [x] Confirm `go.mod` pins `github.com/rushairer/gouno v1.4.0` and `go.sum` contains the verified module and go.mod checksums.
- [x] Confirm the Unreleased changelog is moved into a dated SemVer entry.
- [x] Run the rendered-project verification on the latest Go 1.26.x and 1.27.x patch releases with `actions/setup-go check-latest: true`.
- [x] Confirm rendered-project tidy-without-diff, Codegen individual and composition smoke tests, race, vet, lint, and govulncheck all pass.
- [x] Confirm the project generated from the template builds and the default health route smoke test works.
- [x] Confirm no reachable vulnerability or unreviewed dependency source remains.
- [x] Confirm target tag `v1.4.0` does not exist and will point to the reviewed merge commit; never reuse a tag.
- [x] After pushing the tag, verify `git ls-remote --tags https://github.com/rushairer/gouno-template refs/tags/v1.4.0` resolves to the reviewed tag object/commit.
- [x] Verify `gouno-cli new` can clone the published template tag using `--template-ref v1.4.0` and create a project successfully; the template repository's root `go.mod` is a Go template (`module {{.ModulePath}}`), not a Go module to publish through proxy.golang.org.
- [x] After the tag is public, update the CLI integration test's pinned `--template-ref` from v1.3.0 to v1.4.0 and release the compatible CLI version.
- [x] Update the bilingual Gouno compatibility matrix only after the Template tag is publicly available.


## v1.4.0 evidence (2026-11-05)

- Annotated tag object `75a4d30b3a3eca0dcda339daf1173eaac4c6ef31` peels to validated commit `9f0cba9c8a6be6331f48596c83663efbec5490a2`.
- [Dual-series rendered-project CI](https://github.com/rushairer/gouno-template/actions/runs/41817614153) used Go 1.26.10 / 1.27.3; tidy, individual/suite generation, race (including health-route tests), vet, lint and govulncheck passed. Govulncheck reported zero reachable findings and three uncalled module advisories; those are not a zero-advisory claim.
- Core v1.4.0 resolves to `20efba0af181cb3e03c5bcb199576bfcdec1bd7f` through the public module proxy, with module sum `h1:cCs2eGv6XCJMN4q2RFaFPCL1Ip0PY21RSHp3mT4tAYY=` and go.mod sum `h1:S2DmlM4BlX3BLoeRw7s66aABDJDexP4Gj/xov28MmvA=` matching this template.
- Public pinned Bootstrap passed with the real CLI, automatic tidy, byte-for-byte `.gouno` comparison, health-route race tests and vet. [CLI integration](https://github.com/rushairer/gouno-cli/actions/runs/42010113946) additionally passed both live and published pinned-template generation on both Go series.
- [GitHub Release](https://github.com/rushairer/gouno-template/releases/tag/v1.4.0) exists. The bilingual compatibility update is committed in gouno-doc `18c8773d862cb7f4c2e32d0eba69572f1e6df075`.

This evidence update follows publication; it does not move the release tag.

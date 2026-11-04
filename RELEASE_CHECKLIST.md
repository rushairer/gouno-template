# Release Checklist

- [ ] Confirm the official Gouno Core v1.4.0 tag resolves to the reviewed Core commit and its module checksums match the verified public module.
- [ ] Confirm `go.mod` pins `github.com/rushairer/gouno v1.4.0` and `go.sum` contains the verified module and go.mod checksums.
- [ ] Confirm the Unreleased changelog is moved into a dated SemVer entry.
- [ ] Run the rendered-project verification on the latest Go 1.26.x and 1.27.x patch releases with `actions/setup-go check-latest: true`.
- [ ] Confirm rendered-project tidy-without-diff, Codegen individual and composition smoke tests, race, vet, lint, and govulncheck all pass.
- [ ] Confirm the project generated from the template builds and the default health route smoke test works.
- [ ] Confirm no reachable vulnerability or unreviewed dependency source remains.
- [ ] Confirm target tag `v1.4.0` does not exist and will point to the reviewed merge commit; never reuse a tag.
- [ ] After pushing the tag, verify `git ls-remote --tags https://github.com/rushairer/gouno-template refs/tags/v1.4.0` resolves to the reviewed tag object/commit.
- [ ] Verify `gouno-cli new` can clone the published template tag using `--template-ref v1.4.0` and create a project successfully; the template repository's root `go.mod` is a Go template (`module {{.ModulePath}}`), not a Go module to publish through proxy.golang.org.
- [ ] After the tag is public, update the CLI integration test's pinned `--template-ref` from v1.3.0 to v1.4.0 and release the compatible CLI version.
- [ ] Update the bilingual Gouno compatibility matrix only after the Template tag is publicly available.

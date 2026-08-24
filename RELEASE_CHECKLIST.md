# Release Checklist

- [ ] The rendered project passes `bash scripts/verify-template.sh` on Go 1.25.x and 1.26.x.
- [ ] `CHANGELOG.md` has a dated release entry and target tag/release do not exist.
- [ ] The worktree is clean and `git diff --check` passes.
- [ ] The reusable CI workflow SHA is a released `gouno-doc` commit.

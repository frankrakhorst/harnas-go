# Definition of Done

A task/PR is done when all of the following hold:

- [ ] Code compiles (`go build ./...`) and `gofmt`/`golangci-lint` pass with no warnings.
- [ ] Tests added/updated for new behavior; `go test ./...` passes; coverage ≥ 80%.
- [ ] Todo file `todos/<ticket-key>.md` checkboxes updated (`- [x]`).
- [ ] Branch named `feature-<issue-key>-<description>`; PR title links issue key.
- [ ] One concern per PR; no unrelated changes.
- [ ] No new dependencies without approval; stdlib preferred.
- [ ] No secrets committed; config via env/flags with validation.
- [ ] Exported symbols have doc comments; errors wrapped with context.
- [ ] README/CHANGELOG updated for user-facing changes.
- [ ] Lint and tests green in CI before merge.
- [ ] User approval obtained before commit/push.
- [ ] Pull request merged to main.
# Agents

## Rules

- Never modify AGENTS.md unless the user explicitly asks and confirms the change.
- Work to be done is recorded in `todos/<ticket-key>.md` using Markdown checkboxes (`- [ ]`); mark items `- [x]` when done.
- Branches are named `feature-<issue-key>-<description>`."
- Commit and push must always be confirmed/approved by the user before executing.
- Ask questions as soon as possible.
- Put long-lasting tasks in the background.
- Follow `gofmt` and `golangci-lint`; keep code style consistent.
- Every new feature must include tests; maintain coverage ≥80%.
- Use Conventional Commits (`feat:`, `fix:`, `chore:`, ...).
- One concern per PR; link the issue key in title/description.
- No new dependencies without approval; prefer the standard library.
- Never commit secrets; use environment variables or a secret manager.
- Update README/CHANGELOG for user-facing changes.
- Wrap and return errors with context (`fmt.Errorf("...: %w", err)`); never swallow them; no panics in library code.
- Document goroutine ownership; avoid global mutable state.
- Split files when they become hard to navigate.
- Define interfaces where they are consumed, not where implemented.
- Pass `context.Context` through public functions; respect cancellation.
- Use structured logging (slog); no `fmt.Println` in production code.
- Load config via env/flags with sensible defaults; validate at startup.
- Version public APIs; don't break without deprecation.
- All PRs must pass lint and tests before merge.
- Exported symbols require doc comments.
- Validate and sanitize all external input; defend against injection.
- Use least-privilege credentials and rotate secrets regularly.
- Keep dependencies updated; address CVEs promptly.
- Use `time.Time` with location; avoid `time.Now()` in tests (inject a clock).
- Make zero values useful; design for the zero value.
- Avoid `init()` unless registration requires it.
- Keep business logic out of transport/HTTP handlers.

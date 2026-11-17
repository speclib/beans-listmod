# Tasks

## 1. Ancestor suggestion

- [x] 1.1 Extend the suggested-parent helper to walk the parent chain and choose the nearest open milestone, falling back to the nearest open epic, skipping closed ancestors; verify with table-driven tests in `internal/tui/list_test.go`
- [x] 1.2 Pass all beans and implicit statuses from the list `c` action into the helper; verify with the list create test and `go test ./internal/tui`

## 2. Integration

- [x] 2.1 Run `go vet ./...` and `go test ./...`

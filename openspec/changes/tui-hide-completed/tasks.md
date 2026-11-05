# Tasks

## 1. Completed filter helper

- [x] 1.1 Add an `excludeCompleted(beans []*bean.Bean, implicitStatuses map[string]string) []*bean.Bean` helper in `internal/tui/list.go` that drops beans with status `completed` or implicit status `completed` and keeps everything else
- [x] 1.2 Add table-driven tests for the helper in `internal/tui/list_test.go` covering explicit `completed`, implicit `completed`, `scrapped` by status, implicit `scrapped`, active beans, and a missing implicit entry; verify with `go test ./internal/tui -run TestExcludeCompleted`

## 2. Toggle and interface state

- [x] 2.1 Add `hideCompleted bool` to `listModel` and a `v` case in the key switch, guarded by `FilterState() != list.Filtering`, that flips the flag and returns `m.loadBeans`; verify in `mise run beans-tui` that `v` reloads the list and typing `v` inside `/` filters the text instead
- [x] 2.2 Apply `excludeCompleted` to `filteredBeans` in `loadBeans` after implicit statuses are computed and before `BuildTree`, leaving `allBeans` unfiltered for tree context; verify completed beans disappear when hidden and scrapped beans remain
- [x] 2.3 Extract the title logic into one helper used by `View` and `ViewConstrained`, appending `[completed hidden]` when hidden, and add a unit test for the title combinations; verify with `go test ./internal/tui -run TestListTitle`
- [x] 2.4 Append `v show completed` to the footer help in every footer branch while hidden, and add the `v` line to the help overlay in `internal/tui/help.go`; verify the footer hint and `?` overlay in `mise run beans-tui`

## 3. Integration check

- [x] 3.1 Run `go test ./...` and confirm all packages pass; if `internal/web/dist` is missing, create it with `mkdir -p internal/web/dist && touch internal/web/dist/.gitkeep` first
- [x] 3.2 Smoke test `mise run beans-tui` with at least one completed, one scrapped, and one implicit-completed bean: toggle `v` twice, check title and footer, open help, and confirm tree context and tag filtering still work

## Workflow follow-up

- Archive the change with `openspec archive tui-hide-completed --yes` after review.

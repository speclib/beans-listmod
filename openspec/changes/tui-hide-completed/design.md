# Design

## Context

See proposal.md for motivation and the `tui-bean-list` spec for required behavior.

The TUI list loads beans in `internal/tui/list.go`. `loadBeans` queries the resolver with the optional tag filter, queries all beans separately for tree context, computes implicit statuses with `resolver.Core.ImplicitStatus`, then calls `ui.BuildTree` and `ui.FlattenTree`. Statuses are configured in `pkg/config/config.go`; `completed` and `scrapped` are the terminal statuses, and `pkg/beancore/links.go` resolves an implicit status to the nearest terminal ancestor.

## Goals / Non-Goals

**Goals:**

- Keep the toggle local to the list model, with no backend, GraphQL, config, or persistence changes.
- Reuse the implicit statuses already computed for rendering so hidden beans match what the list displays.
- Keep the filter logic in a pure helper that is easy to test.

**Non-Goals:**

- Changing which statuses exist or how implicit status is computed.
- Persisting the toggle across sessions or applying it to the web UI, board, backlog, or CLI.
- Changing tag filtering, sorting, or tree building.

## Decisions

### Session state on `listModel`

Add `hideCompleted bool` to `listModel`, default `false`. The `v` case is handled in the existing `tea.KeyMsg` switch, guarded by `FilterState() != list.Filtering` like the other shortcuts. Toggling returns `m.loadBeans` as the command so the list rebuilds from the resolver.

Alternative: persist the setting in `.beans.yml`. Rejected because the request is a quick, session-only view toggle, and persistence adds config surface and cross-surface expectations.

### Filter after implicit status computation

In `loadBeans`, compute implicit statuses as today, then pass `filteredBeans` through a helper before `BuildTree`:

```go
func excludeCompleted(beans []*bean.Bean, implicitStatuses map[string]string) []*bean.Bean
```

A bean is excluded when its own status is `completed` or its implicit status is `completed`. `scrapped` is not excluded. `allBeans` stays unfiltered so `BuildTree` can still attach ancestors as dimmed context.

Alternative: filter inside `ui.BuildTree`. Rejected because tree building is shared and should keep its current contract; the visibility policy belongs to the list model.

### Title and footer state

Extract a `title()` method used by both `View` and `ViewConstrained` to avoid a third copy of the title logic. It returns `Beans`, adds `[tag: x]`, and adds `[completed hidden]` when hidden.

When hidden, append `v show completed` to the footer help in every branch (selection, active filter, default) so the hint survives tag filtering and multi-select. The help overlay gains a `v` line. No transient status message is set, because `statusMessage` replaces the footer help and would hide the hint.

### Tests

Add table-driven tests for `excludeCompleted`: explicit `completed`, implicit `completed`, `scrapped` by status, implicit `scrapped`, active beans, and a missing implicit entry. This covers the behavior without driving the Bubbletea model.

## Risks / Trade-offs

- [Completed ancestor appears dimmed] A hidden completed parent can still appear when a visible scrapped child is under it, because `BuildTree` adds ancestors as context. Accepted: the tree stays readable and the bean is clearly non-matching.
- [Cursor position after toggle] Rebuilding the list matches the tag filter behavior; the cursor may land on a different bean. Accepted for a session toggle.
- [Footer width] The extra hint lengthens an already dense footer on narrow terminals. The hint only appears while hidden, and the help overlay remains the complete reference.
- [Name collision with a future CLI/web toggle] Keeping the state in the list model avoids implying that other surfaces share it.

## Migration Plan

None. The feature is additive and defaults to current behavior.

## Open Questions

None.

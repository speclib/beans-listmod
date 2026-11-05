# Proposal

## Why

The TUI bean list shows completed beans among active work. Once a bean is done it adds noise to daily browsing, and the list has no way to set it aside without leaving the TUI or changing bean status.

## What Changes

- Add a `v` shortcut to the TUI bean list that hides or shows completed beans.
- Hide beans with status `completed` and beans whose implicit status is `completed` (a terminal ancestor). Scrapped beans stay visible.
- Keep the toggle session-only and start with completed beans visible, so behavior is unchanged until the user presses `v`.
- Show the state: `[completed hidden]` in the list title, a `v show completed` footer hint, and a `v` line in the help overlay.
- Ignore `v` while the list filter is active, matching other single-key shortcuts.
- Leave filtering, sorting, tree building, backend queries, and persistence untouched.

## Capabilities

### New Capabilities

- `tui-bean-list`: visibility controls and keyboard behavior of the bean list in the terminal UI.

### Modified Capabilities

<!-- No existing specs. -->

## Impact

- `internal/tui/list.go`: toggle state, filtering in `loadBeans`, title and footer text.
- `internal/tui/help.go`: help overlay entry.
- `internal/tui/list_test.go`: tests for the completed-bean filter.
- No API, GraphQL, configuration, or on-disk changes. No new dependencies.

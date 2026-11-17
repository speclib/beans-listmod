# Proposal

## Why

Creating a bean from the list with `c` ignores where the cursor is, so the new bean always starts without a parent. When the cursor sits on a milestone or epic, that row is almost always the intended parent, and the user has to create the bean and then reassign the parent by hand. The create popup also gives no way to set or clear a parent before the bean exists.

## What Changes

- Pressing `c` on an open milestone or epic preselects that bean as the parent of the new bean. Closed rows and every other type create without a parent.
- The create popup gains a parent row showing the parent's type and full title, or `(none)` when unset. The title input and the parent row are the two focusable fields, cycled with `tab` and `shift+tab`; the title keeps focus on open.
- `enter` on the parent row opens the existing parent picker stacked over the popup. `esc` returns with the parent unchanged, a selection returns focus to the title, and the picker's `(No Parent)` entry clears the parent.
- The picker filters candidates for the new bean's type, which remains the configured default (`task`), so milestones, epics and features are eligible.
- Footer hints follow the focused field.
- New beans are still created as `draft`, with the parent set in the same create mutation.

Out of scope: type selection in the create popup, cursor preselection in the web UI, and changing the create flow outside the list view.

## Capabilities

### New Capabilities

- `tui-create-bean`: creating a bean from the TUI list, including how the cursor row preselects a parent and how the create popup displays, changes and removes that parent.

### Modified Capabilities

None.

## Impact

- `internal/tui/list.go`: the `c` action passes the cursor row as a suggested parent.
- `internal/tui/createmodal.go`: parent row, focus handling and picker entry.
- `internal/tui/tui.go`: create message carries the parent, picker returns to the popup, background renders the stacked modal.
- `internal/tui/parentpicker.go`: reused unchanged apart from how it is invoked for a not-yet-created bean.
- `internal/tui/help.go`: shortcut text for the popup if needed.
- Tests in `internal/tui`.
- No API, storage or configuration changes.

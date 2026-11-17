# Proposal

## Why

Preselection only works when the cursor is on a milestone or epic row. Creating a bean from a child inside a milestone is the common case, and it still starts without a parent even though the milestone is visible in the tree above the cursor.

## What Changes

- A descendant row now preselects its nearest open milestone ancestor, or the nearest open epic ancestor when no milestone exists.
- A milestone or epic row keeps preselecting itself.
- Closed ancestors are skipped, including ancestors closed implicitly by an ancestor of their own.
- Rows with no container ancestor still start without a parent.

## Capabilities

### New Capabilities

None.

### Modified Capabilities

- `tui-create-bean`: the cursor row suggestion now also resolves container ancestors for descendant rows.

## Impact

- `internal/tui/list.go`: the suggested-parent helper walks the parent chain.
- `internal/tui/list_test.go`: ancestor cases.
- No API, storage or configuration changes.

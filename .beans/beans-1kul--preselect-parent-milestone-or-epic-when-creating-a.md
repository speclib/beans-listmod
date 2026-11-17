---
# beans-1kul
title: Preselect parent milestone or epic when creating a bean from the list
status: in-progress
type: feature
priority: normal
created_at: 2026-11-17T18:17:12Z
updated_at: 2026-11-17T19:29:56Z
---

When creating a bean from the list view with `c` while the cursor is on a milestone or epic, that bean should become the parent of the new bean.

## Behavior

- Cursor on a `milestone` or `epic` row that is not closed: preselect that bean as the parent of the new bean.
- Cursor on a closed row (`completed` or `scrapped`, including implicit statuses), or on any other type: create without a parent.
- The create popup shows the preselected parent, so the user sees where the bean will land before pressing enter.
- New beans are still created with status `draft`; the parent is set in the same create mutation.
- A cursor on a `feature` does not preselect it as parent.

## Create popup

The popup has two focusable fields: the title input and a parent row.

- Title input has focus when the popup opens.
- `tab` cycles title, parent, title. `shift+tab` reverses.
- The parent row shows the parent type and its full title, wrapped once when the title is long. Without a parent it shows `(none)` in muted text.
- `enter` on the title creates the bean with the current parent.
- `enter` on the parent row opens the existing parent picker stacked over the popup. `esc` returns with the parent unchanged, and a selection returns focus to the title.
- The picker filters parents for the new bean's type, which is the configured default (`task` today), so milestones, epics and features are eligible.
- The picker's `(No Parent)` entry removes the parent.
- Footer hints track focus: `enter create  tab parent  esc cancel` on the title, `enter change  tab title  esc cancel` on the parent row.

The typed title survives the picker because the create modal stays in the app model while the picker is open. The background renderer for pickers must be able to show the create modal over the list, so the picker can stack two overlays deep.

## OpenSpec change

`preselect-parent-on-create`

## Out of scope

The web UI create form keeps its existing parent dropdown without cursor preselection. That is a separate change.

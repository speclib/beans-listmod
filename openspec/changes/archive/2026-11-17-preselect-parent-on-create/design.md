# Design

## Context

See proposal.md for motivation and the `tui-create-bean` spec for required behavior.

The list key handler currently returns an empty `openCreateModalMsg{}` for `c`, so the cursor row never reaches the create modal. `createModalModel` holds a single `textinput.Model`, sends `beanCreatedMsg{title}` on enter, and has no parent state. The app model already keeps `createModal` while a modal state is active, and `overlayModal` can dim any rendered background, so a picker can be stacked over the popup without a special renderer.

`parentPickerModel` already handles eligibility, cycle prevention, sorting, a `(No Parent)` entry, and preselection of a current parent. It is built from bean IDs, a display title, bean types and a current parent.

## Goals / Non-Goals

**Goals:**

- Carry the cursor row's suggested parent into the create popup without an extra lookup in the list model.
- Keep the typed title intact while the parent picker is open.
- Reuse the existing parent picker and its eligibility rules.

**Non-Goals:**

- Adding a type field to the create popup.
- Changing how parents are set for existing beans.
- Web UI preselection.

## Decisions

### The create message carries the suggested parent

`openCreateModalMsg` gains a `parentID string` field. The list key handler fills it only when the selected item is an open milestone or epic, using both the bean's own status and the item's `implicitStatus`. Closed rows and other types pass an empty string. This keeps the closed-parent rule in the list view, where implicit status is already computed, instead of teaching the create modal about tree ancestry.

Alternative considered: let the create modal ask the resolver for a parent later. Rejected because the cursor's implicit status is list state, not bean state available to the modal.

### The modal owns the parent ID and resolves its label

`createModalModel` gains `parentID`, a `resolver`, and a focus field with `focusTitle` and `focusParent` values. `View` resolves the parent ID to its type and title for display. Resolution is in-memory, so rendering stays cheap, and one source of truth serves both the preselected parent and one chosen from the picker.

Alternative considered: pass type and title through every message. Rejected because it duplicates state and can drift from the picker's selection.

### The picker is stacked over the popup

Opening the picker from the popup sets `previousState` to `viewCreateModal`, so `closeParentPickerMsg` and `parentSelectedMsg` return to the popup naturally. `parentSelectedMsg` branches when `previousState` is `viewCreateModal`: it records the parent ID and returns to the popup instead of calling `SetParent`, since the bean does not exist yet. `getBackgroundView` gains a `viewCreateModal` case that renders the popup over the list, which lets `overlayModal` stack the picker two levels deep.

Alternative considered: a dedicated create-parent message and picker mode. Rejected because it would duplicate the picker's existing return handling for no behavior gain.

### The picker is built for the new bean's default type

The popup opens the picker with no bean IDs, the typed title as the display title, and one bean type from `cfg.GetDefaultType()`. With no IDs there are no selected beans and no descendants to exclude, so every bean of a valid parent type is offered. `(No Parent)` clears the parent through the existing entry.

### Focus is explicit and limited to two fields

`tab` and `shift+tab` cycle between title and parent. `enter` on the title creates, `enter` on the parent row opens the picker, and the picker returns focus to the title. The footer hints are rendered from the focus value. No other key is intercepted, so typing a title is unaffected.

Alternative considered: a `ctrl+p` shortcut with no focus model. Rejected because the parent row would not be discoverable and removing a parent would stay hidden in the picker.

### Long parent titles wrap once

The parent row renders `Parent  <type> <full title>` and lets lipgloss wrap it within the modal width, accepting a one-line growth in modal height. Truncation was rejected because parent titles are often the only way to tell similar milestones apart.

## Risks / Trade-offs

- **Two stacked overlays dim the create popup**, so focus inside it is less visible while the picker is open. Accepted because the picker is modal and the popup returns unchanged.
- **The picker's item list is built once per open** and does not live-refresh if beans change while it is open. This matches existing picker behavior and the window is short.
- **Empty title still closes the popup on enter**, as today. The focus model must not turn that into a silent create with no title.
- **A typed title passed to the picker can be long**, which affects the picker subtitle only; eligibility is unaffected.

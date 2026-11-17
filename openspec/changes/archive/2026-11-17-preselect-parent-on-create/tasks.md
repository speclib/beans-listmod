# Tasks

## 1. Suggested parent from the cursor row

- [x] 1.1 Add a helper in `internal/tui/list.go` that returns the cursor item's bean ID only when its type is `milestone` or `epic` and neither its status nor implicit status is completed or scrapped, and unit-test milestone, epic, closed, implicit-closed and other-type cases with `go test ./internal/tui -run TestSuggestedParent`
- [x] 1.2 Extend `openCreateModalMsg` with a parent ID and make the list `c` action fill it from the helper, verifying the new field with a unit test on the key handling and `go test ./internal/tui`

## 2. Parent row and focus in the create popup

- [x] 2.1 Give `createModalModel` a parent ID, a resolver and a focus value, pass the preselected parent into `newCreateModalModel`, and unit-test that the title starts focused and that `tab` and `shift+tab` cycle the focus
- [x] 2.2 Render the parent row with its type, full title and `(none)` fallback using the modal width, and update the footer hints from the focus value; verify with unit tests on the rendered view
- [x] 2.3 Make `enter` on the title create and `enter` on the parent row emit the existing `openParentPickerMsg` with the typed title and the configured default type, keeping empty-title enter closing the popup; verify with unit tests on `Update`

## 3. Stacked parent picker

- [x] 3.1 Render the create popup as the picker background by adding a `viewCreateModal` case to `getBackgroundView`, and verify the parent picker returns to the popup on `esc` with the title intact
- [x] 3.2 Record the chosen parent in the popup when `parentSelectedMsg` arrives while `previousState` is `viewCreateModal`, returning focus to the title instead of calling `SetParent`; verify with a unit test that selection updates the displayed parent and cancellation leaves it unchanged

## 4. Persist the parent

- [x] 4.1 Add the parent ID to `beanCreatedMsg` and pass it as `CreateBeanInput.Parent` in the create handler, keeping status `draft`; verify with a unit test on the message flow and `go test ./internal/tui`
- [x] 4.2 Run `go vet ./...` and `go test ./...` and confirm the whole suite passes

## Workflow follow-up

- Archive the change after the project's review requirements are satisfied.
- Update `.beans` with the implementation outcome.

# Spec Delta

## Purpose

Defines how creating a bean from the terminal UI list uses the cursor row as a suggested parent, and how the create popup lets the user confirm, change or clear that parent before the bean exists.

## ADDED Requirements

### Requirement: The cursor row suggests a parent

When the user starts creating a bean from the list while the cursor is on an open milestone or epic row, the create popup SHALL start with that bean as the parent. On a closed row, on any other bean type, or on an empty list, the popup SHALL start without a parent.

#### Scenario: Milestone row preselects the milestone

- **WHEN** the user presses `c` with the cursor on a milestone whose status is not completed or scrapped
- **THEN** the create popup opens with that milestone as the parent

#### Scenario: Epic row preselects the epic

- **WHEN** the user presses `c` with the cursor on an epic whose status is not completed or scrapped
- **THEN** the create popup opens with that epic as the parent

#### Scenario: Closed row does not preselect

- **WHEN** the user presses `c` with the cursor on a completed or scrapped milestone or epic
- **THEN** the create popup opens without a parent

#### Scenario: Other types do not preselect

- **WHEN** the user presses `c` with the cursor on a task, bug or feature row
- **THEN** the create popup opens without a parent

### Requirement: The create popup shows the parent

The create popup SHALL show the selected parent's type and full title. When no parent is selected it SHALL show `(none)`. A long parent title SHALL wrap once rather than truncate.

#### Scenario: Preselected parent is visible

- **WHEN** the create popup opens with a parent
- **THEN** it shows the parent's type and full title before the user presses enter

#### Scenario: No parent is visible as none

- **WHEN** the create popup opens without a parent
- **THEN** it shows `(none)` in muted text

### Requirement: The parent can be changed or removed before creating

The create popup SHALL make the title input and the parent row focusable, with the title focused on open. `tab` SHALL cycle to the next field and `shift+tab` to the previous one. `enter` on the parent row SHALL open a parent picker over the popup, and the picker's `(No Parent)` entry SHALL clear the parent.

#### Scenario: Change the parent

- **WHEN** the user focuses the parent row, presses enter, and chooses another eligible bean
- **THEN** the popup shows the chosen bean as parent

#### Scenario: Cancel the picker

- **WHEN** the user opens the parent picker and presses `esc`
- **THEN** the popup returns with its previous parent unchanged
- **AND** focus returns to the title

#### Scenario: Remove the parent

- **WHEN** the user focuses the parent row, opens the picker, and chooses `(No Parent)`
- **THEN** the popup shows `(none)` and the bean is created without a parent

#### Scenario: Picker only offers eligible parents

- **WHEN** the parent picker opens from the create popup
- **THEN** it lists only beans whose type may parent the new bean's type, which is the configured default type

#### Scenario: Hints follow the focused field

- **WHEN** focus is on the title
- **THEN** the popup shows `enter create`, `tab parent` and `esc cancel`
- **AND WHEN** focus is on the parent row
- **THEN** the popup shows `enter change`, `tab title` and `esc cancel`

### Requirement: The chosen parent is saved with the new bean

Creating the bean SHALL persist the parent shown in the popup in the same create operation. New beans SHALL keep status `draft`.

#### Scenario: Created with a parent

- **WHEN** the user presses enter on the title with a parent shown
- **THEN** the new bean is created with status `draft` and that parent

#### Scenario: Created without a parent

- **WHEN** the user presses enter on the title with `(none)` shown
- **THEN** the new bean is created with status `draft` and no parent

# Spec Delta

## Purpose

Controls which beans the terminal UI list shows, including a user-toggleable filter for completed work.

## ADDED Requirements

### Requirement: Completed beans can be hidden from the list

When completed beans are hidden, the TUI bean list SHALL omit beans whose status is `completed` and beans whose implicit status is `completed`. Beans whose status or implicit status is `scrapped` SHALL remain visible.

#### Scenario: Explicitly completed bean is hidden

- **WHEN** the user hides completed beans
- **THEN** a bean with status `completed` no longer appears in the list

#### Scenario: Bean with completed ancestor is hidden

- **WHEN** the user hides completed beans and a bean inherits an implicit status of `completed` from an ancestor
- **THEN** that bean no longer appears in the list

#### Scenario: Scrapped beans remain visible

- **WHEN** the user hides completed beans
- **THEN** beans with status `scrapped` or an implicit status of `scrapped` still appear in the list

### Requirement: The `v` key toggles completed visibility

The list view SHALL toggle the hidden state of completed beans when the user presses `v`. The list SHALL ignore `v` while its filter is active. The toggle SHALL start with completed beans visible and SHALL NOT persist across sessions.

#### Scenario: First press hides, second press shows

- **WHEN** the user presses `v` with completed beans visible
- **THEN** completed beans are hidden
- **AND WHEN** the user presses `v` again
- **THEN** completed beans are visible again

#### Scenario: Filter input receives the key

- **WHEN** the list filter is active and the user types `v`
- **THEN** the filter text receives `v` and completed visibility does not change

#### Scenario: A new session starts visible

- **WHEN** the user starts the TUI in a new session
- **THEN** completed beans are visible regardless of the previous session's toggle state

### Requirement: Hidden state is visible in the interface

While completed beans are hidden, the list title SHALL include `[completed hidden]`, the footer SHALL include a `v show completed` hint, and the help overlay SHALL describe `v` as hiding or showing completed beans.

#### Scenario: Title marks the hidden state

- **WHEN** completed beans are hidden
- **THEN** the list title includes `[completed hidden]`

#### Scenario: Footer offers the way back

- **WHEN** completed beans are hidden
- **THEN** the footer includes `v show completed`

#### Scenario: Help documents the shortcut

- **WHEN** the user opens the help overlay
- **THEN** it lists `v` with a description of hiding or showing completed beans

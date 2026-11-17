# Spec Delta

## MODIFIED Requirements

### Requirement: The cursor row suggests a parent

When the user starts creating a bean from the list, the create popup SHALL start with the cursor row's container as parent: an open milestone or epic row itself, otherwise the nearest open milestone ancestor, otherwise the nearest open epic ancestor. On a closed row, a row with no such ancestor, or an empty list, it SHALL start without a parent.

#### Scenario: Milestone row preselects the milestone

- **WHEN** the user presses `c` with the cursor on a milestone whose status is not completed or scrapped
- **THEN** the create popup opens with that milestone as the parent

#### Scenario: Epic row preselects the epic

- **WHEN** the user presses `c` with the cursor on an epic whose status is not completed or scrapped
- **THEN** the create popup opens with that epic as the parent

#### Scenario: Child inside a milestone preselects the milestone

- **WHEN** the user presses `c` with the cursor on a bean whose parent chain contains an open milestone
- **THEN** the create popup opens with that milestone as the parent

#### Scenario: Milestone wins over an intermediate epic

- **WHEN** the parent chain is a task under a feature under an epic under a milestone
- **THEN** the create popup opens with the milestone as the parent

#### Scenario: Epic is the fallback container

- **WHEN** the parent chain contains an open epic and no milestone
- **THEN** the create popup opens with that epic as the parent

#### Scenario: Closed ancestors are skipped

- **WHEN** the nearest ancestor is completed or scrapped, explicitly or implicitly
- **THEN** the popup uses the next open container ancestor, or no parent when none remains

#### Scenario: No container ancestor

- **WHEN** the cursor is on a bean with no milestone or epic ancestor
- **THEN** the create popup opens without a parent

#### Scenario: Closed row does not preselect

- **WHEN** the user presses `c` with the cursor on a completed or scrapped milestone or epic
- **THEN** the create popup opens without a parent

#### Scenario: Other types do not preselect

- **WHEN** the user presses `c` with the cursor on a task, bug or feature row whose parent chain has no milestone or epic
- **THEN** the create popup opens without a parent

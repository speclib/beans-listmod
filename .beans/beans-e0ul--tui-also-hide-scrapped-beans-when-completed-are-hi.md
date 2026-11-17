---
# beans-e0ul
title: 'TUI: also hide scrapped beans when completed are hidden'
status: in-progress
type: task
priority: normal
created_at: 2026-11-13T14:57:50Z
updated_at: 2026-11-17T19:18:57Z
---

The v toggle in the TUI list hides completed beans but keeps scrapped ones. Scrapped beans should be hidden too.

- [x] Update tui-hide-completed OpenSpec change
- [x] Hide scrapped (explicit and implicit) in the list filter
- [x] Update title, footer, help text
- [x] Update tests

## Summary of Changes

- Renamed excludeCompleted to excludeClosed; it now drops completed and scrapped beans (explicit and implicit)
- Title shows [completed/scrapped hidden], footer and help updated
- Tests cover scrapped and draft cases
- Updated the tui-hide-completed OpenSpec change (proposal, spec, design, tasks)

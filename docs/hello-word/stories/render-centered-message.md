# Story — Render centered message

## User story
As a Visitor, I want to see the stored message centered on the page, so that the single-screen proof matches the approved design.

## In scope
- One public page only.
- Fetch the message value from backend data already served for this module.
- Render the backend message centered horizontally and vertically on plain white background with black text.
- Show only the message text, with no navigation, no extra sections, and no animation.

## Out of scope
- Message editing, admin UI, or any other interaction.
- Additional pages, routes, menus, or chrome.
- Any styling beyond plain white background, black text, and centering.
- Changing how the message is stored or fetched; that belongs to the backend/storage story.

## UI scope
- Single screen: the approved single-message page.
- States covered: default, empty, loading, and error.
- Default state shows the backend-provided message centered in the viewport.
- Empty and error states show no stale or fallback text.

## Acceptance criteria
1. Given backend returns `Hello Word`, when visitor opens page, then page shows `Hello Word`.
2. Given page renders on a standard viewport, when visitor views page, then message is centered both horizontally and vertically.
3. Given page renders, when visitor inspects styling, then background is white and text is black.
4. Given page renders, when visitor checks page structure, then no navigation, extra page content, or animation exists.
5. Given backend returns empty string or fails, when page loads, then page shows no stale or fallback message.

## Dependencies
- `Store and serve message` must land first.
- Backend message endpoint must be available.
- PostgreSQL seed row must exist with the canonical message text `Hello Word`.

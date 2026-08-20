# Story — Store and serve message

## User story
As a Visitor, I want one stored message value to be served by backend and shown on page, so frontend is not hardcoded.

## In scope
- One PostgreSQL row containing canonical message text `Hello Word`.
- Backend read path that returns stored message for page use.
- Frontend fetches backend data and renders only that message.
- Empty, missing-row, and backend-failure handling required by SRS.

## Out of scope
- Message editing or admin UI.
- Navigation, extra routes, or extra pages.
- Any styling beyond plain white screen with black text and centering.
- Animation or decorative UI.

## UI scope
- Single message page only.
- States covered: default, loading, empty, and error.
- Use approved design system: one full-viewport shell, centered H1-style message, no extra chrome.
- No other screens or interactions.

## Acceptance criteria
1. Given database has exactly one stored message row with non-empty text, when backend reads message, then backend returns `Hello Word` unchanged.
2. Given frontend requests page data, when backend response contains message, then response includes stored text and not frontend fallback copy.
3. Given page loads normally and stored row is present, when visitor opens page, then visitor sees one message value sourced from backend.
4. Given stored message text is empty, when backend processes record, then backend rejects empty content and page does not render blank text.
5. Given stored message row is missing, when backend or frontend requests message, then system returns empty-data error state and shows no stale text.
6. Given PostgreSQL is unavailable, when backend requests data, then backend returns error state and frontend shows no partial or stale message.
7. Given viewport is small or large, when page renders, then message stays centered horizontally and vertically without horizontal scroll caused by layout.
8. Given backend returns message, when visitor inspects page, then page shows plain white background and black text only, with no navigation, extra sections, or animation.

## Dependencies
- PostgreSQL available.
- Backend read path implemented before frontend can render real content.
- Canonical seed value remains `Hello Word` unless SRS changes.
- Story depends on approved design and design system for page layout rules.

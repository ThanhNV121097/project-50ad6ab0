# Test Cases — Store and serve message

Risk level: low. One public stored message, but tests still cover read path, empty content, missing row, and Postgres outage because SRS names those failure states.

## Acceptance criteria coverage

- AC-1: backend returns `Hello Word` when database has stored message row.
- AC-2: backend response contains stored text, not hardcoded frontend copy, when frontend requests page data.
- AC-3: visitor sees one message value sourced from backend when page loads normally and stored row is present.
- Invalid input: backend rejects empty stored message content and page does not render blank text.
- Boundary: backend serves one printable line of normal length unchanged.
- Not found: backend returns empty-data error state and frontend shows no stale text when row is missing.
- Upstream failure: PostgreSQL unavailable returns error state and frontend shows no content rather than partial message.
- Not permitted: not applicable; page is public.
- Conflict: not applicable; only one current value is shown.

## Scenarios

**Scenario**: Store row returns stored text
**Given** database has one stored message row with text `Hello Word`
**When** backend reads message
**Then** backend returns `Hello Word`

**Scenario**: Frontend uses backend message, not fallback copy
**Given** frontend requests page data and backend response contains stored text `Hello Word`
**When** page data is fetched
**Then** response includes stored text and does not use hardcoded frontend copy

**Scenario**: Page shows backend message on normal load
**Given** page loads normally and stored row is present
**When** visitor opens page
**Then** visitor sees one message value sourced from backend

**Scenario**: Empty stored message is rejected
**Given** stored message text is empty
**When** backend reads message
**Then** backend rejects empty content and page does not render blank text

**Scenario**: Normal printable line is served unchanged
**Given** stored message text is one printable line of normal length
**When** backend reads message
**Then** backend serves it unchanged

**Scenario**: Missing row shows empty-data error state
**Given** message row is missing
**When** backend reads message
**Then** backend returns empty-data error state and frontend shows no stale text

**Scenario**: PostgreSQL outage shows no partial message
**Given** PostgreSQL is unavailable
**When** backend tries to read message
**Then** backend returns error state and frontend shows no content rather than partial message

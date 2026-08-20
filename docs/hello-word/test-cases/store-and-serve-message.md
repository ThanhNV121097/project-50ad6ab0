# Test Cases — Store and serve message

Risk level: low. Scope is one stored public message, one backend read path, one frontend fetch path. Coverage focuses on acceptance criteria and named failure behaviors.

## Automated cases

### Scenario: backend returns stored message row
**Given** database has exactly one stored message row with text `Hello Word`
**When** backend reads message
**Then** backend returns `Hello Word`

Trace: HELLO-WORD-001 AC-1

### Scenario: frontend uses backend message, not hardcoded copy
**Given** frontend requests page data and backend response contains stored message `Hello Word`
**When** page data is fetched
**Then** response includes stored text and does not replace it with frontend fallback copy

Trace: HELLO-WORD-001 AC-2

### Scenario: page shows backend-provided message on normal load
**Given** page loads normally and stored row is present with text `Hello Word`
**When** visitor opens page
**Then** visitor sees one message value, `Hello Word`, sourced from backend

Trace: HELLO-WORD-001 AC-3

### Scenario: backend rejects empty stored message text
**Given** stored message text is empty
**When** backend tries to read or serve message
**Then** backend rejects empty content and page renders no blank text

Trace: HELLO-WORD-001 failure case: Invalid input

### Scenario: backend serves one printable line unchanged
**Given** stored message text is one printable line of normal length
**When** backend reads message
**Then** backend returns same text unchanged

Trace: HELLO-WORD-001 boundary case: Boundary

### Scenario: missing row returns empty-data error state
**Given** message row is missing
**When** backend reads message
**Then** backend returns empty-data error state and frontend shows no stale text

Trace: HELLO-WORD-001 failure case: Not found

### Scenario: PostgreSQL unavailable returns error state
**Given** PostgreSQL is unavailable
**When** backend tries to read message
**Then** backend returns error state and frontend shows no content rather than partial message

Trace: HELLO-WORD-001 failure case: Upstream failure

## Manual case

### Scenario: public page has no permission gate
**Given** actor is Visitor opening public page
**When** visitor requests page
**Then** page is accessible without login and shows stored message

Trace: HELLO-WORD-001 role: Visitor; failure case: Not permitted is not applicable because page is public

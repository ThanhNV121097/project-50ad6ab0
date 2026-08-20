# Test Cases — Render centered message

Risk level: low. One-screen read-only page, so focus on content source, centering, styling, and failure states named in SRS.

## AC coverage

### Scenario: Show backend message
**Given** backend returns `Hello Word` and page loads normally
**When** visitor opens page
**Then** page shows `Hello Word` and not frontend fallback copy
**Traceability**: HELLO-WORD-002 AC-1

### Scenario: Center message in viewport
**Given** backend returns `Hello Word` and page renders on standard viewport
**When** visitor views page
**Then** message is centered both horizontally and vertically
**Traceability**: HELLO-WORD-002 AC-2

### Scenario: Use plain white background and black text
**Given** backend returns `Hello Word` and page renders
**When** visitor inspects page styling
**Then** background is white and text is black only
**Traceability**: HELLO-WORD-002 AC-3

### Scenario: No navigation, extra content, or animation
**Given** backend returns `Hello Word` and page renders
**When** visitor checks page structure
**Then** page contains no navigation, extra sections, extra pages, or animation
**Traceability**: HELLO-WORD-002 AC-4

## Named failure and boundary cases

### Scenario: Empty backend message shows no text
**Given** backend returns empty string
**When** visitor opens page
**Then** page shows no message text and no placeholder copy
**Traceability**: HELLO-WORD-002 failure: Invalid input

### Scenario: Centering holds on small and large viewport
**Given** backend returns `Hello Word`
**When** visitor views page at 320px wide and at a large desktop viewport
**Then** message stays centered with no horizontal page scroll caused by layout
**Traceability**: HELLO-WORD-002 failure: Boundary

### Scenario: Backend message unavailable shows empty state
**Given** backend message request fails
**When** visitor opens page
**Then** page shows no stale or partial message and remains plain white
**Traceability**: HELLO-WORD-002 failure: Upstream failure

### Scenario: No edit path exists
**Given** visitor opens page
**When** visitor looks for navigation or editing UI
**Then** no edit path, admin UI, or nonexistent permission gate exists
**Traceability**: HELLO-WORD-002 failure: Not permitted

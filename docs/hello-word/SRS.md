# SRS — hello-word

Module: `hello-word`
Last updated: 2025-08-14
Design: [View the approved design](http://localhost:8080/design/50ad6ab0-a40c-48ab-995a-90edbbaadd21)
Design system: `design/design-system.md`

> One file per module, at `docs/hello-word/SRS.md`. It covers only the functions
> that belong to this module. Never write `docs/SRS.md`.

## 1. Purpose

`hello-word` exists to show one stored message on one page. It proves the app can
read content from PostgreSQL, serve it through the backend, and render it in the
browser with no extra navigation, chrome, or decorative UI. If this module does
not exist, there is no product behavior beyond an empty shell.

## 2. Actors

| Actor | Who they are | What they may do in this module |
|---|---|---|
| Visitor | Any person opening the page | View the single public page and receive the stored message |
| System | Backend, database, and frontend working together | Store the message, fetch it, and render it centered |

## 3. Scope

**In scope** — the functions specified below, by their plan titles:

- Store and serve message
- Render centered message

**Out of scope** — name what a reader would reasonably expect here and say
where it lives instead. This section prevents the same argument twice.

- Navigation, menus, routes, or extra pages — deliberately not built; project is one-screen only.
- Message editing or admin UI — belongs to no current module and is not part of requested proof.
- Styling beyond horizontal and vertical centering on plain white background with black text — deliberately not built.

## 4. Functional requirements

### 4.1 Store and serve message

**Requirement HELLO-WORD-001 — Store one message row**

*As a* Visitor, *I want to* receive message text from stored data, *so that* page content is not hardcoded in frontend.

Behaviour:

1. On startup, the system has exactly one stored message record available for read access.
2. The stored message contains the text `Hello Word`.
3. The frontend does not supply its own fallback text when the stored message is available.

**Acceptance criteria** — each maps one-to-one onto a test case in `docs/hello-word/test-cases/store-and-serve-message.md`. Given/When/Then, no compound conditions: one behaviour per criterion.

| # | Given | When | Then |
|---|---|---|---|
| AC-1 | database has stored message row | backend reads message | backend returns `Hello Word` |
| AC-2 | frontend requests page data | backend response contains message | response includes stored text, not hardcoded frontend copy |
| AC-3 | page loads normally | stored row is present | visitor sees one message value sourced from backend |

**Failure, boundary and permission behaviour** — the part most often skipped
and most often the source of bugs. Every row needs a defined outcome; "should
not happen" is not an outcome.

| Case | Condition | Expected behaviour |
|---|---|---|
| Invalid input | stored message text is empty | backend rejects empty content and page does not render blank text |
| Boundary | stored message text is one printable line of normal length | backend serves it unchanged |
| Not found | message row is missing | backend returns empty-data error state and frontend shows no stale text |
| Not permitted | actor lacks permission to view public page | not applicable; page is public |
| Conflict | two writers change same message | last write wins for stored row, because only one current value is shown |
| Upstream failure | PostgreSQL unavailable | backend returns error state; frontend shows no content rather than partial message |

**Data touched** — the fields this function reads and writes, in product terms.
The physical schema is TL's job in `docs/architecture/erd.md`; this is the list
that document has to satisfy.

| Field | Type | Required | Rule |
|---|---|---|---|
| message text | text | yes | exactly one current row; content must be non-empty and render as plain text |

### 4.2 Render centered message

**Requirement HELLO-WORD-002 — Center message on plain screen**

*As a* Visitor, *I want to* see the stored message centered on the page, *so that* the page matches the one-screen proof.

Behaviour:

1. The page shows only the backend-provided message text.
2. The message is horizontally centered and vertically centered in the viewport.
3. The page uses plain white background and black text only.
4. The page has no navigation, no extra sections, and no animation.

**Acceptance criteria** — each maps one-to-one onto a test case in `docs/hello-word/test-cases/render-centered-message.md`. Given/When/Then, no compound conditions: one behaviour per criterion.

| # | Given | When | Then |
|---|---|---|---|
| AC-1 | backend returns `Hello Word` | visitor opens page | page shows `Hello Word` |
| AC-2 | page renders on standard viewport | visitor views page | message is centered both horizontally and vertically |
| AC-3 | page renders | visitor inspects styling | background is white and text is black |
| AC-4 | page renders | visitor checks page structure | no navigation, extra page content, or animation exists |

**Failure, boundary and permission behaviour** — the part most often skipped
and most often the source of bugs. Every row needs a defined outcome; "should
not happen" is not an outcome.

| Case | Condition | Expected behaviour |
|---|---|---|
| Invalid input | backend returns empty string | page renders no message text rather than placeholder copy |
| Boundary | viewport is small or large | centering still holds without horizontal scroll caused by layout |
| Not found | backend message is unavailable | page shows empty-state failure, not a guessed fallback |
| Not permitted | actor attempts to reach nonexistent editing UI | not applicable; no edit path exists |
| Conflict | backend value changes while page is open | next fetch can show new value; page never merges multiple values |
| Upstream failure | backend request fails | page shows no stale or partial message and remains plain white |

**Data touched** — the fields this function reads and writes, in product terms.
The physical schema is TL's job in `docs/architecture/erd.md`; this is the list
that document has to satisfy.

| Field | Type | Required | Rule |
|---|---|---|---|
| message text | text | yes | displayed exactly as returned, with no extra labels or decoration |

## 5. Screens

The design is the source of truth for appearance; this section maps functions
onto it so nothing in the design is unaccounted for and nothing specified here
is missing from the design.

| Screen | Section in the design | Functions it serves | States that must exist |
|---|---|---|---|
| Single message page | Approved design preview | HELLO-WORD-001, HELLO-WORD-002 | default, loading, empty, error |

## 6. Non-functional requirements

Only what is real for this module. Delete rows that do not apply rather than
inventing a number nobody will check.

| Area | Requirement |
|---|---|
| Performance | Page content renders within 2 seconds on a typical connection after backend response is available |
| Accessibility | Message is readable with 4.5:1 contrast or better, and page works without mouse input |
| Responsive | Works at 320px wide and up without horizontal page scroll |
| Privacy | No personal data is stored; only one public message string is persisted |

## 7. Dependencies and assumptions

- **Depends on:** PostgreSQL, for storing the single message row.
- **Depends on:** Backend read path, for serving stored message to frontend.
- **Assumption:** `Hello Word` remains the canonical initial message unless the stakeholder later changes the product copy; if false, stored seed text changes and the SRS must be revised.

| Open question | Proposed default | Who decides |
|---|---|---|
| Should missing message row be treated as empty-data error or auto-seeded on startup? | Auto-seeded on startup so page always has one row | Stakeholder / TL |

## 8. Traceability

Every plan item in this module appears exactly once, and every requirement id
traces to a test case. A gap in this table is a gap in the build.

| Plan item | Requirement ids | Test cases |
|---|---|---|
| Store and serve message | HELLO-WORD-001 | `test-cases/store-and-serve-message.md` |
| Render centered message | HELLO-WORD-002 | `test-cases/render-centered-message.md` |

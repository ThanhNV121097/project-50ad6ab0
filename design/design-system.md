# Design System — hello-word

> Source of truth: the approved `index.html` (preview: approved design).
> Every value below is extracted from it. Changing a value here without changing the approved design is a defect.

Last updated: 2025-02-14

## 1. Foundations

### 1.1 Color

Semantic tokens. Name by job, never by hue.

| Token | Value | Used for |
|---|---|---|
| `--color-bg` | `#ffffff` | Page background |
| `--color-text` | `#000000` | Body text |
| `--color-surface` | `#ffffff` | Surface behind centered content; same as page background in this design |
| `--color-border` | `#000000` | Not used in approved design |
| `--color-text-muted` | `#000000` | Not used in approved design |
| `--color-primary` | `#000000` | Not used in approved design |
| `--color-primary-text` | `#ffffff` | Not used in approved design |
| `--color-success` | `#000000` | Not used in approved design |
| `--color-warning` | `#000000` | Not used in approved design |
| `--color-danger` | `#000000` | Not used in approved design |
| `--color-focus` | `#000000` | Not used in approved design |

#### Contrast audit

Every text-on-background pair actually used. Body text ≥ 4.5:1, large text (≥ 18.66px bold or ≥ 24px) ≥ 3:1, UI borders ≥ 3:1.

| Foreground | Background | Ratio | Passes |
|---|---|---|---|
| `--color-text` | `--color-bg` | `21:1` | AA / AA Large |
| `--color-text` | `--color-surface` | `21:1` | AA / AA Large |

### 1.2 Spacing

Base unit: `4px`. Every margin, padding, and gap in the product uses one of these.

| Token | Value |
|---|---|
| `--space-1` | `4px` |
| `--space-2` | `8px` |
| `--space-3` | `12px` |
| `--space-4` | `16px` |
| `--space-6` | `24px` |
| `--space-8` | `32px` |
| `--space-12` | `48px` |

### 1.3 Typography

Font families (include the fallback stack and how the font is loaded):

- Body: `Arial, Helvetica, sans-serif`
- Headings: `Arial, Helvetica, sans-serif`
- Mono: not used

| Token | Size | Line height | Weight | Used for |
|---|---|---|---|---|
| `--text-xs` | not used | not used | not used | not used |
| `--text-sm` | not used | not used | not used | not used |
| `--text-base` | not used | not used | not used | not used |
| `--text-lg` | not used | not used | not used | not used |
| `--text-xl` | not used | not used | not used | not used |
| `--text-2xl` | not used | not used | not used | not used |
| `--text-3xl` | `clamp(2.5rem, 8vw, 6rem)` | `1` | `400` | Main message |

Heading levels are used in order and never skipped for visual sizing. This design uses only one heading level.

### 1.4 Radius, border, shadow, motion

| Token | Value | Used for |
|---|---|---|
| `--radius-sm` | not used | not used |
| `--radius-md` | not used | not used |
| `--radius-lg` | not used | not used |
| `--radius-full` | not used | not used |
| `--border-width` | not used | not used |
| `--shadow-sm` | not used | not used |
| `--shadow-md` | not used | not used |
| `--shadow-lg` | not used | not used |
| `--duration-fast` | not used | not used |
| `--duration-base` | not used | not used |
| `--easing` | not used | not used |

Motion respects `prefers-reduced-motion: reduce`: no motion exists in this design.

### 1.5 Layout and breakpoints

| Name | Min width | Container | Columns | Gutter |
|---|---|---|---|---|
| `sm` | not used | not used | not used | not used |
| `md` | not used | not used | not used | not used |
| `lg` | not used | not used | not used | not used |
| `xl` | not used | not used | not used | not used |

Z-index scale (only these values are allowed):

| Layer | Value |
|---|---|
| Base | `0` |
| Sticky header | not used |
| Dropdown | not used |
| Modal backdrop | not used |
| Modal | not used |
| Toast | not used |

## 2. Components

One subsection per reusable component. Every component lists **all** states.

### 2.1 Page shell

**Purpose** — Full viewport centering shell for the single message. Use only for this one-screen app; do not reuse for other layouts.

**Anatomy** — `[body] [main] [h1]`

**Variants**

| Variant | Tokens | When to use |
|---|---|---|
| default | `--color-bg`, `--color-text`, `--text-3xl` | Only screen in product |

**Sizes**

| Size | Height | Padding | Text token |
|---|---|---|---|
| default | `100vh` | `0` | `--text-3xl` |

**States** — every row must be filled in.

| State | Visual change | Tokens |
|---|---|---|
| Default | White page, black centered message | `--color-bg`, `--color-text`, `--text-3xl` |
| Hover | No hover state | none |
| Focus (keyboard) | No interactive focus target in shell | none |
| Active / pressed | No active state | none |
| Disabled | No disabled state | none |
| Loading | No loading state | none |
| Error | No error state | none |
| Empty | Empty view is same as default because message is the only content | `--color-bg`, `--color-text` |

**Accessibility** — semantic `main` landmark and one `h1`; no interactive controls, so no keyboard interaction or hit-target rule applies.

## 3. Content and formatting

- Voice and tone in one line: plain, neutral, no marketing copy.
- Date, time, number, and currency formats: not used.
- Capitalization rule for buttons, headings, and labels: title case not required; message keeps exact product text `Hello Word`.
- Empty-state and error-message wording pattern: not used.

## 4. Known deviations

Places where the approved design does not follow its own rules or the anti-patterns in `references/ai-defaults.md`. Record, do not silently fix.

| Where | Deviation | Why it stands | Follow-up |
|---|---|---|---|
| Foundations / color | `--color-surface`, `--color-primary`, and other semantic tokens are listed as not used, because design has only background and text | Single-screen product uses only two colors | None |
| Foundations / typography | Only one heading size exists and it is an H1-sized message, so heading ramp does not apply fully | Approved design has one heading only | None |
| Components | No interactive components exist, so hover/focus/active/disabled/loading/error states are mostly N/A | Product scope has no controls | None |

## 5. Change log

| Date | Change | Design PR |
|---|---|---|
| 2025-02-14 | Initial design system extracted from approved single-screen mockup | pending |

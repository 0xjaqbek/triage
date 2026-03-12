# Help Bubbles Toggle — Design Document

**Date:** 2026-03-12
**Approach:** CSS-only bubbles with JS toggle (Approach 1 — no dependencies)

## Overview

A toggle switch below the "Rozpocznij Triage" button (and in the header after event starts) that shows/hides chat-bubble-style help overlays on key UI elements. Bubbles overlap the elements they describe and can be individually dismissed with ✕.

**Audience:** Both first-time users/experiment participants and field paramedics needing a refresher.

## Toggle

- **Setup screen:** Below the "Rozpocznij Triage" button
- **After event starts:** In the header area (near Info button), always accessible
- **Styled as:** Slider switch with label (translated via i18n)
- **Persistence:** `localStorage` key `triageHelpOn` (boolean) — only toggle state is remembered
- **Dismissed bubbles:** Tracked in-memory only (JS Set). NOT persisted across sessions
- **Toggle OFF → ON:** Resets all dismissals, all bubbles reappear
- **New session:** All bubbles show if toggle is ON

## Bubble Styling

- Background: `var(--surface)` with border `var(--border2)`
- Speech-bubble tail/arrow pointing toward the described element
- `✕` close button in top-right corner
- Text color: `var(--text2)` (muted, doesn't compete with main UI)
- `z-index: 50` (above content, below header z:100 and modals)
- Positioned absolutely, overlapping the target element
- Max-width: ~280px, font-size: 12-13px
- Responsive: full-width with no arrow on very small screens
- Animation: CSS opacity fade-in/fade-out on show/dismiss

## Bubble Inventory (15 bubbles)

### Setup Screen (4 bubbles)

| ID | Target | Content |
|---|---|---|
| `help-lang` | Language selector | "Select language" displayed in ALL 7 languages simultaneously |
| `help-event-name` | Event name input | What to type here — incident name, location |
| `help-start-btn` | Rozpocznij Triage button | Starts the event and begins timing |
| `help-toggle` | Help toggle itself | Explains what the toggle does (show/hide help) |

### Segregacja View (7 bubbles)

| ID | Target | Content |
|---|---|---|
| `help-stats-bar` | Header stats bar (T1-T4 badges) | What each color count represents |
| `help-nav` | Navigation tabs | Switch between the 3 main sections |
| `help-wizard` | Triage wizard box | Follow the START algorithm step by step to classify a patient |
| `help-result` | Result card (after triage) | Shows classification result; add notes, confirm, or override color manually |
| `help-patient-list` | Patient list ("Pacjenci na miejscu") | Triaged patients currently waiting at scene |
| `help-transport-list` | Transport section ("W transporcie") | Patients currently being transported to hospitals |
| `help-delivered-btn` | Dostarczony button | Marks patient as delivered to hospital. The EMS team becomes available again and reappears in the dispatch list for a new assignment |

### Dysponowanie View (3 bubbles)

| ID | Target | Content |
|---|---|---|
| `help-dispatch-form` | Dispatch form (3 selects + button) | Select patient, EMS team, and hospital to dispatch transport |
| `help-manage-teams` | Teams panel | Add/remove EMS teams available on scene |
| `help-manage-hospitals` | Hospitals panel | Add/remove destination hospitals |

### Raport View (1 bubble)

| ID | Target | Content |
|---|---|---|
| `help-report` | Report section | Auto-generated event summary; can copy text or send via SMS/email |

## i18n

- All bubble texts translated in all 7 languages via the existing `TRANSLATIONS` object and `data-i18n` system
- **Exception:** `help-lang` bubble shows "Select language" in all 7 languages simultaneously (not translated based on current language — since the user may not understand the current one)

## Technical Implementation

- **CSS:** `.help-bubble` class with positioning, arrow pseudo-element, fade transition. Bubbles visible only when `body.help-on` class is present AND the bubble hasn't been dismissed
- **JS:** Toggle adds/removes `body.help-on`. Dismissed bubble IDs stored in a `Set`. Toggle OFF→ON clears the Set. `applyStaticTranslations()` extended to also translate bubble texts
- **HTML:** Each bubble is a `<div class="help-bubble" data-help-id="help-xxx">` placed inside or adjacent to its target element's container
- **No external dependencies** — fits the single-file PWA architecture

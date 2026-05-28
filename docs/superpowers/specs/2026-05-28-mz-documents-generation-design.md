# Design: MZ Document Generation (Zał. 12, 13, 14)

**Date:** 2026-05-28
**Status:** Approved
**Scope:** `index.html` only — single-file offline PWA

---

## Overview

After ending an MCI event (`eventEndTime` is set), the user can generate the 3 official post-incident documents required by the Polish Ministry of Health Procedura v2.3 (Zał. 12, 13, 14) as a single downloadable HTML file, ready to print as PDF.

---

## User Flow

1. User sets `eventEndTime` in the Report tab (existing functionality)
2. Button **"📄 Generuj dokumenty MZ"** appears in the Report tab, below the end-event button
3. User clicks → supplement modal opens
4. User fills in missing fields and/or skips sections
5. User clicks **"Generuj i pobierz"**
6. File `dokumenty-mci-YYYY-MM-DD.html` is downloaded
7. User opens file in browser → prints to PDF or paper

---

## Supplement Modal

A single scrollable modal with 5 collapsible sections. Each section has a **"Pomiń →"** toggle — skipping a section renders that section's fields as blank lines (`________`) in the output.

### Section 1 — Dane ogólne
- Dysponent (name of dispatch center) — text input
- WKRM województwo — text input
- Autor raportu — text input
- Wysłano do WKRM — checkbox (TAK/NIE)

### Section 2 — Typy ZRM
One row per team in `teams[]`:
- Team name (read-only label)
- Type: dropdown → S / P / LZRM

### Section 3 — Szczegóły KAM
Pre-filled from `kamName`, `eventTime`, `eventEndTime`:
- Imię i nazwisko KAM — pre-filled from `kamName`
- Rola: dropdown → Ratownik Medyczny / Lekarz / Pielęgniarka
- Pełnił funkcję od: pre-filled from `eventTime`
- Pełnił funkcję do: pre-filled from `eventEndTime`

*(App tracks a single KAM. The form supports multiple — this generates one row.)*

### Section 4 — Sposoby transportu
Auto-computed values shown as read-only, user fills in breakdown:
- Łącznie poszkodowanych — auto (`patients.length`)
- Łącznie przetransportowanych — auto (`patients.filter(p => p.hospitalId).length`)
- Pozostali na miejscu — auto (`patients.filter(p => p.status === 'na_miejscu').length`)

User fills in (defaults to 0, all numeric inputs):
- Drogą powietrzną
- Drogą naziemną bez ZRM (inne środki transportu)
- Drogą wodną
- Udzielono pomocy na miejscu, brak transportu
- Nie wyrazili zgody na transport
- Odmówili udzielenia pomocy
- Samodzielnie oddalili się

*Drogą naziemną przez ZRM = max(0, total transported − (air + non-ZRM + water)), computed automatically. No strict validation — if inputs exceed total, naziemne ZRM renders as 0.*

### Section 5 — Dane szpitali
One collapsible card per hospital in `hospitals[]`. Pre-filled: name, redCapacity, yellowCapacity. The section-level "Pomiń →" toggle skips all hospital supplement fields (all render as `________`). Individual hospital cards can be expanded/collapsed for readability but cannot be individually skipped — skip is all-or-nothing for the section.

Per hospital:
- Nr telefonu — text input
- Stanowiska ITM (z respiratorem): 1h ___ / 2h ___
- Wolne respiratory (poza ITM): 1h ___ / 2h ___
- Stanowiska IN: 1h ___ / 2h ___
- Bloki operacyjne: 1h ___ / 2h ___

Oddziały (fixed list, numeric inputs):
- OIT: 1h ___ / 2h ___
- N-chir (neurochirurgia): 1h ___ / 2h ___
- Ortopedia: 1h ___ / 2h ___
- Chirurgia ogólna: 1h ___ / 2h ___
- Chirurgia naczyniowa: 1h ___ / 2h ___
- SOR: 1h ___ / 2h ___

---

## Generated HTML File

**Filename:** `dokumenty-mci-YYYY-MM-DD.html`
**Structure:** 3 documents separated by CSS `page-break-after: always`

```
[Zał. 12 — Tabela Dyslokacji Poszkodowanych]
          ← page break →
[Zał. 13 — Tabela Szpitali]
          ← page break →
[Zał. 14 — Raport z przebiegu zdarzenia]
```

Self-contained: all styles inline in `<style>` tag, no external resources.

---

## Document 1 — Tabela Dyslokacji Poszkodowanych (Zał. 12)

**Header:** "ZAŁĄCZNIK NR 12 — TABELA DYSLOKACJI POSZKODOWANYCH" + event name + date + KAM name

**Structure:** 4 tables, one per triage group, color-coded headers:

| Group | Header color | Columns |
|---|---|---|
| CZERWONI | Red | Nr, Obrażenia ciała, Dodatkowy opis, ZRM, Szpital |
| ŻÓŁCI | Yellow | Nr, Obrażenia ciała, Dodatkowy opis, ZRM, Szpital |
| ZIELONI | Green | Nr, Obrażenia ciała, Dodatkowy opis, ZRM, Szpital |
| CZARNI | Black/grey | Nr, Obrażenia ciała |

**Patient ordering:** sequential 1, 2, 3... within each group (not P-001 tags)

**Body injuries column:** app has 9 zones × 7 injury types. Rendered as compact text per active zone:
```
Głowa: złamanie, krwawienie
Kl. piersiowa: oparzenie
```
If no injuries recorded: cell is blank.

**ZRM column:** `teams.find(t => t.id === p.teamId)?.name || '________'`
**Szpital column:** `hospitals.find(h => h.id === p.hospitalId)?.name || '________'`
**CZARNI:** no ZRM or Szpital columns per procedura specification.

---

## Document 2 — Tabela Szpitali (Zał. 13)

**Header:** "ZAŁĄCZNIK NR 13 — TABELA SZPITALE" + event name + date

One card per hospital (from `hospitals[]`). Pre-filled fields shown with values; skipped/empty fields shown as `________`.

Card structure per hospital:
```
Nazwa: [h.name]           Nr telefonu: [supplement or ________]
Możliwość przyjęcia:      CZERWONEJ: [h.redCapacity]   ŻÓŁTEJ: [h.yellowCapacity]

DOSTĘPNOŚĆ (przy wstrzymaniu planowych zabiegów):
                    W ciągu 1 godz.   W ciągu 2 godz.   Uwagi
Stanowiska ITM      [___]             [___]
Wolne RESP.         [___]             [___]
Stanowiska IN       [___]             [___]
Bloki operacyjne    [___]             [___]

ODDZIAŁY:
OIT                 [___]             [___]
N-chir              [___]             [___]
Ortopedia           [___]             [___]
Chirurgia ogólna    [___]             [___]
Chirurgia naczynio. [___]             [___]
SOR                 [___]             [___]
```

---

## Document 3 — Raport z przebiegu zdarzenia (Zał. 14)

**Header:** "ZAŁĄCZNIK NR 14 — RAPORT Z PRZEBIEGU ZDARZENIA Z DUŻĄ LICZBĄ POSZKODOWANYCH"

### Sekcja 1. Dysponent
`supplement.dysponent || '________'`

### Sekcja 2. Data zdarzenia
- Data: derived from `eventTime`
- Godzina przyjęcia powiadomienia: `eventTime`

### Sekcja 3. Data zakończenia
- `eventEndTime`

### Sekcja 4. GDM
- Imię i nazwisko: `gdmName || '________'`
- KOD GDM: `'________'` (not tracked by app)

### Sekcja 5. ZRM
Counts by type, computed from supplement section 2:

| Typ | Liczba |
|-----|--------|
| "S" | count of teams marked S |
| "P" | count of teams marked P |
| "LZRM" | count of teams marked LZRM |

Kryptonimy: all team names joined, e.g. `ZRM W-01, ZRM W-02, LPR-01`

### Sekcja 6. KAM
One row from supplement section 3:

| Imię i nazwisko | Kryptonim ZRM | Od godz. | Do godz. | Lekarz | Ratownik | Pielęgniarka |
|---|---|---|---|---|---|---|
| kamName | — | eventTime | eventEndTime | TAK/NIE | TAK/NIE | TAK/NIE |

### Sekcja 7. Zestawienie poszkodowanych

Auto-computed totals + supplement transport breakdown:

| Kategoria | Liczba |
|---|---|
| Łącznie poszkodowanych | `patients.length` |
| Drogą powietrzną | supplement |
| Drogą naziemną przez ZRM | computed = total transported − (air + non-ZRM ground + water) |
| Drogą naziemną bez ZRM | supplement |
| Drogą wodną | supplement |
| Na miejscu bez transportu | supplement |
| Nie wyrazili zgody | supplement |
| Odmówili udzielenia pomocy | supplement |
| Samodzielnie oddalili się | supplement |

### Sekcja 8. Wysłanie do WKRM
`supplement.sentToWkrm ? 'TAK' : 'NIE'` — do WKRM województwa: `supplement.wkrmProvince || '________'`

### Sekcja 9. Autor raportu
`supplement.reportAuthor || '________'`
Godzina: time of generation (at download click)

### Załącznik do Raportu — Tabela per-szpital
Patients grouped by hospital, columns: Nr ZRM, Czerwoni count, Żółci count, Zieloni count.

---

## JS Architecture

New functions added to `index.html`:

```js
// State for modal inputs
let mzSupplement = {};

function openMzModal()           // initialize mzSupplement, render modal, show overlay
function renderMzModal()         // build modal HTML from teams/hospitals
function saveMzSupplements()     // read all modal inputs into mzSupplement
function generateMzHtml()        // builds full HTML string for all 3 docs
function downloadMzDocs()        // saveMzSupplements() → generateMzHtml() → blob download

// Sub-generators called by generateMzHtml():
function genZal12()              // Tabela Dyslokacji
function genZal13()              // Tabela Szpitali
function genZal14()              // Raport
function injuryText(injuries)    // formats injury zones to readable text
```

New HTML element: `<div id="mzModal" class="modal-overlay">...</div>`

Trigger in `renderReport()`:
```js
if (eventEndTime) {
  // show "Generuj dokumenty MZ" button
}
```

---

## Constraints

- No external libraries — pure vanilla JS + inline CSS
- Works fully offline (no fetch, no CDN)
- Generated file is self-contained HTML (~30–80 KB depending on patient count)
- Polish language only for generated documents (procedura is PL-only)
- App currently tracks one KAM — Zał. 14 Sekcja 6 generates one KAM row
- App body diagram has 9 zones; procedura Zał. 12 has 13 — rendered as text, visually acceptable

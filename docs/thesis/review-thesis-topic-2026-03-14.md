# Thesis Topic Review — 2026-03-14

## Current title

*Porównanie efektywności cyfrowego systemu wspomagania segregacji START z metodą tradycyjną w warunkach symulowanego zdarzenia masowego*

## What's strong

- **PICOT/FINER** are well-formulated — the research question is focused and answerable
- **Crossover design** is the right call for n=40-50 — each participant as their own control gives you statistical power you wouldn't have otherwise
- **Two complementary pillars** (experiment + survey) give the thesis real depth
- **Chapters 1-3 are well-written** — the literature review is thorough, the methodology section is detailed and rigorous
- The **practical motivation** (you're a working paramedic who built the tool) is a genuine strength, not a weakness

## Issues worth addressing

### 1. Conflict of interest / bias
You developed the app **and** you're testing it. This is the #1 thing a reviewer will flag. You address it nowhere explicitly. Add a subsection in 3.8 (Aspekty etyczne) or 3.9 (Ograniczenia) that names this directly and explains your mitigations: blinded scoring of reports, validated klucz referencyjny, standardized SUS instrument, crossover eliminating between-subject bias.

### 2. Card-based simulation is a double-edged sword
The "karty pacjentów" approach tests **algorithmic knowledge**, not **field performance**. That's fine — but the title says "w warunkach symulowanego zdarzenia masowego," which overpromises. Either:
- Upgrade the simulation (mannequins, actors, time pressure, noise) — harder but more credible
- Or adjust the title to reflect what you're actually testing, e.g.:

> *Wpływ cyfrowego wspomagania na trafność i czas segregacji START — badanie eksperymentalne z użyciem scenariuszy kartkowych i ankieta wśród polskich ratowników medycznych*

### 3. H1 (app is faster) may backfire
Typing on a phone screen can be **slower** than scribbling on a card, especially for users unfamiliar with the app (5 min training only). If the app turns out slower, that's not a failure — it's a finding. But your hypotheses are all directional (app > traditional). Consider:
- Making H1 non-directional ("differs" instead of "shortens") — more defensible
- Or keeping it directional but preparing a strong discussion for the case where H1 is rejected (the report completeness advantage may matter more than raw speed)

### 4. "Kompletność raportu" (H3) is almost trivially true
The app auto-generates a 10/10 report by design. The traditional method requires hand-writing all of it. This isn't really testing the user — it's testing whether software can generate text. Consider:
- Acknowledging this in the hypotheses section (it's expected, not surprising)
- Framing H3 as measuring the **practical advantage** rather than testing a genuine uncertainty
- Or adding a qualitative dimension: was the traditional report **usable** by a dispatch coordinator, even if incomplete?

### 5. Sample composition matters
"Studenci ratownictwa + czynni ratownicy" is a heterogeneous group. Students may perform very differently from 10-year veterans. Your crossover controls for within-subject variance, but you should plan a **subgroup analysis** (students vs. professionals) — this could be one of your most interesting findings.

### 6. The survey is solid but generic
22 questions, snowball sampling, Google Forms — this is standard. Consider adding one concrete element that ties it to the app: include a 2-minute video or screenshots of TRIAGE MCI in the survey, then ask "would you use this specific tool?" instead of just "would you use *a* digital tool?" This makes the survey directly relevant to the thesis instead of being a standalone attitude study.

### 7. Missing: effect size estimation / power analysis for crossover
You have a power analysis for the survey (n≈88 for chi-square), but **not for the experiment** — which is the core of the thesis. For a paired t-test with d=0.5, α=0.05, power=0.80, you need n≈34. With 40-50 participants you're fine, but you should state this explicitly in 3.7.

## Suggested revised titles

If you want to keep it ambitious:
> *Efektywność cyfrowego wspomagania segregacji START — eksperyment symulacyjny i postawy polskich ratowników medycznych wobec narzędzi cyfrowych w zdarzeniach masowych*

If you want it tighter and more accurate:
> *Porównanie trafności, czasu i kompletności raportowania segregacji START z użyciem aplikacji mobilnej i metody tradycyjnej — badanie eksperymentalne crossover*

## Priority actions

1. **Decide on simulation fidelity** — cards vs. something more realistic. This shapes everything.
2. **Add the bias/conflict-of-interest section** — small effort, big credibility gain.
3. **Run the power analysis for the crossover** and document it in chapter 3.
4. **Start building the actual experiment materials** — the karty pacjentów, klucz referencyjny, SUS w polskiej wersji, formularz zgody. These take longer than you think.

---

## Update: Analysis after reading Procedura MZ v2.3 (2026-03-14)

### Key discovery

The official MCI procedure (Ministry of Health, v2.3, June 15 2024, approved by Marek Kos — Podsekretarz Stanu) defines three **paper-based forms** that are mandatory during every MCI with ≥10 persons:

| Official form | Załącznik | TRIAGE MCI equivalent |
|---|---|---|
| TABELA DYSLOKACJI POSZKODOWANYCH | zał. 12 | Patient list: color groups, numbering (P-001...), injury markings, ZRM assignment, hospital destination |
| TABELA SZPITALE | zał. 13 | Hospital list with capacity tracking |
| Raport GDM z przebiegu zdarzenia | zał. 14 | Auto-generated event report (times, patient counts, ZRM list, hospital summary) |

### What this changes

**The thesis framing shifts fundamentally.** TRIAGE MCI is not just "an app that helps with START" — it is a **digital implementation of the officially mandated MCI documentation workflow**. This makes the research question much stronger: instead of testing a generic "digital vs paper triage," the thesis tests whether a digital tool can improve compliance with the actual regulatory framework.

### New actions based on this discovery

1. **Update Chapter 2 (Literature review)** — cite the procedure as the governing regulation, describe TABELA DYSLOKACJI / TABELA SZPITALE / Raport GDM, note that these are currently paper-based ✅ DONE
2. **Reframe the research question** — ground it in the official procedure rather than generic "digital vs traditional" ✅ DONE
3. **Redesign H3 (report completeness)** — replace the arbitrary 10-item checklist with the actual Raport GDM requirements from zał. 14 ✅ DONE
4. **Redesign experiment scenarios** — test the full KAM workflow (segregation + ZRM assignment + hospital assignment + report), not just triage decisions ✅ DONE
5. **Use the procedure's evaluation framework** — adapt the Karty ocen (3-level scale: prawidłowa / nieoptymalna / wymaga poprawy) for experiment evaluation ✅ DONE (3.4.8 + observer form in scenariusze)
6. **Update the survey** — add questions about familiarity with Procedura MZ v2.3, experience with TABELA DYSLOKACJI, attitudes toward digitizing official MCI forms ✅ DONE (Sekcja E, 5 questions)
7. **Consider revised title** ✅ DONE:
   > *Cyfrowe wspomaganie realizacji Procedury postępowania na wypadek zdarzenia z dużą liczbą poszkodowanych — porównanie aplikacji TRIAGE MCI z metodą papierową w badaniu eksperymentalnym*

### Revised priority order

1. ~~Convert the procedure docx to markdown for easy reference~~ ✅ IN PROGRESS (background agent)
2. Map each field in TABELA DYSLOKACJI / TABELA SZPITALE / Raport GDM to app features ✅ DONE (table in 3.6.2)
3. ~~Update Chapter 2 with procedure citation and analysis~~ ✅ DONE (new section 2.3.6)
4. ~~Redesign experiment to test full KAM/ZRM workflow~~ ✅ DONE (updated Faza 2 in 3.4.4)
5. ~~Rebuild report completeness checklist from official Raport GDM template~~ ✅ DONE (3.4.6)
6. ~~Update survey with procedure-specific questions~~ ✅ DONE (Sekcja E)
7. ~~Add bias/conflict-of-interest section to Chapter 3~~ ✅ DONE (3.8.4 — 4 bias types + mitigations)
8. ~~Add power analysis for crossover to Chapter 3~~ ✅ DONE (3.7.8 — paired t-test, n=34 min, n=40-50 justified)
9. ~~Update experiment materials with full KAM workflow~~ ✅ DONE (updated instructions, Raport GDM checklist, ZRM/hospital lists, paper Tabela Dyslokacji template, Karta ocen observer form)

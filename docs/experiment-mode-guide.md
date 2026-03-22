# Tryb eksperymentalny TRIAGE MCI — przewodnik

## Co to jest

Tryb eksperymentalny to ukryta funkcja wbudowana w osobne kopie aplikacji (`experiment.html` i `dyspozytor/experiment.html`). Rejestruje każdą akcję uczestnika z dokładnością do milisekund — bez widocznej różnicy w działaniu aplikacji. Uczestnik nie wie, że jego działania są logowane (poza małym migającym wskaźnikiem REC).

## Pliki

| Plik | Opis |
|------|------|
| `experiment.html` | Kopia głównej aplikacji z wbudowanym loggerem akcji |
| `dyspozytor/experiment.html` | Kopia strony dyspozytora — generuje linki SMS kierujące do `experiment.html` |

Pliki produkcyjne (`index.html`, `dyspozytor/index.html`) nie zawierają żadnego kodu eksperymentalnego — pozostają czyste.

---

## Jak się dostać do trybu eksperymentalnego

### Z głównej aplikacji (`index.html`)
1. Otwórz panel informacyjny (przycisk **Info**)
2. W prawym górnym rogu info-boxa znajduje się ledwo widoczna ikona 🔬 (15% przezroczystości)
3. **Tapnij ją 3 razy** w ciągu 800ms → przekierowanie do `experiment.html`

### Ze strony dyspozytora (`dyspozytor/index.html`)
1. Obok tytułu „Dyspozytor" w prawym górnym rogu jest ikona 🔬
2. **Tapnij 3 razy** → przekierowanie do `dyspozytor/experiment.html`

---

## Przebieg eksperymentu krok po kroku

### Przygotowanie (przed sesją)

1. **Otwórz `dyspozytor/experiment.html`** (przez triple-tap na stronie dyspozytora lub bezpośrednio wpisując URL)
2. **Wypełnij dane scenariusza:**
   - Nazwa zdarzenia (np. „Wypadek autobus A4 km 312")
   - GDM (Główny Dyspozytor Medyczny)
   - Zespoły ZRM (dodaj z listy scenariusza)
   - Szpitale z pojemnością 🔴/🟡 (dodaj z listy scenariusza)
3. **Wygeneruj podgląd** i **wyślij SMS** na urządzenie uczestnika
   - Link w SMS wskazuje na `experiment.html?i=BASE64DATA`
   - Uczestnik tapnie link → otworzy się wersja eksperymentalna z zaimportowanymi danymi

### Uruchomienie nagrywania

4. Na urządzeniu uczestnika: otwórz panel Info → **triple-tap 🔬** (lub jeśli już jesteś na `experiment.html`, otwórz Info i tapnij 🔬)
5. Pojawi się panel **Experiment Mode**:
   - **Participant ID** — wpisz identyfikator uczestnika (np. P01, S12)
   - **Scenario** — wpisz oznaczenie scenariusza (np. A1, B2)
6. Kliknij **START REC**
7. W prawym górnym rogu ekranu pojawi się **migający czerwony wskaźnik REC**

### Sesja uczestnika

8. Uczestnik pracuje normalnie: segregacja → dysponowanie → raport
9. Wszystkie akcje są logowane w tle — uczestnik nie musi nic dodatkowego robić
10. **NIE zamykaj aplikacji** w trakcie sesji — log jest w pamięci

### Eksport danych

11. Po zakończeniu sesji: **tapnij wskaźnik REC** (prawy górny róg)
12. Pojawi się panel eksportu z informacjami:
    - ID uczestnika
    - Scenariusz
    - Liczba zalogowanych zdarzeń
    - Liczba pacjentów
13. Kliknij **📥 Export JSON** → plik zostanie pobrany jako:
    `exp_[participantID]_[scenario]_[data-czas].json`
14. Opcjonalnie: **⏹ Stop Recording** zatrzymuje nagrywanie

### Między sesjami

15. Przed kolejnym uczestnikiem:
    - Wyczyść dane przeglądarki (localStorage) lub kliknij „Nowe zdarzenie"
    - Powtórz od kroku 4

---

## Co jest logowane

| Akcja | Pole `action` | Szczegóły (`detail`) | Znaczenie badawcze |
|-------|---------------|----------------------|-------------------|
| Start eksperymentu | `experiment_start` | participantId, scenario | Znacznik początku sesji |
| Start zdarzenia | `event_start` | eventName, kamName | Początek pomiaru czasu |
| Zmiana zakładki | `tab_switch` | from, to | Podział na fazy: segregacja / dysponowanie / raport |
| Nowy pacjent (kreator) | `wizard_start` | patientNumber | Początek triażu pacjenta |
| Odpowiedź TAK/NIE | `wizard_answer` | step, answer (YES/NO) | Ścieżka decyzyjna START — porównanie z kluczem referencyjnym |
| Wynik triażu | `triage_result` | category (T1/T2/T3/T4), reason | Przypisana kategoria |
| Potwierdzenie pacjenta | `patient_confirmed` | tag, category, sex, age, path, hasInjuries | Koniec segregacji pacjenta — pomiar czasu |
| Wysyłka transportu | `dispatch_transport` | patient, category, team, hospital | Dysponowanie — dokładność i czas |
| Pacjent dostarczony | `patient_delivered` | patient, hospital | Zakończenie transportu |
| Retriaż | `retriage` | patient, from, to | Zmiana kategorii — wzorce decyzyjne |
| Zmiana szpitala | `hospital_change` | patient, from, to | Korekty dysponowania |
| Zamknięcie zdarzenia | `event_end` | endTime | Koniec sesji — czas całkowity |
| Cofnięcie zamknięcia | `event_end_reverted` | — | Korekta |
| Stop eksperymentu | `experiment_stop` | — | Znacznik końca |

---

## Struktura wyeksportowanego pliku JSON

```json
{
  "meta": {
    "participantId": "P01",
    "scenario": "A1",
    "exportTime": "2026-04-15T14:32:01.203Z",
    "eventName": "Wypadek autobus A4 km 312",
    "eventTime": "14:25:03",
    "eventEndTime": "14:48:22",
    "totalEvents": 87
  },
  "log": [
    {
      "t": 1711373521203,
      "ts": "14:32:01",
      "action": "wizard_start",
      "detail": { "patientNumber": 1 }
    },
    {
      "t": 1711373524891,
      "ts": "14:32:04",
      "action": "wizard_answer",
      "detail": { "step": "walking", "answer": "NO" }
    }
  ],
  "state": {
    "patients": [ ... ],
    "transports": [ ... ],
    "teams": [ ... ],
    "hospitals": [ ... ]
  }
}
```

### Pola kluczowe

- **`t`** — czas w milisekundach (Unix timestamp) — do precyzyjnych obliczeń czasu
- **`ts`** — czas w formacie czytelnym (HH:MM:SS) — do szybkiego podglądu
- **`action`** — typ akcji (patrz tabela wyżej)
- **`detail`** — szczegóły zależne od typu akcji
- **`state`** — pełen snapshot stanu aplikacji w momencie eksportu

---

## Wykorzystanie w pracy magisterskiej

### H1 — Czas segregacji

Obliczanie czasu z logu:
- **Czas segregacji pacjenta** = `patient_confirmed.t` − `wizard_start.t` (dla tego samego patientNumber)
- **Czas fazy segregacji** = suma czasów wszystkich pacjentów, lub: czas od pierwszego `wizard_start` do ostatniego `patient_confirmed`
- **Czas fazy dysponowania** = czas od pierwszego `tab_switch` na dispatch do ostatniego `dispatch_transport`
- **Czas raportu** = czas od `tab_switch` na report do `event_end`
- **Czas całkowity** = `event_end.t` − `event_start.t`

Precyzja milisekundowa eliminuje błąd ludzkiego pomiaru stoperem.

### H2 — Dokładność segregacji

Z logu odczytasz:
- Pełną ścieżkę decyzyjną dla każdego pacjenta (sekwencja `wizard_answer`)
- Przypisaną kategorię (`patient_confirmed.detail.category`)
- Porównanie z kluczem referencyjnym z `scenariusze-eksperyment.md`

Możesz ustalić nie tylko CZY uczestnik przypisał złą kategorię, ale GDZIE w algorytmie START popełnił błąd.

### H3 — Kompletność raportu

`state` zawiera pełen stan końcowy:
- Wszystkich pacjentów z danymi (tag, kategoria, płeć, wiek, obrażenia, notatki)
- Wszystkie transporty (pacjent → zespół → szpital → czas)
- Historię retriażu i zmian szpitali

Porównaj automatycznie z 10-punktową checklistą kompletności raportu GDM (zał. 14).

### Dane behawioralne (bonus do dyskusji)

- **Czas wahania** — przerwy między `wizard_answer` (długi czas = niepewność)
- **Kolejność nawigacji** — sekwencja `tab_switch` (czy uczestnik wraca do segregacji po dysponowaniu?)
- **Częstość retriażu** — ile razy zmieniana jest kategoria
- **Czas per pacjent** — krzywa uczenia się (czy 12. pacjent jest szybszy niż 1.?)

---

## Przetwarzanie danych (po zebraniu)

Po zebraniu 40-50 plików JSON z sesji eksperymentalnych, przetworzenie w Pythonie:

```python
import json, glob, statistics

files = glob.glob('exp_*.json')
results = []

for f in files:
    data = json.load(open(f))
    log = data['log']

    # Czas per pacjent
    starts = [e for e in log if e['action'] == 'wizard_start']
    confirms = [e for e in log if e['action'] == 'patient_confirmed']

    for s, c in zip(starts, confirms):
        duration_ms = c['t'] - s['t']
        results.append({
            'participant': data['meta']['participantId'],
            'scenario': data['meta']['scenario'],
            'patient': c['detail']['tag'],
            'category': c['detail']['category'],
            'duration_s': duration_ms / 1000
        })

# Średni czas segregacji per uczestnik
# Porównanie z kluczem referencyjnym
# Export do CSV → analiza w SPSS/R
```

To jest przykładowy zarys — dostosuj do swoich potrzeb analitycznych.

---

## Uwagi praktyczne

- **Log jest w pamięci** — zamknięcie/odświeżenie strony kasuje log. Zawsze eksportuj PRZED zamknięciem sesji.
- **Tryb offline działa** — experiment.html jest w cache SW (v48+). Eksport JSON pobiera plik lokalnie.
- **Uczestnik nie widzi logowania** — jedyny widoczny element to mały migający REC w rogu (można go zignorować).
- **Dane z importu SMS** — jeśli uczestnik otworzył experiment.html przez link SMS, dane scenariusza (zespoły, szpitale) są już załadowane. Nie trzeba ich wprowadzać ręcznie.
- **Identyfikator uczestnika** — wpisz przed sesją zgodnie ze swoim arkuszem randomizacji (np. P01 = osoba 1, scenariusz A w ramieniu 1).

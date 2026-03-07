# Platforma badawcza TRIAGE MCI — projekt techniczny

**Data:** 2026-03-07
**Status:** Do realizacji w osobnym repozytorium

## 1. Cel

Webowa platforma do zbierania danych badawczych do pracy magisterskiej:
- **Ankieta ogólnopolska** — publiczny formularz (zastępuje Google Forms)
- **Panel eksperymentu** — prowadzący wpisuje wyniki uczestników
- **Kwestionariusz SUS** — uczestnik wypełnia po eksperymencie
- **Dashboard** — podgląd wyników + eksport CSV

## 2. Architektura

```
┌─────────────────────────────────────────────────┐
│  Frontend (HTML/CSS/JS — vanilla)                │
│                                                   │
│  /ankieta        — publiczny formularz 22 pytań  │
│  /sus            — formularz SUS (link z kodem)  │
│  /panel          — panel prowadzącego (hasło)    │
│  /dashboard      — wyniki + eksport (hasło)      │
└──────────────────┬──────────────────────────────┘
                   │ fetch('/api/...')
┌──────────────────▼──────────────────────────────┐
│  API (Node.js)                                    │
│                                                   │
│  POST /api/survey          — zapis ankiety       │
│  POST /api/sus             — zapis SUS           │
│  POST /api/participants    — rejestracja         │
│  POST /api/sessions        — zapis sesji         │
│  GET  /api/results         — pobranie wyników    │
│  GET  /api/export/:type    — eksport CSV         │
│  POST /api/auth            — logowanie panelu    │
└──────────────────┬──────────────────────────────┘
                   │ @neondatabase/serverless
┌──────────────────▼──────────────────────────────┐
│  Neon Postgres                                    │
└─────────────────────────────────────────────────┘
```

### Tech stack

| Warstwa | Technologia | Uzasadnienie |
|---------|-------------|--------------|
| Frontend | HTML + CSS + Vanilla JS | Spójne z umiejętnościami (TRIAGE MCI) |
| Backend | Node.js + Express (lub Hono) | Prosty REST API, minimalna konfiguracja |
| Baza danych | Neon Postgres (serverless) | Darmowy tier, Postgres, zero utrzymania |
| Driver | `@neondatabase/serverless` | Oficjalny driver Neon dla Node.js |
| Hosting | Vercel (zalecane) lub Railway | Darmowy tier, obsługa static + API |
| Auth | Proste hasło w env (panel) | Jedno hasło dla prowadzącego, brak kont użytkowników |

## 3. Schemat bazy danych

```sql
-- ============================================
-- ANKIETA OGÓLNOPOLSKA
-- ============================================

CREATE TABLE survey_responses (
  id SERIAL PRIMARY KEY,
  created_at TIMESTAMPTZ DEFAULT NOW(),

  -- Sekcja A: Dane demograficzne
  age_range VARCHAR(20) NOT NULL,          -- '20-25','26-30','31-35','36-40','41-45','45+'
  gender VARCHAR(20) NOT NULL,             -- 'kobieta','mezczyzna','nie_podaje'
  education VARCHAR(40) NOT NULL,          -- 'licencjat','magister','inne'
  education_other TEXT,
  professional_status VARCHAR(40) NOT NULL,-- 'student','ratownik','oba','inny'
  professional_status_other TEXT,
  work_experience VARCHAR(20) NOT NULL,    -- 'brak','<1','1-3','4-7','8-15','15+'
  workplace VARCHAR(30) NOT NULL,          -- 'P','S','SOR','LPR','inne','nie_dotyczy'
  workplace_other TEXT,
  city_size VARCHAR(20) NOT NULL,          -- '<10k','10-50k','50-100k','100-500k','500k+'

  -- Sekcja B: Doświadczenie z MCI
  mci_training BOOLEAN NOT NULL,
  triage_systems TEXT[],                   -- ['START','JumpSTART','SALT','Sieve','inny','nie_pamietam']
  triage_systems_other TEXT,
  real_mci VARCHAR(20) NOT NULL,           -- 'nie','1','2-3','4+'
  simulation_count VARCHAR(20) NOT NULL,   -- 'nie','1','2-3','4+'
  preparedness_rating INT NOT NULL CHECK (preparedness_rating BETWEEN 1 AND 5),

  -- Sekcja C: Stosunek do narzędzi cyfrowych
  uses_medical_apps VARCHAR(20) NOT NULL,  -- 'regularnie','sporadycznie','nie'
  medical_apps_names TEXT,
  heard_of_triage_apps BOOLEAN NOT NULL,
  app_usefulness_rating INT NOT NULL CHECK (app_usefulness_rating BETWEEN 1 AND 5),
  would_use_on_real_mci VARCHAR(20) NOT NULL, -- 'zdecydowanie_tak','raczej_tak','trudno','raczej_nie','zdecydowanie_nie'
  app_requirements TEXT,

  -- Sekcja D: Bariery
  barriers TEXT[] NOT NULL,                -- wielokrotny wybór
  barriers_other TEXT,
  worst_barrier VARCHAR(60) NOT NULL,
  worst_barrier_other TEXT,
  adoption_factors TEXT[] NOT NULL,        -- wielokrotny wybór
  adoption_factors_other TEXT,
  digital_future VARCHAR(20) NOT NULL      -- 'zdecydowanie_tak','raczej_tak','trudno','raczej_nie','zdecydowanie_nie'
);

-- ============================================
-- EKSPERYMENT CROSSOVER
-- ============================================

CREATE TABLE participants (
  id SERIAL PRIMARY KEY,
  code VARCHAR(10) UNIQUE NOT NULL,        -- 'U01', 'U02', ...
  created_at TIMESTAMPTZ DEFAULT NOW(),
  sequence VARCHAR(2) NOT NULL CHECK (sequence IN ('AB', 'BA')),
  status VARCHAR(20) NOT NULL,             -- 'student', 'ratownik'
  age_range VARCHAR(20),
  gender VARCHAR(20),
  work_experience VARCHAR(20),
  has_real_mci_experience BOOLEAN
);

CREATE TABLE experiment_sessions (
  id SERIAL PRIMARY KEY,
  participant_id INT NOT NULL REFERENCES participants(id),
  session_number INT NOT NULL CHECK (session_number IN (1, 2)),
  method VARCHAR(20) NOT NULL CHECK (method IN ('tradycyjna', 'aplikacja')),
  scenario VARCHAR(2) NOT NULL CHECK (scenario IN ('S1', 'S2')),
  created_at TIMESTAMPTZ DEFAULT NOW(),

  -- Czas (timer wbudowany w panel)
  timer_started_at TIMESTAMPTZ,            -- moment kliknięcia START
  timer_stopped_at TIMESTAMPTZ,            -- moment kliknięcia STOP
  time_seconds INT NOT NULL,               -- obliczony automatycznie (stop - start)

  -- Kompletność raportu (10 elementów, każdy 0 lub 1)
  report_event_name BOOLEAN DEFAULT FALSE,
  report_start_time BOOLEAN DEFAULT FALSE,
  report_total_count BOOLEAN DEFAULT FALSE,
  report_category_breakdown BOOLEAN DEFAULT FALSE,
  report_patient_list BOOLEAN DEFAULT FALSE,
  report_transport_info BOOLEAN DEFAULT FALSE,
  report_hospitals BOOLEAN DEFAULT FALSE,
  report_teams BOOLEAN DEFAULT FALSE,
  report_transport_status BOOLEAN DEFAULT FALSE,
  report_report_time BOOLEAN DEFAULT FALSE,

  UNIQUE(participant_id, session_number)
);

CREATE TABLE classifications (
  id SERIAL PRIMARY KEY,
  session_id INT NOT NULL REFERENCES experiment_sessions(id),
  patient_code VARCHAR(10) NOT NULL,       -- 'S1-01', 'S1-02', ..., 'S2-01', ...
  assigned_category VARCHAR(2) NOT NULL CHECK (assigned_category IN ('T1','T2','T3','T4')),

  UNIQUE(session_id, patient_code)
);

-- Klucz referencyjny (wypełniony raz, niezmienny)
CREATE TABLE reference_key (
  patient_code VARCHAR(10) PRIMARY KEY,    -- 'S1-01', 'S2-12', ...
  correct_category VARCHAR(2) NOT NULL CHECK (correct_category IN ('T1','T2','T3','T4')),
  start_path TEXT                          -- opis ścieżki START
);

-- ============================================
-- KWESTIONARIUSZ SUS
-- ============================================

CREATE TABLE sus_responses (
  id SERIAL PRIMARY KEY,
  participant_id INT NOT NULL REFERENCES participants(id) UNIQUE,
  created_at TIMESTAMPTZ DEFAULT NOW(),

  -- 10 stwierdzeń SUS (skala 1-5)
  q1 INT NOT NULL CHECK (q1 BETWEEN 1 AND 5),
  q2 INT NOT NULL CHECK (q2 BETWEEN 1 AND 5),
  q3 INT NOT NULL CHECK (q3 BETWEEN 1 AND 5),
  q4 INT NOT NULL CHECK (q4 BETWEEN 1 AND 5),
  q5 INT NOT NULL CHECK (q5 BETWEEN 1 AND 5),
  q6 INT NOT NULL CHECK (q6 BETWEEN 1 AND 5),
  q7 INT NOT NULL CHECK (q7 BETWEEN 1 AND 5),
  q8 INT NOT NULL CHECK (q8 BETWEEN 1 AND 5),
  q9 INT NOT NULL CHECK (q9 BETWEEN 1 AND 5),
  q10 INT NOT NULL CHECK (q10 BETWEEN 1 AND 5),

  -- Wynik obliczony automatycznie
  sus_score NUMERIC(5,2),

  -- Pytania dodatkowe
  preferred_method VARCHAR(20),            -- 'tradycyjna','aplikacja','trudno'
  preference_reason TEXT
);

-- ============================================
-- DANE REFERENCYJNE (INSERT raz)
-- ============================================

INSERT INTO reference_key (patient_code, correct_category, start_path) VALUES
  -- Scenariusz 1
  ('S1-01', 'T3', 'Chodzi → T3'),
  ('S1-02', 'T3', 'Chodzi → T3'),
  ('S1-03', 'T3', 'Chodzi → T3'),
  ('S1-04', 'T3', 'Chodzi → T3'),
  ('S1-05', 'T2', 'Nie chodzi → Oddycha (16/min) → CRT <2s → Wykonuje polecenia → T2'),
  ('S1-06', 'T2', 'Nie chodzi → Oddycha (22/min) → CRT <2s → Wykonuje polecenia → T2'),
  ('S1-07', 'T2', 'Nie chodzi → Oddycha (24/min) → CRT <2s → Wykonuje polecenia → T2'),
  ('S1-08', 'T1', 'Nie chodzi → Oddycha (36/min >30) → T1'),
  ('S1-09', 'T1', 'Nie chodzi → Oddycha (26/min) → CRT >4s → T1'),
  ('S1-10', 'T1', 'Nie chodzi → Oddycha (18/min) → CRT <2s → Nie wykonuje poleceń → T1'),
  ('S1-11', 'T4', 'Nie chodzi → Nie oddycha → Udrożnienie → Brak oddechu → T4'),
  ('S1-12', 'T4', 'Nie chodzi → Nie oddycha → Udrożnienie → Brak oddechu → T4'),
  -- Scenariusz 2
  ('S2-01', 'T3', 'Chodzi → T3'),
  ('S2-02', 'T3', 'Chodzi → T3'),
  ('S2-03', 'T3', 'Chodzi → T3'),
  ('S2-04', 'T3', 'Chodzi → T3'),
  ('S2-05', 'T2', 'Nie chodzi → Oddycha (20/min) → CRT <2s → Wykonuje polecenia → T2'),
  ('S2-06', 'T2', 'Nie chodzi → Oddycha (22/min) → CRT <2s → Wykonuje polecenia → T2'),
  ('S2-07', 'T2', 'Nie chodzi → Oddycha (18/min) → CRT <2s → Wykonuje polecenia → T2'),
  ('S2-08', 'T1', 'Nie chodzi → Oddycha (34/min >30) → T1'),
  ('S2-09', 'T1', 'Nie chodzi → Nie oddycha → Udrożnienie → Oddycha (14/min) → T1'),
  ('S2-10', 'T1', 'Nie chodzi → Oddycha (28/min) → CRT >4s → T1'),
  ('S2-11', 'T4', 'Nie chodzi → Nie oddycha → Udrożnienie → Brak oddechu → T4'),
  ('S2-12', 'T4', 'Nie chodzi → Nie oddycha → Udrożnienie → Brak oddechu → T4');
```

## 4. Endpointy API

### 4.1. Ankieta

```
POST /api/survey
Body: { age_range, gender, education, ..., digital_future }
Response: { success: true, id: 123 }

GET /api/survey/count
Response: { count: 147 }
```

Walidacja: wszystkie pola obowiązkowe obecne, rating w zakresie 1-5, tablice niepuste.

### 4.2. Uczestnicy eksperymentu

```
POST /api/participants
Body: { code: "U01", sequence: "AB", status: "student", age_range, gender, work_experience, has_real_mci_experience }
Response: { success: true, id: 1 }

GET /api/participants
Response: [{ id, code, sequence, status, ... }]
```

### 4.3. Sesje eksperymentalne

```
POST /api/sessions
Body: {
  participant_code: "U01",
  session_number: 1,
  method: "tradycyjna",
  scenario: "S1",
  timer_started_at: "2026-04-15T14:32:10.000Z",
  timer_stopped_at: "2026-04-15T14:39:13.000Z",
  time_seconds: 423,    // auto-obliczony lub wpisany ręcznie
  classifications: [
    { patient_code: "S1-01", assigned_category: "T3" },
    { patient_code: "S1-02", assigned_category: "T3" },
    ...  // 12 sztuk
  ],
  report_scores: {
    event_name: true,
    start_time: false,
    total_count: true,
    category_breakdown: true,
    patient_list: true,
    transport_info: false,
    hospitals: false,
    teams: false,
    transport_status: false,
    report_time: false
  }
}
Response: {
  success: true,
  accuracy: 91.67,           // % prawidłowych (11/12)
  correct_count: 11,
  total: 12,
  errors: [
    { patient_code: "S1-10", assigned: "T2", correct: "T1", type: "undertriage" }
  ],
  report_completeness: 4     // suma punktów /10
}
```

Serwer automatycznie:
- Porównuje `assigned_category` z `reference_key` → oblicza trafność
- Klasyfikuje błędy jako overtriage / undertriage
- Sumuje punkty raportu

### 4.4. Kwestionariusz SUS

```
POST /api/sus
Body: {
  participant_code: "U01",
  q1: 4, q2: 2, q3: 5, q4: 1, q5: 4, q6: 2, q7: 5, q8: 1, q9: 4, q10: 1,
  preferred_method: "aplikacja",
  preference_reason: "Szybciej i pewniej..."
}
Response: { success: true, sus_score: 87.5 }
```

Serwer automatycznie oblicza wynik SUS:
```js
function calculateSUS(q1,q2,q3,q4,q5,q6,q7,q8,q9,q10) {
  const odd  = (q1-1) + (q3-1) + (q5-1) + (q7-1) + (q9-1);
  const even = (5-q2) + (5-q4) + (5-q6) + (5-q8) + (5-q10);
  return (odd + even) * 2.5;
}
```

### 4.5. Wyniki i eksport

```
GET /api/results
Headers: { Authorization: "Bearer <PANEL_PASSWORD>" }
Response: {
  survey: { count, ... },
  experiment: {
    participants_count,
    sessions_completed,
    summary: {
      traditional: { avg_time, avg_accuracy, avg_report },
      app: { avg_time, avg_accuracy, avg_report }
    }
  },
  sus: { avg_score, count }
}

GET /api/export/survey       → CSV ze wszystkimi odpowiedziami ankiety
GET /api/export/experiment   → CSV z wynikami eksperymentu
GET /api/export/sus          → CSV z wynikami SUS
GET /api/export/all          → ZIP ze wszystkimi CSV
```

### 4.6. Autentykacja panelu

```
POST /api/auth
Body: { password: "..." }
Response: { token: "..." }   // prosty JWT lub losowy token w pamięci
```

Hasło ustawione w zmiennej środowiskowej `PANEL_PASSWORD`. Jeden użytkownik, jedno hasło.

## 5. Strony frontend

### 5.1. `/ankieta` — Formularz ankiety ogólnopolskiej

- **Dostęp:** Publiczny (bez logowania)
- **Treść:** Dokładne odwzorowanie 22 pytań z `ankieta-ogolnopolska.md`
- **Strona wstępna:** Informacja o badaniu + zgoda (przycisk "Rozpocznij")
- **Walidacja:** Na frontendzie (pola obowiązkowe, zakresy) + na backendzie
- **Logika warunkowa:** Pytanie 9 widoczne tylko gdy pyt. 8 = "Tak"
- **Strona końcowa:** Podziękowanie + link do TRIAGE MCI
- **UX:** Podział na sekcje (4 strony), pasek postępu, responsywne
- **Zabezpieczenie przed duplikatami:** Fingerprint urządzenia (localStorage flag) — nie blokuje, ale ostrzega

### 5.2. `/sus` — Kwestionariusz SUS

- **Dostęp:** Link z kodem uczestnika, np. `/sus?code=U01`
- **Treść:** 10 stwierdzeń SUS (radio 1-5) + 2 pytania dodatkowe
- **Kod uczestnika:** Wstępnie wypełniony z URL, edytowalny
- **Po wysłaniu:** Wyświetla obliczony wynik SUS z interpretacją

### 5.3. `/panel` — Panel prowadzącego eksperyment

- **Dostęp:** Hasło (ekran logowania)
- **Moduły:**

#### 5.3.1. Rejestracja uczestnika

```
┌─────────────────────────────────┐
│  Nowy uczestnik                  │
│                                  │
│  Kod: [U__]  (auto-sugestia)    │
│  Sekwencja: (•) AB  ( ) BA     │
│    [losuj automatycznie]        │
│  Status: ( ) Student (•) Ratow. │
│  Wiek: [▾ 26-30]               │
│  Płeć: [▾ Mężczyzna]           │
│  Staż: [▾ 4-7 lat]             │
│  Doświadczenie MCI: [▾ Nie]    │
│                                  │
│  [Zarejestruj]                  │
└─────────────────────────────────┘
```

Przycisk "losuj" automatycznie przydziela AB lub BA (balansując grupy 50/50).

#### 5.3.2. Wpis sesji

```
┌────────────────────────────────────────────────────────┐
│  Sesja eksperymentalna                                  │
│                                                         │
│  Uczestnik: [▾ U01 — Student, AB]                      │
│  Sesja: (•) 1  ( ) 2                                   │
│  → Metoda: TRADYCYJNA  Scenariusz: S1  (auto z AB+nr) │
│                                                         │
│  Timer:  ┌──────────────────────────────────────┐      │
│          │       ⏱  07:03                        │      │
│          │                                       │      │
│          │   [▶ START]          [⏹ STOP]        │      │
│          └──────────────────────────────────────┘      │
│  Czas: 423 sek (auto z timera, edytowalny ręcznie)    │
│                                                         │
│  Klasyfikacje:                                          │
│  ┌──────────┬──────┬─────────────────┐                 │
│  │ Pacjent  │ Kat. │ Ref.  │ Status  │                 │
│  ├──────────┼──────┼───────┼─────────┤                 │
│  │ S1-01    │ [▾]  │  T3   │         │                 │
│  │ S1-02    │ [▾]  │  T3   │         │                 │
│  │ ...      │      │       │         │                 │
│  │ S1-12    │ [▾]  │  T4   │         │                 │
│  └──────────┴──────┴───────┴─────────┘                 │
│                                                         │
│  Raport — kompletność:                                  │
│  ☐ Nazwa zdarzenia    ☐ Czas rozpoczęcia               │
│  ☐ Liczba poszkod.    ☐ Podział T1-T4                  │
│  ☐ Lista pacjentów    ☐ Info o transportach            │
│  ☐ Szpitale           ☐ Zespoły ZRM                    │
│  ☐ Status transportów ☐ Czas raportu                   │
│                                                         │
│  [Zapisz sesję]                                         │
│                                                         │
│  → Wynik: 10/12 (83.3%), raport: 4/10                  │
│    Błędy: S1-09 (T2→T1 overtriage), S1-10 (T2→T1 ...)│
└────────────────────────────────────────────────────────┘
```

Logika automatyczna:
- Gdy wybierzesz uczestnika i numer sesji → metoda i scenariusz wypełniają się automatycznie z sekwencji:
  - AB, sesja 1 → tradycyjna, S1
  - AB, sesja 2 → aplikacja, S2
  - BA, sesja 1 → aplikacja, S1
  - BA, sesja 2 → tradycyjna, S2
- **Timer wbudowany:**
  - Przycisk **START** → zapisuje `timer_started_at` (Date.now()), uruchamia wyświetlacz MM:SS odliczający w górę
  - Przycisk **STOP** → zapisuje `timer_stopped_at`, oblicza `time_seconds = Math.round((stop - start) / 1000)`
  - Wynik wyświetlony i wpisany automatycznie w pole czasu
  - Pole czasu pozostaje **edytowalne ręcznie** (awaryjnie, np. jeśli zapomnisz kliknąć START)
  - Timer działa na frontendzie (JS `setInterval`), timestampy wysyłane do bazy
- Kolumna "Ref." (kategoria referencyjna) widoczna ale nie edytowalna — do kontroli po wpisaniu
- Po zapisie: natychmiastowy feedback z wynikiem trafności i listą błędów

#### 5.3.3. Lista uczestników

```
┌──────────────────────────────────────────────────────────────┐
│  Uczestnicy (24/40)                                          │
│                                                               │
│  Kod  │ Sekw. │ Status   │ Sesja 1      │ Sesja 2     │ SUS │
│  U01  │ AB    │ Student  │ ✅ 423s 83%  │ ✅ 312s 92% │ 87  │
│  U02  │ BA    │ Ratownik │ ✅ 298s 100% │ ✅ 387s 75% │ 72  │
│  U03  │ AB    │ Student  │ ✅ 510s 67%  │ ⏳ —        │ —   │
│  U04  │ BA    │ Ratownik │ ⏳ —         │ —           │ —   │
│  ...                                                         │
└──────────────────────────────────────────────────────────────┘
```

### 5.4. `/dashboard` — Wyniki i eksport

- **Dostęp:** Hasło (to samo co panel)
- **Sekcje:**

#### Ankieta
- Liczba odpowiedzi (live counter)
- Rozkłady odpowiedzi na kluczowe pytania (wykresy słupkowe)
- Średnia ocena przydatności, gotowość do adopcji

#### Eksperyment
- Tabela zbiorcza: czas / trafność / kompletność raportu × metoda
- Średnie, mediany, SD
- Overtriage / undertriage wg metody

#### SUS
- Średni wynik SUS
- Rozkład wyników (histogram)
- Preferencja metody (pie chart)

#### Eksport
- Przyciski: `[Pobierz CSV — ankieta]` `[Pobierz CSV — eksperyment]` `[Pobierz CSV — SUS]` `[Pobierz wszystko ZIP]`
- Format CSV gotowy do importu do SPSS/Statistica/R

## 6. Obliczenia automatyczne

### 6.1. Trafność klasyfikacji

```js
// Po zapisie sesji
function calculateAccuracy(classifications, referenceKey) {
  let correct = 0;
  let overtriage = 0;
  let undertriage = 0;
  const errors = [];
  const priority = { T1: 4, T2: 3, T3: 2, T4: 1 };

  for (const c of classifications) {
    const ref = referenceKey[c.patient_code];
    if (c.assigned_category === ref) {
      correct++;
    } else {
      const type = priority[c.assigned_category] > priority[ref]
        ? 'overtriage' : 'undertriage';
      errors.push({
        patient_code: c.patient_code,
        assigned: c.assigned_category,
        correct: ref,
        type
      });
      if (type === 'overtriage') overtriage++;
      else undertriage++;
    }
  }

  return {
    correct,
    total: classifications.length,
    accuracy: (correct / classifications.length * 100).toFixed(2),
    overtriage,
    undertriage,
    errors
  };
}
```

### 6.2. Wynik SUS

```js
function calculateSUS({ q1,q2,q3,q4,q5,q6,q7,q8,q9,q10 }) {
  const odd  = (q1-1) + (q3-1) + (q5-1) + (q7-1) + (q9-1);
  const even = (5-q2) + (5-q4) + (5-q6) + (5-q8) + (5-q10);
  return (odd + even) * 2.5;
}
```

### 6.3. Kompletność raportu

```js
function calculateReportCompleteness(scores) {
  return Object.values(scores).filter(v => v === true).length; // 0-10
}
```

### 6.4. Timer sesji (frontend)

```js
let timerInterval = null;
let startedAt = null;

function startTimer() {
  startedAt = Date.now();
  document.getElementById('timerDisplay').textContent = '00:00';
  document.getElementById('btnStart').disabled = true;
  document.getElementById('btnStop').disabled = false;

  timerInterval = setInterval(() => {
    const elapsed = Math.floor((Date.now() - startedAt) / 1000);
    const min = String(Math.floor(elapsed / 60)).padStart(2, '0');
    const sec = String(elapsed % 60).padStart(2, '0');
    document.getElementById('timerDisplay').textContent = `${min}:${sec}`;
  }, 1000);
}

function stopTimer() {
  clearInterval(timerInterval);
  const stoppedAt = Date.now();
  const seconds = Math.round((stoppedAt - startedAt) / 1000);

  document.getElementById('btnStart').disabled = false;
  document.getElementById('btnStop').disabled = true;
  document.getElementById('timeSeconds').value = seconds;

  // Zapisz timestampy do wysłania z formularzem
  document.getElementById('timerStartedAt').value = new Date(startedAt).toISOString();
  document.getElementById('timerStoppedAt').value = new Date(stoppedAt).toISOString();
}
```

### 6.5. Automatyczne przypisanie metody/scenariusza

```js
function getSessionDetails(sequence, sessionNumber) {
  const map = {
    'AB-1': { method: 'tradycyjna', scenario: 'S1' },
    'AB-2': { method: 'aplikacja',  scenario: 'S2' },
    'BA-1': { method: 'aplikacja',  scenario: 'S1' },
    'BA-2': { method: 'tradycyjna', scenario: 'S2' },
  };
  return map[`${sequence}-${sessionNumber}`];
}
```

## 7. Format eksportu CSV

### 7.1. experiment.csv

```csv
participant_code,sequence,status,age_range,gender,work_experience,has_real_mci,session,method,scenario,timer_started_at,timer_stopped_at,time_seconds,accuracy_pct,correct_count,overtriage_count,undertriage_count,report_score,S_01,S_02,S_03,S_04,S_05,S_06,S_07,S_08,S_09,S_10,S_11,S_12
U01,AB,student,20-25,mezczyzna,brak,false,1,tradycyjna,S1,2026-04-15T14:32:10Z,2026-04-15T14:39:13Z,423,83.33,10,1,1,4,T3,T3,T3,T3,T2,T2,T2,T1,T2,T1,T4,T4
U01,AB,student,20-25,mezczyzna,brak,false,2,aplikacja,S2,2026-04-15T14:55:01Z,2026-04-15T15:00:13Z,312,91.67,11,1,0,9,T3,T3,T3,T3,T2,T2,T2,T1,T1,T2,T4,T4
```

Każdy wiersz = jeden uczestnik × jedna sesja. Kolumny S_01-S_12 to przypisane kategorie (do analizy macierzy pomyłek).

### 7.2. survey.csv

Jeden wiersz = jeden respondent. Kolumny = wszystkie pola z `survey_responses`.

### 7.3. sus.csv

```csv
participant_code,sequence,status,q1,q2,q3,q4,q5,q6,q7,q8,q9,q10,sus_score,preferred_method,preference_reason
U01,AB,student,4,2,5,1,4,2,5,1,4,1,87.50,aplikacja,Szybciej i pewniej
```

## 8. Bezpieczeństwo

| Aspekt | Rozwiązanie |
|--------|-------------|
| Ankieta — anonimowość | Nie zbieraj IP, nie wymagaj logowania |
| Panel — dostęp | Hasło w env `PANEL_PASSWORD`, prosty token w cookie/header |
| Baza — połączenie | Connection string w env `DATABASE_URL`, nigdy w kodzie |
| CORS | Ogranicz do domeny produkcyjnej |
| Rate limiting | Prosty limit na POST /api/survey (max 1/min z tego samego fingerprinta) |
| SQL injection | Parametryzowane zapytania (prepared statements) — NIGDY konkatenacja |
| Dane osobowe | Brak — ankieta anonimowa, eksperyment na kodach (U01, U02) |

## 9. Struktura plików (sugestia)

```
research-platform/
├── public/
│   ├── index.html          — landing page z linkami
│   ├── ankieta.html        — formularz ankiety
│   ├── sus.html            — formularz SUS
│   ├── panel.html          — panel prowadzącego
│   ├── dashboard.html      — wyniki i eksport
│   ├── css/
│   │   └── style.css       — wspólne style (dark theme jak TRIAGE MCI)
│   └── js/
│       ├── ankieta.js
│       ├── sus.js
│       ├── panel.js
│       └── dashboard.js
├── api/
│   ├── server.js           — Express app + routing
│   ├── db.js               — Neon connection pool
│   ├── routes/
│   │   ├── survey.js
│   │   ├── participants.js
│   │   ├── sessions.js
│   │   ├── sus.js
│   │   ├── results.js
│   │   └── export.js
│   └── utils/
│       ├── calculations.js — SUS score, accuracy, report score
│       └── auth.js         — middleware sprawdzające hasło
├── schema.sql              — pełny schemat bazy (z tego dokumentu)
├── seed.sql                — INSERT reference_key (z tego dokumentu)
├── package.json
├── .env.example            — DATABASE_URL, PANEL_PASSWORD
└── README.md
```

## 10. Deployment (Vercel)

```json
// vercel.json
{
  "rewrites": [
    { "source": "/api/(.*)", "destination": "/api/server.js" }
  ]
}
```

1. `npm init -y && npm install express @neondatabase/serverless`
2. Skopiuj `DATABASE_URL` z Neon Dashboard → Vercel Environment Variables
3. Ustaw `PANEL_PASSWORD` w Vercel Environment Variables
4. `vercel deploy`

## 11. Kolejność implementacji

1. **Baza danych** — utwórz projekt w Neon, uruchom `schema.sql` + `seed.sql`
2. **API** — `server.js` + routes, przetestuj Postmanem/curlem
3. **Ankieta** — najpilniejsza, bo zbieranie trwa 4-6 tygodni
4. **Panel** — przed dniem eksperymentu
5. **SUS** — przed dniem eksperymentu
6. **Dashboard + eksport** — może być po zebraniu danych

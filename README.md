# TRIAGE MCI

System segregacji i zarządzania zdarzeniem masowym (Mass Casualty Incident).

Progresywna aplikacja webowa (PWA) działająca w pełni offline na urządzeniach z systemem iOS, Android i Windows.

![Triage Icon](triage.jpg)

---

## Spis treści

- [Opis](#opis)
- [Funkcje](#funkcje)
- [Instalacja](#instalacja)
  - [Online (zalecane)](#online-zalecane)
  - [Offline — Windows](#offline--windows)
  - [Offline — iOS / Android](#offline--ios--android)
- [Jak używać](#jak-używać)
- [Struktura projektu](#struktura-projektu)
- [Technologie](#technologie)
- [English](#english)

---

## Opis

TRIAGE MCI to narzędzie do segregacji poszkodowanych podczas zdarzeń masowych, oparte na algorytmie **START** (Simple Triage and Rapid Treatment).

Aplikacja umożliwia:
- Segregację pacjentów wg protokołu START (T1–T4)
- Zarządzanie transportem i dysponowanie zespołów ratowniczych
- Generowanie raportów ze zdarzenia
- Pracę w pełni offline — bez potrzeby połączenia z internetem

Wszystkie dane są przechowywane lokalnie w przeglądarce i przetrwają zamknięcie / restart aplikacji.

---

## Funkcje

### Segregacja (Triage)
- 6-krokowy kreator decyzyjny START
- Klasyfikacja do 4 kategorii:
  - **T1 (Czerwony)** — Natychmiastowy — zagrożenie życia
  - **T2 (Żółty)** — Pilny — poważny, stabilny
  - **T3 (Zielony)** — Odroczony — lekkie obrażenia
  - **T4 (Czarny)** — Zgon / Expectant
- Ręczna zmiana koloru segregacji
- Automatyczne tagi pacjentów (P-001, P-002, ...)
- Notatki o obrażeniach i mechanizmie urazu

### Dysponowanie (Dispatch)
- Przypisywanie zespołów ZRM do pacjentów
- Kierowanie do szpitali
- Zarządzanie listą zespołów i szpitali
- Historia transportów
- Śledzenie statusu: Na miejscu → Transportowany → Dostarczony

### Raport
- Podsumowanie statystyk wg kategorii segregacji
- Grupowanie pacjentów wg szpitali
- Generowanie raportu tekstowego
- Kopiowanie do schowka jednym kliknięciem

### PWA i tryb offline
- Pełna funkcjonalność bez internetu
- Instalacja na ekranie głównym (iOS, Android, Windows)
- Automatyczny zapis danych w localStorage
- Wznawianie zdarzenia po zamknięciu aplikacji
- Potwierdzenia przed destrukcyjnymi akcjami (usuwanie pacjentów, zespołów, szpitali)

---

## Instalacja

### Online (zalecane)

1. Otwórz aplikację w przeglądarce (Chrome / Edge / Safari)
2. Zainstaluj jako PWA:
   - **Chrome / Edge**: kliknij ikonę instalacji na pasku adresu
   - **iOS Safari**: tapnij Udostępnij (⎙) → Dodaj do ekranu początkowego
3. Aplikacja działa offline od razu po pierwszym uruchomieniu

### Offline — Windows

Dla komputerów bez dostępu do internetu:

1. Skopiuj cały folder projektu na pendrive
2. Na docelowym komputerze uruchom `triage-server.exe`
3. Przeglądarka otworzy się automatycznie
4. Zainstaluj PWA z paska adresu
5. Zamknij serwer — aplikacja działa samodzielnie

Szczegółowa instrukcja: [INSTALACJA-OFFLINE.txt](INSTALACJA-OFFLINE.txt)

### Offline — iOS / Android

1. Podłącz urządzenie mobilne do tej samej sieci WiFi co komputer z uruchomionym `triage-server.exe`
2. Na komputerze sprawdź jego adres IP w sieci lokalnej
3. Na telefonie otwórz `http://<adres-ip>:8080`
4. Zainstaluj aplikację (Android: baner instalacji, iOS: Udostępnij → Dodaj do ekranu początkowego)

Alternatywnie: użyj hotspotu z telefonu, uruchom serwer na laptopie, otwórz na drugim urządzeniu.

---

## Jak używać

### Rozpoczęcie zdarzenia
1. Wpisz nazwę zdarzenia (np. "Wypadek A4 km 312")
2. Kliknij **ROZPOCZNIJ TRIAGE**

### Segregacja pacjenta
1. Odpowiadaj TAK/NIE na pytania kreatora START
2. Wynik segregacji wyświetli się automatycznie
3. Opcjonalnie zmień kolor ręcznie lub dodaj notatki
4. Kliknij **POTWIERDŹ I DODAJ** aby zapisać pacjenta

### Dysponowanie transportu
1. Przejdź do zakładki **DYSPONOWANIE**
2. Wybierz pacjenta, zespół ZRM i szpital docelowy
3. Kliknij **DYSPONUJ TRANSPORT**
4. Oznacz pacjenta jako dostarczonego po dotarciu do szpitala

### Wznawianie zdarzenia
- Przy ponownym uruchomieniu aplikacja wykryje zapisane dane
- Kliknij **KONTYNUUJ** aby wrócić do zdarzenia
- Kliknij **NOWE ZDARZENIE** aby rozpocząć od nowa (wymaga potwierdzenia)

---

## Struktura projektu

```
triage/
├── index.html              # Aplikacja (HTML + CSS + JS w jednym pliku)
├── manifest.json           # Manifest PWA
├── sw.js                   # Service Worker (cache-first, tryb offline)
├── triage.jpg              # Logo źródłowe
├── triage-server.exe       # Serwer lokalny do instalacji offline
├── INSTALACJA-OFFLINE.txt  # Instrukcja instalacji offline
├── icons/
│   ├── icon-192.png        # Ikona PWA 192x192
│   ├── icon-512.png        # Ikona PWA 512x512
│   └── apple-touch-icon.png # Ikona iOS 180x180
├── server/
│   └── main.go             # Kod źródłowy serwera
└── docs/
    └── plans/              # Dokumentacja techniczna
```

---

## Technologie

- **HTML/CSS/JS** — bez frameworków, bez zależności zewnętrznych
- **Service Worker** — strategia cache-first, pełny tryb offline
- **Web App Manifest** — instalacja na ekranie głównym
- **localStorage** — persystencja danych między sesjami
- **Go** — przenośny serwer HTTP do instalacji offline

---

## English

### What is TRIAGE MCI?

A Progressive Web App (PWA) for managing mass casualty incidents using the START triage protocol. Fully functional offline on iOS, Android, and Windows.

### Features
- **START triage wizard** — 6-step decision tree classifying patients into T1 (Red/Immediate), T2 (Yellow/Urgent), T3 (Green/Delayed), T4 (Black/Deceased)
- **Dispatch system** — assign ambulance teams to patients, route to hospitals, track transport status
- **Reports** — summary statistics, per-hospital grouping, copy-to-clipboard text reports
- **Full offline support** — service worker caches all assets, works without internet after first visit
- **Data persistence** — all state saved to localStorage, survives app restarts
- **Confirmation dialogs** — prevents accidental deletion of patients, teams, or hospitals
- **Cross-platform install** — installable as PWA on Android, iOS, and Windows

### Offline installation (no internet)

For devices without internet access, a portable Go server (`triage-server.exe`, 5.7 MB) is included. Copy the project folder to the target device, run the exe, install the PWA from localhost, and close the server. The app runs independently after that. See [INSTALACJA-OFFLINE.txt](INSTALACJA-OFFLINE.txt) for details.

### Tech stack
Vanilla HTML/CSS/JS (no frameworks), Service Worker, Web App Manifest, localStorage, Go (portable HTTP server).

---

## Kontakt

Zgłoszenia i propozycje ulepszeń: [jaqbek.eth@gmail.com](mailto:jaqbek.eth@gmail.com)

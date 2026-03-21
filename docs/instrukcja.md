# Instrukcja obsługi TRIAGE MCI

## 1. Czym jest TRIAGE MCI

TRIAGE MCI to aplikacja webowa (PWA) do segregacji i zarządzania zdarzeniem masowym (MCI) oparta na algorytmie START. Umożliwia:
- Klasyfikację pacjentów w 4 kategoriach (T1–T4)
- Zarządzanie transportem do szpitali
- Śledzenie pojemności szpitali
- Generowanie i wysyłanie raportów
- Import danych od dyspozytora

Aplikacja działa w pełni offline. Wszystkie dane przechowywane są wyłącznie na urządzeniu użytkownika — żadne informacje nie są wysyłane na zewnętrzne serwery.

Obsługuje 7 języków: polski, angielski, włoski, francuski, niemiecki, czeski, portugalski.

---

## 2. Rozpoczęcie zdarzenia

### Ekran startowy

Po uruchomieniu aplikacji pojawi się ekran konfiguracji zdarzenia:

1. **Wybór języka** — selektor u góry ekranu. Zmiana języka natychmiast tłumaczy cały interfejs.
2. **Nazwa zdarzenia** — pole obowiązkowe. Wpisz lokalizację i typ zdarzenia (np. „Wypadek A4 km 312"). Pojawi się w raportach.
3. **KAM** — pole opcjonalne. Kierujący Akcją Medyczną — imię i nazwisko osoby koordynującej działania na miejscu.
4. **Pokaż podpowiedzi** — przełącznik włączający dymki z kontekstowymi wskazówkami w całej aplikacji. Przydatne przy pierwszym użyciu.
5. **ROZPOCZNIJ TRIAGE** — rozpoczyna zdarzenie i uruchamia zegar.

### Wznowienie zdarzenia

Jeśli w pamięci urządzenia istnieje zapisane zdarzenie, przy uruchomieniu pojawi się opcja:
- **KONTYNUUJ** — wznawia zapisane zdarzenie ze wszystkimi danymi
- **NOWE ZDARZENIE** — kasuje poprzednie dane i rozpoczyna od nowa

---

## 3. Segregacja (algorytm START)

### Nawigacja

Aplikacja ma 3 zakładki w górnej części ekranu:
- **SEGREGACJA** — klasyfikacja pacjentów
- **DYSPONOWANIE** — transport i zarządzanie
- **RAPORT** — podsumowanie i eksport

Pasek statystyk pod nawigacją pokazuje liczbę pacjentów w każdej kategorii: 🔴 T1, 🟡 T2, 🟢 T3, ⚫ T4 oraz sumę.

### Kreator segregacji START

Kreator prowadzi przez 6-krokowy algorytm START. Na każdym kroku odpowiadasz TAK lub NIE:

**Krok 1 — Czy pacjent chodzi?**
- TAK → T3 Zielony (odroczony)
- NIE → przejdź do kroku 2

**Krok 2 — Czy oddycha spontanicznie?**
- TAK → przejdź do kroku 4 (ocena częstości oddechów)
- NIE → przejdź do kroku 3

**Krok 3 — Udrożnij drogi oddechowe. Czy oddycha po udrożnieniu?**
- TAK → T1 Czerwony (natychmiastowy)
- NIE → T4 Czarny (zgon)

**Krok 4 — Czy częstość oddechów jest prawidłowa (≤ 30/min)?**
- TAK → przejdź do kroku 5
- NIE → T1 Czerwony (natychmiastowy)

**Krok 5 — Czy nawrót kapilarny jest prawidłowy (≤ 2 sek.) i tętno na tt. promieniowej obecne?**
- TAK → przejdź do kroku 6
- NIE → T1 Czerwony (natychmiastowy)

**Krok 6 — Czy spełnia proste polecenia?**
- TAK → T2 Żółty (pilny)
- NIE → T1 Czerwony (natychmiastowy)

Każde pytanie ma podpowiedź (hint) widoczną pod pytaniem — opisuje jak ocenić dany parametr.

### Karta wyniku

Po zakończeniu segregacji pojawia się karta wyniku z:

- **Kolorowy znacznik** — duży kwadrat z kategorią (T1/T2/T3/T4)
- **Ścieżka decyzyjna** — pokazuje przebieg algorytmu
- **Płeć** — przyciski M (mężczyzna), K (kobieta), ? (nieznana)
- **Wiek** — pole numeryczne z przyciskami -5, -1, +1, +5 (zakres 0–150)
- **Obrażenia ciała** — przycisk otwierający diagram obrażeń (patrz sekcja 4)
- **Notatki** — pole na dodatkowe informacje (obrażenia, mechanizm urazu itp.)

Dostępne akcje:
- **✓ ZATWIERDŹ PACJENTA** — dodaje pacjenta do listy
- **↺ OD NOWA** — resetuje kreator do kroku 1
- **ZMIEŃ KOLOR RĘCZNIE** — otwiera rząd 4 przycisków (T1–T4) do ręcznej zmiany kategorii

### Lista pacjentów na miejscu

Poniżej kreatora wyświetla się lista „Pacjenci na miejscu" posortowana wg priorytetu (T1 → T2 → T3 → T4).

Każda karta pacjenta zawiera:
- Kolorowy znacznik (np. P-001)
- Kategorię, płeć/wiek, czas rejestracji
- Pole notatek (edytowalne)
- Przycisk obrażeń (🩻) z podsumowaniem
- 4 kropki retriażu (T1–T4) — kliknij aby zmienić kategorię
- Przycisk usunięcia (✕)
- Znacznik retriażu (↺n) jeśli kategoria była zmieniana

### Lista w transporcie

Sekcja „W transporcie" pojawia się gdy pacjenci są w drodze do szpitala. Każda karta pokazuje:
- Znacznik pacjenta z kolorem kategorii
- Trasę: Zespół → Szpital → czas wyjazdu
- Przycisk „zmień" — zmiana szpitala docelowego
- Przycisk „✓ Dostarczony" — oznacza dostarczenie i zwalnia zespół ZRM

---

## 4. Diagram obrażeń ciała

Przycisk „🩻 Obrażenia ciała" otwiera interaktywny diagram ciała.

### Widok diagramu

Wyświetlane są dwa zarysy ciała obok siebie:
- **PRZÓD** (front) — głowa, klatka piersiowa, brzuch, ramiona, nogi
- **TYŁ** (back) — głowa, górna część pleców, dolna część pleców, ramiona, nogi

Po bokach widoczne są oznaczenia stron: **L** (lewa) i **P** (prawa), odwrócone dla widoku tyłu.

Kliknij na strefę ciała aby ją wybrać.

### Panel obrażeń strefy

Po wybraniu strefy otwiera się panel z 7 typami urazów do zaznaczenia:
- Złamanie
- Krwawienie
- Oparzenie
- Rana
- Opuchlizna
- Ból
- Amputacja

Zaznaczone typy oznaczone są ✓. Każdy można włączyć/wyłączyć kliknięciem.

Na dole panelu:
- **✕ NIE** — odrzuca zmiany w tej strefie
- **✓ TAK** — zatwierdza obrażenia dla wybranej strefy

### Zatwierdzanie diagramu

Po zaznaczeniu obrażeń we wszystkich potrzebnych strefach:
- **✓ TAK** — zatwierdza cały diagram obrażeń
- **✕ NIE** — odrzuca wszystkie zmiany i wraca do stanu sprzed otwarcia

Zaznaczone obrażenia widoczne są:
- Na karcie pacjenta jako skrócone podsumowanie (np. „🩻 Głowa: Złamanie, Krwawienie")
- W raporcie przy danym pacjencie

---

## 5. Dysponowanie

### Modalne okna konfiguracji

Przy pierwszym wejściu w zakładkę Dysponowanie pojawiają się kolejno 3 modale:

**1. Kierownik Działań Medycznych (GDM)**
- Pole na imię i nazwisko GDM
- Przycisk importu danych od dyspozytora (patrz sekcja 6)
- POMIŃ lub ZAPISZ

**2. Zespoły ZRM**
- Pytanie: „Czy znasz zespoły ZRM zadysponowane na miejsce?"
- TAK → formularz dodawania zespołów (nazwa + przycisk dodaj)
- NIE → użyj domyślnych zespołów (S-01, S-02, P-01, P-02, P-03, LPR)
- ZATWIERDŹ lub POMIŃ (DOMYŚLNE)

**3. Szpitale docelowe**
- Formularz dodawania szpitali z pojemnością:
  - Nazwa szpitala
  - 🔴 Pojemność czerwona (T1) — liczba łóżek
  - 🟡 Pojemność żółta (T2) — liczba łóżek
- ZATWIERDŹ lub POMIŃ (DOMYŚLNE)

Każdy modal pojawia się tylko raz na zdarzenie.

### Formularz dysponowania

Sekcja „Zadysponuj transport" zawiera:

1. **Pacjent** — lista rozwijana z pacjentami na miejscu, posortowana wg priorytetu
2. **Zespół ZRM** — lista dostępnych zespołów (zajęte zespoły są ukryte)
3. **Szpital** — lista szpitali z informacją o zajętości:
   - Format: „Nazwa szpitala [🔴 zajęte/wszystkie 🟡 zajęte/wszystkie]"
   - Szpitale bez wolnych miejsc w kategorii pacjenta oznaczone ostrzeżeniem
4. **DYSPONUJ TRANSPORT** — przycisk aktywny gdy wszystkie pola wypełnione

### Panele zarządzania

Pod formularzem znajdują się dwa panele (na szerokich ekranach obok siebie, na wąskich pod sobą):

**Szpitale** (łóżka WSZYSTKIE):
- Lista szpitali z edytowalnymi polami pojemności (🔴 i 🟡)
- Dodawanie nowych szpitali
- Usuwanie szpitali (✕)
- Zwijanie/rozwijanie listy

**Zespoły ZRM:**
- Lista zespołów z przyciskami usuwania
- Dodawanie nowych zespołów
- Zwijanie/rozwijanie listy

### Historia transportów

Na dole sekcji dysponowania (domyślnie zwinięta). Rozwiń aby zobaczyć chronologiczną listę wszystkich transportów z:
- Znacznik pacjenta z kolorem
- Zespół → Szpital → czas wyjazdu
- Status (w transporcie / dostarczony)

---

## 6. Import danych od dyspozytora

Funkcja pozwala dyspozytorowi przesłać dane (nazwę zdarzenia, GDM, zespoły ZRM, szpitale z pojemnością) do ratownika w terenie.

### Metoda 1 — Link SMS (jedno kliknięcie)

Dyspozytor wysyła SMS z linkiem w formacie:
```
https://0xjaqbek.github.io/triage/?i=DANE_BASE64
```

Po kliknięciu linku:
1. Aplikacja otwiera się z cache (nie wymaga internetu jeśli zainstalowana)
2. Pojawia się podgląd danych: nazwa zdarzenia, GDM, zespoły, szpitale
3. **IMPORTUJ** — importuje dane i otwiera segregację
4. **ODRZUĆ** — ignoruje dane

### Metoda 2 — Wklejenie danych

Jeśli link nie otwiera aplikacji (zablokowane urządzenia):
1. Otwórz aplikację, przejdź do Dysponowanie
2. W modalu GDM kliknij **📥 IMPORTUJ DANE OD DYSPOZYTORA**
3. Wklej treść SMS (cały link lub sam kod Base64)
4. Podgląd danych pojawi się automatycznie
5. **IMPORTUJ** — zatwierdza dane

### Import w trakcie zdarzenia

Jeśli zdarzenie jest już rozpoczęte i masz posegregowanych pacjentów:
- **Pacjenci** — zachowani (nigdy nie są usuwani)
- **KAM** — zachowany
- **GDM** — zastąpiony danymi dyspozytora
- **Zespoły ZRM** — zastąpione danymi dyspozytora
- **Szpitale** — zastąpione danymi dyspozytora
- **Nazwa zdarzenia** — połączone: „Twoja nazwa | Nazwa dyspozytora (dyspozytor)"

Żółte ostrzeżenie w podglądzie informuje o zachowaniu pacjentów i KAM.

---

## 7. Tryb dyspozytora

Dostępny przez link „Tryb dyspozytora →" na ekranie startowym lub pod adresem `/dyspozytor/`.

Strona pozwala dyspozytorowi przygotować i wysłać dane dla zespołu w terenie.

### Formularz

1. **Nazwa zdarzenia** — np. „MCI Kraków-Nowa Huta"
2. **Główny Dyspozytor Medyczny (GDM)** — imię i nazwisko
3. **Zespoły ZRM** — dodawaj po jednym (nazwa + przycisk „+")
4. **Szpitale docelowe** — nazwa + pojemność 🔴 i 🟡 (+ przycisk dodaj)

### Wysyłanie

- **📋 Podgląd danych** — pokazuje podsumowanie z rozmiarem w znakach i liczbą SMS
- **📱 Wyślij SMS** — otwiera aplikację SMS z gotowym linkiem
- **📋 Kopiuj dane** — kopiuje link do schowka

Dane kodowane są w Base64 (kompatybilność z SMS GSM-7 — do 1530 znaków w 10 segmentach SMS). Pomieści ok. 20 szpitali i 20 zespołów.

---

## 8. Raport

### Raport wizualny

Sekcja „Raport dla dyspozytora" generuje się automatycznie i zawiera:

**Nagłówek:**
- Nazwa zdarzenia, czas rozpoczęcia, czas raportu
- KAM i GDM (jeśli podane)
- Czas zakończenia (jeśli zdarzenie zakończone)

**Podsumowanie:**
- Kolorowe znaczniki z liczbą pacjentów w każdej kategorii
- Łączna liczba pacjentów

**Szpitale** (osobny blok dla każdego szpitala z transportami):
- Nazwa szpitala, liczba pacjentów, info o pojemności
- Lista pacjentów z: znacznik, wiek/płeć, zespół, czas, status (W TRANSPORCIE / DOSTARCZONY)
- Historia retriażu (np. ↺T1→T2→T1)
- Historia zmian szpitala (np. ⇄Szpital1→Szpital2)
- Podsumowanie obrażeń

**Pacjenci na miejscu** (jeśli ktoś pozostał):
- Lista z: znacznik, wiek/płeć, kategoria, notatki, obrażenia

### Raport tekstowy

Przycisk „📋 Raport tekstowy" otwiera modal z pełnym raportem w formacie tekstowym. Zawiera kompletną listę pacjentów z notatkami, płcią, wiekiem i ścieżką segregacji.

### Wysyłanie raportu

- **📋 Kopiuj do schowka** — kopiuje pełny tekst raportu
- **📤 Wyślij** — otwiera opcje:
  - **SMS** — otwiera natywną aplikację SMS z treścią raportu
  - **Email** — otwiera klienta email z tematem i treścią raportu

### Zakończenie zdarzenia

Przycisk **ZAKOŃCZ ZDARZENIE** w nagłówku raportu:
- Zapisuje czas zakończenia
- Zmienia się na **COFNIJ ZAKOŃCZENIE** (odwracalne)
- Czas zakończenia pojawia się w raporcie

---

## 9. Retriaż i zmiana szpitala

### Retriaż

Na karcie pacjenta w sekcji „Pacjenci na miejscu" znajdują się 4 kolorowe kropki (T1–T4). Kliknięcie innej kategorii niż aktualna:
1. Otwiera okno potwierdzenia: „Zmienić kategorię pacjenta P-001 z T1 na T2?"
2. **POTWIERDŹ** — zmienia kategorię, zapisuje w historii
3. **ANULUJ** — bez zmian

Historia retriażu:
- Znacznik ↺n na karcie pacjenta (n = liczba zmian)
- Pełna historia w raporcie (np. T1→T2→T1 z godzinami)

### Zmiana szpitala w transporcie

Na karcie pacjenta „W transporcie" przycisk „zmień":
1. Otwiera modal z aktualnym szpitalem
2. Lista rozwijana z dostępnymi szpitalami (z informacją o pojemności)
3. **POTWIERDŹ** — zmienia szpital, zapisuje w historii
4. **ANULUJ** — bez zmian

Historia zmian szpitala widoczna w raporcie (⇄Szpital1→Szpital2).

---

## 10. Instalacja i tryb offline

### Instalacja jako PWA

**Android (Chrome):**
1. Otwórz aplikację w Chrome
2. Tapnij ⋮ (menu) w prawym górnym rogu
3. Wybierz „Dodaj do ekranu głównego"
4. Tapnij „Zainstaluj"

**iOS (Safari):**
1. Otwórz aplikację w Safari
2. Tapnij ⎙ (Udostępnij)
3. Przewiń w dół i wybierz „Dodaj do ekranu początkowego"

Po instalacji aplikacja działa jako osobna aplikacja z pełnym trybem offline.

### Tryb offline

Aplikacja po pierwszym załadowaniu zapisuje się w cache urządzenia. Wszystkie funkcje działają bez połączenia z internetem:
- Segregacja pacjentów
- Dysponowanie transportu
- Generowanie raportów
- Import danych (wklejanie)

Internet potrzebny jest tylko do: pierwszego pobrania, aktualizacji, oraz wysyłki SMS/email (przez natywne aplikacje urządzenia).

### Pobranie jako plik HTML

Przycisk **⬇ Pobierz** na ekranie startowym pobiera aplikację jako samodzielny plik HTML. Można go przenieść na pendrive i uruchomić na dowolnym komputerze z przeglądarką.

---

## 11. Prywatność

- Wszystkie dane przechowywane są **wyłącznie na urządzeniu** (localStorage)
- Żadne dane **nie są wysyłane** na serwery zewnętrzne
- Brak analityki, śledzenia, telemetrii
- Brak plików cookie do śledzenia
- Funkcje SMS i email korzystają z natywnych aplikacji urządzenia — aplikacja nie wysyła wiadomości bezpośrednio
- Usunięcie danych: użyj „Zakończ zdarzenie" lub wyczyść dane przeglądarki

Pełna polityka prywatności dostępna pod adresem `/privacy/`.

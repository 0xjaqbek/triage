# Skrypt eksperymentu — Standardowy protokół prowadzącego (rola DM/GDM)

**Wersja:** 2.0
**Data:** 2026-03-15
**Przeznaczenie:** Wyłącznie dla prowadzącego eksperyment. NIE udostępniać uczestnikom.

---

## A. Zasady ogólne

### A.1. Rola prowadzącego

Prowadzący pełni jednocześnie dwie funkcje:
1. **Obserwator** — mierzy czas, ocenia wyniki, wypełnia Kartę ocen
2. **Dyspozytor Medyczny (DM)** — symuluje uproszczoną rolę DM-W/GDM z Procedury MZ v2.3

### A.2. Zasady standaryzacji

- **Czytaj z karty.** Każdy komunikat DM odczytuj dosłownie z tego skryptu. Nie improwizuj.
- **Styl radiowy.** Zwracaj się do uczestnika: „Kierujący Akcją Medyczną". Uczestnik zwraca się do Ciebie: „Dyspozytor" lub „Dyspozytor Wysyłający".
- **Nie pomagaj.** Nigdy nie sugeruj kategorii segregacyjnej, nie podpowiadaj dyslokacji, nie komentuj decyzji uczestnika.
- **Nie poganiaj.** Nie mów „proszę się pospieszyć", „zostało mało czasu" itp.
- **Pytania spoza skryptu → odpowiedź domyślna:** „Proszę opierać się na informacjach, które Pan/Pani posiada."
- **Timer:** Osobny stoper na każdą podfazę (A, C1, C2, D). Start/stop. Wymiany DM nie są mierzone. Zapisuj czasy na Arkuszu czasów.

### A.3. Identyczne warunki

Poniższe elementy MUSZĄ być identyczne dla każdego uczestnika, niezależnie od metody i scenariusza:
- Treść komunikatów DM (ten skrypt)
- Dane o szpitalach (Karta informacyjna szpitali)
- Lista ZRM (4 zespoły)
- Czas instruktażu (5 min aplikacja + 2 min rola DM)
- Pomieszczenie i warunki

---

## B. Checklist przygotowania (przed każdą sesją)

### Materiały na stole uczestnika

**Sesja tradycyjna (papierowa):**
- [ ] Koperta z 12 kartami pacjentów (potasowane!)
- [ ] Papierowa Tabela Dyslokacji (szablon)
- [ ] Pusta kartka A4 + długopis (na raport)
- [ ] ~~Karta informacyjna szpitali~~ (podasz po segregacji, nie wcześniej!)

**Sesja z aplikacją:**
- [ ] Koperta z 12 kartami pacjentów (potasowane!)
- [ ] Smartfon z TRIAGE MCI v2.4.0 — wersja eksperymentalna (`experiment.html`), ekran startowy
- [ ] Przygotowany SMS z danymi scenariuszowymi (z `dyspozytor/experiment.html`) — wyślij PRZED sesją, uczestnik tapnie link w Podfazie B
- [ ] ~~Karta informacyjna szpitali~~ (uczestnik otrzyma dane przez SMS, nie kartę papierową)

### Materiały prowadzącego

- [ ] Ten skrypt (wydrukowany)
- [ ] Stoper z funkcją międzyczasów
- [ ] Arkusz czasów (jeden wiersz na sesję)
- [ ] Klucz referencyjny — segregacja (z scenariusze-eksperyment.md)
- [ ] Klucz referencyjny — dyslokacja (z scenariusze-eksperyment.md)
- [ ] Karta ocen (formularz obserwatora, jeden na sesję)
- [ ] Lista kontrolna kompletności raportu (Raport GDM)
- [ ] Karta informacyjna szpitali (do podania słownie w sesji tradycyjnej)
- [ ] Przygotowane linki SMS z danymi scenariuszowymi (Scenariusz 1 i 2) — wygenerowane z `dyspozytor/experiment.html`
- [ ] Dostęp do urządzenia uczestnika w celu aktywacji trybu REC (lub instrukcja dla uczestnika)

---

## C. Przebieg sesji — skrypt krok po kroku

---

### FAZA 0: Wprowadzenie (przed pierwszą sesją danego uczestnika)

**Czas: ~7 minut. Timer NIE działa.**

#### Krok 0.1 — Zgoda i cel badania

Odczytaj:

> Dziękuję za udział w badaniu. Celem tego eksperymentu jest porównanie dwóch metod realizacji zadań Kierującego Akcją Medyczną podczas zdarzenia z dużą liczbą poszkodowanych: metody papierowej i metody cyfrowej z wykorzystaniem aplikacji TRIAGE MCI.
>
> Udział jest dobrowolny. Może Pan/Pani wycofać się na każdym etapie bez podania przyczyny. Dane są anonimowe — wyniki przypisane są do kodu uczestnika, nie do nazwiska. Czy wyraża Pan/Pani zgodę na udział?

*(Uzyskaj podpis na formularzu zgody.)*

#### Krok 0.2 — Instrukcja aplikacji (5 min, tylko przed pierwszą sesją)

Odczytaj:

> Teraz pokażę Panu/Pani krótko aplikację TRIAGE MCI. Będzie Pan/Pani z niej korzystać w jednej z dwóch sesji.

Pokaż na telefonie demo:
1. Ekran startowy → pola: nazwa zdarzenia + KAM (Kierujący Akcją Medyczną) → ROZPOCZNIJ TRIAGE
2. Kreator START — pytania TAK/NIE → wynik segregacji → opcjonalnie: płeć (M/K/?) i wiek (stepper ±1/±5) → ZATWIERDŹ
3. Lista pacjentów — kropki kolorów = zmiana kategorii (retriage) z potwierdzeniem
4. Moduł DYSPONOWANIE — dodawanie szpitali (z pojemnością RED/YELLOW) i ZRM, przypisanie pacjenta do ZRM i szpitala, przycisk „zmień" przy szpitalu w transporcie
5. Modal GDM — pojawia się raz przy pierwszym wejściu w DYSPONOWANIE (imię i nazwisko GDM)
6. Moduł RAPORT — generowanie raportu + przycisk ZAKOŃCZ ZDARZENIE (odwracalny)

**Nie komentuj zalet ani wad aplikacji. Pokaż, jak działają przyciski.**

#### Krok 0.3 — Wyjaśnienie roli DM

Odczytaj:

> W trakcie każdej sesji będę pełnił rolę Dyspozytora Medycznego. Proszę traktować naszą komunikację jak łączność radiową. Po zakończeniu segregacji proszę przekazać mi raport z rozpoznania wstępnego — ile poszkodowanych w poszczególnych grupach. Ja przekażę Panu/Pani informacje o zadysponowanych zespołach ZRM i możliwościach szpitali. Następnie wspólnie ustalimy dyslokację poszkodowanych. Na końcu poproszę o sporządzenie raportu ze zdarzenia.
>
> Proszę pamiętać: w trakcie segregacji pracuje Pan/Pani samodzielnie. Kontakt ze mną następuje po jej zakończeniu.

#### Krok 0.4 — Wyjaśnienie pomiaru

Odczytaj:

> Będę mierzyć czas osobno w każdym etapie zadania: segregacja, dyslokacja i raport. Nie ma limitu czasu, ale proszę pracować tak szybko, jak to możliwe przy zachowaniu dokładności. Czas naszej wzajemnej komunikacji nie jest mierzony.
>
> Czy ma Pan/Pani pytania zanim rozpoczniemy?

*(Odpowiedz na pytania organizacyjne. NIE odpowiadaj na pytania merytoryczne dotyczące START ani procedury.)*

---

### FAZA 1: Sesja

**Zasada pomiaru czasu:** Każda podfaza mierzona **osobnym stoperem** (start/stop). Między podfazami stoper nie działa — wymiana DM jest stałą i nie jest mierzona.

---

#### PODFAZA A: Segregacja pierwotna

**Odczytaj opis scenariusza** (z scenariusze-eksperyment.md — odpowiedni dla danej sesji).

Następnie odczytaj:

> Pełni Pan/Pani funkcję Kierującego Akcją Medyczną. Na miejscu zdarzenia znajduje się 12 poszkodowanych wymagających segregacji medycznej algorytmem START.
>
> [SESJA TRADYCYJNA]: Proszę otworzyć kopertę. Na podstawie kart pacjentów proszę przeprowadzić segregację i odnotować wyniki na Tabeli Dyslokacji — kolumny: numer, kategoria, płeć/wiek, obrażenia. Po zakończeniu segregacji proszę o przekazanie mi raportu z rozpoznania wstępnego.
>
> [SESJA Z APLIKACJĄ]: Proszę wpisać nazwę zdarzenia i swoje dane jako KAM w aplikacji, a następnie rozpocząć triage. Proszę otworzyć kopertę. Na podstawie kart pacjentów proszę przeprowadzić segregację w aplikacji TRIAGE MCI, przechodząc kreator START dla każdego pacjenta. Przy każdym pacjencie może Pan/Pani opcjonalnie zaznaczyć płeć i wiek. Po zakończeniu segregacji proszę o przekazanie mi raportu z rozpoznania wstępnego.
>
> Zaczynamy.

**→ START stopera [A]**

**W trakcie Podfazy A:**
- NIE inicjuj kontaktu z uczestnikiem
- Obserwuj i notuj na Karcie ocen
- Jeśli uczestnik się odezwie:

| Pytanie uczestnika | Odpowiedź DM |
|---|---|
| „Czy ten pacjent chodzi?" | „Informacja ta znajduje się na karcie." |
| „Co się stanie po udrożnieniu dróg oddechowych?" | „Informacja ta znajduje się na karcie." |
| Jakiekolwiek pytanie o stan pacjenta | „Wszystkie potrzebne informacje są na karcie pacjenta." |
| „Czy mogę już przekazać raport?" | „Proszę zakończyć segregację wszystkich poszkodowanych, a następnie przekazać raport." |
| „Ile mam czasu?" | „Nie ma limitu czasu. Proszę pracować tak szybko, jak to możliwe." |

**Trigger zakończenia:** Uczestnik mówi wariant: „Skończyłem segregację" / „Mam raport z rozpoznania" / „Gotowy do raportu"

**→ STOP stopera [A] → zapisz czas segregacji**

---

#### PODFAZA B: Raport z rozpoznania wstępnego + przekazanie danych

**⏱ Stoper NIE działa — wymiana DM jest stałą.**

**Krok B.1 — Przyjęcie raportu**

Odczytaj:

> Kierujący Akcją Medyczną, proszę o raport z rozpoznania wstępnego. Ile poszkodowanych w poszczególnych grupach segregacyjnych?

*(Uczestnik powinien podać: liczbę czerwonych, żółtych, zielonych, czarnych i łączną.)*

**Krok B.2 — Potwierdzenie i powtórzenie**

Odczytaj (wstawiając podane przez uczestnika liczby):

> Przyjmuję. Poszkodowani: grupa czerwona — [X], grupa żółta — [X], grupa zielona — [X], grupa czarna — [X]. Łącznie [X] poszkodowanych. Potwierdzam.

**WAŻNE:** Nie koryguj! Nawet jeśli uczestnik podał błędne liczby (np. 2 czerwonych zamiast 3), powtórz dokładnie to, co powiedział. Błąd w raporcie jest częścią mierzonego wyniku.

**Krok B.3 — Informacja o ZRM**

Odczytaj:

> Na miejsce zdarzenia zadysponowane zostały następujące zespoły ratownictwa medycznego:
> - ZRM P-01 Kraków — zespół podstawowy — na miejscu
> - ZRM P-02 Kraków — zespół podstawowy — na miejscu
> - ZRM S-01 Kraków — zespół specjalistyczny — na miejscu
> - LZRM Kraków — lotniczy zespół ratownictwa medycznego — na miejscu
>
> Łącznie cztery zespoły. Wszystkie obecne na miejscu zdarzenia.

**Krok B.4 — Informacja o szpitalach**

Odczytaj:

> Przekazuję informacje o szpitalach wyznaczonych do przyjęcia poszkodowanych.
>
> Szpital Uniwersytecki: SOR, Centrum Urazowe, bloki operacyjne — chirurgiczny i ortopedyczny, trzy stanowiska intensywnej terapii z respiratorem. Czas dojazdu dwanaście minut. Może przyjąć do trzech poszkodowanych z grupy czerwonej i do trzech z grupy żółtej.
>
> Szpital Miejski numer dwa: SOR, blok operacyjny — chirurgia ogólna, dwa stanowiska intensywnej terapii z respiratorem. Czas dojazdu osiemnaście minut. Może przyjąć do dwóch poszkodowanych z grupy czerwonej i do dwóch z grupy żółtej.
>
> Szpital Powiatowy: Izba Przyjęć, jedno stanowisko intensywnej terapii. Bez bloku operacyjnego. Czas dojazdu dwadzieścia pięć minut. Nie przyjmuje poszkodowanych z grupy czerwonej. Może przyjąć do trzech poszkodowanych z grupy żółtej.

**[SESJA TRADYCYJNA]:** *(Wręcz uczestnikowi wydrukowaną Kartę informacyjną szpitali.)*

**[SESJA Z APLIKACJĄ]:** *(Pomiń wręczanie karty. Przejdź do kroku B.5.)*

**Krok B.5 — Import danych od dyspozytora (tylko sesja z aplikacją)**

**⏱ Stoper NIE działa — import danych jest stałą procedurą, nie mierzoną zmienną.**

Odczytaj:

> [SESJA Z APLIKACJĄ]: Jako dyspozytor medyczny wysyłam teraz na Pana/Pani urządzenie SMS z danymi logistycznymi — informacjami o zespołach ZRM i szpitalach z ich pojemnościami. Proszę otworzyć otrzymanego SMS-a i tapnąć link. Dane zostaną zaimportowane automatycznie do aplikacji.
>
> [SESJA TRADYCYJNA]: *(pomiń ten krok)*

*(Wyślij przygotowany wcześniej SMS z linkiem z `dyspozytor/experiment.html`. Poczekaj, aż uczestnik tapnie link i dane się zaimportują. Jeśli link nie otwiera aplikacji — poproś uczestnika o otwarcie aplikacji ręcznie i wklejenie danych z SMS-a w modalu GDM → przycisk „Importuj dane od dyspozytora". Gdy dane są zaimportowane — przejdź do B.6.)*

**Krok B.6 — Polecenie dyslokacji**

Odczytaj:

> Proszę o przeprowadzenie dyslokacji poszkodowanych z grupy czerwonej i żółtej. Proszę przypisać poszkodowanych do zespołów ZRM i wskazać szpitale docelowe.

---

#### PODFAZA C: Dyslokacja (dwie tury)

Dyslokacja przebiega w **dwóch turach**, odzwierciedlając realistyczny cykl dysponowania ZRM:
- **Tura 1:** Uczestnik dysponuje 4 poszkodowanych (do 4 dostępnych ZRM)
- **Tura 2:** Po komunikacie DM o powrocie ZRM, uczestnik dysponuje pozostałych poszkodowanych

---

**TURA 1 — pierwsze zadysponowanie**

**→ START stopera [C1]** (natychmiast po wypowiedzeniu polecenia dyslokacji w B.6)

**Uczestnik pracuje:**
- Sesja tradycyjna: wypełnia kolumny ZRM i Szpital na Tabeli Dyslokacji
- Sesja z aplikacją: używa modułu DYSPONOWANIE (przy pierwszym wejściu pojawi się modal GDM — uczestnik wpisuje imię i nazwisko GDM lub pomija)

**Prowadzący (DM) jest dostępny.** Odpowiadaj na pytania wg poniższej tabeli:

| Pytanie uczestnika | Odpowiedź DM |
|---|---|
| „Czy mogę wysłać dwóch pacjentów jednym ZRM?" | „Jeden ZRM transportuje jednego poszkodowanego z grupy czerwonej. W uzasadnionych przypadkach możliwy jest transport dwóch poszkodowanych z grupy żółtej jednym ZRM." |
| „Co z pacjentami z grupy zielonej?" | „Poszkodowani z grupy zielonej mogą być transportowani innym środkiem transportu. Proszę skupić się na dyslokacji poszkodowanych z grupy czerwonej i żółtej." |
| „Co z pacjentami z grupy czarnej?" | „Poszkodowani z grupy czarnej pozostają na miejscu zdarzenia. Nie wymagają dysponowania ZRM." |
| „Czy mogą przyjechać dodatkowe ZRM?" | „Dysponuje Pan/Pani czterema zespołami. Nie ma możliwości zadysponowania dodatkowych zespołów w tym czasie." |
| „Czy szpital może przyjąć więcej pacjentów?" | „Proszę opierać się na informacjach o pojemności szpitali, które przekazałem." |
| „Który szpital jest najlepszy dla tego pacjenta?" | „Decyzja o dyslokacji należy do Kierującego Akcją Medyczną." |
| „Aplikacja pokazuje BRAK MIEJSC" | „To ostrzeżenie informacyjne. Decyzja o dyslokacji należy do Kierującego Akcją Medyczną." |
| „Jak wpisać GDM?" | „Proszę wpisać imię i nazwisko Głównego Dyspozytora Medycznego lub pominąć." |
| „Jak zmienić szpital?" | „Proszę kliknąć przycisk «zmień» obok nazwy szpitala w sekcji transportu." |
| Uczestnik przekazuje plan dyslokacji | „Przyjmuję. [Powtórz przypisanie]. Potwierdzam dyslokację." |

**Trigger zakończenia Tury 1:** Uczestnik zadysponował 4 poszkodowanych i mówi wariant: „Czterech zadysponowanych" / „Brak wolnych ZRM" / „Pozostali oczekują" — **LUB** próbuje zadysponować 5. pacjenta (wtedy przypomnij: „Wszystkie cztery zespoły ZRM są w drodze do szpitali.")

**→ STOP stopera [C1] → zapisz czas dyslokacji Tura 1**

---

**KOMUNIKAT DM — powrót ZRM:**

**⏱ Stoper NIE działa — wymiana DM jest stałą.**

Odczytaj:

> Kierujący Akcją Medyczną, informuję: wszystkie cztery zespoły ratownictwa medycznego dostarczyły poszkodowanych do szpitali i są w drodze powrotnej na miejsce zdarzenia. ZRM P-01, ZRM P-02, ZRM S-01 oraz LZRM — dostępne do ponownego zadysponowania. Proszę kontynuować dyslokację pozostałych poszkodowanych.

Następnie odczytaj instrukcję odpowiednią dla metody:

> **[SESJA TRADYCYJNA]:** Proszę kontynuować uzupełnianie Tabeli Dyslokacji dla pozostałych poszkodowanych.
>
> **[SESJA Z APLIKACJĄ]:** Proszę oznaczyć wszystkie cztery zespoły ZRM jako „dostarczony" w aplikacji, aby zwolnić je do ponownego zadysponowania. Następnie proszę przypisać pozostałych poszkodowanych.

---

**TURA 2 — ponowne zadysponowanie**

**→ START stopera [C2]** (natychmiast po zakończeniu komunikatu DM)

Uczestnik dysponuje pozostałych poszkodowanych (z grupy T2) przy użyciu 4 zwróconych ZRM. Do zadysponowania pozostaje 2 poszkodowanych — uczestnik decyduje, które ZRM przypisać (test: czy ponownie użyje ZRM S/LZRM do T2, czy racjonalnie przydzieli zespoły podstawowe).

**Przebieg wg metody:**
- **Sesja tradycyjna:** Uczestnik dopisuje ZRM i szpital w kolejnych wierszach Tabeli Dyslokacji.
- **Sesja z aplikacją:** Uczestnik najpierw klika „dostarczony" przy każdym z 4 ZRM (zmienia status z „w trasie" na „dostępny"), następnie przypisuje pozostałych poszkodowanych w module DYSPONOWANIE.

**Tabela odpowiedzi DM — jak w Turze 1**, plus:

| Pytanie uczestnika | Odpowiedź DM |
|---|---|
| „Czy te ZRM mogą jechać do innego szpitala niż poprzednio?" | „Tak, zespoły mogą być zadysponowane do dowolnego szpitala." |
| „Czy szpitale mają jeszcze wolne miejsca?" | „Pojemność szpitali podałem w informacji — proszę uwzględnić poszkodowanych już zadysponowanych w Turze 1." |
| „Jak oznaczyć dostarczony w aplikacji?" | „Proszę kliknąć przy każdym zespole przycisk oznaczający dostarczenie poszkodowanego do szpitala." |
| „Nie widzę wolnych ZRM w aplikacji" | „Proszę najpierw oznaczyć zespoły jako dostarczone — wtedy będą ponownie dostępne." |

**Trigger zakończenia Tury 2:** Uczestnik mówi wariant: „Dyslokacja zakończona" / „Wszyscy przypisani" / „Gotowe"

**→ STOP stopera [C2] → zapisz czas dyslokacji Tura 2**

---

#### PODFAZA D: Raport końcowy

Odczytaj:

> Kierujący Akcją Medyczną, proszę o sporządzenie raportu z przebiegu zdarzenia.

**→ START stopera [D]**

- Sesja tradycyjna: uczestnik pisze raport odręcznie na kartce A4
- Sesja z aplikacją: uczestnik generuje raport w module RAPORT, a następnie klika ZAKOŃCZ ZDARZENIE

**DM nie kontaktuje się z uczestnikiem w tej podfazie.**

Jeśli uczestnik pyta:

| Pytanie | Odpowiedź DM |
|---|---|
| „Co ma zawierać raport?" | „Raport powinien zawierać podsumowanie zdarzenia: liczbę poszkodowanych, podział na grupy, przypisanie zespołów i szpitali, oraz dane o przebiegu akcji. Proszę też oznaczyć zakończenie zdarzenia." |
| „Czy raport z aplikacji wystarczy?" | „Proszę wygenerować raport w aplikacji i zakończyć zdarzenie." |
| „Jak zakończyć zdarzenie?" | „Proszę kliknąć przycisk ZAKOŃCZ ZDARZENIE w sekcji RAPORT." |

**Trigger zakończenia:** Uczestnik mówi: „Raport gotowy" / „Skończyłem"

**→ STOP stopera [D] → zapisz czas raportu**

---

#### Zbieranie materiałów

1. Sesja tradycyjna: Zbierz Tabelę Dyslokacji + kartę z raportem + karty pacjentów
2. Sesja z aplikacją: Zrób zrzut ekranu / eksport raportu z aplikacji. Zbierz karty pacjentów.
3. Wpisz czasy na Arkusz czasów
4. Wypełnij Kartę ocen (obserwatora) dla tej sesji

---

### FAZA 2: Przerwa (15 minut)

> Dziękuję. Teraz przerwa — piętnaście minut. Proszę nie rozmawiać o szczegółach zadania z innymi uczestnikami.

*(W trakcie przerwy: przygotuj materiały do sesji 2 — nowy scenariusz, nowa metoda.)*

---

### FAZA 3: Sesja 2

Powtórz FAZĘ 1 (Podfazy A-D) z:
- Drugim scenariuszem
- Drugą metodą (jeśli sesja 1 była tradycyjna → sesja 2 z aplikacją, i odwrotnie)
- Identycznymi komunikatami DM

**Jedyna różnica:** Przed sesją 2 NIE powtarzaj instrukcji aplikacji ani wyjaśnienia roli DM. Powiedz tylko:

> Przechodzimy do drugiej sesji. Tym razem [metoda tradycyjna / z aplikacją]. Scenariusz jest inny, ale zasady takie same. Gotowy/Gotowa?

---

### FAZA 4: Kwestionariusz (po obu sesjach)

> Dziękuję za ukończenie obu sesji. Teraz proszę o wypełnienie krótkiego kwestionariusza — zajmie to około dziesięć minut.

Wręcz:
- Kwestionariusz SUS (10 pytań, dot. aplikacji TRIAGE MCI)
- Pytania o preferencje (którą metodę wybrałby/wybrałaby na realnym zdarzeniu i dlaczego)
- Dane demograficzne i zawodowe

---

## D. Biblioteka odpowiedzi DM — pełna tabela referencyjna

### D.1. Pytania proceduralne

| Pytanie | Odpowiedź |
|---|---|
| „Co mam zrobić najpierw?" | „Proszę rozpocząć od segregacji medycznej poszkodowanych." |
| „Ile mam czasu?" | „Nie ma limitu czasu. Proszę pracować tak szybko, jak to możliwe." |
| „Czy muszę zrobić raport?" | „Tak, raport jest ostatnim etapem zadania." |
| „Czy muszę przypisać wszystkich pacjentów?" | „Proszę przypisać do transportu poszkodowanych z grupy czerwonej i żółtej." |
| „W jakiej kolejności mam robić?" | „Proszę najpierw przeprowadzić segregację, potem przekazać mi raport, następnie przejść do dyslokacji, a na końcu sporządzić raport." |

### D.2. Pytania kliniczne

| Pytanie | Odpowiedź |
|---|---|
| Cokolwiek o stanie pacjenta | „Wszystkie potrzebne informacje są na karcie pacjenta." |
| „Czy ten pacjent chodzi?" | „Informacja ta znajduje się na karcie." |
| „Co się stanie po udrożnieniu?" | „Informacja ta znajduje się na karcie." |
| „Ten pacjent jest stabilny?" | „Proszę dokonać oceny na podstawie informacji z karty." |

### D.3. Pytania logistyczne

| Pytanie | Odpowiedź |
|---|---|
| „Czy mogą przyjechać dodatkowe ZRM?" | „Dysponuje Pan/Pani czterema zespołami. Nie ma możliwości zadysponowania dodatkowych w tym czasie." |
| „Czy szpital może przyjąć więcej?" | „Proszę opierać się na informacjach o pojemności szpitali, które przekazałem." |
| „Czy mogę wysłać dwóch jednym ZRM?" | „Jeden ZRM transportuje jednego poszkodowanego z grupy czerwonej. W uzasadnionych przypadkach możliwy jest transport dwóch z grupy żółtej jednym ZRM." |
| „Który ZRM jest najlepszy?" | „Decyzja o przypisaniu ZRM należy do Kierującego Akcją Medyczną." |
| „Gdzie wysłać tego pacjenta?" | „Decyzja o dyslokacji należy do Kierującego Akcją Medyczną." |
| „Jak daleko jest szpital?" | „Czasy dojazdu podałem w informacji o szpitalach — proszę sprawdzić na karcie." |
| „Kiedy wrócą ZRM?" | [Tura 1]: „Poinformuję, gdy zespoły będą dostępne." / [Tura 2]: „Wszystkie cztery zespoły są dostępne." |
| „Czy te ZRM mogą jechać do innego szpitala?" | „Tak, zespoły mogą być zadysponowane do dowolnego szpitala." |
| „Czy szpitale mają jeszcze wolne miejsca po Turze 1?" | „Proszę uwzględnić poszkodowanych już zadysponowanych. Pojemność szpitali podałem w informacji." |
| „Aplikacja pokazuje BRAK MIEJSC RED/YELLOW" | „To ostrzeżenie informacyjne. Decyzja o dyslokacji należy do Kierującego Akcją Medyczną." |
| „Jak wpisać płeć i wiek?" | „Przy zatwierdzaniu pacjenta może Pan/Pani opcjonalnie wybrać płeć i ustawić wiek." |
| „Jak zmienić kategorię pacjenta?" | „Proszę kliknąć odpowiednią kropkę koloru przy pacjencie." |
| „Jak zmienić szpital w transporcie?" | „Proszę kliknąć przycisk «zmień» obok nazwy szpitala." |
| „Jak zakończyć zdarzenie?" | „Proszę kliknąć przycisk ZAKOŃCZ ZDARZENIE w sekcji RAPORT." |
| „Co wpisać w polu KAM?" | „Proszę wpisać imię i nazwisko Kierującego Akcją Medyczną." |
| „Co to jest GDM?" | „Główny Dyspozytor Medyczny. Proszę wpisać moje dane: [podaj swoje imię i nazwisko jako prowadzącego]." |

### D.4. Odpowiedzi ZAKAZANE

**NIGDY nie mów:**
- Sugestii dot. kategorii segregacyjnej („Ten pacjent wygląda na T1")
- Sugestii dot. dyslokacji („Ja bym wysłał go do Uniwersyteckiego")
- Komentarzy wartościujących („Dobrze", „Nieprawidłowo", „Hmm...")
- Presji czasowej („Proszę się pospieszyć", „Inni zrobili to szybciej")
- Porównań z innymi uczestnikami
- Pochwał ani krytyki w trakcie sesji

---

## E. Trening prowadzącego

### Przed pierwszym uczestnikiem:

1. **Przeczytaj cały skrypt 2 razy** — zaznacz miejsca, gdzie musisz odczytać tekst dosłownie
2. **Przeprowadź 2 sesje próbne** — poproś kolegę/koleżankę o odegranie roli uczestnika
3. **Przećwicz obsługę stopera** — osobny start/stop na każdą podfazę (A, C1, C2, D); upewnij się, że potrafisz szybko resetować i uruchamiać stoper między fazami
4. **Wydrukuj materiały** — sprawdź, czy masz wystarczającą liczbę kopii na cały dzień badań
5. **Przygotuj stanowisko** — stół, krzesło, cisza, brak rozpraszaczy

### Częste pułapki:

| Pułapka | Jak unikać |
|---|---|
| Zapomnienie o wręczeniu Karty szpitali | Karta leży na osobnym stosie, wręczasz ją po Kroku B.4 |
| Korekta błędnej segregacji uczestnika | NIE koryguj. Powtórz dokładnie to, co uczestnik powiedział |
| Różne tempo mówienia w różnych sesjach | Czytaj z karty w jednakowym tempie |
| Zapomnienie o uruchomieniu/zatrzymaniu stopera | Wydrukuj ściągawkę: A→start, „skończyłem segregację"→stop A, po B.6→start C1, „brak ZRM"→stop C1, po komunikacie→start C2, „gotowe"→stop C2, po poleceniu raportu→start D, „raport gotowy"→stop D |
| Pominięcie kroku B.5 (import SMS) | Sesja z aplikacją: po podaniu danych o szpitalach wyślij SMS z linkiem. POCZEKAJ, aż uczestnik tapnie link i dane się zaimportują |
| Rozmowa towarzyska w trakcie przerwy | Ograniczaj kontakt w przerwie do minimum |

---

## F. Karta informacyjna szpitali (do wydrukowania i wręczenia uczestnikowi)

*Wytnij i wręcz uczestnikowi w Podfazie B, po przekazaniu informacji ustnie.*

```
┌──────────────────────────────────────────────────────────────────────┐
│                    INFORMACJA O SZPITALACH                          │
│              (Dyspozytor Medyczny → Kierujący Akcją Medyczną)       │
├──────────────────────────────┬──────────┬──────┬──────┬──────┬──────┤
│ Szpital                      │ Oddziały │  OR  │ ITM  │ CZER │ ŻÓŁ │
├──────────────────────────────┼──────────┼──────┼──────┼──────┼──────┤
│ Szpital Uniwersytecki        │ SOR, CU  │  2   │  3   │ do 3 │ do 3│
│ Czas dojazdu: 12 min        │ chir.    │      │      │      │     │
│                              │ ortop.   │      │      │      │     │
├──────────────────────────────┼──────────┼──────┼──────┼──────┼──────┤
│ Szpital Miejski nr 2        │ SOR      │  1   │  2   │ do 2 │ do 2│
│ Czas dojazdu: 18 min        │ chir. og.│      │      │      │     │
├──────────────────────────────┼──────────┼──────┼──────┼──────┼──────┤
│ Szpital Powiatowy            │ IP       │  0   │  1   │  0   │ do 3│
│ Czas dojazdu: 25 min        │          │      │      │      │     │
├──────────────────────────────┴──────────┴──────┴──────┴──────┴──────┤
│ OR = bloki operacyjne | ITM = stanowiska intensywnej terapii       │
│ CZER = przyjęcie grupy czerwonej | ŻÓŁ = przyjęcie grupy żółtej   │
└──────────────────────────────────────────────────────────────────────┘
```

---

## G. Arkusz czasów — szablon

| ID | Sesja | Metoda | Scenariusz | t_A (segregacja) | t_C1 (dyslokacja T1) | t_C2 (dyslokacja T2) | t_D (raport) | t_suma | Uwagi |
|---|---|---|---|---|---|---|---|---|---|
| U__ | 1 |  |  | | | | | | |
| U__ | 2 |  |  | | | | | | |

**Pomiar:** Każda faza mierzona osobnym stoperem (start/stop). Wymiany DM (Podfaza B w tym B.5 — import danych SMS, komunikat o powrocie ZRM) NIE są mierzone — to stałe, identyczne dla każdego uczestnika. W sesji z aplikacją dodatkowo działa logger eksperymentalny (tryb REC) rejestrujący akcje z precyzją milisekundową — dane eksportowane po sesji jako JSON (patrz sekcja J).

**Obliczenia:**
- t_dyslokacja = t_C1 + t_C2
- t_suma = t_A + t_C1 + t_C2 + t_D

Wartość t_suma jest bezpośrednio porównywalna między metodami, bo nie zawiera czasu wymian DM.

---

## H. Kolejność sesji — lista randomizacyjna

Przed badaniem wypełnij tabelę randomizacyjną:

| ID uczestnika | Sekwencja | Sesja 1: metoda | Sesja 1: scenariusz | Sesja 2: metoda | Sesja 2: scenariusz |
|---|---|---|---|---|---|
| U01 | AB | Tradycyjna | S1 | Aplikacja | S2 |
| U02 | BA | Aplikacja | S1 | Tradycyjna | S2 |
| U03 | AB | Tradycyjna | S1 | Aplikacja | S2 |
| U04 | BA | Aplikacja | S1 | Tradycyjna | S2 |
| ... | ... | ... | ... | ... | ... |

Losowanie: użyj generatora liczb losowych lub metody kopertowej. Przygotuj kompletną listę PRZED pierwszym uczestnikiem.

---

## I. Checklist zakończenia (po każdym uczestniku)

- [ ] Oba zestawy kart pacjentów odebrane i odłożone
- [ ] Tabela Dyslokacji (papierowa) zebrana i oznaczona kodem uczestnika
- [ ] Raport papierowy zebrany i oznaczony kodem uczestnika
- [ ] **EKSPORT JSON z loggera eksperymentalnego:** tapnij REC (prawy górny róg) → Export JSON → zapisz plik `exp_[ID]_[scenariusz]_[data].json`
- [ ] **STOP nagrywania:** w panelu eksportu kliknij „Stop Recording"
- [ ] Arkusz czasów wypełniony dla obu sesji
- [ ] Karta ocen wypełniona dla obu sesji
- [ ] Kwestionariusz SUS + preferencje + dane demograficzne zebrane
- [ ] Aplikacja zresetowana (nowe zdarzenie, puste listy) dla następnego uczestnika
- [ ] Nowy komplet kart pacjentów przygotowany i potasowany

## J. Aktywacja trybu eksperymentalnego (przed sesją z aplikacją)

Tryb eksperymentalny rejestruje każdą akcję uczestnika z precyzją milisekundową. Aktywuj go **przed** rozpoczęciem sesji:

1. Otwórz aplikację na urządzeniu uczestnika (`experiment.html` — upewnij się, że to wersja eksperymentalna, nie produkcyjna)
2. Otwórz panel Info (przycisk Info na ekranie startowym)
3. Tapnij ikonę 🔬 w prawym górnym rogu info-boxa — otworzy się panel Experiment Mode
4. Wpisz **Participant ID** (kod uczestnika, np. U01) i **Scenario** (np. S1 lub S2)
5. Kliknij **START REC** — w prawym górnym rogu ekranu pojawi się migający czerwony wskaźnik REC
6. Zamknij panel Info — uczestnik może rozpocząć sesję

**UWAGA:** Nie zamykaj ani nie odświeżaj aplikacji w trakcie sesji — log jest przechowywany w pamięci.

Po zakończeniu sesji: eksportuj JSON (patrz Checklist zakończenia, sekcja I).

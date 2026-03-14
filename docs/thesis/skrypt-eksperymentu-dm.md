# Skrypt eksperymentu — Standardowy protokół prowadzącego (rola DM/GDM)

**Wersja:** 1.0
**Data:** 2026-03-14
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
- **Timer:** Używaj stopera z międzyczasami (split/lap). Zapisuj każdy timestamp na Arkuszu czasów.

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
- [ ] Smartfon z TRIAGE MCI (nowe zdarzenie, pusta lista pacjentów)
- [ ] Zespoły ZRM i szpitale wstępnie skonfigurowane w aplikacji
- [ ] ~~Karta informacyjna szpitali~~ (podasz po segregacji, nie wcześniej!)

### Materiały prowadzącego

- [ ] Ten skrypt (wydrukowany)
- [ ] Stoper z funkcją międzyczasów
- [ ] Arkusz czasów (jeden wiersz na sesję)
- [ ] Klucz referencyjny — segregacja (z scenariusze-eksperyment.md)
- [ ] Klucz referencyjny — dyslokacja (z scenariusze-eksperyment.md)
- [ ] Karta ocen (formularz obserwatora, jeden na sesję)
- [ ] Lista kontrolna kompletności raportu (Raport GDM)
- [ ] Karta informacyjna szpitali (do wręczenia uczestnikowi w Podfazie B)

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
1. Ekran główny → ROZPOCZNIJ TRIAGE
2. Kreator START — pytania TAK/NIE → wynik segregacji
3. Moduł DYSPONOWANIE — przypisanie ZRM i szpitala
4. Moduł RAPORT — generowanie raportu

**Nie komentuj zalet ani wad aplikacji. Pokaż, jak działają przyciski.**

#### Krok 0.3 — Wyjaśnienie roli DM

Odczytaj:

> W trakcie każdej sesji będę pełnił rolę Dyspozytora Medycznego. Proszę traktować naszą komunikację jak łączność radiową. Po zakończeniu segregacji proszę przekazać mi raport z rozpoznania wstępnego — ile poszkodowanych w poszczególnych grupach. Ja przekażę Panu/Pani informacje o zadysponowanych zespołach ZRM i możliwościach szpitali. Następnie wspólnie ustalimy dyslokację poszkodowanych. Na końcu poproszę o sporządzenie raportu ze zdarzenia.
>
> Proszę pamiętać: w trakcie segregacji pracuje Pan/Pani samodzielnie. Kontakt ze mną następuje po jej zakończeniu.

#### Krok 0.4 — Wyjaśnienie pomiaru

Odczytaj:

> Będę mierzyć czas od momentu otwarcia koperty z kartami pacjentów. Nie ma limitu czasu, ale proszę pracować tak szybko, jak to możliwe przy zachowaniu dokładności. Czas jest mierzony w kilku etapach, które będę oznaczać.
>
> Czy ma Pan/Pani pytania zanim rozpoczniemy?

*(Odpowiedz na pytania organizacyjne. NIE odpowiadaj na pytania merytoryczne dotyczące START ani procedury.)*

---

### FAZA 1: Sesja (timer działa)

---

#### PODFAZA A: Segregacja pierwotna

**Odczytaj opis scenariusza** (z scenariusze-eksperyment.md — odpowiedni dla danej sesji).

Następnie odczytaj:

> Pełni Pan/Pani funkcję Kierującego Akcją Medyczną. Na miejscu zdarzenia znajduje się 12 poszkodowanych wymagających segregacji medycznej algorytmem START.
>
> [SESJA TRADYCYJNA]: Proszę otworzyć kopertę. Na podstawie kart pacjentów proszę przeprowadzić segregację i odnotować wyniki na Tabeli Dyslokacji — kolumny: numer, kategoria, płeć/wiek, obrażenia. Po zakończeniu segregacji proszę o przekazanie mi raportu z rozpoznania wstępnego.
>
> [SESJA Z APLIKACJĄ]: Proszę otworzyć kopertę. Na podstawie kart pacjentów proszę przeprowadzić segregację w aplikacji TRIAGE MCI, przechodząc kreator START dla każdego pacjenta. Po zakończeniu segregacji proszę o przekazanie mi raportu z rozpoznania wstępnego.
>
> Zaczynamy.

**→ URUCHOM TIMER (T_start)**

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

**Trigger zakończenia Podfazy A:** Uczestnik mówi wariant: „Skończyłem segregację" / „Mam raport z rozpoznania" / „Gotowy do raportu"

**→ ZAPISZ MIĘDZYCZAS (T_segregacja)**

---

#### PODFAZA B: Raport z rozpoznania wstępnego + przekazanie danych

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

*(Wręcz uczestnikowi wydrukowaną Kartę informacyjną szpitali.)*

**Krok B.5 — Polecenie dyslokacji**

Odczytaj:

> Proszę o przeprowadzenie dyslokacji poszkodowanych z grupy czerwonej i żółtej. Proszę przypisać poszkodowanych do zespołów ZRM i wskazać szpitale docelowe.

**→ ZAPISZ MIĘDZYCZAS (T_raport_DM)**

---

#### PODFAZA C: Dyslokacja

**Uczestnik pracuje:**
- Sesja tradycyjna: wypełnia kolumny ZRM i Szpital na Tabeli Dyslokacji
- Sesja z aplikacją: używa modułu DYSPONOWANIE

**Prowadzący (DM) jest dostępny.** Odpowiadaj na pytania wg poniższej tabeli:

| Pytanie uczestnika | Odpowiedź DM |
|---|---|
| „Czy mogę wysłać dwóch pacjentów jednym ZRM?" | „Jeden ZRM transportuje jednego poszkodowanego z grupy czerwonej. W uzasadnionych przypadkach możliwy jest transport dwóch poszkodowanych z grupy żółtej jednym ZRM." |
| „Co z pacjentami z grupy zielonej?" | „Poszkodowani z grupy zielonej mogą być transportowani innym środkiem transportu. Proszę skupić się na dyslokacji poszkodowanych z grupy czerwonej i żółtej." |
| „Co z pacjentami z grupy czarnej?" | „Poszkodowani z grupy czarnej pozostają na miejscu zdarzenia. Nie wymagają dysponowania ZRM." |
| „Czy mogą przyjechać dodatkowe ZRM?" | „Dysponuje Pan/Pani czterema zespołami. Nie ma możliwości zadysponowania dodatkowych zespołów w tym czasie." |
| „Czy szpital może przyjąć więcej pacjentów?" | „Proszę opierać się na informacjach o pojemności szpitali, które przekazałem." |
| „Który szpital jest najlepszy dla tego pacjenta?" | „Decyzja o dyslokacji należy do Kierującego Akcją Medyczną." |
| Uczestnik przekazuje plan dyslokacji | „Przyjmuję. [Powtórz przypisanie]. Potwierdzam dyslokację." |

**Trigger zakończenia Podfazy C:** Uczestnik mówi wariant: „Dyslokacja zakończona" / „Wszyscy przypisani" / „Gotowe"

**→ ZAPISZ MIĘDZYCZAS (T_dyslokacja)**

---

#### PODFAZA D: Raport końcowy

Odczytaj:

> Kierujący Akcją Medyczną, proszę o sporządzenie raportu z przebiegu zdarzenia.

- Sesja tradycyjna: uczestnik pisze raport odręcznie na kartce A4
- Sesja z aplikacją: uczestnik generuje raport w module RAPORT

**DM nie kontaktuje się z uczestnikiem w tej podfazie.**

Jeśli uczestnik pyta:

| Pytanie | Odpowiedź DM |
|---|---|
| „Co ma zawierać raport?" | „Raport powinien zawierać podsumowanie zdarzenia: liczbę poszkodowanych, podział na grupy, przypisanie zespołów i szpitali, oraz dane o przebiegu akcji." |
| „Czy raport z aplikacji wystarczy?" | „Proszę wygenerować raport w aplikacji." |

**Trigger zakończenia Podfazy D:** Uczestnik mówi: „Raport gotowy" / „Skończyłem"

**→ ZATRZYMAJ TIMER (T_raport_końcowy = T_total)**

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
3. **Przećwicz obsługę stopera** — upewnij się, że potrafisz zapisywać międzyczasy bez przerywania sesji
4. **Wydrukuj materiały** — sprawdź, czy masz wystarczającą liczbę kopii na cały dzień badań
5. **Przygotuj stanowisko** — stół, krzesło, cisza, brak rozpraszaczy

### Częste pułapki:

| Pułapka | Jak unikać |
|---|---|
| Zapomnienie o wręczeniu Karty szpitali | Karta leży na osobnym stosie, wręczasz ją po Kroku B.4 |
| Korekta błędnej segregacji uczestnika | NIE koryguj. Powtórz dokładnie to, co uczestnik powiedział |
| Różne tempo mówienia w różnych sesjach | Czytaj z karty w jednakowym tempie |
| Zapomnienie o międzyczasie | Timer z dźwiękowym potwierdzeniem (beep) pomaga |
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

| ID | Sesja | Metoda | Scenariusz | T_start | T_segregacja | T_raport_DM | T_dyslokacja | T_raport_końcowy | Czas_łączny | Uwagi |
|---|---|---|---|---|---|---|---|---|---|---|
| U__ | 1 |  |  | : | : | : | : | : | | |
| U__ | 2 |  |  | : | : | : | : | : | | |

**Obliczenia:**
- Czas segregacji = T_segregacja − T_start
- Czas wymiany DM = T_raport_DM − T_segregacja
- Czas dyslokacji = T_dyslokacja − T_raport_DM
- Czas raportu = T_raport_końcowy − T_dyslokacja
- Czas łączny = T_raport_końcowy − T_start

**Uwaga:** Czas wymiany DM jest stałą (identyczne komunikaty), więc do porównania metod używaj: Czas łączny − Czas wymiany DM.

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
- [ ] Zrzut ekranu / eksport z aplikacji zapisany
- [ ] Arkusz czasów wypełniony dla obu sesji
- [ ] Karta ocen wypełniona dla obu sesji
- [ ] Kwestionariusz SUS + preferencje + dane demograficzne zebrane
- [ ] Aplikacja zresetowana (nowe zdarzenie) dla następnego uczestnika
- [ ] Nowy komplet kart pacjentów przygotowany i potasowany

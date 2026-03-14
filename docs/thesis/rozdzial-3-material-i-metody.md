# Rozdział 3. Materiał i metody

## 3.1. Cel i uzasadnienie badania

Celem niniejszej pracy było porównanie efektywności cyfrowego systemu wspomagania segregacji medycznej i dokumentacji zdarzenia z dużą liczbą poszkodowanych (aplikacja TRIAGE MCI) z obowiązującą metodą papierową w warunkach symulowanego zdarzenia masowego. Badanie miało charakter dwuczęściowy: obejmowało eksperyment symulacyjny z zastosowaniem schematu crossover oraz ogólnopolską ankietę skierowaną do ratowników medycznych.

Uzasadnienie podjęcia badania wynika z przesłanek zidentyfikowanych w przeglądzie literatury (rozdział 2):

1. Algorytm START, choć ugruntowany w praktyce, charakteryzuje się znaczącymi wskaźnikami overtriage (30-50%) i undertriage (5-15%), których jednym ze źródeł jest błąd ludzki wynikający ze stresu, zmęczenia i niedostatecznego doświadczenia [Jenkins i in., 2008].
2. Obowiązująca *Procedura postępowania na wypadek wystąpienia zdarzenia z dużą liczbą poszkodowanych* (MZ v2.3, 2024) definiuje obowiązkowe narzędzia dokumentacyjne — Tabelę Dyslokacji Poszkodowanych, Tabelę Szpitali i Raport GDM — które funkcjonują wyłącznie w formie papierowej, co stanowi potencjalne ograniczenie operacyjne.
3. Cyfryzacja procesu segregacji i dokumentacji ma potencjał do standaryzacji, automatyzacji i usprawnienia raportowania, lecz brakuje badań porównawczych oceniających ten potencjał w kontekście polskiego systemu ratownictwa medycznego i obowiązujących regulacji.
4. Istniejące rozwiązania cyfrowe nie spełniają łącznie wszystkich wymagań praktyki polowej (kompletność funkcjonalna, praca offline, dostępność w języku polskim, bezpłatność) ani nie realizują funkcjonalności odpowiadającej narzędziom z Procedury MZ v2.3.

## 3.2. Pytania badawcze i hipotezy

### 3.2.1. Pytanie badawcze w formacie PICOT

| Element | Definicja |
|---------|-----------|
| **P** (Population) | Ratownicy medyczni i studenci ratownictwa medycznego (n=40-50) |
| **I** (Intervention) | Segregacja z użyciem aplikacji TRIAGE MCI |
| **C** (Comparison) | Segregacja metodą tradycyjną bez wsparcia cyfrowego |
| **O** (Outcome) | Czas segregacji, trafność klasyfikacji (T1-T4), kompletność raportu, subiektywna ocena użyteczności |
| **T** (Time) | Jednorazowa symulacja z crossover (dwa scenariusze) |

### 3.2.2. Szczegółowe pytania badawcze

Na podstawie powyższego sformułowano następujące pytania badawcze:

**PB1:** Czy zastosowanie aplikacji TRIAGE MCI wpływa na czas potrzebny do przeprowadzenia segregacji medycznej grupy poszkodowanych w porównaniu z metodą tradycyjną?

**PB2:** Czy trafność klasyfikacji segregacyjnej (zgodność z kluczem referencyjnym) różni się między metodą cyfrową a metodą tradycyjną?

**PB3:** Czy kompletność raportu ze zdarzenia jest wyższa przy zastosowaniu aplikacji niż przy metodzie tradycyjnej?

**PB4:** Jak ratownicy medyczni i studenci ratownictwa oceniają użyteczność aplikacji TRIAGE MCI w porównaniu z metodą tradycyjną?

**PB5:** Jaki jest stosunek polskich ratowników medycznych do cyfrowych narzędzi wspomagających segregację i jakie bariery wdrożeniowe identyfikują?

### 3.2.3. Hipotezy badawcze

Na podstawie przesłanek z przeglądu literatury sformułowano następujące hipotezy:

**H1 (główna):** Średni czas segregacji grupy poszkodowanych jest istotnie statystycznie krótszy przy zastosowaniu aplikacji TRIAGE MCI niż przy metodzie tradycyjnej.

**H2:** Trafność klasyfikacji segregacyjnej (odsetek prawidłowo przypisanych kategorii T1-T4) jest istotnie wyższa przy zastosowaniu aplikacji niż przy metodzie tradycyjnej.

**H3:** Kompletność raportu ze zdarzenia (mierzona standaryzowaną listą kontrolną) jest istotnie wyższa przy zastosowaniu aplikacji niż przy metodzie tradycyjnej.

**H4:** Uczestnicy badania oceniają aplikację TRIAGE MCI jako narzędzie o wysokiej użyteczności (wynik powyżej 68 punktów w skali SUS).

**Hipoteza zerowa (H0)** dla hipotez H1-H3: Nie istnieje statystycznie istotna różnica między metodą cyfrową a tradycyjną w zakresie badanej zmiennej.

### 3.2.4. Ocena pytania badawczego według kryteriów FINER

| Kryterium | Ocena |
|-----------|-------|
| **F** (*Feasible*) | Autor ma bezpośredni dostęp do grupy 40-50 uczestników (współstudenci kierunku ratownictwo medyczne oraz współpracownicy — czynni ratownicy medyczni). Aplikacja TRIAGE MCI jest gotowa i przetestowana. Scenariusze symulacyjne są możliwe do przygotowania przy minimalnych nakładach materialnych. |
| **I** (*Interesting*) | Cyfryzacja ratownictwa medycznego jest aktualnym i dynamicznie rozwijającym się obszarem, budzącym zainteresowanie zarówno środowiska akademickiego, jak i praktyków. |
| **N** (*Novel*) | W polskiej literaturze brak badań porównawczych oceniających efektywność aplikacji mobilnych do segregacji medycznej w zdarzeniach masowych. Badanie wypełnia zidentyfikowaną lukę badawczą. |
| **E** (*Ethical*) | Badanie przeprowadzane jest w warunkach symulacyjnych, bez udziału realnych pacjentów. Udział jest dobrowolny, a dane anonimowe. Nie występują zagrożenia etyczne. |
| **R** (*Relevant*) | Wyniki badania mogą mieć bezpośrednie przełożenie na praktykę szkoleniową i operacyjną w polskim systemie ratownictwa medycznego. |

## 3.3. Schemat badania

Badanie zaprojektowano jako dwuczęściowe, z komplementarnymi źródłami danych:

**Część I — Eksperyment symulacyjny** (dane pierwotne, badanie dynamiczne)
Badanie eksperymentalne z zastosowaniem schematu crossover (krzyżowego), w którym każdy uczestnik pełnił rolę zarówno grupy eksperymentalnej, jak i kontrolnej. Schemat ten pozwolił na eliminację zmienności międzyosobniczej i uzyskanie istotnych statystycznie wyników przy relatywnie niewielkiej próbie.

**Część II — Ankieta ogólnopolska** (dane pierwotne, badanie przekrojowe)
Anonimowa ankieta internetowa skierowana do ratowników medycznych w Polsce, mająca na celu zbadanie postaw, doświadczeń i barier związanych z cyfryzacją segregacji medycznej.

Poniższy schemat ilustruje ogólną strukturę badania:

```
BADANIE
├── Część I: Eksperyment crossover (n=40-50)
│   ├── Grupa A: Scenariusz 1 (tradycyjnie) → Scenariusz 2 (aplikacja)
│   ├── Grupa B: Scenariusz 1 (aplikacja) → Scenariusz 2 (tradycyjnie)
│   ├── Pomiar: czas, trafność, kompletność raportu
│   └── Kwestionariusz SUS + preferencje
│
└── Część II: Ankieta ogólnopolska (cel: n=100-200+)
    ├── Doświadczenie z MCI
    ├── Znajomość algorytmu START
    ├── Stosunek do narzędzi cyfrowych
    └── Gotowość do adopcji i bariery
```

## 3.4. Część I — Eksperyment symulacyjny

### 3.4.1. Grupa badana

Do udziału w eksperymencie zaproszono dwie grupy uczestników:

**Grupa A — studenci ratownictwa medycznego** (n=20-30): Studenci drugiego stopnia (studia magisterskie) kierunku ratownictwo medyczne, posiadający co najmniej tytuł licencjata w zakresie ratownictwa medycznego oraz ukończone szkolenie z zakresu segregacji medycznej w zdarzeniach masowych.

**Grupa B — czynni ratownicy medyczni** (n≈20): Ratownicy medyczni zatrudnieni w systemie Państwowego Ratownictwa Medycznego, posiadający co najmniej roczne doświadczenie zawodowe.

**Kryteria włączenia:**
- ukończone szkolenie z zakresu algorytmu START;
- posiadanie smartfona z systemem Android lub iOS;
- wyrażenie świadomej zgody na udział w badaniu.

**Kryteria wyłączenia:**
- wcześniejsze korzystanie z aplikacji TRIAGE MCI (aby uniknąć efektu wprawy);
- odmowa udziału lub wycofanie zgody na dowolnym etapie badania.

### 3.4.2. Randomizacja

Uczestnicy zostali losowo przydzieleni do dwóch sekwencji eksperymentalnych za pomocą prostej randomizacji (metoda kopertowa lub generator liczb losowych):

- **Sekwencja AB:** Scenariusz 1 metodą tradycyjną → przerwa 15 min → Scenariusz 2 z aplikacją TRIAGE MCI
- **Sekwencja BA:** Scenariusz 1 z aplikacją TRIAGE MCI → przerwa 15 min → Scenariusz 2 metodą tradycyjną

Randomizacja kolejności miała na celu kontrolowanie efektu uczenia się (*carry-over effect*) — gdyby wszyscy uczestnicy zaczynali od tej samej metody, wynik drugiego scenariusza mógłby być zawyżony ze względu na nabyte doświadczenie.

### 3.4.3. Scenariusze symulacyjne

Przygotowano dwa scenariusze zdarzenia masowego o zbliżonej trudności i strukturze. Każdy scenariusz obejmował 12 poszkodowanych z następującym rozkładem kategorii segregacyjnych:

| Kategoria | Scenariusz 1 | Scenariusz 2 |
|-----------|-------------|-------------|
| T1 (czerwony) | 3 | 3 |
| T2 (żółty) | 3 | 3 |
| T3 (zielony) | 4 | 4 |
| T4 (czarny) | 2 | 2 |
| **Razem** | **12** | **12** |

**Scenariusz 1: Wypadek autobusu na drodze krajowej**
Autobus miejski zderzył się z samochodem ciężarowym. Na miejscu 12 poszkodowanych o zróżnicowanym stopniu obrażeń. Każdy poszkodowany został opisany na indywidualnej karcie zawierającej: płeć, przybliżony wiek, opis obrażeń widocznych, stan świadomości, obecność oddechu, zdolność chodzenia oraz dodatkowe informacje kliniczne (np. masywne krwawienie z kończyny, siniaki na klatce piersiowej, widoczna deformacja kończyny).

**Scenariusz 2: Wybuch gazu w budynku wielorodzinnym**
Eksplozja gazu w czteropiętrowym bloku mieszkalnym. Na miejscu 12 poszkodowanych. Analogiczna struktura kart pacjentów jak w scenariuszu 1, z odmiennym mechanizmem urazowym (oparzenia, urazy od odłamków, urazy inhalacyjne, przysypanie gruzem).

Scenariusze zaprojektowano tak, aby:
- zawierały porównywalny rozkład kategorii segregacyjnych;
- każdy krok algorytmu START był testowany co najmniej dwukrotnie;
- obejmowały przypadki jednoznaczne i niejednoznaczne (graniczne między kategoriami);
- były realistyczne i klinicznie wiarygodne;
- różniły się kontekstem i mechanizmem urazu, aby zminimalizować efekt pamięciowy.

Prawidłowość kategoryzacji każdego przypadku (klucz referencyjny) została ustalona niezależnie przez dwóch doświadczonych ratowników medycznych. W przypadku rozbieżności przeprowadzono dyskusję i osiągnięto konsensus.

### 3.4.4. Przebieg eksperymentu

Eksperyment przeprowadzono według następującego protokołu:

**Faza 0: Przygotowanie (przed dniem eksperymentu)**
- Instalacja aplikacji TRIAGE MCI na smartfonach uczestników z sekwencji BA (lub przygotowanie urządzeń zastępczych);
- Przygotowanie kart pacjentów, formularzy papierowych, zegarów, identyfikatorów;
- Randomizacja uczestników do sekwencji AB lub BA.

**Faza 1: Wprowadzenie (15 minut)**
- Wyjaśnienie celu badania i uzyskanie pisemnej zgody uczestników;
- Krótka instrukcja dotycząca aplikacji TRIAGE MCI (5 minut) — prezentacja interfejsu i podstawowych funkcji. Instrukcja była standaryzowana i jednakowa dla wszystkich uczestników;
- Prowadzący informował uczestników, że w trakcie sesji będzie pełnił rolę Dyspozytora Medycznego (DM) / Głównego Dyspozytora Medycznego (GDM) i krótko opisywał protokół komunikacji KAM–DM (kiedy uczestnik powinien się spodziewać kontaktu ze strony DM i kiedy sam powinien inicjować komunikację);
- Wyjaśnienie zasad oceny i dokumentacji wyników.

**Faza 2: Sesja 1 (bez limitu czasu)**
Każdy uczestnik indywidualnie realizował uproszczony zakres zadań Kierującego Akcją Medyczną (KAM) zgodnie z Procedurą MZ v2.3, w interakcji z prowadzącym pełniącym rolę Dyspozytora Medycznego (DM). Sesja składała się z czterech podfaz:

**Podfaza A: Segregacja pierwotna.** Uczestnik samodzielnie przeprowadzał segregację 12 poszkodowanych (na kartach) algorytmem START. Prowadzący (DM) nie inicjował kontaktu w tej podfazie. Czas mierzony od momentu otwarcia koperty z kartami.

**Podfaza B: Raport z rozpoznania wstępnego.** Po zakończeniu segregacji uczestnik przekazywał prowadzącemu (DM) ustny raport z rozpoznania wstępnego: liczba poszkodowanych w poszczególnych grupach segregacyjnych. Prowadzący potwierdzał odbiór raportu, przekazywał informacje o zadysponowanych ZRM oraz szczegółowe dane o możliwościach szpitali (wg standaryzowanej karty informacyjnej, identycznej dla wszystkich uczestników i obu metod).

**Podfaza C: Dyslokacja.** Uczestnik przypisywał poszkodowanych z grupy czerwonej i żółtej do zespołów ZRM oraz wskazywał szpitale docelowe, z uwzględnieniem informacji o pojemności szpitali otrzymanych od DM. Metodą tradycyjną: wypełnianie Tabeli Dyslokacji. Z aplikacją: moduł DYSPONOWANIE. Prowadzący (DM) był dostępny do udzielania odpowiedzi na pytania logistyczne (wg standaryzowanej biblioteki odpowiedzi).

**Podfaza D: Raport końcowy.** Na polecenie prowadzącego (DM) uczestnik sporządzał raport ze zdarzenia. Metodą tradycyjną: na kartce. Z aplikacją: za pomocą modułu RAPORT.

Metoda realizacji zależała od przypisanej sekwencji: tradycyjnie (papierowa Tabela Dyslokacji + ręczny raport na kartce) lub z aplikacją TRIAGE MCI.

Czas sesji mierzono w podziale na podfazy (A, B, C, D) za pomocą timera z funkcją międzyczasów, co umożliwiało analizę, w której fazie metoda cyfrowa lub papierowa wykazuje przewagę.

**Faza 3: Przerwa (15 minut)**
Czas na odpoczynek i neutralizację efektu zmęczenia.

**Faza 4: Sesja 2 (bez limitu czasu)**
Analogicznie do sesji 1, ale z drugim scenariuszem i drugą metodą (zgodnie z przypisaną sekwencją).

**Faza 5: Kwestionariusz (10 minut)**
Po zakończeniu obu sesji każdy uczestnik wypełniał kwestionariusz obejmujący:
- skalę SUS (System Usability Scale) dla aplikacji TRIAGE MCI;
- pytania o preferencje (którą metodę wybrałby na realnym zdarzeniu i dlaczego);
- dane demograficzne i zawodowe.

### 3.4.5. Zmienne badawcze

**Zmienna niezależna:**
Metoda segregacji (dwa poziomy: tradycyjna vs. aplikacja TRIAGE MCI).

**Zmienne zależne:**

| Zmienna | Typ | Jednostka/Skala | Sposób pomiaru |
|---------|-----|-----------------|----------------|
| Czas segregacji | Ilościowa ciągła | Sekundy | Stoper — od rozpoczęcia do zakończenia |
| Trafność klasyfikacji | Ilościowa dyskretna | Odsetek (0-100%) | Zgodność z kluczem: liczba prawidłowych / 12 × 100% |
| Kompletność raportu | Ilościowa dyskretna | Punkty (0-10) | Standaryzowana lista kontrolna |
| Ocena użyteczności (SUS) | Ilościowa ciągła | Punkty (0-100) | Kwestionariusz SUS |
| Preferencja metody | Jakościowa nominalna | Tradycyjna / Cyfrowa | Pytanie zamknięte |
| Czas segregacji (podfaza A) | Ilościowa ciągła | Sekundy | Timer — od otwarcia koperty do deklaracji zakończenia segregacji |
| Czas dyslokacji (podfaza C) | Ilościowa ciągła | Sekundy | Timer — od otrzymania danych o szpitalach do deklaracji zakończenia |
| Czas raportu (podfaza D) | Ilościowa ciągła | Sekundy | Timer — od polecenia DM do deklaracji ukończenia raportu |
| Jakość dyslokacji | Ilościowa dyskretna | Punkty (0-12) | Kompletność (0-6 pkt) + zgodność z pojemnością szpitali (0-6 pkt) |

**Zmienne kontrolowane:**
- Kolejność metod (randomizacja sekwencji AB/BA);
- Trudność scenariuszy (identyczny rozkład kategorii, walidacja przez ekspertów);
- Czas instruktażu (standaryzowany, 5 minut);
- Warunki przeprowadzenia (to samo pomieszczenie, te same warunki dla wszystkich uczestników).

**Zmienne zakłócające potencjalne:**
- Doświadczenie zawodowe (staż pracy, liczba zdarzeń masowych);
- Biegłość w obsłudze smartfona;
- Poziom znajomości algorytmu START przed badaniem.

Zmienne zakłócające kontrolowano poprzez: (a) schemat crossover (każdy uczestnik jest swoją własną kontrolą), (b) rejestrację zmiennych demograficznych i zawodowych umożliwiającą analizę podgrup, (c) randomizację kolejności.

### 3.4.6. Lista kontrolna kompletności raportu

Kompletność raportu ze zdarzenia oceniano za pomocą standaryzowanej listy kontrolnej opartej na wymaganiach **Raportu GDM z przebiegu zdarzenia** zdefiniowanego w *Procedurze postępowania na wypadek wystąpienia zdarzenia z dużą liczbą poszkodowanych* [Ministerstwo Zdrowia, 2024, zał. nr 14]. Lista obejmuje 10 elementów odpowiadających obowiązkowym składnikom raportu, z których każdy oceniano binarnie (0 — brak, 1 — obecny):

1. Oznaczenie dyspozytorni medycznej / identyfikacja zdarzenia
2. Data i godzina rozpoczęcia akcji ratowniczej
3. Data i godzina zakończenia akcji ratowniczej
4. Dane kierującego akcją medyczną (KAM)
5. Liczba, rodzaj i oznaczenie ZRM biorących udział w akcji
6. Zestawienie liczby poszkodowanych z podziałem na grupy segregacyjne (czerwona, żółta, zielona, czarna)
7. Zestawienie danych poszkodowanych (numer, kategoria, obrażenia)
8. Przypisanie ZRM do poszkodowanych (dyslokacja)
9. Zestawienie szpitali docelowych z przypisanymi poszkodowanymi
10. Dane osoby sporządzającej raport / czas sporządzenia

Maksymalny wynik: 10 punktów. Lista została opracowana bezpośrednio na podstawie wymagań Raportu GDM [zał. nr 14] z Procedury MZ v2.3, co zapewnia zgodność mierzonej kompletności z obowiązującym standardem regulacyjnym. Zastosowanie oficjalnego wzorca zamiast autorskiej listy kontrolnej wzmacnia trafność treściową pomiaru i umożliwia interpretację wyników w kategoriach zgodności z obowiązującą procedurą.

### 3.4.7. Kwestionariusz SUS (System Usability Scale)

Do oceny użyteczności aplikacji TRIAGE MCI zastosowano kwestionariusz SUS (*System Usability Scale*) opracowany przez Brooke'a [1996]. SUS jest jednym z najszerzej stosowanych standaryzowanych narzędzi oceny użyteczności systemów informatycznych, walidowanym w ponad 1300 badaniach [Bangor i in., 2009].

Kwestionariusz składa się z 10 stwierdzeń ocenianych na pięciostopniowej skali Likerta (1 — zdecydowanie się nie zgadzam, 5 — zdecydowanie się zgadzam). Stwierdzenia naprzemiennie sformułowane są pozytywnie i negatywnie, co zapobiega tendencji do automatycznego zaznaczania skrajnych odpowiedzi.

Wynik SUS mieści się w zakresie 0-100 punktów i interpretowany jest następująco [Bangor i in., 2009]:
- poniżej 50: niska użyteczność (nieakceptowalna);
- 50-68: marginalna użyteczność;
- 68-80: dobra użyteczność;
- powyżej 80: wysoka użyteczność (doskonała).

Na potrzeby badania kwestionariusz SUS przetłumaczono na język polski, zachowując oryginalną strukturę 10 stwierdzeń.

### 3.4.8. Adaptacja Kart ocen z Procedury MZ v2.3

Jako dodatkowe narzędzie ewaluacyjne zastosowano adaptację **Kart ocen** zdefiniowanych w Procedurze postępowania na wypadek wystąpienia zdarzenia z dużą liczbą poszkodowanych [Ministerstwo Zdrowia, 2024, zał. nr 15-26]. W oryginale Karty ocen stanowią ustrukturyzowany formularz wypełniany przez uczestników akcji ratowniczej po jej zakończeniu, służący identyfikacji obszarów wymagających poprawy.

Na potrzeby niniejszego badania zaadaptowano format Kart ocen, stosując oryginalną trzystopniową skalę:

| Ocena | Znaczenie |
|-------|-----------|
| **Prawidłowa** | Zadanie zrealizowane poprawnie, bez istotnych uchybień |
| **Nieoptymalna** | Zadanie zrealizowane, ale z uchybieniami wpływającymi na efektywność |
| **Wymaga poprawy** | Zadanie zrealizowane z istotnymi brakami lub niezrealizowane |

Ocenie podlegały trzy obszary, odpowiadające wymiarom ocenianym w oryginalnych Kartach ocen:

1. **Przepływ informacji** — czy dane o poszkodowanych (kategoria, numer, obrażenia) były rejestrowane w sposób umożliwiający szybki dostęp i odczyt?
2. **Podejmowanie decyzji** — czy klasyfikacja segregacyjna i dyslokacja poszkodowanych były realizowane w sposób systematyczny i zgodny z algorytmem START?
3. **Dokumentacja i raportowanie** — czy raport końcowy zawierał elementy wymagane przez Raport GDM [zał. nr 14] i był sporządzony w sposób czytelny?

Oceny dokonywał niezależny obserwator (autor badania lub wyznaczony asystent) na podstawie obserwacji przebiegu sesji eksperymentalnej oraz analizy wytworzonej dokumentacji (tradycyjnej lub cyfrowej). Każdy uczestnik otrzymywał odrębną ocenę dla sesji tradycyjnej i sesji z aplikacją, co umożliwiało porównanie jakości realizacji zadań KAM między obiema metodami w ramach schematu crossover.

Zastosowanie adaptacji oficjalnego narzędzia ewaluacyjnego z Procedury MZ v2.3 wzmacnia ekologiczną trafność badania — uczestnik jest oceniany według kryteriów, które obowiązują w realnej praktyce ratowniczej.

### 3.4.9. Rola Dyspozytora Medycznego (DM) w eksperymencie

W celu zwiększenia ekologicznej trafności badania i odwzorowania realistycznego przepływu informacji zdefiniowanego w Procedurze MZ v2.3, prowadzący eksperyment pełnił jednocześnie rolę Dyspozytora Medycznego (DM) — uproszczoną, łączącą funkcje DM-W i GDM.

**Uzasadnienie wprowadzenia roli DM:**
W warunkach realnego zdarzenia masowego KAM nie działa w izolacji — prowadzi ciągłą wymianę informacji z DM-W, przekazując raporty z rozpoznania i segregacji, odbierając informacje o zadysponowanych ZRM i możliwościach szpitali, wspólnie planując dyslokację poszkodowanych [Ministerstwo Zdrowia, 2024, rozdziały 8 i 13]. Eksperyment bez tego elementu testowałby jedynie umiejętność wypełniania formularzy, a nie realizację zadań KAM w kontekście procedury.

**Poziom wierności symulacji:**
Przyjęto umiarkowany poziom wierności (*moderate fidelity*): DM uczestniczył w ustrukturyzowanej wymianie informacji na granicach faz (po segregacji, przed dyslokacją, przed raportem), ale nie symulował ciągłej korespondencji radiowej ani wielostronnnej koordynacji. Takie podejście zapewnia realizm w kluczowych momentach decyzyjnych bez wprowadzania niekontrolowanej zmienności.

**Standaryzacja roli DM:**
Wszystkie wypowiedzi DM były skryptowane — prowadzący odczytywał standaryzowane komunikaty z protokołu eksperymentu. Dla pytań uczestników spoza skryptu opracowano bibliotekę standaryzowanych odpowiedzi. DM nigdy nie udzielał wskazówek dotyczących klasyfikacji segregacyjnej ani dyslokacji — jedynie przekazywał informacje i potwierdzał odbiór raportów. Identyczność komunikatów DM we wszystkich sesjach (niezależnie od metody i scenariusza) zapewniała, że rola DM stanowiła kontrolowaną stałą, a nie zmienną.

## 3.5. Część II — Ankieta ogólnopolska

### 3.5.1. Cel i charakter ankiety

Ankieta stanowiła komplementarne źródło danych, rozszerzające zakres badania poza grupę eksperymentalną. O ile eksperyment crossover dostarczał obiektywnych danych pomiarowych na temat efektywności aplikacji, ankieta miała na celu zbadanie szerszego kontekstu: postaw, doświadczeń i barier związanych z cyfryzacją segregacji medycznej wśród polskich ratowników.

Ankieta miała charakter badania przekrojowego (*cross-sectional*) — dane zbierano jednorazowo, uzyskując migawkowy obraz badanego zjawiska w określonym punkcie czasowym.

### 3.5.2. Populacja i próba

**Populacja docelowa:** Ratownicy medyczni zatrudnieni w systemie Państwowego Ratownictwa Medycznego w Polsce oraz studenci ratownictwa medycznego na studiach licencjackich i magisterskich.

**Metoda doboru próby:** Dobór celowy z elementami kuli śnieżnej (*snowball sampling*). Ankietę dystrybuowano poprzez:
- zamknięte grupy ratowników medycznych w mediach społecznościowych (Facebook);
- fora branżowe i grupy dyskusyjne;
- kontakty zawodowe autora w środowisku ratowniczym;
- prośby o dalsze udostępnianie (efekt kuli śnieżnej).

**Planowana wielkość próby:** Minimum 100 respondentów, cel optymalny 200+. Wielkość próby ustalono na podstawie analizy mocy testu: przy zakładanym średnim rozmiarze efektu (d=0,5), poziomie istotności α=0,05 i mocy testu 0,80, minimalna wielkość próby dla testu chi-kwadrat z dwiema zmiennymi dychotomicznymi wynosi n≈88 [Cohen, 1992].

**Kryteria włączenia:**
- posiadanie kwalifikacji ratownika medycznego (dyplom licencjata lub magistra ratownictwa medycznego) LUB status studenta ratownictwa medycznego (co najmniej II rok studiów);
- wyrażenie zgody na udział przez rozpoczęcie wypełniania ankiety.

### 3.5.3. Narzędzie badawcze — kwestionariusz ankiety

Ankieta została opracowana przez autora na potrzeby niniejszego badania. Kwestionariusz składał się z następujących sekcji:

**Sekcja A: Dane demograficzne i zawodowe (7 pytań)**
1. Wiek (przedziały: 20-25, 26-30, 31-35, 36-40, 41-45, powyżej 45 lat)
2. Płeć (kobieta, mężczyzna, nie chcę podawać)
3. Wykształcenie (licencjat ratownictwa medycznego, magister ratownictwa medycznego, inne — jakie?)
4. Status zawodowy (student ratownictwa, czynny ratownik medyczny, oba, inny)
5. Staż pracy w ratownictwie medycznym (brak, poniżej 1 roku, 1-3 lata, 4-7 lat, 8-15 lat, powyżej 15 lat)
6. Miejsce pracy (zespół podstawowy „P", zespół specjalistyczny „S", SOR/IP, LPR, inne, nie dotyczy)
7. Wielkość miejscowości pracy (do 10 tys., 10-50 tys., 50-100 tys., 100-500 tys., powyżej 500 tys.)

**Sekcja B: Doświadczenie ze zdarzeniami masowymi (5 pytań)**
8. Czy uczestniczył/a Pan/Pani w szkoleniu z zakresu segregacji medycznej w zdarzeniach masowych? (Tak / Nie)
9. Jaki system segregacji był omawiany na szkoleniu? (START, JumpSTART, SALT, Sieve/Sort, inny, nie pamiętam) — możliwość wielokrotnego wyboru
10. Czy uczestniczył/a Pan/Pani w realnym zdarzeniu masowym (>5 poszkodowanych)? (Tak — ile razy? / Nie)
11. Czy uczestniczył/a Pan/Pani w ćwiczeniach/symulacjach zdarzenia masowego? (Tak — ile razy? / Nie)
12. Jak ocenia Pan/Pani swoje przygotowanie do działania w zdarzeniu masowym? (skala Likerta 1-5: 1 — zdecydowanie nieprzygotowany, 5 — w pełni przygotowany)

**Sekcja C: Stosunek do narzędzi cyfrowych (6 pytań)**
13. Czy korzysta Pan/Pani z aplikacji medycznych w pracy zawodowej? (Tak — regularnie / Tak — sporadycznie / Nie)
14. Jeśli tak, z jakich? (pytanie otwarte)
15. Czy słyszał/a Pan/Pani o aplikacjach mobilnych do segregacji medycznej w zdarzeniach masowych? (Tak / Nie)
16. Jak ocenia Pan/Pani przydatność aplikacji mobilnej wspierającej segregację START w zdarzeniu masowym? (skala Likerta 1-5: 1 — zdecydowanie nieprzydatna, 5 — zdecydowanie przydatna)
17. Czy chciałby/chciałaby Pan/Pani korzystać z takiej aplikacji na realnym zdarzeniu? (Zdecydowanie tak / Raczej tak / Trudno powiedzieć / Raczej nie / Zdecydowanie nie)
18. Jakie wymagania powinna spełniać taka aplikacja, aby była przydatna w praktyce? (pytanie otwarte)

**Sekcja D: Bariery i obawy (4 pytania)**
19. Jakie dostrzega Pan/Pani przeszkody w stosowaniu aplikacji mobilnych podczas zdarzeń masowych? (wielokrotny wybór + „inne"):
    - Ryzyko awarii/rozładowania telefonu
    - Trudności z obsługą w rękawiczkach
    - Brak czasu na obsługę aplikacji
    - Problemy z łącznością/internetem
    - Brak zaufania do technologii w sytuacjach krytycznych
    - Brak oficjalnego zatwierdzenia/rekomendacji
    - Przyzwyczajenie do metod tradycyjnych
    - Inne (jakie?)
20. Które z powyższych barier uważa Pan/Pani za najpoważniejszą? (jedna odpowiedź)
21. Co mogłoby przekonać Pana/Panią do stosowania takiej aplikacji? (wielokrotny wybór):
    - Oficjalna rekomendacja KGP/Konsultanta Krajowego
    - Szkolenie z obsługi aplikacji
    - Możliwość pracy offline
    - Pozytywne wyniki badań naukowych
    - Rekomendacja kolegów z zespołu
    - Inne (jakie?)
22. Czy uważa Pan/Pani, że cyfrowe narzędzia wspomagające segregację staną się standardem w ciągu najbliższych 5-10 lat? (Zdecydowanie tak / Raczej tak / Trudno powiedzieć / Raczej nie / Zdecydowanie nie)

**Sekcja E: Znajomość Procedury MZ v2.3 i narzędzi dokumentacyjnych (5 pytań)**
23. Czy zna Pan/Pani aktualną *Procedurę postępowania na wypadek wystąpienia zdarzenia z dużą liczbą poszkodowanych* (MZ v2.3, 2024)? (Tak, znam dobrze / Tak, słyszałem/am o niej / Nie znam)
24. Czy ćwiczył/a Pan/Pani wypełnianie Tabeli Dyslokacji Poszkodowanych [zał. nr 12] z Procedury MZ? (Tak — na szkoleniu / Tak — na realnym zdarzeniu / Nie)
25. Jak ocenia Pan/Pani praktyczność papierowej Tabeli Dyslokacji Poszkodowanych w warunkach zdarzenia masowego? (skala Likerta 1-5: 1 — zdecydowanie niepraktyczna, 5 — w pełni praktyczna)
26. Czy uważa Pan/Pani, że cyfrowa wersja Tabeli Dyslokacji Poszkodowanych i Raportu GDM byłaby praktyczna w warunkach zdarzenia? (Zdecydowanie tak / Raczej tak / Trudno powiedzieć / Raczej nie / Zdecydowanie nie)
27. Jakie elementy Procedury MZ v2.3 najtrudniej realizować w warunkach zdarzenia masowego? (wielokrotny wybór):
    - Wypełnianie Tabeli Dyslokacji Poszkodowanych
    - Prowadzenie Tabeli Szpitali
    - Sporządzenie Raportu GDM
    - Komunikacja KAM–DM-W
    - Koordynacja dyslokacji z wieloma szpitalami
    - Inne (jakie?)

Łączna liczba pytań: 27. Szacowany czas wypełniania: 10-15 minut.

### 3.5.4. Pilotaż ankiety

Przed właściwym badaniem przeprowadzono pilotaż kwestionariusza na grupie 5-10 osób (ratownicy medyczni niezaangażowani w badanie główne) w celu:
- weryfikacji zrozumiałości pytań;
- oceny czasu wypełniania;
- identyfikacji błędów technicznych formularza elektronicznego;
- zbierania uwag o brakujących opcjach odpowiedzi.

Na podstawie wyników pilotażu wprowadzono ewentualne korekty w treści i kolejności pytań.

### 3.5.5. Platforma i dystrybucja

Ankietę przygotowano w formie elektronicznej za pomocą narzędzia Google Forms (bezpłatne, RODO-kompatybilne przy odpowiedniej konfiguracji). Formularz skonfigurowano tak, aby:
- nie zbierał adresów e-mail (anonimowość);
- wymagał odpowiedzi na pytania kluczowe (pola obowiązkowe);
- uniemożliwiał wielokrotne wypełnienie z tego samego urządzenia;
- zawierał klauzulę informacyjną o przetwarzaniu danych na pierwszej stronie.

Okres zbierania odpowiedzi zaplanowano na 4-6 tygodni.

## 3.6. Opis narzędzia badawczego — aplikacja TRIAGE MCI

### 3.6.1. Charakterystyka techniczna

Aplikacja TRIAGE MCI jest progresywną aplikacją internetową (PWA — *Progressive Web App*) opracowaną przez autora niniejszej pracy. Główne cechy techniczne:

- **Technologia:** HTML5, CSS3, JavaScript (Vanilla JS) — jednoplikowa aplikacja;
- **Instalacja:** Możliwość instalacji na ekranie głównym smartfona (Android, iOS) lub komputera (Windows, macOS) bez konieczności pobierania ze sklepu z aplikacjami;
- **Praca offline:** Pełna funkcjonalność bez dostępu do internetu dzięki mechanizmowi Service Worker;
- **Przechowywanie danych:** Lokalnie na urządzeniu użytkownika (localStorage), bez transmisji na zewnętrzne serwery;
- **Wielojęzyczność:** Interfejs dostępny w 7 językach (polski, angielski, włoski, francuski, niemiecki, czeski, portugalski);
- **Licencja:** Otwarte oprogramowanie (GPLv3), bezpłatne.

### 3.6.2. Funkcjonalność i odpowiedniość z Procedurą MZ v2.3

Aplikacja realizuje trzy główne funkcje, z których każda odpowiada konkretnemu narzędziu dokumentacyjnemu zdefiniowanemu w obowiązującej *Procedurze postępowania na wypadek wystąpienia zdarzenia z dużą liczbą poszkodowanych* [Ministerstwo Zdrowia, 2024]:

**Moduł segregacji (SEGREGACJA)** — odpowiednik: **Tabela Dyslokacji Poszkodowanych** [zał. nr 12]
Interaktywny kreator prowadzący użytkownika krok po kroku przez algorytm START. Na każdym etapie wyświetlane jest pytanie kliniczne z podpowiedzią. Użytkownik odpowiada TAK/NIE, a aplikacja automatycznie klasyfikuje pacjenta do odpowiedniej kategorii (T1-T4). Każdy pacjent otrzymuje automatyczny numer indywidualny (P-001, P-002...), analogicznie do numeracji w Tabeli Dyslokacji. Po zakończeniu segregacji możliwe jest dodanie notatek o obrażeniach oraz ręczna korekta kategorii (override). Ścieżka decyzyjna jest automatycznie rejestrowana. Podział pacjentów na grupy kolorystyczne (czerwona, żółta, zielona, czarna) odpowiada grupom zdefiniowanym w procedurze.

**Moduł dysponowania (DYSPONOWANIE)** — odpowiednik: **Tabela Szpitali** [zał. nr 13] + część dyslokacyjna **Tabeli Dyslokacji** [zał. nr 12]
Umożliwia przypisanie pacjenta do zespołu ratownictwa medycznego (ZRM) i wskazanie szpitala docelowego — realizując funkcję dyslokacji poszkodowanych, którą Procedura MZ v2.3 definiuje jako wspólne zadanie KAM i DM-W. Zawiera edytowalną listę zespołów i szpitali, predefiniowaną dla danego kraju/języka. Wyświetla statusy transportów (w trakcie / dostarczony) i historię dysponowania.

**Moduł raportowania (RAPORT)** — odpowiednik: **Raport GDM z przebiegu zdarzenia** [zał. nr 14]
Automatyczne generowanie raportu ze zdarzenia zawierającego elementy wymagane przez Procedurę MZ v2.3: zestawienie liczbowe poszkodowanych z podziałem na grupy segregacyjne, rozkład pacjentów wg szpitali docelowych, przypisanie ZRM do pacjentów, statusy transportów, listę pacjentów pozostających na miejscu. Raport dostępny w formie wizualnej (w aplikacji) i tekstowej (do skopiowania do schowka lub wysłania).

Poniższa tabela zestawia odpowiedniość funkcjonalną:

| Narzędzie z Procedury MZ v2.3 | Załącznik | Moduł TRIAGE MCI | Forma w procedurze | Forma w aplikacji |
|---|---|---|---|---|
| Tabela Dyslokacji Poszkodowanych | zał. 12 | Segregacja | Papierowa | Cyfrowa (automatyczna) |
| Tabela Szpitali | zał. 13 | Dysponowanie | Papierowa | Cyfrowa (edytowalna) |
| Raport GDM z przebiegu zdarzenia | zał. 14 | Raport | Papierowa | Cyfrowa (auto-generowana) |

### 3.6.3. Uzasadnienie wyboru technologii PWA

Decyzja o zastosowaniu technologii PWA (zamiast aplikacji natywnej) była podyktowana następującymi przesłankami:

- **Niezależność od platformy** — jedna wersja aplikacji działa na Androidzie, iOS i Windows;
- **Brak konieczności instalacji ze sklepu** — eliminacja bariery wejścia (konieczność posiadania konta Google/Apple, zgody organizacji);
- **Natychmiastowa aktualizacja** — użytkownik zawsze korzysta z najnowszej wersji;
- **Pełna funkcjonalność offline** — dzięki mechanizmowi Service Worker aplikacja po pierwszym otwarciu działa bez internetu;
- **Prywatność danych** — brak konieczności tworzenia konta, logowania ani przesyłania danych na serwer.

## 3.7. Analiza statystyczna

### 3.7.1. Oprogramowanie

Analizę statystyczną przeprowadzono z użyciem oprogramowania [UWAGA: uzupełnić — np. Statistica 13.3, IBM SPSS Statistics 29, R 4.3 lub inne dostępne dla autora].

### 3.7.2. Statystyki opisowe

Dla zmiennych ilościowych obliczono: średnią arytmetyczną, odchylenie standardowe, medianę, wartości minimalne i maksymalne oraz kwartyle (Q1, Q3). Dla zmiennych jakościowych obliczono częstości bezwzględne i względne (odsetki).

### 3.7.3. Weryfikacja normalności rozkładu

Normalność rozkładu zmiennych ilościowych weryfikowano testem Shapiro-Wilka (dla prób n<50). Za kryterium normalności przyjęto poziom istotności p>0,05.

### 3.7.4. Testy hipotez — eksperyment crossover

Dla porównania wyników między metodami (w schemacie crossover) zastosowano:

- **Test t-Studenta dla prób zależnych** (*paired t-test*) — w przypadku spełnienia założenia o normalności rozkładu różnic. Test ten porównuje średnie wartości dwóch pomiarów wykonanych na tych samych uczestnikach, co odpowiada schematowi crossover.
- **Test Wilcoxona** (*Wilcoxon signed-rank test*) — nieparametryczny odpowiednik testu t, stosowany w przypadku niespełnienia założenia normalności.

Dla porównania proporcji (np. odsetek prawidłowych klasyfikacji) zastosowano test McNemara.

Dodatkowo obliczono:
- **Rozmiar efektu** (d Cohena dla prób zależnych) — w celu oceny praktycznego znaczenia różnic;
- **95% przedział ufności** dla różnicy średnich.

### 3.7.5. Testy hipotez — ankieta

Dla danych ankietowych zastosowano:
- **Test chi-kwadrat** (χ²) — dla porównania rozkładów zmiennych jakościowych między podgrupami (np. stosunek do cyfryzacji wg stażu pracy);
- **Test U Manna-Whitneya** — dla porównania zmiennych porządkowych (skala Likerta) między dwiema grupami;
- **Test Kruskala-Wallisa** — dla porównania zmiennych porządkowych między trzema lub więcej grupami;
- **Korelacja Spearmana** — dla oceny związku między zmiennymi porządkowymi.

### 3.7.6. Poziom istotności

We wszystkich testach przyjęto poziom istotności α=0,05. Wyniki o wartości p<0,05 uznawano za statystycznie istotne.

### 3.7.7. Analiza efektu kolejności (carry-over)

W schemacie crossover istnieje ryzyko efektu przenoszenia (*carry-over effect*) — doświadczenie z pierwszej sesji może wpływać na wyniki w drugiej. W celu weryfikacji tego efektu przeprowadzono:

- Porównanie wyników sesji 1 między grupami AB i BA (test t dla prób niezależnych lub test U Manna-Whitneya) — brak istotnej różnicy potwierdza skuteczność randomizacji;
- Analizę interakcji metoda × kolejność w dwuczynnikowej analizie wariancji (ANOVA) dla pomiarów powtarzanych — istotna interakcja wskazywałaby na efekt carry-over.

### 3.7.8. Analiza mocy testu dla eksperymentu crossover

Wielkość próby dla eksperymentu crossover ustalono na podstawie a priori analizy mocy testu (*power analysis*) dla testu t-Studenta dla prób zależnych (paired t-test), który stanowi podstawowy test statystyczny w schemacie crossover.

**Założenia:**
- Poziom istotności: α = 0,05 (test dwustronny)
- Moc testu: 1 − β = 0,80
- Zakładany rozmiar efektu: d = 0,50 (efekt średni wg klasyfikacji Cohena [1992])
- Korelacja między pomiarami: r = 0,50 (konserwatywne założenie)

**Obliczenie minimalnej wielkości próby:**

Dla testu t dla prób zależnych minimalna wielkość próby obliczana jest ze wzoru:

n = (Z_α/2 + Z_β)² / d² × (1 − r) × 2

Przy powyższych założeniach:
- Z_α/2 = 1,96 (dla α = 0,05, test dwustronny)
- Z_β = 0,84 (dla mocy 0,80)

n = (1,96 + 0,84)² / 0,50² × (1 − 0,50) × 2 ≈ **34 osoby**

Alternatywnie, stosując standardowy wzór G*Power dla paired t-test z d = 0,50 i mocą 0,80, minimalna wymagana próba wynosi n = 34 [Faul i in., 2007].

**Uzasadnienie planowanej wielkości próby:**

Planowana próba n = 40-50 uczestników przekracza minimalną wymaganą wielkość o 18-47%, co zapewnia:
- zapas na ewentualne wykluczenia (np. nieprawidłowo wypełnione formularze, wycofanie zgody);
- wystarczającą moc statystyczną do wykrycia efektu średniego (d ≥ 0,50) we wszystkich hipotezach H1-H3;
- możliwość przeprowadzenia eksploracyjnej analizy podgrup (studenci vs. czynni ratownicy), choć taka analiza będzie miała ograniczoną moc statystyczną przy podziale na dwie podgrupy.

W przypadku uzyskania efektu mniejszego niż zakładany (d < 0,50), próba n = 40-50 zapewnia moc testu odpowiednio: 0,63 dla d = 0,40, 0,44 dla d = 0,30. Ograniczenie to zostanie uwzględnione w dyskusji wyników.

## 3.8. Aspekty etyczne

### 3.8.1. Zgoda uczestników

Wszyscy uczestnicy eksperymentu wyrazili pisemną, świadomą zgodę na udział w badaniu po zapoznaniu się z informacją obejmującą: cel badania, przebieg, dobrowolność udziału, prawo do wycofania się na każdym etapie bez podania przyczyny oraz sposób przetwarzania danych.

Uczestnicy ankiety wyrażali zgodę przez rozpoczęcie wypełniania kwestionariusza, po zapoznaniu się z klauzulą informacyjną na pierwszej stronie formularza.

### 3.8.2. Anonimowość i ochrona danych

Dane eksperymentalne gromadzono z użyciem kodów identyfikacyjnych (np. U01, U02...), bez rejestrowania danych osobowych. Kwestionariusze papierowe (zgody) przechowywano oddzielnie od danych pomiarowych, uniemożliwiając powiązanie wyników z tożsamością uczestnika.

Ankieta internetowa nie zbierała adresów e-mail ani adresów IP. Dane przechowywano na serwerach Google (Google Forms) zgodnie z polityką prywatności Google Workspace i przepisami RODO.

### 3.8.3. Bezpieczeństwo uczestników

Badanie przeprowadzono w warunkach symulacyjnych, bez udziału realnych pacjentów. Uczestnicy nie byli narażeni na żadne ryzyko fizyczne ani psychologiczne wykraczające poza standardowe warunki szkolenia. Badanie nie wymagało opinii komisji bioetycznej, gdyż nie stanowiło eksperymentu medycznego w rozumieniu ustawy z dnia 5 grudnia 1996 r. o zawodach lekarza i lekarza dentysty.

### 3.8.4. Konflikt interesów i potencjalne źródła stronniczości

Autor niniejszej pracy jest jednocześnie twórcą i deweloperem aplikacji TRIAGE MCI, która stanowi przedmiot interwencji eksperymentalnej. Ta podwójna rola — badacza i twórcy narzędzia — stanowi potencjalne źródło stronniczości (*bias*), które należy jawnie zidentyfikować i kontrolować. Poniżej przedstawiono analizę potencjalnych zagrożeń oraz zastosowane mechanizmy mitygacji:

**1. Stronniczość w projektowaniu badania (*design bias*)**
- *Zagrożenie:* Badanie mogło zostać zaprojektowane w sposób faworyzujący aplikację (np. zbyt łatwe scenariusze, zbyt krótki czas na metodę tradycyjną).
- *Mitygacja:* Scenariusze symulacyjne zostały zwalidowane przez dwóch niezależnych, doświadczonych ratowników medycznych niebędących współtwórcami aplikacji. Rozkład kategorii segregacyjnych (T1:3, T2:3, T3:4, T4:2) jest identyczny w obu scenariuszach. Schemat crossover eliminuje zmienność międzyosobniczą — każdy uczestnik jest swoją własną kontrolą, co utrudnia systemowe faworyzowanie jednej metody.

**2. Stronniczość w ocenie wyników (*assessment bias*)**
- *Zagrożenie:* Autor mógłby nieświadomie oceniać wyniki korzystniej dla sesji z aplikacją.
- *Mitygacja:* Trafność segregacji oceniana jest na podstawie obiektywnego klucza referencyjnego (zgodność/niezgodność z kluczem — pomiar binarny, niemożliwy do subiektywnej interpretacji). Kompletność raportu oceniana jest za pomocą standaryzowanej listy kontrolnej opartej na Raporcie GDM [zał. nr 14 Procedury MZ v2.3] — pomiar binarny (element obecny/brak). Czas mierzony stoperem przez obserwatora. Wynik SUS obliczany jest algorytmicznie ze standaryzowanego kwestionariusza. W miarę możliwości ocenę raportów przeprowadzi drugi, niezależny oceniacz (*inter-rater reliability*).

**3. Stronniczość w prezentacji wyników (*reporting bias*)**
- *Zagrożenie:* Autor mógłby selektywnie przedstawiać wyniki korzystne dla aplikacji.
- *Mitygacja:* Wszystkie hipotezy (H1-H4) zostały sformułowane a priori, przed przeprowadzeniem badania. Wyniki zostaną przedstawione w pełni, niezależnie od tego, czy potwierdzają, czy obalają postawione hipotezy. Dane surowe zostaną udostępnione w aneksie pracy.

**4. Efekt entuzjazmu (*enthusiasm effect*)**
- *Zagrożenie:* Autor, prezentując aplikację uczestnikom, mógłby nieświadomie wzbudzać pozytywne nastawienie do narzędzia.
- *Mitygacja:* Instrukcja dotycząca aplikacji jest standaryzowana (5 minut, identyczna treść dla wszystkich uczestników) i ogranicza się do opisu funkcji, bez wartościowania ani porównywania z metodą tradycyjną.

Powyższe mechanizmy nie eliminują ryzyka stronniczości w sposób absolutny — całkowite wyeliminowanie tego ryzyka wymagałoby przeprowadzenia badania przez niezależny zespół badawczy, co wykracza poza możliwości pracy magisterskiej. Transparentne ujawnienie potencjalnego konfliktu interesów oraz zastosowanie opisanych mechanizmów mitygacji stanowią standardową praktykę w badaniach, w których twórca narzędzia jest jednocześnie badaczem [Ioannidis, 2005].

## 3.9. Ograniczenia metodologiczne

Zaprojektowane badanie obarczone jest następującymi ograniczeniami, które należy uwzględnić przy interpretacji wyników:

1. **Warunki symulacyjne vs. realne zdarzenie.** Symulacja z kartami pacjentów nie odwzorowuje w pełni warunków realnego zdarzenia masowego (stres, hałas, wielozadaniowość, warunki atmosferyczne, kontakt z poszkodowanymi). Wyniki mogą nie być w pełni transferowalne do sytuacji realnej.

2. **Efekt uczenia się w crossover.** Mimo randomizacji kolejności i zastosowania dwóch różnych scenariuszy, nie można całkowicie wykluczyć wpływu doświadczenia z pierwszej sesji na wyniki drugiej.

3. **Dobór próby w ankiecie.** Metoda kuli śnieżnej i dystrybucja przez media społecznościowe mogą prowadzić do nadreprezentacji osób aktywnych w mediach społecznościowych i potencjalnie bardziej otwartych na nowe technologie, co może zawyżać pozytywne oceny narzędzi cyfrowych.

4. **Brak zaślepienia.** Ze względu na charakter interwencji (widoczna różnica między kartką a smartfonem) zaślepienie uczestników nie było możliwe. Świadomość uczestników, że ich wyniki są mierzone, mogła wpływać na ich zachowanie (efekt Hawthorne'a).

5. **Jednorazowy pomiar.** Badanie nie uwzględnia krzywej uczenia się — wyniki mogłyby być inne po dłuższym okresie korzystania z aplikacji. Badanie podłużne (longitudinalne) stanowiłoby wartościowe uzupełnienie w przyszłości.

6. **Specyficzność narzędzia.** Wyniki dotyczą konkretnej aplikacji (TRIAGE MCI) i nie mogą być automatycznie uogólniane na inne rozwiązania cyfrowe do segregacji medycznej.

7. **Konflikt interesów.** Autor jest twórcą badanego narzędzia, co stanowi potencjalne źródło stronniczości. Zastosowane mechanizmy mitygacji (klucz referencyjny, standaryzowane narzędzia pomiarowe, schemat crossover, transparentne raportowanie) ograniczają, ale nie eliminują tego ryzyka (patrz 3.8.4).

Powyższe ograniczenia zostały szczegółowo omówione w rozdziale 6 (Dyskusja).

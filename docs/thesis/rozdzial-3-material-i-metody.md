# Rozdział 3. Materiał i metody

## 3.1. Cel i uzasadnienie badania

Celem niniejszej pracy było porównanie efektywności cyfrowego systemu wspomagania segregacji medycznej (aplikacja TRIAGE MCI) z metodą tradycyjną (bez wsparcia cyfrowego) w warunkach symulowanego zdarzenia masowego. Badanie miało charakter dwuczęściowy: obejmowało eksperyment symulacyjny z zastosowaniem schematu crossover oraz ogólnopolską ankietę skierowaną do ratowników medycznych.

Uzasadnienie podjęcia badania wynika z przesłanek zidentyfikowanych w przeglądzie literatury (rozdział 2):

1. Algorytm START, choć ugruntowany w praktyce, charakteryzuje się znaczącymi wskaźnikami overtriage (30-50%) i undertriage (5-15%), których jednym ze źródeł jest błąd ludzki wynikający ze stresu, zmęczenia i niedostatecznego doświadczenia [Jenkins i in., 2008].
2. Cyfryzacja procesu segregacji ma potencjał do standaryzacji, automatyzacji dokumentacji i usprawnienia raportowania, lecz brakuje badań porównawczych oceniających ten potencjał w kontekście polskiego systemu ratownictwa medycznego.
3. Istniejące rozwiązania cyfrowe nie spełniają łącznie wszystkich wymagań praktyki polowej (kompletność funkcjonalna, praca offline, dostępność w języku polskim, bezpłatność).

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
- Wyjaśnienie zasad oceny i dokumentacji wyników.

**Faza 2: Sesja 1 (bez limitu czasu)**
Każdy uczestnik indywidualnie przeprowadzał segregację 12 poszkodowanych (na kartach) zgodnie z przypisaną metodą (tradycyjną lub z aplikacją). Mierzone zmienne:
- **Czas** — od momentu rozpoczęcia (otrzymanie kompletu kart pacjentów) do momentu zakończenia (deklaracja ukończenia segregacji i raportu). Mierzony stoperem przez obserwatora.
- **Trafność** — porównanie przypisanych kategorii z kluczem referencyjnym.
- **Raport** — uczestnik proszony o sporządzenie raportu ze zdarzenia (metodą tradycyjną: na kartce; z aplikacją: za pomocą wbudowanej funkcji raportowania).

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

Kompletność raportu ze zdarzenia oceniano za pomocą standaryzowanej listy kontrolnej obejmującej 10 elementów, z których każdy oceniano binarnie (0 — brak, 1 — obecny):

1. Nazwa/opis zdarzenia
2. Czas rozpoczęcia działań
3. Całkowita liczba poszkodowanych
4. Podział na kategorie segregacyjne (T1-T4) z liczbami
5. Lista pacjentów z przypisanymi kategoriami
6. Informacja o pacjentach zadysponowanych do transportu
7. Wskazanie szpitali docelowych
8. Przypisanie zespołów ZRM do pacjentów
9. Status transportów (w trakcie / dostarczony)
10. Czas sporządzenia raportu

Maksymalny wynik: 10 punktów. Lista została opracowana na podstawie wytycznych dotyczących raportowania ze zdarzeń masowych obowiązujących w polskim systemie ratownictwa medycznego.

### 3.4.7. Kwestionariusz SUS (System Usability Scale)

Do oceny użyteczności aplikacji TRIAGE MCI zastosowano kwestionariusz SUS (*System Usability Scale*) opracowany przez Brooke'a [1996]. SUS jest jednym z najszerzej stosowanych standaryzowanych narzędzi oceny użyteczności systemów informatycznych, walidowanym w ponad 1300 badaniach [Bangor i in., 2009].

Kwestionariusz składa się z 10 stwierdzeń ocenianych na pięciostopniowej skali Likerta (1 — zdecydowanie się nie zgadzam, 5 — zdecydowanie się zgadzam). Stwierdzenia naprzemiennie sformułowane są pozytywnie i negatywnie, co zapobiega tendencji do automatycznego zaznaczania skrajnych odpowiedzi.

Wynik SUS mieści się w zakresie 0-100 punktów i interpretowany jest następująco [Bangor i in., 2009]:
- poniżej 50: niska użyteczność (nieakceptowalna);
- 50-68: marginalna użyteczność;
- 68-80: dobra użyteczność;
- powyżej 80: wysoka użyteczność (doskonała).

Na potrzeby badania kwestionariusz SUS przetłumaczono na język polski, zachowując oryginalną strukturę 10 stwierdzeń.

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

Łączna liczba pytań: 22. Szacowany czas wypełniania: 8-12 minut.

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

### 3.6.2. Funkcjonalność

Aplikacja realizuje trzy główne funkcje:

**Moduł segregacji (SEGREGACJA):** Interaktywny kreator prowadzący użytkownika krok po kroku przez algorytm START. Na każdym etapie wyświetlane jest pytanie kliniczne z podpowiedzią. Użytkownik odpowiada TAK/NIE, a aplikacja automatycznie klasyfikuje pacjenta do odpowiedniej kategorii (T1-T4). Po zakończeniu segregacji możliwe jest dodanie notatek oraz ręczna korekta kategorii (override). Ścieżka decyzyjna jest automatycznie rejestrowana.

**Moduł dysponowania (DYSPONOWANIE):** Umożliwia przypisanie pacjenta do zespołu ratownictwa medycznego (ZRM) i wskazanie szpitala docelowego. Zawiera edytowalną listę zespołów i szpitali, predefiniowaną dla danego kraju/języka. Wyświetla statusy transportów (w trakcie / dostarczony) i historię dysponowania.

**Moduł raportowania (RAPORT):** Automatyczne generowanie kompletnego raportu ze zdarzenia, obejmującego: podsumowanie liczbowe (T1-T4), rozkład pacjentów wg szpitali docelowych, statusy transportów, listę pacjentów pozostających na miejscu. Raport dostępny w formie wizualnej (w aplikacji) i tekstowej (do skopiowania do schowka lub wysłania).

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

## 3.8. Aspekty etyczne

### 3.8.1. Zgoda uczestników

Wszyscy uczestnicy eksperymentu wyrazili pisemną, świadomą zgodę na udział w badaniu po zapoznaniu się z informacją obejmującą: cel badania, przebieg, dobrowolność udziału, prawo do wycofania się na każdym etapie bez podania przyczyny oraz sposób przetwarzania danych.

Uczestnicy ankiety wyrażali zgodę przez rozpoczęcie wypełniania kwestionariusza, po zapoznaniu się z klauzulą informacyjną na pierwszej stronie formularza.

### 3.8.2. Anonimowość i ochrona danych

Dane eksperymentalne gromadzono z użyciem kodów identyfikacyjnych (np. U01, U02...), bez rejestrowania danych osobowych. Kwestionariusze papierowe (zgody) przechowywano oddzielnie od danych pomiarowych, uniemożliwiając powiązanie wyników z tożsamością uczestnika.

Ankieta internetowa nie zbierała adresów e-mail ani adresów IP. Dane przechowywano na serwerach Google (Google Forms) zgodnie z polityką prywatności Google Workspace i przepisami RODO.

### 3.8.3. Bezpieczeństwo uczestników

Badanie przeprowadzono w warunkach symulacyjnych, bez udziału realnych pacjentów. Uczestnicy nie byli narażeni na żadne ryzyko fizyczne ani psychologiczne wykraczające poza standardowe warunki szkolenia. Badanie nie wymagało opinii komisji bioetycznej, gdyż nie stanowiło eksperymentu medycznego w rozumieniu ustawy z dnia 5 grudnia 1996 r. o zawodach lekarza i lekarza dentysty.

## 3.9. Ograniczenia metodologiczne

Zaprojektowane badanie obarczone jest następującymi ograniczeniami, które należy uwzględnić przy interpretacji wyników:

1. **Warunki symulacyjne vs. realne zdarzenie.** Symulacja z kartami pacjentów nie odwzorowuje w pełni warunków realnego zdarzenia masowego (stres, hałas, wielozadaniowość, warunki atmosferyczne, kontakt z poszkodowanymi). Wyniki mogą nie być w pełni transferowalne do sytuacji realnej.

2. **Efekt uczenia się w crossover.** Mimo randomizacji kolejności i zastosowania dwóch różnych scenariuszy, nie można całkowicie wykluczyć wpływu doświadczenia z pierwszej sesji na wyniki drugiej.

3. **Dobór próby w ankiecie.** Metoda kuli śnieżnej i dystrybucja przez media społecznościowe mogą prowadzić do nadreprezentacji osób aktywnych w mediach społecznościowych i potencjalnie bardziej otwartych na nowe technologie, co może zawyżać pozytywne oceny narzędzi cyfrowych.

4. **Brak zaślepienia.** Ze względu na charakter interwencji (widoczna różnica między kartką a smartfonem) zaślepienie uczestników nie było możliwe. Świadomość uczestników, że ich wyniki są mierzone, mogła wpływać na ich zachowanie (efekt Hawthorne'a).

5. **Jednorazowy pomiar.** Badanie nie uwzględnia krzywej uczenia się — wyniki mogłyby być inne po dłuższym okresie korzystania z aplikacji. Badanie podłużne (longitudinalne) stanowiłoby wartościowe uzupełnienie w przyszłości.

6. **Specyficzność narzędzia.** Wyniki dotyczą konkretnej aplikacji (TRIAGE MCI) i nie mogą być automatycznie uogólniane na inne rozwiązania cyfrowe do segregacji medycznej.

Powyższe ograniczenia zostały szczegółowo omówione w rozdziale 6 (Dyskusja).

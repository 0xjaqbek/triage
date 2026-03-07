# Rozdział 1. Wstęp

## 1.1. Wprowadzenie

Zdarzenia masowe stanowią jedno z największych wyzwań współczesnego ratownictwa medycznego. Sytuacja, w której liczba poszkodowanych przekracza bieżące możliwości ratownicze, wymaga od personelu medycznego nie tylko wiedzy i umiejętności klinicznych, ale przede wszystkim zdolności do szybkiego podejmowania decyzji w warunkach ekstremalnego stresu, deficytu czasu i ograniczonych zasobów. W takich okolicznościach kluczowym elementem łańcucha ratunkowego staje się segregacja medyczna — triage — czyli proces systematycznej klasyfikacji poszkodowanych według priorytetów leczenia i transportu [Koenig i Schultz, 2010].

Algorytm START (*Simple Triage and Rapid Treatment*), opracowany w 1983 roku w Stanach Zjednoczonych, jest obecnie najszerzej stosowanym narzędziem segregacji medycznej na świecie, w tym w polskim systemie Państwowego Ratownictwa Medycznego [Benson i in., 1996]. Jego konstrukcja opiera się na sekwencyjnej ocenie czterech podstawowych parametrów fizjologicznych: zdolności chodzenia, oddechu, krążenia i stanu neurologicznego. Dzięki prostocie i szybkości wykonania algorytm START może być stosowany przez personel o różnym poziomie doświadczenia, co stanowi jego niekwestionowaną zaletę.

Jednocześnie badania walidacyjne wskazują na istotne ograniczenia tego algorytmu. Wskaźnik overtriage (nadklasyfikacji) szacowany jest na 30-50%, co oznacza, że znaczna część poszkodowanych otrzymuje wyższy priorytet niż wynika to z ich rzeczywistego stanu klinicznego. Undertriage (niedoklasyfikacja) dotyczy 5-15% przypadków i niesie ze sobą poważniejsze konsekwencje — pacjent zakwalifikowany zbyt nisko może nie otrzymać pomocy na czas [Jenkins i in., 2008]. Źródłem tych błędów jest nie tyle sam algorytm, ile czynnik ludzki: stres, zmęczenie, niedostateczne doświadczenie, pominięcie kroków procedury lub błędna interpretacja parametrów klinicznych [Risavi i in., 2001].

Równolegle do wyzwań stojących przed ratownictwem medycznym dynamicznie rozwija się cyfryzacja ochrony zdrowia. Smartfony stały się powszechnym narzędziem pracy personelu medycznego, a aplikacje mobilne wspierające podejmowanie decyzji klinicznych zyskują coraz szersze zastosowanie w medycynie ratunkowej [Boulos i in., 2011; Wallis i in., 2017]. Postęp technologiczny — w szczególności rozwój progresywnych aplikacji internetowych (PWA), zdolnych do pełnego funkcjonowania bez dostępu do sieci — otwiera nowe możliwości w zakresie cyfrowego wspomagania procedur ratowniczych, w tym segregacji medycznej.

Pomimo rosnącego zainteresowania tematem, w polskiej literaturze naukowej brakuje badań empirycznych oceniających efektywność cyfrowych narzędzi segregacji medycznej w porównaniu z metodą tradycyjną. Niniejsza praca stanowi próbę wypełnienia tej luki badawczej.

## 1.2. Cel pracy

Głównym celem niniejszej pracy jest porównanie efektywności cyfrowego systemu wspomagania segregacji medycznej — aplikacji TRIAGE MCI — z metodą tradycyjną (bez wsparcia cyfrowego) w warunkach symulowanego zdarzenia masowego.

Cele szczegółowe obejmują:

1. Porównanie czasu potrzebnego do przeprowadzenia segregacji medycznej grupy poszkodowanych przy zastosowaniu aplikacji TRIAGE MCI i metody tradycyjnej.

2. Porównanie trafności klasyfikacji segregacyjnej (zgodności z kluczem referencyjnym) między obiema metodami.

3. Porównanie kompletności raportu ze zdarzenia generowanego przy pomocy aplikacji i sporządzanego metodą tradycyjną.

4. Ocenę użyteczności aplikacji TRIAGE MCI z perspektywy użytkowników (ratowników medycznych i studentów ratownictwa) z wykorzystaniem standaryzowanego kwestionariusza SUS.

5. Zbadanie postaw, doświadczeń i barier polskich ratowników medycznych wobec cyfrowych narzędzi wspomagających segregację medyczną w zdarzeniach masowych.

## 1.3. Hipotezy badawcze

Na podstawie przesłanek wynikających z przeglądu literatury sformułowano następujące hipotezy:

**H1 (hipoteza główna):** Zastosowanie aplikacji TRIAGE MCI istotnie skraca średni czas segregacji poszkodowanych w porównaniu z metodą tradycyjną.

Uzasadnienie: Algorytm zaimplementowany w aplikacji prowadzi użytkownika krok po kroku, eliminując czas potrzebny na przypominanie sobie kolejności kroków i kryteriów. Automatyczna klasyfikacja po udzieleniu odpowiedzi TAK/NIE jest szybsza niż ręczne przypisywanie kategorii. Ponadto automatyczne generowanie raportu eliminuje czasochłonny etap ręcznego dokumentowania wyników [Kirsch i in., 2012].

**H2:** Trafność klasyfikacji segregacyjnej jest istotnie wyższa przy zastosowaniu aplikacji niż przy metodzie tradycyjnej.

Uzasadnienie: Algorytm w aplikacji jest realizowany zawsze w pełnej, prawidłowej sekwencji, co eliminuje ryzyko pominięcia kroków, błędnej interpretacji kryteriów i subiektywnych odstępstw od procedury. Badania wskazują, że nawet przeszkolony personel popełnia błędy segregacyjne pod wpływem stresu [Risavi i in., 2001], a standaryzacja procesu cyfrowego ma potencjał do redukcji tych błędów.

**H3:** Kompletność raportu ze zdarzenia jest istotnie wyższa przy zastosowaniu aplikacji niż przy metodzie tradycyjnej.

Uzasadnienie: Aplikacja automatycznie gromadzi dane o każdym pacjencie (kategoria, czas, ścieżka decyzyjna, przypisany zespół, szpital docelowy, status transportu) i generuje kompletny raport jednym kliknięciem. Przy metodzie tradycyjnej raport wymaga ręcznego zestawiania informacji z notatek, kart i pamięci, co sprzyja pominięciom i błędom.

**H4:** Uczestnicy badania oceniają aplikację TRIAGE MCI jako narzędzie o wysokiej użyteczności, osiągając wynik powyżej 68 punktów w skali SUS (próg dobrej użyteczności według Bangora i in. [2009]).

Uzasadnienie: Aplikacja została zaprojektowana z myślą o minimalizmie interfejsu, intuicyjnej nawigacji i szybkości obsługi. Technologia PWA eliminuje barierę instalacji i konfiguracji. Praca offline zapewnia niezawodność w warunkach polowych.

## 1.4. Uzasadnienie wyboru tematu

Wybór tematu niniejszej pracy podyktowany był zarówno przesłankami naukowymi, jak i praktycznymi.

### Przesłanki naukowe

Przegląd literatury (rozdział 2) ujawnił istotną lukę badawczą: o ile cyfryzacja segregacji medycznej jest coraz szerzej dyskutowana w literaturze światowej, o tyle badania empiryczne — a w szczególności badania porównawcze z użyciem schematu eksperymentalnego — pozostają nieliczne. W polskiej literaturze naukowej autor nie zidentyfikował ani jednego badania porównującego efektywność cyfrowego narzędzia segregacji z metodą tradycyjną. Niniejsza praca stanowi zatem wkład w rozwój polskiej bazy dowodowej w tym zakresie.

### Przesłanki praktyczne

Autor niniejszej pracy jest czynnym ratownikiem medycznym z doświadczeniem w systemie Państwowego Ratownictwa Medycznego. Motywacją do opracowania aplikacji TRIAGE MCI było osobiste doświadczenie zawodowe i obserwacja, że:

- szkolenia z zakresu segregacji medycznej, choć obowiązkowe, są sporadyczne i nie zawsze skutkują trwałym utrwaleniem algorytmu;
- w warunkach realnego zdarzenia masowego ratownicy nierzadko mają trudności z prawidłową realizacją pełnej sekwencji algorytmu START;
- dokumentacja zdarzenia masowego prowadzona ręcznie jest czasochłonna, niepełna i podatna na błędy;
- istniejące rozwiązania cyfrowe nie spełniają potrzeb polskiego ratownictwa medycznego (brak języka polskiego, brak pracy offline, niekompletna funkcjonalność).

Aplikacja TRIAGE MCI powstała jako odpowiedź na te obserwacje — nie jako produkt komercyjny, lecz jako narzędzie o otwartym kodzie źródłowym, bezpłatne i dostępne dla każdego ratownika medycznego. Niniejsze badanie stanowi próbę naukowej weryfikacji, czy to narzędzie rzeczywiście poprawia efektywność segregacji w porównaniu z metodą tradycyjną.

### Przesłanki społeczne

Zdarzenia masowe, choć stosunkowo rzadkie, mają ogromny wpływ na zdrowie publiczne i bezpieczeństwo obywateli. Każda poprawa efektywności procesu segregacji — nawet o kilka procent — może przełożyć się na szybsze udzielenie pomocy pacjentom w stanie bezpośredniego zagrożenia życia, a w konsekwencji na zmniejszenie śmiertelności i powikłań. Jeśli badanie wykaże, że narzędzia cyfrowe mogą wspierać ratowników w podejmowaniu trafniejszych i szybszych decyzji segregacyjnych, wyniki te mogą stanowić podstawę do rekomendacji włączenia takich narzędzi do standardowego wyposażenia ratownika medycznego.

## 1.5. Zakres i struktura pracy

Niniejsza praca ma charakter badawczy i składa się z ośmiu rozdziałów.

**Rozdział 1 (Wstęp)** przedstawia tło problemu, cele pracy, hipotezy badawcze oraz uzasadnienie wyboru tematu.

**Rozdział 2 (Przegląd literatury)** zawiera systematyczny przegląd piśmiennictwa obejmujący: definicje i epidemiologię zdarzeń masowych, historię i zasady segregacji medycznej, szczegółowy opis algorytmu START i jego wariantów, przegląd innych systemów segregacyjnych, cyfryzację w medycynie ratunkowej oraz analizę istniejących rozwiązań cyfrowych do segregacji.

**Rozdział 3 (Materiał i metody)** opisuje schemat badania, w tym eksperyment crossover (grupa badana, randomizacja, scenariusze symulacyjne, zmienne, protokół) oraz ankietę ogólnopolską (narzędzie badawcze, dystrybucja, próba). Przedstawia również opis techniczny aplikacji TRIAGE MCI, plan analizy statystycznej oraz aspekty etyczne.

**Rozdział 4 (Wyniki — eksperyment crossover)** prezentuje wyniki badania eksperymentalnego: porównanie czasu segregacji, trafności klasyfikacji, kompletności raportów oraz oceny użyteczności (SUS) między metodą cyfrową a tradycyjną. Zawiera analizę statystyczną z weryfikacją hipotez H1-H4.

**Rozdział 5 (Wyniki — ankieta ogólnopolska)** prezentuje wyniki ankiety skierowanej do polskich ratowników medycznych: doświadczenie ze zdarzeniami masowymi, stosunek do narzędzi cyfrowych, gotowość do adopcji i zidentyfikowane bariery wdrożeniowe.

**Rozdział 6 (Dyskusja)** interpretuje uzyskane wyniki w kontekście istniejącej literatury, omawia implikacje praktyczne dla polskiego ratownictwa medycznego, analizuje ograniczenia badania oraz proponuje kierunki przyszłych badań.

**Rozdział 7 (Wnioski)** formułuje syntetyczne wnioski wynikające z przeprowadzonego badania.

**Rozdział 8 (Streszczenie)** zawiera streszczenie pracy w języku polskim i angielskim.

Pracę uzupełniają: bibliografia, spis tabel, spis rycin oraz aneksy (kwestionariusz ankiety, scenariusze symulacyjne, formularz zgody, kwestionariusz SUS w wersji polskiej).

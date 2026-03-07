# Rozdział 6. Dyskusja

## 6.1. Omówienie wyników eksperymentu

### 6.1.1. Czas segregacji (H1)

Wyniki eksperymentu crossover [wykazały / nie wykazały] istotnie krótszy czas segregacji przy zastosowaniu aplikacji TRIAGE MCI w porównaniu z metodą tradycyjną ([M_app] vs. [M_trad] sekund, p=[wartość], d=[wartość]).

Uzyskany wynik [jest spójny z / pozostaje w opozycji do] obserwacji prezentowanych w literaturze światowej. Kilka badań wykazało, że cyfrowe narzędzia wspomagające segregację mogą skracać czas procesu. Cone i in. [2011] zaobserwowali 18-procentową redukcję czasu segregacji przy zastosowaniu aplikacji mobilnej w porównaniu z metodą papierową, argumentując, że eliminacja konieczności przypominania sobie kolejności kroków algorytmu stanowi istotny mechanizm oszczędności czasowej. Podobnie Killeen i in. [2006] w badaniu z zastosowaniem urządzenia PDA do segregacji wykazali szybsze kompletowanie danych, choć w ich badaniu zysk czasowy wynikał głównie z automatyzacji dokumentacji, a nie samego procesu decyzyjnego.

[WARIANT — jeśli H1 potwierdzona:]
Mechanizm redukcji czasu w przypadku aplikacji TRIAGE MCI jest wielowymiarowy. Po pierwsze, aplikacja prowadzi użytkownika krok po kroku przez algorytm START, eliminując czas na refleksję nad prawidłową sekwencją kroków. Po drugie, klasyfikacja do kategorii T1-T4 odbywa się automatycznie po udzieleniu odpowiedzi, co eliminuje czas potrzebny na samodzielne przypisanie kategorii. Po trzecie — i być może najistotniejsze — raport generowany jest automatycznie, co eliminuje czasochłonny etap ręcznej dokumentacji. Ten ostatni mechanizm może tłumaczyć, dlaczego różnica czasowa jest szczególnie wyraźna, gdy jako punkt końcowy przyjmie się ukończenie pełnej procedury (segregacja + raport), a nie samą segregację.

[WARIANT — jeśli H1 niepotwierdzona:]
Brak istotnej różnicy czasowej może wynikać z kilku czynników. Po pierwsze, uczestnicy korzystali z aplikacji po raz pierwszy (po zaledwie 5-minutowej instrukcji), co mogło generować dodatkowy czas na zapoznanie się z interfejsem. Badania z zakresu ergonomii oprogramowania wskazują, że krzywa uczenia się nowych narzędzi może niwelować korzyści efektywnościowe w krótkim horyzoncie czasowym [Nielsen, 1993]. Po drugie, interakcja z ekranem dotykowym (nawigacja, zatwierdzanie odpowiedzi) wymaga czasu, którego nie ma przy mentalnym przetwarzaniu algorytmu przez doświadczonego ratownika. Możliwe, że powtórzony pomiar po okresie wprawy wykazałby istotną różnicę na korzyść aplikacji.

Warto zauważyć, że [ewentualna] analiza podgrup wykazała [większą / mniejszą] korzyść czasową z aplikacji wśród studentów niż wśród doświadczonych ratowników. Obserwacja ta jest spójna z hipotezą, że narzędzia cyfrowe mogą mieć największą wartość dodaną dla osób z mniejszym doświadczeniem, u których ryzyko błędu proceduralnego jest wyższe [Risavi i in., 2001].

### 6.1.2. Trafność klasyfikacji (H2)

[Wyższa / porównywalna] trafność klasyfikacji przy zastosowaniu aplikacji ([M_app]% vs. [M_trad]%, p=[wartość]) stanowi [istotne / niejednoznaczne] potwierdzenie potencjału cyfryzacji w zakresie redukcji błędów segregacyjnych.

W literaturze przedmiotu wskaźniki błędów segregacyjnych przy stosowaniu algorytmu START wahają się od 15 do 50% w zależności od badania, populacji i warunków [Jenkins i in., 2008; Bhalla i in., 2015]. Overtriage (nadklasyfikacja) dominuje nad undertriage (niedoklasyfikacją), co jest zjawiskiem częściowo pożądanym — lepiej jest przeszacować priorytet pacjenta niż go niedoszacować. American College of Surgeons przyjmuje za akceptowalny wskaźnik overtriage do 50% przy jednoczesnym utrzymaniu undertriage poniżej 5% [ACS, 2014].

[Jeśli aplikacja redukuje błędy:]
Redukcja błędów przy zastosowaniu aplikacji jest spójna z mechanizmem działania narzędzia — algorytm zaimplementowany w TRIAGE MCI jest realizowany zawsze w prawidłowej, pełnej sekwencji, co eliminuje dwie główne przyczyny błędów: pominięcie kroków procedury i subiektywną interpretację kryteriów [Kahn i in., 2009]. Warto podkreślić, że aplikacja nie eliminuje wszystkich źródeł błędów — ocena parametrów klinicznych (czy pacjent oddycha? czy CRT < 2 sekundy?) pozostaje w gestii użytkownika. Błędy obserwowane przy użyciu aplikacji wynikały zatem z [nieprawidłowej oceny parametrów klinicznych na kartach / ...], a nie z defektu algorytmu.

[Szczególnie istotny jest wynik w zakresie undertriage:]
[Jeśli dane to potwierdzą:] Redukcja undertriage z [X]% do [Y]% przy zastosowaniu aplikacji jest klinicznie istotna, ponieważ undertriage wiąże się bezpośrednio z ryzykiem opóźnienia pomocy pacjentom w stanie bezpośredniego zagrożenia życia. Nawet niewielka redukcja undertriage może przekładać się na poprawę rokowania pacjentów w warunkach realnego zdarzenia masowego [Frykberg, 2002].

### 6.1.3. Kompletność raportu (H3)

[Przewidywany wyraźny wynik na korzyść aplikacji w zakresie kompletności raportu stanowi / Wynik w zakresie kompletności raportu stanowi] być może najbardziej jednoznaczny argument za stosowaniem narzędzi cyfrowych w segregacji medycznej.

Raportowanie ze zdarzenia masowego jest elementem procesu segregacji, który w warunkach realnych jest szczególnie podatny na błędy i pominięcia. W atmosferze chaosu, stresu i wielozadaniowości ratownik musi jednocześnie zarządzać wieloma poszkodowanymi, koordynować zespoły i utrzymywać łączność — dokumentacja schodzi na dalszy plan [Auf der Heide, 2006]. Badania z zakresu zarządzania kryzysowego konsekwentnie wskazują na niedostateczną jakość dokumentacji ze zdarzeń masowych jako jeden z kluczowych problemów operacyjnych [Killeen i in., 2006].

Aplikacja TRIAGE MCI automatycznie rejestruje dane każdego pacjenta (kategorię segregacyjną, ścieżkę decyzyjną, czas, przypisany zespół ZRM, szpital docelowy, status transportu) i generuje kompletny raport jednym kliknięciem. Ta automatyzacja eliminuje konieczność ręcznego zestawiania informacji z pamięci, notatek i kart, co tłumaczy [przewidywaną] istotną przewagę nad metodą tradycyjną.

Warto podkreślić, że elementy listy kontrolnej takie jak „czas rozpoczęcia działań" i „nazwa/opis zdarzenia" wymagały ręcznego wprowadzenia również w aplikacji — różnica między metodami dotyczyła przede wszystkim elementów automatycznie generowanych (lista pacjentów z kategoriami, podział T1-T4, statusy transportów, przypisania zespołów).

### 6.1.4. Użyteczność — wynik SUS (H4)

Wynik SUS wynoszący [M_SUS] punktów plasuje aplikację TRIAGE MCI w kategorii [dobrej / wysokiej] użyteczności według klasyfikacji Bangora i in. [2009]. Wynik ten jest [porównywalny z / wyższy niż] wyniki raportowane dla innych aplikacji medycznych: średni wynik SUS dla aplikacji mobilnych w ochronie zdrowia mieści się w przedziale 65-78 punktów [Zapata i in., 2015], a dla aplikacji ratownictwa medycznego — 60-75 punktów [Wallis i in., 2017].

[Wysoki wynik SUS jest szczególnie obiecujący, biorąc pod uwagę, że:] uczestnicy korzystali z aplikacji po raz pierwszy, po zaledwie 5-minutowym instruktażu. Sugeruje to, że interfejs aplikacji jest na tyle intuicyjny, że minimalizuje barierę wejścia — cecha kluczowa dla narzędzia przeznaczonego do użycia w warunkach stresu i deficytu czasu.

[Ewentualna różnica w ocenie SUS między studentami a ratownikami] może wskazywać na [np. większą otwartość młodszego pokolenia na technologie mobilne / wyższe oczekiwania doświadczonych ratowników wobec narzędzi pracy].

## 6.2. Omówienie wyników ankiety

### 6.2.1. Doświadczenie ze zdarzeniami masowymi

Wyniki ankiety ogólnopolskiej ujawniają [niepokojący / zróżnicowany] obraz przygotowania polskich ratowników medycznych do działania w zdarzeniach masowych. Fakt, że [jedynie X]% respondentów miało doświadczenie z realnym zdarzeniem masowym, potwierdza obserwacje z literatury — zdarzenia masowe, choć spektakularne, są w praktyce ratowniczej stosunkowo rzadkie, co utrudnia budowanie kompetencji opartych na doświadczeniu [Koenig i Schultz, 2010]. Jednocześnie średnia samoocena przygotowania na poziomie [M]/5 wskazuje na [subiektywne poczucie niedostatecznego przygotowania / umiarkowaną pewność siebie w tym zakresie].

Te ustalenia wzmacniają argument za stosowaniem narzędzi cyfrowych: skoro doświadczenie z realnymi zdarzeniami masowymi jest z natury ograniczone, a szkolenia — choć wartościowe — nie zawsze prowadzą do trwałego utrwalenia algorytmu, narzędzie cyfrowe prowadzące użytkownika krok po kroku przez algorytm START może stanowić cenną „siatkę bezpieczeństwa" (*safety net*) kompensującą braki w doświadczeniu i pamięci proceduralnej.

### 6.2.2. Gotowość do adopcji i bariery

[X]% respondentów wyraziło gotowość do korzystania z aplikacji do segregacji na realnym zdarzeniu, co wskazuje na [pozytywny / umiarkowanie pozytywny / zróżnicowany] stosunek środowiska ratowniczego do cyfryzacji tego procesu.

Analiza barier ujawnia interesujące rozbieżności między obawami respondentów a rzeczywistymi właściwościami nowoczesnych rozwiązań technologicznych. [Jeśli dane to potwierdzą:] Bariera „problemy z łącznością / internetem" jest jedną z najczęściej wskazywanych, a jednocześnie jedną z najbardziej niezasadnych w kontekście aplikacji klasy PWA, które z definicji działają w pełni offline po jednokrotnym pobraniu. Ta obserwacja sugeruje, że edukacja na temat możliwości współczesnych technologii mobilnych jest równie ważna jak rozwój samych narzędzi.

Bariera „brak oficjalnego zatwierdzenia / rekomendacji" wskazuje na kluczową rolę autorytetu instytucjonalnego w procesie wdrażania innowacji w ratownictwie medycznym. Ratownicy medyczni pracują w systemie silnie zhierarchizowanym i zregulowanym, w którym odejście od ustalonych procedur — nawet w celu ich usprawnienia — wymaga sankcji formalnej. Ten wynik ma istotne implikacje praktyczne: nawet doskonałe technologicznie narzędzie cyfrowe może nie zyskać akceptacji bez rekomendacji Konsultanta Krajowego w dziedzinie medycyny ratunkowej lub innego autorytetu instytucjonalnego.

### 6.2.3. Profil respondenta otwartego na cyfryzację

[Na podstawie danych — omówienie czynników predykcyjnych pozytywnego stosunku do cyfryzacji, np.:]

Zidentyfikowane zależności między gotowością do adopcji a zmiennymi demograficzno-zawodowymi wskazują, że [wcześniejsze korzystanie z aplikacji medycznych / młodszy wiek / staż pracy X-Y lat] jest najsilniejszym predyktorem pozytywnego stosunku do cyfryzacji segregacji. Wynik ten jest spójny z teorią dyfuzji innowacji Rogersa [2003], w której wczesni adoptorzy (*early adopters*) charakteryzują się [otwartością na nowości, wyższym poziomem wykształcenia technicznego i gotowością do eksperymentowania].

## 6.3. Integracja wyników obu części badania

Zestawienie wyników eksperymentu crossover z wynikami ankiety ogólnopolskiej pozwala na sformułowanie zintegrowanej oceny potencjału cyfryzacji segregacji medycznej w Polsce.

Z jednej strony, dane eksperymentalne [wskazują na mierzalne korzyści / sugerują potencjalne korzyści] w zakresie [czasu, trafności, kompletności raportu] — są to obiektywne wskaźniki efektywności, trudne do zakwestionowania. Z drugiej strony, dane ankietowe ujawniają, że nawet przy istniejących korzyściach technicznych, wdrożenie wymaga przezwyciężenia barier organizacyjnych, kulturowych i percepcyjnych.

Ta rozbieżność między obiektywnymi korzyściami a subiektywnym postrzeganiem technologii nie jest zjawiskiem nowym w kontekście wdrażania innowacji w ochronie zdrowia. Podobne wzorce opisywano przy wdrażaniu elektronicznej dokumentacji medycznej [Boonstra i Broekhuis, 2010], telemedycyny [Kruse i in., 2018] i systemów wspomagania decyzji klinicznych [Kawamoto i in., 2005].

Kluczowym wnioskiem integracyjnym jest to, że [skuteczne wdrożenie narzędzi cyfrowych w segregacji medycznej wymaga strategii wielowymiarowej, łączącej rozwój technologiczny z edukacją, szkoleniem i wsparciem instytucjonalnym].

## 6.4. Implikacje praktyczne

Na podstawie uzyskanych wyników można sformułować następujące implikacje praktyczne dla polskiego ratownictwa medycznego:

### 6.4.1. Implikacje dla praktyki operacyjnej

1. **Narzędzia cyfrowe jako uzupełnienie, nie zastąpienie.** Wyniki badania [sugerują / potwierdzają], że aplikacje do segregacji medycznej mogą stanowić wartościowe uzupełnienie tradycyjnych metod, ale nie powinny być traktowane jako ich substytut. Ratownik medyczny musi znać algorytm START niezależnie od dostępności narzędzi cyfrowych — aplikacja jest wsparciem, a nie protezą.

2. **Szczególna wartość w raportowaniu.** [Jeśli H3 potwierdzona:] Największa wartość dodana aplikacji dotyczy automatyzacji dokumentacji i raportowania — elementu, który w warunkach realnych jest najczęściej zaniedbywany. Wdrożenie narzędzi cyfrowych mogłoby istotnie poprawić jakość dokumentacji ze zdarzeń masowych.

3. **Wartość dla mniej doświadczonych ratowników.** [Jeśli analiza podgrup to potwierdza:] Aplikacja może mieć szczególną wartość dla ratowników z mniejszym doświadczeniem, u których ryzyko błędów proceduralnych jest wyższe.

### 6.4.2. Implikacje dla szkolenia

1. **Włączenie narzędzi cyfrowych do programu szkoleń.** Wyniki sugerują, że szkolenia z zakresu segregacji medycznej mogłyby być wzbogacone o moduł zapoznający ratowników z dostępnymi narzędziami cyfrowymi.

2. **Aplikacja jako narzędzie dydaktyczne.** Interaktywny charakter aplikacji TRIAGE MCI (krok po kroku przez algorytm) sprawia, że może ona pełnić funkcję narzędzia szkoleniowego, pomagającego w utrwaleniu algorytmu START.

### 6.4.3. Implikacje dla polityki zdrowotnej

1. **Potrzeba rekomendacji instytucjonalnej.** Wyniki ankiety wskazują, że oficjalna rekomendacja ze strony Konsultanta Krajowego w dziedzinie medycyny ratunkowej lub Komendy Głównej PSP byłaby kluczowym czynnikiem sprzyjającym adopcji narzędzi cyfrowych.

2. **Standaryzacja wymagań.** Rozwój rynku aplikacji medycznych dla ratownictwa wymaga opracowania kryteriów oceny takich narzędzi (funkcjonalność, bezpieczeństwo danych, praca offline, dostępność językowa).

## 6.5. Ograniczenia badania

Niniejsze badanie obarczone jest ograniczeniami, które należy uwzględnić przy interpretacji wyników:

### 6.5.1. Ograniczenia eksperymentu

1. **Warunki symulacyjne vs. realne zdarzenie.** Segregacja przeprowadzana na kartach pacjentów w kontrolowanych warunkach nie odwzorowuje w pełni złożoności realnego zdarzenia masowego: stresu, hałasu, widoku poszkodowanych, konieczności wielozadaniowości, warunków atmosferycznych i ograniczeń czasowych wynikających z dynamiki zdarzenia. Transferowalność wyników do warunków polowych pozostaje otwartym pytaniem wymagającym weryfikacji w bardziej realistycznych symulacjach (np. z użyciem aktorów lub manekinów).

2. **Wielkość próby.** Próba n=[N] uczestników, choć wystarczająca dla schematu crossover (w którym każdy uczestnik jest swoją kontrolą), ogranicza moc statystyczną analiz podgrup i generalizowalność wyników.

3. **Efekt nowości.** Uczestnicy korzystali z aplikacji po raz pierwszy, co mogło zarówno spowalniać ich (krzywa uczenia się), jak i przyspieszać (motywacja wynikająca z nowości — efekt Hawthorne'a). Badanie podłużne, uwzględniające wielokrotne użycie aplikacji, pozwoliłoby na ocenę stabilnych efektów po okresie wprawy.

4. **Brak zaślepienia.** Charakter interwencji (smartfon vs. kartka) uniemożliwiał zaślepienie uczestników i obserwatorów, co mogło wprowadzać stronniczość w pomiarze czasu i ocenie kompletności raportu.

5. **Specyficzność narzędzia.** Wyniki dotyczą konkretnej aplikacji (TRIAGE MCI) i nie mogą być automatycznie uogólniane na inne rozwiązania cyfrowe do segregacji medycznej.

### 6.5.2. Ograniczenia ankiety

1. **Dobór próby.** Metoda kuli śnieżnej i dystrybucja przez media społecznościowe mogą prowadzić do nadreprezentacji osób aktywnych online i potencjalnie bardziej otwartych na nowe technologie. Wyniki mogą zatem zawyżać pozytywny stosunek do cyfryzacji w porównaniu z ogółem populacji ratowników medycznych w Polsce.

2. **Samoocena.** Deklaracje respondentów (np. dotyczące gotowości do korzystania z aplikacji) mogą nie przekładać się na rzeczywiste zachowania w warunkach operacyjnych.

3. **Przekrój czasowy.** Badanie przekrojowe uchwytuje postawy w jednym punkcie czasowym — mogą one ulegać zmianom wraz z rozwojem technologii i kolejnymi doświadczeniami.

## 6.6. Kierunki przyszłych badań

Na podstawie uzyskanych wyników i zidentyfikowanych ograniczeń można zaproponować następujące kierunki przyszłych badań:

1. **Badanie w warunkach wysokiej wierności symulacji.** Przeprowadzenie analogicznego eksperymentu z użyciem aktorów (standardized patients) lub zaawansowanych manekinów w warunkach zbliżonych do realnego zdarzenia masowego (hałas, symulacja stresu, ograniczone zasoby).

2. **Badanie podłużne.** Ocena efektywności aplikacji po okresie regularnego korzystania (np. 3-6 miesięcy), uwzględniająca krzywą uczenia się i efekt wprawy.

3. **Badanie wieloośrodkowe.** Replikacja badania na większej, reprezentatywnej próbie ratowników medycznych z różnych regionów Polski.

4. **Integracja z systemami dyspozytorskimi.** Badanie możliwości integracji danych z aplikacji do segregacji z systemami zarządzania dyspozytorskiego (np. SWD PRM), co mogłoby usprawnić przepływ informacji między miejscem zdarzenia a centrum powiadamiania ratunkowego.

5. **Porównanie algorytmów.** Rozszerzenie aplikacji o inne algorytmy segregacyjne (SALT, Sieve/Sort) i porównanie ich efektywności w wersji cyfrowej.

6. **Badanie w kontekście międzynarodowym.** Wielojęzyczność aplikacji TRIAGE MCI (7 języków) umożliwia przeprowadzenie analogicznych badań porównawczych w innych krajach europejskich.

# Návod k obsluze TRIAGE MCI

## 1. Co je TRIAGE MCI

TRIAGE MCI je webová aplikace (PWA) pro třídění a řízení hromadného neštěstí (MCI) založená na algoritmu START. Umožňuje:
- Klasifikaci pacientů do 4 kategorií (T1–T4)
- Řízení transportu do nemocnic
- Sledování kapacity nemocnic
- Generování a odesílání hlášení
- Import dat od dispečera

Aplikace funguje plně offline. Veškerá data jsou uchovávána výhradně na zařízení uživatele — žádné informace nejsou odesílány na externí servery.

Podporuje 7 jazyků: polštinu, angličtinu, italštinu, francouzštinu, němčinu, češtinu, portugalštinu.

---

## 2. Zahájení události

### Úvodní obrazovka

Po spuštění aplikace se zobrazí obrazovka konfigurace události:

1. **Volba jazyka** — výběr v horní části obrazovky. Změna jazyka okamžitě přeloží celé rozhraní.
2. **Název události** — povinné pole. Zadejte lokalizaci a typ události (např. „Nehoda D1 km 312"). Zobrazí se v hlášeních.
3. **KAM** — volitelné pole. Vedoucí zdravotnické akce — jméno a příjmení osoby koordinující činnost na místě.
4. **Zobrazit nápovědu** — přepínač zapínající bubliny s kontextovými tipy v celé aplikaci. Užitečné při prvním použití.
5. **ZAHÁJIT TRIAGE** — zahájí událost a spustí časomíru.

### Obnovení události

Pokud je v paměti zařízení uložena předchozí událost, při spuštění se zobrazí volba:
- **POKRAČOVAT** — obnoví uloženou událost se všemi daty
- **NOVÁ UDÁLOST** — vymaže předchozí data a začne znovu

---

## 3. Třídění (algoritmus START)

### Navigace

Aplikace má 3 záložky v horní části obrazovky:
- **TŘÍDĚNÍ** — klasifikace pacientů
- **DISPONOVÁNÍ** — transport a řízení
- **HLÁŠENÍ** — souhrn a export

Statistický panel pod navigací zobrazuje počet pacientů v každé kategorii: 🔴 T1, 🟡 T2, 🟢 T3, ⚫ T4 a celkem.

### Průvodce tříděním START

Průvodce vás provede 6krokovým algoritmem START. Na každém kroku odpovídáte ANO nebo NE:

**Krok 1 — Chodí pacient?**
- ANO → T3 Zelený (odložený)
- NE → přejděte ke kroku 2

**Krok 2 — Dýchá spontánně?**
- ANO → přejděte ke kroku 4 (hodnocení dechové frekvence)
- NE → přejděte ke kroku 3

**Krok 3 — Zprůchodněte dýchací cesty. Dýchá po zprůchodnění?**
- ANO → T1 Červený (neodkladný)
- NE → T4 Černý (úmrtí)

**Krok 4 — Je dechová frekvence v normě (≤ 30/min)?**
- ANO → přejděte ke kroku 5
- NE → T1 Červený (neodkladný)

**Krok 5 — Je kapilární návrat v normě (≤ 2 s) a pulz na radiální tepně přítomen?**
- ANO → přejděte ke kroku 6
- NE → T1 Červený (neodkladný)

**Krok 6 — Splní jednoduché pokyny?**
- ANO → T2 Žlutý (naléhavý)
- NE → T1 Červený (neodkladný)

Ke každé otázce je nápověda (hint) zobrazená pod otázkou — popisuje, jak daný parametr posoudit.

### Karta výsledku

Po dokončení třídění se zobrazí karta výsledku s:

- **Barevný štítek** — velký čtverec s kategorií (T1/T2/T3/T4)
- **Rozhodovací cesta** — zobrazuje průběh algoritmu
- **Pohlaví** — tlačítka M (muž), Ž (žena), ? (neznámé)
- **Věk** — číselné pole s tlačítky -5, -1, +1, +5 (rozsah 0–150)
- **Poranění těla** — tlačítko otevírající diagram poranění (viz oddíl 4)
- **Poznámky** — pole pro doplňující informace (poranění, mechanismus úrazu atd.)

Dostupné akce:
- **✓ POTVRDIT PACIENTA** — přidá pacienta do seznamu
- **↺ ZNOVU** — resetuje průvodce na krok 1
- **ZMĚNIT BARVU RUČNĚ** — otevře řadu 4 tlačítek (T1–T4) pro ruční změnu kategorie

### Seznam pacientů na místě

Pod průvodcem se zobrazuje seznam „Pacienti na místě" řazený podle priority (T1 → T2 → T3 → T4).

Každá karta pacienta obsahuje:
- Barevný štítek (např. P-001)
- Kategorii, pohlaví/věk, čas registrace
- Pole poznámek (editovatelné)
- Tlačítko poranění (🩻) se souhrnem
- 4 barevné tečky re-triáže (T1–T4) — klikněte pro změnu kategorie
- Tlačítko smazání (✕)
- Značku re-triáže (↺n), pokud byla kategorie měněna

### Seznam v transportu

Sekce „V transportu" se zobrazí, když jsou pacienti na cestě do nemocnice. Každá karta obsahuje:
- Štítek pacienta s barvou kategorie
- Trasu: Tým → Nemocnice → čas odjezdu
- Tlačítko „změnit" — změna cílové nemocnice
- Tlačítko „✓ Doručen" — označí doručení a uvolní tým ZZS

---

## 4. Diagram poranění těla

Tlačítko „🩻 Poranění těla" otevírá interaktivní diagram těla.

### Zobrazení diagramu

Zobrazeny jsou dva obrysy těla vedle sebe:
- **PŘEDNÍ** (front) — hlava, hrudník, břicho, paže, nohy
- **ZADNÍ** (back) — hlava, horní část zad, dolní část zad, paže, nohy

Po stranách jsou viditelná označení stran: **L** (levá) a **P** (pravá), obrácená pro zadní pohled.

Klikněte na zónu těla pro její výběr.

### Panel poranění zóny

Po výběru zóny se otevře panel se 7 typy poranění k zaškrtnutí:
- Zlomenina
- Krvácení
- Popálenina
- Rána
- Otok
- Bolest
- Amputace

Zaškrtnuté typy jsou označeny ✓. Každý lze zapnout/vypnout kliknutím.

Ve spodní části panelu:
- **✕ NE** — zahodí změny v této zóně
- **✓ ANO** — potvrdí poranění pro vybranou zónu

### Potvrzení diagramu

Po zaznačení poranění ve všech potřebných zónách:
- **✓ ANO** — potvrdí celý diagram poranění
- **✕ NE** — zahodí všechny změny a vrátí stav před otevřením

Zaznačená poranění jsou viditelná:
- Na kartě pacienta jako zkrácený souhrn (např. „🩻 Hlava: Zlomenina, Krvácení")
- V hlášení u daného pacienta

---

## 5. Disponování

### Modální okna konfigurace

Při prvním vstupu do záložky Disponování se postupně zobrazí 3 modální okna:

**1. Vedoucí zdravotnické služby (GDM)**
- Pole pro jméno a příjmení GDM
- Tlačítko importu dat od dispečera (viz oddíl 6)
- PŘESKOČIT nebo ULOŽIT

**2. Týmy ZZS**
- Otázka: „Znáte týmy ZZS vyslané na místo?"
- ANO → formulář pro přidávání týmů (název + tlačítko přidat)
- NE → použijí se výchozí týmy (S-01, S-02, P-01, P-02, P-03, LPR)
- POTVRDIT nebo PŘESKOČIT (VÝCHOZÍ)

**3. Cílové nemocnice**
- Formulář pro přidávání nemocnic s kapacitou:
  - Název nemocnice
  - 🔴 Červená kapacita (T1) — počet lůžek
  - 🟡 Žlutá kapacita (T2) — počet lůžek
- POTVRDIT nebo PŘESKOČIT (VÝCHOZÍ)

Každé modální okno se zobrazí pouze jednou za událost.

### Formulář disponování

Sekce „Disponovat transport" obsahuje:

1. **Pacient** — rozbalovací seznam pacientů na místě, řazený podle priority
2. **Tým ZZS** — seznam dostupných týmů (obsazené týmy jsou skryté)
3. **Nemocnice** — seznam nemocnic s informací o obsazenosti:
   - Formát: „Název nemocnice [🔴 obsazeno/celkem 🟡 obsazeno/celkem]"
   - Nemocnice bez volných míst v kategorii pacienta jsou označeny varováním
4. **DISPONOVAT TRANSPORT** — tlačítko aktivní po vyplnění všech polí

### Panely správy

Pod formulářem se nacházejí dva panely (na širokých obrazovkách vedle sebe, na úzkých pod sebou):

**Nemocnice** (lůžka CELKEM):
- Seznam nemocnic s editovatelnými poli kapacity (🔴 a 🟡)
- Přidávání nových nemocnic
- Mazání nemocnic (✕)
- Sbalení/rozbalení seznamu

**Týmy ZZS:**
- Seznam týmů s tlačítky pro smazání
- Přidávání nových týmů
- Sbalení/rozbalení seznamu

### Historie transportů

Ve spodní části sekce disponování (ve výchozím stavu sbalená). Rozbalte pro zobrazení chronologického seznamu všech transportů s:
- Štítek pacienta s barvou
- Tým → Nemocnice → čas odjezdu
- Stav (v transportu / doručen)

---

## 6. Import dat od dispečera

Funkce umožňuje dispečerovi odeslat data (název události, GDM, týmy ZZS, nemocnice s kapacitou) záchranáři v terénu.

### Metoda 1 — SMS odkaz (jedno kliknutí)

Dispečer odešle SMS s odkazem ve formátu:
```
https://0xjaqbek.github.io/triage/?i=DATA_BASE64
```

Po kliknutí na odkaz:
1. Aplikace se otevře z cache (nevyžaduje internet, pokud je nainstalována)
2. Zobrazí se náhled dat: název události, GDM, týmy, nemocnice
3. **IMPORTOVAT** — importuje data a otevře třídění
4. **ZAMÍTNOUT** — ignoruje data

### Metoda 2 — Vložení dat

Pokud odkaz neotevře aplikaci (zablokovaná zařízení):
1. Otevřete aplikaci, přejděte do Disponování
2. V modálním okně GDM klikněte na **📥 IMPORTOVAT DATA OD DISPEČERA**
3. Vložte obsah SMS (celý odkaz nebo samotný kód Base64)
4. Náhled dat se zobrazí automaticky
5. **IMPORTOVAT** — potvrdí data

### Import v průběhu události

Pokud je událost již zahájena a máte vytříděné pacienty:
- **Pacienti** — zachováni (nikdy nejsou mazáni)
- **KAM** — zachován
- **GDM** — nahrazen daty od dispečera
- **Týmy ZZS** — nahrazeny daty od dispečera
- **Nemocnice** — nahrazeny daty od dispečera
- **Název události** — sloučen: „Váš název | Název dispečera (dispečer)"

Žluté varování v náhledu informuje o zachování pacientů a KAM.

---

## 7. Režim dispečera

Dostupný přes odkaz „Režim dispečera →" na úvodní obrazovce nebo na adrese `/dyspozytor/`.

Stránka umožňuje dispečerovi připravit a odeslat data pro tým v terénu.

### Formulář

1. **Název události** — např. „MCI Praha-Háje"
2. **Hlavní dispečer zdravotnické služby (GDM)** — jméno a příjmení
3. **Týmy ZZS** — přidávejte jednotlivě (název + tlačítko „+")
4. **Cílové nemocnice** — název + kapacita 🔴 a 🟡 (+ tlačítko přidat)

### Odeslání

- **📋 Náhled dat** — zobrazí souhrn s velikostí ve znacích a počtem SMS
- **📱 Odeslat SMS** — otevře SMS aplikaci s připraveným odkazem
- **📋 Kopírovat data** — zkopíruje odkaz do schránky

Data jsou kódována v Base64 (kompatibilita s SMS GSM-7 — až 1530 znaků v 10 segmentech SMS). Pojme přibližně 20 nemocnic a 20 týmů.

---

## 8. Hlášení

### Vizuální hlášení

Sekce „Hlášení pro dispečera" se generuje automaticky a obsahuje:

**Záhlaví:**
- Název události, čas zahájení, čas hlášení
- KAM a GDM (pokud jsou zadány)
- Čas ukončení (pokud je událost ukončena)

**Souhrn:**
- Barevné štítky s počtem pacientů v každé kategorii
- Celkový počet pacientů

**Nemocnice** (samostatný blok pro každou nemocnici s transporty):
- Název nemocnice, počet pacientů, informace o kapacitě
- Seznam pacientů s: štítek, věk/pohlaví, tým, čas, stav (V TRANSPORTU / DORUČEN)
- Historie re-triáže (např. ↺T1→T2→T1)
- Historie změn nemocnice (např. ⇄Nemocnice1→Nemocnice2)
- Souhrn poranění

**Pacienti na místě** (pokud někdo zbývá):
- Seznam s: štítek, věk/pohlaví, kategorie, poznámky, poranění

### Textové hlášení

Tlačítko „📋 Textové hlášení" otevře modální okno s úplným hlášením v textovém formátu. Obsahuje kompletní seznam pacientů s poznámkami, pohlavím, věkem a cestou třídění.

### Odeslání hlášení

- **📋 Kopírovat do schránky** — zkopíruje úplný text hlášení
- **📤 Odeslat** — otevře možnosti:
  - **SMS** — otevře nativní SMS aplikaci s obsahem hlášení
  - **Email** — otevře emailového klienta s předmětem a obsahem hlášení

### Ukončení události

Tlačítko **UKONČIT UDÁLOST** v záhlaví hlášení:
- Uloží čas ukončení
- Změní se na **ZRUŠIT UKONČENÍ** (vratné)
- Čas ukončení se zobrazí v hlášení

---

## 9. Re-triáž a změna nemocnice

### Re-triáž

Na kartě pacienta v sekci „Pacienti na místě" se nacházejí 4 barevné tečky (T1–T4). Kliknutí na jinou kategorii než aktuální:
1. Otevře potvrzovací okno: „Změnit kategorii pacienta P-001 z T1 na T2?"
2. **POTVRDIT** — změní kategorii, uloží do historie
3. **ZRUŠIT** — bez změn

Historie re-triáže:
- Značka ↺n na kartě pacienta (n = počet změn)
- Úplná historie v hlášení (např. T1→T2→T1 s časy)

### Změna nemocnice v transportu

Na kartě pacienta „V transportu" tlačítko „změnit":
1. Otevře modální okno s aktuální nemocnicí
2. Rozbalovací seznam s dostupnými nemocnicemi (s informací o kapacitě)
3. **POTVRDIT** — změní nemocnici, uloží do historie
4. **ZRUŠIT** — bez změn

Historie změn nemocnice je viditelná v hlášení (⇄Nemocnice1→Nemocnice2).

---

## 10. Instalace a offline režim

### Instalace jako PWA

**Android (Chrome):**
1. Otevřete aplikaci v Chrome
2. Klepněte na ⋮ (menu) v pravém horním rohu
3. Zvolte „Přidat na domovskou obrazovku"
4. Klepněte na „Nainstalovat"

**iOS (Safari):**
1. Otevřete aplikaci v Safari
2. Klepněte na ⎙ (Sdílet)
3. Posuňte dolů a zvolte „Přidat na plochu"

Po instalaci aplikace funguje jako samostatná aplikace s plným offline režimem.

### Offline režim

Aplikace se po prvním načtení uloží do cache zařízení. Všechny funkce fungují bez připojení k internetu:
- Třídění pacientů
- Disponování transportu
- Generování hlášení
- Import dat (vložení)

Internet je potřeba pouze pro: první stažení, aktualizaci a odesílání SMS/emailů (přes nativní aplikace zařízení).

### Stažení jako HTML soubor

Tlačítko **⬇ Stáhnout** na úvodní obrazovce stáhne aplikaci jako samostatný HTML soubor. Lze jej přenést na flash disk a spustit na libovolném počítači s prohlížečem.

---

## 11. Soukromí

- Veškerá data jsou uchovávána **výhradně na zařízení** (localStorage)
- Žádná data **nejsou odesílána** na externí servery
- Žádná analytika, sledování, telemetrie
- Žádné sledovací cookies
- Funkce SMS a email využívají nativní aplikace zařízení — aplikace neodesílá zprávy přímo
- Smazání dat: použijte „Ukončit událost" nebo vymažte data prohlížeče

Úplné zásady ochrany soukromí jsou dostupné na adrese `/privacy/`.

# Help Bubbles Implementation Plan

> **For Claude:** REQUIRED SUB-SKILL: Use superpowers:executing-plans to implement this plan task-by-task.

**Goal:** Add a toggle that shows/hides 15 chat-bubble-style help overlays across all screens of the TRIAGE MCI app.

**Architecture:** Pure CSS+JS, no dependencies. A `body.help-on` class controls visibility. Bubbles are absolutely-positioned `<div class="help-bubble">` elements with `✕` dismiss. Toggle state persisted in localStorage; dismissed bubbles tracked in-memory only (reset on toggle OFF→ON or new session).

**Tech Stack:** Vanilla HTML/CSS/JS in a single file (`index.html`). Existing i18n system (`TRANSLATIONS` object + `data-i18n` attributes + `t()` function).

**Design doc:** `docs/plans/2026-03-12-help-bubbles-design.md`

---

### Task 1: Add CSS for help bubbles and toggle switch

**Files:**
- Modify: `index.html` — CSS section (after line ~554, after `.btn-start:active`)

**Step 1: Add the help bubble CSS**

Insert after the `.btn-start:active { transform: scale(0.97); }` rule (line 554):

```css
/* HELP BUBBLES */
.help-bubble {
  display: none;
  position: absolute;
  background: var(--surface);
  border: 1px solid var(--border2);
  border-radius: 10px;
  padding: 10px 28px 10px 10px;
  font-size: 12px;
  line-height: 1.4;
  color: var(--text2);
  max-width: 280px;
  z-index: 50;
  box-shadow: 0 4px 20px rgba(0,0,0,0.4);
  animation: helpFadeIn 0.2s ease;
}

body.help-on .help-bubble:not(.dismissed) {
  display: block;
}

.help-bubble .help-close {
  position: absolute;
  top: 4px;
  right: 6px;
  background: none;
  border: none;
  color: var(--text3);
  font-size: 14px;
  cursor: pointer;
  padding: 2px 4px;
  line-height: 1;
}

.help-bubble .help-close:hover {
  color: var(--text);
}

/* Arrow pointing up */
.help-bubble[data-arrow="up"]::before {
  content: '';
  position: absolute;
  top: -6px;
  left: 20px;
  width: 10px;
  height: 10px;
  background: var(--surface);
  border-left: 1px solid var(--border2);
  border-top: 1px solid var(--border2);
  transform: rotate(45deg);
}

/* Arrow pointing down */
.help-bubble[data-arrow="down"]::before {
  content: '';
  position: absolute;
  bottom: -6px;
  left: 20px;
  width: 10px;
  height: 10px;
  background: var(--surface);
  border-right: 1px solid var(--border2);
  border-bottom: 1px solid var(--border2);
  transform: rotate(45deg);
}

/* Arrow pointing left */
.help-bubble[data-arrow="left"]::before {
  content: '';
  position: absolute;
  left: -6px;
  top: 12px;
  width: 10px;
  height: 10px;
  background: var(--surface);
  border-left: 1px solid var(--border2);
  border-bottom: 1px solid var(--border2);
  transform: rotate(45deg);
}

@keyframes helpFadeIn {
  from { opacity: 0; transform: translateY(4px); }
  to { opacity: 1; transform: translateY(0); }
}

/* Responsive: full-width on small screens */
@media (max-width: 400px) {
  .help-bubble {
    max-width: calc(100vw - 24px);
    left: 12px !important;
    right: 12px !important;
  }
  .help-bubble::before { display: none; }
}

/* HELP TOGGLE SWITCH */
.help-toggle {
  display: flex;
  align-items: center;
  gap: 8px;
  cursor: pointer;
  font-size: 12px;
  color: var(--text3);
  user-select: none;
}

.help-toggle-track {
  width: 36px;
  height: 20px;
  background: var(--border);
  border-radius: 10px;
  position: relative;
  transition: background 0.2s;
}

.help-toggle input { display: none; }

.help-toggle input:checked + .help-toggle-track {
  background: var(--green);
}

.help-toggle-track::after {
  content: '';
  position: absolute;
  top: 2px;
  left: 2px;
  width: 16px;
  height: 16px;
  background: var(--white);
  border-radius: 50%;
  transition: transform 0.2s;
}

.help-toggle input:checked + .help-toggle-track::after {
  transform: translateX(16px);
}
```

**Step 2: Verify CSS renders correctly**

Open the app in browser, inspect elements to confirm no style conflicts.

**Step 3: Commit**

```bash
git add index.html
git commit -m "style: add CSS for help bubbles and toggle switch"
```

---

### Task 2: Add help bubble toggle HTML to setup screen and header

**Files:**
- Modify: `index.html` — HTML section

**Step 1: Add toggle to setup screen**

Insert after the `btn-start` button (line 1183), before the `setup-footer` div (line 1184):

```html
  <label class="help-toggle" style="margin-top:12px;">
    <input type="checkbox" id="helpToggleSetup" onchange="toggleHelp(this.checked)">
    <span class="help-toggle-track"></span>
    <span data-i18n="help_toggle_label">Pokaż podpowiedzi</span>
  </label>
```

**Step 2: Add toggle to header**

Insert inside the `header-top` div (after line 1248, after the `.nav` div closing tag):

```html
      <label class="help-toggle" style="margin-left:auto;">
        <input type="checkbox" id="helpToggleHeader" onchange="toggleHelp(this.checked)">
        <span class="help-toggle-track"></span>
        <span data-i18n="help_toggle_label">Pokaż podpowiedzi</span>
      </label>
```

**Step 3: Verify both toggles render**

Open the app. Confirm toggle appears below "ROZPOCZNIJ TRIAGE" on setup. Start event, confirm toggle appears in header.

**Step 4: Commit**

```bash
git add index.html
git commit -m "feat: add help toggle to setup screen and header"
```

---

### Task 3: Add help bubble JS logic (toggle, dismiss, persistence)

**Files:**
- Modify: `index.html` — JS section (after `applyStaticTranslations` function, around line 2259)

**Step 1: Add the JS logic**

Insert after line 2259 (after `applyStaticTranslations` closing brace):

```javascript
// ===== HELP BUBBLES =====
const dismissedBubbles = new Set();

function toggleHelp(on) {
  document.body.classList.toggle('help-on', on);
  document.getElementById('helpToggleSetup').checked = on;
  document.getElementById('helpToggleHeader').checked = on;
  if (on) {
    dismissedBubbles.clear();
    document.querySelectorAll('.help-bubble.dismissed').forEach(el => el.classList.remove('dismissed'));
  }
  localStorage.setItem('triageHelpOn', on ? '1' : '0');
}

function dismissBubble(id) {
  dismissedBubbles.add(id);
  const el = document.querySelector(`.help-bubble[data-help-id="${id}"]`);
  if (el) el.classList.add('dismissed');
}

function initHelp() {
  const on = localStorage.getItem('triageHelpOn') === '1';
  document.getElementById('helpToggleSetup').checked = on;
  document.getElementById('helpToggleHeader').checked = on;
  if (on) document.body.classList.add('help-on');
}
```

**Step 2: Call `initHelp()` on page load**

Find the existing DOMContentLoaded or init block. Add `initHelp();` call alongside existing initialization code. If the app uses an inline init pattern at the bottom of the script, add it there.

**Step 3: Verify toggle works**

Open app. Toggle ON → `body.help-on` class appears. Toggle OFF → class removed. Refresh → state persists.

**Step 4: Commit**

```bash
git add index.html
git commit -m "feat: add help bubble toggle/dismiss JS logic with localStorage"
```

---

### Task 4: Add i18n translations for all 15 help bubbles in all 7 languages

**Files:**
- Modify: `index.html` — TRANSLATIONS object in each language block

**Step 1: Add keys to Polish block (before line 1508)**

Add these keys before the closing `},` of the `pl:` block:

```javascript
    help_toggle_label: 'Pokaż podpowiedzi',
    help_lang: 'Wybierz język / Select language / Seleziona lingua / Choisir la langue / Sprache wählen / Vyberte jazyk / Selecione o idioma',
    help_event_name: 'Wpisz nazwę zdarzenia — np. lokalizacja, typ wypadku. To pojawi się w raporcie.',
    help_start_btn: 'Rozpoczyna zdarzenie i uruchamia stoper.',
    help_toggle_info: 'Włącza lub wyłącza dymki z podpowiedziami.',
    help_stats_bar: 'Liczba pacjentów w każdej kategorii: 🔴 T1 natychmiastowy, 🟡 T2 pilny, 🟢 T3 odroczony, ⚫ T4 zgon.',
    help_nav: 'Przełączaj między trzema sekcjami: segregacja pacjentów, dysponowanie transportu, raport.',
    help_wizard: 'Algorytm START krok po kroku. Odpowiadaj TAK/NIE, aby sklasyfikować pacjenta.',
    help_result: 'Wynik segregacji. Dodaj notatki, zatwierdź pacjenta lub zmień kolor ręcznie.',
    help_patient_list: 'Lista pacjentów na miejscu zdarzenia, posortowana wg priorytetu.',
    help_transport_list: 'Pacjenci aktualnie transportowani do szpitali.',
    help_delivered_btn: 'Oznacza pacjenta jako dostarczonego do szpitala. Zespół ZRM zostaje zwolniony i ponownie pojawia się na liście do dyspozycji.',
    help_dispatch_form: 'Wybierz pacjenta, zespół ZRM i szpital, aby zadysponować transport.',
    help_manage_teams: 'Dodawaj i usuwaj zespoły ZRM dostępne na miejscu zdarzenia.',
    help_manage_hospitals: 'Dodawaj i usuwaj szpitale docelowe.',
    help_report: 'Automatycznie generowany raport ze zdarzenia. Możesz skopiować tekst lub wysłać przez SMS/e-mail.',
```

**Step 2: Add keys to English block (before line 1630)**

```javascript
    help_toggle_label: 'Show hints',
    help_lang: 'Wybierz język / Select language / Seleziona lingua / Choisir la langue / Sprache wählen / Vyberte jazyk / Selecione o idioma',
    help_event_name: 'Enter the event name — e.g. location, type of accident. This will appear in the report.',
    help_start_btn: 'Starts the event and begins the timer.',
    help_toggle_info: 'Turns help bubbles on or off.',
    help_stats_bar: 'Patient count per category: 🔴 T1 immediate, 🟡 T2 delayed, 🟢 T3 minor, ⚫ T4 deceased.',
    help_nav: 'Switch between three sections: patient triage, transport dispatch, report.',
    help_wizard: 'START algorithm step by step. Answer YES/NO to classify the patient.',
    help_result: 'Triage result. Add notes, confirm patient, or manually override the colour.',
    help_patient_list: 'Patients currently at the scene, sorted by priority.',
    help_transport_list: 'Patients currently being transported to hospitals.',
    help_delivered_btn: 'Marks the patient as delivered to the hospital. The EMS team is freed and reappears in the dispatch list for a new assignment.',
    help_dispatch_form: 'Select a patient, EMS team, and hospital to dispatch transport.',
    help_manage_teams: 'Add or remove EMS teams available on scene.',
    help_manage_hospitals: 'Add or remove destination hospitals.',
    help_report: 'Auto-generated event report. Copy the text or send via SMS/email.',
```

**Step 3: Add keys to Italian block (before line 1752)**

```javascript
    help_toggle_label: 'Mostra suggerimenti',
    help_lang: 'Wybierz język / Select language / Seleziona lingua / Choisir la langue / Sprache wählen / Vyberte jazyk / Selecione o idioma',
    help_event_name: 'Inserisci il nome dell\'evento — es. luogo, tipo di incidente. Apparirà nel rapporto.',
    help_start_btn: 'Avvia l\'evento e il cronometro.',
    help_toggle_info: 'Attiva o disattiva i suggerimenti.',
    help_stats_bar: 'Conteggio pazienti per categoria: 🔴 T1 immediato, 🟡 T2 urgente, 🟢 T3 differito, ⚫ T4 deceduto.',
    help_nav: 'Passa tra le tre sezioni: triage pazienti, invio trasporto, rapporto.',
    help_wizard: 'Algoritmo START passo per passo. Rispondi SÌ/NO per classificare il paziente.',
    help_result: 'Risultato del triage. Aggiungi note, conferma il paziente o cambia il colore manualmente.',
    help_patient_list: 'Pazienti attualmente sulla scena, ordinati per priorità.',
    help_transport_list: 'Pazienti attualmente in trasporto verso gli ospedali.',
    help_delivered_btn: 'Segna il paziente come consegnato all\'ospedale. Il team EMS viene liberato e riappare nella lista di invio.',
    help_dispatch_form: 'Seleziona paziente, team EMS e ospedale per inviare il trasporto.',
    help_manage_teams: 'Aggiungi o rimuovi team EMS disponibili sulla scena.',
    help_manage_hospitals: 'Aggiungi o rimuovi ospedali di destinazione.',
    help_report: 'Rapporto evento generato automaticamente. Copia il testo o invia via SMS/email.',
```

**Step 4: Add keys to French block (before line 1874)**

```javascript
    help_toggle_label: 'Afficher les indices',
    help_lang: 'Wybierz język / Select language / Seleziona lingua / Choisir la langue / Sprache wählen / Vyberte jazyk / Selecione o idioma',
    help_event_name: 'Entrez le nom de l\'événement — ex. lieu, type d\'accident. Il apparaîtra dans le rapport.',
    help_start_btn: 'Démarre l\'événement et le chronomètre.',
    help_toggle_info: 'Active ou désactive les bulles d\'aide.',
    help_stats_bar: 'Nombre de patients par catégorie : 🔴 T1 immédiat, 🟡 T2 urgent, 🟢 T3 différé, ⚫ T4 décédé.',
    help_nav: 'Basculer entre les trois sections : triage, envoi de transport, rapport.',
    help_wizard: 'Algorithme START étape par étape. Répondez OUI/NON pour classifier le patient.',
    help_result: 'Résultat du triage. Ajoutez des notes, confirmez le patient ou changez la couleur manuellement.',
    help_patient_list: 'Patients actuellement sur les lieux, triés par priorité.',
    help_transport_list: 'Patients en cours de transport vers les hôpitaux.',
    help_delivered_btn: 'Marque le patient comme livré à l\'hôpital. L\'équipe EMS est libérée et réapparaît dans la liste d\'envoi.',
    help_dispatch_form: 'Sélectionnez un patient, une équipe EMS et un hôpital pour envoyer le transport.',
    help_manage_teams: 'Ajoutez ou supprimez les équipes EMS disponibles sur place.',
    help_manage_hospitals: 'Ajoutez ou supprimez les hôpitaux de destination.',
    help_report: 'Rapport d\'événement généré automatiquement. Copiez le texte ou envoyez par SMS/e-mail.',
```

**Step 5: Add keys to German block (before line 1996)**

```javascript
    help_toggle_label: 'Hinweise anzeigen',
    help_lang: 'Wybierz język / Select language / Seleziona lingua / Choisir la langue / Sprache wählen / Vyberte jazyk / Selecione o idioma',
    help_event_name: 'Ereignisname eingeben — z.B. Ort, Unfallart. Erscheint im Bericht.',
    help_start_btn: 'Startet das Ereignis und die Zeitmessung.',
    help_toggle_info: 'Schaltet die Hilfe-Blasen ein oder aus.',
    help_stats_bar: 'Patientenanzahl pro Kategorie: 🔴 T1 sofort, 🟡 T2 dringend, 🟢 T3 verzögert, ⚫ T4 verstorben.',
    help_nav: 'Zwischen drei Bereichen wechseln: Triage, Transportdisposition, Bericht.',
    help_wizard: 'START-Algorithmus Schritt für Schritt. Antworten Sie JA/NEIN, um den Patienten zu klassifizieren.',
    help_result: 'Triage-Ergebnis. Notizen hinzufügen, Patient bestätigen oder Farbe manuell ändern.',
    help_patient_list: 'Patienten vor Ort, sortiert nach Priorität.',
    help_transport_list: 'Patienten, die derzeit in Krankenhäuser transportiert werden.',
    help_delivered_btn: 'Markiert den Patienten als im Krankenhaus zugestellt. Das Rettungsteam wird frei und erscheint wieder in der Dispositionsliste.',
    help_dispatch_form: 'Patient, Rettungsteam und Krankenhaus auswählen, um den Transport zu disponieren.',
    help_manage_teams: 'Rettungsteams hinzufügen oder entfernen.',
    help_manage_hospitals: 'Zielkrankenhäuser hinzufügen oder entfernen.',
    help_report: 'Automatisch generierter Ereignisbericht. Text kopieren oder per SMS/E-Mail senden.',
```

**Step 6: Add keys to Czech block (before line 2118)**

```javascript
    help_toggle_label: 'Zobrazit nápovědu',
    help_lang: 'Wybierz język / Select language / Seleziona lingua / Choisir la langue / Sprache wählen / Vyberte jazyk / Selecione o idioma',
    help_event_name: 'Zadejte název události — např. místo, typ nehody. Objeví se v hlášení.',
    help_start_btn: 'Spustí událost a časomíru.',
    help_toggle_info: 'Zapíná nebo vypíná bubliny s nápovědou.',
    help_stats_bar: 'Počet pacientů v každé kategorii: 🔴 T1 okamžitý, 🟡 T2 naléhavý, 🟢 T3 odložený, ⚫ T4 zemřelý.',
    help_nav: 'Přepínejte mezi třemi sekcemi: třídění pacientů, dispečink transportu, hlášení.',
    help_wizard: 'Algoritmus START krok za krokem. Odpovězte ANO/NE pro klasifikaci pacienta.',
    help_result: 'Výsledek třídění. Přidejte poznámky, potvrďte pacienta nebo ručně změňte barvu.',
    help_patient_list: 'Pacienti na místě události, seřazení podle priority.',
    help_transport_list: 'Pacienti aktuálně přepravovaní do nemocnic.',
    help_delivered_btn: 'Označí pacienta jako doručeného do nemocnice. Tým ZZS se uvolní a znovu se objeví v seznamu k dispozici.',
    help_dispatch_form: 'Vyberte pacienta, tým ZZS a nemocnici pro odeslání transportu.',
    help_manage_teams: 'Přidávejte nebo odebírejte týmy ZZS dostupné na místě.',
    help_manage_hospitals: 'Přidávejte nebo odebírejte cílové nemocnice.',
    help_report: 'Automaticky vygenerované hlášení z události. Zkopírujte text nebo odešlete přes SMS/e-mail.',
```

**Step 7: Add keys to Portuguese block (before line 2240)**

```javascript
    help_toggle_label: 'Mostrar dicas',
    help_lang: 'Wybierz język / Select language / Seleziona lingua / Choisir la langue / Sprache wählen / Vyberte jazyk / Selecione o idioma',
    help_event_name: 'Insira o nome do evento — ex. local, tipo de acidente. Aparecerá no relatório.',
    help_start_btn: 'Inicia o evento e o cronómetro.',
    help_toggle_info: 'Ativa ou desativa as dicas de ajuda.',
    help_stats_bar: 'Contagem de pacientes por categoria: 🔴 T1 imediato, 🟡 T2 urgente, 🟢 T3 adiado, ⚫ T4 falecido.',
    help_nav: 'Alternar entre as três secções: triagem, despacho de transporte, relatório.',
    help_wizard: 'Algoritmo START passo a passo. Responda SIM/NÃO para classificar o paciente.',
    help_result: 'Resultado da triagem. Adicione notas, confirme o paciente ou altere a cor manualmente.',
    help_patient_list: 'Pacientes atualmente no local, ordenados por prioridade.',
    help_transport_list: 'Pacientes atualmente em transporte para hospitais.',
    help_delivered_btn: 'Marca o paciente como entregue no hospital. A equipa EMS fica disponível e reaparece na lista de despacho.',
    help_dispatch_form: 'Selecione paciente, equipa EMS e hospital para despachar transporte.',
    help_manage_teams: 'Adicione ou remova equipas EMS disponíveis no local.',
    help_manage_hospitals: 'Adicione ou remova hospitais de destino.',
    help_report: 'Relatório do evento gerado automaticamente. Copie o texto ou envie por SMS/e-mail.',
```

**Step 8: Commit**

```bash
git add index.html
git commit -m "feat: add i18n translations for 15 help bubbles in all 7 languages"
```

---

### Task 5: Add help bubble HTML to setup screen (4 bubbles)

**Files:**
- Modify: `index.html` — setup screen HTML (lines 1167-1185)

**Step 1: Add bubbles to setup screen**

The setup screen (`#setup`) needs `position: relative` (already has it). Add bubbles inside the `#setup` div. Each parent element needs `position: relative` if it doesn't have it.

Wrap the language selector and add its bubble. Insert after the `<select>` closing tag (line 1176):

```html
  <div class="help-bubble" data-help-id="help-lang" data-arrow="down" style="position:fixed;top:44px;left:50%;transform:translateX(-50%);z-index:201;text-align:center;">
    <button class="help-close" onclick="dismissBubble('help-lang')">✕</button>
    <span data-i18n="help_lang">Wybierz język / Select language / Seleziona lingua / Choisir la langue / Sprache wählen / Vyberte jazyk / Selecione o idioma</span>
  </div>
```

Insert after the event name input (line 1182):

```html
  <div style="position:relative;width:340px;max-width:90vw;">
    <div class="help-bubble" data-help-id="help-event-name" data-arrow="up" style="top:4px;left:0;right:0;">
      <button class="help-close" onclick="dismissBubble('help-event-name')">✕</button>
      <span data-i18n="help_event_name">Wpisz nazwę zdarzenia — np. lokalizacja, typ wypadku. To pojawi się w raporcie.</span>
    </div>
  </div>
```

Insert after the btn-start button (line 1183), before the help toggle:

```html
  <div style="position:relative;">
    <div class="help-bubble" data-help-id="help-start-btn" data-arrow="up" style="top:4px;left:50%;transform:translateX(-50%);">
      <button class="help-close" onclick="dismissBubble('help-start-btn')">✕</button>
      <span data-i18n="help_start_btn">Rozpoczyna zdarzenie i uruchamia stoper.</span>
    </div>
  </div>
```

Insert after the help toggle label:

```html
  <div style="position:relative;">
    <div class="help-bubble" data-help-id="help-toggle" data-arrow="up" style="top:4px;left:50%;transform:translateX(-50%);">
      <button class="help-close" onclick="dismissBubble('help-toggle')">✕</button>
      <span data-i18n="help_toggle_info">Włącza lub wyłącza dymki z podpowiedziami.</span>
    </div>
  </div>
```

**Step 2: Verify all 4 setup bubbles render**

Open app, enable toggle. All 4 bubbles should appear overlapping their target elements.

**Step 3: Commit**

```bash
git add index.html
git commit -m "feat: add 4 help bubbles to setup screen"
```

---

### Task 6: Add help bubble HTML to header and Segregacja view (7 bubbles)

**Files:**
- Modify: `index.html` — header and triage view HTML (lines 1240-1294)

**Step 1: Add stats bar bubble**

Insert inside `#header`, after the `.stats-bar` div (after line 1243):

```html
      <div class="help-bubble" data-help-id="help-stats-bar" data-arrow="up" style="top:100%;left:0;margin-top:4px;">
        <button class="help-close" onclick="dismissBubble('help-stats-bar')">✕</button>
        <span data-i18n="help_stats_bar">Liczba pacjentów w każdej kategorii: 🔴 T1 natychmiastowy, 🟡 T2 pilny, 🟢 T3 odroczony, ⚫ T4 zgon.</span>
      </div>
```

**Step 2: Add nav bubble**

Insert after the `.nav` div closing tag (after line 1248):

```html
      <div class="help-bubble" data-help-id="help-nav" data-arrow="up" style="top:100%;right:0;margin-top:4px;">
        <button class="help-close" onclick="dismissBubble('help-nav')">✕</button>
        <span data-i18n="help_nav">Przełączaj między trzema sekcjami: segregacja pacjentów, dysponowanie transportu, raport.</span>
      </div>
```

**Step 3: Add wizard bubble**

Insert inside `#triageWizard` section, after the `.wizard-box` div (after line 1259):

```html
        <div class="help-bubble" data-help-id="help-wizard" data-arrow="up" style="top:0;left:0;right:0;">
          <button class="help-close" onclick="dismissBubble('help-wizard')">✕</button>
          <span data-i18n="help_wizard">Algorytm START krok po kroku. Odpowiadaj TAK/NIE, aby sklasyfikować pacjenta.</span>
        </div>
```

**Step 4: Add result card bubble**

Insert inside `#triageResult` section, after the `.override-row` div (after line 1283):

```html
        <div class="help-bubble" data-help-id="help-result" data-arrow="up" style="top:0;left:0;right:0;">
          <button class="help-close" onclick="dismissBubble('help-result')">✕</button>
          <span data-i18n="help_result">Wynik segregacji. Dodaj notatki, zatwierdź pacjenta lub zmień kolor ręcznie.</span>
        </div>
```

**Step 5: Add patient list bubble**

Insert inside the patient list section (after `#patientList` div, line 1288):

```html
        <div class="help-bubble" data-help-id="help-patient-list" data-arrow="up" style="position:relative;margin-top:8px;">
          <button class="help-close" onclick="dismissBubble('help-patient-list')">✕</button>
          <span data-i18n="help_patient_list">Lista pacjentów na miejscu zdarzenia, posortowana wg priorytetu.</span>
        </div>
```

**Step 6: Add transport list bubble**

Insert inside `#transportSection` (after `#transportingList` div, line 1292):

```html
        <div class="help-bubble" data-help-id="help-transport-list" data-arrow="up" style="position:relative;margin-top:8px;">
          <button class="help-close" onclick="dismissBubble('help-transport-list')">✕</button>
          <span data-i18n="help_transport_list">Pacjenci aktualnie transportowani do szpitali.</span>
        </div>
```

**Step 7: Add delivered button bubble**

The delivered button is rendered dynamically by JS in the `renderAll()` function. This bubble needs to be injected by JS when transport patients are rendered. Find the transport list rendering code (around line 2684 where `btn-delivered` is generated) and add the bubble HTML after each delivered button in the template, OR add a single bubble inside `#transportSection` that describes the delivered button concept:

```html
        <div class="help-bubble" data-help-id="help-delivered-btn" data-arrow="up" style="position:relative;margin-top:8px;">
          <button class="help-close" onclick="dismissBubble('help-delivered-btn')">✕</button>
          <span data-i18n="help_delivered_btn">Oznacza pacjenta jako dostarczonego do szpitala. Zespół ZRM zostaje zwolniony i ponownie pojawia się na liście do dyspozycji.</span>
        </div>
```

Place this inside `#transportSection`, after the `help-transport-list` bubble.

**Step 8: Ensure `#header` and parent sections have `position: relative`**

Add `position: relative;` to `#header` CSS if not present (check line 50-57). The `#header` already has `position: sticky` which establishes a positioning context, so this should work.

**Step 9: Verify all 7 bubbles render**

Start event, enable toggle. All bubbles should appear in the Segregacja view. Transport bubbles only visible when transport section is visible.

**Step 10: Commit**

```bash
git add index.html
git commit -m "feat: add 7 help bubbles to header and segregacja view"
```

---

### Task 7: Add help bubble HTML to Dysponowanie view (3 bubbles)

**Files:**
- Modify: `index.html` — dispatch view HTML (lines 1297-1346)

**Step 1: Add dispatch form bubble**

Insert inside the dispatch form section, after the `btn-dispatch` button (after line 1315):

```html
        <div class="help-bubble" data-help-id="help-dispatch-form" data-arrow="up" style="position:relative;margin-top:8px;">
          <button class="help-close" onclick="dismissBubble('help-dispatch-form')">✕</button>
          <span data-i18n="help_dispatch_form">Wybierz pacjenta, zespół ZRM i szpital, aby zadysponować transport.</span>
        </div>
```

**Step 2: Add teams panel bubble**

Insert inside the first `.manage-panel` (teams), after the `.manage-box` closing tag (after line 1328):

```html
          <div class="help-bubble" data-help-id="help-manage-teams" data-arrow="up" style="position:relative;margin-top:8px;">
            <button class="help-close" onclick="dismissBubble('help-manage-teams')">✕</button>
            <span data-i18n="help_manage_teams">Dodawaj i usuwaj zespoły ZRM dostępne na miejscu zdarzenia.</span>
          </div>
```

**Step 3: Add hospitals panel bubble**

Insert inside the second `.manage-panel` (hospitals), after the `.manage-box` closing tag (after line 1338):

```html
          <div class="help-bubble" data-help-id="help-manage-hospitals" data-arrow="up" style="position:relative;margin-top:8px;">
            <button class="help-close" onclick="dismissBubble('help-manage-hospitals')">✕</button>
            <span data-i18n="help_manage_hospitals">Dodawaj i usuwaj szpitale docelowe.</span>
          </div>
```

**Step 4: Verify all 3 bubbles render**

Switch to DYSPONOWANIE tab with toggle ON. All 3 bubbles should appear.

**Step 5: Commit**

```bash
git add index.html
git commit -m "feat: add 3 help bubbles to dysponowanie view"
```

---

### Task 8: Add help bubble HTML to Raport view (1 bubble)

**Files:**
- Modify: `index.html` — report view HTML (lines 1349-1355)

**Step 1: Add report bubble**

Insert inside `#view-report`, after the `.report-header-row` div (after line 1353):

```html
      <div class="help-bubble" data-help-id="help-report" data-arrow="up" style="position:relative;margin-top:8px;">
        <button class="help-close" onclick="dismissBubble('help-report')">✕</button>
        <span data-i18n="help_report">Automatycznie generowany raport ze zdarzenia. Możesz skopiować tekst lub wysłać przez SMS/e-mail.</span>
      </div>
```

**Step 2: Verify bubble renders**

Switch to RAPORT tab with toggle ON. Bubble should appear.

**Step 3: Commit**

```bash
git add index.html
git commit -m "feat: add help bubble to raport view"
```

---

### Task 9: Integration testing and polish

**Files:**
- Modify: `index.html` — potential CSS/JS tweaks

**Step 1: Full flow test**

1. Fresh load (no localStorage) → toggle OFF, no bubbles visible
2. Toggle ON on setup → all 4 setup bubbles appear
3. Dismiss 2 bubbles with ✕ → they disappear
4. Toggle OFF → all bubbles hidden
5. Toggle ON → ALL bubbles reappear (including previously dismissed)
6. Start event → header toggle synced, setup toggle hidden
7. Segregacja view → 5 bubbles visible (transport bubbles hidden until transport exists)
8. Switch to Dysponowanie → 3 bubbles visible
9. Switch to Raport → 1 bubble visible
10. Dismiss some bubbles, refresh page → toggle state persisted, but dismissed bubbles reappear
11. Switch language → all bubble texts update to new language (except `help-lang` which always shows all 7)

**Step 2: Test on mobile viewport**

Resize to 375px width. Bubbles should go full-width, arrows hidden.

**Step 3: Fix any positioning or z-index issues**

Adjust bubble positions if they clip off-screen or overlap badly.

**Step 4: Bump service worker cache version**

In `sw.js`, increment the cache version so returning users get the updated app.

**Step 5: Commit**

```bash
git add index.html sw.js
git commit -m "feat: help bubbles integration testing and SW cache bump"
```

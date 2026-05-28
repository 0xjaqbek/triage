# MZ Document Generation (Zał. 12, 13, 14) Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** After setting `eventEndTime`, a button in the Report tab opens a supplement modal; user fills in missing fields, clicks "Generuj i pobierz", and receives a single HTML file containing all 3 official MZ post-incident documents (Załącznik 12, 13, 14) ready to print as PDF.

**Architecture:** Single-file modification (`index.html` only). New CSS and HTML modal added inline. New JS functions added before `</script>`. `renderReport()` extended to show the trigger button when `eventEndTime` is set.

**Tech Stack:** Vanilla JS, inline CSS, Blob API for download. No external dependencies. Polish-language output only.

---

## File

- Modify: `index.html`
  - CSS insertion before `</style>` at line 1565
  - HTML insertion before `<!-- GDM MODAL -->` at line 2008
  - JS modification in `renderReport()` ending at line 4815
  - JS insertion before `</script>` at line 5785

---

## Task 1: Trigger button in Report tab

**Files:**
- Modify: `index.html` (HTML at line 2003, JS in `renderReport()`)

- [ ] **Step 1: Add button HTML after `<div id="reportContent"></div>`**

Find this exact string in `index.html`:
```html
      <div id="reportContent"></div>
    </div>
  </div>
</div>
```

Replace with:
```html
      <div id="reportContent"></div>
      <div id="btnGenerateMzWrap" style="display:none;padding:0 0 24px">
        <button class="btn-generate-mz" onclick="openMzModal()">📄 Generuj dokumenty MZ — Zał. 12, 13, 14</button>
      </div>
    </div>
  </div>
</div>
```

- [ ] **Step 2: Add `.btn-generate-mz` CSS before `</style>`**

Find `</style>` (line 1565) and insert before it:
```css
/* ===== MZ DOCUMENT GENERATION ===== */
.btn-generate-mz {
  display: block; width: 100%; padding: 14px 20px; margin-top: 4px;
  background: rgba(37,99,235,0.12); border: 2px solid #3b82f6;
  border-radius: 10px; color: #93c5fd; font-size: 13px; font-weight: 700;
  font-family: var(--font); letter-spacing: 0.05em; cursor: pointer;
  text-transform: uppercase; transition: all 0.2s;
}
.btn-generate-mz:hover { background: rgba(37,99,235,0.28); color: #fff; }
```

- [ ] **Step 3: Show/hide button in `renderReport()`**

Find the last line of `renderReport()`:
```js
  document.getElementById('reportContent').innerHTML = html;
}
```

Replace with:
```js
  document.getElementById('reportContent').innerHTML = html;
  const mzWrap = document.getElementById('btnGenerateMzWrap');
  if (mzWrap) mzWrap.style.display = eventEndTime ? '' : 'none';
}
```

- [ ] **Step 4: Verify in browser**

1. Open `index.html`, start an event, add any patient, navigate to Report tab.
2. The "Generuj dokumenty MZ" button must NOT be visible.
3. Click "ZAKOŃCZ ZDARZENIE" and confirm.
4. The button must NOW appear below the report content.
5. Click "COFNIJ ZAKOŃCZENIE" — button must disappear again.

- [ ] **Step 5: Commit**

```bash
git add index.html
git commit -m "feat: add Generuj dokumenty MZ trigger button in Report tab"
```

---

## Task 2: Supplement modal CSS + HTML shell

**Files:**
- Modify: `index.html` (CSS before `</style>`, HTML before `<!-- GDM MODAL -->`)

- [ ] **Step 1: Add modal CSS before `</style>`**

After the `.btn-generate-mz` CSS from Task 1, append:
```css
.mz-section { border: 1px solid var(--border); border-radius: 8px; margin-bottom: 10px; overflow: hidden; }
.mz-section-header {
  display: flex; align-items: center; justify-content: space-between;
  padding: 10px 14px; background: var(--bg);
}
.mz-section-title { font-size: 11px; font-weight: 700; color: var(--white); text-transform: uppercase; letter-spacing: 0.04em; }
.mz-skip-btn {
  font-size: 11px; padding: 3px 10px; border-radius: 5px; border: 1px solid var(--border2);
  background: transparent; color: var(--text2); font-family: var(--font); cursor: pointer; transition: all 0.15s;
}
.mz-skip-btn:hover { border-color: var(--yellow); color: var(--yellow); }
.mz-section.mz-skipped .mz-skip-btn { background: var(--yellow); color: #000; border-color: var(--yellow); font-weight: 700; }
.mz-section.mz-skipped .mz-section-body { display: none; }
.mz-section-body { padding: 10px 14px; }
.mz-field-row { display: flex; align-items: center; gap: 8px; margin-bottom: 8px; flex-wrap: wrap; }
.mz-field-row label { font-size: 11px; color: var(--text2); min-width: 190px; flex-shrink: 0; }
.mz-field-input {
  flex: 1; min-width: 0; background: var(--bg); border: 1px solid var(--border2); border-radius: 6px;
  color: var(--white); padding: 5px 8px; font-size: 12px; font-family: var(--font);
}
.mz-field-input:focus { outline: none; border-color: var(--blue); }
.mz-field-input.mz-number { max-width: 80px; flex: none; }
.mz-field-readonly .mz-auto-value { font-size: 12px; color: var(--white); font-weight: 700; }
.mz-select {
  background: var(--bg); border: 1px solid var(--border2); border-radius: 6px;
  color: var(--white); padding: 5px 8px; font-size: 12px; font-family: var(--font); cursor: pointer;
}
.mz-table { width: 100%; border-collapse: collapse; font-size: 11px; margin-top: 6px; }
.mz-table th, .mz-table td { border: 1px solid var(--border); padding: 4px 8px; text-align: left; }
.mz-table th { background: var(--bg); color: var(--text2); font-weight: 600; }
.mz-table td .mz-field-input { width: 64px; }
.mz-hospital-card { border: 1px solid var(--border); border-radius: 6px; margin-bottom: 8px; overflow: hidden; }
.mz-hospital-card-header {
  padding: 8px 12px; background: var(--bg); cursor: pointer; display: flex;
  align-items: center; justify-content: space-between; font-size: 12px; color: var(--white); font-weight: 600;
}
.mz-hospital-card-header:hover { background: var(--border); }
.mz-hospital-card.mz-collapsed .mz-hospital-card-body { display: none; }
.mz-hospital-card-body { padding: 10px 12px; }
.mz-generate-btn {
  padding: 9px 22px; font-size: 12px; background: var(--red); border: none; border-radius: 8px;
  color: var(--white); font-family: var(--font); font-weight: 700; cursor: pointer; letter-spacing: 0.05em;
  transition: opacity 0.15s;
}
.mz-generate-btn:hover { opacity: 0.85; }
```

- [ ] **Step 2: Add modal HTML before `<!-- GDM MODAL -->`**

Find:
```html
<!-- GDM MODAL -->
```

Insert before it:
```html
<!-- MZ DOCUMENTS MODAL -->
<div class="modal-overlay" id="mzModal">
  <div class="modal" style="max-width:520px">
    <div class="modal-header">
      <div class="modal-title">📄 Generuj dokumenty MZ — Zał. 12 / 13 / 14</div>
      <button class="modal-close" onclick="closeMzModal()">&times;</button>
    </div>
    <div class="modal-body" id="mzModalBody" style="padding:14px 16px">
      <!-- rendered dynamically by renderMzModalBody() -->
    </div>
    <div class="modal-footer" style="justify-content:space-between;align-items:center;padding:12px 16px">
      <span style="font-size:11px;color:var(--text3)">Plik HTML gotowy do druku jako PDF</span>
      <button class="mz-generate-btn" onclick="downloadMzDocs()">⬇ Generuj i pobierz</button>
    </div>
  </div>
</div>

```

- [ ] **Step 3: Add stub JS functions before `</script>`**

Find `</script>` (last occurrence, near line 5785) and insert before it:
```js
// ===== MZ DOCUMENT GENERATION =====
function openMzModal() {
  document.getElementById('mzModalBody').innerHTML = '<p style="color:var(--text2);font-size:12px">Ładowanie...</p>';
  document.getElementById('mzModal').classList.add('open');
}
function closeMzModal() {
  document.getElementById('mzModal').classList.remove('open');
}
function downloadMzDocs() {
  closeMzModal();
  alert('TODO: generateMzHtml()');
}
```

- [ ] **Step 4: Verify in browser**

1. Open `index.html`, start event, end it, go to Report tab.
2. Click "📄 Generuj dokumenty MZ" — modal must open with "Ładowanie..." text.
3. Click × — modal must close.
4. Click "⬇ Generuj i pobierz" — must show alert "TODO: generateMzHtml()".

- [ ] **Step 5: Commit**

```bash
git add index.html
git commit -m "feat: add MZ documents supplement modal shell (CSS + HTML + stub JS)"
```

---

## Task 3: Modal section rendering (`renderMzModalBody`)

**Files:**
- Modify: `index.html` (JS — replace stub `openMzModal` with full implementation)

- [ ] **Step 1: Replace stub `openMzModal` and add helpers**

Find the stub section added in Task 2:
```js
// ===== MZ DOCUMENT GENERATION =====
function openMzModal() {
  document.getElementById('mzModalBody').innerHTML = '<p style="color:var(--text2);font-size:12px">Ładowanie...</p>';
  document.getElementById('mzModal').classList.add('open');
}
function closeMzModal() {
  document.getElementById('mzModal').classList.remove('open');
}
function downloadMzDocs() {
  closeMzModal();
  alert('TODO: generateMzHtml()');
}
```

Replace with:
```js
// ===== MZ DOCUMENT GENERATION =====
function openMzModal() {
  renderMzModalBody();
  document.getElementById('mzModal').classList.add('open');
}

function closeMzModal() {
  document.getElementById('mzModal').classList.remove('open');
}

function toggleMzSection(id) {
  document.getElementById(id).classList.toggle('mz-skipped');
}

function mzSection(id, title, bodyHtml) {
  return `<div class="mz-section" id="${id}">
    <div class="mz-section-header">
      <span class="mz-section-title">${title}</span>
      <button class="mz-skip-btn" onclick="toggleMzSection('${id}')">Pomiń →</button>
    </div>
    <div class="mz-section-body">${bodyHtml}</div>
  </div>`;
}

function renderMzModalBody() {
  const transportedCount = patients.filter(p => p.hospitalId).length;
  const onSceneCount = patients.filter(p => p.status === 'na_miejscu').length;

  // Section 1: Dane ogólne
  const sec1 = mzSection('mz-sec-general', 'Dane ogólne (Raport — Sek. 1, 8, 9)', `
    <div class="mz-field-row"><label>Dysponent (dyspozytornia)</label><input class="mz-field-input" id="mz-dysponent" placeholder="np. WSRM w Warszawie"></div>
    <div class="mz-field-row"><label>WKRM województwo</label><input class="mz-field-input" id="mz-wkrm" placeholder="np. mazowieckie"></div>
    <div class="mz-field-row"><label>Wysłano do WKRM</label><select class="mz-select" id="mz-sent-wkrm"><option value="NIE">NIE</option><option value="TAK">TAK</option></select></div>
    <div class="mz-field-row"><label>Autor raportu</label><input class="mz-field-input" id="mz-author" placeholder="Imię i nazwisko"></div>
  `);

  // Section 2: Typy ZRM
  const zrmRows = teams.length
    ? teams.map(tm => `<div class="mz-field-row"><label>${escHtml(tm.name)}</label><select class="mz-select" id="mz-zrm-${tm.id}"><option value="P">P</option><option value="S">S</option><option value="LZRM">LZRM</option></select></div>`).join('')
    : '<p style="color:var(--text3);font-size:12px">Brak zespołów ZRM</p>';
  const sec2 = mzSection('mz-sec-zrm', 'Typy ZRM (Raport — Sek. 5)', zrmRows);

  // Section 3: KAM
  const sec3 = mzSection('mz-sec-kam', 'Szczegóły KAM (Raport — Sek. 6)', `
    <div class="mz-field-row"><label>Imię i nazwisko KAM</label><input class="mz-field-input" id="mz-kam-name" value="${escHtml(kamName)}"></div>
    <div class="mz-field-row"><label>Rola KAM</label><select class="mz-select" id="mz-kam-role"><option value="Ratownik Medyczny">Ratownik Medyczny</option><option value="Lekarz">Lekarz</option><option value="Pielęgniarka">Pielęgniarka</option></select></div>
    <div class="mz-field-row"><label>Pełnił funkcję od</label><input class="mz-field-input" id="mz-kam-from" value="${escHtml(eventTime)}"></div>
    <div class="mz-field-row"><label>Pełnił funkcję do</label><input class="mz-field-input" id="mz-kam-to" value="${escHtml(eventEndTime)}"></div>
  `);

  // Section 4: Sposoby transportu
  const sec4 = mzSection('mz-sec-transport', 'Sposoby transportu (Raport — Sek. 7)', `
    <div class="mz-field-row mz-field-readonly"><label>Łącznie poszkodowanych (auto)</label><span class="mz-auto-value">${patients.length}</span></div>
    <div class="mz-field-row mz-field-readonly"><label>Przetransportowanych (auto)</label><span class="mz-auto-value">${transportedCount}</span></div>
    <div class="mz-field-row mz-field-readonly"><label>Pozostali na miejscu (auto)</label><span class="mz-auto-value">${onSceneCount}</span></div>
    <div class="mz-field-row"><label>Drogą powietrzną</label><input class="mz-field-input mz-number" id="mz-tr-air" type="number" min="0" value="0"></div>
    <div class="mz-field-row"><label>Drogą naziemną bez ZRM</label><input class="mz-field-input mz-number" id="mz-tr-ground-other" type="number" min="0" value="0"></div>
    <div class="mz-field-row"><label>Drogą wodną</label><input class="mz-field-input mz-number" id="mz-tr-water" type="number" min="0" value="0"></div>
    <div class="mz-field-row"><label>Na miejscu bez transportu</label><input class="mz-field-input mz-number" id="mz-tr-onscene" type="number" min="0" value="0"></div>
    <div class="mz-field-row"><label>Nie wyrazili zgody</label><input class="mz-field-input mz-number" id="mz-tr-no-consent" type="number" min="0" value="0"></div>
    <div class="mz-field-row"><label>Odmówili udzielenia pomocy</label><input class="mz-field-input mz-number" id="mz-tr-refused" type="number" min="0" value="0"></div>
    <div class="mz-field-row"><label>Samodzielnie oddalili się</label><input class="mz-field-input mz-number" id="mz-tr-left" type="number" min="0" value="0"></div>
  `);

  // Section 5: Dane szpitali
  const deptLabels = ['OIT','N-chir (neurochirurgia)','Ortopedia','Chirurgia ogólna','Chirurgia naczyniowa','SOR'];
  const deptKeys  = ['oit','neurochir','ortopedia','chir-ogolna','chir-naczyniowa','sor'];

  const hospCards = hospitals.map(h => {
    const deptRows = deptLabels.map((d, i) => `<tr><td>${d}</td>
      <td><input class="mz-field-input mz-number" id="mz-h${h.id}-dept-${deptKeys[i]}-1h" type="number" min="0" placeholder="—"></td>
      <td><input class="mz-field-input mz-number" id="mz-h${h.id}-dept-${deptKeys[i]}-2h" type="number" min="0" placeholder="—"></td></tr>`).join('');
    return `<div class="mz-hospital-card">
      <div class="mz-hospital-card-header" onclick="this.parentElement.classList.toggle('mz-collapsed')">
        🏥 ${escHtml(h.name)}
        <span style="color:var(--text3);font-size:11px;margin-left:6px">🔴${h.redCapacity} 🟡${h.yellowCapacity}</span>
        <span style="flex:1"></span><span style="font-size:10px;color:var(--text3)">▲</span>
      </div>
      <div class="mz-hospital-card-body">
        <div class="mz-field-row"><label>Nr telefonu</label><input class="mz-field-input" id="mz-h${h.id}-phone" placeholder="________"></div>
        <table class="mz-table">
          <thead><tr><th>Zasób</th><th>1 godz.</th><th>2 godz.</th></tr></thead>
          <tbody>
            <tr><td>Stanowiska ITM (z resp.)</td>
              <td><input class="mz-field-input mz-number" id="mz-h${h.id}-itm-1h" type="number" min="0" placeholder="—"></td>
              <td><input class="mz-field-input mz-number" id="mz-h${h.id}-itm-2h" type="number" min="0" placeholder="—"></td></tr>
            <tr><td>Wolne respiratory (poza ITM)</td>
              <td><input class="mz-field-input mz-number" id="mz-h${h.id}-resp-1h" type="number" min="0" placeholder="—"></td>
              <td><input class="mz-field-input mz-number" id="mz-h${h.id}-resp-2h" type="number" min="0" placeholder="—"></td></tr>
            <tr><td>Stanowiska IN</td>
              <td><input class="mz-field-input mz-number" id="mz-h${h.id}-in-1h" type="number" min="0" placeholder="—"></td>
              <td><input class="mz-field-input mz-number" id="mz-h${h.id}-in-2h" type="number" min="0" placeholder="—"></td></tr>
            <tr><td>Bloki operacyjne</td>
              <td><input class="mz-field-input mz-number" id="mz-h${h.id}-or-1h" type="number" min="0" placeholder="—"></td>
              <td><input class="mz-field-input mz-number" id="mz-h${h.id}-or-2h" type="number" min="0" placeholder="—"></td></tr>
          </tbody>
        </table>
        <table class="mz-table" style="margin-top:8px">
          <thead><tr><th>Oddział</th><th>1 godz.</th><th>2 godz.</th></tr></thead>
          <tbody>${deptRows}</tbody>
        </table>
      </div>
    </div>`;
  }).join('') || '<p style="color:var(--text3);font-size:12px">Brak szpitali</p>';

  const sec5 = mzSection('mz-sec-hospitals', 'Dane szpitali (Tabela Szpitali — Zał. 13)', hospCards);

  document.getElementById('mzModalBody').innerHTML = sec1 + sec2 + sec3 + sec4 + sec5;
}

function downloadMzDocs() {
  closeMzModal();
  alert('TODO: generateMzHtml()');
}
```

- [ ] **Step 2: Verify in browser**

1. Open `index.html`, start an event with at least 2 teams and 2 hospitals, add patients with different triage colors, end the event, open the modal.
2. Must see 5 clearly labelled sections.
3. Click "Pomiń →" on Section 1 — section body must hide, button must turn yellow.
4. Click "Pomiń →" again — section body must reappear.
5. Expand/collapse hospital cards by clicking the header.
6. In browser console run:
```js
// After opening modal:
console.assert(!!document.getElementById('mz-sec-general'), 'sec1 present');
console.assert(!!document.getElementById('mz-sec-hospitals'), 'sec5 present');
console.assert(document.querySelectorAll('.mz-section').length === 5, 'exactly 5 sections');
```

- [ ] **Step 3: Commit**

```bash
git add index.html
git commit -m "feat: implement MZ modal section rendering with 5 supplement sections"
```

---

## Task 4: `collectMzSupplement()`

**Files:**
- Modify: `index.html` (JS — replace `downloadMzDocs` stub)

- [ ] **Step 1: Add `collectMzSupplement` and update `downloadMzDocs` stub**

Find:
```js
function downloadMzDocs() {
  closeMzModal();
  alert('TODO: generateMzHtml()');
}
```

Replace with:
```js
function collectMzSupplement() {
  const v  = id => (document.getElementById(id)?.value  || '').trim();
  const n  = id => parseInt(document.getElementById(id)?.value || '0') || 0;
  const sk = id => !!document.getElementById(id)?.classList.contains('mz-skipped');

  const deptKeys = ['oit','neurochir','ortopedia','chir-ogolna','chir-naczyniowa','sor'];
  const transportedCount = patients.filter(p => p.hospitalId).length;
  const air         = n('mz-tr-air');
  const groundOther = n('mz-tr-ground-other');
  const water       = n('mz-tr-water');

  return {
    skipGeneral:   sk('mz-sec-general'),
    dysponent:     v('mz-dysponent'),
    wkrmProvince:  v('mz-wkrm'),
    sentToWkrm:    v('mz-sent-wkrm') === 'TAK',
    reportAuthor:  v('mz-author'),

    skipZrmTypes:  sk('mz-sec-zrm'),
    zrmTypes:      Object.fromEntries(teams.map(tm => [tm.id, v(`mz-zrm-${tm.id}`) || 'P'])),

    skipKam:       sk('mz-sec-kam'),
    kamNameSup:    v('mz-kam-name') || kamName,
    kamRole:       v('mz-kam-role') || 'Ratownik Medyczny',
    kamFrom:       v('mz-kam-from') || eventTime,
    kamTo:         v('mz-kam-to')   || eventEndTime,

    skipTransport:        sk('mz-sec-transport'),
    transportAir:         air,
    transportGroundZrm:   Math.max(0, transportedCount - air - groundOther - water),
    transportGroundOther: groundOther,
    transportWater:       water,
    transportOnScene:     n('mz-tr-onscene'),
    transportNoConsent:   n('mz-tr-no-consent'),
    transportRefused:     n('mz-tr-refused'),
    transportLeft:        n('mz-tr-left'),

    skipHospitals: sk('mz-sec-hospitals'),
    hospitalData:  Object.fromEntries(hospitals.map(h => [h.id, {
      phone:  v(`mz-h${h.id}-phone`),
      itm1h:  v(`mz-h${h.id}-itm-1h`),  itm2h: v(`mz-h${h.id}-itm-2h`),
      resp1h: v(`mz-h${h.id}-resp-1h`), resp2h: v(`mz-h${h.id}-resp-2h`),
      in1h:   v(`mz-h${h.id}-in-1h`),   in2h:  v(`mz-h${h.id}-in-2h`),
      or1h:   v(`mz-h${h.id}-or-1h`),   or2h:  v(`mz-h${h.id}-or-2h`),
      depts:  Object.fromEntries(deptKeys.map(d => [d, {
        h1: v(`mz-h${h.id}-dept-${d}-1h`),
        h2: v(`mz-h${h.id}-dept-${d}-2h`)
      }]))
    }]))
  };
}

function downloadMzDocs() {
  const sup = collectMzSupplement();
  console.log('[MZ] supplement collected:', sup);
  closeMzModal();
  alert('TODO: generateMzHtml(sup)');
}
```

- [ ] **Step 2: Verify in browser console**

1. Open modal, fill "Dysponent" = "Test Dysp.", skip Section 2 (ZRM types), set KAM role to "Lekarz".
2. In console run:
```js
// With modal open:
const sup = collectMzSupplement();
console.assert(sup.dysponent === 'Test Dysp.', 'dysponent captured');
console.assert(sup.skipZrmTypes === true, 'section 2 skipped');
console.assert(sup.kamRole === 'Lekarz', 'KAM role captured');
console.assert(typeof sup.transportGroundZrm === 'number', 'groundZrm is number');
console.assert(sup.transportGroundZrm >= 0, 'groundZrm non-negative');
```

- [ ] **Step 3: Commit**

```bash
git add index.html
git commit -m "feat: implement collectMzSupplement — reads all modal inputs into supplement object"
```

---

## Task 5: `genZal12()` — Tabela Dyslokacji Poszkodowanych

**Files:**
- Modify: `index.html` (JS — add function before `downloadMzDocs`)

- [ ] **Step 1: Add `genZal12()` before `collectMzSupplement`**

Find:
```js
function collectMzSupplement() {
```

Insert before it:
```js
function genZal12() {
  const ZONE_NAMES = {
    head:'Głowa', chest:'Kl. piersiowa', abdomen:'Brzuch',
    upperBack:'Górna cz. pleców', lowerBack:'Dolna cz. pleców',
    leftArm:'Lewa ręka', rightArm:'Prawa ręka',
    leftLeg:'Lewa noga', rightLeg:'Prawa noga'
  };
  const TYPE_NAMES = {
    fracture:'złamanie', bleeding:'krwawienie', burn:'oparzenie',
    wound:'rana', swelling:'obrzęk', pain:'ból', amputation:'amputacja'
  };

  function injText(inj) {
    if (!inj) return '';
    return Object.entries(inj)
      .filter(([, types]) => types && types.length)
      .map(([zone, types]) => `${ZONE_NAMES[zone]||zone}: ${types.map(ty => TYPE_NAMES[ty]||ty).join(', ')}`)
      .join('<br>');
  }

  const groups = [
    { key:'T1', label:'CZERWONI', bg:'#dc2626', textColor:'#fff' },
    { key:'T2', label:'ŻÓŁCI',    bg:'#ca8a04', textColor:'#000' },
    { key:'T3', label:'ZIELONI',  bg:'#16a34a', textColor:'#fff' },
    { key:'T4', label:'CZARNI',   bg:'#374151', textColor:'#fff', black:true }
  ];

  let html = `<div class="mz-doc-header">
    <div class="mz-doc-num">ZAŁĄCZNIK NR 12</div>
    <h2>TABELA DYSLOKACJI POSZKODOWANYCH</h2>
    <p>Zdarzenie: ${escHtml(eventName||'________')} &nbsp;|&nbsp; Data: ${escHtml(eventTime||'________')} &nbsp;|&nbsp; KAM: ${escHtml(kamName||'________')}</p>
  </div>`;

  let globalNr = 1;
  for (const g of groups) {
    const pts = patients.filter(p => p.triage === g.key);
    const cols = g.black ? 3 : 5;
    html += `<table class="mz-triage-table">
      <thead>
        <tr style="background:${g.bg}">
          <th colspan="${cols}" style="color:${g.textColor};font-size:13px;padding:6px 10px;text-align:left">
            ${g.label} — ${pts.length} poszkodowanych
          </th>
        </tr>
        <tr class="mz-triage-col-header">
          <th style="width:36px">Nr</th>
          <th>Obrażenia ciała</th>
          <th>Dodatkowy opis</th>
          ${g.black ? '' : '<th style="width:90px">ZRM</th><th style="width:130px">Szpital</th>'}
        </tr>
      </thead>
      <tbody>`;
    if (pts.length === 0) {
      html += `<tr><td colspan="${cols}" style="text-align:center;color:#888;padding:10px;font-style:italic">Brak poszkodowanych w tej grupie</td></tr>`;
    } else {
      pts.forEach(p => {
        const team = teams.find(t => t.id === p.teamId);
        const hosp = hospitals.find(h => h.id === p.hospitalId);
        html += `<tr>
          <td style="text-align:center;font-weight:700">${globalNr++}</td>
          <td>${injText(p.injuries)}</td>
          <td>${p.notes ? escHtml(p.notes) : ''}</td>
          ${g.black ? '' : `<td>${team ? escHtml(team.name) : '________'}</td><td>${hosp ? escHtml(hosp.name) : '________'}</td>`}
        </tr>`;
      });
    }
    html += `</tbody></table>`;
  }
  return html;
}

```

- [ ] **Step 2: Verify via browser console**

With an active event that has patients (add at least one T1, one T2, one T4), run in console:
```js
const html12 = genZal12();
console.assert(html12.includes('CZERWONI'), 'T1 group present');
console.assert(html12.includes('ŻÓŁCI'),   'T2 group present');
console.assert(html12.includes('ZIELONI'), 'T3 group present');
console.assert(html12.includes('CZARNI'),  'T4 group present');
// T4 must not have ZRM/Szpital columns
const t4idx = html12.indexOf('CZARNI');
const t4block = html12.slice(t4idx, t4idx + 500);
console.assert(!t4block.includes('<th style="width:90px">ZRM</th>'), 'T4 has no ZRM column');
// Check sequential numbering
const nrMatches = [...html12.matchAll(/<td style="text-align:center;font-weight:700">(\d+)<\/td>/g)].map(m => +m[1]);
console.assert(nrMatches[0] === 1, 'numbering starts at 1');
```

- [ ] **Step 3: Commit**

```bash
git add index.html
git commit -m "feat: implement genZal12 — Tabela Dyslokacji Poszkodowanych HTML generator"
```

---

## Task 6: `genZal13()` — Tabela Szpitali

**Files:**
- Modify: `index.html` (JS — add before `collectMzSupplement`)

- [ ] **Step 1: Add `genZal13(sup)` before `collectMzSupplement`**

Find:
```js
function collectMzSupplement() {
```

Insert before it:
```js
function genZal13(sup) {
  const fill = v => (v !== undefined && v !== null && v !== '') ? escHtml(String(v)) : '________';

  const DEPT_LABELS = ['OIT','N-chir (neurochirurgia)','Ortopedia','Chirurgia ogólna','Chirurgia naczyniowa','SOR'];
  const DEPT_KEYS   = ['oit','neurochir','ortopedia','chir-ogolna','chir-naczyniowa','sor'];

  let html = `<div class="mz-doc-header">
    <div class="mz-doc-num">ZAŁĄCZNIK NR 13</div>
    <h2>TABELA SZPITALE</h2>
    <p>Zdarzenie: ${escHtml(eventName||'________')}</p>
  </div>`;

  if (hospitals.length === 0) {
    html += '<p style="color:#888;font-style:italic">Brak szpitali w zdarzeniu.</p>';
    return html;
  }

  hospitals.forEach(h => {
    const hd = (!sup.skipHospitals && sup.hospitalData && sup.hospitalData[h.id]) ? sup.hospitalData[h.id] : {};
    const f  = k => fill(hd[k]);

    const deptRows = DEPT_LABELS.map((d, i) => {
      const dept = (hd.depts || {})[DEPT_KEYS[i]] || {};
      return `<tr><td>${d}</td><td>${fill(dept.h1)}</td><td>${fill(dept.h2)}</td><td></td></tr>`;
    }).join('');

    html += `<div class="mz-hosp-card-print">
      <div class="mz-hosp-title">${escHtml(h.name)}</div>
      <div class="mz-hosp-meta">
        Nr telefonu: <strong>${fill(hd.phone)}</strong>
        &nbsp;&nbsp;Możliwość przyjęcia — CZERWONEJ: <strong>${h.redCapacity||'___'}</strong>
        &nbsp;ŻÓŁTEJ: <strong>${h.yellowCapacity||'___'}</strong>
      </div>
      <p class="mz-sub-heading">DOSTĘPNOŚĆ (przy wstrzymaniu planowych zabiegów):</p>
      <table class="mz-print-table">
        <thead><tr><th>Zasób</th><th>W ciągu 1 godz.</th><th>W ciągu 2 godz.</th><th>Uwagi</th></tr></thead>
        <tbody>
          <tr><td>Liczba stanowisk ITM (z respiratorem)</td><td>${f('itm1h')}</td><td>${f('itm2h')}</td><td></td></tr>
          <tr><td>Liczba wolnych RESP. (poza stanowiskami ITM)</td><td>${f('resp1h')}</td><td>${f('resp2h')}</td><td></td></tr>
          <tr><td>Liczba stanowisk IN</td><td>${f('in1h')}</td><td>${f('in2h')}</td><td></td></tr>
          <tr><td>Bloki operacyjne</td><td>${f('or1h')}</td><td>${f('or2h')}</td><td></td></tr>
        </tbody>
      </table>
      <p class="mz-sub-heading" style="margin-top:8px">ODDZIAŁY:</p>
      <table class="mz-print-table">
        <thead><tr><th>Oddział</th><th>W ciągu 1 godz.</th><th>W ciągu 2 godz.</th><th>Uwagi</th></tr></thead>
        <tbody>${deptRows}</tbody>
      </table>
    </div>`;
  });

  return html;
}

```

- [ ] **Step 2: Verify via browser console**

With hospitals configured in the app, in console run:
```js
// Test with supplement data
const testSup = {
  skipHospitals: false,
  hospitalData: Object.fromEntries(hospitals.map(h => [h.id, {
    phone: '22-123-4567', itm1h: '5', itm2h: '8',
    resp1h:'2', resp2h:'3', in1h:'10', in2h:'15', or1h:'2', or2h:'3',
    depts: { oit:{h1:'3',h2:'5'}, neurochir:{h1:'1',h2:'2'}, ortopedia:{h1:'2',h2:'3'}, 'chir-ogolna':{h1:'1',h2:'2'}, 'chir-naczyniowa':{h1:'0',h2:'1'}, sor:{h1:'4',h2:'6'} }
  }]))
};
const html13 = genZal13(testSup);
console.assert(html13.includes('ZAŁĄCZNIK NR 13'), 'header present');
console.assert(html13.includes('22-123-4567'), 'phone populated');
console.assert(html13.includes('Chirurgia naczyniowa'), 'dept row present');

// Test with skip
const html13skip = genZal13({ skipHospitals: true, hospitalData: {} });
console.assert(html13skip.includes('________'), 'blanks when skipped');
```

- [ ] **Step 3: Commit**

```bash
git add index.html
git commit -m "feat: implement genZal13 — Tabela Szpitali HTML generator"
```

---

## Task 7: `genZal14()` — Raport z przebiegu zdarzenia

**Files:**
- Modify: `index.html` (JS — add before `collectMzSupplement`)

- [ ] **Step 1: Add `genZal14(sup)` before `collectMzSupplement`**

Find:
```js
function collectMzSupplement() {
```

Insert before it:
```js
function genZal14(sup) {
  const fill = v => (v !== undefined && v !== null && v !== '') ? escHtml(String(v)) : '________';
  const blank = () => '________';

  const counts = { T1:0, T2:0, T3:0, T4:0 };
  patients.forEach(p => counts[p.triage]++);

  // ZRM type counts from supplement
  const zrmTypeCounts = { S:0, P:0, LZRM:0 };
  if (!sup.skipZrmTypes) {
    teams.forEach(tm => { zrmTypeCounts[sup.zrmTypes[tm.id] || 'P']++; });
  }

  // KAM role flags
  const kamRole = sup.skipKam ? '' : (sup.kamRole || '');
  const isLekarz   = kamRole === 'Lekarz'             ? 'TAK' : 'NIE';
  const isRatownik = kamRole === 'Ratownik Medyczny'  ? 'TAK' : 'NIE';
  const isPieleg   = kamRole === 'Pielęgniarka'       ? 'TAK' : 'NIE';

  // Per-hospital patient counts for attachment table
  const byHosp = {};
  patients.filter(p => p.hospitalId).forEach(p => {
    if (!byHosp[p.hospitalId]) byHosp[p.hospitalId] = { T1:0, T2:0, T3:0 };
    if (p.triage !== 'T4') byHosp[p.hospitalId][p.triage]++;
  });

  const genTime = new Date().toLocaleTimeString('pl-PL', { hour:'2-digit', minute:'2-digit' });

  let html = `<div class="mz-doc-header">
    <div class="mz-doc-num">ZAŁĄCZNIK NR 14</div>
    <h2>RAPORT Z PRZEBIEGU ZDARZENIA<br>Z DUŻĄ LICZBĄ POSZKODOWANYCH</h2>
  </div>

  <div class="mz-report-section">
    <div class="mz-report-label">1. Dysponent</div>
    <div class="mz-report-value">${fill(sup.skipGeneral ? '' : sup.dysponent)}</div>
  </div>

  <div class="mz-report-section">
    <div class="mz-report-label">2. Data zdarzenia</div>
    <div class="mz-report-value">
      Data wystąpienia zdarzenia: ${fill(eventTime)}<br>
      Godzina przyjęcia powiadomienia przez DM: ${fill(eventTime)}
    </div>
  </div>

  <div class="mz-report-section">
    <div class="mz-report-label">3. Data zakończenia akcji</div>
    <div class="mz-report-value">${fill(eventEndTime)}</div>
  </div>

  <div class="mz-report-section">
    <div class="mz-report-label">4. GDM</div>
    <div class="mz-report-value">Imię i nazwisko GDM: ${fill(gdmName)} &nbsp;|&nbsp; KOD GDM: ________</div>
  </div>

  <div class="mz-report-section">
    <div class="mz-report-label">5. ZRM</div>
    <div class="mz-report-value">
      <table class="mz-print-table" style="width:auto;margin-bottom:8px">
        <thead><tr><th>Typ ZRM</th><th>Liczba</th></tr></thead>
        <tbody>
          <tr><td>"S"</td><td>${sup.skipZrmTypes ? blank() : zrmTypeCounts.S}</td></tr>
          <tr><td>"P"</td><td>${sup.skipZrmTypes ? blank() : zrmTypeCounts.P}</td></tr>
          <tr><td>"LZRM"</td><td>${sup.skipZrmTypes ? blank() : zrmTypeCounts.LZRM}</td></tr>
        </tbody>
      </table>
      Kryptonimy ZRM biorących udział w zdarzeniu: ${teams.length ? teams.map(t => escHtml(t.name)).join(', ') : blank()}
    </div>
  </div>

  <div class="mz-report-section">
    <div class="mz-report-label">6. KAM</div>
    <div class="mz-report-value">
      <table class="mz-print-table">
        <thead><tr><th>Imię i nazwisko</th><th>Kryptonim ZRM</th><th>Od godz.</th><th>Do godz.</th><th>Lekarz</th><th>Ratownik Med.</th><th>Pielęgniarka</th></tr></thead>
        <tbody><tr>
          <td>${fill(sup.skipKam ? '' : sup.kamNameSup)}</td>
          <td>________</td>
          <td>${fill(sup.skipKam ? '' : sup.kamFrom)}</td>
          <td>${fill(sup.skipKam ? '' : sup.kamTo)}</td>
          <td>${sup.skipKam ? blank() : isLekarz}</td>
          <td>${sup.skipKam ? blank() : isRatownik}</td>
          <td>${sup.skipKam ? blank() : isPieleg}</td>
        </tr></tbody>
      </table>
    </div>
  </div>

  <div class="mz-report-section">
    <div class="mz-report-label">7. Zestawienie poszkodowanych uczestniczących w zdarzeniu</div>
    <div class="mz-report-value">
      <strong>Łączna liczba poszkodowanych: ${patients.length}</strong>
      <table class="mz-print-table" style="margin-top:8px">
        <thead><tr><th>Kategoria</th><th>Liczba</th></tr></thead>
        <tbody>
          <tr><td>Drogą powietrzną</td><td>${sup.skipTransport ? blank() : sup.transportAir}</td></tr>
          <tr><td>Drogą naziemną przez ZRM</td><td>${sup.skipTransport ? blank() : sup.transportGroundZrm}</td></tr>
          <tr><td>Drogą naziemną bez ZRM (inne środki transportu)</td><td>${sup.skipTransport ? blank() : sup.transportGroundOther}</td></tr>
          <tr><td>Drogą wodną</td><td>${sup.skipTransport ? blank() : sup.transportWater}</td></tr>
          <tr><td>Udzielono pomocy na miejscu, brak transportu do SOR/IP/JOS</td><td>${sup.skipTransport ? blank() : sup.transportOnScene}</td></tr>
          <tr><td>Nie wyrazili zgody na transport do SOR/IP/JOS</td><td>${sup.skipTransport ? blank() : sup.transportNoConsent}</td></tr>
          <tr><td>Odmówili udzielenia pomocy (brak zgody na badanie)</td><td>${sup.skipTransport ? blank() : sup.transportRefused}</td></tr>
          <tr><td>Samodzielnie oddalili się z miejsca zdarzenia</td><td>${sup.skipTransport ? blank() : sup.transportLeft}</td></tr>
        </tbody>
      </table>
    </div>
  </div>

  <div class="mz-report-section">
    <div class="mz-report-label">8. Wysłanie do WKRM</div>
    <div class="mz-report-value">
      Raport przesłany: ${sup.skipGeneral ? blank() : (sup.sentToWkrm ? 'TAK' : 'NIE')}
      &nbsp;|&nbsp; do WKRM województwa: ${fill(sup.skipGeneral ? '' : sup.wkrmProvince)}
    </div>
  </div>

  <div class="mz-report-section">
    <div class="mz-report-label">9. Autor raportu</div>
    <div class="mz-report-value">
      Raport przygotował: ${fill(sup.skipGeneral ? '' : sup.reportAuthor)}
      &nbsp;|&nbsp; Godzina: ${genTime}
    </div>
  </div>

  <div class="mz-report-section">
    <div class="mz-report-label">Załącznik — Zestawienie poszkodowanych per podmiot leczniczy</div>
    <div class="mz-report-value">
      <table class="mz-print-table">
        <thead><tr><th>Podmiot leczniczy</th><th>Czerwoni (T1)</th><th>Żółci (T2)</th><th>Zieloni (T3)</th><th>Łącznie</th></tr></thead>
        <tbody>${hospitals.map(h => {
          const hd = byHosp[h.id] || { T1:0, T2:0, T3:0 };
          return `<tr>
            <td>${escHtml(h.name)}</td>
            <td>${hd.T1}</td><td>${hd.T2}</td><td>${hd.T3}</td>
            <td><strong>${hd.T1 + hd.T2 + hd.T3}</strong></td>
          </tr>`;
        }).join('')}</tbody>
      </table>
      <p style="margin-top:6px">
        Podsumowanie segregacji: 🔴 T1=${counts.T1} &nbsp; 🟡 T2=${counts.T2} &nbsp; 🟢 T3=${counts.T3} &nbsp; ⬛ T4=${counts.T4}
      </p>
    </div>
  </div>`;

  return html;
}

```

- [ ] **Step 2: Verify via browser console**

With an event (patients, teams, hospitals, eventTime, eventEndTime, gdmName, kamName set), run:
```js
const testSup = {
  skipGeneral:false, dysponent:'WSRM Warszawa', wkrmProvince:'mazowieckie',
  sentToWkrm:true, reportAuthor:'Jan Kowalski',
  skipZrmTypes:false, zrmTypes: Object.fromEntries(teams.map(t=>[t.id,'P'])),
  skipKam:false, kamNameSup:kamName, kamRole:'Ratownik Medyczny', kamFrom:eventTime, kamTo:eventEndTime,
  skipTransport:false, transportAir:0, transportGroundZrm:2, transportGroundOther:0,
  transportWater:0, transportOnScene:0, transportNoConsent:0, transportRefused:0, transportLeft:0,
  skipHospitals:false, hospitalData:{}
};
const html14 = genZal14(testSup);
console.assert(html14.includes('ZAŁĄCZNIK NR 14'), 'header');
console.assert(html14.includes('WSRM Warszawa'), 'dysponent');
console.assert(html14.includes('mazowieckie'), 'wkrm province');
console.assert(html14.includes('Ratownik Medyczny'), 'KAM role');
console.assert(html14.includes('TAK'), 'sent to WKRM');

// Test skip
const suppSkipAll = { skipGeneral:true, skipZrmTypes:true, skipKam:true, skipTransport:true, skipHospitals:true, zrmTypes:{}, hospitalData:{}, sentToWkrm:false };
const html14s = genZal14(suppSkipAll);
const blankCount = (html14s.match(/________/g)||[]).length;
console.assert(blankCount > 5, 'many blanks when skipped');
```

- [ ] **Step 3: Commit**

```bash
git add index.html
git commit -m "feat: implement genZal14 — Raport z przebiegu zdarzenia HTML generator"
```

---

## Task 8: `generateMzHtml()` + `downloadMzDocs()` — wire up and download

**Files:**
- Modify: `index.html` (JS — replace `downloadMzDocs` stub, add `generateMzHtml`)

- [ ] **Step 1: Add `generateMzHtml(sup)` and replace `downloadMzDocs`**

Find:
```js
function downloadMzDocs() {
  const sup = collectMzSupplement();
  console.log('[MZ] supplement collected:', sup);
  closeMzModal();
  alert('TODO: generateMzHtml(sup)');
}
```

Replace with:
```js
function generateMzHtml(sup) {
  const dateStr = new Date().toLocaleDateString('pl-PL');
  const css = `
*{box-sizing:border-box;margin:0;padding:0}
body{font-family:Arial,Helvetica,sans-serif;font-size:12px;color:#111;background:#fff}
.doc-page{padding:16mm 14mm}
.page-break{page-break-before:always}
@media print{.page-break{page-break-before:always}}
@page{size:A4;margin:0}
.mz-doc-header{margin-bottom:14px;border-bottom:2px solid #111;padding-bottom:8px}
.mz-doc-num{font-size:11px;color:#555;letter-spacing:.08em;font-weight:600;margin-bottom:2px}
.mz-doc-header h2{font-size:15px;font-weight:700;margin:2px 0}
.mz-doc-header p{font-size:11px;color:#444;margin-top:4px}
.mz-triage-table{width:100%;border-collapse:collapse;margin-bottom:14px;font-size:11px}
.mz-triage-table th,.mz-triage-table td{border:1px solid #ccc;padding:4px 6px;vertical-align:top}
.mz-triage-col-header th{background:#f0f0f0;font-weight:600}
.mz-triage-table tbody tr:nth-child(even){background:#fafafa}
.mz-hosp-card-print{border:1px solid #aaa;border-radius:3px;padding:10px;margin-bottom:10px;page-break-inside:avoid}
.mz-hosp-title{font-size:13px;font-weight:700;margin-bottom:4px}
.mz-hosp-meta{font-size:11px;color:#444;margin-bottom:6px}
.mz-sub-heading{font-size:11px;font-weight:700;margin:6px 0 3px}
.mz-print-table{width:100%;border-collapse:collapse;font-size:11px}
.mz-print-table th,.mz-print-table td{border:1px solid #ccc;padding:4px 8px;text-align:left}
.mz-print-table th{background:#f0f0f0;font-weight:600}
.mz-report-section{margin-bottom:10px;border:1px solid #ddd;border-radius:3px;overflow:hidden;page-break-inside:avoid}
.mz-report-label{background:#1a1a2e;color:#fff;font-weight:700;font-size:11px;padding:5px 10px;letter-spacing:.04em}
.mz-report-value{padding:8px 10px;font-size:11px;line-height:1.5}
`;

  return `<!DOCTYPE html>
<html lang="pl">
<head>
  <meta charset="UTF-8">
  <title>Dokumenty MZ — ${escHtml(eventName||'Zdarzenie')} — ${dateStr}</title>
  <style>${css}</style>
</head>
<body>
  <div class="doc-page">${genZal12()}</div>
  <div class="doc-page page-break">${genZal13(sup)}</div>
  <div class="doc-page page-break">${genZal14(sup)}</div>
</body>
</html>`;
}

function downloadMzDocs() {
  const sup = collectMzSupplement();
  const html = generateMzHtml(sup);
  const dateStr = new Date().toISOString().slice(0, 10);
  const blob = new Blob([html], { type: 'text/html;charset=utf-8' });
  const a = document.createElement('a');
  a.href = URL.createObjectURL(blob);
  a.download = `dokumenty-mci-${dateStr}.html`;
  document.body.appendChild(a);
  a.click();
  document.body.removeChild(a);
  URL.revokeObjectURL(a.href);
  closeMzModal();
}
```

- [ ] **Step 2: Verify full flow in browser**

1. Open `index.html`, start an event named "Wypadek A4 km 312", add one T1 patient with injuries (Głowa: złamanie), assign ZRM "W-01" and hospital "Szpital Bródnowski". Add one T4 patient with no team/hospital.
2. End the event. Open the modal. Fill:
   - Dysponent: "WSRM w Warszawie"
   - WKRM: "mazowieckie"
   - Author: "Jan Kowalski"
3. Skip Section 2 (ZRM types) and Section 5 (hospitals).
4. Click "⬇ Generuj i pobierz".
5. File `dokumenty-mci-YYYY-MM-DD.html` must download.
6. Open the downloaded file in browser. Verify:
   - **Page 1 (Zał. 12):** "ZAŁĄCZNIK NR 12", "CZERWONI", patient nr 1 listed, "złamanie" in injury column, "W-01" in ZRM column, "CZARNI" section has only 3 columns (no ZRM/Szpital).
   - **Page 2 (Zał. 13):** "ZAŁĄCZNIK NR 13", hospital name appears, fields show "________" (sections were skipped).
   - **Page 3 (Zał. 14):** "ZAŁĄCZNIK NR 14", "WSRM w Warszawie" in Sekcja 1, "mazowieckie" in Sekcja 8, "Jan Kowalski" in Sekcja 9, ZRM type counts show "________" (skipped).

- [ ] **Step 3: Test print layout**

1. In the downloaded file, open browser Print dialog (Ctrl+P).
2. Select "Save as PDF".
3. Verify page breaks appear between the 3 documents.
4. Verify tables don't overflow page margins.

- [ ] **Step 4: Commit**

```bash
git add index.html
git commit -m "feat: implement generateMzHtml + downloadMzDocs — complete MZ document generation and download"
```

---

## Self-Review Checklist

After completing all tasks, verify against the spec:

- [ ] Zał. 12: 4 triage groups, sequential global numbering, T4 has no ZRM/Szpital columns
- [ ] Zał. 13: all hospitals from `hospitals[]`, capacity auto-filled, supplement fields or blanks
- [ ] Zał. 14: all 9 sekcje present, attachment table per hospital, genTime at generation moment
- [ ] Supplement modal: 5 sections, each skippable, KAM pre-filled from app state
- [ ] Button only visible when `eventEndTime` is set, hidden otherwise
- [ ] Downloaded file is self-contained HTML (no external URLs)
- [ ] `transportGroundZrm = max(0, transported - air - groundOther - water)` ✓
- [ ] Hospital section skip: all hospital supplement fields show `________`

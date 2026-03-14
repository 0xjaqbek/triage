# PWA Conversion Implementation Plan

> **For Claude:** REQUIRED SUB-SKILL: Use superpowers:executing-plans to implement this plan task-by-task.

**Goal:** Convert TRIAGE MCI into an installable PWA with offline support and localStorage persistence on iOS, Android, and Windows.

**Architecture:** Add `manifest.json`, `sw.js`, and icon PNGs alongside existing `index.html`. Modify `index.html` to remove Google Fonts, add PWA meta tags, register service worker, implement localStorage save/restore, and add confirmation dialogs for destructive actions.

**Tech Stack:** Vanilla JS, Service Worker API, Web App Manifest, localStorage

---

### Task 1: Generate PWA Icon Files

**Files:**
- Source: `triage.png`
- Create: `icons/icon-192.png`
- Create: `icons/icon-512.png`
- Create: `icons/apple-touch-icon.png`

**Step 1: Create icons directory**

```bash
mkdir -p icons
```

**Step 2: Generate PNG icons from triage.png**

Use ImageMagick (or `sharp` via Node) to resize `triage.png` to three sizes:
- `icons/icon-192.png` — 192x192
- `icons/icon-512.png` — 512x512
- `icons/apple-touch-icon.png` — 180x180

```bash
magick triage.png -resize 192x192 icons/icon-192.png
magick triage.png -resize 512x512 icons/icon-512.png
magick triage.png -resize 180x180 icons/apple-touch-icon.png
```

If ImageMagick is not available, create a small Node script:

```bash
npm init -y && npm install sharp
node -e "
const sharp = require('sharp');
const sizes = [[192,'icon-192'],[512,'icon-512'],[180,'apple-touch-icon']];
sizes.forEach(([s,n]) => sharp('triage.png').resize(s,s).png().toFile('icons/'+n+'.png'));
"
```

**Step 3: Commit**

```bash
git add icons/
git commit -m "feat: add PWA icon files in 192, 512, and 180px sizes"
```

---

### Task 2: Create Web App Manifest

**Files:**
- Create: `manifest.json`

**Step 1: Create manifest.json**

```json
{
  "name": "TRIAGE MCI – System Zarządzania Zdarzeniem Masowym",
  "short_name": "TRIAGE",
  "description": "System segregacji i zarządzania zdarzeniem masowym",
  "start_url": "./index.html",
  "display": "standalone",
  "background_color": "#0C0F14",
  "theme_color": "#0C0F14",
  "orientation": "portrait-primary",
  "icons": [
    {
      "src": "icons/icon-192.png",
      "sizes": "192x192",
      "type": "image/png",
      "purpose": "any maskable"
    },
    {
      "src": "icons/icon-512.png",
      "sizes": "512x512",
      "type": "image/png",
      "purpose": "any maskable"
    }
  ]
}
```

**Step 2: Commit**

```bash
git add manifest.json
git commit -m "feat: add PWA web app manifest"
```

---

### Task 3: Create Service Worker

**Files:**
- Create: `sw.js`

**Step 1: Create sw.js with cache-first strategy**

```javascript
const CACHE_NAME = 'triage-pwa-v1';
const ASSETS = [
  './',
  './index.html',
  './manifest.json',
  './icons/icon-192.png',
  './icons/icon-512.png',
  './icons/apple-touch-icon.png'
];

// Install: precache all assets
self.addEventListener('install', event => {
  event.waitUntil(
    caches.open(CACHE_NAME)
      .then(cache => cache.addAll(ASSETS))
      .then(() => self.skipWaiting())
  );
});

// Activate: remove old caches
self.addEventListener('activate', event => {
  event.waitUntil(
    caches.keys().then(keys =>
      Promise.all(keys.filter(k => k !== CACHE_NAME).map(k => caches.delete(k)))
    ).then(() => self.clients.claim())
  );
});

// Fetch: cache-first, fallback to network
self.addEventListener('fetch', event => {
  event.respondWith(
    caches.match(event.request)
      .then(cached => cached || fetch(event.request))
  );
});
```

**Step 2: Commit**

```bash
git add sw.js
git commit -m "feat: add service worker with cache-first offline strategy"
```

---

### Task 4: Update index.html — Remove Google Fonts

**Files:**
- Modify: `index.html:8-9` (remove font links)
- Modify: `index.html:30` (update --font variable)

**Step 1: Remove Google Fonts preconnect and stylesheet links**

Delete lines 8-9:
```html
<link rel="preconnect" href="https://fonts.googleapis.com">
<link href="https://fonts.googleapis.com/css2?family=JetBrains+Mono:wght@400;500;600;700;800&display=swap" rel="stylesheet">
```

**Step 2: Update CSS --font variable**

Change line 30 from:
```css
--font: 'JetBrains Mono', 'Fira Code', 'SF Mono', monospace;
```
To:
```css
--font: ui-monospace, 'SF Mono', 'Cascadia Code', 'Segoe UI Mono', Consolas, 'Liberation Mono', monospace;
```

**Step 3: Commit**

```bash
git add index.html
git commit -m "feat: replace Google Fonts with system monospace stack for offline"
```

---

### Task 5: Update index.html — Add PWA Meta Tags and SW Registration

**Files:**
- Modify: `index.html:7` (after `<title>`, add manifest/meta/icons)
- Modify: `index.html` (end of `<script>`, add SW registration)

**Step 1: Add PWA meta tags in `<head>` after the `<title>` tag (line 7)**

Insert after `<title>` line:
```html
<link rel="manifest" href="manifest.json">
<meta name="theme-color" content="#0C0F14">
<meta name="apple-mobile-web-app-capable" content="yes">
<meta name="apple-mobile-web-app-status-bar-style" content="black-translucent">
<meta name="apple-mobile-web-app-title" content="TRIAGE">
<link rel="apple-touch-icon" href="icons/apple-touch-icon.png">
```

**Step 2: Add service worker registration at end of `<script>` block**

Before the closing `</script>` tag (line 1718), add:
```javascript
// ===== SERVICE WORKER =====
if ('serviceWorker' in navigator) {
  navigator.serviceWorker.register('./sw.js');
}
```

**Step 3: Commit**

```bash
git add index.html
git commit -m "feat: add PWA manifest link, iOS meta tags, and service worker registration"
```

---

### Task 6: Add localStorage Persistence

**Files:**
- Modify: `index.html` — add `saveState()` and `loadState()` functions in `<script>`
- Modify: `index.html` — call `saveState()` from all mutating functions
- Modify: `index.html` — call `loadState()` on page load

**Step 1: Add saveState/loadState functions**

Insert after the STATE section (after line 1093, before the `now()` function):

```javascript
// ===== PERSISTENCE =====
function saveState() {
  try {
    localStorage.setItem('triageAppState', JSON.stringify({
      eventName, eventTime, nextId, patients, transports, teams, hospitals, currentView
    }));
  } catch(e) {}
}

function loadState() {
  try {
    const raw = localStorage.getItem('triageAppState');
    if (!raw) return false;
    const s = JSON.parse(raw);
    eventName = s.eventName || '';
    eventTime = s.eventTime || '';
    nextId = s.nextId || 1;
    patients = s.patients || [];
    transports = s.transports || [];
    teams = s.teams || [];
    hospitals = s.hospitals || [];
    currentView = s.currentView || 'triage';
    return true;
  } catch(e) { return false; }
}

function clearState() {
  localStorage.removeItem('triageAppState');
}
```

**Step 2: Call saveState() in every mutating function**

Add `saveState();` at the end of each of these functions (before the closing `}`):
- `startEvent()` — after `renderAll();` (line ~1135)
- `addPatient()` — after `renderAll();` (line ~1305)
- `changeTriage()` — after `renderAll();` (line ~1311)
- `updateNotes()` — after `if (p) p.notes = val;` (line ~1316)
- `removePatient()` — after `renderAll();` (line ~1321)
- `markDelivered()` — after `renderAll();` (line ~1327)
- `dispatchTransport()` — after `renderAll();` (line ~1355)
- `addTeam()` — after `renderAll();` (line ~1371)
- `removeTeam()` — after `renderAll();` (line ~1376)
- `addHospital()` — after `renderAll();` (line ~1384)
- `removeHospital()` — after `renderAll();` (line ~1389)
- `switchView()` — after `renderAll();` (line ~1154)

**Step 3: Add load-on-startup logic**

Replace the inline event listener at line 1142 with an initialization block that checks for saved state:

At the bottom of the `<script>` block (before `</script>`), add:

```javascript
// ===== INIT =====
(function init() {
  if (loadState() && eventName) {
    // Saved event found — ask user to resume or start fresh
    const setup = document.getElementById('setup');
    const resumeDiv = document.createElement('div');
    resumeDiv.className = 'resume-prompt';
    resumeDiv.innerHTML = `
      <div style="margin-top:16px;padding:16px;background:var(--surface);border:1px solid var(--border2);border-radius:10px;">
        <div style="color:var(--yellow-light);font-weight:700;margin-bottom:8px;">⚠ Wykryto zapisane zdarzenie</div>
        <div style="color:var(--text2);font-size:12px;margin-bottom:12px;">${escHtml(eventName)} — rozpoczęte ${escHtml(eventTime)}</div>
        <div style="display:flex;gap:8px;">
          <button class="btn-start" onclick="resumeEvent()" style="flex:1;font-size:12px;padding:10px;">KONTYNUUJ</button>
          <button class="btn-start" onclick="freshEvent()" style="flex:1;font-size:12px;padding:10px;background:var(--border);color:var(--text2);">NOWE ZDARZENIE</button>
        </div>
      </div>
    `;
    setup.appendChild(resumeDiv);
  }
})();
```

**Step 4: Add resumeEvent() and freshEvent() functions**

```javascript
function resumeEvent() {
  document.getElementById('setup').classList.add('hidden');
  document.getElementById('btnDownload').classList.add('hidden');
  document.getElementById('app').classList.remove('hidden');
  updateHeaderTime();
  setInterval(updateHeaderTime, 1000);
  // Restore active view
  switchView(currentView);
  renderAll();
}

function freshEvent() {
  if (!confirm('Czy na pewno chcesz rozpocząć nowe zdarzenie? Wszystkie zapisane dane zostaną usunięte.')) return;
  clearState();
  // Reset all state to defaults
  eventName = '';
  eventTime = '';
  nextId = 1;
  patients = [];
  transports = [];
  teams = [
    { id: 1, name: 'S-01 (specjalistyczny)' },
    { id: 2, name: 'S-02 (specjalistyczny)' },
    { id: 3, name: 'P-01 (podstawowy)' },
    { id: 4, name: 'P-02 (podstawowy)' },
    { id: 5, name: 'P-03 (podstawowy)' },
    { id: 6, name: 'LPR (śmigłowiec)' },
  ];
  hospitals = [
    { id: 1, name: 'SOR Szpital Wojewódzki' },
    { id: 2, name: 'SOR Szpital Miejski' },
    { id: 3, name: 'SOR Szpital Kliniczny' },
    { id: 4, name: 'CU Szpital Powiatowy' },
  ];
  currentView = 'triage';
  // Remove resume prompt and reload setup
  location.reload();
}
```

**Step 5: Commit**

```bash
git add index.html
git commit -m "feat: add localStorage persistence with save/load/clear state"
```

---

### Task 7: Add Confirmation Dialogs for Destructive Actions

**Files:**
- Modify: `index.html` — `removePatient()`, `removeTeam()`, `removeHospital()`

**Step 1: Add confirmation to removePatient()**

Wrap existing logic:
```javascript
function removePatient(id) {
  if (!confirm('Czy na pewno chcesz usunąć tego pacjenta?')) return;
  patients = patients.filter(x => x.id !== id);
  renderAll();
  saveState();
}
```

**Step 2: Add confirmation to removeTeam()**

```javascript
function removeTeam(id) {
  if (!confirm('Czy na pewno chcesz usunąć ten zespół?')) return;
  teams = teams.filter(x => x.id !== id);
  renderAll();
  saveState();
}
```

**Step 3: Add confirmation to removeHospital()**

```javascript
function removeHospital(id) {
  if (!confirm('Czy na pewno chcesz usunąć ten szpital?')) return;
  hospitals = hospitals.filter(x => x.id !== id);
  renderAll();
  saveState();
}
```

**Step 4: Commit**

```bash
git add index.html
git commit -m "feat: add confirmation dialogs before destructive actions"
```

---

### Task 8: Manual Testing

**Step 1: Serve locally and verify**

```bash
npx serve . -p 8080
```

Open `http://localhost:8080` in Chrome and verify:
- [ ] App loads without Google Fonts (system monospace font visible)
- [ ] Service worker registers (DevTools > Application > Service Workers)
- [ ] Manifest is detected (DevTools > Application > Manifest)
- [ ] Icons display correctly in manifest
- [ ] "Install app" prompt appears in Chrome address bar
- [ ] Create an event, add patients — reload page — resume prompt appears
- [ ] Click "KONTYNUUJ" — all data restored correctly
- [ ] Click "NOWE ZDARZENIE" — confirmation dialog appears
- [ ] Confirm — data wiped, fresh setup screen
- [ ] Go offline (DevTools > Network > Offline) — app still works
- [ ] Remove patient — confirmation dialog appears
- [ ] Remove team — confirmation dialog appears
- [ ] Remove hospital — confirmation dialog appears
- [ ] All existing features work: triage wizard, dispatch, report, clipboard copy, download

**Step 2: Final commit (if any fixes needed)**

```bash
git add -A
git commit -m "fix: address issues found during PWA testing"
```

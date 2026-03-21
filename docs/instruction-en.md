# TRIAGE MCI User Manual

## 1. What is TRIAGE MCI

TRIAGE MCI is a web application (PWA) for triage and mass casualty incident (MCI) management based on the START algorithm. It allows you to:
- Classify patients into 4 categories (T1–T4)
- Manage transport to hospitals
- Track hospital capacity
- Generate and send reports
- Import data from a dispatcher

The application works fully offline. All data is stored exclusively on the user's device — no information is sent to external servers.

It supports 7 languages: Polish, English, Italian, French, German, Czech, and Portuguese.

---

## 2. Starting an Incident

### Start Screen

When you launch the application, the incident configuration screen appears:

1. **Language selection** — selector at the top of the screen. Changing the language immediately translates the entire interface.
2. **Incident name** — required field. Enter the location and type of incident (e.g., "Crash A4 km 312"). This will appear in reports.
3. **MCI Commander** — optional field. Medical Action Commander — the name of the person coordinating on-scene operations.
4. **Show hints** — toggle that enables contextual hint bubbles throughout the application. Useful for first-time use.
5. **START TRIAGE** — starts the incident and launches the timer.

### Resuming an Incident

If a saved incident exists in device memory, the following options appear on launch:
- **CONTINUE** — resumes the saved incident with all data
- **NEW INCIDENT** — clears previous data and starts from scratch

---

## 3. Triage (START Algorithm)

### Navigation

The application has 3 tabs at the top of the screen:
- **TRIAGE** — patient classification
- **DISPATCH** — transport and management
- **REPORT** — summary and export

The statistics bar below the navigation shows the number of patients in each category: 🔴 T1, 🟡 T2, 🟢 T3, ⚫ T4, and the total.

### START Triage Wizard

The wizard guides you through the 6-step START algorithm. At each step, you answer YES or NO:

**Step 1 — Can the patient walk?**
- YES → T3 Green (delayed)
- NO → proceed to step 2

**Step 2 — Is the patient breathing spontaneously?**
- YES → proceed to step 4 (respiratory rate assessment)
- NO → proceed to step 3

**Step 3 — Open the airway. Is the patient breathing after airway opening?**
- YES → T1 Red (immediate)
- NO → T4 Black (deceased)

**Step 4 — Is the respiratory rate normal (≤ 30/min)?**
- YES → proceed to step 5
- NO → T1 Red (immediate)

**Step 5 — Is capillary refill normal (≤ 2 sec) and radial pulse present?**
- YES → proceed to step 6
- NO → T1 Red (immediate)

**Step 6 — Can the patient follow simple commands?**
- YES → T2 Yellow (urgent)
- NO → T1 Red (immediate)

Each question has a hint displayed below it — describing how to assess the given parameter.

### Result Card

After completing triage, a result card appears with:

- **Color tag** — a large square with the category (T1/T2/T3/T4)
- **Decision path** — shows the algorithm flow
- **Sex** — buttons M (male), F (female), ? (unknown)
- **Age** — numeric field with buttons -5, -1, +1, +5 (range 0–150)
- **Body injuries** — button to open the injury diagram (see section 4)
- **Notes** — field for additional information (injuries, mechanism of injury, etc.)

Available actions:
- **✓ CONFIRM PATIENT** — adds the patient to the list
- **↺ START OVER** — resets the wizard to step 1
- **CHANGE COLOR MANUALLY** — opens a row of 4 buttons (T1–T4) for manual category override

### On-Scene Patient List

Below the wizard, the "On-scene patients" list is displayed, sorted by priority (T1 → T2 → T3 → T4).

Each patient card contains:
- Color tag (e.g., P-001)
- Category, sex/age, registration time
- Notes field (editable)
- Injury button (🩻) with summary
- 4 retriage dots (T1–T4) — click to change category
- Delete button (✕)
- Retriage badge (↺n) if the category was changed

### In-Transit List

The "In transit" section appears when patients are en route to a hospital. Each card shows:
- Patient tag with category color
- Route: Team → Hospital → departure time
- "change" button — change the destination hospital
- "✓ Delivered" button — marks delivery and releases the EMS team

---

## 4. Body Injury Diagram

The "🩻 Body injuries" button opens an interactive body diagram.

### Diagram View

Two body outlines are displayed side by side:
- **FRONT** — head, chest, abdomen, arms, legs
- **BACK** — head, upper back, lower back, arms, legs

Side labels are visible: **L** (left) and **R** (right), reversed for the back view.

Click on a body zone to select it.

### Zone Injury Panel

After selecting a zone, a panel opens with 7 injury types to mark:
- Fracture
- Bleeding
- Burn
- Wound
- Swelling
- Pain
- Amputation

Checked types are marked with ✓. Each can be toggled on/off by clicking.

At the bottom of the panel:
- **✕ NO** — discards changes for this zone
- **✓ YES** — confirms injuries for the selected zone

### Confirming the Diagram

After marking injuries in all needed zones:
- **✓ YES** — confirms the entire injury diagram
- **✕ NO** — discards all changes and reverts to the state before opening

Marked injuries are visible:
- On the patient card as an abbreviated summary (e.g., "🩻 Head: Fracture, Bleeding")
- In the report for the given patient

---

## 5. Dispatch

### Configuration Modals

When you first enter the Dispatch tab, 3 modals appear in sequence:

**1. Medical Operations Director (GDM)**
- Field for GDM name
- Button to import data from the dispatcher (see section 6)
- SKIP or SAVE

**2. EMS Teams**
- Question: "Do you know the EMS teams dispatched to the scene?"
- YES → form for adding teams (name + add button)
- NO → use default teams (S-01, S-02, P-01, P-02, P-03, LPR)
- CONFIRM or SKIP (DEFAULT)

**3. Destination Hospitals**
- Form for adding hospitals with capacity:
  - Hospital name
  - 🔴 Red capacity (T1) — number of beds
  - 🟡 Yellow capacity (T2) — number of beds
- CONFIRM or SKIP (DEFAULT)

Each modal appears only once per incident.

### Dispatch Form

The "Dispatch transport" section contains:

1. **Patient** — dropdown list of on-scene patients, sorted by priority
2. **EMS Team** — list of available teams (busy teams are hidden)
3. **Hospital** — list of hospitals with occupancy information:
   - Format: "Hospital name [🔴 occupied/total 🟡 occupied/total]"
   - Hospitals with no available beds for the patient's category are flagged with a warning
4. **DISPATCH TRANSPORT** — button active when all fields are filled

### Management Panels

Below the form are two panels (side by side on wide screens, stacked on narrow screens):

**Hospitals** (ALL beds):
- List of hospitals with editable capacity fields (🔴 and 🟡)
- Adding new hospitals
- Removing hospitals (✕)
- Collapsing/expanding the list

**EMS Teams:**
- List of teams with delete buttons
- Adding new teams
- Collapsing/expanding the list

### Transport History

At the bottom of the dispatch section (collapsed by default). Expand to see a chronological list of all transports with:
- Patient tag with color
- Team → Hospital → departure time
- Status (in transit / delivered)

---

## 6. Importing Data from the Dispatcher

This feature allows the dispatcher to send data (incident name, GDM, EMS teams, hospitals with capacity) to the responder in the field.

### Method 1 — SMS Link (one click)

The dispatcher sends an SMS with a link in this format:
```
https://0xjaqbek.github.io/triage/?i=BASE64_DATA
```

After clicking the link:
1. The application opens from cache (no internet required if installed)
2. A data preview appears: incident name, GDM, teams, hospitals
3. **IMPORT** — imports the data and opens triage
4. **DISCARD** — ignores the data

### Method 2 — Pasting Data

If the link does not open the application (locked devices):
1. Open the application, go to Dispatch
2. In the GDM modal, click **📥 IMPORT DATA FROM DISPATCHER**
3. Paste the SMS content (the full link or just the Base64 code)
4. The data preview appears automatically
5. **IMPORT** — confirms the data

### Import During an Active Incident

If the incident is already in progress and you have triaged patients:
- **Patients** — preserved (never deleted)
- **MCI Commander** — preserved
- **GDM** — replaced with dispatcher data
- **EMS Teams** — replaced with dispatcher data
- **Hospitals** — replaced with dispatcher data
- **Incident name** — merged: "Your name | Dispatcher name (dispatcher)"

A yellow warning in the preview informs you that patients and MCI Commander are preserved.

---

## 7. Dispatcher Mode

Accessible via the "Dispatcher mode →" link on the start screen or at `/dyspozytor/`.

This page allows the dispatcher to prepare and send data to the field team.

### Form

1. **Incident name** — e.g., "MCI Krakow-Nowa Huta"
2. **Medical Operations Director (GDM)** — full name
3. **EMS Teams** — add one at a time (name + "+" button)
4. **Destination hospitals** — name + 🔴 and 🟡 capacity (+ add button)

### Sending

- **📋 Preview data** — shows a summary with character count and number of SMS segments
- **📱 Send SMS** — opens the SMS app with a ready-made link
- **📋 Copy data** — copies the link to the clipboard

Data is encoded in Base64 (compatible with GSM-7 SMS — up to 1,530 characters in 10 SMS segments). This can accommodate approximately 20 hospitals and 20 teams.

---

## 8. Report

### Visual Report

The "Report for dispatcher" section is generated automatically and contains:

**Header:**
- Incident name, start time, report time
- MCI Commander and GDM (if provided)
- End time (if the incident is closed)

**Summary:**
- Color tags with patient counts for each category
- Total patient count

**Hospitals** (separate block for each hospital with transports):
- Hospital name, patient count, capacity info
- Patient list with: tag, age/sex, team, time, status (IN TRANSIT / DELIVERED)
- Retriage history (e.g., ↺T1→T2→T1)
- Hospital change history (e.g., ⇄Hospital1→Hospital2)
- Injury summary

**On-scene patients** (if any remain):
- List with: tag, age/sex, category, notes, injuries

### Text Report

The "📋 Text report" button opens a modal with the full report in text format. It includes a complete patient list with notes, sex, age, and triage pathway.

### Sending the Report

- **📋 Copy to clipboard** — copies the full report text
- **📤 Send** — opens options:
  - **SMS** — opens the native SMS app with the report content
  - **Email** — opens the email client with the subject and report content

### Closing the Incident

The **END INCIDENT** button in the report header:
- Saves the end time
- Changes to **UNDO CLOSING** (reversible)
- The end time appears in the report

---

## 9. Retriage and Hospital Change

### Retriage

On the patient card in the "On-scene patients" section, there are 4 colored dots (T1–T4). Clicking a different category than the current one:
1. Opens a confirmation dialog: "Change patient P-001 category from T1 to T2?"
2. **CONFIRM** — changes the category, saves to history
3. **CANCEL** — no changes

Retriage history:
- Badge ↺n on the patient card (n = number of changes)
- Full history in the report (e.g., T1→T2→T1 with timestamps)

### Changing Hospital During Transport

On the patient card in "In transit," the "change" button:
1. Opens a modal with the current hospital
2. Dropdown list of available hospitals (with capacity info)
3. **CONFIRM** — changes the hospital, saves to history
4. **CANCEL** — no changes

Hospital change history is visible in the report (⇄Hospital1→Hospital2).

---

## 10. Installation and Offline Mode

### Installing as a PWA

**Android (Chrome):**
1. Open the application in Chrome
2. Tap ⋮ (menu) in the upper right corner
3. Select "Add to Home screen"
4. Tap "Install"

**iOS (Safari):**
1. Open the application in Safari
2. Tap ⎙ (Share)
3. Scroll down and select "Add to Home Screen"

After installation, the application runs as a standalone app with full offline mode.

### Offline Mode

After the first load, the application is saved in the device's cache. All features work without an internet connection:
- Patient triage
- Transport dispatch
- Report generation
- Data import (pasting)

Internet is only needed for: the initial download, updates, and sending SMS/email (via the device's native apps).

### Download as HTML File

The **⬇ Download** button on the start screen downloads the application as a standalone HTML file. It can be transferred to a USB drive and opened on any computer with a browser.

---

## 11. Privacy

- All data is stored **exclusively on the device** (localStorage)
- No data is **sent** to external servers
- No analytics, tracking, or telemetry
- No tracking cookies
- SMS and email features use the device's native apps — the application does not send messages directly
- To delete data: use "End incident" or clear browser data

The full privacy policy is available at `/privacy/`.

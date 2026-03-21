# Bedienungsanleitung TRIAGE MCI

## 1. Was ist TRIAGE MCI

TRIAGE MCI ist eine Webanwendung (PWA) zur Sichtung und zum Management von Massenanfällen von Verletzten (MCI), basierend auf dem START-Algorithmus. Sie ermöglicht:
- Klassifizierung von Patienten in 4 Kategorien (T1–T4)
- Transportmanagement zu Krankenhäusern
- Überwachung der Krankenhauskapazitäten
- Erstellung und Versand von Berichten
- Import von Daten des Disponenten

Die Anwendung funktioniert vollständig offline. Alle Daten werden ausschließlich auf dem Gerät des Benutzers gespeichert — es werden keine Informationen an externe Server gesendet.

Sie unterstützt 7 Sprachen: Polnisch, Englisch, Italienisch, Französisch, Deutsch, Tschechisch, Portugiesisch.

---

## 2. Ereignis starten

### Startbildschirm

Nach dem Start der Anwendung erscheint der Ereigniskonfigurationsbildschirm:

1. **Sprachauswahl** — Selektor oben auf dem Bildschirm. Ein Sprachwechsel übersetzt die gesamte Oberfläche sofort.
2. **Ereignisname** — Pflichtfeld. Geben Sie den Ort und die Art des Ereignisses ein (z. B. „Unfall A4 km 312"). Erscheint in den Berichten.
3. **EAL** — optionales Feld. Einsatzabschnittsleiter — Name der Person, die die Maßnahmen vor Ort koordiniert.
4. **Hinweise anzeigen** — Schalter zur Aktivierung von kontextuellen Hinweisblasen in der gesamten Anwendung. Nützlich bei der ersten Verwendung.
5. **TRIAGE STARTEN** — startet das Ereignis und aktiviert den Timer.

### Ereignis fortsetzen

Falls ein gespeichertes Ereignis im Gerätespeicher vorhanden ist, erscheint beim Start die Option:
- **FORTSETZEN** — setzt das gespeicherte Ereignis mit allen Daten fort
- **NEUES EREIGNIS** — löscht die vorherigen Daten und beginnt von vorne

---

## 3. Sichtung (START-Algorithmus)

### Navigation

Die Anwendung verfügt über 3 Registerkarten im oberen Bereich des Bildschirms:
- **SICHTUNG** — Patientenklassifizierung
- **DISPONIERUNG** — Transport und Management
- **BERICHT** — Zusammenfassung und Export

Die Statistikleiste unter der Navigation zeigt die Anzahl der Patienten in jeder Kategorie: 🔴 T1, 🟡 T2, 🟢 T3, ⚫ T4 sowie die Gesamtsumme.

### START-Sichtungsassistent

Der Assistent führt durch den 6-stufigen START-Algorithmus. Bei jedem Schritt antworten Sie JA oder NEIN:

**Schritt 1 — Kann der Patient gehen?**
- JA → T3 Grün (aufgeschoben)
- NEIN → weiter zu Schritt 2

**Schritt 2 — Atmet der Patient spontan?**
- JA → weiter zu Schritt 4 (Beurteilung der Atemfrequenz)
- NEIN → weiter zu Schritt 3

**Schritt 3 — Atemwege freimachen. Atmet der Patient nach dem Freimachen?**
- JA → T1 Rot (sofort)
- NEIN → T4 Schwarz (Tod)

**Schritt 4 — Ist die Atemfrequenz normal (≤ 30/min)?**
- JA → weiter zu Schritt 5
- NEIN → T1 Rot (sofort)

**Schritt 5 — Ist der Kapillar-Refill normal (≤ 2 Sek.) und der Radialispuls vorhanden?**
- JA → weiter zu Schritt 6
- NEIN → T1 Rot (sofort)

**Schritt 6 — Befolgt der Patient einfache Anweisungen?**
- JA → T2 Gelb (dringend)
- NEIN → T1 Rot (sofort)

Jede Frage enthält einen Hinweis unterhalb der Frage — er beschreibt, wie der jeweilige Parameter zu beurteilen ist.

### Ergebniskarte

Nach Abschluss der Sichtung erscheint die Ergebniskarte mit:

- **Farbmarkierung** — großes Quadrat mit Kategorie (T1/T2/T3/T4)
- **Entscheidungspfad** — zeigt den Ablauf des Algorithmus
- **Geschlecht** — Schaltflächen M (männlich), W (weiblich), ? (unbekannt)
- **Alter** — numerisches Feld mit Schaltflächen -5, -1, +1, +5 (Bereich 0–150)
- **Körperverletzungen** — Schaltfläche zum Öffnen des Verletzungsdiagramms (siehe Abschnitt 4)
- **Notizen** — Feld für zusätzliche Informationen (Verletzungen, Unfallmechanismus usw.)

Verfügbare Aktionen:
- **✓ PATIENT BESTÄTIGEN** — fügt den Patienten zur Liste hinzu
- **↺ NEU STARTEN** — setzt den Assistenten auf Schritt 1 zurück
- **FARBE MANUELL ÄNDERN** — öffnet eine Reihe von 4 Schaltflächen (T1–T4) zur manuellen Kategorieänderung

### Patientenliste vor Ort

Unterhalb des Assistenten wird die Liste „Patienten vor Ort" angezeigt, sortiert nach Priorität (T1 → T2 → T3 → T4).

Jede Patientenkarte enthält:
- Farbmarkierung (z. B. P-001)
- Kategorie, Geschlecht/Alter, Registrierungszeit
- Notizenfeld (bearbeitbar)
- Verletzungsschaltfläche (🩻) mit Zusammenfassung
- 4 Resichtungspunkte (T1–T4) — klicken Sie, um die Kategorie zu ändern
- Löschschaltfläche (✕)
- Resichtungsmarkierung (↺n), falls die Kategorie geändert wurde

### Transportliste

Der Bereich „Im Transport" erscheint, wenn Patienten auf dem Weg zum Krankenhaus sind. Jede Karte zeigt:
- Patientenmarkierung mit Kategoriefarbe
- Route: Team → Krankenhaus → Abfahrtszeit
- Schaltfläche „ändern" — Änderung des Zielkrankenhauses
- Schaltfläche „✓ Zugestellt" — markiert die Zustellung und gibt das Rettungsteam frei

---

## 4. Körperverletzungsdiagramm

Die Schaltfläche „🩻 Körperverletzungen" öffnet ein interaktives Körperdiagramm.

### Diagrammansicht

Es werden zwei Körperumrisse nebeneinander angezeigt:
- **VORNE** (Front) — Kopf, Brustkorb, Bauch, Arme, Beine
- **HINTEN** (Rücken) — Kopf, oberer Rücken, unterer Rücken, Arme, Beine

An den Seiten sind die Seitenbezeichnungen sichtbar: **L** (links) und **R** (rechts), beim Rückansicht gespiegelt.

Klicken Sie auf eine Körperzone, um sie auszuwählen.

### Zonenverletzungspanel

Nach Auswahl einer Zone öffnet sich ein Panel mit 7 Verletzungstypen zur Markierung:
- Fraktur
- Blutung
- Verbrennung
- Wunde
- Schwellung
- Schmerz
- Amputation

Markierte Typen sind mit ✓ gekennzeichnet. Jeder kann per Klick ein-/ausgeschaltet werden.

Am unteren Rand des Panels:
- **✕ NEIN** — verwirft die Änderungen in dieser Zone
- **✓ JA** — bestätigt die Verletzungen für die ausgewählte Zone

### Diagramm bestätigen

Nach Markierung der Verletzungen in allen erforderlichen Zonen:
- **✓ JA** — bestätigt das gesamte Verletzungsdiagramm
- **✕ NEIN** — verwirft alle Änderungen und kehrt zum Zustand vor dem Öffnen zurück

Markierte Verletzungen sind sichtbar:
- Auf der Patientenkarte als Kurzzusammenfassung (z. B. „🩻 Kopf: Fraktur, Blutung")
- Im Bericht beim jeweiligen Patienten

---

## 5. Disponierung

### Modale Konfigurationsfenster

Beim ersten Aufrufen der Registerkarte Disponierung erscheinen nacheinander 3 Dialoge:

**1. Leitender Notarzt (LNA)**
- Feld für Name des LNA
- Schaltfläche zum Datenimport vom Disponenten (siehe Abschnitt 6)
- ÜBERSPRINGEN oder SPEICHERN

**2. Rettungsteams**
- Frage: „Kennen Sie die zum Einsatzort disponierten Rettungsteams?"
- JA → Formular zum Hinzufügen von Teams (Name + Hinzufügen-Schaltfläche)
- NEIN → Standard-Teams verwenden (S-01, S-02, P-01, P-02, P-03, LPR)
- BESTÄTIGEN oder ÜBERSPRINGEN (STANDARD)

**3. Zielkrankenhäuser**
- Formular zum Hinzufügen von Krankenhäusern mit Kapazität:
  - Krankenhausname
  - 🔴 Rote Kapazität (T1) — Anzahl der Betten
  - 🟡 Gelbe Kapazität (T2) — Anzahl der Betten
- BESTÄTIGEN oder ÜBERSPRINGEN (STANDARD)

Jeder Dialog erscheint nur einmal pro Ereignis.

### Disponierungsformular

Der Bereich „Transport disponieren" enthält:

1. **Patient** — Dropdown-Liste mit Patienten vor Ort, sortiert nach Priorität
2. **Rettungsteam** — Liste verfügbarer Teams (besetzte Teams sind ausgeblendet)
3. **Krankenhaus** — Krankenhausliste mit Belegungsinformation:
   - Format: „Krankenhausname [🔴 belegt/gesamt 🟡 belegt/gesamt]"
   - Krankenhäuser ohne freie Plätze in der Patientenkategorie sind mit Warnung gekennzeichnet
4. **TRANSPORT DISPONIEREN** — Schaltfläche aktiv, wenn alle Felder ausgefüllt sind

### Verwaltungspanels

Unter dem Formular befinden sich zwei Panels (auf breiten Bildschirmen nebeneinander, auf schmalen untereinander):

**Krankenhäuser** (Betten GESAMT):
- Krankenhausliste mit bearbeitbaren Kapazitätsfeldern (🔴 und 🟡)
- Neue Krankenhäuser hinzufügen
- Krankenhäuser löschen (✕)
- Liste ein-/ausklappen

**Rettungsteams:**
- Teamliste mit Löschschaltflächen
- Neue Teams hinzufügen
- Liste ein-/ausklappen

### Transporthistorie

Am unteren Rand des Disponierungsbereichs (standardmäßig eingeklappt). Ausklappen zeigt eine chronologische Liste aller Transporte mit:
- Patientenmarkierung mit Farbe
- Team → Krankenhaus → Abfahrtszeit
- Status (im Transport / zugestellt)

---

## 6. Datenimport vom Disponenten

Diese Funktion ermöglicht dem Disponenten, Daten (Ereignisname, LNA, Rettungsteams, Krankenhäuser mit Kapazität) an den Rettungsdienst vor Ort zu senden.

### Methode 1 — SMS-Link (Ein-Klick)

Der Disponent sendet eine SMS mit einem Link im Format:
```
https://0xjaqbek.github.io/triage/?i=DATEN_BASE64
```

Nach Anklicken des Links:
1. Die Anwendung öffnet sich aus dem Cache (kein Internet erforderlich, wenn installiert)
2. Eine Datenvorschau erscheint: Ereignisname, LNA, Teams, Krankenhäuser
3. **IMPORTIEREN** — importiert die Daten und öffnet die Sichtung
4. **ABLEHNEN** — ignoriert die Daten

### Methode 2 — Daten einfügen

Falls der Link die Anwendung nicht öffnet (gesperrte Geräte):
1. Öffnen Sie die Anwendung, navigieren Sie zu Disponierung
2. Klicken Sie im LNA-Dialog auf **📥 DATEN VOM DISPONENTEN IMPORTIEREN**
3. Fügen Sie den SMS-Inhalt ein (ganzer Link oder nur der Base64-Code)
4. Die Datenvorschau erscheint automatisch
5. **IMPORTIEREN** — bestätigt die Daten

### Import während eines Ereignisses

Falls das Ereignis bereits läuft und Sie gesichtete Patienten haben:
- **Patienten** — bleiben erhalten (werden niemals gelöscht)
- **EAL** — bleibt erhalten
- **LNA** — wird durch die Disponenten-Daten ersetzt
- **Rettungsteams** — werden durch die Disponenten-Daten ersetzt
- **Krankenhäuser** — werden durch die Disponenten-Daten ersetzt
- **Ereignisname** — zusammengeführt: „Ihr Name | Disponenten-Name (Disponent)"

Eine gelbe Warnung in der Vorschau informiert über die Beibehaltung der Patienten und des EAL.

---

## 7. Disponentenmodus

Erreichbar über den Link „Disponentenmodus →" auf dem Startbildschirm oder unter der Adresse `/dyspozytor/`.

Die Seite ermöglicht dem Disponenten, Daten für das Team vor Ort vorzubereiten und zu senden.

### Formular

1. **Ereignisname** — z. B. „MCI Kraków-Nowa Huta"
2. **Leitender Notarzt (LNA)** — Name
3. **Rettungsteams** — einzeln hinzufügen (Name + Schaltfläche „+")
4. **Zielkrankenhäuser** — Name + Kapazität 🔴 und 🟡 (+ Hinzufügen-Schaltfläche)

### Versand

- **📋 Datenvorschau** — zeigt eine Zusammenfassung mit Zeichenanzahl und SMS-Anzahl
- **📱 SMS senden** — öffnet die SMS-Anwendung mit fertigem Link
- **📋 Daten kopieren** — kopiert den Link in die Zwischenablage

Die Daten werden in Base64 kodiert (kompatibel mit SMS GSM-7 — bis zu 1530 Zeichen in 10 SMS-Segmenten). Es können ca. 20 Krankenhäuser und 20 Teams übermittelt werden.

---

## 8. Bericht

### Visueller Bericht

Der Bereich „Bericht für den Disponenten" wird automatisch generiert und enthält:

**Kopfzeile:**
- Ereignisname, Startzeit, Berichtszeit
- EAL und LNA (falls angegeben)
- Endzeit (falls das Ereignis beendet wurde)

**Zusammenfassung:**
- Farbmarkierungen mit Patientenanzahl in jeder Kategorie
- Gesamtzahl der Patienten

**Krankenhäuser** (separater Block für jedes Krankenhaus mit Transporten):
- Krankenhausname, Patientenanzahl, Kapazitätsinformation
- Patientenliste mit: Markierung, Alter/Geschlecht, Team, Zeit, Status (IM TRANSPORT / ZUGESTELLT)
- Resichtungshistorie (z. B. ↺T1→T2→T1)
- Krankenhaus-Änderungshistorie (z. B. ⇄Krankenhaus1→Krankenhaus2)
- Verletzungszusammenfassung

**Patienten vor Ort** (falls noch jemand verblieben ist):
- Liste mit: Markierung, Alter/Geschlecht, Kategorie, Notizen, Verletzungen

### Textbericht

Die Schaltfläche „📋 Textbericht" öffnet ein Dialogfeld mit dem vollständigen Bericht im Textformat. Er enthält die vollständige Patientenliste mit Notizen, Geschlecht, Alter und Sichtungspfad.

### Bericht versenden

- **📋 In Zwischenablage kopieren** — kopiert den vollständigen Berichtstext
- **📤 Senden** — öffnet Optionen:
  - **SMS** — öffnet die native SMS-Anwendung mit dem Berichtsinhalt
  - **E-Mail** — öffnet den E-Mail-Client mit Betreff und Berichtsinhalt

### Ereignis beenden

Die Schaltfläche **EREIGNIS BEENDEN** in der Berichtskopfzeile:
- Speichert die Endzeit
- Ändert sich zu **BEENDIGUNG RÜCKGÄNGIG MACHEN** (umkehrbar)
- Die Endzeit erscheint im Bericht

---

## 9. Resichtung und Krankenhauswechsel

### Resichtung

Auf der Patientenkarte im Bereich „Patienten vor Ort" befinden sich 4 farbige Punkte (T1–T4). Das Anklicken einer anderen als der aktuellen Kategorie:
1. Öffnet ein Bestätigungsfenster: „Kategorie des Patienten P-001 von T1 auf T2 ändern?"
2. **BESTÄTIGEN** — ändert die Kategorie, speichert in der Historie
3. **ABBRECHEN** — keine Änderung

Resichtungshistorie:
- Markierung ↺n auf der Patientenkarte (n = Anzahl der Änderungen)
- Vollständige Historie im Bericht (z. B. T1→T2→T1 mit Uhrzeiten)

### Krankenhauswechsel während des Transports

Auf der Patientenkarte „Im Transport" die Schaltfläche „ändern":
1. Öffnet einen Dialog mit dem aktuellen Krankenhaus
2. Dropdown-Liste mit verfügbaren Krankenhäusern (mit Kapazitätsinformation)
3. **BESTÄTIGEN** — ändert das Krankenhaus, speichert in der Historie
4. **ABBRECHEN** — keine Änderung

Die Krankenhaus-Änderungshistorie ist im Bericht sichtbar (⇄Krankenhaus1→Krankenhaus2).

---

## 10. Installation und Offline-Modus

### Installation als PWA

**Android (Chrome):**
1. Öffnen Sie die Anwendung in Chrome
2. Tippen Sie auf ⋮ (Menü) in der oberen rechten Ecke
3. Wählen Sie „Zum Startbildschirm hinzufügen"
4. Tippen Sie auf „Installieren"

**iOS (Safari):**
1. Öffnen Sie die Anwendung in Safari
2. Tippen Sie auf ⎙ (Teilen)
3. Scrollen Sie nach unten und wählen Sie „Zum Home-Bildschirm"

Nach der Installation funktioniert die Anwendung als eigenständige App mit vollem Offline-Modus.

### Offline-Modus

Die Anwendung speichert sich nach dem ersten Laden im Gerätecache. Alle Funktionen funktionieren ohne Internetverbindung:
- Patientensichtung
- Transportdisponierung
- Berichtserstellung
- Datenimport (Einfügen)

Internet wird nur benötigt für: erstes Herunterladen, Aktualisierung sowie SMS-/E-Mail-Versand (über native Geräte-Anwendungen).

### Download als HTML-Datei

Die Schaltfläche **⬇ Herunterladen** auf dem Startbildschirm lädt die Anwendung als eigenständige HTML-Datei herunter. Diese kann auf einen USB-Stick übertragen und auf jedem Computer mit Browser geöffnet werden.

---

## 11. Datenschutz

- Alle Daten werden **ausschließlich auf dem Gerät** gespeichert (localStorage)
- Es werden **keine Daten** an externe Server gesendet
- Keine Analytik, kein Tracking, keine Telemetrie
- Keine Tracking-Cookies
- SMS- und E-Mail-Funktionen nutzen die nativen Geräte-Anwendungen — die App sendet keine Nachrichten direkt
- Datenlöschung: Verwenden Sie „Ereignis beenden" oder löschen Sie die Browserdaten

Die vollständige Datenschutzerklärung ist unter `/privacy/` verfügbar.

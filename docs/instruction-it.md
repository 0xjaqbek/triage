# Manuale d'uso TRIAGE MCI

## 1. Cos'è TRIAGE MCI

TRIAGE MCI è un'applicazione web (PWA) per il triage e la gestione degli incidenti con numerose vittime (MCI), basata sull'algoritmo START. Consente di:
- Classificare i pazienti in 4 categorie (T1–T4)
- Gestire il trasporto verso gli ospedali
- Monitorare la capacità degli ospedali
- Generare e inviare rapporti
- Importare dati dal dispatching

L'applicazione funziona interamente offline. Tutti i dati vengono memorizzati esclusivamente sul dispositivo dell'utente — nessuna informazione viene inviata a server esterni.

Supporta 7 lingue: polacco, inglese, italiano, francese, tedesco, ceco, portoghese.

---

## 2. Avvio dell'evento

### Schermata iniziale

All'avvio dell'applicazione apparirà la schermata di configurazione dell'evento:

1. **Selezione della lingua** — selettore in alto sullo schermo. La modifica della lingua traduce immediatamente l'intera interfaccia.
2. **Nome dell'evento** — campo obbligatorio. Inserire la posizione e il tipo di evento (es. "Incidente A4 km 312"). Apparirà nei rapporti.
3. **KAM** — campo facoltativo. Responsabile dell'Azione Medica — nome e cognome della persona che coordina le operazioni sul posto.
4. **Mostra suggerimenti** — interruttore che attiva i fumetti con suggerimenti contestuali in tutta l'applicazione. Utile al primo utilizzo.
5. **INIZIA TRIAGE** — avvia l'evento e fa partire il cronometro.

### Ripresa dell'evento

Se nella memoria del dispositivo è presente un evento salvato, all'avvio apparirà l'opzione:
- **CONTINUA** — riprende l'evento salvato con tutti i dati
- **NUOVO EVENTO** — cancella i dati precedenti e ricomincia da zero

---

## 3. Triage (algoritmo START)

### Navigazione

L'applicazione ha 3 schede nella parte superiore dello schermo:
- **TRIAGE** — classificazione dei pazienti
- **DISPATCHING** — trasporto e gestione
- **RAPPORTO** — riepilogo ed esportazione

La barra delle statistiche sotto la navigazione mostra il numero di pazienti in ogni categoria: 🔴 T1, 🟡 T2, 🟢 T3, ⚫ T4 e il totale.

### Procedura guidata di triage START

La procedura guidata accompagna l'utente attraverso l'algoritmo START in 6 passaggi. Ad ogni passaggio si risponde SÌ o NO:

**Passaggio 1 — Il paziente cammina?**
- SÌ → T3 Verde (differito)
- NO → passare al passaggio 2

**Passaggio 2 — Respira spontaneamente?**
- SÌ → passare al passaggio 4 (valutazione della frequenza respiratoria)
- NO → passare al passaggio 3

**Passaggio 3 — Liberare le vie aeree. Respira dopo la disostruzione?**
- SÌ → T1 Rosso (immediato)
- NO → T4 Nero (decesso)

**Passaggio 4 — La frequenza respiratoria è nella norma (≤ 30/min)?**
- SÌ → passare al passaggio 5
- NO → T1 Rosso (immediato)

**Passaggio 5 — Il refill capillare è nella norma (≤ 2 sec.) e il polso radiale è presente?**
- SÌ → passare al passaggio 6
- NO → T1 Rosso (immediato)

**Passaggio 6 — Esegue ordini semplici?**
- SÌ → T2 Giallo (urgente)
- NO → T1 Rosso (immediato)

Ogni domanda ha un suggerimento (hint) visibile sotto la domanda — descrive come valutare il parametro in questione.

### Scheda del risultato

Al termine del triage appare la scheda del risultato con:

- **Indicatore colorato** — grande quadrato con la categoria (T1/T2/T3/T4)
- **Percorso decisionale** — mostra lo svolgimento dell'algoritmo
- **Sesso** — pulsanti M (maschio), F (femmina), ? (sconosciuto)
- **Età** — campo numerico con pulsanti -5, -1, +1, +5 (intervallo 0–150)
- **Lesioni corporee** — pulsante che apre il diagramma delle lesioni (vedi sezione 4)
- **Note** — campo per informazioni aggiuntive (lesioni, meccanismo del trauma ecc.)

Azioni disponibili:
- **✓ CONFERMA PAZIENTE** — aggiunge il paziente alla lista
- **↺ DA CAPO** — reimposta la procedura guidata al passaggio 1
- **CAMBIA COLORE MANUALMENTE** — apre una fila di 4 pulsanti (T1–T4) per la modifica manuale della categoria

### Lista pazienti sul posto

Sotto la procedura guidata è visualizzata la lista "Pazienti sul posto" ordinata per priorità (T1 → T2 → T3 → T4).

Ogni scheda paziente contiene:
- Indicatore colorato (es. P-001)
- Categoria, sesso/età, ora di registrazione
- Campo note (modificabile)
- Pulsante lesioni (🩻) con riepilogo
- 4 punti di re-triage (T1–T4) — cliccare per cambiare la categoria
- Pulsante di eliminazione (✕)
- Indicatore di re-triage (↺n) se la categoria è stata modificata

### Lista in trasporto

La sezione "In trasporto" appare quando i pazienti sono in viaggio verso l'ospedale. Ogni scheda mostra:
- Indicatore del paziente con il colore della categoria
- Percorso: Squadra → Ospedale → ora di partenza
- Pulsante "modifica" — modifica dell'ospedale di destinazione
- Pulsante "✓ Consegnato" — segna la consegna e libera la squadra

---

## 4. Diagramma delle lesioni corporee

Il pulsante "🩻 Lesioni corporee" apre un diagramma interattivo del corpo.

### Vista del diagramma

Vengono visualizzati due contorni del corpo affiancati:
- **ANTERIORE** (front) — testa, torace, addome, braccia, gambe
- **POSTERIORE** (back) — testa, parte superiore della schiena, parte inferiore della schiena, braccia, gambe

Ai lati sono visibili le indicazioni dei lati: **S** (sinistra) e **D** (destra), invertite per la vista posteriore.

Cliccare su una zona del corpo per selezionarla.

### Pannello delle lesioni della zona

Dopo aver selezionato una zona, si apre un pannello con 7 tipi di trauma da contrassegnare:
- Frattura
- Emorragia
- Ustione
- Ferita
- Gonfiore
- Dolore
- Amputazione

I tipi contrassegnati sono indicati con ✓. Ognuno può essere attivato/disattivato con un clic.

In fondo al pannello:
- **✕ NO** — annulla le modifiche in questa zona
- **✓ SÌ** — conferma le lesioni per la zona selezionata

### Conferma del diagramma

Dopo aver contrassegnato le lesioni in tutte le zone necessarie:
- **✓ SÌ** — conferma l'intero diagramma delle lesioni
- **✕ NO** — annulla tutte le modifiche e torna allo stato precedente all'apertura

Le lesioni contrassegnate sono visibili:
- Sulla scheda del paziente come riepilogo abbreviato (es. "🩻 Testa: Frattura, Emorragia")
- Nel rapporto relativo al paziente

---

## 5. Dispatching

### Finestre modali di configurazione

Al primo accesso alla scheda Dispatching appaiono in sequenza 3 finestre modali:

**1. Direttore delle Operazioni Mediche (GDM)**
- Campo per nome e cognome del GDM
- Pulsante di importazione dati dal dispatching (vedi sezione 6)
- SALTA o SALVA

**2. Squadre di soccorso**
- Domanda: "Conosci le squadre di soccorso inviate sul posto?"
- SÌ → modulo per aggiungere le squadre (nome + pulsante aggiungi)
- NO → usa le squadre predefinite (S-01, S-02, P-01, P-02, P-03, LPR)
- CONFERMA o SALTA (PREDEFINITE)

**3. Ospedali di destinazione**
- Modulo per aggiungere ospedali con capacità:
  - Nome dell'ospedale
  - 🔴 Capacità rossa (T1) — numero di posti letto
  - 🟡 Capacità gialla (T2) — numero di posti letto
- CONFERMA o SALTA (PREDEFINITE)

Ogni finestra modale appare solo una volta per evento.

### Modulo di dispatching

La sezione "Disponi il trasporto" contiene:

1. **Paziente** — elenco a discesa con i pazienti sul posto, ordinato per priorità
2. **Squadra di soccorso** — elenco delle squadre disponibili (le squadre occupate sono nascoste)
3. **Ospedale** — elenco degli ospedali con informazioni sull'occupazione:
   - Formato: "Nome ospedale [🔴 occupati/totali 🟡 occupati/totali]"
   - Ospedali senza posti liberi nella categoria del paziente contrassegnati con avviso
4. **DISPONI TRASPORTO** — pulsante attivo quando tutti i campi sono compilati

### Pannelli di gestione

Sotto il modulo si trovano due pannelli (su schermi larghi affiancati, su schermi stretti sovrapposti):

**Ospedali** (posti letto TOTALI):
- Elenco degli ospedali con campi di capacità modificabili (🔴 e 🟡)
- Aggiunta di nuovi ospedali
- Eliminazione degli ospedali (✕)
- Comprimere/espandere l'elenco

**Squadre di soccorso:**
- Elenco delle squadre con pulsanti di eliminazione
- Aggiunta di nuove squadre
- Comprimere/espandere l'elenco

### Cronologia dei trasporti

In fondo alla sezione dispatching (compressa per impostazione predefinita). Espandere per visualizzare l'elenco cronologico di tutti i trasporti con:
- Indicatore del paziente con colore
- Squadra → Ospedale → ora di partenza
- Stato (in trasporto / consegnato)

---

## 6. Importazione dati dal dispatching

La funzione consente al dispatcher di inviare dati (nome dell'evento, GDM, squadre di soccorso, ospedali con capacità) al soccorritore sul campo.

### Metodo 1 — Link SMS (un solo clic)

Il dispatcher invia un SMS con un link nel formato:
```
https://0xjaqbek.github.io/triage/?i=DATI_BASE64
```

Dopo aver cliccato il link:
1. L'applicazione si apre dalla cache (non richiede internet se installata)
2. Appare un'anteprima dei dati: nome dell'evento, GDM, squadre, ospedali
3. **IMPORTA** — importa i dati e apre il triage
4. **RIFIUTA** — ignora i dati

### Metodo 2 — Incollare i dati

Se il link non apre l'applicazione (dispositivi bloccati):
1. Aprire l'applicazione, andare alla sezione Dispatching
2. Nella finestra modale GDM cliccare **📥 IMPORTA DATI DAL DISPATCHER**
3. Incollare il contenuto dell'SMS (l'intero link o solo il codice Base64)
4. L'anteprima dei dati apparirà automaticamente
5. **IMPORTA** — conferma i dati

### Importazione durante l'evento

Se l'evento è già in corso e ci sono pazienti già triaggiati:
- **Pazienti** — mantenuti (non vengono mai eliminati)
- **KAM** — mantenuto
- **GDM** — sostituito con i dati del dispatcher
- **Squadre di soccorso** — sostituite con i dati del dispatcher
- **Ospedali** — sostituiti con i dati del dispatcher
- **Nome dell'evento** — unito: "Il tuo nome | Nome del dispatcher (dispatcher)"

Un avviso giallo nell'anteprima informa sul mantenimento dei pazienti e del KAM.

---

## 7. Modalità dispatcher

Accessibile tramite il link "Modalità dispatcher →" nella schermata iniziale o all'indirizzo `/dyspozytor/`.

La pagina consente al dispatcher di preparare e inviare i dati per la squadra sul campo.

### Modulo

1. **Nome dell'evento** — es. "MCI Kraków-Nowa Huta"
2. **Direttore delle Operazioni Mediche (GDM)** — nome e cognome
3. **Squadre di soccorso** — aggiungere una alla volta (nome + pulsante "+")
4. **Ospedali di destinazione** — nome + capacità 🔴 e 🟡 (+ pulsante aggiungi)

### Invio

- **📋 Anteprima dati** — mostra il riepilogo con la dimensione in caratteri e il numero di SMS
- **📱 Invia SMS** — apre l'applicazione SMS con il link pronto
- **📋 Copia dati** — copia il link negli appunti

I dati sono codificati in Base64 (compatibilità con SMS GSM-7 — fino a 1530 caratteri in 10 segmenti SMS). Può contenere circa 20 ospedali e 20 squadre.

---

## 8. Rapporto

### Rapporto visuale

La sezione "Rapporto per il dispatcher" si genera automaticamente e contiene:

**Intestazione:**
- Nome dell'evento, ora di inizio, ora del rapporto
- KAM e GDM (se indicati)
- Ora di chiusura (se l'evento è concluso)

**Riepilogo:**
- Indicatori colorati con il numero di pazienti in ogni categoria
- Numero totale di pazienti

**Ospedali** (blocco separato per ogni ospedale con i trasporti):
- Nome dell'ospedale, numero di pazienti, informazioni sulla capacità
- Elenco dei pazienti con: indicatore, età/sesso, squadra, ora, stato (IN TRASPORTO / CONSEGNATO)
- Cronologia del re-triage (es. ↺T1→T2→T1)
- Cronologia dei cambi di ospedale (es. ⇄Ospedale1→Ospedale2)
- Riepilogo delle lesioni

**Pazienti sul posto** (se qualcuno è rimasto):
- Elenco con: indicatore, età/sesso, categoria, note, lesioni

### Rapporto testuale

Il pulsante "📋 Rapporto testuale" apre una finestra modale con il rapporto completo in formato testo. Contiene l'elenco completo dei pazienti con note, sesso, età e percorso di triage.

### Invio del rapporto

- **📋 Copia negli appunti** — copia il testo completo del rapporto
- **📤 Invia** — apre le opzioni:
  - **SMS** — apre l'applicazione SMS nativa con il contenuto del rapporto
  - **Email** — apre il client email con oggetto e contenuto del rapporto

### Chiusura dell'evento

Il pulsante **CHIUDI EVENTO** nell'intestazione del rapporto:
- Salva l'ora di chiusura
- Si trasforma in **ANNULLA CHIUSURA** (reversibile)
- L'ora di chiusura appare nel rapporto

---

## 9. Re-triage e cambio di ospedale

### Re-triage

Sulla scheda del paziente nella sezione "Pazienti sul posto" ci sono 4 punti colorati (T1–T4). Cliccando una categoria diversa da quella attuale:
1. Si apre una finestra di conferma: "Cambiare la categoria del paziente P-001 da T1 a T2?"
2. **CONFERMA** — cambia la categoria e la salva nella cronologia
3. **ANNULLA** — nessuna modifica

Cronologia del re-triage:
- Indicatore ↺n sulla scheda del paziente (n = numero di modifiche)
- Cronologia completa nel rapporto (es. T1→T2→T1 con orari)

### Cambio di ospedale durante il trasporto

Sulla scheda del paziente "In trasporto" il pulsante "modifica":
1. Apre una finestra modale con l'ospedale attuale
2. Elenco a discesa con gli ospedali disponibili (con informazioni sulla capacità)
3. **CONFERMA** — cambia l'ospedale e lo salva nella cronologia
4. **ANNULLA** — nessuna modifica

La cronologia dei cambi di ospedale è visibile nel rapporto (⇄Ospedale1→Ospedale2).

---

## 10. Installazione e modalità offline

### Installazione come PWA

**Android (Chrome):**
1. Aprire l'applicazione in Chrome
2. Toccare ⋮ (menu) nell'angolo in alto a destra
3. Selezionare "Aggiungi alla schermata Home"
4. Toccare "Installa"

**iOS (Safari):**
1. Aprire l'applicazione in Safari
2. Toccare ⎙ (Condividi)
3. Scorrere verso il basso e selezionare "Aggiungi alla schermata Home"

Dopo l'installazione l'applicazione funziona come un'applicazione indipendente con piena modalità offline.

### Modalità offline

L'applicazione, dopo il primo caricamento, viene salvata nella cache del dispositivo. Tutte le funzioni operano senza connessione a internet:
- Triage dei pazienti
- Dispatching dei trasporti
- Generazione dei rapporti
- Importazione dati (incollare)

La connessione a internet è necessaria solo per: il primo download, gli aggiornamenti e l'invio di SMS/email (tramite le applicazioni native del dispositivo).

### Download come file HTML

Il pulsante **⬇ Scarica** nella schermata iniziale scarica l'applicazione come file HTML autonomo. Può essere trasferito su una chiavetta USB e avviato su qualsiasi computer dotato di browser.

---

## 11. Privacy

- Tutti i dati sono memorizzati **esclusivamente sul dispositivo** (localStorage)
- Nessun dato viene **inviato** a server esterni
- Nessuna analitica, tracciamento o telemetria
- Nessun cookie di tracciamento
- Le funzioni SMS ed email utilizzano le applicazioni native del dispositivo — l'applicazione non invia messaggi direttamente
- Eliminazione dei dati: utilizzare "Chiudi evento" o cancellare i dati del browser

L'informativa completa sulla privacy è disponibile all'indirizzo `/privacy/`.

# Manuel d'utilisation TRIAGE MCI

## 1. Qu'est-ce que TRIAGE MCI

TRIAGE MCI est une application web (PWA) de tri et de gestion des incidents à victimes multiples (MCI) basée sur l'algorithme START. Elle permet :
- La classification des patients en 4 catégories (T1–T4)
- La gestion du transport vers les hôpitaux
- Le suivi de la capacité hospitalière
- La génération et l'envoi de rapports
- L'importation de données depuis le régulateur

L'application fonctionne entièrement hors ligne. Toutes les données sont stockées uniquement sur l'appareil de l'utilisateur — aucune information n'est envoyée à des serveurs externes.

Elle prend en charge 7 langues : polonais, anglais, italien, français, allemand, tchèque, portugais.

---

## 2. Démarrage d'un événement

### Écran d'accueil

Au lancement de l'application, l'écran de configuration de l'événement s'affiche :

1. **Choix de la langue** — sélecteur en haut de l'écran. Le changement de langue traduit instantanément toute l'interface.
2. **Nom de l'événement** — champ obligatoire. Saisissez la localisation et le type d'événement (p. ex. « Accident A4 km 312 »). Il apparaîtra dans les rapports.
3. **KAM** — champ facultatif. Directeur des Opérations Médicales — nom et prénom de la personne coordonnant les actions sur le terrain.
4. **Afficher les info-bulles** — interrupteur activant les bulles d'aide contextuelle dans toute l'application. Utile lors de la première utilisation.
5. **COMMENCER LE TRIAGE** — lance l'événement et démarre le chronomètre.

### Reprise d'un événement

Si un événement sauvegardé existe dans la mémoire de l'appareil, une option apparaît au démarrage :
- **CONTINUER** — reprend l'événement sauvegardé avec toutes les données
- **NOUVEL ÉVÉNEMENT** — efface les données précédentes et recommence à zéro

---

## 3. Tri (algorithme START)

### Navigation

L'application dispose de 3 onglets en haut de l'écran :
- **TRI** — classification des patients
- **RÉGULATION** — transport et gestion
- **RAPPORT** — résumé et export

La barre de statistiques sous la navigation affiche le nombre de patients dans chaque catégorie : 🔴 T1, 🟡 T2, 🟢 T3, ⚫ T4 ainsi que le total.

### Assistant de tri START

L'assistant guide à travers l'algorithme START en 6 étapes. À chaque étape, vous répondez OUI ou NON :

**Étape 1 — Le patient marche-t-il ?**
- OUI → T3 Vert (différé)
- NON → passer à l'étape 2

**Étape 2 — Respire-t-il spontanément ?**
- OUI → passer à l'étape 4 (évaluation de la fréquence respiratoire)
- NON → passer à l'étape 3

**Étape 3 — Libérez les voies aériennes. Respire-t-il après la libération ?**
- OUI → T1 Rouge (immédiat)
- NON → T4 Noir (décédé)

**Étape 4 — La fréquence respiratoire est-elle normale (≤ 30/min) ?**
- OUI → passer à l'étape 5
- NON → T1 Rouge (immédiat)

**Étape 5 — Le temps de recoloration capillaire est-il normal (≤ 2 s) et le pouls radial est-il présent ?**
- OUI → passer à l'étape 6
- NON → T1 Rouge (immédiat)

**Étape 6 — Obéit-il à des ordres simples ?**
- OUI → T2 Jaune (urgent)
- NON → T1 Rouge (immédiat)

Chaque question est accompagnée d'une info-bulle visible sous la question — elle décrit comment évaluer le paramètre concerné.

### Fiche de résultat

À la fin du tri, une fiche de résultat s'affiche avec :

- **Marqueur de couleur** — grand carré avec la catégorie (T1/T2/T3/T4)
- **Chemin décisionnel** — montre le déroulement de l'algorithme
- **Sexe** — boutons H (homme), F (femme), ? (inconnu)
- **Âge** — champ numérique avec boutons -5, -1, +1, +5 (plage 0–150)
- **Lésions corporelles** — bouton ouvrant le diagramme des lésions (voir section 4)
- **Notes** — champ pour des informations supplémentaires (lésions, mécanisme du traumatisme, etc.)

Actions disponibles :
- **✓ VALIDER LE PATIENT** — ajoute le patient à la liste
- **↺ RECOMMENCER** — réinitialise l'assistant à l'étape 1
- **CHANGER LA COULEUR MANUELLEMENT** — ouvre une rangée de 4 boutons (T1–T4) pour modifier la catégorie manuellement

### Liste des patients sur place

Sous l'assistant s'affiche la liste « Patients sur place » triée par priorité (T1 → T2 → T3 → T4).

Chaque fiche patient contient :
- Marqueur de couleur (p. ex. P-001)
- Catégorie, sexe/âge, heure d'enregistrement
- Champ de notes (modifiable)
- Bouton de lésions (🩻) avec résumé
- 4 points de re-triage (T1–T4) — cliquez pour changer de catégorie
- Bouton de suppression (✕)
- Marqueur de re-triage (↺n) si la catégorie a été modifiée

### Liste en transport

La section « En transport » apparaît lorsque des patients sont en route vers l'hôpital. Chaque fiche affiche :
- Marqueur du patient avec la couleur de catégorie
- Itinéraire : Équipe → Hôpital → heure de départ
- Bouton « modifier » — changement de l'hôpital de destination
- Bouton « ✓ Livré » — marque la livraison et libère l'équipe

---

## 4. Diagramme des lésions corporelles

Le bouton « 🩻 Lésions corporelles » ouvre un diagramme corporel interactif.

### Vue du diagramme

Deux silhouettes sont affichées côte à côte :
- **AVANT** (face) — tête, thorax, abdomen, bras, jambes
- **ARRIÈRE** (dos) — tête, haut du dos, bas du dos, bras, jambes

Sur les côtés, les indications latérales sont visibles : **G** (gauche) et **D** (droite), inversées pour la vue arrière.

Cliquez sur une zone du corps pour la sélectionner.

### Panneau de lésions de la zone

Après la sélection d'une zone, un panneau s'ouvre avec 7 types de lésions à cocher :
- Fracture
- Hémorragie
- Brûlure
- Plaie
- Œdème
- Douleur
- Amputation

Les types cochés sont marqués par ✓. Chacun peut être activé/désactivé par un clic.

En bas du panneau :
- **✕ NON** — rejette les modifications pour cette zone
- **✓ OUI** — valide les lésions pour la zone sélectionnée

### Validation du diagramme

Après avoir coché les lésions dans toutes les zones nécessaires :
- **✓ OUI** — valide l'ensemble du diagramme des lésions
- **✕ NON** — rejette toutes les modifications et revient à l'état précédant l'ouverture

Les lésions cochées sont visibles :
- Sur la fiche patient sous forme de résumé abrégé (p. ex. « 🩻 Tête : Fracture, Hémorragie »)
- Dans le rapport pour le patient concerné

---

## 5. Régulation

### Fenêtres modales de configuration

Lors du premier accès à l'onglet Régulation, 3 fenêtres modales apparaissent successivement :

**1. Directeur des Secours Médicaux (GDM)**
- Champ pour le nom et prénom du GDM
- Bouton d'importation des données du régulateur (voir section 6)
- PASSER ou ENREGISTRER

**2. Équipes de secours**
- Question : « Connaissez-vous les équipes de secours envoyées sur place ? »
- OUI → formulaire d'ajout d'équipes (nom + bouton ajouter)
- NON → utiliser les équipes par défaut (S-01, S-02, P-01, P-02, P-03, LPR)
- VALIDER ou PASSER (PAR DÉFAUT)

**3. Hôpitaux de destination**
- Formulaire d'ajout d'hôpitaux avec capacité :
  - Nom de l'hôpital
  - 🔴 Capacité rouge (T1) — nombre de lits
  - 🟡 Capacité jaune (T2) — nombre de lits
- VALIDER ou PASSER (PAR DÉFAUT)

Chaque fenêtre modale n'apparaît qu'une seule fois par événement.

### Formulaire de régulation

La section « Affecter un transport » contient :

1. **Patient** — liste déroulante des patients sur place, triée par priorité
2. **Équipe de secours** — liste des équipes disponibles (les équipes occupées sont masquées)
3. **Hôpital** — liste des hôpitaux avec informations d'occupation :
   - Format : « Nom de l'hôpital [🔴 occupés/total 🟡 occupés/total] »
   - Les hôpitaux sans places disponibles dans la catégorie du patient sont signalés par un avertissement
4. **AFFECTER LE TRANSPORT** — bouton actif lorsque tous les champs sont remplis

### Panneaux de gestion

Sous le formulaire se trouvent deux panneaux (côte à côte sur écrans larges, empilés sur écrans étroits) :

**Hôpitaux** (lits TOUS) :
- Liste des hôpitaux avec champs de capacité modifiables (🔴 et 🟡)
- Ajout de nouveaux hôpitaux
- Suppression d'hôpitaux (✕)
- Réduction/expansion de la liste

**Équipes de secours :**
- Liste des équipes avec boutons de suppression
- Ajout de nouvelles équipes
- Réduction/expansion de la liste

### Historique des transports

En bas de la section de régulation (réduit par défaut). Développez pour voir la liste chronologique de tous les transports avec :
- Marqueur du patient avec couleur
- Équipe → Hôpital → heure de départ
- Statut (en transport / livré)

---

## 6. Importation des données du régulateur

Cette fonction permet au régulateur d'envoyer des données (nom de l'événement, GDM, équipes de secours, hôpitaux avec capacité) à l'équipe sur le terrain.

### Méthode 1 — Lien SMS (un seul clic)

Le régulateur envoie un SMS avec un lien au format :
```
https://0xjaqbek.github.io/triage/?i=DONNÉES_BASE64
```

Après avoir cliqué sur le lien :
1. L'application s'ouvre depuis le cache (ne nécessite pas d'internet si elle est installée)
2. Un aperçu des données s'affiche : nom de l'événement, GDM, équipes, hôpitaux
3. **IMPORTER** — importe les données et ouvre le tri
4. **REJETER** — ignore les données

### Méthode 2 — Collage des données

Si le lien n'ouvre pas l'application (appareils verrouillés) :
1. Ouvrez l'application, accédez à Régulation
2. Dans la fenêtre modale GDM, cliquez sur **📥 IMPORTER LES DONNÉES DU RÉGULATEUR**
3. Collez le contenu du SMS (lien complet ou code Base64 seul)
4. L'aperçu des données s'affiche automatiquement
5. **IMPORTER** — valide les données

### Importation en cours d'événement

Si l'événement est déjà en cours et que vous avez des patients triés :
- **Patients** — conservés (jamais supprimés)
- **KAM** — conservé
- **GDM** — remplacé par les données du régulateur
- **Équipes de secours** — remplacées par les données du régulateur
- **Hôpitaux** — remplacés par les données du régulateur
- **Nom de l'événement** — fusionné : « Votre nom | Nom du régulateur (régulateur) »

Un avertissement jaune dans l'aperçu informe de la conservation des patients et du KAM.

---

## 7. Mode régulateur

Accessible via le lien « Mode régulateur → » sur l'écran d'accueil ou à l'adresse `/dyspozytor/`.

Cette page permet au régulateur de préparer et d'envoyer des données à l'équipe sur le terrain.

### Formulaire

1. **Nom de l'événement** — p. ex. « MCI Kraków-Nowa Huta »
2. **Directeur des Secours Médicaux (GDM)** — nom et prénom
3. **Équipes de secours** — ajoutez-les une par une (nom + bouton « + »)
4. **Hôpitaux de destination** — nom + capacité 🔴 et 🟡 (+ bouton ajouter)

### Envoi

- **📋 Aperçu des données** — affiche le résumé avec la taille en caractères et le nombre de SMS
- **📱 Envoyer par SMS** — ouvre l'application SMS avec le lien prêt à l'emploi
- **📋 Copier les données** — copie le lien dans le presse-papiers

Les données sont encodées en Base64 (compatibilité SMS GSM-7 — jusqu'à 1530 caractères en 10 segments SMS). Cela permet d'inclure environ 20 hôpitaux et 20 équipes.

---

## 8. Rapport

### Rapport visuel

La section « Rapport pour le régulateur » se génère automatiquement et contient :

**En-tête :**
- Nom de l'événement, heure de début, heure du rapport
- KAM et GDM (si renseignés)
- Heure de fin (si l'événement est terminé)

**Résumé :**
- Marqueurs de couleur avec le nombre de patients dans chaque catégorie
- Nombre total de patients

**Hôpitaux** (un bloc distinct par hôpital avec les transports) :
- Nom de l'hôpital, nombre de patients, informations de capacité
- Liste des patients avec : marqueur, âge/sexe, équipe, heure, statut (EN TRANSPORT / LIVRÉ)
- Historique de re-triage (p. ex. ↺T1→T2→T1)
- Historique des changements d'hôpital (p. ex. ⇄Hôpital1→Hôpital2)
- Résumé des lésions

**Patients sur place** (s'il en reste) :
- Liste avec : marqueur, âge/sexe, catégorie, notes, lésions

### Rapport textuel

Le bouton « 📋 Rapport textuel » ouvre une fenêtre modale avec le rapport complet au format texte. Il contient la liste complète des patients avec notes, sexe, âge et chemin de tri.

### Envoi du rapport

- **📋 Copier dans le presse-papiers** — copie le texte complet du rapport
- **📤 Envoyer** — ouvre les options :
  - **SMS** — ouvre l'application SMS native avec le contenu du rapport
  - **Email** — ouvre le client de messagerie avec l'objet et le contenu du rapport

### Clôture de l'événement

Le bouton **CLÔTURER L'ÉVÉNEMENT** dans l'en-tête du rapport :
- Enregistre l'heure de fin
- Se transforme en **ANNULER LA CLÔTURE** (réversible)
- L'heure de fin apparaît dans le rapport

---

## 9. Re-triage et changement d'hôpital

### Re-triage

Sur la fiche patient dans la section « Patients sur place » se trouvent 4 points de couleur (T1–T4). Un clic sur une catégorie différente de la catégorie actuelle :
1. Ouvre une fenêtre de confirmation : « Changer la catégorie du patient P-001 de T1 à T2 ? »
2. **CONFIRMER** — change la catégorie, enregistre dans l'historique
3. **ANNULER** — aucune modification

Historique de re-triage :
- Marqueur ↺n sur la fiche patient (n = nombre de modifications)
- Historique complet dans le rapport (p. ex. T1→T2→T1 avec les heures)

### Changement d'hôpital en cours de transport

Sur la fiche patient « En transport », bouton « modifier » :
1. Ouvre une fenêtre modale avec l'hôpital actuel
2. Liste déroulante des hôpitaux disponibles (avec informations de capacité)
3. **CONFIRMER** — change l'hôpital, enregistre dans l'historique
4. **ANNULER** — aucune modification

L'historique des changements d'hôpital est visible dans le rapport (⇄Hôpital1→Hôpital2).

---

## 10. Installation et mode hors ligne

### Installation en tant que PWA

**Android (Chrome) :**
1. Ouvrez l'application dans Chrome
2. Appuyez sur ⋮ (menu) en haut à droite
3. Sélectionnez « Ajouter à l'écran d'accueil »
4. Appuyez sur « Installer »

**iOS (Safari) :**
1. Ouvrez l'application dans Safari
2. Appuyez sur ⎙ (Partager)
3. Faites défiler vers le bas et sélectionnez « Sur l'écran d'accueil »

Après l'installation, l'application fonctionne comme une application autonome avec un mode hors ligne complet.

### Mode hors ligne

L'application, après son premier chargement, se sauvegarde dans le cache de l'appareil. Toutes les fonctions sont disponibles sans connexion internet :
- Tri des patients
- Régulation du transport
- Génération des rapports
- Importation des données (collage)

L'internet n'est nécessaire que pour : le premier téléchargement, les mises à jour, et l'envoi de SMS/email (via les applications natives de l'appareil).

### Téléchargement en fichier HTML

Le bouton **⬇ Télécharger** sur l'écran d'accueil télécharge l'application sous forme de fichier HTML autonome. Il peut être transféré sur une clé USB et exécuté sur n'importe quel ordinateur disposant d'un navigateur.

---

## 11. Confidentialité

- Toutes les données sont stockées **uniquement sur l'appareil** (localStorage)
- Aucune donnée **n'est envoyée** à des serveurs externes
- Pas d'analytique, de suivi, ni de télémétrie
- Pas de cookies de pistage
- Les fonctions SMS et email utilisent les applications natives de l'appareil — l'application n'envoie pas de messages directement
- Suppression des données : utilisez « Clôturer l'événement » ou effacez les données du navigateur

La politique de confidentialité complète est disponible à l'adresse `/privacy/`.

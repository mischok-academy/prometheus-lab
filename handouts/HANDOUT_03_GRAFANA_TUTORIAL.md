# Grafana Dashboard erstellen - Step by Step

**Schritt-für-Schritt Anleitung für Anfänger**

---

## 🎯 Ziel
Am Ende hast du ein eigenes Dashboard mit 5 Panels, das die wichtigsten Metriken zeigt.

---

## 📍 Schritt 1: Grafana öffnen

1. Browser öffnen
2. Gehe zu: **http://localhost:3000**
3. Melde dich an:
   - **Username:** admin
   - **Passwort:** admin123

Du siehst jetzt die Grafana-Startseite.

---

## 📍 Schritt 2: Neues Dashboard erstellen

### Option A: Über das Plus-Menü
1. Klick auf das **"+"** Symbol (oben links)
2. Wähle **"Create Dashboard"**
3. Klick auf **"Create new dashboard"**

### Option B: Über das Hamburger-Menü
1. Klick auf das **Hamburger-Menü** ☰ (oben links)
2. Wähle **"Dashboards"**
3. Klick auf **"New Dashboard"** (oben rechts, blauer Button)

**Ergebnis:** Du hast ein leeres Dashboard vor dir.

---

## 📍 Schritt 3: Erstes Panel hinzufügen (Uptime)

Ein Panel = Ein Graph/Visualisierung auf dem Dashboard

### Erstes Panel erstellen
1. Klick auf **"Add Panel"** (oben rechts) oder **"Add a new panel"** (große Fläche)
2. Du siehst ein leeres Panel mit Editor unten

### Query schreiben
1. Im Editor (unten) findest du "Metrics"
2. Klick in das Eingabefeld
3. Schreibe: `up`
4. Drücke **Enter** oder warte auf Autocomplete

**Was siehst du:** Für jeden Service ein Punkt (1 = up, 0 = down)

### Visualisierung ändern
1. Oben rechts findest du "Visualization"
2. Aktuelle Auswahl: "Time series"
3. Klick darauf und wähle: **"Stat"**
4. Das Panel zeigt jetzt große Zahlen statt Kurven

### Panel speichern
1. Oben rechts klick **"Apply"**
2. Titel geben: "Services Status" (Klick auf "Panel Title")
3. Oben rechts **"Save"**
4. Name: "App Monitoring"
5. Klick **"Save"**

---

## 📍 Schritt 4: Zweites Panel (CPU-Auslastung)

Das Dashboard ist jetzt gespeichert. Wir fügen mehr Panels hinzu.

### Panel hinzufügen
1. Klick **"Add Panel"** (oben rechts)
2. Klick **"Create new panel"**

### CPU Query
1. Im Editor unten, Feld "Metrics"
2. Schreibe diese Query:
```promql
100 - (avg(irate(node_cpu_seconds_total{mode="idle"}[5m])) * 100)
```
3. Drücke **Enter**

**Was siehst du:** Eine Kurve mit CPU-Auslastung über Zeit

### Visualisierung: Graph beibehalten
- Das sieht gut aus! Bleibe bei "Time series"

### Panel formatieren
1. Rechts findest du "Standard options"
2. Scroll zu **"Unit"**
3. Wähle **"Percent (0-100)"**
4. Jetzt zeigt es `45%` statt `45` an

### Maximal-Wert setzen
1. Scroll zu **"Max"** (unter Unit)
2. Setze Wert auf: `100`
3. Jetzt geht die Y-Achse nur bis 100

### Alarm-Schwelle einziehen
1. Scroll zu **"Thresholds"**
2. Klick **"Add threshold"**
3. Stelle ein:
   - Wert: `80`
   - Farbe: Orange/Rot
4. Das zeigt eine rote Linie bei 80% an

### Speichern
1. Klick **"Apply"**
2. Titel: "CPU Auslastung"
3. Speichern (oben rechts)

---

## 📍 Schritt 5: Drittes Panel (RAM-Speicher)

### Panel hinzufügen
1. Klick **"Add Panel"** → **"Create new panel"**

### RAM Query
```promql
(1 - (node_memory_MemAvailable_bytes / node_memory_MemTotal_bytes)) * 100
```

### Visualisierung: Gauge
1. Wähle: **"Gauge"** (statt Time Series)
2. Das zeigt einen großen Kreis-Messwert

### Formatierung
1. Unit: **"Percent (0-100)"**
2. Min: `0`, Max: `100`
3. Thresholds:
   - Grün bis 60%
   - Gelb 60-80%
   - Rot ab 80%

### Speichern
- Titel: "RAM Nutzung"

---

## 📍 Schritt 6: Viertes Panel (Requests pro Sekunde)

### Panel hinzufügen
1. **"Add Panel"** → **"Create new panel"**

### Request-Rate Query
```promql
rate(app_requests_total[5m])
```

### Einstellungen
1. Visualization: **"Time series"** (Graph)
2. Unit: **"Requests/sec"**
3. Titel: "Request Rate"

---

## 📍 Schritt 7: Fünftes Panel (Error Rate)

### Panel hinzufügen
1. **"Add Panel"** → **"Create new panel"**

### Error-Rate Query
```promql
(sum(rate(app_errors_total[5m])) / sum(rate(app_requests_total[5m]))) * 100
```

### Einstellungen
1. Visualization: **"Stat"**
2. Unit: **"Percent"**
3. Thresholds:
   - Grün < 1%
   - Rot > 1%
4. Titel: "Error Rate"

---

## 🎨 Schritt 8: Layout schöner machen

### Panels bewegen
1. Klick auf eine Panel-Header
2. Ziehe sie an neue Position
3. Panels sortieren sich automatisch neu

### Größe ändern
1. Hoover über Ecke eines Panels
2. Ziehe um die Größe zu ändern

### Gutes Layout
```
┌─────────────────────────────────┐
│  Status    │  CPU  │  RAM       │
│  (oben)    │ (oben) │ (oben)    │
├─────────────────────────────────┤
│  Requests (unten groß)          │
│                                  │
├─────────────────────────────────┤
│  Errors (unten klein)           │
```

---

## 💾 Schritt 9: Alles speichern

1. Oben rechts klick auf **Zahnrad** ⚙️
2. Klick **"Save dashboard"** (oder nur Strg+S)
3. Gib einen Namen ein (falls nicht schon geschehen)
4. Klick **"Save"**

**Fertig!** Dein Dashboard wird automatisch gespeichert.

---

## 🔍 Nützliche Tipps & Tricks

### Dashboard duplizieren
1. Zahnrad ⚙️ → Dashboard settings
2. Scroll zu "Copy dashboard"
3. Neuer Name → "Copy"

### Variablen hinzufügen (fortgeschritten)
Mit Variablen kannst du zwischen Werten wechseln (z.B. verschiedene Server):

1. Zahnrad ⚙️ → Variables
2. **"Add variable"**
3. Name: `instance`
4. Type: `Query`
5. Data source: `Prometheus`
6. Query: `label_values(up, instance)`
7. Save

Jetzt kannst du die Variable in Queries nutzen:
```promql
up{instance="$instance"}
```

### Auto-Refresh einstellen
1. Oben rechts neben Zeit-Selector
2. Klick auf "off"
3. Wähle z.B. "30 seconds"
4. Dashboard aktualisiert automatisch

### Zeit-Range ändern
1. Oben rechts "Last 6 hours" (oder ähnlich)
2. Wähle andere Zeit-Range
3. Oder gib eigene Werte ein

---

## 🆘 Häufige Probleme

### Problem: Query zeigt keine Daten
- [ ] Ist die Metrik wirklich vorhanden? (http://localhost:9090/targets)
- [ ] Zeit-Range stimmt? (Daten könnten älter sein)
- [ ] Syntax korrekt? (PromQL Case-sensitive!)

**Lösung:** Query im Prometheus Graph Tab (http://localhost:9090) testen

### Problem: Panel sieht blöd aus
- [ ] Andere Visualization wählen (Stat, Gauge, Pie, etc.)
- [ ] Unit ändern
- [ ] Min/Max anpassen
- [ ] Thresholds justieren

### Problem: Dashboard lädt zu lange
- [ ] Zu viele Panels? (Maximum ~15)
- [ ] Zu lange Zeit-Range? (Wechsel auf "Last 24 hours")
- [ ] Zu komplexe Queries? (Vereinfachen)

---

## ✨ Coole Features zum Ausprobieren

### Annotations
Markiere besondere Events auf dem Graph (z.B. Deployments)

### Alerts
Lass Grafana dich benachrichtigen, wenn ein Threshold überschritten wird

### Templating
Benutzer können Variablen wechseln (z.B. verschiedene Apps)

---

## 📚 Nächste Schritte

1. Mehrere Dashboards erstellen (pro Team/App eine)
2. Dashboards exportieren und teilen
3. Alerts in Grafana einrichten
4. Mit AlertManager integrieren

---

## 📋 Checkliste: Dein erstes Dashboard

- [ ] Grafana geöffnet
- [ ] Neues Dashboard erstellt
- [ ] Status Panel (uptime)
- [ ] CPU Panel
- [ ] RAM Panel
- [ ] Request-Rate Panel
- [ ] Error-Rate Panel
- [ ] Panels angeordnet
- [ ] Dashboard gespeichert
- [ ] Auto-Refresh aktiviert

**Gratuliere! Du hast dein erstes Grafana-Dashboard gebaut!** 🎉

---

**Nächste Datei:** HANDOUT_04_GLOSSAR.md

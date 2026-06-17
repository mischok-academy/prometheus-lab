# PromQL Cheat Sheet

**Die wichtigsten Prometheus-Queries auf einer Seite**

---

## 📌 Basis-Queries

### Einfache Metrik anschauen
```promql
node_cpu_seconds_total
```
→ Zeigt: Alle CPU-Sekunden (riesige Zahl, nicht aussagekräftig!)

### Mit Label filtern
```promql
node_cpu_seconds_total{mode="idle"}
```
→ Zeigt: Nur CPU-Sekunden im "idle" Modus

### Mehrere Label filtern
```promql
up{job="prometheus"}
up{job="node", instance="localhost:9100"}
```

### Regex-Filter (fortgeschritten)
```promql
node_cpu_seconds_total{mode=~"idle|system"}
```
→ Mode ist "idle" ODER "system"

---

## 🔧 Wichtigste Funktionen

### rate() - Rate für Counter
```promql
# Wie viele Requests pro Sekunde in letzten 5 Min?
rate(app_requests_total[5m])

# Wie viel RAM wird pro Sekunde freigegeben?
rate(node_memory_MemAvailable_bytes[5m])
```
**Wann nutzen:** Bei Counter (Zähler), um saubere Werte zu bekommen

### increase() - Steigung
```promql
# Wie viele Requests in der letzten Stunde insgesamt?
increase(app_requests_total[1h])
```

### irate() - Momentane Rate
```promql
# CPU-Auslastung jetzt (schnelle, kurzfristige Änderung)
irate(node_cpu_seconds_total[5m])
```

---

## 📊 Praktische Formeln

### CPU-Auslastung in %
```promql
100 - (avg by (instance) (irate(node_cpu_seconds_total{mode="idle"}[5m])) * 100)
```
**Was bedeutet das:**
- `irate()` = momentane Rate
- `mode="idle"` = ungenutzter CPU
- `100 - ...` = nutzte CPU ist 100% minus ungenutzter
- Ergebnis: 0-100%

### RAM-Auslastung in %
```promql
(1 - (node_memory_MemAvailable_bytes / node_memory_MemTotal_bytes)) * 100
```
**Was bedeutet das:**
- Verfügbar / Total = Frei-Prozentsatz
- `1 - ...` = Genutzt-Prozentsatz
- `* 100` = In Prozent umrechnen

### Festplatte voll in %
```promql
(1 - (node_filesystem_avail_bytes / node_filesystem_size_bytes)) * 100
```

### Request-Fehlerrate in %
```promql
(sum(rate(app_errors_total[5m])) / sum(rate(app_requests_total[5m])) * 100)
```
→ Zeigt: Wie viel Prozent der Requests sind Fehler?

### Response-Zeit 95. Perzentil
```promql
histogram_quantile(0.95, rate(app_request_duration_seconds_bucket[5m]))
```
→ Zeigt: 95% der Requests sind schneller als X Sekunden

---

## 🎯 Aggregations-Funktionen

### Alle Instanzen zusammen
```promql
# Summe über alle Instanzen
sum(node_memory_MemTotal_bytes)

# Durchschnitt
avg(node_memory_MemTotal_bytes)

# Maximum
max(node_memory_MemTotal_bytes)

# Minimum
min(node_memory_MemTotal_bytes)
```

### Nach Label gruppieren
```promql
# Summe pro Job
sum by (job) (up)

# Durchschnitt pro Instance
avg by (instance) (node_cpu_seconds_total)

# Mehrere Labels
sum by (job, instance) (app_requests_total)
```

### Top 3
```promql
# Die 3 Instanzen mit meisten Speicher
topk(3, node_memory_MemTotal_bytes)

# Die 3 mit wenig Speicher
bottomk(3, node_memory_MemAvailable_bytes)
```

---

## ⏱️ Zeit-Fenster

| Fenster | Code | Wann nutzen |
|---------|------|-----------|
| 1 Minute | `[1m]` | Schnelle Reaktion |
| 5 Minuten | `[5m]` | Standard (bei 15s Scrape) |
| 1 Stunde | `[1h]` | Trends über Zeit |
| 1 Tag | `[1d]` | Langfristige Trends |

**Faustregel:** Mindestens `[2m]`, besser `[5m]` oder `[15m]`

---

## ✅ Häufige Fehler vermeiden

### ❌ FALSCH: Counter-Wert direkt nutzen
```promql
app_requests_total  # NEIN! Zahl ist riesig und steigt immer
```

### ✅ RICHTIG: rate() verwenden
```promql
rate(app_requests_total[5m])  # JA! Zeigt Requests pro Sekunde
```

---

### ❌ FALSCH: Zu kurzes Fenster
```promql
rate(app_requests_total[10s])  # NEIN! Bei 15s Scrape zu kurz
```

### ✅ RICHTIG: Ausreichend langes Fenster
```promql
rate(app_requests_total[5m])  # JA! Mehrere Datenpunkte
```

---

## 🔍 Debug-Queries

### Alle Metriken-Namen sehen
```promql
# Gehe zu "Graph" → Klick auf Feld → Alle Optionen expandieren
# Oder verwende Prometheus UI Autocomplete (Strg+Space)
```

### Targets und Scrape-Status
```
http://localhost:9090/targets
```
→ Zeigt welche Services hochladen und welche down sind

### Test deine Query
```
http://localhost:9090/graph
```
→ Hier kannst du jede Query testen bevor du sie ins Dashboard packst

---

## 📋 Copy-Paste Template

Für schnelle Dashboards:

```promql
# Oben: Uptime
up

# CPU
100 - (avg by (instance) (irate(node_cpu_seconds_total{mode="idle"}[5m])) * 100)

# RAM
(1 - (node_memory_MemAvailable_bytes / node_memory_MemTotal_bytes)) * 100

# Requests pro Sekunde
rate(app_requests_total[5m])

# Fehlerquote
rate(app_errors_total[5m])

# Latenz (p95)
histogram_quantile(0.95, rate(app_request_duration_seconds_bucket[5m]))
```

Kopiere diese Queries in Grafana und baue dein Dashboard!

---

## 🎓 Übungen

### Übung 1: CPU-Auslastung
Schreibe die Query für CPU-Auslastung in %:
<details>
<summary>💡 Lösung</summary>

```promql
100 - (avg(irate(node_cpu_seconds_total{mode="idle"}[5m])) * 100)
```
</details>

### Übung 2: Speicher pro Server
Zeige RAM-Auslastung gruppiert nach Instance:
<details>
<summary>💡 Lösung</summary>

```promql
(1 - (node_memory_MemAvailable_bytes / node_memory_MemTotal_bytes)) * 100
```
(Wird automatisch pro Instance angezeigt)
</details>

### Übung 3: Error-Rate
Wie viel % der App-Requests sind Fehler?
<details>
<summary>💡 Lösung</summary>

```promql
(sum(rate(app_errors_total[5m])) / sum(rate(app_requests_total[5m]))) * 100
```
</details>

---

## 📞 Schnelle Hilfe

**Frage:** "Meine Query zeigt nichts"  
**Antwort:** Prüf ob die Metrik existiert → http://localhost:9090/targets

**Frage:** "Warum so große Zahlen?"  
**Antwort:** Counter nutzen rate() oder increase()

**Frage:** "Kann ich Metriken filtern?"  
**Antwort:** Ja! Mit `{label="wert"}` am Ende

---

**Nächste Datei:** HANDOUT_03_GRAFANA_TUTORIAL.md

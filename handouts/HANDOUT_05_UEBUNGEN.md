# Praktische Übungen: Prometheus Monitoring

**Wende dein Wissen an und löse diese Aufgaben!**

---

## ⏱️ Schwierigkeit & Zeit

- 🟢 **Leicht** (5-10 min): Ideal zum Starten
- 🟡 **Mittel** (15-20 min): Erfordert etwas Nachdenken
- 🔴 **Schwer** (30+ min): Für Profis

---

## 🟢 Übung 1: Die erste Query

**Thema:** PromQL Basics  
**Schwierigkeit:** 🟢 Leicht (5 Min)  
**Ziel:** Deine erste Query schreiben

### Aufgabe
Schreibe eine PromQL-Query, die zeigt, welche Targets gerade UP sind.

### Tipps
- Ein Target hat eine `up` Metrik
- Wert 1 = UP, Wert 0 = DOWN

### Lösung
<details>
<summary>💡 Hier klicken für Lösung</summary>

```promql
up
```

**Erklärung:** Diese einfache Query zeigt den Status aller Targets.

**Erweitert:** Um nur UP-Targets zu zeigen:
```promql
up == 1
```
</details>

### Bonus
Schreibe eine Query, die nur den Node-Exporter zeigt:
<details>
<summary>💡 Bonus-Lösung</summary>

```promql
up{job="node"}
```
</details>

---

## 🟢 Übung 2: CPU in Prometheus testen

**Thema:** Rate-Funktion  
**Schwierigkeit:** 🟢 Leicht (10 Min)  
**Ziel:** CPU-Auslastung visualisieren

### Aufgabe
1. Öffne Prometheus Graph: http://localhost:9090
2. Schreibe eine Query, die CPU-Auslastung in % über die letzten 5 Minuten zeigt
3. Visualisiere es als Graph
4. Beobachte wie die CPU-Kurve sich bewegt

### Tipps
- Nutze die CPU-Query aus dem Cheat Sheet
- Zeitfenster sollte mindestens [5m] sein

### Lösung
<details>
<summary>💡 Hier klicken für Lösung</summary>

```promql
100 - (avg(irate(node_cpu_seconds_total{mode="idle"}[5m])) * 100)
```

**Erklärung:**
- `node_cpu_seconds_total` = CPU-Sekunden
- `{mode="idle"}` = ungenutzter CPU
- `irate()` = momentane Rate
- `100 - ...` = nutzte CPU
- `* 100` = in Prozent

**Was siehst du:** Eine Kurve zwischen 0-100%, die sich in Echtzeit ändert
</details>

---

## 🟢 Übung 3: RAM-Nutzung Dashboard

**Thema:** Grafana Dashboard  
**Schwierigkeit:** 🟢 Leicht (15 Min)  
**Ziel:** Ein einfaches Dashboard bauen

### Aufgabe
1. Öffne Grafana: http://localhost:3000
2. Erstelle ein neues Dashboard
3. Füge ein Panel hinzu mit RAM-Auslastung in %
4. Setze die Visualization auf "Gauge"
5. Speichere das Dashboard als "RAM Monitor"

### Tipps
- RAM-Query findest du im Cheat Sheet
- Unit sollte "Percent" sein
- Min: 0, Max: 100

### Lösung
Folge dem HANDOUT_03_GRAFANA_TUTORIAL.md Schritt für Schritt (Schritt 5)

---

## 🟡 Übung 4: Fehlerrate der Sample-App

**Thema:** Aggregation & Mathematik  
**Schwierigkeit:** 🟡 Mittel (20 Min)  
**Ziel:** Error-Rate berechnen

### Aufgabe
1. Öffne Prometheus Graph
2. Schreibe eine Query, die die Fehlerrate der Sample-App in % berechnet
3. Was ist die aktuelle Fehlerrate?
4. Mache mehrere Requests zu http://localhost:8888 und schau wie die Quote sich ändert

### Tipps
- Fehlerrate = Fehler / Gesamt
- Nutze `sum()` um Total zu berechnen
- Die Metriken sind `app_errors_total` und `app_requests_total`

### Lösung
<details>
<summary>💡 Hier klicken für Lösung</summary>

```promql
(sum(rate(app_errors_total[5m])) / sum(rate(app_requests_total[5m]))) * 100
```

**Erklärung:**
- Zähler: `sum(rate(app_errors_total[5m]))` = Fehler pro Sekunde
- Nenner: `sum(rate(app_requests_total[5m]))` = Requests pro Sekunde
- Division = Fehlerquote
- `* 100` = in Prozent

**Beispiel Ergebnis:** 5 bedeutet 5% der Requests sind Fehler
</details>

---

## 🟡 Übung 5: Alert schreiben

**Thema:** Alert-Regeln  
**Schwierigkeit:** 🟡 Mittel (25 Min)  
**Ziel:** Alert-Regel verstehen und testen

### Aufgabe
1. Öffne `alert-rules.yml` im prometheus-lab Verzeichnis
2. Schaue dir die bestehenden Alerts an
3. Füge einen neuen Alert hinzu: "SampleAppHighErrorRate"
   - Schwellenwert: > 10% Fehlerrate
   - Dauer: 2 Minuten
   - Severity: warning
4. Speichern und Prometheus neuladen: `curl -X POST http://localhost:9090/-/reload`
5. Gehe zu http://localhost:9090/alerts und schau ob der Alert existiert

### Struktur
```yaml
- alert: SampleAppHighErrorRate
  expr: (sum(rate(app_errors_total[5m])) / sum(rate(app_requests_total[5m])) * 100) > 10
  for: 2m
  labels:
    severity: warning
  annotations:
    summary: "Sample App hat hohe Fehlerrate"
    description: "Fehlerrate ist {{ $value }}%"
```

### Lösung
Einfügen in `alert-rules.yml` unter dem `app_alerts` Section

---

## 🟡 Übung 6: Custom Dashboard mit 3 Panels

**Thema:** Grafana (Fortgeschritten)  
**Schwierigkeit:** 🟡 Mittel (30 Min)  
**Ziel:** Professionelles Monitoring-Dashboard

### Aufgabe
Erstelle ein Dashboard "App Performance Monitor" mit:

1. **Panel 1: Request Rate** (Graph)
   - Query: `rate(app_requests_total[5m])`
   - Unit: Requests/sec
   
2. **Panel 2: Error Rate** (Stat)
   - Query: `(sum(rate(app_errors_total[5m])) / sum(rate(app_requests_total[5m]))) * 100`
   - Unit: Percent
   - Thresholds: < 5% green, > 5% red

3. **Panel 3: Response Time P95** (Gauge)
   - Query: `histogram_quantile(0.95, rate(app_request_duration_seconds_bucket[5m]))`
   - Unit: seconds
   - Max: 1

### Anforderungen
- [ ] Dashboard heißt "App Performance Monitor"
- [ ] 3 Panels haben aussagekräftige Titel
- [ ] Alle Queries funktionieren und zeigen Daten
- [ ] Dashboard ist gespeichert

### Lösung
Folge HANDOUT_03_GRAFANA_TUTORIAL.md für die detaillierten Schritte

---

## 🔴 Übung 7: Alerting Integration

**Thema:** Vollständiges Alerting-System  
**Schwierigkeit:** 🔴 Schwer (45 Min)  
**Ziel:** Alerts bis zum Versand verstehen

### Aufgabe
1. **Alert in Prometheus definieren**
   ```yaml
   - alert: NodeHighCPU
     expr: 100 - (avg(irate(node_cpu_seconds_total{mode="idle"}[5m])) * 100) > 50
     for: 1m
   ```

2. **AlertManager konfigurieren**
   - Öffne `alertmanager.yml`
   - Konfiguriere einen einfachen Webhook Receiver

3. **Testen**
   - Triggere Alert: `docker-compose exec -T node-exporter stress --cpu 1 --timeout 120s`
   - Schau Prometheus Alerts: http://localhost:9090/alerts
   - Schau AlertManager: http://localhost:9093

4. **Dokumentieren**
   - Wann wurde der Alert getriggert?
   - Welcher Severity-Level?
   - Wie lange war er aktiv?

### Bonus: Slack-Integration
- Erstelle Slack Webhook
- Trage URL in `alertmanager.yml` ein
- Erhalte Slack-Notifications

### Lösung
Siehe `alertmanager.yml` - dort sind Kommentare für Slack/Email

---

## 🔴 Übung 8: Eigene Metrik hinzufügen

**Thema:** Custom Metriken  
**Schwierigkeit:** 🔴 Schwer (60 Min)  
**Ziel:** Verstehe wie man Metriken exportiert

### Aufgabe
1. **Sample-App modifizieren**
   - Öffne `sample-app/main.go`
   - Füge eine neue Metrik `app_uptime_seconds` hinzu
   - Sie sollte zeigen wie lange die App läuft

2. **Rebuild**
   ```bash
   docker-compose restart sample-app
   ```

3. **Testen**
   - Öffne http://localhost:8888/metrics
   - Ist die neue Metrik sichtbar?
   - Erhöht sich der Wert?

4. **Visualisieren**
   - Erstelle Panel in Grafana mit der neuen Metrik
   - Visualization: Stat (große Zahl)

### Hint
Nutze eine Variable mit `time.Now()` und `time.Since()` zur Uptime-Berechnung

### Lösung
<details>
<summary>💡 Hier klicken für Lösung</summary>

In `main.go` nach `recordMetrics()`:

```go
var startTime = time.Now()

func metricsHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	
	uptime := int64(time.Since(startTime).Seconds())
	
	// ... andere Metriken ...
	
	fmt.Fprintf(w, "# HELP app_uptime_seconds Application uptime\n")
	fmt.Fprintf(w, "# TYPE app_uptime_seconds gauge\n")
	fmt.Fprintf(w, "app_uptime_seconds %d\n", uptime)
}
```
</details>

---

## 🟡 Übung 9: PromQL Treasure Hunt

**Thema:** PromQL Fortgeschritten  
**Schwierigkeit:** 🟡 Mittel (30 Min)  
**Ziel:** Verschiedene PromQL-Techniken kombinieren

### Aufgabe: Finde folgende Informationen mit PromQL

1. **Welcher Service hat den meisten RAM?**
   <details>
   <summary>💡 Lösung</summary>
   
   ```promql
   topk(1, container_memory_usage_bytes{name!=""})
   ```
   </details>

2. **Wie viele Requests hatte die Sample-App in der letzten Stunde insgesamt?**
   <details>
   <summary>💡 Lösung</summary>
   
   ```promql
   increase(app_requests_total[1h])
   ```
   </details>

3. **Welcher Service war am längsten DOWN?**
   <details>
   <summary>💡 Lösung (schwer)</summary>
   
   Brauchst du die Logs für - PromQL kann das nicht direkt
   </details>

4. **Wie viel GB RAM sind in Summe verfügbar?**
   <details>
   <summary>💡 Lösung</summary>
   
   ```promql
   sum(node_memory_MemAvailable_bytes) / 1024 / 1024 / 1024
   ```
   </details>

5. **Welcher Port der Sample-App wird am meisten genutzt?**
   <details>
   <summary>💡 Lösung</summary>
   
   ```promql
   topk(1, app_processed_items)
   ```
   </details>

---

## 🟢 Übung 10: Troubleshooting Challenge

**Thema:** Debugging  
**Schwierigkeit:** 🟢 Leicht (15 Min)  
**Ziel:** Probleme diagnostizieren

### Szenario
Ein Service antwortet nicht mehr. Was machst du?

### Schritte
1. **Target-Status prüfen**
   ```
   http://localhost:9090/targets
   ```
   - Ist der Service DOWN?

2. **Logs anschauen**
   ```bash
   docker-compose logs SERVICE_NAME
   ```
   - Gibt es Error-Messages?

3. **Service neu starten**
   ```bash
   docker-compose restart SERVICE_NAME
   ```

4. **Status überprüfen**
   ```bash
   docker-compose ps
   ```

### Aufgabe
Starte diese Challenge:

```bash
# 1. Stoppe einen Service bewusst
docker-compose stop sample-app

# 2. Was siehst du in Prometheus Targets?
# Beantworte: Ist der Service als DOWN sichtbar?

# 3. Starte ihn wieder
docker-compose start sample-app

# 4. Wie lange dauert es bis er wieder als UP angezeigt wird?
```

### Lösung
<details>
<summary>💡 Hier klicken für Erklärung</summary>

**Was passiert:**
1. Wenn Service stoppt: Prometheus kann nicht connecten → DOWN
2. AlertManager sendet Alert nach `for:` Zeit
3. Wenn Service wieder started: Nach 1-2 Scrape-Cycles (15-30s) → UP

**Das ist normal!**
</details>

---

## 📊 Übersicht: Alle Übungen

| # | Titel | Schwierigkeit | Zeit | Topics |
|---|-------|---------------|------|--------|
| 1 | Erste Query | 🟢 | 5 min | PromQL Basics |
| 2 | CPU in Prometheus | 🟢 | 10 min | rate(), irate() |
| 3 | RAM Dashboard | 🟢 | 15 min | Grafana Basics |
| 4 | Fehlerrate | 🟡 | 20 min | Aggregation |
| 5 | Alert schreiben | 🟡 | 25 min | Alert Rules |
| 6 | App Dashboard | 🟡 | 30 min | Grafana Pro |
| 7 | Alerting System | 🔴 | 45 min | Vollständiges System |
| 8 | Custom Metrik | 🔴 | 60 min | Code ändern |
| 9 | PromQL Treasure Hunt | 🟡 | 30 min | PromQL Mix |
| 10 | Troubleshooting | 🟢 | 15 min | Debugging |

---

## ✅ Empfohlene Reihenfolge

**Tag 1 (nach Vortrag):**
- Übung 1 ✅
- Übung 2 ✅
- Übung 3 ✅

**Tag 2 (nächster Tag):**
- Übung 4 ✅
- Übung 5 ✅
- Übung 10 ✅

**Wenn Zeit / Bonus:**
- Übung 6 ✅
- Übung 9 ✅

**Fortgeschritten:**
- Übung 7 ✅
- Übung 8 ✅

---

## 💡 Tipps zum Lernen

1. **Schreibe Queries selbst** - Kopieren nutzt dir nichts
2. **Experimentiere** - Prometheus-Abfragen sind sicher
3. **Fehler sind ok** - Error-Messages sind Lernmaterial
4. **Dokumentiere** - Schreibe auf welche Query welches zeigt
5. **Teile mit anderen** - Erkläre deine Lösung jemandem

---

## 🏆 Abschluss-Challenge

Wenn du alle Übungen gemacht hast:

**Erstelle dein eigenes Monitoring-Dashboard für einen beliebigen Service:**
1. Wähle einen Service (z.B. Database, Cache, Queue)
2. Recherchiere welche Metriken wichtig sind
3. Schreibe die PromQL-Queries
4. Baue ein Dashboard in Grafana
5. Schreibe Alerts dafür
6. Dokumentiere alles

**Bonus:** Präsentiere dein Dashboard den anderen Teilnehmern! 🎉

---

**Nächste Datei:** HANDOUT_06_TROUBLESHOOTING.md

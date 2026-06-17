# Glossar: Prometheus Monitoring Begriffe

**Alle wichtigen Begriffe, die du kennen solltest**

---

## A

### Aggregation
**Definition:** Zusammenfassung von Daten  
**Beispiel:** `sum()` addiert mehrere Metriken zusammen  
**Im Kontext:** "Aggregiere alle CPU-Werte über alle Server"

### Alert
**Definition:** Warnung/Benachrichtigung wenn Schwellenwert überschritten  
**Beispiel:** "Sende Email wenn CPU > 80%"  
**Im Kontext:** AlertManager versendet Alerts

### AlertManager
**Definition:** Komponente von Prometheus, die Alerts verwaltet  
**Aufgaben:** Empfangen, Filtern, Gruppieren, Versenden  
**Port:** 9093

---

## B

### Bucket
**Definition:** Behälter im Histogram für einen Wertebereich  
**Beispiel:** Bucket "0-0.1 Sekunden", "0.1-1 Sekunde", etc.  
**Im Kontext:** Response-Zeit in verschiedene Buckets aufteilen

---

## C

### cAdvisor
**Definition:** Container Advisor - Tool zur Container-Überwachung  
**Funktion:** Sammelt Docker-/Kubernetes-Metriken  
**Port:** 8080  
**Alternative zu:** Prometheus Container-Exporter

### Counter
**Definition:** Metriken-Typ, der nur nach oben zählt  
**Beispiel:** Gesamt-HTTP-Requests seit Serverstart  
**Regel:** Nur 0 oder Anstieg, nie Rückgang  
**Nutze rate() darauf!**

---

## D

### Dashboard
**Definition:** Sammlungen von Visualisierungen  
**Wo:** Grafana  
**Nutzen:** Übersicht über mehrere Metriken auf einen Blick  
**Beispiel:** "Production Cluster Overview"

### Datenquelle / Data Source
**Definition:** Konfiguration einer Prometheus-Instanz in Grafana  
**Zweck:** Grafana weiß wo Daten zu finden sind  
**Im Kontext:** "Prometheus" ist eine Datenquelle

---

## E

### Exporter
**Definition:** Software, die Metriken im Prometheus-Format bereitstellt  
**Arten:**
- Node Exporter (Host-Metriken)
- Custom Exporter (App-Metriken)
- Database Exporter (DB-Metriken)

**Port:** Verschieden (Node: 9100, custom: beliebig)

---

## F

### Gauge
**Definition:** Metriken-Typ, der beliebig hoch/runter gehen kann  
**Beispiel:** CPU-Auslastung (45% jetzt, dann 60%)  
**Im Kontext:** "Die CPU-Gauge zeigt derzeit 75%"

---

## G

### Grafana
**Definition:** Visualisierungs- und Dashboard-Tool  
**Kompatibel mit:** Prometheus, InfluxDB, Elasticsearch, etc.  
**Hauptnutzen:** Schöne Dashboards bauen  
**Port:** 3000

---

## H

### Histogram
**Definition:** Metriken-Typ, der Wertverteilung zeigt  
**Beispiel:** Response-Zeiten: 100 Requests < 0.1s, 500 < 1s, 50 > 5s  
**Nutze für:** Latenz, Durchsatz, Größen  
**Funktion:** `histogram_quantile()` um Perzentile zu berechnen

### Host
**Definition:** Ein Server/Computer in deinem System  
**Im Kontext:** "Der Host ubuntu-server-1 hat 80% RAM-Auslastung"

---

## I

### Instance
**Definition:** Eine einzelne Metrik-Quelle (Server, Pod, etc.)  
**Label:** Immer in Metriken vorhanden `{instance="server:9100"}`  
**Im Kontext:** "Die Instance prometheus:9090 ist down"

### irate()
**Definition:** Instant Rate - momentane Änderungsrate  
**Unterschied zu rate():** Berücksichtigt nur letzte 2 Datenpunkte  
**Wann nutzen:** Für schnelle, kurzfristige Änderungen

---

## J

### Job
**Definition:** Gruppe von Instances (gleicher Service-Typ)  
**Beispiel:** Job "node" hat Instances [server-1, server-2, server-3]  
**Label:** Immer in Metriken vorhanden `{job="node"}`  
**Im Kontext:** "Alle Instances im Job 'prometheus' sind up"

---

## K

### KPI
**Definition:** Key Performance Indicator - wichtige Messgröße  
**Beispiel KPIs:**
- Verfügbarkeit > 99.9%
- Response-Time < 200ms
- Error-Rate < 1%

**Im Kontext:** "Das KPI für Verfügbarkeit ist überschritten"

---

## L

### Label
**Definition:** Tag-Name für eine Metrik (Key-Value Pair)  
**Beispiel:** `{method="GET", status="200", instance="server-1"}`  
**Nutze für:** Filterung und Gruppierung  
**Syntax:** `metric_name{label1="wert1", label2="wert2"}`

---

## M

### Metrik
**Definition:** Messwert mit Zeitstempel  
**Teile:** Name + Labels + Wert + Zeit  
**Beispiel:** "CPU bei 14:00 = 45%"  
**Im Kontext:** "Prometheus speichert Zeitreihen-Metriken"

### Monitoring
**Definition:** Kontinuierliche Überwachung von Systemen  
**Ziele:**
- Probleme früh erkennen
- Performance nachvollziehen
- Verfügbarkeit sicherstellen

---

## N

### Node Exporter
**Definition:** Exporter für Host-System-Metriken  
**Metriken:** CPU, RAM, Disk, Network, etc.  
**Port:** 9100  
**Wer entwickelt:** Prometheus-Projekt

---

## O

### Observer Pattern
**Definition:** Prometheus-Architektur: Server zieht Daten ab  
**Im Kontext:** "Prometheus observiert/beobachtet alle Services"

---

## P

### Percentile
**Definition:** Wert unter dem X% aller Messwerte liegen  
**Beispiele:**
- P50 (Median): Hälfte ist schneller, Hälfte langsamer
- P95: 95% sind schneller als dieser Wert
- P99: 99% sind schneller als dieser Wert

**Funktion:** `histogram_quantile(0.95, ...)`

### PromQL
**Definition:** Prometheus Query Language  
**Nutzen:** Daten aus Prometheus abfragen und filtern  
**Syntax:** `metric{label="wert"}` + Funktionen

### Prometheus
**Definition:** Zeit-Reihen-Datenbank und Monitoring-System  
**Funktion:** Speichert Metriken und beantwortet Abfragen  
**Port:** 9090  
**Besonderheit:** Pull-basiert (zieht Daten ab)

### Pushgateway
**Definition:** Komponente zum Empfang von Metriken (Push statt Pull)  
**Nutze für:** Kurzlebige Batch-Jobs, die Metriken pushen  
**Port:** 9091

---

## Q

### Query
**Definition:** Abfrage nach bestimmten Metriken  
**Beispiel:** `rate(http_requests_total[5m])`  
**Sprache:** PromQL  
**Im Kontext:** "Schreib eine Query, die CPU zeigt"

---

## R

### Regex
**Definition:** Regulärer Ausdruck zum Filtern  
**Syntax:** `metric{label=~"pattern"}`  
**Beispiel:** `{job=~"api|web"}` = Job ist "api" ODER "web"

### Rate
**Definition:** Änderungsrate pro Sekunde  
**Funktion:** `rate(counter[5m])`  
**Wann nutzen:** Bei Counter, um aussagekräftige Werte zu bekommen  
**Im Kontext:** "Rate ist 1000 Requests/Sekunde"

### Retention
**Definition:** Wie lange Prometheus Daten speichert  
**Default:** 15 Tage  
**Einstellung:** `--storage.tsdb.retention.time=30d`  
**Im Kontext:** "Nach 30 Tagen werden alte Metriken gelöscht"

---

## S

### Scrape
**Definition:** Prozess zum Abrufen von Metriken  
**Ablauf:**
1. Prometheus geht zu `http://service:port/metrics`
2. Liest rohe Metriken
3. Speichert mit Zeitstempel

**Interval:** Standard 15 Sekunden  
**Im Kontext:** "Der Prometheus Scrape-Interval ist 30 Sekunden"

### Service Discovery
**Definition:** Automatische Erkennung von Services zum Scrapen  
**Typen:** Statisch, Kubernetes, Consul, AWS, etc.  
**Im Kontext:** Lab nutzt statische Config

### SINGLE Point of Failure
**Definition:** Ein Element, dessen Ausfall das ganze System bricht  
**Beispiel:** Nur ein Prometheus = Single Point of Failure  
**Lösung:** Mehrere Prometheus-Instanzen + Remote Storage

### Summary
**Definition:** Metriken-Typ wie Histogram, aber mit Quantilen  
**Unterschied:** Quantile werden client-seitig berechnet  
**Wann nutzen:** Wenn Histogram zu komplex ist

---

## T

### Target
**Definition:** Ein Service/Exporter, von dem Prometheus Metriken holt  
**Beispiel:** Node Exporter auf server-1 ist ein Target  
**Status:** UP oder DOWN  
**Im Kontext:** "Das Target node-exporter zeigt DOWN"

### Threshold
**Definition:** Schwellenwert für Warnung  
**Beispiel:** "Warne wenn CPU > 80%"  
**Wo konfiguriert:** Alert-Regeln in Prometheus oder Grafana  
**Im Kontext:** "Setze einen Threshold für die Fehlerrate"

### Time Series / Zeitreihe
**Definition:** Sequenz von Metriken mit Zeitstempeln  
**Beispiel:**
```
14:00 → CPU = 45%
14:01 → CPU = 48%
14:02 → CPU = 52%
```

**Im Kontext:** "Prometheus ist eine Zeit-Reihen-Datenbank"

### TSDB
**Definition:** Time Series Data Base  
**Im Kontext:** Prometheus nutzt eine lokale TSDB

---

## U

### Uptime
**Definition:** Wie lange ein Service läuft / verfügbar ist  
**Metriken:** "9 days, 5 hours, 30 minutes"  
**Im Kontext:** "Der Service hat 99.99% Uptime"

---

## V

### Visualization
**Definition:** Art wie Daten dargestellt werden  
**Typen in Grafana:**
- Time Series (Graph)
- Stat (große Zahl)
- Gauge (Messuhr)
- Table (Tabelle)
- Pie Chart
- Heatmap

**Im Kontext:** "Wähle die Visualization 'Stat' für dieses Panel"

---

## W

### Web UI
**Definition:** Webinterface von Prometheus  
**URL:** http://localhost:9090  
**Features:** Metrics Browser, Graph, Alerts, Targets

---

## Z

### Zeitfenster
**Definition:** Zeit-Range für Queries  
**Syntax:** `[1m]`, `[5m]`, `[1h]`, `[1d]`  
**Im Kontext:** "Nutze ein Zeitfenster von [5m] für diese Query"

---

## 🔍 Schnelle Nachschlag-Tabelle

| Begriff | Typ | Kurz-Erklärung |
|---------|-----|----------------|
| Counter | Metrik | Nur hoch zählen |
| Gauge | Metrik | Beliebig rauf/runter |
| Histogram | Metrik | Verteilung |
| Summary | Metrik | Quantile |
| rate() | Funktion | Änderung pro Sekunde |
| increase() | Funktion | Steigung über Zeit |
| sum() | Funktion | Addition |
| avg() | Funktion | Durchschnitt |
| Job | Label | Gruppe von Instanzen |
| Instance | Label | Einzelne Quelle |
| Alert | Regel | Warnung bei Schwellenwert |
| Dashboard | UI | Sammlung von Panels |
| Exporter | Software | Gibt Metriken aus |
| PromQL | Sprache | Abfrage-Sprache |

---

**Nächste Datei:** HANDOUT_05_UEBUNGEN.md

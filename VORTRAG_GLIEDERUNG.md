# Prometheus Monitoring
## Vortragsgliederung: Prometheus Monitoring

**Dauer:** 120 Minuten (2 Stunden)  
**Zielgruppe:** Fachinformatiker Systemintegration  
**Format:** Präsentation mit Live-Demo

---

## 📋 Inhaltsübersicht

| Abschnitt | Dauer | Seiten |
|-----------|-------|--------|
| Einstieg & Motivation | 10 min | 1-2 |
| Grundlagen | 20 min | 3-5 |
| Prometheus Architektur | 15 min | 6-8 |
| Installation & Konfiguration | 20 min | 9-11 |
| Live-Demo | 25 min | - |
| PromQL & Abfragen | 15 min | 12-14 |
| Grafana & Visualisierung | 10 min | 15-16 |
| Best Practices & Häufige Fehler | 5 min | 17 |
| Fragen & Diskussion | 5 min | - |

**Gesamt: 125 Minuten**

---

## 1️⃣ Einstieg & Motivation (10 Min)

### 1.1 Problemstellung
- **Szenario:** Server läuft, aber wann merkst du, dass etwas kaputt geht?
- **Herausforderung:** Mehrere Server, verschiedene Services, manuelle Überwachung ist nicht skalierbar
- **Frage an Teilnehmer:** "Wer hat schon mal einen Service komplett nicht bemerkt, der down war?"

### 1.2 Was ist Monitoring?
- Überwachung von IT-Infrastruktur
- Früherkennung von Problemen
- Erhöhung von Verfügbarkeit und Zuverlässigkeit

### 1.3 Warum Prometheus?
- Modernes, Open-Source Monitoring
- Leicht zu lernen und einzusetzen
- Industry-Standard (Netflix, Google, Kubernetes)
- Kostenlos
- Große Community

### 1.4 Lernziele
- Prometheus verstehen und einrichten
- Metriken erfassen und visualisieren
- Alerts konfigurieren
- Praktische Übungen durchführen

**Übergangsfrage:** "Okay, schauen wir uns an, wie das funktioniert..."

---

## 2️⃣ Grundlagen Monitoring (20 Min)

### 2.1 Kern-Begriffe

#### Metriken
- **Definition:** Messwerte von Systemen (Zahlen)
- **Beispiele:**
  - CPU-Auslastung: 75%
  - RAM-Verbrauch: 8 GB / 16 GB
  - Netzwerk-Traffic: 1.2 MB/s
  - HTTP-Requests: 450 pro Sekunde

#### Zeitreihen
- **Definition:** Metriken mit Zeitstempel
- **Beispiel:** CPU um 14:00 = 45%, CPU um 14:01 = 52%, ...
- **Vorteil:** Trends erkennen

#### Indikatoren (KPIs)
- **Definition:** Metriken, die für dich wichtig sind
- **Beispiele:**
  - Verfügbarkeit > 99.9%
  - Response-Zeit < 200ms
  - Error-Rate < 1%

### 2.2 Monitoring-Ansätze

#### Pull-basiert (Prometheus)
```
Prometheus → Server 1 (Metriken abholen)
         → Server 2 (Metriken abholen)
         → Server 3 (Metriken abholen)
```
- ✅ Prometheus kontrolliert den Flow
- ✅ Einfacher zu debuggen
- ✅ Lastverteilung natürlich

#### Push-basiert (Alternative)
```
Server 1 → Datenbank (Metriken senden)
Server 2 → Datenbank (Metriken senden)
Server 3 → Datenbank (Metriken senden)
```
- ❌ Gefahr von Überlastung
- ✅ Für kurzlebige Jobs besser

### 2.3 Metriken-Typen in Prometheus

#### 1. Counter (Zähler)
- Zählt nur nach oben
- Beispiel: Gesamt-Anfragen seit Start
```
http_requests_total = 1000000
```

#### 2. Gauge (Messwert)
- Kann hoch und runter gehen
- Beispiel: aktuelle CPU-Auslastung
```
cpu_usage_percent = 45
```

#### 3. Histogram (Verteilung)
- Verteilt Werte in Buckets
- Beispiel: HTTP-Response-Zeiten
```
response_time = 0.001s, 0.002s, 0.150s, 0.500s, ...
```

#### 4. Summary (Zusammenfassung)
- Ähnlich wie Histogram, aber mit Perzentilen
- Beispiel: 95. Perzentil der Response-Zeit = 200ms

**Merkhilfe:** 
- Counter = Schrittweise nur hoch (Tastendruck-Zähler)
- Gauge = Auto-Tacho (beliebig hoch/runter)

### 2.4 Typischer Ablauf
1. **Exporter** sammelt Metriken vom System
2. **Prometheus** zieht diese Metriken regelmäßig ab
3. **Speicherung** in lokaler Datenbank (TSDB)
4. **Grafana** zeigt die Daten schön an
5. **AlertManager** sendet Alerts bei Problemen

---

## 3️⃣ Prometheus Architektur (15 Min)

### 3.1 Komponenten

```
┌─────────────────────────────────────────────┐
│         Prometheus Server (TSDB)            │
│  - Metriken speichern                       │
│  - PromQL-Abfragen verarbeiten              │
│  - Alerting evaluieren                      │
└──────────────────┬──────────────────────────┘
                   │
        ┌──────────┼──────────┬──────────┐
        │          │          │          │
   ┌────▼────┐ ┌──▼───┐ ┌────▼────┐ ┌─▼────┐
   │ Node    │ │Alert │ │ Grafana │ │Push  │
   │Exporter │ │Manager       │ │Gateway
   │(9100)   │ │(9093)│ │(3000)   │ │(9091)│
   └────┬────┘ └──────┘ └────┬────┘ └──────┘
        │                    │
        │             Visualisierung
   Metriken
```

### 3.2 Prometheus Server
- **Port:** 9090
- **Funktion:** 
  - Zentrale Komponente
  - Speichert Zeitreihen-Daten
  - Evaluiert Alert-Regeln
  - Bietet Web-UI
- **Daten-Retention:** Konfigurierbar (z.B. 30 Tage)

### 3.3 Exporters (Metriken-Sammler)

#### Node Exporter
- Hardware und OS-Metriken
- CPU, RAM, Festplatte, Netzwerk
- Port: 9100

#### cAdvisor
- Container-Metriken
- Docker/Kubernetes
- Port: 8080

#### Custom Exporters
- Deine eigene App
- z.B. Datenbankserver, Webserver

### 3.4 AlertManager
- **Port:** 9093
- **Funktion:**
  - Empfängt Alerts von Prometheus
  - Dedupliziert/gruppiert
  - Sendet Benachrichtigungen (Email, Slack, SMS)
  - Definiert Eskalationen

### 3.5 Grafana
- **Port:** 3000
- **Funktion:**
  - Dashboards erstellen
  - Abfragen visualisieren
  - Mehrere Datenquellen
  - Alerts anzeigen

### 3.6 Scrape-Prozess
```
Prometheus (alle 15 Sekunden):
1. Gehe zur Adresse http://server:9100/metrics
2. Parse die Rohdaten
3. Speichere in lokaler DB mit Zeitstempel
4. Lösche alte Daten nach Retention-Zeit
```

---

## 4️⃣ Installation & Konfiguration (20 Min)

### 4.1 Installation mit Docker Compose
- **Warum Docker?** 
  - Einfach
  - Reproduzierbar
  - Isoliert
  - Perfekt für Lab/Entwicklung

**Schritte zeigen:**
```bash
# 1. Clone lab environment
cd /home/moritzkraus/src/prometheus-lab

# 2. Starten
docker compose up -d

# 3. Status prüfen
docker compose ps
```

### 4.2 Prometheus Konfiguration

**Datei:** `prometheus.yml`

```yaml
global:
  scrape_interval: 15s        # Alle 15 Sekunden scrapen
  evaluation_interval: 15s     # Alert-Regeln alle 15 Sekunden prüfen

scrape_configs:
  - job_name: 'prometheus'    # Sich selbst monitoren
    static_configs:
      - targets: ['localhost:9090']

  - job_name: 'node'          # Node Exporter
    static_configs:
      - targets: ['localhost:9100']
```

### 4.3 Service Discovery
- **Statisch:** IPs hart eingecoded (Lab)
- **Dynamisch:** Automatisch Server finden
  - Consul
  - Kubernetes
  - Cloud-Provider (AWS, Azure, GCP)

### 4.4 Retention-Richtlinie
- **Standard:** 15 Tage
- **Berechnung:** 
  - 1 Serie × 8 Bytes = ~8 Bytes pro Sekunde
  - 1000 Serien × 8 Bytes = ~80 KB pro Sekunde
  - 30 Tage = 2.6 GB Speicherplatz

**Demo:** 
```bash
# Prometheus mit 30 Tagen Retention starten
prometheus --storage.tsdb.retention.time=30d
```

### 4.5 Häufige Konfigurationen
- **Scrape-Timeout:** Wie lange auf Antwort warten (Default: 10s)
- **Scrape-Interval:** Wie oft scrapen (Default: 15s)
- **Labels hinzufügen:**
```yaml
scrape_configs:
  - job_name: 'prod-server'
    static_configs:
      - targets: ['prod.example.com:9100']
        labels:
          environment: 'production'
          team: 'backend'
```

---

## 5️⃣ Live-Demo (25 Min)

### 5.1 Vorbereitung
- Lab ist bereits gestartet
- Alle Container sind laufend

### 5.2 Demo-Ablauf

#### Teil 1: Prometheus Web UI (8 Min)
```
1. Browser öffnen: http://localhost:9090
2. "Graph" Tab zeigen
3. Query schreiben: up
   - Zeigt alle Targets und deren Status
4. Query: node_cpu_seconds_total
   - Zeigt CPU-Metriken
5. Graph anzeigen lassen (letzte Stunde)
6. Targets anzeigen: http://localhost:9090/targets
   - Alle 7 Services sollten grün sein
7. Service ausschalten (z.B. docker compose stop sample-app)
   - Target wird rot (DOWN)
8. Service wieder anschalten
   - Target wird grün
```

#### Teil 2: Sample App Metriken (8 Min)
```
1. Browser: http://localhost:8888/metrics
   - Rohe Prometheus-Metriken zeigen
2. Metriken-Typen identifizieren:
   - app_requests_total (Counter)
   - app_processed_items (Gauge)
   - app_request_duration_seconds (Histogram)
3. Mehrfach Request auf http://localhost:8888
4. Zurück zu Prometheus
5. Query: app_requests_total
   - Rate sehen und erklären
```

#### Teil 3: Erstes Dashboard in Grafana (9 Min)
```
1. Browser: http://localhost:3000 (admin/admin123)
2. Create Dashboard
3. Panel hinzufügen
4. Queries:
   - up (alle Targets)
   - node_memory_MemAvailable_bytes / node_memory_MemTotal_bytes * 100 (RAM%)
   - rate(node_cpu_seconds_total{mode="system"}[5m]) * 100 (CPU)
5. Visualisierungen wechseln (Graph, Table, Gauge)
6. Dashboard speichern
```

### 5.3 Interaktion
- "Wer möchte eine eigene Query schreiben?"
- "Welche Metrik interessiert euch?"

---

## 6️⃣ PromQL & Abfragen (15 Min)

### 6.1 Grundlagen PromQL

**PromQL = Prometheus Query Language**

#### Einfache Queries

```promql
# Alle Metriken eines Types
node_cpu_seconds_total

# Mit Label-Filter
node_cpu_seconds_total{mode="idle"}

# Mit Regex-Filter
node_memory_MemAvailable_bytes{job=~"node|sample-app"}
```

#### Operatoren

```promql
# Arithmetik
node_memory_MemAvailable_bytes / 1024 / 1024 / 1024  # In GB umrechnen

# Vergleiche
up == 1  # Nur "up" Targets
node_cpu_seconds_total{mode="idle"} > 100000
```

### 6.2 Funktionen

#### Rate (für Counter)
```promql
# Wie viele Requests pro Sekunde in den letzten 5 Minuten?
rate(app_requests_total[5m])

# Spitzenwert in letzter Stunde?
max_over_time(app_requests_total[1h])
```

#### Aggregation
```promql
# Durchschnittliche CPU über alle Instanzen
avg(node_cpu_seconds_total)

# Summe aller Requests
sum(rate(app_requests_total[5m]))

# Top 3 Memory-verbraucher
topk(3, node_memory_MemTotal_bytes)
```

#### Prozentile (für Histograms)
```promql
# 95. Perzentil der Response-Zeit
histogram_quantile(0.95, rate(app_request_duration_seconds_bucket[5m]))

# Median (50. Perzentil)
histogram_quantile(0.5, rate(app_request_duration_seconds_bucket[5m]))
```

### 6.3 Praxis-Beispiele

#### CPU-Auslastung in %
```promql
100 - (avg by (instance) (irate(node_cpu_seconds_total{mode="idle"}[5m])) * 100)
```

#### RAM-Auslastung in %
```promql
(1 - (node_memory_MemAvailable_bytes / node_memory_MemTotal_bytes)) * 100
```

#### Fehlerrate in %
```promql
sum(rate(app_errors_total[5m])) / sum(rate(app_requests_total[5m])) * 100
```

#### Request-Dauer 99. Perzentil
```promql
histogram_quantile(0.99, rate(app_request_duration_seconds_bucket[5m]))
```

### 6.4 Häufige Fehler
- ❌ Counter-Summe nutzen (Wert steigt immer)
- ✅ `rate()` verwenden statt raw-Werte
- ❌ Zu kurze Zeit-Range (z.B. `[10s]` bei 15s Interval)
- ✅ Mindestens `[2m]` oder `[5m]` nutzen

---

## 7️⃣ Grafana & Visualisierung (10 Min)

### 7.1 Dashboard-Konzept
- **Was:** Sammlung von Panels/Visualisierungen
- **Verwendung:** Übersicht über Systeme
- **Beispiele:**
  - Prod-Cluster Status
  - App Performance
  - Infrastruktur-Übersicht

### 7.2 Panel-Typen

| Typ | Verwendung | Beispiel |
|-----|-----------|---------|
| Graph | Zeitreihen | CPU über Zeit |
| Gauge | Single Value | RAM-Nutzung % |
| Table | Tabellarisch | Top Errors |
| Heatmap | Intensität | Traffic-Muster |
| Stat | Große Zahl | 99.99% Uptime |

### 7.3 Dashboard Best Practices
- **Titel:** Klar und aussagekräftig
- **Labels:** Zeige Einheit (%, ms, GB)
- **Farben:** Rot = Bad, Grün = Good
- **Größe:** Wichtige Panels größer
- **Layout:** Logische Anordnung (Top Critical, dann Details)

### 7.4 Variable & Templating
```
Dashboard mit Dropdown zum Wechsel zwischen:
- Verschiedenen Servern
- Verschiedenen Umgebungen (prod/staging)
- Verschiedenen Applikationen
```

### 7.5 Alerts in Grafana
- Basierend auf Metriken
- Notification Channels (Email, Slack, PagerDuty)
- Integration mit AlertManager

---

## 8️⃣ Best Practices & Häufige Fehler (5 Min)

### 8.1 Best Practices ✅

| Praktik | Begründung |
|---------|-----------|
| **Beschreibende Metrik-Namen** | `http_request_duration_seconds` nicht `time` |
| **Labels nutzen** | `{service="auth", env="prod"}` statt separater Metriken |
| **Retention planen** | Disk-Platz berücksichtigen |
| **Alerts testen** | Falsche Alerts sind schlimmer als keine |
| **Dokumentation** | Was bedeutet die Metrik? |
| **Regelmäßig aufräumen** | Alte Targets löschen |

### 8.2 Häufige Fehler ❌

| Fehler | Problem | Lösung |
|--------|---------|--------|
| Zu viele Labels | Speicher-Explosion | Nur nötige Labels |
| Zu hoher Scrape-Interval | Alte Daten | 15-30s ist Standard |
| Alerts ohne Kontext | Wer liest das? | Beschreibung + Handlung |
| Keine Redundanz | Single Point of Failure | Remote Storage nutzen |
| Metriken nicht normalisiert | Vergleiche unmöglich | Prozente, Raten, Latenzen |

### 8.3 Monitoring-Regeln (SRE-Konzept)
1. **USE-Methode** (für Ressourcen):
   - **U**tilization: Wie viel nutzen wir?
   - **S**aturation: Wie voll ist es?
   - **E**rrors: Wie viele Fehler?

2. **RED-Methode** (für Services):
   - **R**ate: Requests pro Sekunde
   - **E**rrors: Fehlerrate
   - **D**uration: Response-Zeit

---

## 9️⃣ Fragen & Diskussion (5 Min)

### 9.1 Mögliche Fragen der Teilnehmer
- "Was kostet Prometheus?" → Kostenlos & Open-Source
- "Brauchen wir das in unserem Betrieb?" → Kommt auf Größe an
- "Kann man Prometheus hacken?" → Wie alles, mit Authentifizierung schützen
- "Wo speichert Prometheus die Daten?" → Lokal auf Festplatte
- "Kann man Metriken löschen?" → Nur per Komplett-Reset

### 9.2 Gesprächsstarter
- "Welche Metrik möchtet ihr als erstes monitoren?"
- "Welche Probleme seht ihr im Monitoring von Hand?"
- "Wo seht ihr Monitoring in eurem zukünftigen Job?"

---

## 📚 Materialien & Handouts

### Mitgeben zu den Teilnehmern
1. **Quick Reference Card**
   - PromQL-Cheat-Sheet (A5)
   - Docker-Befehle
   - Wichtige Ports

2. **Lab-Setup Instructions**
   - README.md ausdrucken
   - Docker Compose file
   - Sample Queries

3. **Weitere Ressourcen (Links)**
   - prometheus.io/docs
   - grafana.com/docs
   - YouTube: "Prometheus for Beginners"

### Prüfungsfragen (optional)
1. Erkläre den Unterschied zwischen Push- und Pull-Monitoring
2. Was ist eine Zeitreihe?
3. Schreibe eine PromQL-Query für CPU-Auslastung
4. Wo speichert Prometheus Daten?
5. Warum ist Rate() für Counter wichtig?

---

## 🎯 Lernziele Kontrolle

Nach dem Vortrag sollten Teilnehmer...

- [ ] Prometheus-Architektur erklären können
- [ ] Einfache Metriken-Queries schreiben
- [ ] Ein Dashboard in Grafana erstellen
- [ ] Alert-Regeln verstehen
- [ ] Lab-Umgebung selbst starten können
- [ ] Monitoring-Konzepte auf ihr Projekt anwenden

---

## ⏱️ Zeitplan für Moderator

| Zeit | Aktivität | Notizen |
|------|-----------|---------|
| 00:00 | Start + Einstieg | Motivation |
| 00:10 | Grundlagen | Begriffe klar machen |
| 00:30 | Architektur | Diagramm zeigen |
| 00:45 | Konfiguration | Kurz halten |
| 01:10 | **LIVE-DEMO** | Wichtigste Wahrnehmung! |
| 01:35 | PromQL | Mit Beispielen |
| 01:50 | Grafana + Best Practices | Schnell durchgehen |
| 02:00 | Fragen | Großzügig Zeit nehmen |

---

## 💡 Moderationsmittel

### Visual Aids
- [ ] Folien mit Diagrammen (Architecture)
- [ ] Live-Demo-Umgebung (Docker Compose läuft)
- [ ] Großer Bildschirm für Code/Queries
- [ ] Handouts ausgedruckt

### Interaktivität
- [ ] "Wer hat schon Monitoring benutzt?" (Frage stellen)
- [ ] Teilnehmer ihre eigenen Queries schreiben lassen
- [ ] Gemeinsam Dashboard erstellen
- [ ] "Findet den Fehler" im Config

### Timing-Tricks
- Wenn vorne: Demo zeigen → strafft ab
- Wenn hinten: Demo länger machen → gewinnt Zeit
- Bei 1h 50min: Fragen schneller beantworten

---

## 📞 Support nach dem Vortrag

### Was kann man Teilnehmern in die Hand geben?
1. **Lab-Repository-Link:** prometheus-lab auf GitHub/GitLab
2. **Slack-Channel:** Für Fragen
3. **Kontakt:** Deine Email für Follow-Ups
4. **Bonus-Aufgaben:**
   - Alert für eigene Metrik schreiben
   - Custom Dashboard bauen
   - Zu Hause Lab starten

---

**Viel Erfolg beim Vortrag! 🚀**

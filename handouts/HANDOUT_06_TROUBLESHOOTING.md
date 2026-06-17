# Troubleshooting Guide: Häufige Probleme & Lösungen

**Wenn etwas nicht funktioniert - hier hilft man dir!**

---

## 🔴 Problem: Lab startet nicht

### Symptom
```bash
docker-compose up -d
# Error: ... (irgendein Fehler)
```

### Diagnostik
```bash
# 1. Status überprüfen
docker-compose ps

# 2. Logs anschauen
docker-compose logs
```

### Lösungen

#### Lösung 1: Port ist belegt
**Symptom:** "bind: address already in use"

```bash
# Finde was auf dem Port läuft
lsof -i :9090  # Für Prometheus
lsof -i :3000  # Für Grafana

# Entweder:
# A) Kill den Prozess (risky)
kill -9 PID

# B) Oder warte bis Prozess stoppt
# C) Oder ändere Port in docker-compose.yml
```

#### Lösung 2: Docker läuft nicht
**Symptom:** "Cannot connect to Docker daemon"

```bash
# Docker starten (Linux)
sudo systemctl start docker

# Oder: Desktop-App öffnen (Mac/Windows)
```

#### Lösung 3: Zu wenig Speicher
**Symptom:** Container crashen immer wieder

```bash
# RAM für Docker erhöhen (Docker Desktop Settings)
# Minimum: 4 GB RAM
# Empfohlen: 8 GB RAM
```

---

## 🟡 Problem: Einige Container crashen

### Symptom
```
prometheus-lab      Up
grafana-lab         Up
sample-app-lab      Restarting (1) 5 seconds ago  ← Problem!
```

### Diagnostik
```bash
docker-compose logs sample-app --tail=50
```

### Häufige Fehler

#### Fehler: "Connection refused"
**Bedeutung:** Service versucht sich zu conecten, aber wer hört nicht

**Lösung:**
```bash
# 1. Prüfe ob Target läuft
docker-compose ps

# 2. Starte alle Services
docker-compose restart

# 3. Warte 30 Sekunden, dann nochmal prüfen
```

#### Fehler: "OOM" (Out Of Memory)
**Bedeutung:** Container hat nicht genug RAM

**Lösung:**
```bash
# Erhöhe Docker RAM-Limit
# In docker-compose.yml bei problematischem Service:
deploy:
  resources:
    limits:
      memory: 2G
```

#### Fehler: Exit Code 1
**Bedeutung:** Generischer Fehler, braucht Log-Analyse

**Lösung:**
```bash
# Lies ganze Logs
docker-compose logs SERVICE_NAME

# Suche nach "ERROR" oder "error"
docker-compose logs SERVICE_NAME | grep -i error
```

---

## 🔴 Problem: Prometheus antwortet nicht

### Symptom
- Browser: http://localhost:9090 → Keine Antwort
- Oder: Connection refused

### Diagnostik
```bash
# 1. Läuft der Container?
docker-compose ps prometheus

# 2. Logs anschauen
docker-compose logs prometheus

# 3. Port prüfen
curl http://localhost:9090  # Sollte HTML zurückgeben
```

### Lösungen

#### Lösung 1: Prometheus neu starten
```bash
docker-compose restart prometheus

# Warten Sie 10 Sekunden
sleep 10

# Test
curl http://localhost:9090/-/healthy
```

#### Lösung 2: Config ist kaputt
**Symptom:** "config load error"

```bash
# Prüfe Syntax
cat prometheus.yml

# Oder: Fehler in Logs lesen und Typo fixen
```

#### Lösung 3: Disk voll
**Symptom:** Prometheus speichert keine Daten mehr

```bash
# Prüfe Disk
df -h

# Wenn voll: Docker volumes cleanen
docker-compose down -v
docker system prune -a  # ⚠️ ACHTUNG: Löscht alles!
```

---

## 🟡 Problem: Targets zeigen DOWN

### Symptom
Prometheus → Targets → Service ist RED/DOWN

### Diagnostik
```bash
# 1. Läuft Service im Container?
docker-compose ps

# 2. Kann man die Metriken manuell abholen?
curl http://localhost:9100/metrics  # Node Exporter
curl http://localhost:8888/metrics  # Sample App
```

### Häufige Ursachen

#### Problem: Service ist nicht gestartet
```bash
# Restart
docker-compose restart SERVICE_NAME

# Warte 15 Sekunden (1 Scrape-Intervall)
```

#### Problem: Port ist falsch in config
**Symptom:** prometheus.yml hat falschen Port

```yaml
# FALSCH:
targets: ['localhost:9999']  # ← Port existiert nicht!

# RICHTIG:
targets: ['localhost:9100']  # ← Stimmt mit Service überein
```

**Lösung:**
```bash
# Fix in prometheus.yml
nano prometheus.yml

# Reload Prometheus
curl -X POST http://localhost:9090/-/reload
```

#### Problem: Firewall blockiert
```bash
# Unter Linux: UFW
sudo ufw allow 9090/tcp

# Oder: Docker Network Issue
docker network ls
docker network inspect prometheus-lab_monitoring
```

---

## 🟡 Problem: Grafana zeigt keine Daten

### Symptom
- Grafana lädt
- Aber Dashboards sind leer
- Queries zeigen "No data"

### Diagnostik
```bash
# 1. Grafana läuft?
curl http://localhost:3000/api/health

# 2. Kann sich Grafana mit Prometheus connecten?
# Gehe zu: http://localhost:3000 → Datasources → Prometheus
# → "Test" Button klicken
```

### Lösungen

#### Lösung 1: Datasource ist falsch konfiguriert
**Symptom:** "Datasource error"

```bash
# In Grafana:
# 1. Gehe zu Datasources
# 2. Edit "Prometheus"
# 3. URL muss sein: http://prometheus:9090
#    NICHT: http://localhost:9090 (weil Docker Network!)
```

#### Lösung 2: Metriken existieren nicht
**Symptom:** Query schreibt: "Metric does not exist"

```bash
# 1. Gehe zu Prometheus Graph: http://localhost:9090
# 2. Test deine Query dort
# 3. Wenn da auch nichts kommt → Metrik existiert nicht
```

#### Lösung 3: Zeitfenster ist zu eng
**Symptom:** "No data" aber Metrik existiert

```promql
# FALSCH (zu kurz):
rate(app_requests_total[10s])

# RICHTIG (mindestens 2-3 × Scrape-Interval):
rate(app_requests_total[1m])
```

#### Lösung 4: Scrape-Interval ist zu groß
**Symptom:** Sehr wenige Datenpunkte, Graph sieht blockich aus

```yaml
# In prometheus.yml:
global:
  scrape_interval: 15s  # Weniger als 15s ist gut
```

---

## 🟡 Problem: Alert funktioniert nicht

### Symptom
- Alert-Rule sieht korrekt aus
- Aber Alert wird nicht getriggert
- Oder: AlertManager zeigt "No Alerts"

### Diagnostik
```bash
# 1. Prüfe Alert Status
curl http://localhost:9090/api/v1/alerts | jq

# 2. Prüfe AlertManager
curl http://localhost:9093/api/v1/alerts | jq

# 3. Logs
docker-compose logs prometheus | grep -i alert
```

### Lösungen

#### Lösung 1: Alert-Expr ist falsch
**Symptom:** Alert ist "Inactive"

```bash
# Test die Expression in Prometheus Graph
# Wenn dort kein Ergebnis → Expression ist falsch

# Beispiel FALSCH:
expr: app_errors_total > 1000  # ← Counter! Wert ist riesig

# Beispiel RICHTIG:
expr: rate(app_errors_total[5m]) > 10  # ← Rate!
```

#### Lösung 2: Bedingung ist zu streng
**Symptom:** Alert bleibt "Pending"

```yaml
# Wenn: for: 1m ist und Schwellenwert nur kurz überschritten
# → Alert wird nicht getriggert

# LÖSUNG: Entweder Zeit verkürzen oder Schwellenwert senken
```

#### Lösung 3: AlertManager empfängt Alerts nicht
**Symptom:** Prometheus zeigt Alerts, aber AlertManager ist leer

```bash
# Prüfe AlertManager Config
cat alertmanager.yml

# Sollte Routes definieren
# Starte AlertManager neu
docker-compose restart alertmanager
```

---

## 🔴 Problem: Performance ist schlecht

### Symptom
- Prometheus antwortet langsam
- Dashboard lädt lange
- Grafana wird träge

### Diagnostik
```bash
# CPU Check
docker stats  # Zeigt CPU/RAM Nutzung

# Disk Check
du -h prometheus_data/  # Größe der Datenbank

# Query Performance (in Prometheus):
# Gehe zu Targets/Health und schau auf Scrape-Zeit
```

### Lösungen

#### Lösung 1: Zu viele Metriken
**Problem:** Jede Konfiguration generiert neue Zeitreihen

```yaml
# REDUZIERE scrape_configs oder nutze relabel_configs
relabel_configs:
  - source_labels: [__name__]
    regex: 'go_.*'
    action: drop  # Ignoriere go_* Metriken
```

#### Lösung 2: Queries sind zu komplex
**Problem:** rate()/aggregations in Grafana

**Lösung:**
```promql
# FALSCH: Sehr komplexe Query
rate(sum(rate(app_requests_total{job="app"}[5m]))[10m:5s])

# RICHTIG: Einfache Query
rate(app_requests_total[5m])
```

#### Lösung 3: Storage ist voll
**Problem:** Festplatte ~90% voll

```bash
# Prüfe Größe
du -sh prometheus_data/

# Reduziere Retention-Zeit
# In prometheus.yml:
--storage.tsdb.retention.time=7d  # Statt 30d
```

#### Lösung 4: Zu viele Dashboards
**Problem:** Jedes Dashboard lädt viele Queries

**Lösung:**
- Reduziere Anzahl Panels pro Dashboard
- Nutze Row Collapse (Gruppe Panel in Sections)
- Erhöhe Refresh-Interval (statt "Every 10s" → "Every 30s")

---

## 🟢 Problem: Docker-Compose Syntax-Fehler

### Symptom
```
docker-compose up -d
# Error: error parsing docker-compose.yml
```

### Diagnostik
```bash
# YAML Syntax prüfen
yamllint docker-compose.yml  # Falls yamllint installiert

# Oder: Einfach anschauen und auf Indentation prüfen
```

### Lösungen

#### Häufige Fehler
1. **Falsche Indentation** (Tabs statt Spaces)
   ```yaml
   services:
     prometheus:  # ← 2 Spaces nicht 1 Tab!
   ```

2. **Fehlender Doppelpunkt**
   ```yaml
   ports  # FALSCH
   ports: # RICHTIG
   ```

3. **Unpassende Anführungszeichen**
   ```yaml
   command: "sh -c 'echo hello'"  # Gemischte Quotes ok
   ```

### Fix
```bash
# 1. Nutze Editor mit YAML-Highlighting (VS Code)
# 2. Oder: Online YAML Validator
# 3. Oder: Fragen Sie ChatGPT um Syntax zu überprüfen
```

---

## 🟡 Problem: Metriken sind verrauscht/springen

### Symptom
- Metriken zeigen unlogische Sprünge
- Z.B. CPU springt 0% → 100% → 0%
- Keine smoothe Kurve

### Diagnostik
```promql
# Ist die rohe Metrik auch verrauscht?
# Gehe zu Prometheus Graph und teste:
node_cpu_seconds_total
# vs.
rate(node_cpu_seconds_total[5m])
```

### Lösungen

#### Lösung 1: Rate ist zu kurz
```promql
# FALSCH:
rate(node_cpu_seconds_total[30s])  # Zu kurz!

# RICHTIG:
rate(node_cpu_seconds_total[5m])   # Mindestens 3 Daten-Punkte
```

#### Lösung 2: Counter resettet
**Symptom:** Plötzliche Sprünge nach oben (Dienst neustartete)

```promql
# Nutze increase() statt rate() bei Neustarts
increase(app_requests_total[5m])
```

#### Lösung 3: Nutze Smoothing
```promql
# Glattes Ergebnis über 2 Fenster gemittelt
(rate(metric[5m]) + rate(metric[10m])) / 2
```

---

## 📋 Checkliste: Wenn nichts funktioniert

- [ ] Docker läuft? (`docker ps`)
- [ ] Container sind UP? (`docker-compose ps`)
- [ ] Ports stimmen? (9090, 3000, 8888)
- [ ] Firewall erlaubt Zugriff?
- [ ] Logs gelesen? (`docker-compose logs`)
- [ ] System-Ressourcen ok? (`free -h`, `df -h`)
- [ ] Alles neu gestartet? (`docker-compose restart`)
- [ ] Netzwerk ok? (`docker network ls`)

---

## 🆘 Letzte Rettung: Alles neu

```bash
# ⚠️ WARNUNG: Das löscht ALLE Daten!

# Stoppe alles
docker-compose down

# Lösche Volumes (Daten)
docker-compose down -v

# Starte frisch
docker-compose up -d

# Warte 30 Sekunden
sleep 30

# Prüfe
docker-compose ps
```

---

## 💬 Wenn du nicht weiterkommst

1. **Lies die Logs vollständig** (nicht nur letzte 5 Zeilen)
2. **Googele den Fehler** (in Anführungszeichen)
3. **Frag einen erfahreneren Kollegen**
4. **Stack Overflow oder GitHub Issues**
5. **Prometheus Dokumentation:** prometheus.io/docs

---

## 📞 Support-Kontakt

**Bei Fragen:**
- Lehrer: [Email]
- Projekt: prometheus-lab auf GitHub
- Dokumentation: HANDOUT_01_QUICK_START.md

---

**Nächste Datei:** HANDOUT_07_ZUSAMMENFASSUNG.md

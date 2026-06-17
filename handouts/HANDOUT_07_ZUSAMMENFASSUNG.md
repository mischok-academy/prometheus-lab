# Prometheus Monitoring - Zusammenfassung

**Das musst du wissen - Die Essenz in 1 Seite**

---

## 🎯 Was ist Prometheus?

**Prometheus** = Monitoring-System zur Überwachung von IT-Infrastruktur

**Wie es funktioniert:**
```
Exporter (sammelt)  →  Prometheus (speichert)  →  Grafana (zeigt an)
   Metriken               Zeitreihen               Dashboards
```

---

## 📊 4 Metriken-Typen

| Typ | Wofür | Beispiel | rate()? |
|-----|-------|---------|---------|
| **Counter** | Zähler | HTTP-Requests total | ✅ JA |
| **Gauge** | Messwert | CPU % jetzt | ❌ NEIN |
| **Histogram** | Verteilung | Response-Zeit Buckets | ✅ Für Quantile |
| **Summary** | Perzentile | P95 Response-Zeit | ✅ Ähnlich |

**Merksatz:** Bei Counter immer `rate()` nehmen!

---

## 🚀 Quick Commands

```bash
# Lab starten
docker compose up -d

# Logs sehen
docker compose logs SERVICE_NAME

# Status
docker compose ps

# Neu starten
docker compose restart SERVICE_NAME

# Alles killen
docker compose down -v
```

---

## 📍 Wichtigste URLs

| Service | URL | Login |
|---------|-----|-------|
| Prometheus | http://localhost:9090 | - |
| Prometheus Targets | http://localhost:9090/targets | - |
| Prometheus Graph | http://localhost:9090/graph | - |
| Grafana | http://localhost:3000 | admin/admin123 |
| Sample App | http://localhost:8888 | - |
| Sample App Metrics | http://localhost:8888/metrics | - |
| AlertManager | http://localhost:9093 | - |

---

## ⚡ Die wichtigsten PromQL-Funktionen

### rate() - Rate pro Sekunde
```promql
rate(app_requests_total[5m])  # Requests/Sekunde
```
**Wann:** Bei Counter, um saubere Daten zu bekommen

### avg() / sum() / max() - Aggregation
```promql
avg(node_cpu_seconds_total)   # Durchschnitt
sum(app_errors_total)         # Summe
max(node_memory_MemTotal_bytes)  # Maximum
```

### by() - Gruppierung
```promql
sum by (job) (up)  # Gruppiere pro Job
```

### Mathe
```promql
# CPU-Auslastung in %
100 - (avg(irate(node_cpu_seconds_total{mode="idle"}[5m])) * 100)

# RAM-Auslastung in %
(1 - (node_memory_MemAvailable_bytes / node_memory_MemTotal_bytes)) * 100

# Error-Rate in %
(sum(rate(app_errors_total[5m])) / sum(rate(app_requests_total[5m]))) * 100
```

---

## 📈 Grafana in 3 Schritten

1. **Dashboard erstellen**
   - Home → Create Dashboard → Create new Dashboard

2. **Panel hinzufügen**
   - Add Panel → Metrics-Query schreiben

3. **Speichern**
   - Click "Apply" → Click "Save"

**Visualisierungen:** Time Series (Graph), Stat, Gauge, Table, Pie, Heatmap

---

## 🚨 Alerts schreiben

**Wo:** `alert-rules.yml`

**Struktur:**
```yaml
- alert: MeinAlert
  expr: metrik > 100           # Bedingung
  for: 2m                      # Dauer
  labels:
    severity: critical         # Wichtigkeit
  annotations:
    summary: "Beschreibung"    # Was ist los?
```

**Severities:** critical > warning > info

---

## 🔍 Debugging-Checkliste

- [ ] Docker läuft? (`docker ps`)
- [ ] Services UP? (`docker compose ps`)
- [ ] Ports erreichbar? (`curl http://localhost:9090`)
- [ ] Logs gelesen? (`docker compose logs`)
- [ ] Warte 30s? (Scrape-Interval!)
- [ ] Firewall blockiert? (`ufw allow PORT`)
- [ ] Disk voll? (`df -h`)
- [ ] RAM genug? (`free -h`)

---

## 💡 Best Practices

✅ **Richtig:**
- PromQL: `rate(counter[5m])` bei Counter nutzen
- Alerts: Mit Severity und Beschreibung
- Labels: Sparsam aber aussagekräftig
- Retention: 7-30 Tage planen
- Dashboards: Max 10-15 Panels pro Board

❌ **Falsch:**
- Counter direkt nutzen ohne rate()
- Alerts ohne Beschreibung
- Zu viele Labels (= Speicher-Explosion)
- Zu lange Retention (= Disk voll)
- 50+ Panels auf einem Dashboard

---

## 📋 Das Minimum zum Anfangen

**Starte mit nur diesen 5 Queries:**

```promql
# 1. Sind alle Services UP?
up

# 2. Wie ist die CPU?
100 - (avg(irate(node_cpu_seconds_total{mode="idle"}[5m])) * 100)

# 3. Wie viel RAM?
(1 - (node_memory_MemAvailable_bytes / node_memory_MemTotal_bytes)) * 100

# 4. Wie viele Requests?
rate(app_requests_total[5m])

# 5. Fehlerquote?
(sum(rate(app_errors_total[5m])) / sum(rate(app_requests_total[5m]))) * 100
```

---

## 🎓 Lernpfad

### Woche 1: Basics
- [ ] Prometheus Architektur verstehen
- [ ] Erste Query schreiben
- [ ] Ein einfaches Dashboard bauen

### Woche 2: Vertiefung
- [ ] PromQL Funktionen lernen (rate, sum, avg)
- [ ] Alerts schreiben und testen
- [ ] Ein Pro-Dashboard bauen

### Woche 3: Mastery
- [ ] Custom Metriken hinzufügen
- [ ] Vollständiges Alerting einrichten
- [ ] Monitoring-Strategie entwickeln

---

## 🔗 Wichtige Links

- **Prometheus Docs:** https://prometheus.io/docs/
- **PromQL Docs:** https://prometheus.io/docs/prometheus/latest/querying/basics/
- **Grafana Docs:** https://grafana.com/docs/
- **Node Exporter:** https://github.com/prometheus/node_exporter
- **Alerting:** https://prometheus.io/docs/alerting/latest/

---

## 🆘 Ich bin stuck!

**1. Symptom googeln:** "prometheus [dein error]"  
**2. Logs lesen:** `docker compose logs SERVICE | grep -i error`  
**3. Test in Prometheus Graph:** http://localhost:9090/graph  
**4. Frag einen erfahreneren Kollegen**  
**5. Schau Handouts:** HANDOUT_06_TROUBLESHOOTING.md

---

## 📝 Notizen für mich

```markdown
# Meine Erkenntnisse

## Was ich gelernt habe:
- ...

## Was ich noch lernen möchte:
- ...

## Meine wichtigsten Queries:
1. ...
2. ...
3. ...

## Meine Probleme & Lösungen:
- Problem: ...
  Lösung: ...
```

---

## ✅ Hast du alles verstanden?

### Test dein Wissen:

**F1:** Was ist der Unterschied zwischen Counter und Gauge?  
<details><summary>A</summary>Counter geht nur hoch, Gauge beliebig</summary></details>

**F2:** Wann nutze ich rate()?  
<details><summary>A</summary>Bei Counter, um Wert pro Sekunde zu bekommen</summary></details>

**F3:** Wie heißt die Abfrage-Sprache von Prometheus?  
<details><summary>A</summary>PromQL (Prometheus Query Language)</summary></details>

**F4:** Auf welchem Port läuft Grafana?  
<details><summary>A</summary>3000</summary></details>

**F5:** Was ist ein Target?  
<details><summary>A</summary>Ein Service/Exporter, von dem Prometheus Metriken holt</summary></details>

**Wenn du alle richtig hattest:** 🎉 Du bist bereit!

---

## 🚀 Nächste Schritte

1. **Heute:** Lab starten und erste Query schreiben
2. **Morgen:** Dashboard bauen
3. **Nächste Woche:** Alert in Grafana einrichten
4. **Ziel:** Eigenes Produktions-Monitoring aufbauen

---

## 📞 Kontakt & Support

**Fragen?**
- Schau zuerst: HANDOUT_06_TROUBLESHOOTING.md
- Dann: Frag Lehrer oder einen erfahreneren Kollegen
- Oder: Schau Prometheus Dokumentation

**Feedback:**
- Was war unklar?
- Was fehlte?
- Was war zu viel?

---

## 🎓 Viel Erfolg beim Lernen!

**"The best way to learn is by doing."**

Baue deine ersten Dashboards, schreibe deine ersten Alerts, und lerne dabei! 💪

---

**Zusammenfassung der Handouts:**

1. **HANDOUT_01_QUICK_START.md** - Los gehts in 5 Minuten
2. **HANDOUT_02_PROMQL_CHEATSHEET.md** - Alle Queries auf einer Seite
3. **HANDOUT_03_GRAFANA_TUTORIAL.md** - Dashboards bauen
4. **HANDOUT_04_GLOSSAR.md** - Alle Begriffe erklärt
5. **HANDOUT_05_UEBUNGEN.md** - 10 praktische Aufgaben
6. **HANDOUT_06_TROUBLESHOOTING.md** - Wenn was kaputt geht
7. **HANDOUT_07_ZUSAMMENFASSUNG.md** - Diese Datei (Die Essenz)

**Viel Spaß beim Vortrag und Lernen!** 🚀

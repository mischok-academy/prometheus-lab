# Prometheus Monitoring - Quick Start Guide

---

## 🚀 Los geht's in 5 Minuten

### 1. Lab starten
```bash
git clone https://github.com/mischok-academy/prometheus-lab.git
cd /home/moritzkraus/src/prometheus-lab
docker compose up -d
```

### 2. Services öffnen
Öffne diese URLs in deinem Browser:

| Service | URL | Login |
|---------|-----|-------|
| **Prometheus** | http://localhost:9090 | keine |
| **Grafana** | http://localhost:3000 | admin / admin123 |
| **Sample App** | http://localhost:8888 | keine |

### 3. Erste Query schreiben
1. Gehe zu Prometheus: http://localhost:9090
2. Klick auf "Graph" Tab
3. Schreibe in das Eingabefeld: `up`
4. Klick "Execute"
5. Du siehst alle Targets und ihren Status

---

## 📊 Wichtigste Konzepte (5 min Überblick)

### Was ist eine Metrik?
Eine **Metrik** = eine Messgröße mit Zeitstempel

```
CPU-Auslastung um 14:00 Uhr = 45%
RAM-Verbrauch um 14:01 Uhr = 8 GB
HTTP-Requests pro Sekunde um 14:02 Uhr = 1250
```

### 4 Metriken-Typen

| Typ | Beispiel | Merkhilfe |
|-----|----------|-----------|
| **Counter** | Gesamtrequests: 1.000.000 | 📈 Nur nach oben |
| **Gauge** | CPU jetzt: 45% | 📊 Beliebig hoch/runter |
| **Histogram** | Response-Zeit: 0.1-1.5s | 📉 Verteilung |
| **Summary** | P95 Response: 500ms | 📐 Percentile |

### Wie es funktioniert

```
1. Exporter         2. Prometheus      3. Grafana/PromQL
   sammelt    →        speichert    →     zeigt an
   Metriken          Zeitreihen          & befragt
```

---

## 🎯 3 wichtigste Commands

### Status überprüfen
```bash
docker compose ps
```
✅ Alle sollten "Up" sein

### Logs anschauen
```bash
docker compose logs -f prometheus
docker compose logs -f sample-app
```

### Alles neustarten
```bash
docker compose restart
```

---

## 💡 Erste Übung (10 Min)

### Task: Ein Dashboard erstellen

1. **Grafana öffnen:** http://localhost:3000
2. **Einloggen:** admin / admin123
3. **Dashboard erstellen:**
   - Home → Create Dashboard → Create new Dashboard
4. **Panel hinzufügen:**
   - Click "Add Panel"
5. **Query schreiben:**
   - Under "Metrics", schreibe: `up`
6. **Visualisieren:**
   - Wähle "Stat" (oben rechts)
7. **Speichern:**
   - Oben rechts "Save"

**Ergebnis:** Du hast dein erstes Dashboard erstellt! 🎉

---

## 📝 Checkliste vor/nach dem Vortrag

- [ ] Docker installiert?
- [ ] Lab startet ohne Fehler?
- [ ] Alle 7 Services laufen?
- [ ] Kann ich Prometheus öffnen?
- [ ] Kann ich Grafana öffnen (admin/admin123)?
- [ ] Sample App liefert Metriken?

---

## 🆘 Schnelle Fehlerbehebung

### Problem: Container crasht
```bash
docker compose logs CONTAINER_NAME
# CONTAINER_NAME durch echten Namen ersetzen
# z.B.: docker compose logs prometheus
```

### Problem: Port ist belegt
```bash
# Prüfe welcher Process Port 9090 nutzt
lsof -i :9090
# Oder: stoppe alles und starte neu
docker compose down
docker compose up -d
```

### Problem: Keine Metriken sichtbar
1. Prüfe ob alle Targets "UP" sind
2. Warte 30 Sekunden (Prometheus braucht Zeit zum Scrapen)
3. Neuladen (Ctrl+F5)

---

## 📚 Was du noch lernen wirst

✅ PromQL - Abfrage-Sprache  
✅ Dashboards bauen  
✅ Alerts schreiben  
✅ Best Practices  
✅ Praktische Troubleshooting  

---

**Nächste Datei:** HANDOUT_02_PROMQL_CHEATSHEET.md


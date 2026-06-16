# Prometheus Monitoring Lab

A complete Docker Compose environment for testing and learning Prometheus monitoring, including Grafana, AlertManager, and various exporters.

## Architecture

```
┌─────────────────┐
│   Prometheus    │  Metrics collection & storage
│   (port 9090)   │
└────────┬────────┘
         │
    ┌────┴────────────┬──────────────────┬──────────────┐
    │                 │                  │              │
┌───▼────┐     ┌─────▼─────┐    ┌──────▼──┐    ┌─────▼──────┐
│ Grafana │     │   Alert   │    │ Sample  │    │  Node      │
│(3000)  │     │  Manager  │    │  App    │    │ Exporter   │
│         │     │ (9093)    │    │ (8888)  │    │ (9100)     │
└─────────┘     └───────────┘    └─────────┘    └────────────┘

     ┌───────────────┐         ┌────────────┐
     │   cAdvisor    │         │ Pushgateway│
     │   (8080)      │         │   (9091)   │
     └───────────────┘         └────────────┘
```

## Services

| Service | Port | Description |
|---------|------|-------------|
| **Prometheus** | 9090 | Time-series database and monitoring system |
| **Grafana** | 3000 | Visualization and dashboarding (admin/admin123) |
| **Node Exporter** | 9100 | Host system metrics (CPU, memory, disk, etc.) |
| **cAdvisor** | 8080 | Container metrics |
| **AlertManager** | 9093 | Alert routing and management |
| **Pushgateway** | 9091 | For scraping short-lived batch jobs |
| **Sample App** | 8888 | Example Go app with custom metrics |

## Quick Start

### 1. Start the Lab

```bash
docker-compose up -d
```

Wait for all containers to be healthy:

```bash
docker-compose ps
```

### 2. Access the Services

- **Prometheus**: http://localhost:9090
  - View targets: http://localhost:9090/targets
  - Query metrics: http://localhost:9090/graph

- **Grafana**: http://localhost:3000
  - Default login: `admin` / `admin123`
  - Prometheus datasource is auto-configured

- **Sample App**: http://localhost:8888
  - Metrics endpoint: http://localhost:8888/metrics

- **AlertManager**: http://localhost:9093

## Common PromQL Queries

### Node Metrics
```promql
# CPU usage percentage
100 - (avg by (instance) (irate(node_cpu_seconds_total{mode="idle"}[5m])) * 100)

# Memory usage percentage
(1 - (node_memory_MemAvailable_bytes / node_memory_MemTotal_bytes)) * 100

# Disk usage percentage
(1 - (node_filesystem_avail_bytes{fstype!~"tmpfs"} / node_filesystem_size_bytes{fstype!~"tmpfs"})) * 100

# Network traffic
rate(node_network_receive_bytes_total[5m])
```

### Container Metrics
```promql
# Container CPU usage
rate(container_cpu_usage_seconds_total{name!=""}[5m])

# Container memory usage
container_memory_usage_bytes{name!=""}

# Container restart count
container_last_seen
```

### Application Metrics
```promql
# Request rate
rate(app_requests_total[5m])

# Error rate
rate(app_errors_total[5m])

# Request duration p95
histogram_quantile(0.95, rate(app_request_duration_seconds_bucket[5m]))

# Processed items
app_processed_items
```

## Configuration Files

### `prometheus.yml`
Configures Prometheus scrape targets and alerting rules. Add new targets in the `scrape_configs` section:

```yaml
- job_name: 'my-service'
  static_configs:
    - targets: ['my-service:9090']
  scrape_interval: 15s
```

### `alert-rules.yml`
Defines alert rules for different components. Edit to add your own alerts.

### `alertmanager.yml`
Configures alert routing and notification channels. Supports:
- Slack webhooks
- Email notifications
- PagerDuty
- Custom webhooks

To enable Slack notifications:
```yaml
global:
  slack_api_url: 'https://hooks.slack.com/services/YOUR/WEBHOOK/URL'
```

## Common Tasks

### Add a New Target

1. Update `prometheus.yml`:
```yaml
- job_name: 'my-app'
  static_configs:
    - targets: ['my-app:9090']
```

2. Reload Prometheus:
```bash
curl -X POST http://localhost:9090/-/reload
```

### Create a Custom Dashboard

1. Open Grafana: http://localhost:3000
2. Click "+" → "Dashboard"
3. Add panels querying Prometheus metrics
4. Save the dashboard

### Test Alerts

Trigger a high CPU alert:
```bash
# From any container
docker-compose exec prometheus wget -O- http://node-exporter:9100/metrics | grep node_cpu
```

Or manually create an alert condition:
```bash
# Use PromQL to check alert status
curl 'http://localhost:9090/api/v1/query?query=ALERTS'
```

### Push Metrics to Pushgateway

```bash
# Example: Push a batch job metric
curl -X POST --data-binary @- http://localhost:9091/metrics/job/myjob <<EOF
# HELP my_metric A custom metric
# TYPE my_metric gauge
my_metric 42
EOF
```

### View Prometheus Targets

```bash
curl http://localhost:9090/api/v1/targets | jq .
```

### Check AlertManager Alerts

```bash
curl http://localhost:9093/api/v1/alerts | jq .
```

## Development

### Modify Sample App

Edit `sample-app/main.go` and restart:

```bash
docker-compose restart sample-app
```

### Add Custom Metrics to Sample App

```go
var myMetric = prometheus.NewGaugeVec(
    prometheus.GaugeOpts{
        Name: "my_custom_metric",
        Help: "My custom metric",
    },
    []string{"label"},
)

func init() {
    prometheus.MustRegister(myMetric)
}
```

### Enable Slack Notifications

1. Create a Slack webhook: https://api.slack.com/apps
2. Update `alertmanager.yml` with your webhook URL
3. Restart AlertManager:

```bash
docker-compose restart alertmanager
```

## Cleanup

### Stop All Services
```bash
docker-compose down
```

### Stop and Remove Volumes
```bash
docker-compose down -v
```

### View Logs

```bash
# All services
docker-compose logs -f

# Specific service
docker-compose logs -f prometheus
docker-compose logs -f grafana
docker-compose logs -f alertmanager
```

## Troubleshooting

### Prometheus targets showing as "DOWN"

```bash
# Check target health
curl http://localhost:9090/api/v1/targets

# Verify service is running
docker-compose ps

# Check service logs
docker-compose logs prometheus
docker-compose logs node-exporter
```

### Grafana can't connect to Prometheus

```bash
# Verify Prometheus is healthy
curl http://localhost:9090/-/healthy

# Check Grafana logs
docker-compose logs grafana
```

### AlertManager not sending notifications

```bash
# Check AlertManager logs
docker-compose logs alertmanager

# Verify configuration
docker-compose exec alertmanager amtool config routes
```

### Sample app metrics not appearing

```bash
# Check if service is running
docker-compose logs sample-app

# Verify metrics endpoint
curl http://localhost:8888/metrics
```

## Learning Resources

- [Prometheus Documentation](https://prometheus.io/docs/)
- [PromQL Query Language](https://prometheus.io/docs/prometheus/latest/querying/basics/)
- [AlertManager Configuration](https://prometheus.io/docs/alerting/latest/configuration/)
- [Grafana Dashboards](https://grafana.com/grafana/dashboards/)
- [Node Exporter Metrics](https://github.com/prometheus/node_exporter)

## Notes

- Data is persisted in Docker volumes: `prometheus_data`, `grafana_data`, `alertmanager_data`
- All services use a shared `monitoring` network
- Prometheus retention is set to 30 days
- Health checks are configured for critical services
- Services auto-restart on failure

## License

This lab environment is provided as-is for educational purposes.

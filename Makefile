.PHONY: help up down logs restart clean health test query prometheus grafana alertmanager

help:
	@echo "Prometheus Lab - Docker Compose Commands"
	@echo ""
	@echo "Usage:"
	@echo "  make up              Start all services"
	@echo "  make down            Stop all services"
	@echo "  make logs            View logs from all services"
	@echo "  make logs-SRVICE     View logs from specific service (e.g., make logs-prometheus)"
	@echo "  make restart         Restart all services"
	@echo "  make restart-SERVICE Restart specific service (e.g., make restart-prometheus)"
	@echo "  make clean           Stop and remove volumes"
	@echo "  make health          Check service health status"
	@echo "  make test            Run test queries"
	@echo "  make prometheus      Open Prometheus in browser"
	@echo "  make grafana         Open Grafana in browser"
	@echo "  make alertmanager    Open AlertManager in browser"
	@echo "  make query QUERY=... Execute PromQL query (e.g., make query QUERY='up')"

up:
	docker-compose up -d
	@echo "Waiting for services to start..."
	@sleep 5
	@docker-compose ps

down:
	docker-compose down

logs:
	docker-compose logs -f

logs-%:
	docker-compose logs -f $*

restart:
	docker-compose restart

restart-%:
	docker-compose restart $*

clean:
	docker-compose down -v
	@echo "Cleaned up all services and volumes"

health:
	@echo "Service Health Status:"
	@docker-compose ps
	@echo ""
	@echo "Prometheus health:"
	@curl -s http://localhost:9090/-/healthy || echo "Prometheus not responding"
	@echo ""
	@echo "Grafana health:"
	@curl -s http://localhost:3000/api/health || echo "Grafana not responding"

test:
	@echo "Testing Prometheus connectivity..."
	@curl -s http://localhost:9090/api/v1/query?query=up | jq '.data.result | length' | xargs echo "Targets up:"
	@echo ""
	@echo "Testing sample app..."
	@curl -s http://localhost:8888/health && echo "Sample app is healthy" || echo "Sample app not responding"
	@echo ""
	@echo "Testing AlertManager..."
	@curl -s http://localhost:9093/api/v1/alerts | jq '. | length' | xargs echo "Active alerts:"

prometheus:
	@command -v xdg-open >/dev/null 2>&1 && xdg-open http://localhost:9090 || \
	command -v open >/dev/null 2>&1 && open http://localhost:9090 || \
	echo "Please open http://localhost:9090 in your browser"

grafana:
	@command -v xdg-open >/dev/null 2>&1 && xdg-open http://localhost:3000 || \
	command -v open >/dev/null 2>&1 && open http://localhost:3000 || \
	echo "Please open http://localhost:3000 in your browser"

alertmanager:
	@command -v xdg-open >/dev/null 2>&1 && xdg-open http://localhost:9093 || \
	command -v open >/dev/null 2>&1 && open http://localhost:9093 || \
	echo "Please open http://localhost:9093 in your browser"

query:
	@if [ -z "$(QUERY)" ]; then \
		echo "Usage: make query QUERY='your_promql_query'"; \
		echo "Example: make query QUERY='up'"; \
	else \
		curl -s 'http://localhost:9090/api/v1/query?query=$(QUERY)' | jq '.'; \
	fi

reload-prometheus:
	@curl -X POST http://localhost:9090/-/reload && echo "Prometheus configuration reloaded"

push-metric:
	@if [ -z "$(JOB)" ] || [ -z "$(VALUE)" ]; then \
		echo "Usage: make push-metric JOB=job_name VALUE=value"; \
		echo "Example: make push-metric JOB=test_job VALUE=123"; \
	else \
		curl -X POST --data-binary "@-" http://localhost:9091/metrics/job/$(JOB) <<< "test_metric $(VALUE)"; \
	fi

targets:
	@curl -s http://localhost:9090/api/v1/targets | jq '.data.activeTargets[] | {job: .labels.job, instance: .labels.instance, health: .health}'

alerts:
	@curl -s http://localhost:9093/api/v1/alerts | jq '.data[] | {alertname: .labels.alertname, severity: .labels.severity, state: .state}'

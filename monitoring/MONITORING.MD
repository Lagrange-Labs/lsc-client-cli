# Overview

This document provides a sample template for monitoring an Optimism attestation node using Prometheus and Grafana. The template assumes the attestation node is running on its own machine, and Prometheus metrics are exposed on port `8080`. Operators running multiple attestation nodes on the same machine can modify this template as needed.

## Prometheus Configuration

The default configuration can be found in `prometheus.yml` file.

**The scrape target is set to `localhost:8080`. Modify this to match the IP address of the machine and port where the metrics are exposed.**

## Alert Configuration

The alert is set to trigger if the `gap` field, which is computed by taking the difference between `client_current_batch_number` and `client_commit_batch_number`, exceeds `5`.

> [!NOTE]
> If you are running multiple attestation nodes on the same machine, update the prometheus.yml file to include all relevant targets.

## Setup Instructions

1. Navigate to the Monitoring Directory

```bash
cd monitoring
```

2. Start the Docker Containers for Prometheus and Grafana

```bash
docker compose up -d
```

3. To Shut Down the Docker Containers

```bash
docker compose down
```

## Conclusion

This template provides a starting point for monitoring Optimism attestation nodes with Prometheus and Grafana. Customize it according to your deployment setup to ensure effective monitoring and alerting.

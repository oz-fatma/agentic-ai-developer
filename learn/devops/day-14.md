# Day 14: Docker Compose — Multi-Container Applications

**Phase:** Containerization with Docker (Days 11-15)

## Phase Context

Real apps need databases, caches, and APIs together. Docker Compose defines and runs multi-container stacks from a single YAML file on your local machine.

## Tasks

### 1. Write docker-compose.yml

```yaml
# docker-compose.yml
services:
  web:
    image: nginx:alpine
    ports:
      - "8080:80"
    depends_on:
      - api
  api:
    image: python:3.12-slim
    command: python -m http.server 9000
    expose:
      - "9000"
```

### 2. Start and scale the stack

```bash
docker compose up -d
docker compose ps
docker compose logs api
docker compose up -d --scale api=2
docker compose down
```

### 3. Override with environment file

```bash
cat > .env <<EOF
COMPOSE_PROJECT_NAME=devops-lab
EOF
docker compose config
```

## Notes / Cheat Sheet

| Item | Description |
|---|---|
| `docker compose up -d` | Start stack in background |
| `docker compose down` | Stop and remove containers |
| `depends_on` | Startup order (not health) |
| `expose` vs `ports` | Internal vs published ports |
| `docker compose config` | Validate merged YAML |

## Summary

Compose orchestrates local multi-service environments. Use it to mirror production topology before deploying to the cloud.

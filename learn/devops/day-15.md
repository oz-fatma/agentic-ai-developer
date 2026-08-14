# Day 15: Docker Networking and Volumes — Persistence and Connectivity

**Phase:** Containerization with Docker (Days 11-15)

## Phase Context

Containers are ephemeral; volumes persist data. Networks isolate and connect services. Today you practice both for stateful, multi-service applications.

## Tasks

### 1. Create and use a named volume

```bash
docker volume create pgdata
docker run -d --name db \
  -e POSTGRES_PASSWORD=secret \
  -v pgdata:/var/lib/postgresql/data \
  postgres:16-alpine
docker volume inspect pgdata
```

### 2. Custom bridge network

```bash
docker network create app-net
docker run -d --name api --network app-net nginx:alpine
docker run -d --name proxy --network app-net nginx:alpine
docker network inspect app-net
```

### 3. Bind mount for development

```bash
mkdir html && echo "<h1>DevOps</h1>" > html/index.html
docker run -d -p 8080:80 -v $(pwd)/html:/usr/share/nginx/html nginx:alpine
curl localhost:8080
```

## Notes / Cheat Sheet

| Item | Description |
|---|---|
| Named volume | Managed storage, survives container delete |
| Bind mount | Maps host path into container |
| Bridge network | Default isolated network |
| `-v name:path` | Attach volume to container |
| `--network` | Connect container to network |

## Summary

Volumes keep data; networks connect services. Use named volumes for production data and bind mounts for live code editing during development.

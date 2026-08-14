# Day 12: Images and Containers — Lifecycle and Inspection

**Phase:** Containerization with Docker (Days 11-15)

## Phase Context

Images are read-only templates; containers are running instances. Today you pull, tag, inspect, and understand the image/container lifecycle that underpins all container workflows.

## Tasks

### 1. Pull, tag, and list images

```bash
docker pull redis:7-alpine
docker tag redis:7-alpine myredis:local
docker images | grep redis
```

### 2. Inspect running containers

```bash
docker run -d --name cache redis:7-alpine
docker inspect cache --format '{{.State.Status}}'
docker exec cache redis-cli ping
docker stats cache --no-stream
```

### 3. Clean up unused resources

```bash
docker stop cache && docker rm cache
docker image prune -f
docker system df
```

## Notes / Cheat Sheet

| Item | Description |
|---|---|
| `docker pull` | Download image from registry |
| `docker tag` | Create alias for an image |
| `docker inspect` | JSON metadata about object |
| `docker exec` | Run command inside running container |
| `docker prune` | Remove unused data |

## Summary

Images are layered templates; containers add a writable layer on top. Use inspect, exec, and stats to debug running workloads.

# Day 11: Docker Introduction — Install and First Containers

**Phase:** Containerization with Docker (Days 11-15)

## Phase Context

Docker packages apps with dependencies into portable containers. Today you install Docker, run your first containers, and learn the core CLI workflow that every DevOps engineer uses daily.

## Tasks

### 1. Install and verify Docker

```bash
docker --version
docker info | head -20
docker run hello-world
```

### 2. Run and manage containers

```bash
docker run -d --name web -p 8080:80 nginx:alpine
docker ps
docker logs web
curl -s localhost:8080 | head -5
docker stop web && docker rm web
```

### 3. Explore images locally

```bash
docker images
docker pull alpine:3.19
docker run -it --rm alpine:3.19 sh -c "cat /etc/os-release"
```

## Notes / Cheat Sheet

| Item | Description |
|---|---|
| `docker run` | Create and start a container |
| `-d` | Detached (background) mode |
| `-p host:container` | Publish port |
| `docker ps -a` | List all containers |
| `docker rm` | Remove stopped container |

## Summary

You verified Docker, ran nginx in a container, mapped ports, and cleaned up. Containers are ephemeral — stop and remove them when labs are done.

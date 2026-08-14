# Day 13: Dockerfiles — Build Custom Images

**Phase:** Containerization with Docker (Days 11-15)

## Phase Context

Dockerfiles define how to build images reproducibly. Today you write a Dockerfile, understand layer caching, and build a custom image from source.

## Tasks

### 1. Create a simple Dockerfile

```dockerfile
# Dockerfile
FROM python:3.12-slim
WORKDIR /app
COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt
COPY . .
EXPOSE 8000
CMD ["python", "-m", "http.server", "8000"]
```

### 2. Build and run your image

```bash
echo "flask==3.0.0" > requirements.txt
echo "# demo app" > app.py
docker build -t myapp:1.0 .
docker run -d -p 8000:8000 --name myapp myapp:1.0
curl -I localhost:8000
```

### 3. Use .dockerignore and rebuild

```bash
echo "__pycache__" > .dockerignore
docker build -t myapp:1.1 .
docker history myapp:1.1
```

## Notes / Cheat Sheet

| Item | Description |
|---|---|
| `FROM` | Base image |
| `COPY` / `ADD` | Add files to image |
| `RUN` | Execute command at build time |
| `CMD` / `ENTRYPOINT` | Default runtime command |
| `docker build -t` | Build and tag image |

## Summary

Dockerfiles turn source code into portable images. Order layers for cache hits: install dependencies before copying application code.

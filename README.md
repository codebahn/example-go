# Docker container

A minimal Go HTTP service with a multi-stage Docker build. Demonstrates container builds and smoke testing on [Forgejo Actions](https://forgejo.org/docs/latest/user/actions/).

## CI workflow

Runs on every push ([`.forgejo/workflows/ci.yml`](.forgejo/workflows/ci.yml)):

1. Builds the Docker image
2. Starts the container and verifies `/health` responds

Uncomment the push steps in `ci.yml` to ship images to a container registry.

## Run locally

```
go run main.go
# or
docker build -t myapp . && docker run -p 8080:8080 myapp
```

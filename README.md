# Go

A minimal Go HTTP service with tests and a multi-stage Docker build. Demonstrates `go vet`, `go test`, container builds, and smoke testing on [Forgejo Actions](https://forgejo.org/docs/latest/user/actions/).

## CI workflow

Runs on every push ([`.forgejo/workflows/ci.yml`](.forgejo/workflows/ci.yml)):

1. Runs `go vet` and `go test`
2. Builds the Docker image
3. Starts the container and verifies `/health` responds

Uncomment the push steps in `ci.yml` to ship images to a container registry.

## Run locally

```
go test ./...
go run main.go
# or
docker build -t myapp . && docker run -p 8080:8080 myapp
```

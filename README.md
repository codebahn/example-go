# Go

A minimal Go HTTP service with tests and a multi-stage Docker build.

## CI workflow

The workflow lives at [`.github/workflows/ci.yml`](.github/workflows/ci.yml) and declares `runs-on: ubuntu-latest`. Both are GitHub defaults, unchanged. Nothing in this repository was renamed or rewritten to run here:

- `.github/workflows/` is read directly. There is no `.forgejo/` directory.
- `ubuntu-latest` routes to a Codebahn runner in Paris.
- `actions/checkout` and `actions/setup-go` resolve from GitHub automatically.

On every push and pull request:

1. Runs `go vet` and `go test`
2. Builds the Docker image
3. Starts the container and verifies `/health` responds

Uncomment the push steps in `ci.yml` to ship images to the built-in container registry. The job token authenticates automatically, so there are no secrets to configure.

Need more than the 3 GB default? Set `runs-on: codebahn-medium` for 7 GB. See [CI runners](https://codebahn.net/docs/features/ci-runners/).

## Run locally

```
go test ./...
go run main.go
# or
docker build -t myapp . && docker run -p 8080:8080 myapp
```

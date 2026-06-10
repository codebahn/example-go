# Docker container example

A minimal Go HTTP service with a multi-stage Docker build, built and smoke-tested on every push using Codebahn's managed runners. The workflow builds the image, tags it with the commit SHA, and verifies the health endpoint responds before the job passes. Uncomment the push step and add registry secrets to ship images to your own registry.

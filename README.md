# miniKMS

_A lightweight, developer‑friendly Key Management Service for local development, CI, and integration testing._

[![Go Reference](https://pkg.go.dev/badge/github.com/michaeljmartin28/minikms/package/sdk/go.svg)](https://pkg.go.dev/github.com/michaeljmartin28/minikms/package/sdk/go)
[![npm version](https://img.shields.io/npm/v/@minikms/sdk.svg)](https://www.npmjs.com/package/@minikms/sdk)
[![PyPI version](https://img.shields.io/pypi/v/minikms.svg)](https://pypi.org/project/minikms/)

![Docker Pulls](https://img.shields.io/docker/pulls/michaeljmartin28/minikms)
![Docker Image Size](https://img.shields.io/docker/image-size/michaeljmartin28/minikms/0.2.0)

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

miniKMS is a minimal, pluggable, cloud‑shaped KMS emulator written in Go.
It provides a consistent API for key lifecycle operations, encryption, decryption, and data key generation without requiring AWS, GCP, Azure, or Vault.
It’s designed for developers who want a KMS experience locally, with zero cloud dependencies.


## Features

### Core
- AES‑GCM encryption engine
- Versioned key storage (BoltDB backend)
- Envelope encryption (GenerateDataKey / DecryptDataKey)
- Key lifecycle operations (create, rotate, enable, disable)
- Deterministic, test‑friendly behavior
- Unified error taxonomy

### Transports
- Full HTTP/JSON API
- Full gRPC API (protobuf definitions included)

### Developer Experience
- Docker‑friendly runtime
- Test store for isolated engine tests
- Clean, transport‑agnostic core engine

### Provider Modes (planned)
- Generic mode (simple, cloud‑agnostic API)
- AWS‑like KMS mode
- GCP KMS mode
- Azure Key Vault crypto mode
- Vault Transit mode

### SDKs

**Go:**  
[![Go SDK](https://img.shields.io/badge/pkg.go.dev-reference-00ADD8?logo=go)](https://pkg.go.dev/github.com/michaeljmartin28/minikms/package/sdk/go)  
[Source](package/sdk/go)

**Node:**  
[![npm version](https://img.shields.io/npm/v/@minikms/sdk.svg?logo=npm)](https://www.npmjs.com/package/@minikms/sdk)  
[Source](package/sdk/node)

**Python:**   
[![PyPI](https://img.shields.io/pypi/v/minikms.svg?logo=python&logoColor=white)](https://pypi.org/project/minikms/)  
[Source](package/sdk/python)

### Docker
![Docker Pulls](https://img.shields.io/docker/pulls/michaeljmartin28/minikms)
![Docker Image Size](https://img.shields.io/docker/image-size/michaeljmartin28/minikms/0.2.0)

miniKMS is now available as a lightweight Docker image (~9.5MB):

#### Run the server
```bash
docker run -p 8080:8080 michaeljmartin28/minikms:0.2.0
```

Or always pull the latest stable version:
```bash
docker run -p 8080:8080 michaeljmartin28/minikms:latest
```
Build from source
```bash
docker build \
  --build-arg VERSION=0.2.0 \
  -t minikms:0.2.0 .
```

## Why miniKMS?

Cloud KMS systems are powerful but difficult to run locally.
miniKMS gives you a tiny, reproducible, dependency‑free KMS that behaves like the real thing — perfect for:
- Local development
- CI pipelines
- Integration tests
- Security demos
- Cloud migration testing
- Learning KMS concepts

If you’ve ever mocked AWS KMS with hand‑rolled stubs or struggled to test envelope encryption locally, miniKMS is built for you.

## Project Status

miniKMS is under active development.

The following components are complete:
- Core engine
- Storage layer
- Crypto provider
- HTTP transport
- gRPC transport
- Test infrastructure
- Node SDK (npm)
- Go SDK
- Python SDK (pypi)
- Public Docker image

Currently in progress:
- Demo page

Planned next:
- AWS‑like provider
- Web admin portal
- Additional cloud provider modes
- Adding health/version endpoints
- Adding user-specific support

## Documentation

For detailed architecture and design decisions, see:

- `docs/design.md` — the authoritative reference for engine, storage, crypto, transports, and future extensions.

Roadmap and planning are tracked in the GitHub Project board.




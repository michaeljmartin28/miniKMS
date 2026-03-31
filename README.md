# miniKMS

_A lightweight, developer‑friendly Key Management Service for local development, CI, and integration testing._

[![Go Reference](https://pkg.go.dev/badge/github.com/michaeljmartin28/minikms/package/sdk/go.svg)](https://pkg.go.dev/github.com/michaeljmartin28/minikms/package/sdk/go)
[![npm version](https://img.shields.io/npm/v/@minikms/sdk.svg)](https://www.npmjs.com/package/@minikms/sdk)
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

Currently in progress:
- Go SDK
- Node + Python SDKs
- Docker image
- Demo page

Planned next:
- AWS‑like provider
- Web admin portal
- Additional cloud provider modes


## Documentation

For detailed architecture and design decisions, see:

- `docs/design.md` — the authoritative reference for engine, storage, crypto, transports, and future extensions.

Roadmap and planning are tracked in the GitHub Project board.




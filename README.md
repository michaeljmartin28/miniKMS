# miniKMS

_A lightweight, developer‑friendly Key Management Service for local development and CI._

miniKMS is a minimal, pluggable, cloud‑shaped KMS emulator written in Go.

It provides a simple, consistent way to test encryption, key lifecycle flows, and cloud‑style KMS integrations without running AWS, GCP, Azure, or Vault.

## Features (planned)

- Lightweight AES‑GCM encryption engine
- File‑based key storage
- REST API for key creation, encryption, and decryption
- Envelope encryption (data keys)
- Audit logging
- Provider modes:
- Generic mode (clean, simple API)
- AWS KMS mode
- GCP KMS mode
- Azure Key Vault crypto mode
- Vault Transit mode
- Docker‑first developer experience
- Go client SDK (later JS/Python)

## Why miniKMS?

Cloud KMS systems are powerful but heavy to mock locally.
miniKMS gives developers a tiny, local, reproducible KMS that behaves like the real thing perfect for:

- Local development
- CI pipelines
- Integration testing
- Security demos
- Cloud migration testing
- Learning KMS

## Project Status

- Early planning and architecture design.
- Implementation begins after v0.1 design is finalized.

## Documentation

- docs/vision.md — project goals and philosophy
- docs/architecture.md — core engine, provider interface, storage model
- docs/roadmap.md — release plan and milestones

# miniKMS Architecture & Design Document

## 1. Overview

miniKMS is a lightweight, developer‑friendly Key Management Service designed for local development, CI pipelines, and integration testing. It provides a cloud‑shaped API surface (REST + gRPC) with pluggable provider modes that emulate AWS KMS, GCP KMS, Azure Key Vault Crypto, and Vault Transit.
The goal is not to replace real cloud KMS systems, but to offer a deterministic, fast, and portable alternative for development environments.

### Goals

- Provide a minimal, consistent KMS API for local use
- Support both REST and gRPC transports
- Offer provider modes that mimic cloud KMS semantics
- Use a simple, embedded storage backend (bbolt)
- Keep the core engine cloud‑agnostic
- Make the system easy to run via Docker
- Provide SDKs for Go, JS/TS, and Python

### Non‑Goals

- No real IAM or policy enforcement
- No HSM integration
- No multi‑region replication
- No multi‑tenant isolation
- No production‑grade durability guarantees
- No signing keys in v0.x (future consideration)

## 2. High‑Level Architecture

```
                +----------------------+
                |      REST API        |
                |   (JSON over HTTP)   |
                +----------+-----------+
                           |
                           v
+-----------------+   +-----------+   +----------------+
| Provider Layer  |-->| CoreKMS   |-->| Storage Layer  |
| (AWS/GCP/Azure) |   | Engine    |   |   (bbolt)      |
+-----------------+   +-----------+   +----------------+
                           ^
                           |
                +----------+-----------+
                |       gRPC API       |
                |   (Protocol Buffers) |
                +----------------------+

```

### Key Principles

- **CoreKMS is the center of the system**  
  All operations flow through it, regardless of transport or provider mode.

- **Providers translate, not implement**  
  They adapt cloud‑specific semantics to the core engine.

- **Storage is abstracted**  
  bbolt is the default backend, but the interface allows future backends.

- **REST and gRPC share the same core types**  
  Prevents divergence between transports.

## 3. Core Interfaces

### CoreKMS

The main engine interface:

- `Encrypt`
- `Decrypt`
- `GenerateKey`
- `RotateKey`
- `DisableKey`
- `EnableKey`
- `GenerateDataKey`
- `DecryptDataKey`

  _This interface is stable and transport‑agnostic._

### KeyStore

Abstracts storage:

- `CreateKey`
- `GetKey`
- `ListKeys`
- `SaveKeyVersion`

  _bbolt implements this interface._

### Provider

Defines cloud‑specific behavior:

- request translation
- response translation
- error mapping
- route mounting

_Providers never bypass the core engine._

## 4. Data Model

### Key Metadata

- id (UUID)
- arn or provider‑specific identifier
- state (Enabled, Disabled)
- created_at
- versions (list of version IDs)

### Key Version

- version_number
- material (AES‑256 key)
- created_at

### Ciphertext Format

A simple, deterministic structure:  
`version` | `nonce` | `ciphertext` | `tag`

_Encoded as base64 for REST._

## 5. Storage Model (bbolt)

### Buckets

```
keys/ -> key metadata (JSON)
key_versions/ -> versioned key material (binary)
```

### Key Layout

- `keys/<key_id>` -> metadata JSON
- `key_versions/<key_id>/<version>` -> raw key bytes

### Rationale

- Simple
- Fast
- Single‑file persistence
- Easy to mount in Docker

## 6. API Model

### REST Endpoints

- `POST /keys`
- `GET /keys`
- `POST /keys/{id}/rotate`
- `POST /keys/{id}/disable`
- `POST /keys/{id}/enable`
- `POST /encrypt`
- `POST /decrypt`
- `POST /datakey`
- `POST /datakey/decrypt`

### gRPC Service

- `Encrypt`
- `Decrypt`
- `GenerateKey`
- `RotateKey`
- `GenerateDataKey`
- `DecryptDataKey`

### Error Model

#### REST:

```json
{
  "error": "KeyDisabled",
  "message": "Key is disabled"
}
```

#### gRPC:

- Use canonical gRPC status codes
- Map provider errors appropriately

## 7. Provider Mode Strategy

### Provider Responsibilities

- Define cloud‑specific request/response shapes
- Translate to/from core types
- Mount provider‑specific routes
- Generate provider‑specific identifiers (ARNs, resource names)
  Provider Modes
- AWS KMS
- GCP KMS
- Azure Key Vault Crypto
- Vault Transit (future)

#### Selection

CLI flag or config:

- `--provider=aws`
- `--provider=gcp`
- `--provider=azure`

## 8. Configuration Model

### CLI Flags

- `--rest-port`
- `--grpc-port`
- `--db-path`
- `--provider`
- `--log-level`
  ### Environment Variables
- `MINIKMS_REST_PORT`
- `MINIKMS_GRPC_PORT`
- `MINIKMS_DB_PATH`
  ### Defaults
- `REST: 8080`
- `gRPC: 9090`
- `DB: /data/minikms.db`

## 9. Logging & Audit

### Audit Log Fields

- timestamp
- operation
- key_id
- version
- success/failure
- request_id

### Format

JSON, one entry per line.

## 10. Testing Strategy

### Unit Tests

- CoreKMS crypto operations
- Storage layer
- Provider translation

### Integration Tests

- REST endpoints
- gRPC endpoints
- Envelope encryption flows

### Golden Tests

- AWS/GCP/Azure provider responses
- Error shapes

### Test Utilities

- in‑memory KeyStore
- deterministic nonce generator (test‑only)

## 11. Release Strategy

### Versioning

#### Semantic Versioning:

- v0.1.0 → Core engine
- v0.2.0 → Key lifecycle
- v0.3.0 → Envelope encryption
- v0.4.0 → AWS provider
- v0.5.0 → GCP/Azure providers
- v1.0.0 → SDKs + plugin interface

### Changelog

- Maintain [Unreleased] section
- Promote to version block on release

### Tagging

- git tag v0.1.0
- git push --tags

# Architecture

This document describes the initial architecture for the tool. v0.1

## Core Components

- Key store (file‑based to start)
- Crypto engine (AES‑GCM)
- Key lifecycle (create, rotate, disable)
- Envelope encryption
- Audit logging

### API Layer

miniKMS exposes two transport layers that map to the same core request/response types and call the same engine.

- REST API (JSON over HTTP)
- gRPC API (Protocol Buffers)

### Providers

- Provider Interface
- Providers adapt cloud‑specific semantics to the core engine.
- Providers
  - AWSProvider
  - GCPProvider
  - AzureProvider
  - VaultProvider

Each provider:

- Defines its own routes
- Translates requests to the core engine
- Translates responses from the core engine

### Storage Backend (v0.1):

- bbolt (default)
- File‑based, single file, easy to mount in Docker

#### Future options:

- SQLite
- In‑memory (for CI)
- Pluggable interface for custom stores

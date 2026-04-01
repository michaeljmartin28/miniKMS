# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

---
## 0.2.0 — 2026-03-31
- Added Docker distribution
- Fixed npm SDK bug
- Unified versioning across all SDKs
- Added server version injection


## [Unreleased - v0.1.3] – 31 March 2026

### Added
- Full end‑to‑end test coverage for Node and Python SDKs
- Published Go, Node, and Python SDKs to pkg.go.dev, npm, and PyPI
- Added consistent `KeyMetadata` return type across all transports and SDKs

### Changed
- Updated Go, Node, and Python SDKs to match new `CreateKey` response shape
- Regenerated and corrected gRPC protobuf definitions to match engine behavior
- Improved consistency between HTTP and gRPC metadata responses
- Cleaned up SDK type definitions and binary handling across languages

### Fixed
- Resolved mismatches between SDK request/response types and backend models
- Fixed gRPC service definitions to use correct proto field names and structures
- Corrected several SDK bugs discovered during cross‑language testing
- Ensured consistent versioning and AAD handling across all SDKs

---

## [Unreleased v0.1.0] – 24 March 2026

### Added
- Complete core engine implementation  
  - Key creation, rotation, enabling/disabling  
  - Encryption and decryption  
  - Data key generation and decryption (envelope encryption)  
  - Versioning model with deterministic behavior  
- BoltDB-backed storage layer  
  - Key metadata bucket  
  - Key version bucket  
  - Serialization/deserialization helpers  
  - Test store for isolated engine tests  
- AES‑GCM crypto provider  
  - Deterministic envelope encryption  
  - Versioned key material handling  
- Full HTTP transport layer  
  - All endpoints implemented  
  - JSON request/response mapping  
  - Error → HTTP status mapping  
  - Comprehensive handler tests  
- Full gRPC transport layer  
  - Complete proto definitions  
  - RPC implementations for all engine operations  
  - Error → gRPC status mapping  
  - Initial gRPC handler tests  
- Dual‑server bootstrap (HTTP + gRPC)  
- Configuration system (`DefaultConfig`)  
- Architecture documentation, planning notes, and milestone structure  
- Test helpers and utilities for engine + transport layers  

### Changed
- Normalized algorithm, key state, and version naming across engine, proto, and transports
- Standardized timestamp formatting and version fields
- Improved consistency in request/response mapping between HTTP and gRPC
- Refined storage bucket layout and serialization logic
- Cleaned up error taxonomy and unified error mapping across transports

### Fixed
- Corrected mismatches between proto definitions and engine types
- Fixed BoltDB bucket initialization edge cases
- Resolved several transport-level error mapping inconsistencies
- Addressed minor naming inconsistencies (CreatedAt, version fields, etc.)



## [Unreleased] 13-March-2026

### Added

- Initial project scaffolding
- Core architecture docs
- Planning issues and milestones

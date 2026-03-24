# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

---

## [Unreleased] – 24 March 2026
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
- Error -> HTTP status mapping
- Comprehensive handler tests
- Full gRPC transport layer
- Complete proto definitions
- RPC implementations for all engine operations
- Error -> gRPC status mapping
- Initial gRPC handler tests
- Dual‑server bootstrap
- HTTP + gRPC servers running concurrently
- Shared engine instance
- Configuration system (DefaultConfig)
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

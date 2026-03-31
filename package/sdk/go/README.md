[![Go Reference](https://pkg.go.dev/badge/github.com/michaeljmartin28/minikms/package/sdk/go.svg)](https://pkg.go.dev/github.com/michaeljmartin28/minikms/package/sdk/go)[![Go Report Card](https://goreportcard.com/badge/github.com/michaeljmartin28/minikms)](https://goreportcard.com/report/github.com/michaeljmartin28/minikms)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](../../../LICENSE)

# miniKMS Official Go SDK

The official Go client for **miniKMS**, a lightweight Key Management Service designed for demos, local development, and portfolio‑grade backend/security engineering projects.

The SDK provides a clean, mechanical API for:

- Creating and managing customer master keys (CMKs)
- Encrypting and decrypting data
- Generating and decrypting data keys (envelope encryption)
- Rotating, enabling, and disabling keys
- Retrieving key metadata

It’s dependency‑minimal, concurrency‑safe, and structured the way Go developers expect.

---

## Installation

```bash
go get github.com/michaeljmartin28/minikms/sdk/go
```

## Quickstart

```go
package main

import (
    "context"
    "fmt"
    "github.com/michaeljmartin28/minikms/sdk/go/kms"
)

func main() {
    ctx := context.Background()

    client := kms.NewClient("http://localhost:8080")

    // Create a new CMK
    key, err := client.CreateKey(ctx, kms.CreateKeyParams{
        Algorithm: "AES-256_GCM",
        Name:      "example-key",
    })
    if err != nil {
        panic(err)
    }

    fmt.Println("Created key:", key.KeyID)

    // Encrypt some data
    enc, err := client.Encrypt(ctx, key.KeyID, kms.EncryptParams{
        Plaintext: []byte("hello world"),
    })
    if err != nil {
        panic(err)
    }

    // Decrypt it
    dec, err := client.Decrypt(ctx, key.KeyID, kms.DecryptParams{
        Ciphertext: enc.Ciphertext,
    })
    if err != nil {
        panic(err)
    }

    fmt.Println("Decrypted:", string(dec.Plaintext))
}
```

## Client

```go
client := kms.NewClient("http://localhost:8080")
```

The client is safe for concurrent use and performs:

- JSON marshaling/unmarshaling
- HTTP request construction
- error propagation
- typed response decoding

## Key Lifecycle

### Creating Keys

```go
key, err := client.CreateKey(ctx, kms.CreateKeyParams{
    Name:      "my-key",
    Algorithm: "AES-256-GCM",
})
```

Returns:

```go
type Key struct {
    KeyID     string
    Version   uint32
    CreatedAt time.Time
}
```

### Enable a key:

```go
_, err := client.EnableKey(ctx, key.KeyID)
```

### Disable a key:

```go
_, err := client.DisableKey(ctx, key.KeyID)
```

### Rotate a key:

```go
rot, err := client.RotateKey(ctx, key.KeyID)
fmt.Println(rot.Version)

```

Each returns updated metadata.

## Encrypting Data

```go
enc, err := client.Encrypt(ctx, key.KeyID, kms.EncryptParams{
    Plaintext:      []byte("hello world"),
    AdditionalData: []byte("optional AAD"),
})
```

Returns:

```go
type EncryptResponse struct {
    Ciphertext []byte
    AdditionalData []byte
    Version    uint32
    Algorithm  string
}
```

## Decrypting Data

```go
dec, err := client.Decrypt(ctx, key.KeyID, kms.DecryptParams{
    Ciphertext: enc.Ciphertext,
    AdditionalData: enc.AdditionalData,
    Version:    enc.Version,
})
```

Returns:

```go
type DecryptResponse struct {
    Plaintext []byte
}
```

## Data Keys (Envelope Encryption)

### Generate a data key

```go
dk, err := client.GenerateDataKey(ctx, key.KeyID)
fmt.Println(dk.PlaintextKey)  // []byte
fmt.Println(dk.EncryptedKey)  // base64-encoded []byte
```

### Decrypt a data key:

```go
pt, err := client.DecryptDataKey(ctx, key.KeyID, kms.DecryptDataKeyParams{
    EncryptedKey: dk.EncryptedKey,
})
fmt.Println(pt.PlaintextKey)
```

## Error Handling

All SDK methods return Go error values.
Future versions will include typed errors such as:

- `ErrKeyDisabled`
- `ErrKeyNotFound`
- `ErrInvalidCiphertext`

## Versioning

All SDK calls target:

```Code
/v1/...
```

When the API introduces /v2, the SDK will expose a new client or options without breaking existing code.

License
MIT

[![Go Reference](https://pkg.go.dev/badge/github.com/michaeljmartin28/minikms/sdk/go.svg)](https://pkg.go.dev/github.com/michaeljmartin28/minikms/sdk/go)[![Go Report Card](https://goreportcard.com/badge/github.com/michaeljmartin28/minikms)](https://goreportcard.com/report/github.com/michaeljmartin28/minikms)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](../../../LICENSE)

# miniKMS Go SDK

A lightweight Go client for interacting with miniKMS, a minimal Key Management Service designed for secure local development, demos, and portfolio‑grade backend/security engineering work.

The SDK provides a clean, mechanical API for:

- Creating and managing customer master keys (CMKs)
- Encrypting and decrypting data
- Generating and decrypting data keys (envelope encryption)
- Rotating, enabling, and disabling keys
- Retrieving key metadata

It’s dependency‑minimal, concurrency‑safe, and structured the way Go developers expect.

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

## Creating Keys

```go
key, err := client.CreateKey(ctx, kms.CreateKeyParams{
    Algorithm: "AES256_GCM",
    Name:      "my-key",
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

## Encrypting Data

```go
resp, err := client.Encrypt(ctx, keyID, kms.EncryptParams{
    Plaintext: []byte("secret"),
})
```

Returns ciphertext and metadata:

```go
type EncryptResponse struct {
    Ciphertext []byte
    Version    uint32
    KeyID      string
    Algorithm  string
}
```

## Decrypting Data

```go
resp, err := client.Decrypt(ctx, keyID, kms.DecryptParams{
    Ciphertext: ciphertext,
})
```

## Data Keys (Envelope Encryption)

Generate a data key

```go
dk, err := client.GenerateDataKey(ctx, keyID, kms.GenerateDataKeyParams{
    Bytes: 32,
})
```

Returns:

- plaintext DEK
- encrypted DEK
- metadata

Decrypt a data key:

```go
pt, err := client.DecryptDataKey(ctx, keyID, kms.DecryptDataKeyParams{
    EncryptedDataKey: dk.EncryptedDataKey,
})
```

## Key Lifecycle

Enable a key:

```go
client.EnableKey(ctx, keyID)
```

Disable a key:

```go
client.DisableKey(ctx, keyID)
```

Rotate:

```go
client.RotateKey(ctx, keyID)
```

Each returns updated metadata.

## Error Handling

All SDK methods return Go error values.
Future versions will include typed errors such as:

- `ErrKeyDisabled`
- `ErrKeyNotFound`
- `ErrInvalidCiphertext`

## Versioning

All SDK calls target:

```
/v1/...
```

When the API introduces /v2, the SDK will expose a new client or options without breaking existing code.

License
MIT

[![PyPI version](https://img.shields.io/pypi/v/minikms.svg)](https://pypi.org/project/minikms/)
[![PyPI downloads](https://img.shields.io/pypi/dm/minikms.svg)](https://pypi.org/project/minikms/)
![Python](https://img.shields.io/badge/Python-3.8+-blue)
![httpx](https://img.shields.io/badge/httpx-0.27+-blue)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/michaeljmartin28/minikms/blob/main/LICENSE)

# miniKMS Official Python SDK

The official Python client for **miniKMS**, a lightweight Key Management Service designed for demos, local development, and portfolio‑grade security engineering projects.

The SDK provides a clean, mechanical API for:

- Creating and managing customer master keys (CMKs)
- Encrypting and decrypting data
- Generating and decrypting data keys (envelope encryption)
- Rotating, enabling, and disabling keys
- Retrieving key metadata

---

## Installation

```bash
pip install minikms
```

## Quicksart

```python
from minikms import MiniKMS, CreateKeyRequest, EncryptRequest, DecryptRequest

kms = MiniKMS("http://localhost:8080")

# Create a key
key = kms.create_key(CreateKeyRequest(
    name="example",
    algorithm="AES-256-GCM",
))

# Encrypt data
enc = kms.encrypt(key.keyId, EncryptRequest(
    plaintext="hello world",
    additionalData="",
))

# Decrypt data
dec = kms.decrypt(key.keyId, DecryptRequest(
    ciphertext=enc.ciphertext,
    additionalData="",
    version=enc.version,
))

print(dec.plaintext)  # "hello world"
```

## Client

```python
from minikms import MiniKMS

kms = MiniKMS(base_url="http://localhost:8080")
```

Example:

```python
kms = MiniKMS("http://localhost:8080")
```

## Key Lifecycle

### Creating Keys

```python
from minikms import CreateKeyRequest

key = kms.create_key(CreateKeyRequest(
    name="my-key",
    algorithm="AES-256-GCM",
))
```

### Enable a key

```python
kms.enable_key(key.keyId)
```

### Disable a key

```python
kms.disable_key(key.keyId)
```

### Rotate a key

```python
resp = kms.rotate_key(key.keyId)
print(resp.version)
```

Each returns updated metadata.

## Encrypting Data

```python
from minikms import EncryptRequest

enc = kms.encrypt(key.keyId, EncryptRequest(
    plaintext="hello world",
    additionalData="optional AAD",
))
```

Returns ciphertext and metadata:

```python
{
  "ciphertext": str,   # base64
  "additionalData": str,
  "version": int,
  "algorithm": str,
}
```

## Decrypting Data

```python
from minikms import DecryptRequest

dec = kms.decrypt(key.keyId, DecryptRequest(
ciphertext=enc.ciphertext,
additionalData="optional AAD",
version=enc.version,
))
```

Returns:

```python
{
"plaintext": str # utf8 decoded
}
```

## Data Keys (Envelope Encryption)

### Generate a data key

```python
dk = kms.generate_data_key(key.keyId)
```

Returns:

```python
{
  "plaintextKey": str,   # utf8 decoded
  "encryptedKey": str,   # base64
  "keyVersion": int,
}
```

### Decrypt a data key

```python
dk2 = kms.decrypt_data_key(key.keyId, dk.encryptedKey)
print(dk2.plaintextKey)
```

## Error Handling

All errors raised by the SDK are httpx.HTTPStatusError or connection‑level exceptions.

Example:

```python
try:
    kms.encrypt("bad-id", EncryptRequest(plaintext="test", additionalData=""))
except Exception as err:
    print("Error:", err)
```

## Type Support

This SDK ships with full Python type hints via @dataclass models:

- request/response types
- key metadata
- encryption types
- data key types

## Versioning

All SDK calls target:

```Code
/v1/...
```

When the API introduces `/v2`, the SDK will expose a new client or options without breaking existing code.

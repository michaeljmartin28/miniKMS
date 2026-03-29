[![npm version](https://img.shields.io/npm/v/@minikms/sdk.svg)](https://www.npmjs.com/package/@minikms/sdk)
[![npm downloads](https://img.shields.io/npm/dm/@minikms/sdk.svg)](https://www.npmjs.com/package/@minikms/sdk)
![node-current](https://img.shields.io/node/v/@minikms/sdk)
![TypeScript](https://img.shields.io/badge/TypeScript-5.x-blue)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/michaeljmartin28/minikms/blob/main/LICENSE)

# miniKMS Official Node.js SDK

The official Node.js client for miniKMS, a lightweight Key Management Service designed for demos, local development, and portfolio‑grade security engineering projects.

The SDK provides a clean, mechanical API for:

- Creating and managing customer master keys (CMKs)
- Encrypting and decrypting data
- Generating and decrypting data keys (envelope encryption)
- Rotating, enabling, and disabling keys
- Retrieving key metadata

## Installation

```bash
npm install @minikms/sdk
```

## Quickstart

```ts
import { Client } from "@minikms/sdk";

const kms = new Client("http://localhost:8080");

const run = async () => {
  // Create a key
  const key = await kms.createKey({
    name: "example",
    algorithm: "AES-256-GCM",
  });

  // Encrypt data
  const enc = await kms.encrypt(key.keyId, {
    plaintext: "hello world",
  });

  // Decrypt data
  const dec = await kms.decrypt(key.keyId, {
    ciphertext: enc.ciphertext,
    version: enc.version,
  });

  console.log(dec.plaintext); // "hello world"
};

run();
```

## Client

```ts
const kms = new Client(baseURL: string);
```

### Example:

```ts
const kms = new Client("http://localhost:8080");
```

## Key Lifecycle

### Creating Keys

```ts
const key = await kms.createKey({
  name: "my-key",
  algorithm: "AES-256-GCM",
});
```

### Enable a key:

```ts
await kms.enableKey(key.keyId);
```

### Disable a key:

```ts
await kms.disableKey(key.keyId);
```

### Rotate:

```ts
await kms.rotateKey(key.keyId);
```

Each returns updated metadata.

## Encrypting Data

```ts
const enc = await kms.encrypt(key.keyId, {
  plaintext: "hello world",
  additionalData: "optional AAD",
});
```

Returns ciphertext and metadata:

```ts
{
  ciphertext: string; // base64
  version: number;
  keyId: string;
  algorithm: string;
}
```

## Decrypting Data

```ts
const dec = await kms.decrypt(key.keyId, {
  ciphertext: enc.ciphertext,
  version: enc.version,
});
```

Returns:

```ts
{
  plaintext: string; // utf8 decoded
}
```

## Data Keys (Envelope Encryption)

Generate a data key

```ts
const dk = await kms.generateDataKey(key.keyId, {
  additionalData: "optional AAD",
});
```

Returns:

```ts
{
  plaintextDEK: string; // utf8 decoded
  encryptedDEK: string; // base64
  version: number;
}
```

Decrypt a data key:

```ts
const dk2 = await kms.decryptDataKey(key.keyId, {
  encryptedDEK: dk.encryptedDEK,
  version: dk.version,
});
```

## Error Handling

All errors thrown by the SDK are instances of MiniKMSError:

```ts
try {
  await kms.encrypt("bad-id", { plaintext: "test" });
} catch (err) {
  if (err instanceof MiniKMSError) {
    console.error(err.status, err.body);
  }
}
```

## TypeScript Support

This SDK ships with full TypeScript definitions:

- request/response types
- key metadata
- encryption types
- data key types
- error types

Everything is strongly typed and editor‑friendly.

## Versioning

All SDK calls target:

```
/v1/...
```

When the API introduces /v2, the SDK will expose a new client or options without breaking existing code.

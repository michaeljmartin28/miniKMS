from dataclasses import dataclass

@dataclass(frozen=True)
class KeyMetadata:
    keyId: str
    createdAt: str
    state: str
    algorithm: str
    latestVersion: int

@dataclass(frozen=True)
class RotateKeyResponse:
    version: int

@dataclass(frozen=True)
class EnableKeyResponse:
    keyMetadata: KeyMetadata

@dataclass(frozen=True)
class DisableKeyResponse:
    keyMetadata: KeyMetadata

@dataclass(frozen=True)
class CreateKeyRequest:
    name: str
    algorithm: str

@dataclass(frozen=True)
class EncryptRequest:
    plaintext: str
    additionalData: str

@dataclass(frozen=True)
class EncryptResponse:
    ciphertext: str
    additionalData: str
    version: int 
    algorithm: str

@dataclass(frozen=True)
class DecryptRequest:
    ciphertext: str
    additionalData: str
    version: int 

@dataclass(frozen=True)
class DecryptResponse:
    plaintext: str

@dataclass(frozen=True)
class GenerateDataKeyResponse:
    plaintextKey: str
    encryptedKey: str
    keyVersion: int

@dataclass(frozen=True)
class DecryptDataKeyResponse:
    plaintextKey: str



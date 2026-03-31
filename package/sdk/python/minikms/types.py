from dataclasses import dataclass
from typing import Optional

@dataclass
class KeyMetadata:
    keyId: str
    createdAt: str
    state: str
    algorithm: str
    latestVersion: int

@dataclass
class RotateKeyResponse:
    version: int

@dataclass
class EnableKeyResponse:
    keyMetadata: KeyMetadata

@dataclass
class DisableKeyResponse:
    keyMetadata: KeyMetadata

@dataclass
class CreateKeyRequest:
    name: str
    algorithm: str

@dataclass
class CreateKeyResponse:
  keyId: str
  name: str
  algorithm: str
  createdAt: str

@dataclass
class EncryptRequest:
    plaintext: str
    additionalData: Optional[str] = None

@dataclass
class EncryptResponse:
    keyId: str
    ciphertext: str
    version: int 
    algorithm: str
    additionalData: Optional[str] = None

@dataclass
class DecryptRequest:
    ciphertext: str
    version: int 
    additionalData: Optional[str] = None

@dataclass
class DecryptResponse:
    plaintext: str

@dataclass
class GenerateDataKeyRequest:
    additionalData: Optional[str] = None

@dataclass
class GenerateDataKeyResponse:
    plaintextDEK: str
    encryptedDEK: str
    version: int

@dataclass
class DecryptDataKeyRequest:
    encryptedDEK: str
    version: int
    additionalData: Optional[str] = None

@dataclass
class DecryptDataKeyResponse:
    plaintextDEK: str



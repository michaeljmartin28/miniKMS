import httpx
from .types import (
    KeyMetadata,
    RotateKeyResponse,
    EnableKeyResponse,
    DisableKeyResponse,
    CreateKeyRequest,
    EncryptRequest,
    EncryptResponse,
    DecryptRequest,
    DecryptResponse,
    GenerateDataKeyResponse,
    DecryptDataKeyResponse,
)

class MiniKMS:
    def __init__(self, base_url: str):
        self.base_url = base_url.rstrip("/")
        self.client = httpx.Client()

    def create_key(self, req: CreateKeyRequest) -> KeyMetadata:
        resp = self.client.post(
            f"{self.base_url}/v1/keys",
            json=req.__dict__,
        )
        resp.raise_for_status()
        return KeyMetadata(**resp.json())

    def rotate_key(self, key_id: str) -> RotateKeyResponse:
        resp = self.client.post(f"{self.base_url}/v1/keys/{key_id}/rotate")
        resp.raise_for_status()
        return RotateKeyResponse(**resp.json())

    def enable_key(self, key_id: str) -> EnableKeyResponse:
        resp = self.client.post(f"{self.base_url}/v1/keys/{key_id}/enable")
        resp.raise_for_status()
        return EnableKeyResponse(**resp.json())

    def disable_key(self, key_id: str) -> DisableKeyResponse:
        resp = self.client.post(f"{self.base_url}/v1/keys/{key_id}/disable")
        resp.raise_for_status()
        return DisableKeyResponse(**resp.json())

    def encrypt(self, key_id: str, req: EncryptRequest) -> EncryptResponse:
        resp = self.client.post(
            f"{self.base_url}/v1/keys/{key_id}/encrypt",
            json=req.__dict__,
        )
        resp.raise_for_status()
        return EncryptResponse(**resp.json())

    def decrypt(self, key_id: str, req: DecryptRequest) -> DecryptResponse:
        resp = self.client.post(
            f"{self.base_url}/v1/keys/{key_id}/decrypt",
            json=req.__dict__,
        )
        resp.raise_for_status()
        return DecryptResponse(**resp.json())

    def generate_data_key(self, key_id: str) -> GenerateDataKeyResponse:
        resp = self.client.post(
            f"{self.base_url}/v1/keys/{key_id}/generate-data-key"
        )
        resp.raise_for_status()
        return GenerateDataKeyResponse(**resp.json())

    def decrypt_data_key(self, key_id: str, encrypted_key: str) -> DecryptDataKeyResponse:
        resp = self.client.post(
            f"{self.base_url}/v1/keys/{key_id}/decrypt-data-key",
            json={"encryptedKey": encrypted_key},
        )
        resp.raise_for_status()
        return DecryptDataKeyResponse(**resp.json())

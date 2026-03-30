import httpx
from .types import GenerateDataKeyResponse, DecryptDataKeyResponse

class MiniKMS:
    def __init__(self, base_url: str):
        self.base_url = base_url.rstrip("/")
        self.client = httpx.Client()

    def generate_data_key(self, key_id: str) -> GenerateDataKeyResponse:
        resp = self.client.post(f"{self.base_url}/v1/keys/{key_id}:generateDataKey")
        resp.raise_for_status()
        return GenerateDataKeyResponse(**resp.json())

    def decrypt_data_key(self, key_id: str, encrypted_key: str) -> DecryptDataKeyResponse:
        resp = self.client.post(
            f"{self.base_url}/v1/keys/{key_id}:decryptDataKey",
            json={"encryptedKey": encrypted_key},
        )
        resp.raise_for_status()
        return DecryptDataKeyResponse(**resp.json())

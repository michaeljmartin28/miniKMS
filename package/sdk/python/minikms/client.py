import httpx
import base64
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
    GenerateDataKeyRequest,
    GenerateDataKeyResponse,
    DecryptDataKeyRequest,
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
        data = resp.json()
        return EnableKeyResponse(
            keyMetadata=KeyMetadata(**resp.json())
        )
    
    def disable_key(self, key_id: str) -> DisableKeyResponse:
        resp = self.client.post(f"{self.base_url}/v1/keys/{key_id}/disable")
        resp.raise_for_status()
        return DisableKeyResponse(
            keyMetadata=KeyMetadata(**resp.json())
        )

    def encrypt(self, key_id: str, req: EncryptRequest) -> EncryptResponse:

        payload = {
            "plaintext": base64.b64encode(req.plaintext).decode("ascii"),
            "additionalData": base64.b64encode(req.additionalData).decode("ascii") if req.additionalData else None,
        }

        resp = self.client.post(
            f"{self.base_url}/v1/keys/{key_id}/encrypt",
            json=payload,
        )
        resp.raise_for_status()
        enc_resp = EncryptResponse(**resp.json())
        enc_resp.ciphertext = base64.b64decode(enc_resp.ciphertext)
        if enc_resp.additionalData:
            enc_resp.additionalData = base64.b64decode(enc_resp.additionalData)
        return enc_resp

    def decrypt(self, key_id: str, req: DecryptRequest) -> DecryptResponse:
        payload = {
            "ciphertext": base64.b64encode(req.ciphertext).decode("ascii"),
            "additionalData": base64.b64encode(req.additionalData).decode("ascii") if req.additionalData else None,
            "version": req.version
        }
        resp = self.client.post(
            f"{self.base_url}/v1/keys/{key_id}/decrypt",
            json=payload,
        )
        resp.raise_for_status()
        dec_resp =  DecryptResponse(**resp.json())
        dec_resp.plaintext = base64.b64decode(dec_resp.plaintext)
        return dec_resp

    def generate_data_key(self, key_id: str, req: GenerateDataKeyRequest) -> GenerateDataKeyResponse:
        payload = {
            "additionalData": base64.b64encode(req.additionalData).decode("ascii") if req.additionalData else None,
        }
        resp = self.client.post(
            f"{self.base_url}/v1/keys/{key_id}/generate-data-key",
            json=payload,
        )
        resp.raise_for_status()
        dk_resp = GenerateDataKeyResponse(**resp.json())
        dk_resp.plaintextDEK = base64.b64decode(dk_resp.plaintextDEK)
        dk_resp.encryptedDEK = base64.b64decode(dk_resp.encryptedDEK)
        return dk_resp

    def decrypt_data_key(self, key_id: str, req: DecryptDataKeyRequest) -> DecryptDataKeyResponse:
        payload = {
                "encryptedDEK": base64.b64encode(req.encryptedDEK).decode("ascii"),
                "version": req.version,
                "additionalData": base64.b64encode(req.additionalData).decode("ascii") if req.additionalData else None,
            }
        resp = self.client.post(
            f"{self.base_url}/v1/keys/{key_id}/decrypt-data-key",
            json=payload,
        )
        resp.raise_for_status()
        print(resp.json())

        dk_resp =  DecryptDataKeyResponse(**resp.json())
        dk_resp.plaintextDEK = base64.b64decode(dk_resp.plaintextDEK)
        return dk_resp


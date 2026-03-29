import {
  CreateKeyRequest,
  CreateKeyResponse,
  EncryptRequest,
  EncryptResponse,
  DecryptRequest,
  DecryptResponse,
  GenerateDataKeyRequest,
  GenerateDataKeyResponse,
  DecryptDataKeyRequest,
  DecryptDataKeyResponse,
  RotateKeyResponse,
  EnableKeyResponse,
  DisableKeyResponse,
} from "./types.js";

import { MiniKMSError } from "./errors.js";

export class Client {
  constructor(private baseUrl: string) {}

  private async do<T>(method: string, path: string, body?: any): Promise<T> {
    console.log(body);
    const res = await fetch(`${this.baseUrl}${path}`, {
      method,
      headers: { "Content-Type": "application/json" },
      body: body ? JSON.stringify(body) : undefined,
    });

    if (!res.ok) {
      const text = await res.text();
      throw new MiniKMSError(`HTTP: ${res.status}`, res.status, text);
    }

    return res.json() as Promise<T>;
  }

  private toBase64(input?: string): string | undefined {
    return input ? Buffer.from(input, "utf8").toString("base64") : undefined;
  }

  private fromBase64(input: string): string {
    return Buffer.from(input, "base64").toString("utf8");
  }

  createKey(req: CreateKeyRequest) {
    return this.do<CreateKeyResponse>("POST", "/v1/keys", req);
  }

  encrypt(keyId: string, req: EncryptRequest) {
    const body = {
      plaintext: this.toBase64(req.plaintext),
      additionalData: this.toBase64(req.additionalData),
    };
    return this.do<EncryptResponse>("POST", `/v1/keys/${keyId}/encrypt`, body);
  }

  decrypt(keyId: string, req: DecryptRequest) {
    const body = {
      ciphertext: req.ciphertext,
      additionalData: this.toBase64(req.additionalData),
      version: req.version,
    };
    return this.do<DecryptResponse>(
      "POST",
      `/v1/keys/${keyId}/decrypt`,
      body,
    ).then((res) => ({ plaintext: this.fromBase64(res.plaintext) }));
  }

  generateDataKey(keyId: string, req: GenerateDataKeyRequest) {
    const body = {
      additionalData: this.toBase64(req.additionalData),
    };
    return this.do<GenerateDataKeyResponse>(
      "POST",
      `/v1/keys/${keyId}/generate-data-key`,
      body,
    ).then((res) => ({
      plaintextDEK: this.fromBase64(res.plaintextDEK),
      encryptedDEK: res.encryptedDEK,
      version: res.version,
    }));
  }

  decryptDataKey(keyId: string, req: DecryptDataKeyRequest) {
    const body = {
      encryptedDEK: req.encryptedDEK, // already base64
      version: req.version,
      additionalData: this.toBase64(req.additionalData),
    };
    return this.do<DecryptDataKeyResponse>(
      "POST",
      `/v1/keys/${keyId}/decrypt-data-key`,
      body,
    ).then((res) => ({
      plaintextDEK: this.fromBase64(res.plaintextDEK),
    }));
  }

  rotateKey(keyId: string) {
    return this.do<RotateKeyResponse>("POST", `/v1/keys/${keyId}/rotate`);
  }

  enable(keyId: string) {
    return this.do<EnableKeyResponse>("POST", `/v1/keys/${keyId}/enable`);
  }

  disable(keyId: string) {
    return this.do<DisableKeyResponse>("POST", `/v1/keys/${keyId}/disable`);
  }
}

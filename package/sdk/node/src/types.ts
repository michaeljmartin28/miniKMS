export interface KeyMetadata {
  keyId: string;
  algorithm: string;
  createdAt: string;
  state: string;
  latestVersion: number;
}

export interface CreateKeyRequest {
  name: string;
  algorithm: string;
}

export interface EncryptRequest {
  plaintext: string;
  additionalData?: string;
}

export interface EncryptResponse {
  ciphertext: string;
  version: number;
  keyId: string;
  algorithm: string;
}

export interface DecryptRequest {
  ciphertext: string;
  additionalData?: string;
  version: number;
}

export interface DecryptResponse {
  plaintext: string;
}

export interface GenerateDataKeyRequest {
  additionalData?: string;
}

export interface GenerateDataKeyResponse {
  plaintextDEK: string;
  encryptedDEK: string;
  version: number;
}

export interface DecryptDataKeyRequest {
  encryptedDEK: string;
  version: number;
  additionalData?: string;
}

export interface DecryptDataKeyResponse {
  plaintextDEK: string;
}

export interface RotateKeyResponse {
  version: number;
}

export type EnableKeyResponse = KeyMetadata;
export type DisableKeyResponse = KeyMetadata;
export type CreateKeyResponse = KeyMetadata;

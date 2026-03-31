from minikms import (
    MiniKMS,
    CreateKeyRequest,
    EncryptRequest,
    DecryptRequest,
    GenerateDataKeyRequest,
    DecryptDataKeyRequest
)

def main():
    kms = MiniKMS("http://localhost:8080")

    print(" Creating key ")
    key = kms.create_key(CreateKeyRequest(
        name="python-test-key",
        algorithm="AES-256-GCM",
    ))
    print("KeyMetadata:", key)

    key_id = key.keyId

    print("\n Encrypting ")
    plaintext = b"hello from python"
    aad = b"test-aad"

    enc = kms.encrypt(key_id, EncryptRequest(
        plaintext=plaintext,
        additionalData=aad,
    ))
    print("Ciphertext (bytes):", enc.ciphertext)
    print("Version:", enc.version)

    print("\n Decrypting ")
    dec = kms.decrypt(key_id, DecryptRequest(
        ciphertext=enc.ciphertext,
        additionalData=aad,
        version=enc.version,
    ))
    print("Decrypted plaintext:", dec.plaintext)
    assert dec.plaintext == plaintext

    print("\n Generating Data Key ")
    dk = kms.generate_data_key(key_id, GenerateDataKeyRequest(
        additionalData=aad,
    ))
    print("Plaintext DEK:", dk.plaintextDEK)
    print("Encrypted DEK:", dk.encryptedDEK)
    print("Version:", dk.version)

    print("\n Decrypting Data Key ")
    ddk = kms.decrypt_data_key(key_id, DecryptDataKeyRequest(
        encryptedDEK=dk.encryptedDEK,
        version=key.latestVersion,
        additionalData=aad,
    ))
    print("Decrypted DEK:", ddk.plaintextDEK)
    assert ddk.plaintextDEK == dk.plaintextDEK

    print("\n Rotating Key ")
    rot = kms.rotate_key(key_id)
    print("RotateKeyResponse:", rot)

    print("\n Disabling Key ")
    dis = kms.disable_key(key_id)
    print("DisableKeyResponse:", dis)

    print("\n Enabling Key ")
    en = kms.enable_key(key_id)
    print("EnableKeyResponse:", en)

    print("\n ALL TESTS PASSED ")

if __name__ == "__main__":
    main()

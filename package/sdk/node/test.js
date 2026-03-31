import { Client } from "./dist/index.js";

const kms = new Client("http://localhost:8080");

const run = async () => {
  console.log(" Creating key ");
  const key = await kms.createKey({
    name: "node-test-key",
    algorithm: "AES-256-GCM",
  });
  console.log("KeyMetadata:", key);

  const keyId = key.keyId;

  console.log("\n Encrypting ");
  const plaintext = Buffer.from("hello from node");
  const aad = Buffer.from("node-aad");

  const enc = await kms.encrypt(keyId, {
    plaintext,
    additionalData: aad,
  });
  console.log("Encrypted:", enc);

  console.log("\n Decrypting ");
  const dec = await kms.decrypt(keyId, {
    ciphertext: enc.ciphertext,
    additionalData: aad,
    version: enc.version,
  });
  console.log("Decrypted:", dec);
  console.log("Decrypted plaintext:", dec.plaintext.toString());

  console.log("\n Generating Data Key ");
  const dk = await kms.generateDataKey(keyId);
  console.log("Plaintext DEK:", dk.plaintextDEK);
  console.log("Encrypted DEK:", dk.encryptedDEK);

  console.log("\n Decrypting Data Key ");
  const ddk = await kms.decryptDataKey(keyId, {
    encryptedDEK: dk.encryptedDEK,
    version: dk.version,
    additionalData: Buffer.from(""),
  });
  console.log("Decrypted DEK:", ddk.plaintextDEK);

  console.log("\n Rotating Key ");
  const rot = await kms.rotateKey(keyId);
  console.log("RotateKeyResponse:", rot);

  console.log("\n Disabling Key ");
  const dis = await kms.disable(keyId);
  console.log("DisableKeyResponse:", dis);

  console.log("\n Enabling Key ");
  const en = await kms.enable(keyId);
  console.log("EnableKeyResponse:", en);

  console.log("\n ALL NODE TESTS PASSED ");
};

run().catch((err) => {
  console.error("Test failed:", err);
});

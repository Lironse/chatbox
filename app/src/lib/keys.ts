export async function generateKeyPair(): Promise<{ publicKey: string, privateKey: string }> {
    // Generate a key pair
    const keyPair = await crypto.subtle.generateKey(
      {
        name: 'RSA-OAEP',
        modulusLength: 2048,
        publicExponent: new Uint8Array([0x01, 0x00, 0x01]),
        hash: 'SHA-256'
      },
      true,
      ['encrypt', 'decrypt']
    );
  
    // Export the public key
    const publicKey = await crypto.subtle.exportKey('spki', keyPair.publicKey);
  
    // Export the private key
    const privateKey = await crypto.subtle.exportKey('pkcs8', keyPair.privateKey);
  
    // Convert the exported keys to Base64 strings
    const publicKeyString = btoa(String.fromCharCode(...new Uint8Array(publicKey)));
    const privateKeyString = btoa(String.fromCharCode(...new Uint8Array(privateKey)));
  
    return { publicKey: publicKeyString, privateKey: privateKeyString };
}
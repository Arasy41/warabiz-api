const crypto = require('crypto');

function parseSecretKey(secretKey) {
    while (secretKey.length < 32) {
        secretKey += ' ';
    }
    return Buffer.from(secretKey);
}

// Encode using AES-GCM algorithm and secret key
function encryptAESGCM(message, secretKey) {
    const byteKey = parseSecretKey(secretKey);
    const iv = crypto.randomBytes(12);
    const cipher = crypto.createCipheriv('aes-256-gcm', byteKey, iv);
    const encrypted = Buffer.concat([cipher.update(message, 'utf8'), cipher.final()]);
    const tag = cipher.getAuthTag();
    return {
        iv: iv.toString('base64'),
        encryptedData: Buffer.concat([encrypted, tag]).toString('base64'),
    };
}

function decryptAESGCM(encryptedData, iv, secretKey) {
    let key = parseSecretKey(secretKey);
    const decipher = crypto.createDecipheriv('aes-256-gcm', key, Buffer.from(iv, 'base64'));
    const decodedData = Buffer.from(encryptedData, 'base64');

    const tag = decodedData.slice(-16);
    const encrypted = decodedData.slice(0, -16);

    decipher.setAuthTag(tag);
    const decrypted = Buffer.concat([decipher.update(encrypted), decipher.final()]);

    return decrypted.toString('utf8');
}

const message = 'Hello, World!';
const secretKey = 'isth1sb0nd5ur3ty';

const encryptedData = encryptAESGCM(message, secretKey);
console.log('IV:', encryptedData.iv);
console.log('Encrypted Data:', encryptedData.encryptedData);

const decryptedMessage = decryptAESGCM(encryptedData.encryptedData, encryptedData.iv, secretKey);
console.log('Decrypted Message:', decryptedMessage);

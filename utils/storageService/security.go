package storageService

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
)

// initialized in main.go or in the kms package
var hostBytesHMAC []byte
var hostBytesDEK []byte

// Create JWE using AES GCM 256 for the DEK
func EncryptPlaintextAndSerializeRawJWE(plainTextData *string, dekBytes []byte) (serializedRaw json.RawMessage, errMsg string) {
	if plainTextData == nil {
		return nil, "Plaintext data cannot be nil"
	}

	block, err := aes.NewCipher(dekBytes[:])
	if err != nil {
		return nil, fmt.Sprintf("Failed to initialize AES cipher: %v", err)
	}

	// Generate a random nonce
	nonce := make([]byte, 12)
	// You would normally populate nonce with a cryptographically secure random sequence.

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Sprintf("Failed to initialize AES GCM: %v", err)
	}

	ciphertext := aesgcm.Seal(nil, nonce, []byte(*plainTextData), nil)

	// Here we manually craft a JWE structure. Normally you'd also include more fields like protected headers.
	jwe := map[string]string{
		"nonce":      base64.RawURLEncoding.EncodeToString(nonce),
		"ciphertext": base64.RawURLEncoding.EncodeToString(ciphertext),
		// Add more fields like protected headers, etc. as per JWE specification
	}

	rawJWE, err := json.Marshal(jwe)
	if err != nil {
		return nil, fmt.Sprintf("Failed to serialize JWE: %v", err)
	}

	return rawJWE, ""
}

// to be used in main.go or in the kms package
var InitHostDEK = func(seed []byte) {
	// Derive the encryption key from seed using SHA-256 (it gives us a 32 byte output for AES 256)
	keyArray := sha256.Sum256(seed)
	hostBytesDEK = keyArray[:]
}

// The HMAC key is the seed received to generate pseudo-anonymized data
var InitHostHMAC = func(seed []byte) {
	hostBytesHMAC = seed
}

var getHostPrivateBytesHMAC = func() []byte {
	return hostBytesHMAC
}

var getHostPrivateBytesDEK = func() []byte {
	return hostBytesDEK
}

func EncryptToRawJWE(plainTextData *string) (serializedRaw json.RawMessage, errMsg string) {
	// both recipient and sender are the same (the host)
	return EncryptPlaintextAndSerializeRawJWE(plainTextData, getHostPrivateBytesDEK())
}

var DecryptByRawJWE = func(serializedRawJWE *json.RawMessage) (decryptedDatabytes, protectedHeaderBytes []byte) {
	components := strings.Split(string(*serializedRawJWE), ".")

	if len(components) != 5 {
		return nil, nil // Error case, serializedRawJWE doesn't have expected format
	}

	protectedHeader, err := base64.RawURLEncoding.DecodeString(components[0])
	if err != nil {
		return nil, nil // Error handling here
	}

	iv, err := base64.RawURLEncoding.DecodeString(components[2])
	if err != nil {
		return nil, nil // Error handling here
	}

	ciphertext, err := base64.RawURLEncoding.DecodeString(components[3])
	if err != nil {
		return nil, nil // Error handling here
	}

	// Assuming hostBytesDEK is initialized and is the correct decryption key
	block, err := aes.NewCipher(hostBytesDEK)
	if err != nil {
		return nil, nil // Error handling here
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, nil // Error handling here
	}

	decryptedData, err := aesgcm.Open(nil, iv, ciphertext, nil)
	if err != nil {
		return nil, nil // Error handling here
	}

	return decryptedData, protectedHeader
}

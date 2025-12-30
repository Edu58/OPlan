package crypto

import (
	"crypto/rand"
	"crypto/sha512"
	"crypto/subtle"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

const (
	memory      = 64 * 1024 // 64 MB
	iterations  = 3
	parallelism = 2
	saltLength  = 16
	keyLength   = 32
)

// Hash a string using SHA512
func HashStringSHA512(s string) string {
	h := sha512.Sum512([]byte(s))
	return hex.EncodeToString(h[:])
}

// HashPassword hashes a password using Argon2id
func HashPassword(password string) (string, error) {
	// Generate random salt
	salt := make([]byte, saltLength)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}

	// Hash the password
	hash := argon2.IDKey([]byte(password), salt, iterations, memory, parallelism, keyLength)

	// Encode to base64
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	// Format: $argon2id$v=19$m=65536,t=3,p=2$salt$hash
	encodedHash := fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version, memory, iterations, parallelism, b64Salt, b64Hash)

	return encodedHash, nil
}

// VerifyPassword verifies a password against its hash
func VerifyPassword(password, encodedHash string) (bool, error) {
	// Extract parameters from encoded hash
	parts := strings.Split(encodedHash, "$")
	if len(parts) != 6 {
		return false, fmt.Errorf("invalid hash format")
	}

	var version int
	var m, t uint32
	var p uint8
	_, err := fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", &m, &t, &p)
	if err != nil {
		return false, err
	}

	_, err = fmt.Sscanf(parts[2], "v=%d", &version)

	if err != nil {
		return false, err
	}

	if version != argon2.Version {
		return false, fmt.Errorf("incompatible argon2 version")
	}

	// Decode salt and hash
	salt, err := base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil {
		return false, err
	}

	decodedHash, err := base64.RawStdEncoding.DecodeString(parts[5])
	if err != nil {
		return false, err
	}

	// Hash the password with the same parameters
	newHash := argon2.IDKey([]byte(password), salt, t, m, p, uint32(len(decodedHash)))

	// Compare hashes (constant-time comparison)
	return subtle.ConstantTimeCompare(decodedHash, newHash) == 1, nil
}

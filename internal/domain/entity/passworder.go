package entity

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/abc-valera/netsly-golang/internal/domain/util/coderr"
	"golang.org/x/crypto/argon2"
)

var ErrPasswordDontMatch = coderr.NewCodeMessage(coderr.CodeInvalidArgument, "Provided password doesn't match the original one")

type IPassworder interface {
	// HashPassword returns hash of the provided password
	HashPassword(password string) (string, error)

	// CheckPassword checks if provided password matches provided hash,
	// if matches returns nil, else returns error
	CheckPassword(password, hash string) error
}

func newPassworder(dep IDependency) IPassworder {
	return &argon2idPassworder{dep}
}

const (
	memory      uint32 = 64 * 1024
	iterations  uint32 = 3
	parallelism uint8  = 2
	saltLength  uint32 = 16
	keyLength   uint32 = 32
)

// Argon2 is currently considered the most secure hashing algorithm.
// It has three variants:
//   - Argon2d, which maximizes resistance to GPU cracking attacks;
//   - Argon2i, which is optimized to resist side-channel attacks;
//   - Argon2id, which is a hybrid of both.
//
// The [OWASP Password Storage Cheat Sheet]
// recommends using the hybrid Argon2id algorithmfor password storage
// with a minimum configuration of 19 MiB of memory,
// an iteration count of 2, and 1 degree of parallelism.
//
// [OWASP Password Storage Cheat Sheet]: https://cheatsheetseries.owasp.org/cheatsheets/Password_Storage_Cheat_Sheet.html#password-hashing-algorithms
type argon2idPassworder struct {
	IDependency
}

func (argon2idPassworder) HashPassword(password string) (string, error) {
	// Generate a cryptographically secure random salt.
	salt, err := generateRandomBytes(saltLength)
	if err != nil {
		return "", err
	}

	// Pass the plaintext password, salt and parameters to the argon2.IDKey
	// function. This will generate a hash of the password using the Argon2id variant.
	hash := argon2.IDKey([]byte(password), salt, iterations, memory, parallelism, keyLength)

	// Base64 encode the salt and hashed password.
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	// Return a string using the standard encoded hash representation.
	encodedHash := fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, memory, iterations, parallelism, b64Salt, b64Hash)

	return encodedHash, nil
}

func (argon2idPassworder) CheckPassword(password, encodedHash string) error {
	// Extract the parameters, salt and derived key from the encoded password hash.
	salt, hash, err := decodeHash(encodedHash)
	if err != nil {
		return coderr.NewInternalErr(err)
	}

	// Derive the key from the other password using the same parameters.
	otherHash := argon2.IDKey([]byte(password), salt, iterations, memory, parallelism, keyLength)

	// Check that the contents of the hashed passwords are identical. Note
	// that we are using the subtle.ConstantTimeCompare() function for this
	// to help prevent timing attacks.
	if subtle.ConstantTimeCompare(hash, otherHash) == 1 {
		return nil
	}
	return ErrPasswordDontMatch
}

func generateRandomBytes(n uint32) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func decodeHash(encodedHash string) (salt, hash []byte, err error) {
	vals := strings.Split(encodedHash, "$")
	if len(vals) != 6 {
		return nil, nil, ErrPasswordDontMatch
	}

	var version int
	_, err = fmt.Sscanf(vals[2], "v=%d", &version)
	if err != nil {
		return nil, nil, err
	}
	if version != argon2.Version {
		return nil, nil, ErrPasswordDontMatch
	}

	salt, err = base64.RawStdEncoding.Strict().DecodeString(vals[4])
	if err != nil {
		return nil, nil, err
	}

	hash, err = base64.RawStdEncoding.Strict().DecodeString(vals[5])
	if err != nil {
		return nil, nil, err
	}

	return salt, hash, nil
}

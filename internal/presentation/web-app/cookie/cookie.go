package cookie

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"net/http"

	"github.com/abc-valera/netsly-api-golang/internal/domain/coderr"
)

const (
	AccessTokenKey  = "netsly_access_token"
	RefreshTokenKey = "netsly_refresh_token"
)

// NOTE: should be env variable
var secretKey = []byte("9aa5e83710a7ddf90aa2b426e82ccfb5ccf5890c771a590e0d47dcb5119f034b")

var (
	ErrNoCookie     = coderr.NewInternal(errors.New("no cookie"))
	ErrInvalidValue = coderr.NewInternal(errors.New("invalid cookie value"))
)

// Get returns the value associated with the key
func Get(r *http.Request, key string) (string, error) {
	cookie, err := r.Cookie(key)
	if err != nil {
		return "", ErrNoCookie
	}

	// Decode the base64-encoded cookie value
	signedValueBytes, err := base64.URLEncoding.DecodeString(cookie.Value)
	if err != nil {
		return "", ErrInvalidValue
	}
	signedValue := string(signedValueBytes)

	// Check if the cookie value contains a HMAC signature
	if len(signedValue) < sha256.Size {
		return "", ErrInvalidValue
	}

	// Split the signature and original cookie value
	signature := signedValue[:sha256.Size]
	value := signedValue[sha256.Size:]

	// Recalculate the HMAC signature of the cookie name and original value
	mac := hmac.New(sha256.New, secretKey)
	mac.Write([]byte(key))
	mac.Write([]byte(value))
	expectedSignature := mac.Sum(nil)

	// Check if the recalculated signature matches the one in the cookie
	if !hmac.Equal([]byte(signature), expectedSignature) {
		return "", ErrInvalidValue
	}

	return value, nil
}

// Set sets the value associated with the key (with HMAC signature).
func Set(w http.ResponseWriter, key string, value string) {
	// Calculate a HMAC signature of the cookie name and value with the secret key
	mac := hmac.New(sha256.New, secretKey)
	mac.Write([]byte(key))
	mac.Write([]byte(value))
	signature := mac.Sum(nil)

	// Prepend the signature to the cookie value
	value = string(signature) + value

	// Encode the cookie value using base64
	value = base64.URLEncoding.EncodeToString([]byte(value))

	cookie := http.Cookie{
		Name:     key,
		Value:    value,
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
		MaxAge:   60 * 60 * 24 * 365,
		Secure:   true,
	}

	http.SetCookie(w, &cookie)
}

func Delete(w http.ResponseWriter, key string) {
	cookie := http.Cookie{
		Name:     key,
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		MaxAge:   -1,
		Secure:   true,
	}

	http.SetCookie(w, &cookie)
}

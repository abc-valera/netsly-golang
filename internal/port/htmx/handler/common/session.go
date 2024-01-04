package common

import "github.com/gofiber/fiber/v2"

const (
	AccessTokenKey  = "flug_access_token"
	RefreshTokenKey = "flug_refresh_token"
)

// Get returns the value associated with the key
func Get(ctx *fiber.Ctx, key string) string {
	return "unimplemented"
}

// Set sets the value associated with the key
func Set(ctx *fiber.Ctx, key string, value string) {
}

// Delete deletes the value associated with the key
func Delete(ctx *fiber.Ctx, key string) {
}

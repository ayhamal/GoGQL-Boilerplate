package middleware

// import (
// 	"github.com/ayhamal/gogql-boilerplate/env"

// 	"github.com/gofiber/fiber/v2"
// 	jwtware "github.com/gofiber/jwt/v3"
// )

// // Protected protect routes
// func Protected(env *env.Env) fiber.Handler {
// 	return jwtware.New(jwtware.Config{
// 		SigningKey:   []byte(env.App.SigningKey),
// 		ErrorHandler: jwtError,
// 	})
// }

// func jwtError(c *fiber.Ctx, err error) error {
// 	if err.Error() == "Missing or malformed JWT" {
// 		return c.Status(fiber.StatusBadRequest).
// 			JSON(fiber.Map{"status": "error", "message": "Missing or malformed JWT", "data": nil})
// 	}
// 	return c.Status(fiber.StatusUnauthorized).
// 		JSON(fiber.Map{"status": "error", "message": "Invalid or expired JWT", "data": nil})
// }

package middleware

// import (
// 	"log"

// 	"github.com/gofiber/fiber/v2"
// 	"github.com/golang-jwt/jwt/v4"
// )

// func LogMiddleware(c *fiber.Ctx) error {
//     // Log the IP and User-Agent of the request
//     ip := c.IP()
//     ua := c.Get("User-Agent")
//     us := c.Locals("user").(*jwt.Token)
//     claims := us.Claims.(jwt.MapClaims)

//     // log.Println(claims)
    
//     log.Printf("%s - Request from IP %s with User Agent %s\n", claims["iss"], ip, ua)

//     // Call the next handler function
//     return c.Next()
// }
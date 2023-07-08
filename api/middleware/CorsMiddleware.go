package middleware

import "github.com/rs/cors"

// BuildHandler build cors handler
func BuildCorsHandler() *cors.Cors {
	// CORS setup
	return cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4000", "http://localhost:8080"},
		AllowCredentials: true,
		Debug:            false,
	})
}
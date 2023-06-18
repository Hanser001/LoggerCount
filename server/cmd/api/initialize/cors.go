package initialize

import (
	"github.com/hertz-contrib/cors"
	"summer/server/shared/consts"
	"time"
)

// InitCors return cors.Config.
func InitCors() cors.Config {
	return cors.Config{
		AllowOrigins:     []string{consts.CorsAddress},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization", "Token", "Accept"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
}

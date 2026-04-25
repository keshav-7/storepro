package middleware

import (
	"github.com/gin-gonic/gin"
)

type User struct {
	ID       uint
	TenantID uint
	Name     string
	Email    string
	Password string
	Role     string
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// token := c.GetHeader("Authorization")

		// claims := ParseJWT(token)

		// c.Set("user_id", claims.UserID)
		// c.Set("tenant_id", claims.TenantID)

		// c.Next()
	}
}

func ParseJWT(token string) any {
	panic("unimplemented")
}


package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/o1egl/paseto/v2"
	"github.com/redis/go-redis/v9"
	"os"
	"strings"
)

// paseto section
func AuthMiddleware(redisC *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHeader := c.GetHeader("Authorization")

		if tokenHeader == "" {
			c.AbortWithStatusJSON(401, gin.H{"Error": "Unauthorized"})
			return
		}

		tokenHeader = strings.Replace(tokenHeader, "Bearer ", "", 1)

		// Parse Paseto
		var jsonToken map[string]interface{}
		footer := ""

		// decrypt
		symetricKey := os.Getenv("PASETO_SECRET")
		err := paseto.Decrypt(tokenHeader, []byte(symetricKey), &jsonToken, &footer)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"Error": "Token is invalid"})
			return
		}

		// got an email
		email, found := jsonToken["email"]
		if !found {
			c.AbortWithStatusJSON(500, gin.H{"Error": "Email not found in token"})
			return
		}

		// cv email to string
		emailStr, ok := email.(string)
		if !ok {
			c.AbortWithStatusJSON(500, gin.H{"Error": "Failed to convert email to string"})
			return
		}

		//get email on redis in hereee
		ctx := context.Background()
		tokenInRedis, err := redisC.Get(ctx, "userEmail:"+emailStr).Result()

		if err != nil || tokenInRedis != tokenHeader {
			c.AbortWithStatusJSON(401, gin.H{"Error": "Token expired or user logged out"})
			return
		}

		// Set user
		c.Set("user", emailStr)

		c.Next()
	}
}

//jwt section

//func AuthMiddleware() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		tokenHeader := c.GetHeader("Authorization")
//
//		if tokenHeader == "" {
//			c.AbortWithStatusJSON(401, gin.H{"Error": "Unauthorized"})
//			return
//		}
//
//		//replace
//		tokenHeader = strings.Replace(tokenHeader, "Bearer ", "", 1)
//
//		//parsejwt
//		secretjwt := os.Getenv("JWT_SECRET")
//		token, err := helper.ParseJWT(tokenHeader)
//		if err != nil {
//			c.AbortWithStatusJSON(500, gin.H{"Error": "Failed to parse token"})
//			return
//		}
//
//		//claims jwt
//		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
//			//set
//			c.Set("user", claims["user"])
//			return
//		}
//
//		c.Next()
//	}
//}


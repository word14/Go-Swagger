package middleware

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func JwtMiddleware(context *gin.Context) {
	authToken := context.GetHeader("Authorization")[7:]
	secret := []byte("yusuftampan")
	parsedToken, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return secret, nil
	})

	if err != nil {
		fmt.Println(err)
		context.AbortWithStatusJSON(401, gin.H{
			"message": "Unauthorized", //err.Error(),
			"success": false,
		})
	} else if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		fmt.Println(claims["foo"], claims["exp"])
		context.Next()
	}
}

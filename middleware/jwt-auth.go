package middleware

import(
	"log"
	"net/http"
	"serviceOpname-v2/service"
	"serviceOpname-v2/config/entity/helper"
	
	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
)

func AuthorizeJWT(jwtService service.JWTService) gin.HandlerFunc {
	return func(c *gin.Context){
		authHeader := c.GetHeader("Authorization")
		if authHeader == ""{
			response := helper.BuildErrorResponse("Failed to process", "No toke found", nil)
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
		token, err := jwtService.ValidateToken(authHeader)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claim[userID] : ", claims["userID"])
			log.Println("Claim[issuer] : ", claims["issuer"])
		}else{
			log.Println(err)
			response := helper.BuildErrorResponse("Token is not valid", err.Error(), nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
	}
}
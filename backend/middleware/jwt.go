package middleware

import "C"
import (
	"backend/models"
	setting "backend/settings"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"time"
)

type CustomClaims struct {
	Username string `json:"username"`
	UserId   int64  `json:"userId"`
	jwt.RegisteredClaims
}

// GenToken generates the jwt
func GenToken(username string, userId int64) (string, error) {
	expire := time.Now().Add(setting.TokenExpireDuration)

	// self-defined claim
	claims := CustomClaims{
		username,
		userId,
		jwt.RegisteredClaims{
			Issuer: "sky",
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := t.SignedString(setting.Secret)

	j := models.Jwts{Token: token, Expire: expire}
	models.DB.Create(&j)
	return token, err
}

// ParseToken parses the jwt
func ParseToken(tokenString string) (*CustomClaims, error) {
	// parse the token
	// if it's self-defined Claim struct, use method ParseWithClaims()
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return setting.Secret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

// JWTAuthMiddleware auths the jwt
func JWTAuthMiddleware(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusOK, gin.H{
			"state": "fail",
			"msg":   "Authorization in Request Header is empty",
		})
		c.Abort()
		return
	}
	mc, err := ParseToken(authHeader)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"state": "fail",
			"msg":   "ineffective token",
		})
		c.Abort()
		return
	}
	var jwts models.Jwts
	models.DB.Where("token = ?", authHeader).First(&jwts)
	if jwts.Token != "" {
		if jwts.Expire.After(time.Now()) {
			jwts.Expire = time.Now().Add(setting.TokenExpireDuration)
			models.DB.Save(&jwts)
		} else {
			// forcefully delete the jwt in the datatable
			models.DB.Unscoped().Delete(&jwts)
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"state": "fail",
			"msg":   "ineffective token",
		})
		c.Abort()
		return
	}

	c.Set("username", mc.Username)
	c.Set("uid", mc.UserId)
	c.Next()
}

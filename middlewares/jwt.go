package middlewares

import (

	"errors"

	jwt "github.com/dgrijalva/jwt-go"

	"github.com/gin-gonic/gin"

	"github.com/nj-jay/httpServer/models"

	"net/http"

	"strings"

	"time"

)

type MyClaims struct {

	models.Login

	jwt.StandardClaims

}

const  (

	TokenExpireDuration = time.Hour * 2

)

var (

	MySecret = []byte("cloudcoder.com")

)

func GenToken(username, password string) (string, error) {


	mc := MyClaims{

		Login: models.Login{Username:username, Password:password},

		StandardClaims: jwt.StandardClaims{

			ExpiresAt:time.Now().Add(TokenExpireDuration).Unix(),

			Issuer: "nj-jay",

		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, mc)

	return token.SignedString(MySecret)

}

func ParseToken(tokenString string) (*MyClaims, error) {

	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {

		return MySecret, nil

	})

	if err != nil {

		return nil, err

	}

	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid { // 校验token

		return claims, nil
	}

	return nil, errors.New("invalid token")

}

func JWTAuthMiddleware() gin.HandlerFunc {


	return func(c *gin.Context) {

		authHeader := c.Request.Header.Get("Authorization")

		if authHeader == "" {

			c.IndentedJSON(http.StatusOK, gin.H {

				"code" : 2003,

				"meg" : "请求头中auth为空",

			})

			c.Abort()

			return

		}

		parts := strings.SplitN(authHeader, " ", 2)

		if !(len(parts) == 2 && parts[0] == "Bearer") {

			c.JSON(http.StatusOK, gin.H{

				"code": 2004,

				"msg":  "请求头中auth格式有误",

			})

			c.Abort()

			return

		}

		mc, err := ParseToken(parts[1])

		if err != nil {

			c.JSON(http.StatusOK, gin.H{

				"code": 2005,

				"msg":  "无效的Token",

			})

			c.Abort()

			return

		}

		c.Set("username", mc.Username)

		c.Set("password", mc.Password)

		c.Next()

	}

}
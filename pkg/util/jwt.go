package util

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"seckill-server/pkg/app"
	"seckill-server/pkg/consts"
	"time"
)

var jwtSecret []byte

type Claims struct {
	PkId     string `json:"pk_id"`
	UserName string `json:"user_name"`
	jwt.StandardClaims
}

// GenerateToken generate tokens used for auth
func GenerateToken(pkId string, userName string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Duration(60) * time.Minute) // 1一个小时的过期时间

	claims := Claims{
		pkId,
		userName,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "quick-pass",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	//err = token_cache.SetCacheToken(agency, username, token)
	if err != nil {
		log.Println("token_cache.SetCacheToken:", err)
	}
	return token, err
}

// ParseToken parsing token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

// 载荷，可以加一些自己需要的信息
type CustomClaims struct {
	ID          int64  `json:"userId"`
	AccountName string `json:"accountName"`
	RoleId      int    `json:"role_id"`
	RoleName    string `json:"role_name"`
	jwt.StandardClaims
}

// JWT is jwt middleware
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, code := checkToken(c)
		if code == http.StatusOK {
			// 继续交由下一个路由处理,并将解析出的信息传递下去
			c.Set(consts.TokenKey, claims)
			return
		}

		switch code {
		case http.StatusBadRequest:
			app.UnauthorizedResp(c, code, "token不能为空！")
			c.Abort()
		case http.StatusFound:
			app.UnauthorizedResp(c, code, "token过期！")
			c.Abort()
		default:
			app.UnauthorizedResp(c, code, "token解析错误！")
			c.Abort()
		}
	}
}

func checkToken(c *gin.Context) (*Claims, int) {
	token := c.GetHeader(consts.HeaderToken)
	if token == "" {
		return nil, http.StatusBadRequest
	}

	claims, err := ParseToken(token)
	if err != nil {
		switch err.(*jwt.ValidationError).Errors {
		case jwt.ValidationErrorExpired:
			return nil, http.StatusFound
		default:
			return nil, http.StatusUnauthorized
		}
	}

	return claims, http.StatusOK
}

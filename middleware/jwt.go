package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-cli/config"
	"github.com/gin-cli/pkg/app"
	"github.com/gin-cli/pkg/e"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"time"
)

const (
	ActiveUser = "Claims"
)

//JWT as a middleware to validate user's authority.
func JWT() gin.HandlerFunc {
	return func(context *gin.Context) {
		var code int
		rep := app.NewRep(context)
		token := context.Request.Header.Get(config.Config.JWT.HeaderKey)
		if token == "" {
			rep.Ok(e.ERROR_AUTH_NOT_FOUND)
			context.Abort()
			return
		}
		claims, code := parseToken(token)
		if code != e.SUCCESS {
			rep.Ok(code)
			context.Abort()
			return
		}
		context.Set(ActiveUser, claims)
		context.Next()
	}
}

type Claims interface {
	GenerateToken() (string, error)
}

type CustomClaims struct {
	jwt.StandardClaims
	UUID     uuid.UUID
	NickName string
	RoleId   string
}

func NewClaims(id uuid.UUID, name, roleId string) Claims {
	expire := time.Now().Add(time.Hour * config.Config.JWT.ExpireTime).Unix()
	return &CustomClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expire,
			Issuer:    config.Config.JWT.Issuer,
			IssuedAt:  time.Now().Unix(),
		},
		UUID:     id,
		NickName: name,
		RoleId:   roleId,
	}
}

func (c *CustomClaims) GenerateToken() (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return t.SignedString([]byte(config.Config.JWT.Secret))
}

func parseToken(token string) (*CustomClaims, int) {
	var code int
	t, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Config.JWT.Secret), nil
	})
	if err != nil {
		code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
	}

	if claims, ok := t.Claims.(*CustomClaims); ok && t.Valid {
		if !isExpire(claims.ExpiresAt) {
			if claims.Issuer == config.Config.JWT.Issuer {
				return claims, e.SUCCESS
			} else {
				code = e.ERROR_AUTH
			}
		} else {
			code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
		}
	}
	return nil, code
}

func isExpire(expire int64) bool {
	if expire - time.Now().Unix() < 0 {
		return true
	}
	return false
}
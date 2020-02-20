package middle

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/vtcc/voice-note-admin/router/base"
	"strings"
)

func Auth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			ck, err := c.Cookie("token")
			if err == nil {
				token, _ := jwt.Parse(ck.Value, func(token *jwt.Token) (i interface{}, err error) {
					return []byte(base.JWT_HASH_KEY), err
				})

				if token.Valid {
					claims := token.Claims.(jwt.MapClaims)
					c.Set("uid", claims["jti"])
					return next(c)
				}
			}

			if strings.Contains(c.Path(), "/login") || strings.Contains(c.Path(), "/static/*") {
				return next(c)
			}

			return c.Redirect(302, "/login")
		}
	}
}


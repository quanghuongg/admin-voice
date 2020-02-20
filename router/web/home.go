package web

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/vtcc/voice-note-admin/model"
	"github.com/vtcc/voice-note-admin/router/base"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
	"time"
)

func Home(c echo.Context) error {

	data := base.ViewData{
		RouteName: "home/index",
	}

	return c.Render(http.StatusOK, "home/index", data)
}


func Login(c echo.Context) error {

	data := base.ViewData{
		RouteName: "home/login",
	}

	if c.Get("uid") != nil {
		c.Redirect(302, "/")
	}

	return c.Render(http.StatusOK, "login", data)
}

func Logout(c echo.Context) error {

	ck := http.Cookie{
		HttpOnly: true,
		Name: "token",
		MaxAge: -1,
		Expires: time.Now().Truncate(100),
	}

	c.SetCookie(&ck)

	return c.Redirect(302, "/login")
}

func LoginPost(c echo.Context) error {

	email := c.FormValue("email")
	password := c.FormValue("password")

	var user model.User
	u, err := user.GetUserByPhoneOrEmail(email)
	if err != nil {
		return c.Redirect(302, "/login")
	}

	if u.RoleId == 1 && bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)) == nil {
		token := jwt.New(jwt.SigningMethodHS256)
		token.Claims = jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour + 100).Unix(),
			Id: strconv.Itoa(u.Id),
			Subject: u.FullName,
		}

		key, err := token.SignedString([]byte(base.JWT_HASH_KEY))
		if err == nil {
			ck := http.Cookie{
				HttpOnly: true,
				Name: "token",
				Value: key,
			}

			c.SetCookie(&ck)

			return c.Redirect(302, "/")
		}
	}

	return c.Redirect(302, "/login")
}

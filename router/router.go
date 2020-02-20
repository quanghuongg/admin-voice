package routers

import (
	"github.com/labstack/echo"
	"github.com/vtcc/voice-note-admin/config"
	"github.com/vtcc/voice-note-admin/router/api"
	"github.com/vtcc/voice-note-admin/router/web"
	"html/template"
	"io"
)

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	//if viewContext, isMap := data.(map[string]interface{}); isMap {
	//	viewContext["reverse"] = c.Echo().Reverse
	//}

	if !config.IsProd() {
		t.templates, _ = template.ParseGlob("views/**/*")
	}
	return t.templates.ExecuteTemplate(w, name + ".html", data)
}

func Register(e *echo.Echo) {

	e.Debug = true
	tp, _ := template.ParseGlob("views/*/*")
	t := &TemplateRenderer{tp}
	e.Renderer = t

	e.Static("/static", config.AppConfig().StaticPath)
	if config.IsProd() {
		e.HTTPErrorHandler = HTTPErrorHandler
	}

	// for web
	e.GET("/", web.Home)
	e.GET("/login", web.Login)
	e.POST("/loginPost", web.LoginPost)
	e.GET("/logout", web.Logout)
	e.GET("/stats/record", web.Records)
	e.GET("/stats/user", web.Users)

	// for api
	e.GET("/api/stats/record", api.RecordStats)
	e.GET("/api/stats/user", api.UserStats)
	e.GET("/api/record/detail", api.RecordDetail)
}

func HTTPErrorHandler(err error, c echo.Context) {
	c.Redirect(302, "/")
}
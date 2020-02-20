package web

import (
	"github.com/labstack/echo"
	"github.com/vtcc/voice-note-admin/config"
	"github.com/vtcc/voice-note-admin/router/base"
	"net/http"
)

func Records(c echo.Context) error {

	data := base.ViewData{
		RouteName:    "stats/record",
		VoiceNoteUrl: config.AppConfig().VoiceNoteUrl,
	}

	return c.Render(http.StatusOK, "stats/record", data)
}

func Users(c echo.Context) error {

	data := base.ViewData{
		RouteName:    "stats/user",
		VoiceNoteUrl: config.AppConfig().VoiceNoteUrl,
	}

	return c.Render(http.StatusOK, "stats/user", data)
}

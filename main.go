package main

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/vtcc/voice-note-admin/config"
	_ "github.com/vtcc/voice-note-admin/model"
	"github.com/vtcc/voice-note-admin/router"
	"github.com/vtcc/voice-note-admin/router/middle"
	"os"
	"os/signal"
	"time"
)


func main() {

	e := echo.New()
	routers.Register(e)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Secure())
	e.Use(middle.Auth())


	go func() {
		if err := e.Start(config.AppConfig().Listen); err != nil {
			fmt.Println("Error start server!")
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		fmt.Println("Error Fatal!")
	}
}

package main

import (
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/nekomeowww/bo"
)

func NewServer1() *http.Server {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	return &http.Server{
		Addr:        ":8080",
		Handler:     e,
		ReadTimeout: time.Second * 10,
	}
}

func NewServer2() *http.Server {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Hello, World!",
		})
	})

	return &http.Server{
		Addr:        ":8081",
		Handler:     e,
		ReadTimeout: time.Second * 10,
	}
}

func main() {
	app := bo.New()

	app.Add(func(ctx context.Context, lifeCycle bo.LifeCycle) error {
		srv := NewServer1()

		lifeCycle.Append(bo.Hook{
			OnStart: func(ctx context.Context) error {
				return srv.ListenAndServe()
			},
			OnStop: func(ctx context.Context) error {
				return srv.Shutdown(ctx)
			},
		})

		return nil
	})

	app.Add(func(ctx context.Context, lifeCycle bo.LifeCycle) error {
		srv := NewServer2()

		lifeCycle.Append(bo.Hook{
			OnStart: func(ctx context.Context) error {
				return srv.ListenAndServe()
			},
			OnStop: func(ctx context.Context) error {
				return srv.Shutdown(ctx)
			},
		})

		return nil
	})

	app.Start()
}

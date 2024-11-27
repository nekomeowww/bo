# `bo`

[![Go Reference](https://pkg.go.dev/badge/github.com/nekomeowww/bo.svg)](https://pkg.go.dev/github.com/nekomeowww/bo)
![](https://github.com/nekomeowww/bo/actions/workflows/ci.yaml/badge.svg)
[![](https://goreportcard.com/badge/github.com/nekomeowww/bo)](https://goreportcard.com/report/github.com/nekomeowww/bo)

🚀 bo, BootKit for easily bootstrapping multi-goroutine applications, CLIs, can be used as drop-in replacement of `uber/fx`

- Forget about `signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)`
- Forget about `context.WithCancel(context.Background())`
- Forget about error passing through multiple layers

```go
bo.New()
bo.Add(YourService())
bo.Add(AnotherService())
bo.Start()
```

That's it!

## Getting Started

```shell
go get github.com/nekomeowww/bo
```

## Usage

```go
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
```

## 🤠 Spec

GoDoc: [https://godoc.org/github.com/nekomeowww/bo](https://godoc.org/github.com/nekomeowww/bo)

## 👪 Other family members of `anyo`

- [nekomeowww/xo](https://github.com/nekomeowww/xo): Mega utility & helper & extension library for Go
- [nekomeowww/fo](https://github.com/nekomeowww/fo): Functional programming utility library for Go
- [nekomeowww/tgo](https://github.com/nekomeowww/tgo): Telegram bot framework for Go
- [nekomeowww/wso](https://github.com/nekomeowww/wso): WebSocket utility library for Go

## 🎆 Other cool related Golang projects I made & maintained

- [Kollama - Ollama Operator](https://github.com/knoway-dev/knoway): Kubernetes Operator for managing Ollama instances across multiple clusters
- [lingticio/llmg](https://github.com/lingticio/llmg): LLM Gateway with gRPC, WebSocket, and RESTful API adapters included.
- [knoway-dev/knoway](https://github.com/knoway-dev/knoway): Kubernetes-first LLM Gateway

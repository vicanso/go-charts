package main

import (
	"bytes"
	"embed"
	"time"

	"github.com/vicanso/elton"
	"github.com/vicanso/elton/middleware"
	"github.com/vicanso/go-charts"
)

//go:embed web
var webFS embed.FS

func main() {
	e := elton.New()

	e.Use(middleware.NewDefaultError())
	e.Use(middleware.NewDefaultBodyParser())

	assetFS := middleware.NewEmbedStaticFS(webFS, "web")
	e.GET("/static/*", middleware.NewStaticServe(assetFS, middleware.StaticServeConfig{
		// 缓存服务器缓存一个小时
		SMaxAge:             5 * time.Minute,
		DenyQueryString:     true,
		DisableLastModified: true,
		EnableStrongETag:    true,
	}))

	e.GET("/", func(c *elton.Context) error {
		buf, err := webFS.ReadFile("web/index.html")
		if err != nil {
			return err
		}
		c.SetContentTypeByExt(".html")
		c.BodyBuffer = bytes.NewBuffer(buf)
		return nil
	})
	e.POST("/", func(c *elton.Context) error {
		buf, err := charts.RenderEChartsToSVG(string(c.RequestBody))
		if err != nil {
			return err
		}
		c.BodyBuffer = bytes.NewBuffer(buf)
		return nil
	})

	err := e.ListenAndServe(":3000")
	if err != nil {
		panic(err)
	}
}

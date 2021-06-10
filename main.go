package main

import (
	"embed"
	"github.com/gin-gonic/gin"
	"io/fs"
	"net/http"
)

//go:embed web/dist
var dist embed.FS

func setupRouter() *gin.Engine {
	r := gin.Default()

	assets, err := fs.Sub(dist, "web/dist/assets")
	if err != nil {
		panic(err)
	}
	r.StaticFS("/assets/", http.FS(assets))

	index, err := fs.Sub(dist,"web/dist")
	if err != nil {
		panic(err)
	}
	r.GET("/", func(c *gin.Context) {
		c.FileFromFS("/", http.FS(index))
	})

	return r
}

func main() {
	r := setupRouter()

	r.Run(":8080")
}

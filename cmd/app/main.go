package main

import (
	"github.com/feitian124/go-stater/web"
	"github.com/gin-gonic/gin"
	"io/fs"
	"net/http"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	assets, err := fs.Sub(web.Assets, "dist/assets")
	if err != nil {
		panic(err)
	}
	r.StaticFS("/assets/", http.FS(assets))

	index, err := fs.Sub(web.Assets,"dist")
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

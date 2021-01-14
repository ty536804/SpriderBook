package Router

import (
	"Book/Controller/Book"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func InitRouter() *gin.Engine {
	gin.DisableConsoleColor()

	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdin)

	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	r.GET("/menus", Book.Index)
	return r
}

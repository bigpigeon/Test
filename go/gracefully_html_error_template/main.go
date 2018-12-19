package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	g := gin.Default()
	g.LoadHTMLGlob("templates/*")
	g.Use(func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				//c.HTML(http.StatusInternalServerError, "error.tmpl", nil)
				for i := 0; i < 3; i++ {
					time.Sleep(1 * time.Second)
					c.Writer.Write([]byte("1\n"))
				}
			}
		}()

		c.Next()
	})
	g.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{"var": 4})
	})
	g.Run(":18001")
}

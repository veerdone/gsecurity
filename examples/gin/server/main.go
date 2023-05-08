package main

import (
	"github.com/gin-gonic/gin"
	"github.com/veerdone/gsecurity"
	"github.com/veerdone/gsecurity/adaptor/ginadaptor"
	"log"
	"strconv"
)

func main() {
	g := gin.Default()

	g.POST("/login", func(c *gin.Context) {
		queryId := c.Query("id")
		id, _ := strconv.ParseInt(queryId, 10, 64)
		gsecurity.LoginAndSet(id, ginadaptor.New(c))
		c.JSON(200, gin.H{"msg": "ok"})
	})

	g.GET("/isLogin", func(c *gin.Context) {
		isLogin := gsecurity.IsLogin(ginadaptor.New(c))

		c.JSON(200, gin.H{"isLogin": isLogin})
	})

	err := g.Run()
	if err != nil {
		log.Fatalln(err)
	}
}

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rjmateus/suma-updater/web"
)

func main() {
	r := gin.Default()
	web.InitRoutes(r)
	r.Run(":8088")

	// https://hoohoo.top/blog/20210530112304-golang-tutorial-introduction-gin-html-template-and-how-integration-with-bootstrap/

}

package main

import (
	// "net/http"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {

	/* Set the router as the default one provided by Gin */
	router = gin.Default()

	/*
		Process the templates at the start so that they don't have to be loaded from the disk again.
		this makes serving HTML pages very fast
	*/
	router.LoadHTMLGlob("../template/*")

	router.GET("/", showIndexPage)
	router.GET("/article/view/:article_id", getArticle)
	router.Run()
}

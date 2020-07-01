package router

import "github.com/gin-gonic/gin"

func New() *gin.Engine {
	router := gin.Default()
	router.GET("/info/:slug", GetUrlInfo)
	router.GET("/url/:slug", GetShortenedUrl)
	

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	return router
}
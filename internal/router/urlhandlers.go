package router

import (
	"context"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tludlow/URL-GO/internal/database"
)

type Url struct {
	id int
	slug string
	link string
	expires time.Time
}

///v1/:slug
func GetShortenedUrl(c *gin.Context) {
	slug := c.Param("slug")

	//Get the value of the link based on the slug from the db
	var link string
	var expires time.Time
	err := database.DB.QueryRow(context.Background(), "SELECT link, expires FROM url WHERE slug = $1", slug).Scan(&link, &expires)
	if err != nil {
		log.Printf("error getting slug info from db - %v", err.Error())
		c.JSON(200, gin.H{
			"error": "Error getting link from slug",
		})
		return
	}

	c.JSON(200, gin.H{
		"redirect": link,
		"expires": expires.String(),
	})
	// c.Redirect(http.StatusMovedPermanently, link)
}

///v1/info/:infoSlug
func GetUrlInfo(c *gin.Context) {
	slug := c.Param("slug")

	var url Url
	err := database.DB.QueryRow(context.Background(), "SELECT * FROM url WHERE slug = $1", slug).Scan(&url.id, &url.slug, &url.link, &url.expires)
	if err != nil {
		log.Printf("error getting slug info from db - %v", err.Error())
		c.JSON(200, gin.H{
			"error": "Error getting info about " + slug,
		})
		return
	}

	_, _ = database.DB.Exec(context.Background(), "UPDATE url SET expires = (expires + '01:00:00'::interval) WHERE slug = $1", slug)

	c.JSON(200, gin.H{
		"id": url.id,
		"slug": url.slug,
		"link": url.link,
		"expires": url.expires,
	})
}
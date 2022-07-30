package handler

import (
	"net/http"
	"news/platform/newsfeed"

	"github.com/gin-gonic/gin"
)

func NewsfeedGet(feed newsfeed.Getter) gin.HandlerFunc {
	return func(c *gin.Context) {
		result := feed.GetAll()

		c.JSON(http.StatusOK, result)
	}
}

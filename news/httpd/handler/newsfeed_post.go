package handler

import (
	"net/http"
	"news/platform/newsfeed"

	"github.com/gin-gonic/gin"
)

type newsfeedPost struct {
	Title string `json:"title"`
	Post  string `json:"post"`
}

func NewsfeedPost(feed newsfeed.Adder) gin.HandlerFunc {
	return func(c *gin.Context) {

		requestBody := newsfeedPost{}
		c.Bind(&requestBody)

		item := newsfeed.Item{
			Title: requestBody.Title,
			Post:  requestBody.Post,
		}
		feed.Add(item)

		c.Status(http.StatusNoContent)

		// c.JSON(http.StatusOK, map[string]string{
		// 	"hello": "mom",
		// })
	}
}

package Posts

import (
	"net/http"
	Post "redditClone/pkg/common/models/posts"

	"github.com/gin-gonic/gin"
)

type UpdatePostRequestBody struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Flairs  string `json:"flairs"`
}

func (h handler) UpdatePost(c *gin.Context) {
	id := c.Param("id")
	body := UpdatePostRequestBody{}

	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	Post := Post.New()
	Post.UUID = id

	if result := h.DB.First(&Post); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	if body.Content != "" {
		Post.Content = body.Content
	}

	if body.Flairs != "" {
		Post.Flairs = body.Flairs
	}

	if body.Title != "" {
		Post.Title = body.Title
	}

	h.DB.Save(&Post)

	c.JSON(http.StatusOK, &Post)
}

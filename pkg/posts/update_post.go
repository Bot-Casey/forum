package Posts

import (
	"net/http"

	"redditClone/pkg/common/models"

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

	var Post models.Post

	if result := h.DB.First(&Post, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	Post.Content = body.Content
	Post.Flairs = body.Flairs
	Post.Title = body.Title

	h.DB.Save(&Post)

	c.JSON(http.StatusOK, &Post)
}

package Posts

import (
	"net/http"

	"redditClone/pkg/common/models"

	"github.com/gin-gonic/gin"
)

func (h handler) GetPost(c *gin.Context) {
	var Post models.Post
	Post.UUID = c.Param("id")

	result := h.DB.First(&Post)

	if result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)

		return
	}

	c.JSON(http.StatusOK, &Post)
}

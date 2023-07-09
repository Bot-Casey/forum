package Posts

import (
	"net/http"

	"redditClone/pkg/common/models"

	"github.com/gin-gonic/gin"
)

func (h handler) DeletePost(c *gin.Context) {
	id := c.Param("id")

	var Post models.Post

	if result := h.DB.First(&Post, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	h.DB.Delete(&Post)

	c.Status(http.StatusOK)
}

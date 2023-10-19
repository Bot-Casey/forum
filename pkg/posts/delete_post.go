package Posts

import (
	"net/http"
	Post "redditClone/pkg/common/models/posts"

	"github.com/gin-gonic/gin"
)

func (h handler) DeletePost(c *gin.Context) {
	id := c.Param("id")

	Post := Post.New()

	if result := h.DB.First(Post, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	h.DB.Delete(&Post)

	c.Status(http.StatusOK)
	return
}

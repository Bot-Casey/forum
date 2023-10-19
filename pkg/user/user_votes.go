package 

import (
	"net/http"

	"redditClone/pkg/common/models/votes"
	"redditClone/pkg/common/models/user"

	"github.com/gin-gonic/gin"
)

func (h handler) GetPost(c *gin.Context) {
	var User models.User

	var Post models.Votes
	Post.UUID = c.Param("id")

	result := h.DB.First(&Post)

	if result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)

		return
	}

	c.JSON(http.StatusOK, &Post)
}

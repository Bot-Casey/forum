package Comments

import (
	"net/http"
	"redditClone/pkg/common/models"

	"github.com/gin-gonic/gin"
)

func (h handler) GetPost(c *gin.Context) {
	var Comment models.Comment
	Comment.UUID = c.Param("id")

	result := h.DB.First(&Comment)

	if result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)

		return
	}

	c.JSON(http.StatusOK, &Comment)
}

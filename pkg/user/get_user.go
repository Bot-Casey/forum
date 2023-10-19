package User

import (
	"net/http"
	User "redditClone/pkg/common/models/users"

	"github.com/gin-gonic/gin"
)

func (h handler) GetUser(c *gin.Context) {
	User := User.New()
	User.UUID = c.Param("id")

	result := h.DB.First(&User)

	if result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)

		return
	}

	c.JSON(http.StatusOK, &User)
}

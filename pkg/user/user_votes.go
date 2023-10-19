package User

import (
	"errors"
	"net/http"
	"strconv"

	User "redditClone/pkg/common/models/users"

	"github.com/gin-gonic/gin"
)

func (h handler) GetUserVotes(c *gin.Context) {
	limit := 0

	if query, err := c.GetQuery("limit"); !err {
		c.AbortWithError(http.StatusBadRequest, errors.New("limit not set"))
		return
	} else if limit, err := strconv.Atoi(query); err != nil {
		c.AbortWithError(http.StatusBadRequest, errors.New("limit not valid"))
		return
	} else if limit > 100 {
		c.AbortWithError(http.StatusBadRequest, errors.New("limit greater than 100"))
		return
	}

	h.DB.Limit(limit)

	User := User.New()
	User.UUID = c.Param("id")

	result := h.DB.First(&Post)

	if result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)

		return
	}

	c.JSON(http.StatusOK, &Post)
}

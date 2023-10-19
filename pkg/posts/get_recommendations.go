package Posts

import (
	"errors"
	"net/http"
	Post "redditClone/pkg/common/models/posts"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h handler) GetRecommendations(c *gin.Context) {
	result := h.DB
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

	result.Limit(limit)

	if query, err := c.GetQuery("filter"); !err {
		switch query {
		case "hot":
			result.Scopes(CreatedWithinPastDay, SortByUpVotes)
		case "votes":
			result.Order("vote_count")
		case "new":
			result.Where("? - ")
		default:
			c.AbortWithError(
				http.StatusBadRequest,
				errors.New("filter not valid (hot|votes|new)"),
			)
			return
		}
	}

	Posts := Post.Posts{}

	result.Find(&Posts)

	if result.Error != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"items": &Posts,
	})
	return
}

package Votes

import (
	"errors"
	"net/http"

	"redditClone/pkg/common/models"

	"github.com/gin-gonic/gin"
)

type enum_resource_type string

func (r enum_resource_type) String() (string, error) {
	const (
		post    enum_resource_type = "post"
		comment enum_resource_type = "comment"
	)

	switch r {
	case post:
		return "post", nil
	case comment:
		return "comment", nil
	}
	return "", errors.New("invalid resource_type")
}

type AddVoteRequestBody struct {
	Resource      string             `json:"resource"`
	Resource_Type enum_resource_type `json:"Resource_type"`
	User          string             `json:"User"`
}

func (h handler) AddVote(c *gin.Context) {
	body := AddVoteRequestBody{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	resource_type, err := enum_resource_type.String(body.Resource_Type)
	if err != nil {
		c.JSON(http.StatusBadRequest, "{error: invalid resource_type}")
		c.AbortWithError(http.StatusBadRequest, errors.New("invalid resource_type"))
		return
	}

	var Vote models.Vote_Relations

	Vote.Resource = body.Resource
	Vote.User = body.User
	Vote.Resource_Type = resource_type

	var Post models.Post
	Post.UUID = Vote.Resource

	if result := h.DB.First(&Post); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	Post.IncrementVote()

	if result := h.DB.Save(&Post); result.Error != nil {
		c.AbortWithError(http.StatusConflict, result.Error)
		return
	}

	if result := h.DB.Create(&Vote); result.Error != nil {
		c.AbortWithError(http.StatusInternalServerError, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &Vote)
}

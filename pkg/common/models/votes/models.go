package votes

import (
	"errors"
	"time"
)

type Vote_Relations struct {
	Resource      string    `json:"resource" gorm:"primaryKey"`
	Resource_Type string    `json:"resource_type"`
	User          string    `json:"user" gorm:"primaryKey"`
	CreatedAt     time.Time `json:"created_at"`
}

func New() Vote_Relations {
	return Vote_Relations{}
}

type Votes []Vote_Relations

type Enum_resource_type string

func (r Enum_resource_type) String() (string, error) {
	const (
		post    Enum_resource_type = "post"
		comment Enum_resource_type = "comment"
	)

	switch r {
	case post:
		return "post", nil
	case comment:
		return "comment", nil
	}
	return "", errors.New("invalid resource_type")
}

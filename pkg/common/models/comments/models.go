package comments

import (
	"time"

	"github.com/google/uuid"
)

type Comment struct {
	UUID       string    `json:"uuid" gorm:"primaryKey"`
	Parent     string    `json:"parent"`
	ParentType string    `json:"parent_type"`
	VoteCount  int       `json:"votes"`
	Flairs     string    `json:"flairs"`
	Author     string    `json:"author"`
	UpdatedAt  time.Time `json:"updated_at"`
	CreatedAt  time.Time `json:"created_at"`
}

func New() Comment {
	C := Comment{}

	C.UUID = uuid.New().String()
	C.VoteCount = 1

	return C
}

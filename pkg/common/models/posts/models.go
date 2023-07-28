package posts

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	UUID      string     `json:"uuid" gorm:"primaryKey"`
	Title     string     `json:"title"`
	Author    string     `json:"author"`
	Content   string     `json:"content"`
	VoteCount int        `json:"votes"`
	Flairs    string     `json:"flairs"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func New() Post {
	return Post{
		UUID:      uuid.New().String(),
		VoteCount: 1,
	}
}

func (p *Post) IncrementVote() *Post {
	p.VoteCount += 1

	return p
}

func (p *Post) DecrementVote() *Post {
	p.VoteCount -= 1

	return p
}

type Vote_Relations struct {
	Resource      string    `json:"resource" gorm:"primaryKey"`
	Resource_Type string    `json:"resource_type"`
	User          string    `json:"user" gorm:"primaryKey"`
	CreatedAt     time.Time `json:"created_at"`
}

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

func (C Comment) New() Comment {
	C.UUID = uuid.New().String()
	C.VoteCount = 1

	return C
}

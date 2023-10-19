package user

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UUID       string     `json:"uuid" gorm:"primaryKey"`
	Name       string     `json:"title"`
	Email      string     `json:"email"`
	TotalKarma string     `json:"author"`
	Profile    string     `json:"content"`
	Tagline    string     `json:"tagline"`
	CreatedAt  *time.Time `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at"`
}

func New() User {
	U := User{}
	U.UUID = uuid.New().String()
	return U
}

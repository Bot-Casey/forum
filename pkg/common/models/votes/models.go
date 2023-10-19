package votes

import "time"

type Vote_Relations struct {
	Resource      string    `json:"resource" gorm:"primaryKey"`
	Resource_Type string    `json:"resource_type"`
	User          string    `json:"user" gorm:"primaryKey"`
	CreatedAt     time.Time `json:"created_at"`
}

func New() Vote_Relations {
	return Vote_Relations{}
}

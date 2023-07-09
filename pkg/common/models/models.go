package models

import "github.com/google/uuid"

type Post struct {
	UUID      string `json:"uuid" gorm:"primaryKey"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Content   string `json:"content"`
	VoteCount int    `json:"votes"`
	Flairs    string `json:"flairs"`
}

func (p *Post) Builder() *Post {
	return &Post{
		UUID:      uuid.New().String(),
		VoteCount: 1,
	}
}

func (p *Post) SetUuid(uuid string) *Post {
	p.UUID = uuid

	return p
}

func (p *Post) SetTitle(title string) *Post {
	p.Title = title

	return p
}

func (p *Post) SetAuthor(author string) *Post {
	p.Author = author

	return p
}

func (p *Post) SetContent(content string) *Post {
	p.Content = content

	return p
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
	Resource      string `json:"resource" gorm:"primaryKey"`
	Resource_Type string `json:"resource_type"`
	User          string `json:"user" gorm:"primaryKey"`
}

type Comment struct {
	UUID      string `json:"uuid" gorm:"primaryKey"`
	Parent    string `json:"parent"`
	Post      string `json:"post"`
	VoteCount int    `json:"votes"`
	Flairs    string `json:"flairs"`
	Author    string `json:"author"`
}

func (C Comment) New() Comment {
	C.UUID = uuid.New().String()
	C.VoteCount = 1

	return C
}

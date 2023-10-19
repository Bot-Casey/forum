package db

import (
	"log"

	Comment "redditClone/pkg/common/models/comments"
	Post "redditClone/pkg/common/models/posts"
	User "redditClone/pkg/common/models/users"
	Vote "redditClone/pkg/common/models/votes"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(url string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(Post.New())
	db.AutoMigrate(Comment.New())
	db.AutoMigrate(Vote.New())
	db.AutoMigrate(User.New())

	return db
}

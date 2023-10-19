package main

import (
	"redditClone/pkg/common/db"
	Posts "redditClone/pkg/posts"
	Votes "redditClone/pkg/votes"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"fmt"
)

func main() {
	viper.SetConfigFile("../pkg/common/envs/.env")
	viper.ReadInConfig()

	port := viper.Get("PORT")
	dbUrl := viper.Get("DB_URL")

	fmt.Print(port)
	fmt.Print(dbUrl)

	r := gin.Default()
	h := db.Init("postgres://test:test@localhost:5432/test")

	Posts.RegisterRoutes(r, h)
	Votes.RegisterRoutes(r, h)
	// register more routes here

	r.Run()
}

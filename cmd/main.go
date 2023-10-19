package main

import (
	"redditClone/pkg/common/db"
	Posts "redditClone/pkg/posts"
	Votes "redditClone/pkg/votes"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("../pkg/common/envs/.env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)
	dbUrl := viper.Get("DB_URL").(string)

	r := gin.Default()
	h := db.Init(dbUrl)

	Posts.RegisterRoutes(r, h)
	Votes.RegisterRoutes(r, h)
	// register more routes here

	r.Run(port)
}

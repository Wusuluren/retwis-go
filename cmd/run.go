package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wusuluren/retwis"
	"github.com/Sirupsen/logrus"
	"os"
)

func RunServer() {
	mux := gin.Default()
	mux.Any("/", retwis.IndexHandle)
	mux.Any( "/index", retwis.IndexHandle)
	mux.Any("/welcome", retwis.WelcomeHandle)
	mux.Any( "/home", retwis.HomeHandle)
	mux.Any( "/register", retwis.RegisterHandle)
	mux.Any( "/login", retwis.LoginHandle)
	mux.Any( "/post", retwis.PostHandle)
	mux.Any("/timeline", retwis.TimelineHandle)
	mux.Any("/logout", retwis.LogoutHandle)
	mux.Any("/profile", retwis.ProfileHandle)
	mux.Any( "/follow", retwis.FollowHandle)

	curPath, _ := os.Getwd()
	mux.Static("/static", curPath)

	if err := mux.Run("0.0.0.0:8000"); err != nil {
		logrus.Fatal(err)
	}
}

func main() {
	RunServer()
}

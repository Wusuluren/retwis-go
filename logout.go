package retwis

import (
	"github.com/gin-gonic/gin"
)

func LogoutHandle(c *gin.Context) {
	if !isLoggedIn(c) {
		tempRedirect(c, "index")
		return
	}

	r, _ := redisLink()
	newauthsecret := getrand()
	userid := User.Get(c, "id")
	oldauthsecret, _ := r.HGet("user:"+userid, "auth")

	r.HSet("user:"+userid, "auth", newauthsecret)
	r.HSet("auths", newauthsecret, userid)
	r.HDel("auths", oldauthsecret)

	User.Del(c)

	tempRedirect(c, "index")
}

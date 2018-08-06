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
	oldauthsecret, _ := r.Hget("user:"+userid, "auth")

	r.Hset("user:"+userid, "auth", newauthsecret)
	r.Hset("auths", newauthsecret, userid)
	r.Hdel("auths", oldauthsecret)

	User.Del(c)

	tempRedirect(c, "index")
}

package retwis

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func LogoutHandle(c *gin.Context) {
	if !isLoggedIn(c) {
		c.Redirect(http.StatusTemporaryRedirect, "index")
		return
	}

	r, _ := redisLink()
	newauthsecret := getrand()
	userid := User["id"]
	oldauthsecret, _ := r.Hget("user:"+userid, "auth")

	r.Hset("user:"+userid, "auth", newauthsecret)
	r.Hset("auths", newauthsecret, userid)
	r.Hdel("auths", oldauthsecret)

	c.Redirect(http.StatusTemporaryRedirect, "index")
}

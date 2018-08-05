package retwis

import (
	"github.com/gin-gonic/gin"
	"time"
	"net/http"
)

func LoginHandle(c *gin.Context) {
	// Form sanity checks
	if gt(c, "username") == "" || gt(c, "password") == "" {
		goback(c, "You need to enter both username and password to login.")
		return
	}

	// The form is ok, check if the username is available
	username := gt(c, "username")
	password := gt(c, "password")
	r, _ := redisLink()
	var userid string
	if user_id, err := r.Hget("users", username); err != nil || user_id == "" {
		goback(c, "Wrong username or password")
		return
	} else {
		userid = user_id
	}
	realpassword, _ := r.Hget("user:"+userid, "password")
	if realpassword != password {
		goback(c, "Wrong username or password")
		return
	}

	// Username / password OK, set the cookie and redirect to index
	authsecret, _ := r.Hget("user:"+userid, "auth")
	setcookie(c, "auth", authsecret, int(time.Now().Unix() + 3600*24*365))
	c.Redirect(http.StatusTemporaryRedirect, "index")
}

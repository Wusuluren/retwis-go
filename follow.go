package retwis

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
	"fmt"
	"net/url"
)

func FollowHandle(c *gin.Context) {
	r, _ := redisLink()
	username, _ := r.Hget("user:"+gt(c, "uid"), "username")
	if !isLoggedIn(c) || gt(c, "uid") == "" || gt(c, "f") == "" ||
		username == "" {
		c.Redirect(http.StatusTemporaryRedirect, "index")
		return
	}

	f, _ := strconv.Atoi(gt(c, "f"))
	uid := gt(c, "uid")
	if uid != User["id"] {
		if f != 0 {
			r.Zadd("followers:"+uid, time.Now().Unix(), User["id"])
			r.Zadd("following:"+User["id"], time.Now().Unix(), uid)
		} else {
			r.Zrem("followers:"+uid, User["id"])
			r.Zrem("following:"+User["id"], uid)
		}
	}
	c.Redirect(http.StatusTemporaryRedirect, fmt.Sprintf("profile?u=%s", url.PathEscape(username)))
}

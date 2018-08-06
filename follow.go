package retwis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/url"
	"strconv"
	"time"
)

func FollowHandle(c *gin.Context) {
	r, _ := redisLink()
	username, _ := r.Hget("user:"+gt(c, "uid"), "username")
	if !isLoggedIn(c) || gt(c, "uid") == "" || gt(c, "f") == "" ||
		username == "" {
		tempRedirect(c, "index")
		return
	}

	f, _ := strconv.Atoi(gt(c, "f"))
	uid := gt(c, "uid")
	userid := User.Get(c, "id")
	if uid != userid {
		if f != 0 {
			r.Zadd("followers:"+uid, time.Now().Unix(), userid)
			r.Zadd("following:"+userid, time.Now().Unix(), uid)
		} else {
			r.Zrem("followers:"+uid, userid)
			r.Zrem("following:"+userid, uid)
		}
	}
	tempRedirect(c, fmt.Sprintf("profile?u=%s", url.PathEscape(username)))
}

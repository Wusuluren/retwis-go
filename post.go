package retwis

import (
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

func PostHandle(c *gin.Context) {
	if !isLoggedIn(c) || gt(c, "status") == "" {
		tempRedirect(c, "index")
		return
	}

	r, _ := redisLink()
	userid := User.Get(c, "id")
	postid, _ := r.Incr("next_post_id")
	status := strings.Replace(gt(c, "status"), "\n", " ", -1)
	r.Hmset("post:"+postid, "user_id", userid, "time", time.Now().Unix(), "body", status)
	followers, _ := r.Zrange("followers:"+userid, 0, -1)
	followers = append(followers, userid) //??

	for _, fid := range followers {
		r.Lpush("post:"+fid.(string), postid)
	}
	// Push the post on the timeline, and trim the timeline to the
	// newest 1000 elements.
	r.Lpush("timeline", postid)
	r.Ltrim("timeline", 0, 1000)

	tempRedirect(c, "index")
}

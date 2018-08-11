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
	r.HMSet("post:"+int64ToString(postid), "user_id", userid, "time", time.Now().Unix(), "body", status)
	followers, _ := r.ZRange("followers:"+userid, 0, -1)
	followers = append(followers, userid) //??

	for _, fid := range followers {
		r.LPush("posts:"+fid, int64ToString(postid))
	}
	// Push the post on the timeline, and trim the timeline to the
	// newest 1000 elements.
	r.LPush("timeline", int64ToString(postid))
	r.LTrim("timeline", 0, 1000)

	tempRedirect(c, "index")
}

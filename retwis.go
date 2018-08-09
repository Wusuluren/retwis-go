package retwis

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/url"
	"strings"
	"time"
)

func getrand() int {
	rand.Seed(time.Now().Unix())
	data := rand.Intn(16)
	return data
}

func isLoggedIn(c *gin.Context) bool {
	if User.IsSet(c) {
		return true
	}

	if authcookie, err := c.Cookie("auth"); err == nil {
		r, _ := redisLink()
		if userid, err := r.HGet("auths", authcookie); err == nil {
			if authedcookie, err := r.HGet("user:"+userid, "auth"); err == nil {
				if authedcookie != authcookie {
					return false
				}
			}
			loadUserInfo(c, userid)
			return true
		}
	}
	return false
}

func loadUserInfo(c *gin.Context, userid string) bool {
	r, _ := redisLink()
	if !User.IsSet(c) {
		User.Init(c)
	}
	User.Set(c, "id", userid)
	if username, err := r.HGet("user:"+userid, "username"); err == nil {
		User.Set(c, "username", username)
	}
	return true
}

func redisLink() (*Redis, error) {
	r := &Redis{}
	conn, err := redis.Dial("tcp", "192.168.116.129:6379")
	r.Conn = conn
	return r, err
}

func g(c *gin.Context, param string) interface{} {
	if value, err := c.Cookie(param); err == nil {
		return value
	}
	if value := c.PostForm(param); value != "" {
		return value
	}
	if value, exists := c.Get(param); exists {
		return value
	}
	return nil
}

func gt(c *gin.Context, param string) string {
	val := g(c, param)
	if val == nil {
		return ""
	}
	if _, ok := val.(string); !ok {
		return ""
	}
	return strings.TrimSpace(val.(string))
}

func goback(c *gin.Context, msg string) {
	body := fmt.Sprintf(`<div id ="error">%s<br>
<a href="javascript:history.back()">Please return back and try again</a></div>`, msg)
	renderBody(c, body)
}

func strElapsed(t int64) string {
	var unit string
	d := time.Now().Unix() - t
	if d < 60 {
		return fmt.Sprintf("%d seconds", d)
	}
	if d < 3600 {
		m := d / 60
		if m > 1 {
			unit = "s"
		}
		return fmt.Sprintf("%d minute%s", m, unit)
	}
	d = d / (3600 * 24)
	if d > 1 {
		unit = "s"
	}
	return fmt.Sprintf("%d day%s", d, unit)
}

func parsePost(post map[string]string) string {
	var content string
	content += post["user_id"] + " "
	content += post["time"] + " "
	content += post["body"] + " "
	return content
}

func showPost(c *gin.Context, id string) bool {
	r, _ := redisLink()
	post, _ := r.HGetAll("post:" + id)
	if len(post) == 0 {
		return false
	}

	userid := post["user_id"]
	username, _ := r.HGet("user:"+userid, "username")
	elapsed := post["time"]
	userlink := fmt.Sprintf(`<a class="%s" href="profile?u=%s">%s</a>`,
		username, url.PathEscape(username), username)

	c.Writer.WriteString(fmt.Sprintf(`<div class="post">%s %s<br>`,
		userlink, parsePost(post)))
	c.Writer.WriteString(fmt.Sprintf(`<i>posted %s ago via web</i></div>`,
		elapsed))
	return true
}

func showUserPosts(c *gin.Context, userid string, start, count int) bool {
	r, _ := redisLink()
	var key string
	if userid == "-1" {
		key = "timeline"
	} else {
		key = "posts:" + userid
	}
	posts, _ := r.LRange(key, start, start+count)
	ctr := 0
	for _, p := range posts {
		if showPost(c, p) {
			ctr++
		}
		if ctr == count {
			break
		}
	}
	return len(posts) == count+1
}

func showUserPostsWithPagination(c *gin.Context, username, userid string, start, count int) {
	thispage := c.Request.URL

	next := start + 10
	prev := start - 10
	nextlink, prevlink := "", ""
	if prev < 10 {
		prev = 0
	}

	var u string
	if username != "" {
		u = "&u=" + url.PathEscape(username)
	} else {
		u = ""
	}
	if showUserPosts(c, userid, start, count) {
		nextlink = fmt.Sprintf(`<a href="%s?start=%d%s">Older posts &raquo;</a>`,
			thispage, next, u)
	}
	if start > 0 {
		var sep string
		if nextlink != "" {
			sep = " | "
		}
		prevlink = fmt.Sprintf(`<a href="%s?start=%d%s">&laquo; Newer posts</a>%s`,
			thispage, prev, u, sep)
	}
	if nextlink != "" || prevlink != "" {
		c.Writer.WriteString(fmt.Sprintf(`<div class="rightlink">%s %s</div>`, prevlink, nextlink))
	}
}

func showLastUsers(c *gin.Context) string {
	var content string
	r, _ := redisLink()
	users, _ := r.ZRevRange("users_by_time", 0, 9)
	c.Writer.WriteString("<div>")
	for _, u := range users {
		c.Writer.WriteString(fmt.Sprintf(`<a class="username" href="profile?u=%s">%s</a>`,
			url.PathEscape(u), u))
	}
	c.Writer.WriteString("</div><br>")
	return content
}

package retwis

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"strconv"
)

func ProfileHandle(c *gin.Context) {
	c.Writer.WriteString(generateHeader(c))

	r, _ := redisLink()
	userid, _ := r.Hget("users", gt(c, "u"))
	if gt(c, "u") == "" || userid == "" {
		c.Redirect(http.StatusTemporaryRedirect, "index")
		return
	}
	c.Writer.WriteString(fmt.Sprintf(`<h2 class="username">%s</h2>`, gt(c, "u")))
	if isLoggedIn(c) && User["id"] != userid {
		isfollowing, _ := r.Zscore("following:"+User["id"], userid)
		if isfollowing == 0 {
			c.Writer.WriteString(fmt.Sprintf(`<a href="follow?uid=%s&f=1" class="button">Follow this user</a>`, userid))
		} else {
			c.Writer.WriteString(fmt.Sprintf(`<a href="follow?uid=%s&f=0" class="button">Stop following</a>`, userid))
		}
	}

	var start int
	if gt(c, "start") != "" {
		start, _ = strconv.Atoi(gt(c, "start"))
	}
	showUserPostsWithPagination(c, gt(c, "u"), userid, start, 10)
	c.Writer.WriteString(generateFooter())
	renderEnd(c)
}

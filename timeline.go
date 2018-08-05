package retwis

import (
	"github.com/gin-gonic/gin"
)

func TimelineHandle(c *gin.Context) {
	c.Writer.WriteString(generateHeader(c))
	c.Writer.WriteString(`<h2>Timeline</h2>
<i>Latest registered users (an example of sorted sets)</i><br>`)
	showLastUsers(c)
	c.Writer.WriteString(`<i>Latest 50 messages from users aroud the world!</i><br>`)
	showUserPosts(c, "-1", 0, 50)
	c.Writer.WriteString(generateFooter())
	renderEnd(c)
}

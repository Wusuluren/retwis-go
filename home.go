package retwis

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
)

func HomeHandle(c *gin.Context) {
	if !isLoggedIn(c) {
		c.Redirect(http.StatusOK, "index")
		return
	}

	r, _ := redisLink()
	followers, _ := r.Zcard("followers:"+User["id"])
	following, _ := r.Zcard("following:"+User["id"])
	body := fmt.Sprintf(`<div id="postform">
<form method="POST" action="post">
%s, what you are doing?
<br>
<table>
<tr><td><textarea cols="70" rows="3" name="status"></textarea></td></tr>
<tr><td align="right"><input type="submit" name="doit" value="Update"></td></tr>
</table>
</form>
<div id="homeinfobox">
%s followers<br>
%s following<br>
</div>
</div>
`, User["username"], followers, following)
	var start int
	if gt(c, "start") == "" {
		start = 0
	} else {
		start = stringToInt(gt(c, "start"))
	}
	showUserPostsWithPagination(c, User["username"], User["id"], start, 10)
	renderBody(c, body)
}

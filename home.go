package retwis

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func HomeHandle(c *gin.Context) {
	if !isLoggedIn(c) {
		tempRedirect(c, "index")
		return
	}

	r, _ := redisLink()
	userid := User.Get(c, "id")
	followers, _ := r.Zcard("followers:" + userid)
	following, _ := r.Zcard("following:" + userid)
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
`, User.Get(c, "username"), followers, following)
	var start int
	if gt(c, "start") == "" {
		start = 0
	} else {
		start = stringToInt(gt(c, "start"))
	}
	showUserPostsWithPagination(c, User.Get(c, "username"), userid, start, 10)
	renderBody(c, body)
}

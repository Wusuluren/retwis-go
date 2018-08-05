package retwis

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func generateNavbar(c *gin.Context) string {
	var logout string
	if isLoggedIn(c) {
		logout = `<a href="logout">logout</a>`
	}
	return fmt.Sprintf(`<div id="navbar">
<a href="index">home</a>
| <a href="timeline">timeline</a>
%s
</div>`, logout)
}

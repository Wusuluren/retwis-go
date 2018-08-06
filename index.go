package retwis

import (
	"github.com/gin-gonic/gin"
)

func IndexHandle(c *gin.Context) {
	if !isLoggedIn(c) {
		tempRedirect(c, "welcome")
	} else {
		tempRedirect(c, "home")
	}
}

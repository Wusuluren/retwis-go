package retwis

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexHandle(c *gin.Context) {
	if !isLoggedIn(c) {
		c.Redirect(http.StatusTemporaryRedirect, "welcome")
	} else {
		c.Redirect(http.StatusTemporaryRedirect, "home")
	}
}
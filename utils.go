package retwis

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func renderHtml(c *gin.Context, html string) {
	c.Header("Content-Type", "text/html; charset=utf-8")
	c.String(http.StatusOK, html)
}

func renderBody(c *gin.Context, body string) {
	c.Header("Content-Type", "text/html; charset=utf-8")
	html := generateHeader(c) + body + generateFooter()
	c.String(http.StatusOK, html)
}

func renderEnd(c *gin.Context) {
	c.Header("Content-Type", "text/html; charset=utf-8")
	c.Status(http.StatusOK)
}

func intToString(i int) string {
	return int64ToString(int64(i))
}

func int64ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}

func stringToInt(s string) int {
	if i, err := strconv.Atoi(s); err == nil {
		return i
	} else {
		return 0
	}
}

func setcookie(c *gin.Context, name, value string, maxAge int) {
	c.SetCookie(name, value, maxAge, "", "", false, false)
}
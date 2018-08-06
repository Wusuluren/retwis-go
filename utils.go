package retwis

import (
	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserType map[string]map[string]string

var User UserType

func init() {
	User = make(UserType)
}

func getAuthFromContext(c *gin.Context) string {
	if auth, err := c.Cookie("auth"); err == nil {
		return auth
	}
	return ""
}

func (User UserType) Get(c *gin.Context, key string) string {
	auth := getAuthFromContext(c)
	if auth != "" {
		if user, existed := User[auth]; existed {
			return user[key]
		}
	}
	return ""
}

func (User UserType) Set(c *gin.Context, key, val string) {
	auth := getAuthFromContext(c)
	if auth != "" {
		User[auth][key] = val
	}
}

func (User UserType) Del(c *gin.Context) {
	auth := getAuthFromContext(c)
	if auth != "" {
		delete(User, auth)
	}
}

func (User UserType) IsSet(c *gin.Context) bool {
	auth := getAuthFromContext(c)
	if auth != "" {
		if _, existed := User[auth]; existed {
			return true
		}
	}
	return false
}

func (User UserType) Init(c *gin.Context) {
	auth := getAuthFromContext(c)
	if auth != "" {
		if _, existed := User[auth]; !existed {
			User[auth] = make(map[string]string)
		}
	}
}

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

func UserMiddleware(c *gin.Context) {
	logrus.Info("UserMiddleware", c)
}

func tempRedirect(c *gin.Context, location string) {
	c.Redirect(http.StatusTemporaryRedirect, location)
}

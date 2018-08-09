package retwis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/url"
	"time"
)

func RegisterHandle(c *gin.Context) {
	// Form sanity checks
	if gt(c, "username") == "" || gt(c, "password") == "" || gt(c, "password2") == "" {
		goback(c, "Every field of the registration form is needed!")
		return
	}

	// The form is ok, check if the username is available
	username := gt(c, "username")
	password := gt(c, "password")
	r, _ := redisLink()
	if val, err := r.HGet("users", username); err == nil && val != "" {
		goback(c, "Sorry the selected username is already in use.")
		return
	}

	// Everything is ok, Register the user!
	userid, _ := r.Incr("next_user_id")
	authsecret := getrand()
	r.HSet("users", username, userid)
	r.HMSet("user:"+int64ToString(userid),
		"username", username,
		"password", password,
		"auth", authsecret)
	r.HSet("auths", authsecret, userid)

	r.ZAdd("users_by_time", time.Now().Unix(), username)

	// User registered! Login her / him.
	setcookie(c, "auth", intToString(authsecret), int(time.Now().Unix()+3600*24*365))

	body := fmt.Sprintf(`
<h2>Welcome aboard!</h2>
Hey %s, now you have an account, <a href="index">a good start is to write your first message!</a>.
`,
		url.PathEscape(username))
	renderBody(c, body)
}

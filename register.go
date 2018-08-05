package retwis

import (
	"github.com/gin-gonic/gin"
	"time"
	"fmt"
	"net/url"
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
	if val, err := r.Hget("users", username); err == nil && val != "" {
		goback(c, "Sorry the selected username is already in use.")
		return
	}

	// Everything is ok, Register the user!
	userid, _ := r.Incr("next_user_id")
	authsecret := getrand()
	r.Hset("users", username, userid)
	r.Hmset("user:"+userid,
		"username", username,
		"password", password,
		"auth", authsecret)
	r.Hset("auths", authsecret, userid)

	r.Zadd("users_by_time", time.Now().Unix(), username)

	// User registered! Login her / him.
	setcookie(c, "auth", intToString(authsecret), int(time.Now().Unix() + 3600*24*365))

	body := fmt.Sprintf(`
<h2>Welcome aboard!</h2>
Hey %s, now you have an account, <a href="index">a good start is to write your first message!</a>.
`,
	url.PathEscape(username))
	renderBody(c, body)
}

package middlewares

import (
	b64 "encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func SessionChecker(fn gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("token")
		if err != nil {
			c.Redirect(http.StatusMovedPermanently, "/auth")
			//c.JSON(http.StatusUnauthorized, nil)
			return
		}
		ok := checkToken(token)
		if !ok {
			//c.JSON(http.StatusUnauthorized, nil)
			c.Redirect(http.StatusMovedPermanently, "/auth")
			return
		} else {
			fn(c)
		}

	}
}

func CheckBasicAuth(credentials string) (string, error) {
	cred, _ := b64.URLEncoding.DecodeString(credentials)
	credential := strings.Split(string(cred), ":")
	// check ipa if true return token else return err
	fmt.Println(credential)
	return "ZAEBISTOKEN", nil
}

func checkToken(token string) bool {
	return true
}

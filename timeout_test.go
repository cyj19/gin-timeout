/**
 * @Author: cyj19
 * @Date: 2021/12/26 20:05
 */

// Test cases of ContextTimeout Middleware

package timeout

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestContextTimeout(t *testing.T) {
	r := gin.Default()
	timeout := 5 * time.Second
	opt := Option{
		Timeout: &timeout,
		Code:    500,
		Msg:     []byte("handle timeout"),
	}

	r.Use(ContextTimeout(opt))
	r.GET("/ping", func(c *gin.Context) {
		time.Sleep(6 * time.Second)
		c.String(http.StatusOK, "pong")
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	r.ServeHTTP(w, req)
	fmt.Println(w.Body.String())
}

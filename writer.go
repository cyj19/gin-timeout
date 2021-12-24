/**
 * @Author: cyj19
 * @Date: 2021/12/24 16:48
 */

package timeout

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
)

type timeoutWriter struct {
	gin.ResponseWriter
	h           http.Header   // 正常响应头
	wbuf        *bytes.Buffer // 正常响应内容
	mu          sync.Mutex
	timedOut    bool
	wroteHeader bool
	code        int
}

func (tw *timeoutWriter) Header() http.Header {
	return tw.h
}

func checkWriteHeaderCode(code int) {
	if code < 100 || code > 999 {
		panic(fmt.Sprintf("invalid WriteHeader code %v", code))
	}
}

func (tw *timeoutWriter) writeHeaderLocked(code int) {
	checkWriteHeaderCode(code)

	switch {
	case tw.timedOut:
		return
	case tw.wroteHeader:
		return
	default:
		tw.wroteHeader = true
		tw.code = code
		return
	}
}

func (tw *timeoutWriter) WriteHeader(code int) {
	tw.mu.Lock()
	tw.mu.Unlock()
	tw.writeHeaderLocked(code)
}

func (tw *timeoutWriter) Write(p []byte) (int, error) {
	tw.mu.Lock()
	tw.mu.Unlock()
	if tw.timedOut {
		return 0, nil
	}
	if !tw.wroteHeader {
		tw.WriteHeader(http.StatusOK)
	}
	// 正常响应内容写入到wbuf
	return tw.wbuf.Write(p)
}

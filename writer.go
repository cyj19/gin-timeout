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

type TimeoutWriter struct {
	gin.ResponseWriter
	h           http.Header   // response header
	wbuf        *bytes.Buffer // response content
	mu          sync.Mutex
	timedOut    bool
	wroteHeader bool
	code        int // response code
}

func NewTimeoutWriter(w gin.ResponseWriter, buf *bytes.Buffer) *TimeoutWriter {
	return &TimeoutWriter{
		wbuf:           buf,
		ResponseWriter: w,
		h:              make(http.Header),
	}
}

func (tw *TimeoutWriter) Header() http.Header {
	return tw.h
}

func checkWriteHeaderCode(code int) {
	if code < 100 || code > 999 {
		panic(fmt.Sprintf("invalid WriteHeader code %v", code))
	}
}

func (tw *TimeoutWriter) writeHeaderLocked(code int) {
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

func (tw *TimeoutWriter) WriteHeader(code int) {
	tw.mu.Lock()
	tw.mu.Unlock()
	tw.writeHeaderLocked(code)
}

func (tw *TimeoutWriter) Write(p []byte) (int, error) {
	tw.mu.Lock()
	tw.mu.Unlock()
	if tw.timedOut {
		return 0, nil
	}
	if !tw.wroteHeader {
		tw.WriteHeader(http.StatusOK)
	}
	// normal response content is written to wbuf
	return tw.wbuf.Write(p)
}

func (tw *TimeoutWriter) Reset() {
	tw.wbuf.Reset()
}

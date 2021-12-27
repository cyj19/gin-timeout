/**
 * @Author: cyj19
 * @Date: 2021/12/24 17:18
 */

// Package timeout middleware for gin
package timeout

import (
	"bytes"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ContextTimeout(opt Option) gin.HandlerFunc {
	return func(c *gin.Context) {
		var tw = &timeoutWriter{
			wbuf:           bytes.NewBufferString(""),
			ResponseWriter: c.Writer,
			h:              make(http.Header),
		}
		c.Writer = tw
		ctx, cancel := context.WithTimeout(c.Request.Context(), *opt.Timeout)
		defer cancel()

		done := make(chan struct{})
		panicChan := make(chan interface{}, 1)
		go func() {
			defer func() {
				if p := recover(); p != nil {
					panicChan <- p
				}
			}()
			c.Next()
			close(done)
		}()

		select {
		case p := <-panicChan:
			panic(p)
		case <-ctx.Done():
			tw.mu.Lock()
			tw.mu.Unlock()
			tw.ResponseWriter.WriteHeader(opt.Code)
			_, _ = tw.ResponseWriter.Write(opt.Msg)
			tw.timedOut = true
			c.Abort()
		case <-done:
			tw.mu.Lock()
			tw.mu.Unlock()
			dst := c.Writer.Header()
			// add the header of timeoutWriter to the original header
			for k, vv := range tw.h {
				dst[k] = vv
			}
			if !tw.wroteHeader {
				tw.code = http.StatusOK
			}
			tw.ResponseWriter.WriteHeader(tw.code)
			_, _ = tw.ResponseWriter.Write(tw.wbuf.Bytes())
		}
	}
}

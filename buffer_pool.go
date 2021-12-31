/**
 * @Author: cyj19
 * @Date: 2021/12/30 17:02
 */

package timeout

import (
	"bytes"
	"sync"
)

type BufferPool struct {
	pool sync.Pool
}

func NewBufferPool() *BufferPool {
	return &BufferPool{}
}

func (p *BufferPool) Get() *bytes.Buffer {
	buf := p.pool.Get()
	if buf == nil {
		buf = &bytes.Buffer{}
		p.pool.Put(buf)
	}
	return buf.(*bytes.Buffer)
}

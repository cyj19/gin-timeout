/**
 * @Author: cyj19
 * @Date: 2021/12/31 9:09
 */

package timeout

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestBufferPool(t *testing.T) {
	bufPool := NewBufferPool()
	buf := bufPool.Get()
	assert.NotEqual(t, nil, buf)
}

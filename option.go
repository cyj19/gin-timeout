/**
 * @Author: cyj19
 * @Date: 2021/12/24 17:17
 */

package timeout

import "time"

// Option 超时配置
type Option struct {
	Timeout *time.Duration
	Code    int
	Msg     []byte
}

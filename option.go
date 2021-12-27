/**
 * @Author: cyj19
 * @Date: 2021/12/24 17:17
 */

package timeout

import "time"

// Option timeout configuration
type Option struct {
	Timeout *time.Duration // The timeout time is generally configured in the configuration file. In order to facilitate hot update, use the pointer
	Code    int            // response code
	Msg     string         // response message
}

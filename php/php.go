// php的时间用法
package php

import "time"

// Time return int64的时间戳
func Time() int64 {
	return time.Now().Unix()
}

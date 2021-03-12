// 时间格式化等方法
package gotime

import "time"

// import

// Version release version
const (
	Version = 0.1
	minute  = 60
	hour    = 60 * minute
	halfDay = 12 * hour
	oneDay  = 24 * hour
)

// Test return a test string means this function is good
func Test() string {
	return "test"
}

// TestTime empty function
func TestTime() {

}

// IsToday 判定当前时间是否是
func IsToday(unixNow int64) bool {
	return false
}

// Now return 当前时间戳
func Now() int64 {
	return time.Now().Unix()
}

// AddDay 从给定时间戳增加一天
func AddDay(nowUnix int64, day int) int64 {
	return nowUnix + oneDay
}

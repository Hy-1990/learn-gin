package utils

import (
	"time"
)

//time.Time类型转化为时间字符串、时间戳
func GetTime(tm time.Time) (string, int64) {
	_timeStr := tm.Format("2006-01-02 15:04:05")
	_timeStamp := tm.Unix()
	return _timeStr, _timeStamp
}

//时间戳转化为时间字符串
func TimestampToDatetime(tm int64) string {
	_timeUnix := tm
	return time.Unix(_timeUnix, 0).Format("2006-01-02 15:04:05")
}

//时间戳转time.Time类型
func TimestampToTime(tm int64) time.Time {
	_timeUnix := TimestampToDatetime(tm)
	_formatTime, _ := time.Parse("2006-01-02 15:04:05", _timeUnix)
	return _formatTime

}

//字符串时间转时间戳
func DatetimeToTime(dateTime string) (tm time.Time) {
	_formatTime, _ := time.Parse("2006-01-02 15:04:05", dateTime)
	return _formatTime
}

//字符串时间转时间戳
func DatetimeToTimestamp(dataTime string) (tm int64) {
	_formatTime, _ := time.Parse("2006-01-02 15:04:05", dataTime)
	return _formatTime.Unix()
}

// Package dates /*
package dates

import (
	"time"
)

const (
	// DayFormatter 日期格式
	DayFormatter = "2006-01-02"
	// DayTimeMinuteFormatter 日期时间格式--分
	DayTimeMinuteFormatter = "2006-01-02 15:04"
	// DayTimeSecondFormatter 日期时间格式--秒
	DayTimeSecondFormatter = "2006-01-02 15:04:05"
	// DayTimeMillisecondFormatter 日期时间格式--毫秒
	DayTimeMillisecondFormatter = "2006-01-02 15:04:05.sss"
	// TimeSecondFormatter 时间格式--秒
	TimeSecondFormatter = "15:04:05"
)

// NowFormatted 默认格式2006-01-02 15:04:05
func NowFormatted() string {
	return time.Now().Format(DayTimeSecondFormatter)
}

func NowLayout(layout string) string {
	return time.Now().Format(layout)
}

func Layout(date time.Time, layout string) string {
	return date.Format(layout)
}

func DefaultLayout(time time.Time) string {
	return time.Format(DayTimeSecondFormatter)
}

func FromDefaultLayout(str string) time.Time {
	loc, _ := time.LoadLocation("Local")
	theTime, _ := time.ParseInLocation(DayTimeSecondFormatter, str, loc)
	return theTime
}

// NowMillisecond 当前的毫秒时间戳
func NowMillisecond() int64 {
	return time.Now().UnixNano() / 1e6
}

func PastDayDate(pastDay int) time.Time {
	return time.Now().AddDate(0, 0, -pastDay)
}

func FutureDayDate(futureDay int) time.Time {
	return time.Now().AddDate(0, 0, futureDay)
}

func WeekStartDayDate() time.Time {
	now := time.Now()
	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}
	weekStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	return weekStart
}

func MonthStartDayDate() time.Time {
	year, month, _ := time.Now().Date()
	monthStart := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	return monthStart
}

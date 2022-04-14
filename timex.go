package timex

import (
	"errors"
	"strings"
	"time"
)

const DateLayout = "2006-01-02"
const DatetimeLayout = "2006-01-02 15:04:05"
const DateHourMinLayout = "2006-01-02 15:04"

func Now() time.Time {
	return time.Now()
}

func NowUnix() int64 {
	return time.Now().Unix()
}

func Date(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, time.Local)
}

func DateOf(tm time.Time) time.Time {
	return Date(tm.Year(), tm.Month(), tm.Day())
}

func DateHourOf(tm time.Time) time.Time {
	return time.Date(tm.Year(), tm.Month(), tm.Day(), tm.Hour(), 0, 0, 0, time.Local)
}

func DateHalfHourOf(tm time.Time) time.Time {
	if tm.Minute() >= 30 {
		return time.Date(tm.Year(), tm.Month(), tm.Day(), tm.Hour(), 30, 0, 0, time.Local)
	}
	return time.Date(tm.Year(), tm.Month(), tm.Day(), tm.Hour(), 0, 0, 0, time.Local)
}

func DateHourMinOf(tm time.Time) time.Time {
	return time.Date(tm.Year(), tm.Month(), tm.Day(), tm.Hour(), tm.Minute(), 0, 0, time.Local)
}

func DateOfTimestamp(ts int64) time.Time {
	return DateOf(Unix(ts))
}

func DateOfToday() time.Time {
	return DateOf(Now())
}

func DateOfTodayUnix() int64 {
	return DateOfToday().Unix()
}

func ZeroOfToday() time.Time {
	return DateOfToday()
}

func ZeroOfTodayUnix() int64 {
	return ZeroOfToday().Unix()
}

func Yesterday() time.Time {
	return Now().AddDate(0, 0, -1)
}

func YesterdayUnix() int64 {
	return Yesterday().Unix()
}

func DateOfYesterday() time.Time {
	return DateOf(Yesterday())
}

func DateOfYesterdayUnix() int64 {
	return DateOfYesterday().Unix()
}

func ZeroOfYesterday() time.Time {
	return DateOfYesterday()
}

func ZeroOfYesterdayUnix() int64 {
	return ZeroOfYesterday().Unix()
}

func Tomorrow() time.Time {
	return Now().AddDate(0, 0, 1)
}

func DateOfTomorrow() time.Time {
	return DateOf(Tomorrow())
}

func DateOfTomorrowUnix() int64 {
	return DateOfTomorrow().Unix()
}

func ZeroOfTomorrow() time.Time {
	return DateOfTomorrow()
}

func ZeroOfTomorrowUnix() int64 {
	return ZeroOfTomorrow().Unix()
}

func DaysLater(days int) time.Time {
	return Now().AddDate(0, 0, days)
}

func Unix(timestamp int64) time.Time {
	return time.Unix(timestamp, 0)
}

// Parse 将字符串解析成本地时间
func Parse(value string) (time.Time, error) {
	return time.ParseInLocation(DatetimeLayout, value, time.Local)
}

// ParseDate 将字符串解析成本地日期
func ParseDate(value string) (time.Time, error) {
	return time.ParseInLocation(DateLayout, value, time.Local)
}

func Format(tm time.Time) string {
	return tm.Format(DatetimeLayout)
}

func FormatDate(tm time.Time) string {
	return tm.Format(DateLayout)
}

func FormatDateHourMin(tm time.Time) string {
	return tm.Format(DateHourMinLayout)
}

func FormatFromTimestamp(ts int64) string {
	return time.Unix(ts, 0).Format(DatetimeLayout)
}

func FormatDateFromTimestamp(ts int64) string {
	return time.Unix(ts, 0).Format(DateLayout)
}

func FormatDateHourMinFromTimestamp(ts int64) string {
	return time.Unix(ts, 0).Format(DateHourMinLayout)
}

func ParseISO8601(value string) (time.Time, error) {
	var tm time.Time
	if !strings.HasSuffix(value, "Z") {
		return tm, errors.New("invalid format")
	}

	dotPos := strings.Index(value, ".")
	if dotPos > 0 {
		value = value[:dotPos] + "Z"
	}

	tm, err := time.Parse("2006-01-02T15:04:05Z", value)
	if err != nil {
		return tm, err
	}
	return time.Unix(tm.Unix(), 0), nil
}

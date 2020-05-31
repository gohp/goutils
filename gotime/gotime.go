package gotime

import (
	"fmt"
	"strings"
	"time"
)

type Gotime struct {
	time.Time
}

func (g Gotime) MarshalJSON() ([]byte, error){
	stamp := fmt.Sprintf("\"%s\"", FormatDatetime(g, "YYYY-MM-DD HH:mm:ss"))
	return []byte(stamp), nil
}

// Format time.Time struct to string
// MM - month - 01
// M - month - 1, single bit
// DD - day - 02
// D - day 2
// YYYY - year - 2006
// YY - year - 06
// HH - 24 hours - 03
// H - 24 hours - 3
// hh - 12 hours - 03
// h - 12 hours - 3
// mm - minute - 04
// m - minute - 4
// ss - second - 05
// s - second = 5

func FormatDatetime(t time.Time, format string) string {
	res := strings.Replace(format, "MM", t.Format("01"), -1)
	res = strings.Replace(res, "M", t.Format("1"), -1)
	res = strings.Replace(res, "DD", t.Format("02"), -1)
	res = strings.Replace(res, "D", t.Format("2"), -1)
	res = strings.Replace(res, "YYYY", t.Format("2006"), -1)
	res = strings.Replace(res, "YY", t.Format("06"), -1)
	res = strings.Replace(res, "HH", fmt.Sprintf("%02d", t.Hour()), -1)
	res = strings.Replace(res, "H", fmt.Sprintf("%d", t.Hour()), -1)
	res = strings.Replace(res, "hh", t.Format("03"), -1)
	res = strings.Replace(res, "h", t.Format("3"), -1)
	res = strings.Replace(res, "mm", t.Format("04"), -1)
	res = strings.Replace(res, "m", t.Format("4"), -1)
	res = strings.Replace(res, "ss", t.Format("05"), -1)
	res = strings.Replace(res, "s", t.Format("5"), -1)
	return res
}

func Now() string {
	return FormatDatetime(time.Now(), "YYYY-MM-DD HH:mm:ss")
}

func NowUnix() int64 {
	return time.Now().Unix()
}

// TodayStart eg: 2018-01-01 00:00:00
func TodayStart() string {
	now := time.Now()
	tm := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0,  time.Local)
	return FormatDatetime(tm, "YYYY-MM-DD HH:mm:ss")
}

// TodayEnd eg: 2018-01-01 23:59:59
func TodayEnd() string {
	now := time.Now()
	tm := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 1e9-1,  time.Local)
	return FormatDatetime(tm, "YYYY-MM-DD HH:mm:ss")
}

func ToUnix() int64 {
	panic("1")
}
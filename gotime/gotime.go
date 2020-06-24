package gotime

import (
	"fmt"
	"strings"
	"time"
)

const (
	RFC3339 = "YYYY-MM-DDTHH:mm:ss+08:00"
	TT      = "YYYY-MM-DD HH:mm:ss"
	Y_M_D   = "YYYY-MM-DD"
	YMD     = "YYYYMMDD"
	HMS     = "HH:mm:ss"
)

type Time time.Time

func (g Time) MarshalJSON() ([]byte, error) {
	// 判空
	if time.Time(g).UnixNano() == (time.Time{}).UnixNano() {
		return []byte("\"\""), nil
	}
	stamp := fmt.Sprintf("\"%s\"", g.Format(TT))
	return []byte(stamp), nil
}

func (g *Time) UnmarshalJSON(input []byte) error {
	if string(input) == "" {
		*g = Time{}
		return nil
	}
	tt, err := time.ParseInLocation(`"`+"2006-01-02 15:04:05"+`"`, string(input), time.Local)
	*g = Time(tt)
	return err
}

func (g Time) Time() time.Time {
	return time.Time(g)
}

func (g Time) Format(format string) string {
	return FormatDatetime(time.Time(g), format)
}

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
	tm := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	return FormatDatetime(tm, "YYYY-MM-DD HH:mm:ss")
}

// TodayEnd eg: 2018-01-01 23:59:59
func TodayEnd() string {
	now := time.Now()
	tm := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 1e9-1, time.Local)
	return FormatDatetime(tm, "YYYY-MM-DD HH:mm:ss")
}

func ToUnix() int64 {
	panic("1")
}

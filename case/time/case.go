package time

import (
	"fmt"
	"time"
)

var TW = time.FixedZone("TW", +8*3600)

func GetDayStartTime(t time.Time, timeZone int64) time.Time {
	loc := time.FixedZone("", int(timeZone))
	now := time.Now().In(loc)
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc)
}

func Run() {
	now := time.Now().In(TW)
	fmt.Println(now)

	cstTest := GetDayStartTime(time.Now(), 8*3600)
	utcTest := GetDayStartTime(time.Now(), 0)
	fmt.Println(cstTest, cstTest.Unix())
	fmt.Println(utcTest, utcTest.Unix())
}

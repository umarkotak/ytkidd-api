package utils

import (
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/umarkotak/ytkidd-api/config"
)

func StringMustInt64(s string) int64 {
	i, _ := strconv.ParseInt(s, 10, 64)
	return i
}

func StringMustInt(s string) int {
	i, _ := strconv.ParseInt(s, 10, 64)
	return int(i)
}

func StringMustDate(s string) time.Time {
	t, err := time.Parse("2006-01-02", s)
	loc, _ := time.LoadLocation(config.Get().DbTimezone)
	t = t.In(loc)
	if err != nil {
		logrus.Error(err)
	}
	return t
}

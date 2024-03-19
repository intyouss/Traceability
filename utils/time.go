package utils

import (
	"strconv"
	"strings"
	"time"
)

// TimeFormat 时间格式化
func TimeFormat(t time.Time) string {
	since := time.Since(t)
	switch {
	case since < time.Minute: // 如果是一分钟内的时间，返回刚刚
		return "刚刚"
	case since < time.Hour: // 如果是一小时内的时间，返回是几分钟前
		return strings.Split(since.String(), "m")[0] + "分钟前"
	case since < 24*time.Hour: // 如果是超过一个小时的时间，返回是几小时前
		return strings.Split(since.String(), "h")[0] + "小时前"
	case since < 7*24*time.Hour: // 如果超过一天但是在一周内，返回是几天前
		x, _ := strconv.Atoi(strings.Split(since.String(), "h")[0])
		return strconv.Itoa(x/24) + "天前"
	case since < 21*24*time.Hour: // 如果超过一周，但不超过三周，返回是几周前
		x, _ := strconv.Atoi(strings.Split(since.String(), "h")[0])
		return strconv.Itoa(x/(7*24)) + "周前"
	default: // 如果超过三周，返回年月日
		return t.Format("2006-01-02")
	}
}

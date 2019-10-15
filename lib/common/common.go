package common

import (
	"os"
	"time"
)

//把年月日时分秒的时间转化为时间戳
func TimeToTimestamp(layout, timestr string) (int64, error) {
	loc, err := time.LoadLocation("Local")
	if err != nil {
		return -1, err
	}
	times, err := time.ParseInLocation(layout, timestr, loc)
	if err != nil {
		return -1, err
	}
	trantimestamp := times.Unix()
	return trantimestamp, nil
}

/*
判断指定的文件是否存在
*/
func FileExists(filePath string) (bool, error) {
	_, err := os.Stat(filePath)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, nil
}

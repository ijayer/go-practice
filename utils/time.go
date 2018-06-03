package utils

import "time"

const (
	DateFormat = "yyyy-MM-dd hh:mm:ss"
	TimeLayout = "2006-01-02 15:04:05"
	DateLayout = "2006-01-02"
)

func Now() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func Date() string {
	return time.Now().Format("2006-01-02")
}

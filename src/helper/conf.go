package helper

import "time"

type conf struct {
	DayCount    int64
	Host        string
	Port        int
	Database    string
	Username    string
	Password    string
	SSLMode     string
	HourCounter int64
	DayCounter  int64
	WeekCounter int64

	HourSecond time.Duration
	DaySecond  time.Duration
	WeekSecond time.Duration
}

var Conf *conf

func InitConf() {
	Conf = &conf{
		DayCount:    1,
		Host:        "127.0.0.1",
		Port:        5432,
		Database:    "student-scoring-case",
		Username:    "postgres",
		Password:    "tayitkan",
		SSLMode:     "disable",
		HourCounter: 0,
		DayCounter:  0,
		WeekCounter: 0,
		HourSecond:  10 * time.Second,
		DaySecond:   240 * time.Second,
		WeekSecond:  1680 * time.Second,
	}
}

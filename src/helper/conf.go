package helper

import "time"

type conf struct {
	Host     string
	Port     int
	Database string
	Username string
	Password string
	SSLMode  string

	HourSecond time.Duration
}

var Conf *conf

func InitConf() {
	Conf = &conf{
		Host:       "127.0.0.1",
		Port:       5432,
		Database:   "student-scoring-case",
		Username:   "postgres",
		Password:   "tayitkan",
		SSLMode:    "disable",
		HourSecond: 10 * time.Second,
	}
}

type timeStruct struct {
	HourCounter int
	DayCounter  int
	WeekCounter int
}

var TimeStruct *timeStruct

func InıtTimeStruct() {
	TimeStruct = &timeStruct{
		HourCounter: 0,
		DayCounter:  0,
		WeekCounter: 0,
	}
}

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
var ConInfo *PgConnectionInfo

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
	ConInfo = &PgConnectionInfo{
		Host:     Conf.Host,
		Port:     Conf.Port,
		Database: Conf.Database,
		Username: Conf.Username,
		Password: Conf.Password,
		SSLMode:  Conf.SSLMode,
	}
}

type timeStruct struct {
	HourCounter int
	DayCounter  int
	WeekCounter int
}

var TimeStruct *timeStruct

func InitTimeStruct() {
	TimeStruct = &timeStruct{
		HourCounter: 0,
		DayCounter:  0,
		WeekCounter: 0,
	}
}

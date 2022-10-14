package helper

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
	}
}

package main

import (
	"errors"
	"fmt"
	"github.com/burakkuru5534/src/helper"
	"github.com/go-co-op/gocron"
	"time"

	_ "github.com/lib/pq"
)

var DayCount = 1

func main() {

	conInfo := helper.PgConnectionInfo{
		Host:     "127.0.0.1",
		Port:     5432,
		Database: "student-scoring-case",
		Username: "postgres",
		Password: "tayitkan",
		SSLMode:  "disable",
	}

	db, err := helper.NewPgSqlxDbHandle(conInfo, 10)
	if err != nil {
		errors.New("create db handle error.")
	}
	err = db.Ping()
	if err != nil {
		errors.New("ping db error.")
	}

	// Create Appplication Service
	err = helper.InitApp(db)
	if err != nil {
		errors.New("init app error.")
	}

	s := gocron.NewScheduler(time.UTC)
	s.Every(10).Seconds().Do(helper.DailyProcess)
	s.Every(70).Seconds().Do(helper.WeeklyProcess)
	s.StartAsync()

	for i := 0; i <= 0; {
		fmt.Println("Select an option:")
		fmt.Println("1.To end the day")
		fmt.Println("2.Score a student")
		fmt.Println("3.List students")
		fmt.Println("4.Exit")
		fmt.Print("Enter your choice:")
		var choice int
		// Taking input from user
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			fmt.Println("End of the day")
			DayCount++

		case 2:
			fmt.Println("Score a student")
		case 3:
			fmt.Println("List students")
		case 4:
			fmt.Println("Exit")
			i++
		}

		if DayCount%7 == 0 {
			helper.WeeklyProcess()
		}

	}

}

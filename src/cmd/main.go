package main

import (
	"errors"
	"fmt"
	"github.com/burakkuru5534/src/helper"
	"github.com/burakkuru5534/src/model"
	"time"

	_ "github.com/lib/pq"
)

func main() {

	helper.InitConf()

	conInfo := helper.PgConnectionInfo{
		Host:     helper.Conf.Host,
		Port:     helper.Conf.Port,
		Database: helper.Conf.Database,
		Username: helper.Conf.Username,
		Password: helper.Conf.Password,
		SSLMode:  helper.Conf.SSLMode,
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

	ticker1hour := time.NewTicker(1 * time.Second)
	ticker1day := time.NewTicker(24 * time.Second)
	ticker7day := time.NewTicker(168 * time.Second)
	myMatrix := make([][]int64, 24)
	for i := range myMatrix {
		myMatrix[i] = make([]int64, 7)
	}
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case _ = <-ticker1hour.C:
				fmt.Println("HourCounter: ", helper.Conf.HourCounter)
				myMatrix[helper.Conf.HourCounter][helper.Conf.DayCounter] = 1
				fmt.Println(fmt.Sprintf("myMatrix[%d][%d]: %d", helper.Conf.HourCounter, helper.Conf.DayCounter, myMatrix[helper.Conf.HourCounter][helper.Conf.DayCounter]))
				helper.Conf.HourCounter += 1

			case _ = <-ticker1day.C:
				for i := range myMatrix {
					if myMatrix[i][helper.Conf.DayCounter] == 1 {
						//give point to student
						fmt.Println("give point to student for hour: ", i)
					}
				}
				helper.Conf.HourCounter = 0
				helper.Conf.DayCounter += 1
				fmt.Println("DayCounter: ", helper.Conf.DayCounter)
				fmt.Println("HourCounter: ", helper.Conf.HourCounter)

			case _ = <-ticker7day.C:
				helper.Conf.HourCounter = 0
				helper.Conf.DayCounter = 0
				helper.Conf.WeekCounter += 1
				myMatrix = make([][]int64, 24)
				for i := range myMatrix {
					myMatrix[i] = make([]int64, 7)
				}
				fmt.Println("WeekCounter: ", helper.Conf.WeekCounter)
				fmt.Println("hourCounter: ", helper.Conf.HourCounter)
				fmt.Println("dayCounter: ", helper.Conf.DayCounter)
			}
		}
	}()

	//// Tickers can be stopped like timers. Once a ticker
	//// is stopped it won't receive any more values on its
	//// channel. We'll stop ours after 1600ms.
	//time.Sleep(1000600 * time.Millisecond)
	//ticker1hour.Stop()
	//ticker7day.Stop()
	//done <- true
	//fmt.Println("Ticker stopped")
	//

	var choice int
	for isExit := false; !isExit; {
		fmt.Println("Select an option:")
		fmt.Println("1.To end the day")
		fmt.Println("2.Score a student")
		fmt.Println("3.List students")
		fmt.Println("4.To end the week")
		fmt.Println("5.Exit")
		fmt.Print("Enter your choice:")
		// Taking input from user
		_, err = fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("Please enter a valid number")
			panic(err)
		}
		switch choice {
		case 1:
			fmt.Println("End of the day")
			helper.Conf.HourCounter = 0
			helper.Conf.DayCounter += 1
			if helper.Conf.DayCounter == 7 {
				helper.Conf.DayCounter = 0
				helper.Conf.WeekCounter += 1
				myMatrix = make([][]int64, 24)
				for i := range myMatrix {
					myMatrix[i] = make([]int64, 7)
				}
			}

		case 2:
			fmt.Println("Score a student")
			var studentNumber, point int64
			fmt.Print("Enter student number:")
			_, err = fmt.Scanln(&studentNumber)
			if err != nil {
				fmt.Println("Please enter a valid number")
				panic(err)
			}
			fmt.Print("Enter point:")
			_, err = fmt.Scanln(&point)
			if err != nil {
				fmt.Println("Please enter a valid number")
				panic(err)
			}
			student := model.Student{}
			student.GivePointToStudent(studentNumber, point)

		case 3:
			fmt.Println("List students")
			//student := model.Student{}
			//student.ListStudents()
		case 4:
			fmt.Println("End of the week")
			helper.Conf.HourCounter = 0
			helper.Conf.DayCounter = 0
			helper.Conf.WeekCounter += 1
		case 5:
			fmt.Println("Exit")
			isExit = true
		}

	}

}

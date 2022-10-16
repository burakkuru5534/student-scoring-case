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

	ticker1hour := time.NewTicker(helper.Conf.HourSecond) // 10 second
	ticker1day := time.NewTicker(helper.Conf.DaySecond)   // 240 second
	ticker7day := time.NewTicker(helper.Conf.WeekSecond)  // 1680 second
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

				myMatrix[helper.Conf.HourCounter][helper.Conf.DayCounter] = 1

				if myMatrix[helper.Conf.HourCounter][helper.Conf.DayCounter] == 1 {

					students := model.GetGroupAStudentList()
					for i, student := range students {

						if i < 4 {
							err = student.GivePointToStudent(student.Number, 1)
							if err != nil {
							}
						} else if i < 8 {
							err = student.GivePointToStudent(student.Number, 2)
							if err != nil {
							}
						} else if i < 10 {
							err = student.GivePointToStudent(student.Number, 3)
							if err != nil {
							}
						}
					}
				}
				helper.Conf.HourCounter += 1
				if helper.Conf.HourCounter == 23 {
					helper.Conf.HourCounter = 0
					helper.Conf.DayCounter += 1
					if helper.Conf.DayCounter == 6 {
						helper.Conf.HourCounter = 0
						helper.Conf.DayCounter = 0
						helper.Conf.WeekCounter += 1
					}
				}

			case _ = <-ticker1day.C:

				helper.Conf.HourCounter = 0
				helper.Conf.DayCounter += 1
				if helper.Conf.DayCounter == 6 {
					helper.Conf.HourCounter = 0
					helper.Conf.DayCounter = 0
					helper.Conf.WeekCounter += 1
				}
				myMatrix = make([][]int64, 24)
				for i := range myMatrix {
					myMatrix[i] = make([]int64, 7)
				}

			case _ = <-ticker7day.C:
				helper.Conf.HourCounter = 0
				helper.Conf.DayCounter = 0
				helper.Conf.WeekCounter += 1
				myMatrix = make([][]int64, 24)
				for i := range myMatrix {
					myMatrix[i] = make([]int64, 7)
				}

			}
		}
	}()

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

			if point > 5 || point < -5 {
				fmt.Println("Please enter a valid number")
				continue
			}
			student := model.Student{}
			err = student.GivePointToStudent(studentNumber, point)
			if err != nil {
				fmt.Println("Error: ", err)
			}
			err = student.MakeChangeIfFirstStudentOfBMoreThanLastStudentOfA()
			if err != nil {
				fmt.Println("Error: ", err)
			}

		case 3:
			fmt.Println("List students")
			student := model.Student{}
			students := student.ListStudents()
			for _, student := range students {
				fmt.Println(student)
			}
		case 4:
			helper.Conf.HourCounter = 0
			helper.Conf.DayCounter = 0
			helper.Conf.WeekCounter += 1

			student := model.Student{}
			err = student.ClearPoints()
			if err != nil {
				fmt.Println("Error: ", err)
			}
			myMatrix = make([][]int64, 24)
			for i := range myMatrix {
				myMatrix[i] = make([]int64, 7)
			}

		case 5:
			fmt.Println("Exit")
			isExit = true
		}

	}

}

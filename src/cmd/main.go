package main

import (
	"errors"
	"fmt"
	"github.com/burakkuru5534/src/helper"
	"github.com/burakkuru5534/src/model"
	_ "github.com/lib/pq"
)

func main() {

	helper.InitConf()
	helper.InıtTimeStruct()

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

	myMatrix := make([][]int64, 24)
	for i := range myMatrix {
		myMatrix[i] = make([]int64, 7)
	}

	ticker1hour := helper.NewTicker(0, helper.Conf.HourSecond) // 10 second
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case _ = <-ticker1hour.C:

				myMatrix[helper.TimeStruct.HourCounter][helper.TimeStruct.DayCounter] = 1
				fmt.Println("hour: ", helper.TimeStruct.HourCounter, " day: ", helper.TimeStruct.DayCounter, " week: ", helper.TimeStruct.WeekCounter)
				if myMatrix[helper.TimeStruct.HourCounter][helper.TimeStruct.DayCounter] == 1 {

					students := model.GetGroupAStudentList()
					for i, student := range students {
						var condition bool
						switch condition {

						case i < 4:
							err = student.GivePointToStudent(student.Number, 1)
							if err != nil {
								errors.New("give point error.")
							}
						case i < 8:
							err = student.GivePointToStudent(student.Number, 2)
							if err != nil {
								errors.New("give point error.")
							}
						case i < 10:
							err = student.GivePointToStudent(student.Number, 3)
							if err != nil {
								errors.New("give point error.")
							}
						}
					}
				}
				helper.TimeStruct.HourCounter += 1
				if helper.TimeStruct.HourCounter > 23 {
					helper.TimeStruct.HourCounter = 0
					helper.TimeStruct.DayCounter += 1

					if helper.TimeStruct.DayCounter > 6 {
						helper.TimeStruct.HourCounter = 0
						helper.TimeStruct.DayCounter = 0
						helper.TimeStruct.WeekCounter += 1

					}
					myMatrix = make([][]int64, 24)
					for i := range myMatrix {
						myMatrix[i] = make([]int64, 7)
					}
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

			helper.TimeStruct.HourCounter = 0
			helper.TimeStruct.DayCounter += 1
			if helper.TimeStruct.DayCounter > 6 {
				helper.TimeStruct.DayCounter = 0
				helper.TimeStruct.WeekCounter += 1
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
			helper.TimeStruct.HourCounter = 0
			helper.TimeStruct.DayCounter = 0
			helper.TimeStruct.WeekCounter += 1

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

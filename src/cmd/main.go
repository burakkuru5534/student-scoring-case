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
	helper.InitTimeStruct()

	db, err := helper.NewPgSqlxDbHandle(*helper.ConInfo, 10)
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

	//student operations
	err = mainProcess()
	if err != nil {
		errors.New("main process error.")
	}

}

func mainProcess() error {

	// we will give points to students with this matrix
	myMatrix := make([][]int64, 24)
	for i := range myMatrix {
		myMatrix[i] = make([]int64, 7)
	}

	//ticker to provide hourly give points to students and also track the time
	ticker1hour := helper.NewTicker(0, helper.Conf.HourSecond) // 10 second
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case _ = <-ticker1hour.C:

				// set 1 the current hour and day's index
				// if teacher end the day early we ll increase the day counter so previous day point's will be until last hour we give 1 for it.
				//after ending the day we ll start to give 1 to other day's hours.
				myMatrix[helper.TimeStruct.HourCounter][helper.TimeStruct.DayCounter] = 1

				// if current hour and day matrix value equal to 1 then we can give points for that hour.
				if myMatrix[helper.TimeStruct.HourCounter][helper.TimeStruct.DayCounter] == 1 {

					// get all group A students from db. bc we'll give points to only group A students
					var student model.Student
					students, _ := student.GetStudentsOfGroupA()
					for i, student := range students {
						switch true {

						case i < 2:
							//we get students from db with order by point and number asc
							//so first 4 students will have 1 points
							err := student.GivePointToStudent(student.Number, 3)
							if err != nil {
								errors.New("give point error.")

							}
						case i < 6:
							//next 4 students will have 2 points
							err := student.GivePointToStudent(student.Number, 2)
							if err != nil {
								errors.New("give point error.")
							}
						case i < 10:
							//last 2 students will have 3 points
							err := student.GivePointToStudent(student.Number, 1)
							if err != nil {
								errors.New("give point error.")
							}
						}
					}
				}
				//after giving points to students we ll increase the hour counter
				helper.TimeStruct.HourCounter += 1
				if helper.TimeStruct.HourCounter > 23 {
					//if hour counter is bigger than 23 we ll increase the day counter and reset the hour counter
					helper.TimeStruct.HourCounter = 0
					helper.TimeStruct.DayCounter += 1

					if helper.TimeStruct.DayCounter > 6 {
						//if day counter is bigger than 6 we ll increase the week counter and reset the day counter and also hour counter
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
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("Please enter a valid number")
		}

		switch choice {
		case 1:

			//end the day
			//reset hour counter and increase day counter
			helper.TimeStruct.HourCounter = 0
			helper.TimeStruct.DayCounter += 1

			if helper.TimeStruct.DayCounter > 6 {
				//if day counter is bigger than 6 we ll increase the week counter and reset the day counter
				//we already reset the hour counter when we end the day
				helper.TimeStruct.DayCounter = 0
				helper.TimeStruct.WeekCounter += 1

				//we need to reset the matrix and also clear students points
				student := model.Student{}
				err = student.ClearPoints()
				if err != nil {
					fmt.Println("Error: ", err)
					return err
				}

				//initialize the matrix
				myMatrix = make([][]int64, 24)
				for i := range myMatrix {
					myMatrix[i] = make([]int64, 7)
				}
			}

		case 2:

			var studentNumber, point int64

			fmt.Print("Enter student number:")
			_, err = fmt.Scanln(&studentNumber)
			if err != nil {
				fmt.Println("Please enter a valid number")
				return err
			}

			fmt.Print("Enter point:")
			_, err = fmt.Scanln(&point)
			if err != nil {
				fmt.Println("Please enter a valid number")
				return err
			}

			//rule for give point to student: can't be greater than 5 and can't be less than -5
			if point > 5 || point < -5 {
				fmt.Println("Please enter a valid number")
				continue
			}
			student := model.Student{}
			err = student.GivePointToStudent(studentNumber, point)
			if err != nil {
				fmt.Println("Error: ", err)
				return err
			}

			//after giving point we need to check if first student of b much more than last student of a
			// if it is true we need to swap them
			err = student.MakeChangeIfFirstStudentOfBMoreThanLastStudentOfA()
			if err != nil {
				fmt.Println("Error: ", err)
				return err
			}

		case 3:
			//list students
			student := model.Student{}
			students := student.ListStudents()
			for _, student := range students {
				fmt.Println(student)
			}
		case 4:

			//end the week
			helper.TimeStruct.HourCounter = 0
			helper.TimeStruct.DayCounter = 0
			helper.TimeStruct.WeekCounter += 1

			//we need to reset the matrix and also clear students points
			student := model.Student{}
			err = student.ClearPoints()
			if err != nil {
				fmt.Println("Error: ", err)
				return err
			}
			//initialize the matrix
			myMatrix = make([][]int64, 24)
			for i := range myMatrix {
				myMatrix[i] = make([]int64, 7)
			}

		case 5:
			//exit
			fmt.Println("Exit")
			isExit = true
		}

	}

	return nil
}

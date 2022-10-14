package helper

import (
	"github.com/burakkuru5534/src/model"
)

// every beginnig of the day we will give points to students
// there will be 2 groups
// 1. group A
// 2. group B
// group A will have 10 students
// group B will have 5 students
// first 2 students of group A will have 3 points
// next 4 students of group A will have 2 points
// last 4 students of group A will have 1 point

func DailyProcess() {

	students := getGroupAStudentList()
	for i, student := range students {

		if i < 4 {
			givePointToStudent(student.ID, 1)
		} else if i < 8 {
			givePointToStudent(student.ID, 2)
		} else if i < 10 {
			givePointToStudent(student.ID, 3)
		}
	}
}

func givePointToStudent(studentID, point int64) {

	sq := "update student set point = point + $1 where id = $2"
	_, err := App.DB.Exec(sq, point, studentID)
	if err != nil {
		panic(err)
	}

}
func getGroupAStudentList() []model.Student {

	sq := "select * from student where group_name = 'A' order by point,number asc"
	var students []model.Student
	err := App.DB.Select(&students, sq)
	if err != nil {
		panic(err)
	}
	return students
}

/*
   a. 10. 9. 8. 7. sıradaki öğrencilere 1
   b. 6. 5. 4. 3. sıradaki öğrencilere 2
   c. 2. 1. sıradaki öğrencilere 3 puan verilmektedir*/

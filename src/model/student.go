package model

import (
	"fmt"
	"github.com/burakkuru5534/src/helper"
)

type Student struct {
	ID        int64  `db:"id"`
	GroupName string `db:"group_name"`
	Number    int64  `db:"number"`
	Point     int64  `db:"point"`
}

func (s *Student) GivePointToStudent(number, point int64) error {
	s.Number = number
	s.Point = point
	sq := "update student set point = point + $1 where number = $2"
	_, err := helper.App.DB.Exec(sq, s.Point, s.Number)
	if err != nil {
		return err
	}

	return nil
}

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
	students := GetGroupAStudentList()
	for i, student := range students {

		if i < 4 {
			student.GivePointToStudent(student.Number, 1)
		} else if i < 8 {
			student.GivePointToStudent(student.Number, 2)
		} else if i < 10 {
			student.GivePointToStudent(student.Number, 3)
		}
	}
	helper.Conf.DayCount++
	fmt.Println("Day count:", helper.Conf.DayCount)
}
func GetGroupAStudentList() []Student {

	sq := "select * from student where group_name = 'A' order by point,number asc"
	var students []Student
	err := helper.App.DB.Select(&students, sq)
	if err != nil {
		panic(err)
	}
	return students
}

//set 0 all student's point

func (s *Student) WeeklyProcess() {

	sq := "update student set point = 0"
	_, err := helper.App.DB.Exec(sq)
	if err != nil {
		panic(err)
	}

}

func (s *Student) MakeChangeIfFirstStudentOfBMoreThanLastStudentOfA() error {
	sq := "select point,number from student where group_name = 'B' order by point desc limit 1"
	var studentB Student
	err := helper.App.DB.Get(&studentB, sq)
	if err != nil {
		return err
	}
	sq = "select point,number from student where group_name = 'A' order by point asc limit 1"
	var studentA Student
	err = helper.App.DB.Get(&studentA, sq)
	if err != nil {
		return err
	}

	if studentB.Point > studentA.Point {
		// we will change the first student of group B with the last student of group A
		sq = "update student set group_name = 'A' where number = $1"
		_, err = helper.App.DB.Exec(sq, studentB.Number)
		if err != nil {
			return err
		}
		sq = "update student set group_name = 'B' where number = $1"
		_, err = helper.App.DB.Exec(sq, studentA.Number)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Student) ListStudents() []Student {
	sq := "select * from student order by group_name,point desc,number asc"
	var students []Student
	err := helper.App.DB.Select(&students, sq)
	if err != nil {
		panic(err)
	}
	return students
}

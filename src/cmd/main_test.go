package main

import (
	"github.com/burakkuru5534/src/helper"
	"github.com/burakkuru5534/src/model"
	"testing"
)

func TestGivePointToStudent(t *testing.T) {

	err := helper.InitDb()
	if err != nil {
		t.Error(err)
	}

	var student model.Student
	currentPoint, err := student.GetPointOfStudentByNumber(101)
	if err != nil {
		t.Error(err)
	}
	err = student.GivePointToStudent(101, 3)
	if err != nil {
		t.Error(err)
	}

	got, err := student.GetPointOfStudentByNumber(101)
	if err != nil {
		t.Error(err)
	}
	want := currentPoint + 3

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestMakeChangeIfFirstStudentOfBMoreThanLastStudentOfA(t *testing.T) {

	err := helper.InitDb()
	if err != nil {
		t.Error(err)
	}

	var student model.Student

	//all students have 0 points
	err = student.ClearPoints()
	if err != nil {
		t.Error(err)
	}

	//we give -3 points to some students from group A
	err = student.GivePointToStudent(101, -3)
	if err != nil {
		t.Error(err)
	}

	//we give 3 points to some students from group B
	err = student.GivePointToStudent(111, 3)
	if err != nil {
		t.Error(err)
	}

	// 111 should be in group A and 101 should be in group B after this func
	err = student.MakeChangeIfFirstStudentOfBMoreThanLastStudentOfA()
	if err != nil {
		t.Error(err)
	}

	gotFirst, err := student.GetGroupOfStudentByNumber(111)
	if err != nil {
		t.Error(err)
	}
	wantFirst := "A"

	gotSecond, err := student.GetGroupOfStudentByNumber(101)
	if err != nil {
		t.Error(err)
	}
	wantSecond := "B"

	if gotFirst != wantFirst && gotSecond != wantSecond {
		t.Errorf("gotFirst %q, wantFirst %q gotSecond %q, wantSecond %q", gotFirst, wantFirst, gotSecond, wantSecond)
	}
}

func TestClearPoints(t *testing.T) {

	err := helper.InitDb()
	if err != nil {
		t.Error(err)
	}

	var student model.Student
	err = student.ClearPoints()
	if err != nil {
		t.Error(err)
	}

	got, err := student.GetPointOfStudentByNumber(101)
	if err != nil {
		t.Error(err)
	}
	want := int64(0)

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestGetStudentsOfGroupA(t *testing.T) {

	err := helper.InitDb()
	if err != nil {
		t.Error(err)
	}

	var student model.Student
	got, err := student.GetStudentsOfGroupA()
	if err != nil {
		t.Error(err)
	}
	want := 10

	if len(got) != want {
		t.Errorf("got %q, wanted %q", len(got), want)
	}
}

func TestGetPointOfStudentByNumber(t *testing.T) {

	err := helper.InitDb()
	if err != nil {
		t.Error(err)
	}

	var student model.Student

	err = student.ClearPoints()
	if err != nil {
		t.Error(err)
	}
	got, err := student.GetPointOfStudentByNumber(101)
	if err != nil {
		t.Error(err)
	}
	want := int64(0)

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestGetGroupOfStudentByNumber(t *testing.T) {

	err := helper.InitDb()
	if err != nil {
		t.Error(err)
	}

	var student model.Student
	got, err := student.GetGroupOfStudentByNumber(101)
	if err != nil {
		t.Error(err)
	}
	want := "A"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

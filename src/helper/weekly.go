package helper

//set 0 all student's point

func WeeklyProcess() {

	sq := "update student set point = 0"
	_, err := App.DB.Exec(sq)
	if err != nil {
		panic(err)
	}

}

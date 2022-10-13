package main

import (
	"fmt"
	"github.com/burakkuru5534/src/helper"
	"github.com/go-co-op/gocron"
	"time"
)

func main() {

	s := gocron.NewScheduler(time.UTC)
	s.Every(10).Seconds().Do(helper.DailyProcess)
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
		case 2:
			fmt.Println("Score a student")
		case 3:
			fmt.Println("List students")
		case 4:
			fmt.Println("Exit")
			i++
		}

	}

}

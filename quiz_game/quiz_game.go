package quiz_game

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

func Run() {

	var fileName string
	var timer int

	writeAns := 0
	wrongAns := 0

	flag.StringVar(&fileName, "f", "quiz_game/problems.csv", "used to define filename as flag")
	flag.IntVar(&timer, "t", 10, "timer if required")
	flag.Parse()

	seconds := time.Duration(timer) * time.Second
	done := make(chan bool)

	go func() {
		file, _ := os.Open(fileName)
		csvReader := csv.NewReader(file)
		records, _ := csvReader.ReadAll()

		var userInput string
		for _, eachRecord := range records {
			fmt.Println(eachRecord[0] + "?")

			fmt.Scanln(&userInput)

			if userInput == eachRecord[1] {
				writeAns++
			} else {
				wrongAns++
			}
		}

		done <- true
	}()

	select {
	case <-done:
		fmt.Print("correct ans", writeAns)
	case <-time.After(seconds):
		fmt.Print("correct ans", writeAns)

	}
}

package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("HELLO WORLD")
	csvFile, err := os.Open("sample.csv")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened CSV file")
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	fmt.Println(csvLines)

	consoleReader := bufio.NewReader(os.Stdin)

	correctAnswer := 0

	for i := 0; i < len(csvLines); i++ {
		fmt.Printf("What is %s sir?\n", csvLines[i][0])
		answer, err := consoleReader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(answer)
		ans, err := strconv.Atoi(strings.Trim(answer, "\n"))
		if err != nil {
			fmt.Println(err)
		}
		correctAns, err := strconv.Atoi(csvLines[i][1])
		if err != nil {
			fmt.Println(err)
		}
		if ans == correctAns {
			correctAnswer++
		}
	}
	fmt.Printf("Answered %d out of %d question correctly", correctAnswer, len(csvLines))

}

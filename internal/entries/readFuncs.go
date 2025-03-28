package entries

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func ReadString(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func ReadInt(prompt string) int {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	num, _ := strconv.Atoi(input)
	return num
}

func ReadFloat(prompt string) float64 {
	fmt.Print(prompt + ": ")
	var input float64
	fmt.Scanln(&input)
	return input
}

func ReadDate(prompt string) time.Time {
	fmt.Printf("%s (format: YYYY-MM-DD): ", prompt)
	var dateStr string
	fmt.Scanln(&dateStr)

	// Parse the date string
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		fmt.Println("Invalid date format. Please use YYYY-MM-DD")
		return ReadDate(prompt) // Recursively ask for input until valid
	}

	return date
}

package entries

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

package entries

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ConfirmEntry(t any) bool {
	fmt.Println(t)
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Confirm entry? (y/n): ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	return input == "y"
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var sum float64 = 0
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("🚀 Will sum numbers, if you fancy 🧠 🤯")
	fmt.Println("Type 'done' to stop 👇👇👇")
	for {
		fmt.Print("💁‍♀️ ")
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		if strings.ToLower(input) == "done" {
			break
		}

		number, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("🔞 🔞 🔞 Enter a number or 'done'.")
			continue
		}

		sum += number
	}

	fmt.Printf("🏦 🏦 🏦 The sum is: %.2f\n", sum)
}

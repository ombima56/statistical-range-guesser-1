package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	mathsskill "student/maths-skill"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var data []float64

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		num, err := strconv.ParseFloat(strings.TrimSpace(line), 64)
		if err != nil {
			fmt.Println("Error parsing input:", err)
			return
		}
		data = append(data, num)
		if len(data) > 1 {
			lowerLimit, upperLimit := mathsskill.CalculateRange(data)
			fmt.Printf("%d %d\n", int(lowerLimit), int(upperLimit))
		}
	}
}

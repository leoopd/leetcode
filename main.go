package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	n := 5
	solutionSlice := []string{}

	nString := fmt.Sprintf("%b", n)
	nSlice := strings.Split(nString, "")

	for i := 0; i < len(nSlice); i++ {
		if nSlice[i] == "1" {
			solutionSlice = append(solutionSlice, "0")
		} else {
			solutionSlice = append(solutionSlice, "1")
		}
	}
	solutionString := strings.Join(solutionSlice, "")
	solution, _ := strconv.ParseInt(solutionString, 2, 64)
	fmt.Println(solution)
}

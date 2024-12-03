package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func subreports(position1, position2 int, report []string) ([]string, []string) {
	//fmt.Println(len(report), " ", report)
	report1 := make([]string, len(report)-1)
	report2 := make([]string, len(report)-1)
	pos1 := 0
	pos2 := 0
	iterator := 0
	for _, value := range report {
		if pos1 != position1 || iterator != position1 {
			report1[pos1] = value
			pos1++
		}
		if pos2 != position2 || iterator != position2 {
			report2[pos2] = value
			pos2++
		}
		iterator++
	}
	return report1, report2
}

func check_report(values []string, tolerance bool) bool {

	//values := strings.Split(report, " ")
	diff := 0
	safe := true
	state := "indeterminate"
	//println(tolerance)
	//check the edges:
	edge1, _ := subreports(0, len(values)-1, values)
	//fmt.Println(tolerance)
	if tolerance {
		if check_report(edge1, false) {
			return true
		}
	}
L:
	for i := 0; i < (len(values) - 1); i++ {
		j := i + 1
		currentValue, _ := strconv.Atoi(values[i])
		NextValue, _ := strconv.Atoi(values[j])
		diff = currentValue - NextValue
		if math.Abs(float64(diff)) > 3 || diff == 0 {

			safe = false
			if !tolerance {
				break L
			}
		}
		//fmt.Println(state, diff)
		previous_state := state
		switch {
		case diff < 0:
			state = "increasing"
		case diff > 0:
			state = "decreasing"
		}
		//fmt.Println(state)
		if previous_state != state && i > 0 {
			safe = false
			if !tolerance {
				break L
			}
		}
		if !safe {
			repA, repB := subreports(i, j, values)
			if check_report(repA, false) {
				//println(repA)
				safe = true
				tolerance = false
				//fmt.Println(repA, "passes subRepA")
				break L
			}
			if check_report(repB, false) {
				//println(repB)
				safe = true
				tolerance = false
				//fmt.Println(repB, "passes subRepB")
				break L
			}
			fmt.Println("no valido:", values, "verificar: ", repA, " ", repB)
			break L

		}
	}
	return safe
}

func main() {

	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal("No Inpout File Found")
	}

	reports := strings.Split(string(content), "\n")
	safereports := 0
	for _, report := range reports {

		values := strings.Split(report, " ")
		if check_report(values, true) {
			safereports++
		}
	}
	fmt.Println("safe reports:", safereports)

}

package main

import (
	"flag"
	"fmt"
	"time"

	"advent.com/cmd/day01"
	"advent.com/cmd/day02"
)

var dayMap = map[string]func() int{
	"day01_gold":   day01.Day1Golden,
	"day01_silver": day01.Day1Silver,
	"day02_gold":   day02.Day2_Gold,
}

func main() {
	testMode := flag.Bool("t", false, "Enable performance test mode")
	flag.Parse()

	args := flag.Args()
	var dayToExec string

	if len(args) < 1 {
		dayToExec = menu()
	} else {
		dayToExec = args[0]
	}

	dayFunction := dayMap[dayToExec]

	if *testMode {
		testModeExecution(dayFunction)
	} else {
		fmt.Printf("result: %v", dayFunction())
	}
}

func menu() string {
	userSelect := false
	var selection string
	for !userSelect {
		fmt.Println("Select a day to execute:")
		for key := range dayMap {
			fmt.Printf("%v\n", key)
		}

		_, err := fmt.Scanln(&selection)
		userSelect = true

		if err != nil {
			userSelect = false
			continue
		}

		_, ok := dayMap[selection]
		if !ok {
			userSelect = false
			fmt.Println("Wrong selection")
			fmt.Println()
		}
	}
	return selection
}

func testModeExecution(dayFunction func() int) {
	start := time.Now()
	iterations := 10000
	var totalDuration time.Duration
	var worstCase time.Duration

	for i := 0; i < iterations; i++ {
		start := time.Now()
		dayFunction()
		duration := time.Since(start)

		totalDuration += duration
		if duration > worstCase {
			worstCase = duration
		}
	}

	average := totalDuration / time.Duration(iterations)

	fmt.Println("Test durations time: ", time.Since(start))
	fmt.Println("Average execution time:", average)
	fmt.Println("Wors execution time:", worstCase)
}

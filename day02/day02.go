package day02

import (
	"bufio"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

func Day2_Gold() int {
	file := getInputFile()
	defer file.Close()

	rl := getReports(file)

	return validReportCount(rl)
}

func validReportCount(l [][]string) int {
	s := 0
	for _, r := range l {
		if isValid(r) {
			s++
		}
	}
	return s
}

func isValid(r []string) bool {

	var nl []int
	for _, i := range r {
		n, _ := strconv.Atoi(i)
		nl = append(nl, n)
	}

	if nl[0] < nl[1] {
		return checkIncrease(nl)
	} else {
		return checkDecrease(nl)
	}
}

func checkIncrease(r []int) bool {
	for i := 0; i < len(r)-1; i++ {
		if r[i] >= r[i+1] || r[i+1]-r[i] > 3 {
			return false
		}
	}
	return true
}

func checkDecrease(r []int) bool {
	for i := 0; i < len(r)-1; i++ {
		if r[i] <= r[i+1] || r[i]-r[i+1] > 3 {
			return false
		}
	}
	return true
}

func getReports(file *os.File) [][]string {
	var reportList [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		reportList = append(reportList, strings.Split(scanner.Text(), " "))
	}

	if scanner.Err() != nil {
		return nil
	}
	return reportList

}

func getInputFile() *os.File {
	_, fullPath, _, ok := runtime.Caller(0)
	if !ok {
		return nil
	}
	currentPath := filepath.Dir(fullPath)

	inputName := "input.txt"
	absPath := filepath.Join(currentPath, inputName)
	file, e := os.Open(absPath)
	if e != nil {
		panic(e)
	}
	return file
}

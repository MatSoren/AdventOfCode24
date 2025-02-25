package day01

import (
	"bufio"
	"math"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strconv"
	"strings"
)

func Day1Silver() int {
	file := getInputFile()
	defer file.Close()

	scanner := bufio.NewScanner(file)
	countLines := CountLines(file)
	idList1, idList2 := InitLists(scanner, countLines)

	scoreMap := getScoreMap(&idList2)

	sum := 0
	for _, value := range idList1 {
		sum += value * scoreMap[value]
	}
	return sum
}

func getScoreMap(idList2 *[]int) map[int]int {
	scoreMap := map[int]int{}
	for _, id := range *idList2 {
		scoreMap[id] += 1
	}
	return scoreMap
}

func Day1Golden() int {
	file := getInputFile()
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := CountLines(file)
	idList1, idList2 := InitLists(scanner, lines)

	SortList(idList1)
	SortList(idList2)

	sum := 0
	for index := range len(idList1) {
		dif := int(math.Abs(float64(idList1[index] - idList2[index])))
		sum += dif
	}
	return sum
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

func GetIdOnPosition(idList []string, index uint8) int {
	id, e := strconv.Atoi(idList[index])
	if e != nil {
		panic(e)
	}
	return id
}

func CountLines(file *os.File) int {
	lines := 0
	defer file.Seek(0, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines++
	}

	return lines
}

func InitLists(scanner *bufio.Scanner, lines int) ([]int, []int) {
	idList1 := make([]int, lines)
	idList2 := make([]int, lines)
	index := 0
	for scanner.Scan() {
		inLineIds := strings.Split(scanner.Text(), "   ")
		idList1[index] = GetIdOnPosition(inLineIds, 0)
		idList2[index] = GetIdOnPosition(inLineIds, 1)
		index++
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return idList1, idList2
}

func SortList(list []int) {
	sort.Slice(list, func(i, j int) bool {
		return list[i] < list[j]
	})
}

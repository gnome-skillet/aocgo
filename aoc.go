package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func lineCounter(r io.Reader) (int, error) {
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}
}

func Valid(numbers []int) bool {
	fmt.Println("Valid?", numbers)
	subset := make([]int, len(numbers)-1)
	for i := range numbers {
		for j := 0; j < i; j++ {
			subset[j] = numbers[j]
		}
		for j := (i + 1); j < len(numbers); j++ {
			subset[j-1] = numbers[j]
		}
		fmt.Println("\tsubset", subset)
		if isValid(subset) {
			return true
		}
	}
	return false
}

func isValid(numbers []int) bool {
	diffs := Differences(numbers)
	return isSafe(diffs)
}

func PositiveNegative(number int) int {
	var diff int
	if number < 0 {
		diff = -1
	} else if number > 0 {
		diff = 1
	} else {
		diff = 0
	}
	return diff
}

func isSafe(numbers []int) bool {
	fmt.Println("differences", numbers)
	if len(numbers) <= 0 {
		return true
	}
	prev := PositiveNegative(numbers[0])
	if prev == 0 {
		return false
	}

	for _, v := range numbers {
		curr := PositiveNegative(v)
		if prev != curr || Abs(v) > 3 {
			return false
		}
	}
	fmt.Println("all safe")
	return true
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Differences(numbers []int) []int {
	differences := make([]int, len(numbers)-1)
	for i := range numbers[1:] {
		differences[i] = numbers[i] - numbers[i+1]
	}
	return differences
}

func main() {
	args := os.Args[1:]
	fmt.Println(args)
	fileName := args[0]

	f, err := os.Open(fileName)
	defer f.Close()
	if err != nil {
		panic(err)
	}
	nLines, err := lineCounter(f)
	if err != nil {
		panic(err)
	}

	f.Seek(0, 0)
	numbers := make([][]int, nLines)
	scanner := bufio.NewScanner(f)

	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		numberTokens := strings.Split(line, " ")
		row := make([]int, len(numberTokens))
		for j, v := range numberTokens {
			row[j], err = strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
		}
		numbers[i] = row
		i += 1
	}
	nValid := 0
	for y := range numbers {
		fmt.Println("evaluate", numbers[y])
		if Valid(numbers[y]) {
			fmt.Println(numbers[y], "is valid")
			nValid += 1
		} else {
			fmt.Println(numbers[y], "is not valid")
		}
	}
	fmt.Println("there are", nValid, "valid sequences")
}

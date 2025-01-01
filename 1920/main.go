package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	n := readInt(reader)
	nums := readArray(reader)
	isContain := make(map[int]struct{}, n)
	for _, num := range nums {
		isContain[num] = struct{}{}
	}

	m := readInt(reader)
	targets := readArray(reader)
	results := make([]byte, 0, m)
	for _, target := range targets {
		if _, found := isContain[target]; found {
			results = append(results, '1', '\n')
		} else {
			results = append(results, '0', '\n')
		}
	}

	writer.WriteString(string(results))
}

func readInt(reader *bufio.Reader) int {
	line, _ := reader.ReadString('\n')
	val, _ := strconv.Atoi(strings.TrimSpace(line))
	return val
}

func readArray(reader *bufio.Reader) []int {
	line, _ := reader.ReadString('\n')
	parts := strings.Fields(line)

	arr := make([]int, len(parts))
	for i, part := range parts {
		arr[i], _ = strconv.Atoi(part)
	}

	return arr
}

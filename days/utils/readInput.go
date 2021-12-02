package utils

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadInput(day int) ([]string, error) {
	data, err := os.ReadFile(fmt.Sprintf("input/day%d.txt", day))

	if err != nil {
		return nil, err
	}

	strData := string(data)

	return strings.Split(strData, "\n"), nil
}

func InputToIntLines(data []string) ([]int, error) {
	var buf []int

	for _, line := range data {
		converted, err := strconv.Atoi(line)

		if err != nil {
			return nil, err
		}

		buf = append(buf, converted)
	}

	return buf, nil
}

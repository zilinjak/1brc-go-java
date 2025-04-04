package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"slices"
	"sort"
	"strings"
)

// Round to 1 decimal place, 0.05 rounds up to 0.1, 0.04 rounds down to 0.0
func round(x float64) float64 {
	return math.Floor((x+0.05)*10) / 10
}

func parseStringToInt(s string) (int64, error) {

	// negative
	negative := false
	if len(s) == 5 {
		s = s[1:5]
		negative = true
	}

	var result int64 = 0
	switch size := len(s); size {
	case 3:
		result = int64(s[0])*10 + int64(s[2])
	case 4:
		result = int64(s[0])*100 + int64(s[1])*10 + int64(s[3])
	default:
		return 0, errors.New("invalid value in input")
	}
	if negative {
		result = result * -1
	}
	return result, nil
}

func readFile(filePath string) map[string][]int64 {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var resultMap = make(map[string][]int64, 10000)

	for scanner.Scan() {
		text := scanner.Text()
		splitString := strings.SplitN(text, ";", 2)
		place, value := splitString[0], splitString[1]
		valueFloat, _ := parseStringToInt(value)
		if _, ok := resultMap[place]; ok {
			// just update value
			resultMap[place] = append(resultMap[place], valueFloat)
		} else {
			// allocate and update
			// possible better solution
			// resultMap[place] = []float64{valueFloat}
			resultMap[place] = make([]int64, 0, 1)
			resultMap[place] = append(resultMap[place], valueFloat)
		}
	}
	return resultMap
}

func printResults(result map[string][]int64) {
	// Extract keys and sort them
	keys := make([]string, 0, len(result))
	for k := range result {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	// Iterate over sorted keys
	fmt.Print("{")
	for i, k := range keys {
		var sum int64 = 0
		for _, temp := range result[k] {
			sum += temp
		}
		var totalSize = float64(len(result[k]))
		minTemp := slices.Min(result[k])
		maxTemp := slices.Max(result[k])
		avgTemp := round(round(sum) / totalSize)
		if i == len(keys)-1 {
			fmt.Printf("%s=%.1f/%.1f/%.1f}\n", k, minTemp, avgTemp, maxTemp)
		} else {
			fmt.Printf("%s=%.1f/%.1f/%.1f, ", k, minTemp, avgTemp, maxTemp)
		}
	}

}

func main() {
	var resultMap map[string][]int64 = readFile("measurements.txt")
	printResults(resultMap)
}

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

// Round to 1 decimal place, 0.05 rounds up to 0.1, 0.04 rounds down to 0.0
func round(x float64) float64 {
	return math.Floor((x+0.05)*10) / 10
}

func readFile(filePath string) map[string][]float64 {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var resultMap = make(map[string][]float64, 10000)

	for scanner.Scan() {
		text := scanner.Text()
		splitString := strings.SplitN(text, ";", 2)
		place, value := splitString[0], splitString[1]
		valueFloat, _ := strconv.ParseFloat(strings.TrimSpace(value), 32)
		if _, ok := resultMap[place]; ok {
			// just update value
			resultMap[place] = append(resultMap[place], valueFloat)
		} else {
			// allocate and update
			// possible better solution
			// resultMap[place] = []float64{valueFloat}
			resultMap[place] = make([]float64, 0, 1)
			resultMap[place] = append(resultMap[place], valueFloat)
		}
	}
	return resultMap
}

func printResults(result map[string][]float64) {
	// Extract keys and sort them
	keys := make([]string, 0, len(result))
	for k := range result {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	// Iterate over sorted keys
	fmt.Print("{")
	for i, k := range keys {
		var sum float64 = 0.0
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
	var resultMap map[string][]float64 = readFile("measurements.txt")
	printResults(resultMap)
}

package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func parseInts(values []string) []int {
	var result []int
	for _, v := range values {
		num, _ := strconv.Atoi(v)
		result = append(result, num)
	}

	return result
}

func parseIntsFromLines(lines []string) [][]int {
	var result [][]int
	for _, line := range lines {
		result = append(result, parseInts(strings.Fields(line)))
	}

	return result
}

func generateMaps(data []string) [][][]int {
	var mapArray = make([][][]int, 7)

	for _, section := range data {
		lines := strings.Split(section, "\n")
		switch {
		case strings.HasPrefix(lines[0], "seed-to-soil map"):
			mapArray[0] = parseIntsFromLines(lines[1:])

		case strings.HasPrefix(lines[0], "soil-to-fertilizer map"):
			mapArray[1] = parseIntsFromLines(lines[1:])

		case strings.HasPrefix(lines[0], "fertilizer-to-water map"):
			mapArray[2] = parseIntsFromLines(lines[1:])

		case strings.HasPrefix(lines[0], "water-to-light map"):
			mapArray[3] = parseIntsFromLines(lines[1:])

		case strings.HasPrefix(lines[0], "light-to-temperature map"):
			mapArray[4] = parseIntsFromLines(lines[1:])

		case strings.HasPrefix(lines[0], "temperature-to-humidity map"):
			mapArray[5] = parseIntsFromLines(lines[1:])

		case strings.HasPrefix(lines[0], "humidity-to-location map"):
			mapArray[6] = parseIntsFromLines(lines[1:])
			mapArray[6] = mapArray[6][:len(mapArray[6])-1]
		}
	}

	return mapArray
}

func getSoilForSeed(seedToSoil [][]int, seed int) int {
	for _, row := range seedToSoil {
		if row[1] <= seed && seed < row[1]+row[2] {
			return row[0] + seed - row[1]
		}
	}

	return seed
}

func getfertilizerForSoil(soilToFertilizer [][]int, soil int) int {
	for _, row := range soilToFertilizer {
		if row[1] <= soil && soil <= row[1]+row[2] {
			return row[0] + soil - row[1]
		}
	}

	return soil
}

func getWaterForFertilizer(fertilizerToWater [][]int, fertilizer int) int {
	for _, row := range fertilizerToWater {
		if row[1] <= fertilizer && fertilizer <= row[1]+row[2] {
			return row[0] + fertilizer - row[1]
		}
	}

	return fertilizer
}

func getLightForWater(waterToLight [][]int, water int) int {
	for _, row := range waterToLight {
		if row[1] <= water && water <= row[1]+row[2] {
			return row[0] + water - row[1]
		}
	}

	return water
}

func getTemperatureForLight(lightToTemperature [][]int, light int) int {
	for _, row := range lightToTemperature {
		if row[1] <= light && light <= row[1]+row[2] {
			return row[0] + light - row[1]
		}
	}

	return light
}

func getHumidityForTemperature(temperatureToHumidity [][]int, temperature int) int {
	for _, row := range temperatureToHumidity {
		if row[1] <= temperature && temperature <= row[1]+row[2] {
			return row[0] + temperature - row[1]
		}
	}

	return temperature
}

func getLocationForHumidity(humidityToLocation [][]int, humidity int) int {
	for _, row := range humidityToLocation {
		if row[1] <= humidity && humidity <= row[1]+row[2] {
			return row[0] + humidity - row[1]
		}
	}

	return humidity
}

func partOne(seeds []int, seedToSoil, soilToFertilizer, fertilizerToWater, waterToLight, lightToTemperature, temperatureToHumidity, humidityToLocation [][]int) int {
	lowestLocation := math.MaxInt

	for _, seed := range seeds {
		soil := getSoilForSeed(seedToSoil, seed)
		fertilizer := getfertilizerForSoil(soilToFertilizer, soil)
		water := getWaterForFertilizer(fertilizerToWater, fertilizer)
		light := getLightForWater(waterToLight, water)
		temperature := getTemperatureForLight(lightToTemperature, light)
		humidity := getHumidityForTemperature(temperatureToHumidity, temperature)
		location := getLocationForHumidity(humidityToLocation, humidity)

		if location < lowestLocation {
			lowestLocation = location
		}
	}

	return lowestLocation
}

func generateSeedRange(seeds []int) [][]int {
	var seedRange [][]int

	for i := 0; i < len(seeds); i += 2 {
		seedRange = append(seedRange, []int{seeds[i], seeds[i+1]})
	}
	return seedRange
}

func partTwo(seeds []int, seedToSoil, soilToFertilizer, fertilizerToWater, waterToLight, lightToTemperature, temperatureToHumidity, humidityToLocation [][]int) int {
	lowestLocation := math.MaxInt

	arraySeedRange := generateSeedRange(seeds)

	for _, seedRange := range arraySeedRange {
		for seed := seedRange[0]; seed < seedRange[0]+seedRange[1]; seed++ {
			soil := getSoilForSeed(seedToSoil, seed)
			fertilizer := getfertilizerForSoil(soilToFertilizer, soil)
			water := getWaterForFertilizer(fertilizerToWater, fertilizer)
			light := getLightForWater(waterToLight, water)
			temperature := getTemperatureForLight(lightToTemperature, light)
			humidity := getHumidityForTemperature(temperatureToHumidity, temperature)
			location := getLocationForHumidity(humidityToLocation, humidity)

			if location < lowestLocation {
				lowestLocation = location
			}

		}
	}

	return lowestLocation
}

func main() {
	file, err := os.ReadFile("input")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Cannot open 'input'")
		return
	}

	fileArray := strings.Split(string(file), "\n\n")
	var seeds = parseInts(strings.Fields(strings.Split(fileArray[0], ":")[1]))
	var rangeArray [][][]int = generateMaps(fileArray)

	// Part 1
	fmt.Printf("Part 1: %d\n", partOne(seeds, rangeArray[0], rangeArray[1], rangeArray[2], rangeArray[3], rangeArray[4], rangeArray[5], rangeArray[6]))

	// Part 2
	fmt.Printf("Part 2: %d\n", partTwo(seeds, rangeArray[0], rangeArray[1], rangeArray[2], rangeArray[3], rangeArray[4], rangeArray[5], rangeArray[6]))
}

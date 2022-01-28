package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
	"weave/cost"
	"weave/readings"
)

func main() {
	fmt.Println("Initializing energy cost calculations")

	csvPath := "readings_unsorted_bad.csv"
	input := readCsvFile(csvPath)

	fmt.Println("Collecting readings...")
	r := new(readings.Readings)
	readingsData := r.GetReadings(input)
	fmt.Println("Readings collected")

	fmt.Println("Calculating usage costs...")
	c := new(cost.Cost)
	output := c.CalculateCost(readingsData)
	fmt.Println("Calculations complete")

	resultFile := writeDataToCsv(output)
	fmt.Println("Calculation Results saved to file: " + resultFile)
	fmt.Println("End")
}

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}

func writeDataToCsv(data map[int]float64) string {

	//create unix timestamped csv
	fileName := "cost_" + fmt.Sprint(time.Now().Unix()) + ".csv"
	csvFile, err := os.Create(fileName)

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	defer csvFile.Close()

	//initiate csv writer
	csvWriter := csv.NewWriter(csvFile)
	defer csvWriter.Flush()

	//create header row
	dataRow := []string{
		"id", "cost",
	}

	//write header row
	if err := csvWriter.Write(dataRow); err != nil {
		log.Fatalln("error writing record to file", err)
	}

	//write data entries to csv
	for key, value := range data {
		dataRow = []string{
			strconv.Itoa(key), fmt.Sprintf("%.2f", value),
		}

		if err := csvWriter.Write(dataRow); err != nil {
			log.Fatalln("error writing record to file", err)
		}
	}

	return fileName
}

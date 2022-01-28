package cost

import (
	"fmt"
	"math"
	"sort"
	"time"
	"weave/readings"
)

type Cost struct {
}

func (c Cost) CalculateCost(readingsData readings.Readings) map[int]float64 {

	//create cost map
	costData := make(map[int]float64)

	//loop only till 2nd last reading since usage for last reading cannot be calculated
	for i := 0; i < len(readingsData.Data); i++ {

		//create variables for each reading
		meteringPointId := readingsData.Data[i].MeteringPointId
		meteringType := readingsData.Data[i].MeteringType
		meteringReading := readingsData.Data[i].MeteringReading
		createdAt := readingsData.Data[i].CreatedAt

		//Set map index if not set
		if _, ok := costData[meteringPointId]; ok {
			//do nothing
		} else {
			costData[meteringPointId] = 0
		}

		//calculate usage
		usage := 0
		correctUsageFlag := false
		nextCorrectIndex := 0

		for j := 1; j+i < len(readingsData.Data); j++ {
			//check if same meter id
			if readingsData.Data[i+j].MeteringPointId == meteringPointId {
				usage = readingsData.Data[i+j].MeteringReading - meteringReading

				if usage >= 0 && usage <= j*100 {
					correctUsageFlag = true
					nextCorrectIndex = j + i
					break
				} else {
					fmt.Println("Incorrect usage detected for meter_point_id " + fmt.Sprint(readingsData.Data[i+j].MeteringPointId) + ": " + fmt.Sprint(usage))
				}
			} else {
				//data has moved to next meter id and no further usage can be determined for this meter and therefore according to the data sort
				//done during the fetch of the readings the cost calculations can move on to the next meter and start from there.
				//set loop index to start of next meter
				i = j + i - 1
				break
			}
		}

		//if no further correct usage can be determined
		if !correctUsageFlag {
			continue
		} else {
			//skip calculation process ahead to a correct reading if next reading was correct loop index won't be affected
			i = nextCorrectIndex - 1
		}

		//create variables for cost calculations
		var kWh float64
		var electricityMod float64 = 1000
		var gasMod float64 = 9.769
		var cost float64

		readingTime := time.Unix(int64(createdAt), 0)
		readingWeekday := readingTime.Weekday().String()
		readingDuringPeak := isInTimeRange(readingTime)

		//calculate cost as per meter type
		switch meteringType {
		//electricity
		case 1:
			kWh = float64(usage) / electricityMod

			//calculate cost as per weekday
			if readingWeekday == "Saturday" || readingWeekday == "Sunday" {
				cost = 0.18 * kWh
			} else {
				if readingDuringPeak {
					cost = 0.20 * kWh
				} else {
					cost = 0.18 * kWh
				}
			}
		//gas
		case 2:
			kWh = float64(usage) * gasMod
			cost = 0.06 * kWh
		}

		//round to 2 decimal places
		costData[meteringPointId] += math.Round(cost*100) / 100
	}

	costData = sortResultsByMeterId(costData)

	return costData
}

//Function to sort cost data by meter id
func sortResultsByMeterId(data map[int]float64) map[int]float64 {
	sortedData := make(map[int]float64)

	keys := make([]int, len(data))
	i := 0
	for k := range data {
		keys[i] = k
		i++
	}

	sort.Ints(keys)

	for _, k := range keys {
		sortedData[k] = data[k]
	}

	return sortedData
}

//Function to convert a string to a time variable for time conditions
func stringToTime(str string) time.Time {
	tm, err := time.Parse(time.Kitchen, str)
	if err != nil {
		fmt.Println("Failed to decode time:", err)
	}
	return tm
}

//Function to check if current meter reading is within a specific time range
func isInTimeRange(t time.Time) bool {

	//set between time
	startTimeString := "07:00AM"
	endTimeString := "11:00PM"

	timeNowString := t.Format(time.Kitchen)

	//set time variables for range check
	timeNow := stringToTime(timeNowString)
	start := stringToTime(startTimeString)
	end := stringToTime(endTimeString)

	if timeNow.Before(start) {
		return false
	}

	if timeNow.Before(end) {
		return true
	}

	return false
}

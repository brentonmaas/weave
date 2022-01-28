package readings

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"weave/readings/reading"
)

type Readings struct {
	Data []reading.Reading
}

// ByIdTime implements sort.Interface based on the MeteringPointId field and CreatedAt field.
type ByIdTime []reading.Reading

func (a ByIdTime) Len() int      { return len(a) }
func (a ByIdTime) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByIdTime) Less(i, j int) bool {
	if a[i].MeteringPointId == a[j].MeteringPointId {
		return a[i].CreatedAt < a[j].CreatedAt
	} else {
		return a[i].MeteringPointId < a[j].MeteringPointId
	}
}

func (r Readings) GetReadings(input [][]string) Readings {

	var data []reading.Reading

	//loop through readings and save to array
	for i := 0; i < len(input); i++ {

		if i != 0 {

			meteringPointId, err := strconv.Atoi(input[i][0])
			if err != nil {
				fmt.Println(err)
				os.Exit(2)
			}

			meteringType, err := strconv.Atoi(input[i][1])
			if err != nil {
				fmt.Println(err)
				os.Exit(2)
			}

			meteringReading, err := strconv.Atoi(input[i][2])
			if err != nil {
				fmt.Println(err)
				os.Exit(2)
			}

			createdAt, err := strconv.Atoi(input[i][3])
			if err != nil {
				fmt.Println(err)
				os.Exit(2)
			}

			//create reading object and append to array
			readingData := reading.Reading{
				MeteringPointId: meteringPointId,
				MeteringType:    meteringType,
				MeteringReading: meteringReading,
				CreatedAt:       createdAt,
			}

			data = append(data, readingData)
		}

	}

	sort.Sort(ByIdTime(data))

	r.Data = data

	return r
}

package readings_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"weave/readings"
	"weave/readings/reading"
)

func TestGetReadings(t *testing.T) {
	input := [][]string{
		{"metering_point_id", "type", "reading", "created_at"},
		{"1", "1", "166606", "1415963700"},
		{"1", "1", "166694", "1415964600"},
		{"1", "1", "166714", "1415965500"},
	}

	expected := readings.Readings{
		Data: []reading.Reading{
			{MeteringPointId: 1, MeteringType: 1, MeteringReading: 166606, CreatedAt: 1415963700},
			{MeteringPointId: 1, MeteringType: 1, MeteringReading: 166694, CreatedAt: 1415964600},
			{MeteringPointId: 1, MeteringType: 1, MeteringReading: 166714, CreatedAt: 1415965500},
		},
	}

	r := new(readings.Readings)

	assert.Equal(t, expected, r.GetReadings(input))

}

func TestUnsortedGetReadings(t *testing.T) {
	input := [][]string{
		{"metering_point_id", "type", "reading", "created_at"},
		{"1", "1", "166606", "1415963700"},
		{"2", "2", "166606", "1415963700"},
		{"1", "1", "166694", "1415964600"},
		{"2", "2", "166694", "1415964600"},
		{"1", "1", "166714", "1415965500"},
		{"1", "1", "166713", "1415966400"},
		{"2", "2", "166714", "1415965500"},
		{"1", "1", "166999", "1415967300"},
		{"1", "1", "166863", "1415968200"},
	}

	expected := readings.Readings{
		Data: []reading.Reading{
			{MeteringPointId: 1, MeteringType: 1, MeteringReading: 166606, CreatedAt: 1415963700},
			{MeteringPointId: 1, MeteringType: 1, MeteringReading: 166694, CreatedAt: 1415964600},
			{MeteringPointId: 1, MeteringType: 1, MeteringReading: 166714, CreatedAt: 1415965500},
			{MeteringPointId: 1, MeteringType: 1, MeteringReading: 166713, CreatedAt: 1415966400},
			{MeteringPointId: 1, MeteringType: 1, MeteringReading: 166999, CreatedAt: 1415967300},
			{MeteringPointId: 1, MeteringType: 1, MeteringReading: 166863, CreatedAt: 1415968200},
			{MeteringPointId: 2, MeteringType: 2, MeteringReading: 166606, CreatedAt: 1415963700},
			{MeteringPointId: 2, MeteringType: 2, MeteringReading: 166694, CreatedAt: 1415964600},
			{MeteringPointId: 2, MeteringType: 2, MeteringReading: 166714, CreatedAt: 1415965500},
		},
	}

	r := new(readings.Readings)

	assert.Equal(t, expected, r.GetReadings(input))

}

func TestUnsortedBadGetReadings(t *testing.T) {
	input := [][]string{
		{"metering_point_id", "type", "reading", "created_at"},
		{"1", "1", "166606", "1415963700"},
		{"2", "2", "166606", "1415963700"},
		{"1", "1", "166694", "1415964600"},
		{"2", "2", "166694", "1415964600"},
		{"1", "1", "166714", "1415965500"},
		{"1", "1", "166713", "1415966400"},
		{"2", "2", "166714", "1415965500"},
		{"1", "1", "166999", "1415967300"},
		{"1", "1", "167200", "1415968200"},
	}

	expected := readings.Readings{
		Data: []reading.Reading{
			{MeteringPointId: 1, MeteringType: 1, MeteringReading: 166606, CreatedAt: 1415963700},
			{MeteringPointId: 1, MeteringType: 1, MeteringReading: 166694, CreatedAt: 1415964600},
			{MeteringPointId: 1, MeteringType: 1, MeteringReading: 166714, CreatedAt: 1415965500},
			{MeteringPointId: 1, MeteringType: 1, MeteringReading: 166713, CreatedAt: 1415966400},
			{MeteringPointId: 1, MeteringType: 1, MeteringReading: 166999, CreatedAt: 1415967300},
			{MeteringPointId: 1, MeteringType: 1, MeteringReading: 167200, CreatedAt: 1415968200},
			{MeteringPointId: 2, MeteringType: 2, MeteringReading: 166606, CreatedAt: 1415963700},
			{MeteringPointId: 2, MeteringType: 2, MeteringReading: 166694, CreatedAt: 1415964600},
			{MeteringPointId: 2, MeteringType: 2, MeteringReading: 166714, CreatedAt: 1415965500},
		},
	}

	r := new(readings.Readings)

	assert.Equal(t, expected, r.GetReadings(input))

}

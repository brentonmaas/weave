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
			{1, 1, 166606, 1415963700},
			{1, 1, 166694, 1415964600},
			{1, 1, 166714, 1415965500},
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
			{1, 1, 166606, 1415963700},
			{1, 1, 166694, 1415964600},
			{1, 1, 166714, 1415965500},
			{1, 1, 166713, 1415966400},
			{1, 1, 166999, 1415967300},
			{1, 1, 166863, 1415968200},
			{2, 2, 166606, 1415963700},
			{2, 2, 166694, 1415964600},
			{2, 2, 166714, 1415965500},
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
			{1, 1, 166606, 1415963700},
			{1, 1, 166694, 1415964600},
			{1, 1, 166714, 1415965500},
			{1, 1, 166713, 1415966400},
			{1, 1, 166999, 1415967300},
			{1, 1, 167200, 1415968200},
			{2, 2, 166606, 1415963700},
			{2, 2, 166694, 1415964600},
			{2, 2, 166714, 1415965500},
		},
	}

	r := new(readings.Readings)

	assert.Equal(t, expected, r.GetReadings(input))

}

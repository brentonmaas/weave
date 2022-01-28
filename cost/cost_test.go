package cost_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"weave/cost"
	"weave/readings"
	"weave/readings/reading"
)

func TestCalculateCost1(t *testing.T) {

	input := readings.Readings{
		Data: []reading.Reading{
			{1, 1, 166606, 1415963700},
			{1, 1, 166694, 1415964600},
			{1, 1, 166714, 1415965500},
		},
	}

	expected := map[int]float64{
		1: 0.02,
	}

	c := new(cost.Cost)

	assert.Equal(t, expected, c.CalculateCost(input))

}

func TestCalculateCost2(t *testing.T) {

	input := readings.Readings{
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

	expected := map[int]float64{
		1: 0.05,
		2: 63.3,
	}

	c := new(cost.Cost)

	assert.Equal(t, expected, c.CalculateCost(input))

}

func TestCalculateCost3(t *testing.T) {

	input := readings.Readings{
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

	expected := map[int]float64{
		1: 0.02,
		2: 63.3,
	}

	c := new(cost.Cost)

	assert.Equal(t, expected, c.CalculateCost(input))

}

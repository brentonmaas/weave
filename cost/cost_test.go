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
			{MeteringPointId: 1, MeteringType: 1, MeteringReading: 166606, CreatedAt: 1415963700},
			{MeteringPointId: 1, MeteringType: 1, MeteringReading: 166694, CreatedAt: 11415964600},
			{MeteringPointId: 1, MeteringType: 1, MeteringReading: 166714, CreatedAt: 11415965500},
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
			{MeteringPointId: 1, MeteringType: 1, MeteringReading: 166606, CreatedAt: 11415963700},
			{MeteringPointId: 1, MeteringType: 1, MeteringReading: 166694, CreatedAt: 11415964600},
			{MeteringPointId: 1, MeteringType: 1, MeteringReading: 166714, CreatedAt: 11415965500},
			{MeteringPointId: 1, MeteringType: 1, MeteringReading: 166713, CreatedAt: 11415966400},
			{MeteringPointId: 1, MeteringType: 1, MeteringReading: 166999, CreatedAt: 11415967300},
			{MeteringPointId: 1, MeteringType: 1, MeteringReading: 166863, CreatedAt: 11415968200},
			{MeteringPointId: 2, MeteringType: 2, MeteringReading: 166606, CreatedAt: 11415963700},
			{MeteringPointId: 2, MeteringType: 2, MeteringReading: 166694, CreatedAt: 11415964600},
			{MeteringPointId: 2, MeteringType: 2, MeteringReading: 166714, CreatedAt: 11415965500},
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
			{MeteringPointId: 1, MeteringType: 1, MeteringReading: 166606, CreatedAt: 11415963700},
			{MeteringPointId: 1, MeteringType: 1, MeteringReading: 166694, CreatedAt: 11415964600},
			{MeteringPointId: 1, MeteringType: 1, MeteringReading: 166714, CreatedAt: 11415965500},
			{MeteringPointId: 1, MeteringType: 1, MeteringReading: 166713, CreatedAt: 11415966400},
			{MeteringPointId: 1, MeteringType: 1, MeteringReading: 166999, CreatedAt: 11415967300},
			{MeteringPointId: 1, MeteringType: 1, MeteringReading: 167200, CreatedAt: 11415968200},
			{MeteringPointId: 2, MeteringType: 2, MeteringReading: 166606, CreatedAt: 11415963700},
			{MeteringPointId: 2, MeteringType: 2, MeteringReading: 166694, CreatedAt: 11415964600},
			{MeteringPointId: 2, MeteringType: 2, MeteringReading: 166714, CreatedAt: 11415965500},
		},
	}

	expected := map[int]float64{
		1: 0.02,
		2: 63.3,
	}

	c := new(cost.Cost)

	assert.Equal(t, expected, c.CalculateCost(input))

}

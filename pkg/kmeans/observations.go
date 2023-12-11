package kmeans

import (
	"fmt"
	"math"
)

type Coordinates []float64

type Observation interface {
	Coordinates() Coordinates           // return kmeans.Coordinates
	Distance(point Coordinates) float64 // receives kmeans.Coordinates returns float64
}

type Observations []Observation

func (c Coordinates) Coordinates() Coordinates {
	return Coordinates(c)
}

func (c Coordinates) Distance(p2 Coordinates) float64 {
	var r float64
	for i, v := range c {
		r += math.Pow(v-p2[i], 2)
	}
	return r
}

func (c Observations) Center() (Coordinates, error) {
	var l = len(c)
	if l == 0 {
		return nil, fmt.Errorf("there is no mean for an empty set of points")
	}

	cc := make([]float64, len(c[0].Coordinates()))
	for _, point := range c {
		for j, v := range point.Coordinates() {
			cc[j] += v
		}
	}

	var mean Coordinates
	for _, v := range cc {
		mean = append(mean, v/float64(l))
	}
	return mean, nil
}

func AverageDistance(o Observation, observations Observations) float64 {
	var (
		d float64
		l int
	)

	for _, observation := range observations {
		dist := o.Distance(observation.Coordinates())
		if dist == 0 {
			continue
		}

		l++
		d += dist
	}

	if l == 0 {
		return 0
	}
	return d / float64(l)
}

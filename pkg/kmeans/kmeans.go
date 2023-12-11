package kmeans

import (
	"fmt"
	"math/rand"
)

const (
	DefaultDeltaThreshold     = 0.01
	DefaultIterationThreshold = 100
)

type Kmeans struct {
	deltaThreshold     float64
	iterationThreshold int
}

func NewKmeansWithOptions(deltaThreshold float64, iterationThreshold int) (*Kmeans, error) {
	if deltaThreshold <= 0.0 || deltaThreshold >= 1.0 {
		return nil, fmt.Errorf("delta threshold is out of bounds (must be > 0.0 and < 1.0)")
	}

	if iterationThreshold < 0 {
		return nil, fmt.Errorf("iteration threshold is out of bounds (must be >= 0)")
	}

	return &Kmeans{
		deltaThreshold:     deltaThreshold,
		iterationThreshold: iterationThreshold,
	}, nil
}

func New() *Kmeans {
	km, _ := NewKmeansWithOptions(DefaultDeltaThreshold, DefaultIterationThreshold)
	return km
}

func (m *Kmeans) Partition(dataset Observations, k int, seed int64) (Clusters, error) {
	if k > len(dataset) {
		return Clusters{}, fmt.Errorf("the size of the data set must at least equal to k (%d)", k)
	}

	cc, err := NewClusters(seed, k, dataset)
	if err != nil {
		return cc, err
	}

	points := make([]int, len(dataset))
	changes := 1

	for i := 0; changes > 0; i++ {
		changes = 0
		cc.Reset()

		for p, point := range dataset {
			ci := cc.Nearest(point)
			cc[ci].Append(point)
			if points[p] != ci {
				points[p] = ci
				changes++
			}
		}

		for ci := 0; ci < len(cc); ci++ {
			if len(cc[ci].Observations) == 0 {
				var ri int

				for {
					ri = rand.Intn(len(dataset))
					if len(cc[points[ri]].Observations) > 1 {
						break
					}
				}

				cc[ci].Append(dataset[ri])
				points[ri] = ci

				changes = len(dataset)
			}
		}

		if changes > 0 {
			cc.Recenter()
		}

		if i == m.iterationThreshold ||
			changes < int(float64(len(dataset))*m.deltaThreshold) {
			break
		}
	}

	return cc, nil
}

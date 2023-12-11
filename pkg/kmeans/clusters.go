package kmeans

import (
	"fmt"
	"math/rand"
)

type Cluster struct {
	Center       Coordinates
	Observations Observations
}

type Clusters []Cluster

func NewClusters(seed int64, k int, dataset Observations) (Clusters, error) {
	var c Clusters

	if len(dataset) == 0 || len(dataset[0].Coordinates()) == 0 {
		return c, fmt.Errorf("there must be at least one dimension in the data set")
	}

	if k == 0 {
		return c, fmt.Errorf("k must be greater than 0")
	}

	for i := 0; i < k; i++ {
		var p Coordinates
		for j := 0; j < len(dataset[0].Coordinates()); j++ {
			p = append(p, rand.Float64())
		}

		c = append(c, Cluster{
			Center: p,
		})
	}
	return c, nil
}

func (c *Cluster) Append(point Observation) {
	c.Observations = append(c.Observations, point)
}

func (c *Clusters) Nearest(point Observation) int {
	var ci int
	dist := -1.0

	// Find the nearest cluster for this data point
	for i, cluster := range *c {
		d := point.Distance(cluster.Center)
		if dist < 0 || d < dist {
			dist = d
			ci = i
		}
	}

	return ci
}

func (c *Clusters) Neighbour(point Observation, fromCluster int) (int, float64) {
	var d float64
	nc := -1

	for i, cluster := range *c {
		if i == fromCluster {
			continue
		}

		cd := AverageDistance(point, cluster.Observations)
		if nc < 0 || cd < d {
			nc = i
			d = cd
		}
	}

	return nc, d
}

func (c *Cluster) Recenter() {
	center, err := c.Observations.Center()
	if err != nil {
		return
	}

	c.Center = center
}

func (c Clusters) Recenter() {
	for i := 0; i < len(c); i++ {
		c[i].Recenter()
	}
}

func (c Clusters) Reset() {
	for i := 0; i < len(c); i++ {
		c[i].Observations = Observations{}
	}
}

func (c Cluster) PointsInDimension(n int) Coordinates {
	var v []float64
	for _, p := range c.Observations {
		v = append(v, p.Coordinates()[n])
	}
	return v
}

func (c Clusters) CentersInDimension(n int) Coordinates {
	var v []float64
	for _, cl := range c {
		v = append(v, cl.Center[n])
	}
	return v
}

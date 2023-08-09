package kmeans

import (
	"errors"
	"fmt"
	"image"
	"sort"

	"github.com/muesli/clusters"
	"github.com/muesli/kmeans"
)

func GetDominantColorsHex(rgba image.Image, kgroups int, tolerance float64) (*[]string, error) {
	if kgroups < 3 {
		return nil, errors.New("kgroups can't be less than 3")
	}

	boundsX := float64(rgba.Bounds().Max.Y)
	boundsY := float64(rgba.Bounds().Max.Y)

	iterationsX := int((boundsX * (tolerance * 100)) / boundsX)
	iterationsY := int((boundsY * (tolerance * 100)) / boundsY)

	var ob clusters.Observations

	for y := rgba.Bounds().Min.Y; y < rgba.Bounds().Max.Y; y += iterationsX {
		for x := rgba.Bounds().Min.X; x < rgba.Bounds().Max.X; x += iterationsY {
			c := rgba.At(x, y)
			r, g, b, _ := c.RGBA()
			ob = append(ob, clusters.Coordinates{
				float64(r >> 8),
				float64(g >> 8),
				float64(b >> 8),
			})
		}
	}

	km := kmeans.New()
	kmClusters, _ := km.Partition(ob, kgroups)

	dominantColors := make([]string, kgroups)

	sort.SliceStable(kmClusters, func(i, j int) bool {
		return len(kmClusters[i].Observations) > len(kmClusters[j].Observations)
	})

	for i, cluster := range kmClusters {
		r := uint8(cluster.Center[0])
		g := uint8(cluster.Center[1])
		b := uint8(cluster.Center[2])

		dominantColors[i] = fmt.Sprintf("#%02x%02x%02x", r, g, b)
	}

	return &dominantColors, nil
}

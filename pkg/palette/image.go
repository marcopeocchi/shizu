package palette

import (
	"errors"
	"image"
	"sort"

	"github.com/lucasb-eyer/go-colorful"
	"github.com/muesli/clusters"
	"github.com/muesli/kmeans"
)

func Get(rgba image.Image, kgroups int, tolerance float64) (*[]string, error) {
	if kgroups < 3 {
		return nil, errors.New("kgroups can't be less than 3")
	}

	var (
		boundY      = float64(rgba.Bounds().Max.Y)
		boundX      = float64(rgba.Bounds().Max.X)
		iterationsY = int((boundY * (tolerance * 100)) / boundY)
		iterationsX = int((boundX * (tolerance * 100)) / boundX)

		obs clusters.Observations
	)

	for y := rgba.Bounds().Min.Y; y < rgba.Bounds().Max.Y; y += iterationsY {
		for x := rgba.Bounds().Min.X; x < rgba.Bounds().Max.X; x += iterationsX {
			linear := rgba.At(x, y)

			c, ok := colorful.MakeColor(linear)
			if !ok {
				return nil, errors.New("error converting color")
			}

			l, a, b := c.Lab()

			obs = append(obs, clusters.Coordinates{l, a, b})
		}
	}

	var (
		km             = kmeans.New()
		kmClusters, _  = km.Partition(obs, kgroups)
		dominantColors = make([]string, kgroups)
	)

	sort.SliceStable(kmClusters, func(i, j int) bool {
		return len(kmClusters[i].Observations) > len(kmClusters[j].Observations)
	})

	for i, cluster := range kmClusters {
		var (
			l   = cluster.Center[0]
			a   = cluster.Center[1]
			b   = cluster.Center[2]
			hex = colorful.Lab(l, a, b).Hex()
		)

		dominantColors[i] = hex
	}

	return &dominantColors, nil
}

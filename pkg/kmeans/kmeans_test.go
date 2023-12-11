package kmeans_test

import (
	"image"
	_ "image/jpeg"
	"math/rand"
	"net/http"
	"testing"

	"github.com/marcopeocchi/shizu/pkg/kmeans"
)

const testUrl = "https://w.wallhaven.cc/full/g7/wallhaven-g7joz7.jpg"

func TestKmeans(t *testing.T) {
	res, err := http.DefaultClient.Get(testUrl)
	if err != nil {
		t.Error(err)
	}

	defer res.Body.Close()

	rgba, _, err := image.Decode(res.Body)
	if err != nil {
		t.Error(err)
	}

	var data kmeans.Observations

	for y := rgba.Bounds().Min.Y; y < rgba.Bounds().Max.Y; y++ {
		for x := rgba.Bounds().Min.X; x < rgba.Bounds().Max.X; x++ {
			c := rgba.At(x, y)
			r, g, b, _ := c.RGBA()

			data = append(data, kmeans.Coordinates{
				float64(r >> 8),
				float64(g >> 8),
				float64(b >> 8),
			})
		}
	}

	k := kmeans.New()
	cls, err := k.Partition(data, 5, rand.Int63())
	if err != nil {
		t.Error(err)
	}

	for _, centroid := range cls {
		r := uint8(centroid.Center[0])
		g := uint8(centroid.Center[1])
		b := uint8(centroid.Center[2])

		t.Log(r, g, b)
	}
}

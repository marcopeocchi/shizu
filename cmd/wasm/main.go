package main

import (
	"fmt"
	"image"
	"strings"
	"syscall/js"

	"github.com/marcopeocchi/shizu/pkg/palette"
)

func wasmWrapper() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) < 1 {
			return "invalid arguments"
		}

		var (
			canvas    = make([]byte, args[0].Get("buffer").Get("byteLength").Int())
			kgroups   = args[0].Get("paletteSize").Int()
			width     = args[0].Get("width").Int()
			height    = args[0].Get("height").Int()
			tolerance = args[0].Get("tolerance").Float()
		)

		js.CopyBytesToGo(canvas, args[0].Get("buffer"))

		if len(args) == 3 {
			tolerance = args[2].Float()
		}

		rgba := image.NewRGBA(image.Rect(0, 0, width, height))
		rgba.Pix = canvas

		colors, err := palette.Get(rgba, kgroups, tolerance)
		if err != nil {
			fmt.Println(err)
		}

		return js.ValueOf(strings.Join((*colors), ","))
	})
}

func main() {
	js.Global().Set("getDominantColors", wasmWrapper())
	<-make(chan struct{})
}

package main

import (
	"fmt"
	"strings"
	"syscall/js"

	"github.com/marcopeocchi/shizu/pkg/kmeans"
	"github.com/marcopeocchi/shizu/pkg/utils"
)

func wasmWrapper() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) != 2 {
			return "invalid arguments"
		}

		b := make([]byte, args[0].Get("byteLength").Int())
		js.CopyBytesToGo(b, args[0])

		kgroups := args[1].Int()

		rgba, _, err := utils.DecodeImageFromBytes(b)
		if err != nil {
			fmt.Println(err)
		}

		colors, err := kmeans.GetDominantColorsHex(rgba, kgroups)
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

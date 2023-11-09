type GoRuntime = {
  new(): GoRuntime
  run: (instance: WebAssembly.Instance) => void
  importObject: WebAssembly.Imports
}

type ShizuParams = {
  buffer: Uint8Array
  paletteSize: number
  tolerance: number
  width: number,
  height: number,
}

declare const getDominantColors: (params: ShizuParams) => string
declare const Go: GoRuntime
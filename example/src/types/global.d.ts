type ShizuParams = {
  buffer: Uint8Array
  paletteSize: number
  tolerance: number
  width: number,
  height: number,
}

declare const generatePalette: (params: ShizuParams) => string
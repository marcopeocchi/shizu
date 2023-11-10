export const getPalette = async (src: string, nColors: number, tolerance: number) => {
  const img = new Image()
  img.crossOrigin = 'Anonymous'
  img.src = src

  const canvas = new OffscreenCanvas(img.width, img.height)
  const context = canvas.getContext('2d')

  context.drawImage(img, 0, 0)

  const buffer = context
    .getImageData(0, 0, img.width, img.height, { colorSpace: 'srgb' })
    .data
    .buffer

  return getDominantColors({
    buffer: new Uint8Array(buffer),
    height: img.height,
    width: img.width,
    paletteSize: nColors,
    tolerance: tolerance
  })
    .split(",")
}
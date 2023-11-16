self.onmessage = async (event: any) => {
  const res = await fetch(event.data)
  const blob = await res.blob()

  const img = await createImageBitmap(blob)

  const canvas = new OffscreenCanvas(img.width, img.height)
  const context = canvas.getContext('2d')

  context.drawImage(img, 0, 0)

  const buffer = context
    .getImageData(0, 0, img.width, img.height, { colorSpace: 'srgb' })
    .data
    .buffer

  self.postMessage({
    buffer,
    width: img.width,
    height: img.height,
  })
}
export const getPalette = async (src: string, nColors: number) => {
  const res = await fetch(src)
  const data = await res.blob()
  const buffer = await data.arrayBuffer()
  return getDominantColors(new Uint8Array(buffer), nColors, 0.85).split(",")
}
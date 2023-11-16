import { get } from 'svelte/store'
import {
  groupsStore,
  paletteStore,
  toleranceStore
} from '../store'

export const getPalette = async (src: string) => {
  const img = new Image()
  img.src = src

  const canvas = new OffscreenCanvas(img.width, img.height)
  const context = canvas.getContext('2d')

  context.drawImage(img, 0, 0)

  const buffer = context
    .getImageData(0, 0, img.width, img.height, { colorSpace: 'srgb' })
    .data
    .buffer

  const palette = getDominantColors({
    buffer: new Uint8Array(buffer),
    height: img.height,
    width: img.width,
    paletteSize: get(groupsStore),
    tolerance: get(toleranceStore)
  })
    .split(",")

  paletteStore.set(palette)
}

export const getDataUrlFromFile = async (file: File) =>
  new Promise<string>(resolve => {
    const reader = new FileReader()
    reader.onload = () => {
      resolve(reader.result as string)
      return
    }

    reader.readAsDataURL(file)
  })

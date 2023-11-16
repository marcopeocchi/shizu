import { get } from 'svelte/store'
import {
  groupsStore,
  loadingStore,
  paletteStore,
  toleranceStore
} from '../store'

type WorkerResponse = {
  buffer: ArrayBufferLike,
  height: number,
  width: number,
}

const decoderWorker = new Worker('./src/workers/decoder.worker.ts')

decoderWorker.onmessage = (event: MessageEvent<WorkerResponse>) => {
  const palette = getDominantColors({
    buffer: new Uint8Array(event.data.buffer),
    height: event.data.height,
    width: event.data.width,
    paletteSize: get(groupsStore),
    tolerance: get(toleranceStore)
  })
    .split(",")

  paletteStore.set(palette)
  loadingStore.set(false)
}

export const getPalette = async (src: string) => {
  loadingStore.set(true)
  decoderWorker.postMessage(src)
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

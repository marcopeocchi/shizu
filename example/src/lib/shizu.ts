import { get } from 'svelte/store'
import {
  groupsStore,
  loadingStore,
  paletteStore,
  toleranceStore
} from '../store'

import DecoderWorker from '../workers/decoder?worker'

type WorkerResponse = {
  buffer: ArrayBufferLike
  height: number
  width: number
}

const decoderWorker = new DecoderWorker()

decoderWorker.onmessage = (event: MessageEvent<WorkerResponse>) => {
  const palette = generatePalette({
    buffer: new Uint8Array(event.data.buffer),
    height: event.data.height,
    width: event.data.width,
    paletteSize: get(groupsStore),
    tolerance: get(toleranceStore)
  })
    .split(',')

  paletteStore.set(palette)
  loadingStore.update(() => false)
}

export const getPalette = async (src: string) => {
  loadingStore.update(() => true)
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

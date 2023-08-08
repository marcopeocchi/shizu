type GoRuntime = {
  new(): GoRuntime
  run: (instance: WebAssembly.Instance) => void
  importObject: WebAssembly.Imports
}

declare const getDominantColors: (buffer: Uint8Array, kgroups: number) => string
declare const Go: GoRuntime
import './app.css'
import App from './App.svelte'
import wasmURL from '/shizu.wasm?url'

const go = new Go()

let app: App

WebAssembly.instantiateStreaming(fetch(wasmURL), go.importObject).then((m) => {
  go.run(m.instance)

  app = new App({
    target: document.getElementById('app'),
  })
})

export default app
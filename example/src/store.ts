import { writable } from 'svelte/store'

export const paletteStore = writable<string[]>([])

export const groupsStore = writable(5)

export const toleranceStore = writable(0.9)

export const dataUrlStore = writable('')

export const loadingStore = writable(false)
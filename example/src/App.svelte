<script lang="ts">
  import testImage0 from '/gopher.webp';
  import Palette from './lib/Palette.svelte';
  import { getPalette } from './lib/shizu';

  let groups = 5;
  let tolerance = 0.9;
  let image = testImage0;
</script>

<main class="min-h-screen bg-neutral-950 text-neutral-50">
  <div class="flex gap-2 items-center justify-center">
    {#await getPalette(image, groups, tolerance)}
      <p>Generating palette...</p>
    {:then colors}
      <Palette {colors} />
    {/await}
    <img class="rounded-lg h-[82.5vh] m-8" src={image} alt="" />
  </div>
  <div class="flex items-center justify-center gap-4">
    <input
      bind:value={image}
      class="bg-neutral-800 p-1 rounded ring-neutral-500 ring-2 text-sm"
    />
    <div class="flex flex-col">
      <label for="palette-size" class="block mb-2 text-sm font-medium"
        >Palette size: {groups}</label
      >
      <input
        id="palette-size"
        type="range"
        min="3"
        max="12"
        step="1"
        bind:value={groups}
        class="appearance-none bg-neutral-800 rounded-lg h-2 cursor-pointer accent-blue-400"
      />
    </div>
    <div class="flex flex-col">
      <label for="tolerance-val" class="block mb-2 text-sm font-medium"
        >Tolerance: {tolerance}</label
      >
      <input
        id="tolerance-val"
        type="range"
        min="0.2"
        max="1.0"
        step="0.1"
        bind:value={tolerance}
        class="appearance-none bg-neutral-800 rounded-lg h-2 cursor-pointer accent-blue-400"
      />
    </div>
  </div>
</main>

<script lang="ts">
  import testImage0 from '/gopher.webp';
  import testImage1 from '/leaves.jpg';
  import testImage2 from '/r34.jpg';
  import testImage3 from '/r34-2.jpg';
  import Palette from './lib/Palette.svelte';
  import { getPalette } from './lib/shizu';

  let index = 0;
  let groups = 5;
  let images = [testImage0, testImage1, testImage2, testImage3];
</script>

<main class="min-h-screen bg-neutral-950 text-neutral-50">
  <div class="flex gap-2 items-center justify-center">
    {#await getPalette(images[index], groups)}
      <p>Generating palette...</p>
    {:then colors}
      <Palette {colors} />
    {/await}
    <img class="rounded-lg h-[82.5vh] m-8" src={images[index]} alt="" />
  </div>
  <div class="flex items-center justify-center gap-4">
    <select
      bind:value={index}
      class="bg-neutral-800 p-1 rounded ring-neutral-500 ring-2 text-sm"
    >
      <option value={0}>Photo 1</option>
      <option value={1}>Photo 2</option>
      <option value={2}>Photo 3</option>
    </select>
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
  </div>
</main>

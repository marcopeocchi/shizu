<script lang="ts">
  import { get } from 'svelte/store';
  import FileUpload from './lib/FileUpload.svelte';
  import FullScreenLoader from './lib/FullScreenLoader.svelte';
  import Palette from './lib/Palette.svelte';
  import PaletteSizeSlider from './lib/PaletteSizeSlider.svelte';
  import Placeholder from './lib/Placeholder.svelte';
  import ToleranceSlider from './lib/ToleranceSlider.svelte';
  import { getPalette } from './lib/shizu';
  import { dataUrlStore, loadingStore, paletteStore } from './store';
  import testImage0 from '/gopher.webp';
  import { onMount } from 'svelte';

  let loading = false;

  dataUrlStore.subscribe((url) => {
    getPalette(url);
  });

  const generate = async () => {
    loading = true;
    await getPalette(get(dataUrlStore));
    loading = false;
  };

  onMount(() => {
    dataUrlStore.set(testImage0);
    generate();
  });
</script>

<main class="min-h-screen bg-neutral-950 text-neutral-50">
  <div class="flex gap-2 items-center justify-center">
    {#if loading}
      <Placeholder />
    {:else if $loadingStore}
      <FullScreenLoader />
    {:else}
      <Palette colors={$paletteStore} />
      <img class="rounded-lg h-[82.5vh] m-8" src={$dataUrlStore} alt="" />
    {/if}
  </div>
  <div class="flex items-center justify-center gap-4">
    <PaletteSizeSlider />
    <ToleranceSlider />
    <FileUpload />
    <button
      on:click={generate}
      class="p-1.5 bg-blue-400 rounded text-sm text-black">Generate</button
    >
  </div>
</main>

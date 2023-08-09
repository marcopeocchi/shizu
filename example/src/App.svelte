<script lang="ts">
  import testImage1 from '/test.jpg';
  import testImage2 from '/r34.jpg';
  import testImage3 from '/r34-2.jpg';
  import Palette from './lib/Palette.svelte';
  import { getPalette } from './lib/shizu';

  let index = 0;
  let groups = 5;
  let images = [testImage1, testImage2, testImage3];
</script>

<main>
  <div>
    {#await getPalette(images[index], groups)}
      <p>Generating palette...</p>
    {:then colors}
      <Palette {colors} />
    {/await}
    <img src={images[index]} height="750" alt="" />
  </div>
  <div class="bottom">
    <select bind:value={index}>
      <option value={0}>Photo 1</option>
      <option value={1}>Photo 2</option>
      <option value={2}>Photo 3</option>
    </select>
    <input type="range" min="3" max="12" step="1" bind:value={groups} />
    <div>Palette size: {groups}</div>
  </div>
</main>

<style scoped>
  main {
    min-height: 100%;
  }
  div {
    display: flex;
    justify-content: center;
    align-items: center;
  }
  .bottom {
    gap: 10px;
    margin-top: 15px;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
  }
</style>

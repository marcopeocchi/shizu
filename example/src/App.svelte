<script lang="ts">
  import testImage from '/test.jpeg';
  import Palette from './lib/Palette.svelte';
  import { getPalette } from './lib/shizu';

  const palette = getPalette(testImage, 5);
</script>

<main>
  <div>
    {#await palette}
      <p>Generating palette...</p>
    {:then colors}
      <Palette {colors} />
    {/await}
    <img src={testImage} height="750" alt="" />
  </div>
  <a href="https://unsplash.com/it/foto/SUUvsLyEZjo">
    Photo by Pablo Arenas from unsplash.com
  </a>
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
  a {
    color: #cecece;
    text-decoration: none;
    margin-top: 20px;
    display: flex;
    justify-content: center;
    align-items: center;
  }
</style>

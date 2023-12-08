<script lang="ts">
  import { loadingStore, paletteStore } from '../store';
  import Spinner from './Spinner.svelte';
</script>

{#if $loadingStore}
  <Spinner />
{:else}
  <div class="rounded-lg shadow p-6 bg-neutral-900 ml-8">
    {#each $paletteStore as color, index}
      <div
        class={index === 0
          ? 'rounded-t-lg tile'
          : index === $paletteStore.length - 1
            ? 'rounded-b-lg tile'
            : 'tile'}
        style:height="calc(77vh / {$paletteStore.length}"
        style:background-color={color}
      />
    {/each}
  </div>
{/if}

<style scoped>
  .tile {
    width: 150px;
  }
</style>

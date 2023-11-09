<script lang="ts">
  export let colors: string[];
</script>

<div class="rounded-lg shadow p-8 bg-neutral-900">
  {#each colors as color, index}
    <div
      class={index === 0
        ? 'rounded-t-lg tile'
        : index === colors.length - 1
        ? 'rounded-b-lg tile'
        : 'tile'}
      style:height="{550 / colors.length}px"
      style:background-color={color}
    />
  {/each}
</div>

<style scoped>
  .tile {
    width: 150px;
  }
</style>

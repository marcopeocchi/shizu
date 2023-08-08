<script lang="ts">
  export let colors: string[];
</script>

<div class="container">
  {#each colors as color}
    <div class="tile" style="background-color: {color};" />
  {/each}
</div>

<style scoped>
  .container {
    display: flex;
    flex-direction: column;
  }

  .tile {
    width: 150px;
    height: 150px;
  }
</style>

<script lang="ts">
  import { groupsStore } from '../store';
</script>

<div class="flex flex-col">
  <label for="palette-size" class="block mb-2 text-sm font-medium"
    >Palette size: {$groupsStore}</label
  >
  <input
    id="palette-size"
    type="range"
    min="3"
    max="12"
    step="1"
    bind:value={$groupsStore}
    class="appearance-none bg-neutral-800 rounded-lg h-2 cursor-pointer accent-blue-400"
  />
</div>

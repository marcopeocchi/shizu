<script lang="ts">
  import { toleranceStore } from '../store';
</script>

<div class="flex flex-col">
  <label for="tolerance-val" class="block mb-2 text-sm font-medium"
    >Tolerance: {$toleranceStore}</label
  >
  <input
    id="tolerance-val"
    type="range"
    min="0.2"
    max="1.0"
    step="0.1"
    bind:value={$toleranceStore}
    class="appearance-none bg-neutral-800 rounded-lg h-2 cursor-pointer accent-blue-400"
  />
</div>

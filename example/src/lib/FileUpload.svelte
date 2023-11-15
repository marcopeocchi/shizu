<script lang="ts">
  import { dataUrlStore, loadingStore } from '../store';
  import { getDataUrlFromFile } from './shizu';

  let files: FileList;

  const extensions = ['.jpg', '.jpeg', '.png', '.webp', 'avif'];

  $: if (files) {
    loadingStore.set(true);

    getDataUrlFromFile(files[0]).then((url) => {
      dataUrlStore.set(url);
      loadingStore.set(false);
    });
  }
</script>

<div class="flex items-center gap-2">
  <label
    for="file-upload"
    class="p-1.5 bg-blue-400 rounded text-sm text-black cursor-pointer"
  >
    Open image
  </label>
  <input
    type="file"
    name="upload"
    id="file-upload"
    accept={extensions.join(',')}
    class="hidden"
    bind:files
  />
  <div>
    {files ? files[0].name : ''}
  </div>
</div>

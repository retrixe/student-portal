<script lang="ts">
  import { page } from '$app/state'
  import { Box, Tab, Tabs } from 'heliodor'
  import type { Snippet } from 'svelte'

  const { children }: { children: Snippet } = $props()
</script>

<Box class="page-container">
  <Tabs>
    <Tab href="/portal/circulars" active={page.route.id === '/portal/circulars'}>
      {#if page.route.id === '/portal/circulars'}
        Circulars
      {:else}
        Back to Circulars
      {/if}
    </Tab>
    {#if page.route.id === '/portal/circulars/[id]'}
      <Tab href={`/portal/circulars/${page.params.id}`} active>Circular #{page.params.id}</Tab>
    {/if}
  </Tabs>
  <br />
  {@render children()}
</Box>

<style lang="scss">
  :global(.page-container) {
    flex: 1;
    margin: 24px;
    padding: 16px;
  }
</style>

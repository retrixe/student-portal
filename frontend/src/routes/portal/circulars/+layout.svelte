<script lang="ts">
  import Box from '$lib/lunaria/Box.svelte'
  import Tab from '$lib/lunaria/Tab.svelte'
  import Tabs from '$lib/lunaria/Tabs.svelte'
  import type { Snippet } from 'svelte'

  const { children }: { children: Snippet } = $props()
</script>

<Box class="page-container">
  <Tabs>
    <Tab href="/portal/circulars">Circulars</Tab>
    <div style:flex="1"></div>
    <!-- TODO: Tab for current circular! -->
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

<script lang="ts">
  import Box from '$lib/lunaria/Box.svelte'
  import Tab from '$lib/lunaria/Tab.svelte'
  import Tabs from '$lib/lunaria/Tabs.svelte'
  import type { Snippet } from 'svelte'

  const { children }: { children: Snippet } = $props()
</script>

<Box class="page-container">
  <Tabs>
    <Tab href="/portal/fees">Pending Fees</Tab>
    <div style:flex="1"></div>
    <Tab href="/portal/fees/cart">Cart (4)</Tab>
    <Tab href="/portal/fees/receipts">Receipts</Tab>
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

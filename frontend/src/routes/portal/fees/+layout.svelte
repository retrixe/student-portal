<script lang="ts">
  import { page } from '$app/state'
  import { Box, Tab, Tabs } from 'heliodor'
  import type { Snippet } from 'svelte'

  const { children }: { children: Snippet } = $props()
</script>

<Box class="page-container">
  <Tabs>
    <Tab currentHref={page.route.id} href="/portal/fees">Pending Fees</Tab>
    <div style:flex="1"></div>
    <Tab currentHref={page.route.id} href="/portal/fees/cart">Cart (4)</Tab>
    <Tab currentHref={page.route.id} href="/portal/fees/receipts">Receipts</Tab>
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

<script lang="ts">
  import { page } from '$app/state'
  import { Box, Tab, Tabs } from 'heliodor'
  import type { Snippet } from 'svelte'

  const { children }: { children: Snippet } = $props()
</script>

<Box class="page-container">
  <Tabs>
    <Tab href="/portal/calendar/holidays" active={page.route.id === '/portal/calendar/holidays'}>
      Holidays
    </Tab>
    <Tab href="/portal/calendar/events" active={page.route.id === '/portal/calendar/events'}>
      Events
    </Tab>
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

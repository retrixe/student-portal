<script lang="ts">
  import { page } from '$app/state'
  import { Box, Tab, Tabs } from 'heliodor'
  import type { Snippet } from 'svelte'

  const { children }: { children: Snippet } = $props()
</script>

<Box class="page-container">
  <Tabs>
    <Tab href="/portal/service-requests" active={page.route.id === '/portal/service-requests'}>
      {#if page.route.id === '/portal/service-requests'}
        Service Requests
      {:else}
        Back to Service Requests
      {/if}
    </Tab>
    {#if page.route.id === '/portal/service-requests/[id]'}
      <Tab href={`/portal/service-requests/${page.params.id}`} active>
        Service Request #{page.params.id}
      </Tab>
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
    display: flex;
    flex-direction: column;
  }
</style>

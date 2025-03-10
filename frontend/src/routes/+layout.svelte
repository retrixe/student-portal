<script lang="ts">
  import { assets } from '$app/paths'
  import { page } from '$app/state'
  import type { Snippet } from 'svelte'
  import type { LayoutData } from './$types'
  import { onNavigate } from '$app/navigation'
  import '$lib/lunaria/Baseline.scss'

  const { children }: { data: LayoutData; children: Snippet } = $props()
  const { title, description, image, imageLarge, noIndex } = $derived(page.data)

  onNavigate(navigation => {
    if (!('startViewTransition' in document)) return

    return new Promise(resolve => {
      document.startViewTransition(async () => {
        resolve()
        await navigation.complete
      })
    })
  })
</script>

<svelte:head>
  <title>{title}</title>
  <meta property="og:type" content="website" />
  <meta property="og:title" content={title} />
  <meta property="og:url" content={page.url.origin + page.url.pathname} />
  <meta property="og:image" content={image ?? assets + '/favicon.png'} />
  <meta property="og:description" content={description} />
  <meta name="twitter:title" content={title} />
  <meta name="twitter:card" content={imageLarge ? 'summary_large_image' : 'summary'} />
  <meta name="twitter:image:src" content={image ?? assets + '/favicon.png'} />
  <meta name="twitter:description" content={description} />
  <meta name="Description" content={description} />
  {#if noIndex}
    <meta name="robots" content="noindex,nofollow" />
  {/if}
</svelte:head>

{@render children()}

<style lang="scss">
  :global {
    :root {
      --primary-color: #0080ff;
      --error-color: #ff0042;

      --link-color: #0080ff;
      --background-color: #f5f5f5; /* White smoke */
      --surface-color: #fcfcfc; /* White smoke but brighter */
      --color: #000000;
      --divider-color: #bbb;

      @media (prefers-color-scheme: dark) {
        --link-color: #00bfff;
        --background-color: #0e0e10; /* Jet black */
        --surface-color: #1b1b1b; /* Eerie black */
        --color: #ffffff;
        --divider-color: #666;
      }
    }
  }
</style>

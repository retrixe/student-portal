<script lang="ts">
  import { assets } from '$app/paths'
  import { page } from '$app/state'
  import type { Snippet } from 'svelte'
  import type { LayoutData } from './$types'
  import { onNavigate } from '$app/navigation'

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
    * {
      margin: 0;
      box-sizing: border-box;
    }

    :root {
      --primary-color: #8f00ff;
      --error-color: #ff0042;
    }

    @media (prefers-color-scheme: dark) {
      :root {
        color-scheme: dark;
        --link-color: #df73ff;
        --background-color: #0e0e10; /* Jet black */
        --surface-color: #1b1b1b; /* Eerie black */
        --color: #ffffff;
        --divider-color: #666;
      }
      .github-image {
        filter: brightness(0) invert(1);
      }
    }

    @media (prefers-color-scheme: light) {
      :root {
        --link-color: #8f00ff;
        --background-color: #f5f5f5; /* White smoke */
        --surface-color: #fcfcfc; /* White smoke but brighter */
        --color: #000000;
        --divider-color: #bbb;
      }
    }

    input {
      font: inherit;
    }
    select {
      font: inherit;
    }
    button {
      font: inherit;
    }
    textarea {
      font: inherit;
    }

    body {
      font-family: system-ui, 'Segoe UI', Roboto, Oxygen-Sans, Ubuntu, Cantarell, 'Helvetica Neue',
        Helvetica, Arial, sans-serif, 'Apple Color Emoji', 'Segoe UI Emoji', 'Segoe UI Symbol';
      background-color: var(--background-color);
      color: var(--color);
      width: 100vw;
      max-width: 100%;
      min-height: 100vh;
      display: flex;
      flex-direction: column;
      a {
        color: var(--link-color);
      }
    }
  }
</style>

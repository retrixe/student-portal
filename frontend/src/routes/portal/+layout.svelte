<script lang="ts">
  import { invalidate } from '$app/navigation'
  import Button from '$lib/lunaria/Button.svelte'
  import NavigationRail from '$lib/lunaria/NavigationRail.svelte'
  import ky from 'ky'
  import {
    CalendarDots,
    CalendarX,
    CurrencyInr,
    EnvelopeSimple,
    Exam,
    Mailbox,
    SignOut,
  } from 'phosphor-svelte'
  import type { Snippet } from 'svelte'

  const { children }: { children: Snippet } = $props()
  // const { username } = $derived(page.data)
  $effect(() => {
    // TODO: if (!username) goto('/').catch(console.error)
  })

  async function logout(event: Event) {
    event.preventDefault()
    try {
      await ky.post('api/logout')
      localStorage.removeItem('student-portal:token')
      invalidate('app:auth').catch(console.error)
    } catch (error) {
      alert('Failed to logout!')
      console.error('Failed to logout!', error)
    }
  }
</script>

<div class="container">
  <NavigationRail>
    {#snippet header()}
      <a href="/portal">
        <img src="/favicon.png" alt="logo" />
      </a>
    {/snippet}

    <a href="/portal/profile">
      <Button>
        <EnvelopeSimple weight="bold" size="24px" /> Profile
      </Button>
    </a>
    <a href="/portal/holidays-and-events">
      <Button>
        <CalendarDots weight="bold" size="24px" /> Holidays & Events
      </Button>
    </a>
    <a href="/portal/circulars">
      <Button>
        <Mailbox weight="bold" size="24px" /> Circulars
      </Button>
    </a>
    <a href="/portal/service-requests">
      <Button>
        <EnvelopeSimple weight="bold" size="24px" /> Service Requests
      </Button>
    </a>
    <a href="/portal/results">
      <Button>
        <Exam weight="bold" size="24px" /> Results
      </Button>
    </a>
    <a href="/portal/attendance">
      <Button>
        <CalendarX weight="bold" size="24px" /> Attendance
      </Button>
    </a>
    <a href="/portal/fees">
      <Button>
        <CurrencyInr weight="bold" size="24px" /> Fees
      </Button>
    </a>

    {#snippet footer()}
      <Button onclick={logout}>
        <SignOut weight="bold" size="24px" /> Logout
      </Button>
    {/snippet}
  </NavigationRail>

  <div class="container">
    {@render children()}
  </div>
</div>

<style lang="scss">
  .container {
    flex-grow: 1;
    display: flex;
  }

  img {
    width: 3rem;
    aspect-ratio: 1;
  }

  a {
    display: contents;
  }
</style>

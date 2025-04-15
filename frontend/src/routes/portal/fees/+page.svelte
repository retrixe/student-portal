<script lang="ts">
  import { goto } from '$app/navigation'
  import feeCart from '$lib/state/fees.svelte'
  import { Box, Button } from 'heliodor'

  const onCartAdd = (installment: number) => () => {
    if (feeCart.has(installment)) {
      feeCart.delete(installment)
    } else {
      feeCart.add(installment)
    }
  }

  const onCheckout = (installment: number) => () => {
    feeCart.add(installment)
    goto('/portal/fees/cart').catch(console.error)
  }
</script>

<Box class="item">
  <h1>Installment 1</h1>
  <br />
  <h2>INR 1,00,000</h2>
  <h3>Due by 12 Feb 2024</h3>
  <br />
  <Button onclick={onCartAdd(1)}>
    {#if feeCart.has(1)}
      Remove from Cart
    {:else}
      Add to Cart
    {/if}
  </Button>
  <Button onclick={onCheckout(1)}>Checkout</Button>
</Box>

<br />

<Box class="item">
  <h1>Installment 2</h1>
  <br />
  <h2>INR 1,00,000</h2>
  <h3>Due by 12 Feb 2024</h3>
  <br />
  <Button onclick={onCartAdd(2)}>
    {#if feeCart.has(2)}
      Remove from Cart
    {:else}
      Add to Cart
    {/if}
  </Button>
  <Button onclick={onCheckout(2)}>Checkout</Button>
</Box>

<style lang="scss">
  :global(.item) {
    padding: 24px;
  }
</style>

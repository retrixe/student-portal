<script>
  import Box from '$lib/lunaria/Box.svelte'
  import Button from '$lib/lunaria/Button.svelte'
</script>

<Box class="item">
  <h1>Installment 1</h1>
  <br />
  <h2>INR 1,00,000</h2>
  <h3>Due by 12 Feb 2024</h3>
  <br />
  <Button>Add to Cart</Button>
  <Button>Checkout</Button>
</Box>

<br />

<Box class="item">
  <h1>Installment 2</h1>
  <br />
  <h2>INR 1,00,000</h2>
  <h3>Due by 12 Feb 2024</h3>
  <br />
  <Button>Add to Cart</Button>
  <Button>Checkout</Button>
</Box>

<style lang="scss">
  :global(.item) {
    padding: 24px;
  }
</style>

<script>
  import Box from '$lib/lunaria/Box.svelte'
  import Button from '$lib/lunaria/Button.svelte'
  import { X } from 'phosphor-svelte'
</script>

<Box>
  <table>
    <thead>
      <tr>
        <th class="action-cells">Actions</th>
        <th>Fee Type</th>
        <th>Due Date</th>
        <th>Amount</th>
      </tr>
    </thead>
    <tbody>
      <tr class="item">
        <td class="action-cells">
          <Button id="remove-btn"><X weight="bold" size="24px" /></Button>
        </td>
        <td>Tuition Fee</td>
        <td>2023-10-01</td>
        <td>₹ 50,000</td>
      </tr>
      <tr class="item">
        <td class="action-cells">
          <Button id="remove-btn"><X weight="bold" size="24px" /></Button>
        </td>
        <td>Library Fee</td>
        <td>2023-10-01</td>
        <td>₹ 1,000</td>
      </tr>
      <!-- Add more rows as needed -->
    </tbody>
    <tfoot>
      <tr>
        <td colspan="3">Total</td>
        <td>₹ 51,000</td>
      </tr>
    </tfoot>
  </table>
</Box>

<br />

<Button class="checkout-btn">Checkout</Button>

<style lang="scss">
  table {
    width: 100%;
    table-layout: fixed;
    border-collapse: collapse;
  }

  thead,
  tbody > tr {
    border-bottom: 1px solid #ddd;
  }

  th,
  td {
    padding: 12px;
    text-align: left;
  }

  :global(.checkout-btn) {
    float: right;
  }

  :global(#remove-btn) {
    background-color: transparent;
    padding: 4px;
    cursor: pointer;
    &:hover {
      color: red;
    }
  }

  .action-cells {
    width: 80px;
  }
</style>

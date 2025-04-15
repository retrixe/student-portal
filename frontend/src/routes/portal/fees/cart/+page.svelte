<script lang="ts">
  import feeCart from '$lib/state/fees.svelte'
  import { Box, Button } from 'heliodor'
  import { X } from 'phosphor-svelte'

  interface InstallmentInfo {
    type: string
    amount: number
    dueDate: string
  }

  const installments: Record<number, InstallmentInfo> = {
    1: {
      type: 'Installment 1',
      amount: 100000,
      dueDate: '2024-02-12',
    },
    2: {
      type: 'Installment 2',
      amount: 100000,
      dueDate: '2024-08-12',
    },
  }
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
      {#each feeCart as fee (fee)}
        <tr class="item">
          <td class="action-cells">
            <Button onclick={() => feeCart.delete(fee)} class="remove-btn">
              <X weight="bold" size="24px" />
            </Button>
          </td>
          <td>{installments[fee].type}</td>
          <td>{installments[fee].dueDate}</td>
          <td>{installments[fee].amount}</td>
        </tr>
      {/each}
    </tbody>
    <tfoot>
      <tr>
        <td colspan="3">Total</td>
        <td>₹ {feeCart.values().reduce((total, fee) => total + installments[fee].amount, 0)}</td>
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

  :global(.remove-btn) {
    background-color: transparent !important;
    padding: 4px !important;
    &:hover {
      color: red;
    }
  }

  .action-cells {
    width: 80px;
  }
</style>

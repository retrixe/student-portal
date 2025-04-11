<script>
  import Box from '$lib/lunaria/Box.svelte'
</script>

<Box>
  <table>
    <thead>
      <tr>
        <th class="id-cells">#</th>
        <th>Name</th>
        <th class="date-cells">Date</th>
      </tr>
    </thead>
    <tbody>
      <tr class="item">
        <td class="id-cells">1</td>
        <td>Tuition Fee</td>
        <td class="date-cells">2023-10-01</td>
      </tr>
      <tr class="item">
        <td class="id-cells">2</td>
        <td>Library Fee</td>
        <td class="date-cells">2023-10-01</td>
      </tr>
      <!-- Add more rows as needed -->
    </tbody>
  </table>
</Box>

<br />

<style lang="scss">
  table {
    width: 100%;
    table-layout: fixed;
    border-collapse: collapse;
  }

  thead,
  tbody > tr:not(:last-child) {
    border-bottom: 1px solid #ddd;
  }

  th,
  td {
    padding: 12px;
    text-align: left;
  }

  .id-cells {
    width: 40px;
  }

  .date-cells {
    width: 160px;
  }
</style>

<script lang="ts">
  import { fetchCircularInfo } from '$lib/api/mock'
  import dayjs from 'dayjs'
  import { Box } from 'heliodor'

  interface CircularInfo {
    id: number
    date: string
    subject: string
  }

  const request: Promise<CircularInfo[]> = fetchCircularInfo()

  const formatDateTime = (date: string) => dayjs(date).format('MMM D, YYYY h:mm A')
</script>

{#await request}
  <h1>Loading...</h1>
{:then data}
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
        {#if data.length}
          {#each data as circular (circular.id)}
            <tr class="item">
              <td class="id-cells">
                <a href={`/portal/circulars/${circular.id}`}>{circular.id}</a>
              </td>
              <td>
                <a href={`/portal/circulars/${circular.id}`}>{circular.subject}</a>
              </td>
              <td class="date-cells">{formatDateTime(circular.date)}</td>
            </tr>
          {/each}
        {:else}
          <tr>
            <td colspan="3" style="text-align: center">
              <h3>No circulars available</h3>
            </td>
          </tr>
        {/if}
      </tbody>
    </table>
  </Box>
{:catch error}
  <h1 style:color="red">An error occurred!</h1>
  <h2>{error}</h2>
{/await}

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
    width: 200px;
  }

  a {
    text-decoration: none;
    color: var(--text-color);
  }
</style>

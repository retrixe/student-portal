<script lang="ts">
  import { fetchHolidayInfo } from '$lib/api/mock'
  import dayjs from 'dayjs'
  import { Box } from 'heliodor'

  interface HolidayInfo {
    name: string
    start: string
    end: string
    type: string
  }

  const request: Promise<HolidayInfo[]> = fetchHolidayInfo()

  // FIXME: dayjs being used in a lot of places in the same way
  const formatDateRange = (start: string, end: string) => {
    const startDate = dayjs(start).format('DD/MM/YYYY')
    const endDate = dayjs(end).format('DD/MM/YYYY')
    return startDate === endDate ? startDate : `${startDate} - ${endDate}`
  }
</script>

{#await request}
  <h1>Loading...</h1>
{:then data}
  <Box>
    <table>
      <thead>
        <tr>
          <th>Name</th>
          <th class="date-cells">Date</th>
          <th class="date-cells">Type</th>
        </tr>
      </thead>
      <tbody>
        {#if data.length}
          {#each data as holiday (holiday.name)}
            <tr class="item">
              <td>{holiday.name}</td>
              <td class="date-cells">{formatDateRange(holiday.start, holiday.end)}</td>
              <td class="date-cells">{holiday.type}</td>
            </tr>
          {/each}
        {:else}
          <tr>
            <td colspan="3" style="text-align: center">
              <h3>No holidays available</h3>
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

  .date-cells {
    width: 200px;
  }
</style>

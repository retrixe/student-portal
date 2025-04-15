<script lang="ts">
  import dayjs from 'dayjs'
  import { Box } from 'heliodor'

  interface ServiceRequestInfo {
    id: number
    date: string
    subject: string
  }

  const request: Promise<ServiceRequestInfo[]> = Promise.resolve([
    {
      id: 1,
      date: '2023-10-01',
      subject: 'Tuition Fee',
    },
    {
      id: 2,
      date: '2023-10-01',
      subject: 'Library Fee',
    },
  ])

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
          {#each data as serviceRequest (serviceRequest.id)}
            <tr class="item">
              <td class="id-cells">
                <a href={`/portal/service-requests/${serviceRequest.id}`}>{serviceRequest.id}</a>
              </td>
              <td>
                <a href={`/portal/service-requests/${serviceRequest.id}`}>
                  {serviceRequest.subject}
                </a>
              </td>
              <td class="date-cells">{formatDateTime(serviceRequest.date)}</td>
            </tr>
          {/each}
        {:else}
          <tr>
            <td colspan="3" style="text-align: center">
              <h3>No service requests available</h3>
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

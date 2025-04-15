<script lang="ts">
  import dayjs from 'dayjs'
  import { Box } from 'heliodor'

  interface ServiceRequest {
    id: number
    date: string
    subject: string
    description: string
  }

  const request: Promise<ServiceRequest> = Promise.resolve({
    id: 1,
    date: '2023-10-01',
    subject: 'Tuition Fee',
    description: 'This is a circular regarding the tuition fee for the semester.',
  })

  const formatDateTime = (date: string) => dayjs(date).format('MMM D, YYYY h:mm A')
</script>

{#await request}
  <h1>Loading...</h1>
{:then data}
  <div class="header">
    <h1>{data.subject}</h1>
    <h3>{formatDateTime(data.date)}</h3>
  </div>
  <br />
  <Box class="description-box">
    <p>
      {data.description}
    </p>
  </Box>
{:catch error}
  <h1 style:color="red">An error occurred!</h1>
  <h2>{error}</h2>
{/await}

<style lang="scss">
  .header {
    margin: 0px 16px;
    display: flex;
    align-items: center;
    justify-content: space-between;
  }

  :global(.description-box) {
    padding: 16px;
  }
</style>

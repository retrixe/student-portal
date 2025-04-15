<script lang="ts">
  import { page } from '$app/state'
  import type { Student } from '$lib/api/entities/Student'
  import dayjs from 'dayjs'
  import { Box, Button, Textarea } from 'heliodor'

  interface ServiceRequest {
    id: number
    author: string
    date: string
    subject: string
    description: string
    replies: {
      author: string
      description: string
      date: string
    }[]
  }

  let request: Promise<ServiceRequest> = $state(
    Promise.resolve({
      id: 1,
      date: '2023-10-01',
      subject: 'Tuition Fee',
      author: 'Ibrahim Ansari',
      description: 'This is a circular regarding the tuition fee for the semester.',
      replies: [],
    }),
  )

  let replyMessage = $state('')

  const onReply = (data: ServiceRequest) => {
    request = Promise.resolve({
      ...data,
      replies: [
        ...data.replies,
        {
          author: (page.data.student as Student).name,
          date: '2025-04-15',
          description: replyMessage,
        },
      ],
    })
    replyMessage = ''
  }

  const formatDateTime = (date: string) => dayjs(date).format('MMM D, YYYY h:mm A')
</script>

{#await request}
  <h1>Loading...</h1>
{:then data}
  <Box class="description-box">
    <h1>{data.subject}</h1>
    <div class="space-between">
      <h3>{data.author}</h3>
      <h3>{formatDateTime(data.date)}</h3>
    </div>
    <br />
    <hr />
    <br />
    <p>
      {data.description}
    </p>
  </Box>
  {#each data.replies as reply, i (i)}
    <Box class="description-box">
      <div class="space-between">
        <h3>#{i + 1} - {reply.author}</h3>
        <h3>{formatDateTime(reply.date)}</h3>
      </div>
      <br />
      <hr />
      <br />
      <p>
        {reply.description}
      </p>
    </Box>
  {/each}
  <Textarea class="reply-box" bind:value={replyMessage} placeholder="Write your reply here..." />
  <Button class="reply-btn" onclick={() => onReply(data)}>Reply</Button>
{:catch error}
  <h1 style:color="red">An error occurred!</h1>
  <h2>{error}</h2>
{/await}

<style lang="scss">
  .space-between {
    display: flex;
    justify-content: space-between;
  }

  :global(.description-box) {
    padding: 16px;
    margin-bottom: 2rem;
  }

  :global(.reply-box) {
    width: 100% !important;
  }

  :global(.reply-btn) {
    align-self: flex-end;
  }
</style>

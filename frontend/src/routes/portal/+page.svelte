<script>
  import Box from '$lib/lunaria/Box.svelte'
  import Button from '$lib/lunaria/Button.svelte'
  import ProgessBar from '$lib/lunaria/ProgessBar.svelte'
  import { XCircle } from 'phosphor-svelte'

  // FIXME: Make everything retrievable from the backend via functions i.e. reactive UI
</script>

<div class="container">
  <Box class="card header">
    <img class="user-avatar" src="https://github.com/retrixe.png" alt="logo" />
    <div>
      <h1>Ibrahim Ansari</h1>
      <h3>PRN 3929303209</h3>
      <h4>Semester 4</h4>
      <h4>B.Tech CSE</h4>
    </div>
  </Box>

  <Box class="card">
    <div class="space-between">
      <h1>Attendance Summary</h1>
      <h1 style:color="yellow /* lime or red */">75%</h1>
    </div>
    <br />
    <div class="attendance-indicator">
      <p>Design and Analysis of Algorithms</p>
      <ProgessBar percentage={75} color="yellow" />
      <p>
        <span style:color="yellow">75%</span> (15/20)
      </p>
    </div>
  </Box>

  <Box class="card header">
    <!--
      <CheckCircle size="64px" color="#90EE90" />
      <h1>No pending fees</h1>
    -->
    <!--
      <WarningCircle size="64px" color="#FFCC00" />
      <div style:flex="1">
        <h1>Pending fees</h1>
        <h2>INR 3,00,000 due by May 12, 2025</h2>
      </div>
    -->
    <XCircle size="64px" color="red" />
    <div style:flex="1">
      <h1>Fee payment overdue!</h1>
      <h2>INR 3,00,000 was due by May 12, 2025</h2>
    </div>
    <Button>Pay now</Button>
  </Box>

  <Box class="card">
    <!-- FIXME: This needs working on -->
    <h1>Upcoming Events</h1>
    <div class="space-between">
      <h2>Hackathon</h2>
      <h2>12 Feb 2024</h2>
    </div>
    <div class="space-between">
      <h2>Project Submission</h2>
      <h2>20 Feb 2024</h2>
    </div>
    <div class="space-between">
      <h2>Mid Sem Exams</h2>
      <h2>25 Feb 2024</h2>
    </div>
  </Box>
</div>

<style>
  .container {
    flex: 1;
    display: flex;
    flex-direction: column;
    align-items: center;
    padding: 48px;
    gap: 24px;
  }

  .container > :global(.card) {
    padding: 24px;
    max-width: 960px;
    width: 100%;
  }

  .container > :global(.header) {
    display: flex;
    align-items: center;
    gap: 24px;
  }

  .space-between {
    display: flex;
    justify-content: space-between;
    align-items: center;
  }

  .attendance-indicator {
    display: flex;
    align-items: center;
    gap: 24px;
    > :global(.progress-bar) {
      flex: 1;
    }
  }

  .user-avatar {
    width: 100px;
    aspect-ratio: 1;
    border-radius: 100%;
  }
</style>

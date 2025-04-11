<script lang="ts">
  import Box from '$lib/lunaria/Box.svelte'
  import Button from '$lib/lunaria/Button.svelte'
  import ProgessBar from '$lib/lunaria/ProgessBar.svelte'
  import dayjs from 'dayjs'
  import { CheckCircle, WarningCircle, XCircle } from 'phosphor-svelte'

  const student = {
    image: 'https://github.com/retrixe.png',
    name: 'Ibrahim Ansari',
    prn: '3929303209',
    semester: 4,
    branch: 'B.Tech CSE',
  }

  interface LandingData {
    totalAttendance: number
    courseAttendance: {
      courseCode: string
      courseName: string
      totalClasses: number
      attendedClasses: number
    }[]
    pendingFees?: {
      totalAmount: number
      dueDate: string
    }
    holidays: {
      name: string
      start: string
      end: string
      type: string
    }[]
  }

  const request: Promise<LandingData> = Promise.resolve({
    totalAttendance: 75,
    courseAttendance: [
      {
        courseCode: 'DAA',
        courseName: 'Design and Analysis of Algorithms',
        totalClasses: 20,
        attendedClasses: 15,
      },
    ],
    pendingFees: {
      totalAmount: 300000,
      dueDate: '2025-05-12',
    },
    holidays: [
      { name: 'Hackathon', start: '2024-02-12', end: '2024-02-14', type: 'national' },
      { name: 'Project Submission', start: '2024-02-20', end: '2024-02-20', type: 'national' },
      { name: 'Mid Sem Exams', start: '2024-02-25', end: '2024-02-25', type: 'national' },
    ],
  })

  const formatNumber = (n: number) =>
    n.toLocaleString('en-IN', {
      style: 'currency',
      currency: 'INR',
      trailingZeroDisplay: 'stripIfInteger',
      maximumFractionDigits: 2,
    })

  const formatDateRange = (start: Date, end: Date) => {
    const startDate = dayjs(start).format('DD/MM/YYYY')
    const endDate = dayjs(end).format('DD/MM/YYYY')
    return startDate === endDate ? startDate : `${startDate} - ${endDate}`
  }

  const calculateAttendance = (attendedClasses: number, totalClasses: number): number => {
    const percentage = (attendedClasses / totalClasses) * 100
    return +percentage.toFixed(2)
  }

  const getAttendanceColor = (attendance: number): string =>
    attendance < 75 ? 'red' : attendance < 80 ? 'yellow' : 'lime'
</script>

<div class="container">
  {#await request}
    <Box class="card">
      <h1>Loading...</h1>
    </Box>
  {:then data}
    <Box class="card header">
      <img class="user-avatar" src={student.image} alt="logo" />
      <div>
        <h1>{student.name}</h1>
        <h3>PRN {student.prn}</h3>
        <h4>Semester {student.semester}</h4>
        <h4>{student.branch}</h4>
      </div>
    </Box>

    <Box class="card">
      <div class="space-between">
        <h1>Attendance Summary</h1>
        <h1 style:color={getAttendanceColor(data.totalAttendance)}>
          {data.totalAttendance}%
        </h1>
      </div>
      <br />
      {#each data.courseAttendance as subject (subject.courseCode)}
        <div class="attendance-indicator">
          <p>{subject.courseName}</p>
          <ProgessBar
            percentage={calculateAttendance(subject.attendedClasses, subject.totalClasses)}
            color={getAttendanceColor(
              calculateAttendance(subject.attendedClasses, subject.totalClasses),
            )}
          />
          <p>
            <span
              style:color={getAttendanceColor(
                calculateAttendance(subject.attendedClasses, subject.totalClasses),
              )}
            >
              {calculateAttendance(subject.attendedClasses, subject.totalClasses)}%
            </span>
            ({subject.attendedClasses}/{subject.totalClasses})
          </p>
        </div>
      {/each}
    </Box>

    <Box class="card header">
      {#if !data.pendingFees}
        <CheckCircle size="64px" color="#90EE90" />
        <h1>No pending fees</h1>
      {:else}
        {#if new Date(data.pendingFees.dueDate).getTime() > Date.now()}
          <WarningCircle size="64px" color="#FFCC00" />
          <div style:flex="1">
            <h1>Pending fees</h1>
            <h2>
              {formatNumber(data.pendingFees.totalAmount)} due by {dayjs(
                data.pendingFees.dueDate,
              ).format('MMMM D, YYYY')}
            </h2>
          </div>
        {:else}
          <XCircle size="64px" color="red" />
          <div style:flex="1">
            <h1>Fee payment overdue!</h1>
            <h2>
              {formatNumber(data.pendingFees.totalAmount)} was due by {dayjs(
                data.pendingFees.dueDate,
              ).format('MMMM D, YYYY')}
            </h2>
          </div>
        {/if}
        <Button>Pay now</Button>
      {/if}
    </Box>

    <Box class="card">
      <h1>Upcoming Holidays</h1>
      <br />
      {#each data.holidays as event (event.name)}
        <div class="space-between">
          <h3>{event.name}</h3>
          <h3>
            {formatDateRange(new Date(event.start), new Date(event.end))}
          </h3>
        </div>
      {/each}
    </Box>
  {:catch error}
    <Box class="card">
      <h1 style:color="red">An error occurred!</h1>
      <h2>{error}</h2>
    </Box>
  {/await}
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

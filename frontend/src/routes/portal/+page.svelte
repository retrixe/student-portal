<script lang="ts">
  import { page } from '$app/state'
  import type { Student } from '$lib/api/entities/Student'
  import { fetchLandingData } from '$lib/api/mock'
  import dayjs from 'dayjs'
  import { Box, Button, ProgressBar } from 'heliodor'
  import { CheckCircle, WarningCircle, XCircle } from 'phosphor-svelte'

  const { student } = $derived(page.data) as { student: Student }

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

  const request: Promise<LandingData> = fetchLandingData()

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
      <img class="user-avatar" src={`data:image/png;base64,${student.picture}`} alt="logo" />
      <div>
        <h1>{student.name}</h1>
        <h3>PRN {student.prn}</h3>
        <h4>Semester {student.semester}</h4>
        <h4>{student.programCode}</h4>
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
          <ProgressBar
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
        <a href="/portal/fees">
          <Button>Pay now</Button>
        </a>
      {/if}
    </Box>

    <Box class="card">
      <h1>Upcoming Holidays / Events</h1>
      <br />
      {@const upcomingHolidays = data.holidays.filter(holiday => {
        const timeDifference = new Date(holiday.start).getTime() - Date.now()
        return timeDifference > 0 && timeDifference < 7 * 24 * 60 * 60 * 1000
      })}
      {#if upcomingHolidays.length}
        {#each upcomingHolidays as holiday (holiday.name)}
          <div class="space-between">
            <h3>{holiday.name}</h3>
            <h3>
              {formatDateRange(new Date(holiday.start), new Date(holiday.end))}
            </h3>
          </div>
        {/each}
      {:else}
        <h3>No upcoming holidays this week :(</h3>
      {/if}
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

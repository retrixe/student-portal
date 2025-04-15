<script lang="ts">
  import { page } from '$app/state'
  import type { Student } from '$lib/api/entities/Student'
  import dayjs from 'dayjs'
  import { Box } from 'heliodor'

  const { student } = $derived(page.data) as { student: Student }
</script>

<div class="container">
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
    <h3>E-mail</h3>
    <p>{student.email}</p>
    <br />
    <h3>Phone No</h3>
    <p>{student.phoneNo}</p>
    <br />
    <h3>Aadhaar No</h3>
    <p>{student.aadhaarNo}</p>
    <br />
    <h3>Blood Group</h3>
    <p>{student.bloodGroup}</p>
    <br />
    <h3>Date of Birth</h3>
    <p>{dayjs(student.dob).format('DD/MM/YYYY')}</p>
    <br />
    <h3>Gender</h3>
    <p>{student.gender}</p>
    <br />
    <h3>Admission Date</h3>
    <p>{dayjs(student.admissionDate).format('DD/MM/YYYY')}</p>
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

  .user-avatar {
    width: 100px;
    aspect-ratio: 1;
    border-radius: 100%;
  }
</style>

export const fetchLandingData = () =>
  Promise.resolve({
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
      { name: 'Hackathon', start: '2025-04-19', end: '2025-04-20', type: 'national' },
      { name: 'Project Submission', start: '2025-04-16', end: '2025-04-16', type: 'national' },
      { name: 'Mid Sem Exams', start: '2025-04-25', end: '2025-04-25', type: 'national' },
    ],
  })

export const fetchServiceRequests = () =>
  Promise.resolve([
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

export const fetchEventInfo = () =>
  Promise.resolve([
    { name: 'Hackathon', start: '2024-02-12', end: '2024-02-14', type: 'National' },
    { name: 'Project Submission', start: '2024-02-20', end: '2024-02-20', type: 'National' },
    { name: 'Mid Sem Exams', start: '2024-02-25', end: '2024-02-25', type: 'National' },
  ])

export const fetchHolidayInfo = () =>
  Promise.resolve([
    {
      name: 'Diwali',
      start: '2023-10-01',
      end: '2023-10-05',
      type: 'National',
    },
    {
      name: 'Christmas',
      start: '2023-12-24',
      end: '2023-12-26',
      type: 'National',
    },
    {
      name: 'New Year',
      start: '2024-01-01',
      end: '2024-01-02',
      type: 'National',
    },
  ])

export const fetchCircularInfo = () =>
  Promise.resolve([
    {
      id: 1,
      date: '2023-10-01',
      subject: 'Update on Tuition Fee',
    },
    {
      id: 2,
      date: '2023-10-01',
      subject: 'Attendance Criteria Update',
    },
  ])

export const fetchCircular = () =>
  Promise.resolve({
    id: 1,
    date: '2023-10-01',
    subject: 'Tuition Fee',
    description: 'This is a circular regarding the tuition fee for the semester.',
  })

import type { Student } from '$lib/api/entities/Student'
import ky from '$lib/api/ky'
import type { LayoutLoad } from './$types'

export const load: LayoutLoad = async event => {
  const { fetch } = event
  event.depends('app:auth')

  if (typeof localStorage !== 'object') return {}

  // Ignore errors trying to check for authentication state
  try {
    const req = await ky('', { fetch }).json<{ student?: Student }>()
    return { student: req.student }
  } catch (e) {
    console.error(e)
  }
}

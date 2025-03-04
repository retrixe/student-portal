import ky from '$lib/api/ky'
import type { LayoutLoad } from './$types'

export const load: LayoutLoad = async event => {
  const { fetch } = event
  event.depends('app:auth')

  if (typeof localStorage !== 'object') return {}

  // Ignore errors trying to check for authentication state
  try {
    const req = await ky('', { fetch }).json<{ userId?: string; username?: string }>()
    return { username: req.username, userId: req.userId }
  } catch (e) {
    console.error(e)
  }
}

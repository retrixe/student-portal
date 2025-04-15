import { env } from '$env/dynamic/public'
import type { PageParentData, PageServerLoad } from './$types'

export const load: PageServerLoad<PageParentData> = () => {
  return {
    title: 'Holidays - Student Portal - My ' + env.PUBLIC_COLLEGE_NAME,
    image: '/favicon.png',
    description: 'Holidays on your portal to My ' + env.PUBLIC_COLLEGE_NAME + '.',
  }
}

import { env } from '$env/dynamic/public'
import type { PageParentData, PageServerLoad } from './$types'

export const load: PageServerLoad<PageParentData> = () => {
  return {
    title: 'Service Requests - Student Portal - My ' + env.PUBLIC_COLLEGE_NAME,
    image: '/favicon.png',
    description: 'Service requests on your portal to My ' + env.PUBLIC_COLLEGE_NAME + '.',
  }
}

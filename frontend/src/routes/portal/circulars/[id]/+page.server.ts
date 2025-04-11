import { env } from '$env/dynamic/public'
import type { PageParentData, PageServerLoad } from './$types'

export const load: PageServerLoad<PageParentData> = data => {
  return {
    title: 'Circular #' + data.params.id + ' - Student Portal - My ' + env.PUBLIC_COLLEGE_NAME,
    image: '/favicon.png',
    description:
      'Circular #' + data.params.id + ' on your portal to My ' + env.PUBLIC_COLLEGE_NAME + '.',
  }
}

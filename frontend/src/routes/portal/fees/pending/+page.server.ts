import { env } from '$env/dynamic/public'
import type { PageServerLoad } from './$types'

interface PageData {
  title: string
  image?: string
  imageLarge?: boolean
  description: string
  noIndex?: boolean
}

export const load: PageServerLoad<PageData> = () => {
  return {
    title: 'Pending Fee Payments - Student Portal - My ' + env.PUBLIC_COLLEGE_NAME,
    image: '/favicon.png',
    description: 'Pending fee payments on your portal to My ' + env.PUBLIC_COLLEGE_NAME + '.',
  }
}

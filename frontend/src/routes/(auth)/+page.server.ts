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
    title: 'Login - My ' + env.PUBLIC_COLLEGE_NAME,
    image: '/favicon.png',
    description: 'Log into your account on My ' + env.PUBLIC_COLLEGE_NAME + '.',
  }
}

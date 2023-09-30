import { isAuthenticated } from '~/composables/user';

export default defineNuxtRouteMiddleware((to, from) => {
    // 로그인 권한이 없을때 로그인 페이지 제외
    if (!isAuthenticated.value && to.path !== '/login') {
        return navigateTo('/login')
    }
  }) 
import { useAuthStore } from '~/stores/auth';

export default defineNuxtRouteMiddleware((to, from) => {
    // if (process.client) { // 클라이언트 사이드에서만 실행되도록 조건을 추가....
    //     const authStore = useAuthStore();
    //     if (authStore.isLoading) return;  // 로딩 중이라면 로직을 건너뜁니다.

    //     if (!authStore.isAuthenticated && to.path !== '/login') {
    //         return navigateTo('/login');
    //     }
    // }
});
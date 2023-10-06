import { useAuthStore } from '~/stores/auth';

export default defineNuxtRouteMiddleware((to, from) => {
    const authStore = useAuthStore();
    
    if (!authStore.isAuthenticated && (to.path !== '/login'  &&  to.path !== '/signup')) {
        return navigateTo('/login');
    }
});
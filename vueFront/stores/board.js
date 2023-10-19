// ~/stores/auth.js
import { defineStore } from 'pinia';

export const useAuthStore = defineStore({
  id: 'auth',

  // state 로그인 상태
  state: () => ({
    user: null,
    loginDate: null,
    isAuthenticated: false, //로그인여부
    isLoading: true,  // 로딩 상태 추가
  }),

  // getters
  getters: {
    // 필요한 경우 여기에 getters를 추가할 수 있습니다.
  },

  // actions
  actions: {
    setLoading(Loading) {
      this.isLoading = Loading;
    },

    setUser(newUser) {
      this.user = newUser;
    },

    setLoginDate() {
      this.loginDate = new Date().toLocaleString();
    },

    logout() {
      this.user = null;
      this.loginDate = null;
      this.isAuthenticated = false;
      this.token = null;
      
      // Pinia 내에서 직접 라우팅은 권장되지 않습니다.
      // 대신 컴포넌트나 라우트 후크 내에서 로그아웃 후 리디렉션을 처리하는 것이 좋습니다.
    },

    setAuthenticated(authStatus) {
      this.isAuthenticated = authStatus;
    },

    login(newUser) {
      this.setUser(newUser);
      this.setLoginDate();
      this.isAuthenticated = true;
    },
  },
});


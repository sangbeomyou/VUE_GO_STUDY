import { useAuthStore } from '~/stores/auth';
import axios from 'axios' // Axios를 사용한다고 가정

export default defineNuxtPlugin(async () => {
  if (process.client) { // 클라이언트 사이드에서만 실행되도록 조건을 추가....
    const auth = useAuthStore();
  
    if (!auth.user) {
      try {
        const response = await axios.get('http://localhost:8080/InitUser', { withCredentials: true });
        if (response) {
          console.log(response)
          auth.login({ email: response.data.email, name: response.data.displayName, uid: response.data.rawId });
        }
      } catch (error) {
        console.error('Error fetching user data:', error);
      } finally {
        auth.setLoading(false);  // 데이터 로드가 완료되면 로딩 상태를 false로 설정
      }
    } else {
      auth.setLoading(false);  // 이미 유저 데이터가 있으면 로딩 상태를 false로 설정
    }
  }
});

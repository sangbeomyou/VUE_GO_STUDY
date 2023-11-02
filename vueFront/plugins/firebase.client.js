import { initializeApp, getApps } from "firebase/app";
import { getAuth, onAuthStateChanged } from "firebase/auth";
import { getAnalytics } from "firebase/analytics";
import { useAuthStore } from '~/stores/auth';


export default defineNuxtPlugin(async (nuxtApp) => {
  const config = useRuntimeConfig()
  const authStore = useAuthStore();

  const firebaseConfig = config.public.firebase;


  if (!getApps().length) { // 이미 초기화된 앱이 없는 경우에만 초기화를 진행합니다.
    initializeApp(firebaseConfig);
  }

  const app = getApps()[0]; // 초기화된 Firebase 앱 가져오기

  const analytics = getAnalytics(app);
  console.log(app)

  const auth = getAuth(app);
  if (process.client) {
    onAuthStateChanged(auth, (firebaseUser) => {
      if (firebaseUser) {
        console.log(firebaseUser);
        authStore.login({ email: firebaseUser.email, name: firebaseUser.displayName, uid: firebaseUser.uid });
      }
      else {
        console.log(firebaseUser);
        navigateTo('/Login');
      }
    });
  }

  return {
    provide: {
      fireAuth: auth,
      fireAnalytics: analytics,
    },
  };

});
import { initializeApp, getApps  } from "firebase/app";
import { getAuth, onAuthStateChanged } from "firebase/auth";
import { getAnalytics } from "firebase/analytics";
import { useAuthStore } from '~/stores/auth';

export default defineNuxtPlugin(async(nuxtApp) => {
  const config = useRuntimeConfig()
  const authStore = useAuthStore();

  const firebaseConfig = config.public.firebase;

  const app = initializeApp(firebaseConfig);

  const analytics = getAnalytics(app);
  console.log(app)

  if (!getApps().length) {
    console.log(auth)
  }
  const auth = getAuth(app);

  onAuthStateChanged(auth, (firebaseUser) => {
    if (firebaseUser) {
      console.log("onAuthStateChanged", firebaseUser);
      authStore.login({ email: firebaseUser.email, name: firebaseUser.displayName, uid: firebaseUser.uid });
      navigateTo('/home');
    }
    else  {
      console.log("onAuthStateChanged", firebaseUser);
      navigateTo('/Login');
    }
  }); 

  
  return {
    provide: {
      fireAuth: auth,
      fireAnalytics: analytics,
    },
  };

});
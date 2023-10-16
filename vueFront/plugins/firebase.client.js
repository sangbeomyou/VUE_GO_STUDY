import { initializeApp, getApps  } from "firebase/app";
import { getAuth } from "firebase/auth";
import { getAnalytics } from "firebase/analytics";

export default defineNuxtPlugin(async(nuxtApp) => {
  const config = useRuntimeConfig()

  const firebaseConfig = config.public.firebase;

  const app = initializeApp(firebaseConfig);

  const analytics = getAnalytics(app);
  console.log(app)

  if (!getApps().length) {
    console.log(auth)
  }
  const auth = getAuth(app);

  return {
    provide: {
      fireAuth: auth,
      fireAnalytics: analytics,
    },
  };

});
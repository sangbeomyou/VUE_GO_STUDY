// composables/useFirebase.ts

import { ref, onMounted } from 'vue'
import { initializeApp, getApps, getApp } from 'firebase/app'
import { getAuth, onAuthStateChanged } from 'firebase/auth'
const config = useRuntimeConfig();

const appconfig = {
    apiKey: config.public.apiKey,
    appId: config.public.appId,
    authDomain: config.public.authDomain,
    databaseURL: config.public.databaseURL,
    measurementId: config.public.measurementId,
    messagingSenderId: config.public.messagingSenderId,
    projectId: config.public.projectId,
    storageBucket: config.public.storageBucket,
}

const firebaseApp = (!getApps().length ? initializeApp(appconfig) : getApp())
const auth = getAuth(firebaseApp)

export const useFirebaseAuth = () => {
    const user = ref(null)

    onMounted(() => {
        onAuthStateChanged(auth, (firebaseUser) => {
            user.value = firebaseUser
        })
    })

    const login = async (email, password) => {
        await getAuth().signInWithEmailAndPassword(email, password)
    }

    const logout = async () => {
        await getAuth().signOut()
    }

    return {
        user,
        signIn,
        signOut,
    }
}

import { ref, onMounted } from 'vue';
import { getAuth, onAuthStateChanged, signInWithEmailAndPassword, createUserWithEmailAndPassword } from 'firebase/auth';

export default function useFirebase() {
  const user = ref(null); // user를 ref로 초기화
  const auth = getAuth(); // Firebase Auth 인스턴스를 가져옵니다.

  onMounted(() => {
    onAuthStateChanged(auth, (firebaseUser) => {
      user.value = firebaseUser;
    });
  });

  const login = async (email, password) => {
    await signInWithEmailAndPassword(auth, email, password);
  }

  const logout = async () => {
    await auth.signOut();
  }

  const signUp = async (email, password) => {
    await createUserWithEmailAndPassword(auth, email, password);
  }

  return {
    user,
    signUp,
    login,
    logout,
  };
}

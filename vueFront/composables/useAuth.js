import { ref, onMounted } from 'vue';
import { getAuth, onAuthStateChanged, signInWithEmailAndPassword, createUserWithEmailAndPassword, updateProfile, fetchSignInMethodsForEmail } from 'firebase/auth';

export default function useFirebase() {
  const user = ref(null);
  let auth;

  if (process.client) {
    auth = getAuth(); // 클라이언트 측에서만 Firebase Auth 인스턴스를 가져옵니다.

    onMounted(() => {
      onAuthStateChanged(auth, (firebaseUser) => {
        user.value = firebaseUser;
      });
    });
  }

  const login = async (email, password) => {
    if (auth) {
      try {
        return await signInWithEmailAndPassword(auth, email, password);
      } catch (error) {
        throw error;  // 오류가 발생하면 외부로 전달
      }
    }
  }

  const logout = async () => {
    if (auth) {
      try {
        return await auth.signOut();
      } catch (error) {
        throw error;  // 오류가 발생하면 외부로 전달
      }
    }
  }

  const signUp = async (email, name, password) => {
    if (auth) {
      try {
        const userCredential = await createUserWithEmailAndPassword(auth, email, password);
        const user = userCredential.user;
        // 사용자 프로필 업데이트 (displayName 및 photoURL)
        await updateProfile(user, {
          displayName: name,
          // photoURL: 'http://example.com/photo.png' // 필요한 경우
        });

      } catch (error) {
        throw error;  // 오류가 발생하면 외부로 전달
      }
    }
  }

  const checkEmail = async (email) => {
    if (auth) {
      try {
        return await fetchSignInMethodsForEmail(auth, email);
      } catch (error) {
        throw error;  // 오류가 발생하면 외부로 전달
      }
    }
  }

  return {
    user,
    signUp,
    checkEmail,
    login,
    logout,
  };
}

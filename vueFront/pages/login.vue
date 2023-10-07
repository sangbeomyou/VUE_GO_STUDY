<template>
  <v-container>
    <v-row justify="center" align="center" class="fill-height">
      <v-col cols="12" sm="8" md="6">
        <v-card class="elevation-12">
          <v-toolbar color="primary" dark flat>
            <v-toolbar-title>로그인</v-toolbar-title>
          </v-toolbar>
          <v-card-text>
            <v-text-field label="이메일" id="email" name="email" prepend-icon="mdi-account" type="text"
              :rules="[rules.required]" required v-model="email"></v-text-field>

            <v-text-field label="비밀번호" id="Password" name="password" prepend-icon="mdi-lock" type="password"
              :rules="[rules.required]" required v-model="password"></v-text-field>
          </v-card-text>
          <v-card-actions>
            <v-btn color="primary" @click="signup">회원가입</v-btn>
            <v-spacer></v-spacer>
            <v-btn color="primary" @click="loginclick">로그인</v-btn>

          </v-card-actions>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>
<script setup>
import { ref } from 'vue';
import axios from 'axios';
import { useAuthStore } from '~/stores/auth';
import useAuth from '~/composables/useAuth'

const { login } = useAuth();

definePageMeta({
  layout: false,
  middleware: false
});

const auth = useAuthStore();
const email = ref('');
const password = ref('');

const rules = {
  required: value => !!value || '빈칸을 입력하세요.',
};

const signup = () => {
  navigateTo('/signup');
}

const loginclick = async () => {
  try {
    const userCredential = await login(email.value, password.value);
    // 서버에 토큰 전송
    const token = await userCredential.user.getIdToken();
    const response = await axios.post('http://localhost:8080/public/setToken', { token: token }, {
      withCredentials: true
    });
    console.log(response)
    // Firebase에서 설정된 email displayName을 가져옵니다.
    const userEmail = userCredential.user.email;
    const userName = userCredential.user.displayName;
    const uid = userCredential.user.uid;

    //store 저장
    auth.login({ email: userEmail, name: userName, uid: uid });
    navigateTo('/home');
  } catch (error) {
    // 실패 시 에러 메시지 출력
    if (error.code === 'auth/user-not-found') {
      alert('등록되지 않은 이메일입니다.');
    } else if (error.code === 'auth/wrong-password') {
      alert('비밀번호가 잘못되었습니다.');
    } else if (error.code === 'auth/invalid-login-credentials') {
      alert('로그인 정보가 유효하지 않습니다. 다시 확인해주세요.');
    } else {
      console.log(error.message);
    }
  }
}

// //mysql 회원조회
// const login = async () => {
//   console.log('Logging in with', user_id.value, password.value);
//   try {
//     const response = await axios.post('http://localhost:8080/user/login', {
//       user_id: user_id.value,
//       password: password.value
//     });
//     console.log(response.data);
//     if (response.data.success === "Y") {
//       auth.login(response.data.result, response.data.token);
//       // navigateTo 함수를 호출하려면 여기에 import하거나 define 해야 합니다.
//       navigateTo('/home');
//     } else {
//       alert('아이디와 비밀번호를 확인하세요.');
//     }
//   } catch (error) {
//     console.error(error);
//   }
// };
</script>

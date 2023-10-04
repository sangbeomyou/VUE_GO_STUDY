<template>
  <v-container>
    <v-row justify="center" align="center" class="fill-height">
      <v-col cols="12" sm="8" md="6">
        <v-card class="elevation-12">
          <v-toolbar color="primary" dark flat>
            <v-toolbar-title>로그인</v-toolbar-title>
          </v-toolbar>
          <v-card-text>
            <v-text-field label="아이디" id="user_id" name="user_id" prepend-icon="mdi-account" type="text"
              :rules="[rules.required]" required v-model="user_id"></v-text-field>
            <v-text-field label="비밀번호" id="Password" name="password" prepend-icon="mdi-lock" type="password"
              :rules="[rules.required]" required v-model="password"></v-text-field>
          </v-card-text>
          <v-card-actions>
            <v-btn color="primary" @click="login">로그인</v-btn>
          </v-card-actions>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import axios from 'axios'
import { useAuthStore } from '~/stores/auth';
const auth = useAuthStore();


definePageMeta({
  layout: false,
  middleware: false
});

export default {
  data: () => ({
    user_id: '',
    password: '',
    rules: {
      required: value => !!value || '빈칸을 입력하세요.',
    },
  }),
  methods: {
    async login() {
      console.log('Logging in with', this.user_id, this.password);
      try {
        const response = await axios.post('http://localhost:8080/user/login', {
          user_id: this.user_id,
          password: this.password
        });
        console.log(response.data)
        if (response.data.success === "Y") {
          auth.login(response.data.result, response.data.token);
          navigateTo('/home');  
        } else {
          alert('아이디와 비밀번호를 확인하세요.'); 
        }
      } catch (error) {
        console.error(error); 
      }
    }
  }
}
</script>
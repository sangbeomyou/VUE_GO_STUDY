<template>
    <v-container>
        <v-row justify="center" align="center" class="fill-height">
            <v-col cols="12" sm="8" md="6">
                <v-card class="elevation-12">
                    <v-toolbar color="primary" dark flat>
                        <v-toolbar-title>회원 가입</v-toolbar-title>
                    </v-toolbar>
                    <v-card-text>
                        <v-text-field label="이메일" v-model="email" prepend-icon="mdi-email" :error-messages="emailError"
                        @input="validateEmail"></v-text-field>
                        <v-text-field label="비밀번호" type="password" v-model="password" prepend-icon="mdi-lock"
                            :error-messages="passwordError"></v-text-field>
                        <v-text-field label="비밀번호 확인" type="password" v-model="confirmPassword" :error-messages="confirmPasswordError"
                            prepend-icon="mdi-lock-check" @input="validatePassword"></v-text-field>
                    </v-card-text>
                    <v-card-actions>
                        <v-btn color="primary" @click="register">회원 가입</v-btn>
                    </v-card-actions>
                </v-card>
            </v-col>
        </v-row>
    </v-container>
</template>
  
<script setup>
import { ref } from 'vue';
import useAuth from '~/composables/useAuth'

const email = ref('');
const password = ref('');
const confirmPassword = ref('');

const emailError = ref('');
const passwordError = ref('');
const confirmPasswordError = ref('');

const { signUp } = useAuth();

// 비밀번호 유효성 검사
const validatePassword = () => {
    if (password.value !== confirmPassword.value) {
        passwordError.value = '비밀번호가 일치하지 않습니다.';
    } else {
        passwordError.value = '';
    }
};
// 이메일 정규식
const isValidEmail = (email) => {
    const emailPattern = /^[a-zA-Z0-9._-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,6}$/;
    return emailPattern.test(email);
}
const validateEmail = () => {
    if (!isValidEmail(email.value)) {
        emailError.value = '올바른 이메일 형식이 아닙니다.';
    } else {
        emailError.value = '';
    }
};
// 회원 가입 로직
const register = async  () => {

    // 이메일 빈 값 체크
    if (!email.value) {
        emailError.value = '이메일을 입력하세요.';
        return
    } else {
        emailError.value = '';
    }
    // 비밀번호 빈 값 체크
    if (!password.value) {
        passwordError.value = '비밀번호을 입력하세요.';
        return
    } else {
        passwordError.value = '';
    }
    // 비밀번호 빈 값 체크 
    if (!confirmPassword.value) {
        confirmPasswordError.value = '비밀번호을 입력하세요.';
        return
    } else {
        confirmPasswordError.value = '';
    }

    // TODO: 여기에 회원 가입 로직을 추가하세요 (API 호출 등)
    console.log(`회원가입 요청: ${password.value}, ${email.value}`);
    try {
        // 회원가입 로직 처리
        await signUp(email.value, password.value);  // 가정: signUp 함수는 Promise를 반환
        console.log(`회원가입 성공: ${email.value}, ${password.value}`);
    } catch (error) {
        console.error('회원가입 실패:', error);
        // 실패 시 사용자에게 보여줄 에러 메세지 처리 등...
    }
}

definePageMeta({
    title: '회원 가입',
    layout: false,
    middleware: false
});

</script>
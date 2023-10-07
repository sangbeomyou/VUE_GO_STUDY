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
                        <v-text-field label="이름" v-model="name" prepend-icon="mdi-email"
                            :error-messages="nameError"></v-text-field>
                        <v-text-field label="비밀번호" type="password" v-model="password" prepend-icon="mdi-lock"
                            :error-messages="passwordError"></v-text-field>
                        <v-text-field label="비밀번호 확인" type="password" v-model="confirmPassword"
                            :error-messages="confirmPasswordError" prepend-icon="mdi-lock-check"
                            @input="validatePassword"></v-text-field>
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
const name = ref('');
const password = ref('');
const confirmPassword = ref('');

const emailError = ref('');
const nameError = ref('');
const passwordError = ref('');
const confirmPasswordError = ref('');

const { signUp } = useAuth();

// 이메일 정규식
const isValidPassword = (password) => {
    const passwordPattern = /^(?=.*[a-z])(?=.*\d)(?=.*[!@#$%^&*])[A-Za-z\d!@#$%^&*]{8,}$/;
    return passwordPattern.test(password);
}
// 비밀번호 유효성 검사
const validatePassword = () => {
    if (password.value !== confirmPassword.value) {
        passwordError.value = '비밀번호가 일치하지 않습니다.';
        return;
    }

    if (!isValidPassword(password.value)) {
        passwordError.value = '올바른 비밀번호 형식을 사용해주세요.';
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
const register = async () => {
    // 이메일 빈 값 체크
    if (!email.value) {
        emailError.value = '이메일을 입력하세요.';
        return
    } else {
        emailError.value = '';
    }
    // 이름 빈 값 체크
    if (!name.value) {
        nameError.value = '이름을 입력하세요.';
        return
    } else {
        nameError.value = '';
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

    // 회원가입 로직 처리
    try {
        await signUp(email.value, name.value, password.value);
        alert('회원가입 완료');
        //console.log(`회원가입 성공: ${email.value}`);
        navigateTo('/login');
    } catch (error) {
        if (error.code === 'auth/email-already-in-use') {
            alert('이미 사용 중인 이메일입니다.');
        } else {
            console.log(error.message);
        }
    }
}

definePageMeta({
    title: '회원 가입',
    layout: false,
    middleware: false
});

</script>
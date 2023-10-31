<template>
    <v-spacer></v-spacer>
    <div class="centered-content">
        <div v-if="user">
            {{ user.name }} | 접속 일시: {{ loginDate }}
            <v-btn color="primary" @click="logoutclick">로그아웃</v-btn>
        </div>
    </div>
</template>

<script setup>
import axios from 'axios';

import { useAuthStore } from '~/stores/auth';
import useAuth from '~/composables/useAuth'; //파이어 베이스

const { logout } = useAuth();

const auth = useAuthStore();
const user = auth.user;
const loginDate = auth.loginDate;

const logoutclick = async () => {
    logout(); // 파이어 베이스 로그아웃
    auth.logout();  // store 로그아웃
    await axios.get('http://localhost:8080/public/logout', {withCredentials: true}); // 쿠키 지우기 ㅠㅠ 3번을 해야하네
}
</script>
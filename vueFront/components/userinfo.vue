<template>
    <v-spacer></v-spacer>
    <div class="centered-content">
        <div v-if="user">
            {{ user.name }} | 접속 일시: {{ loginDate }}
            <v-btn color="primary" @click="logout">로그아웃</v-btn>
        </div>
    </div>
</template>

<script setup>
import { useAuthStore } from '~/stores/auth';

const auth = useAuthStore();
const user = auth.user;
const loginDate = auth.loginDate;
const logout = auth.logout;
</script>
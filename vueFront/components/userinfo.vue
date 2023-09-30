<script>
import { user, loginDate, logout } from '~/composables/user';


</script>
<template>
    <v-spacer></v-spacer>
    <div class="centered-content">
        <div v-if="user">
            {{ user[0].user_id }} | 접속 일시: {{ loginDate }}

            <v-btn color="primary" @click="logout">로그아웃</v-btn>
        </div>

    </div>
</template>

<template>
    <v-container>
        <v-row v-if="isLoading" justify="center">
            <v-col cols="12" class="text-center">
                <v-progress-circular indeterminate></v-progress-circular>
            </v-col>
        </v-row>

        <v-row v-else>
            <v-col cols="12">
                <v-card>
                    <v-card-title>
                        <v-row>
                            <v-col cols="8" class="headline">게시글 수정</v-col>
                        </v-row>
                    </v-card-title>

                    <v-card-text>
                        <v-text-field v-model="title" label="제목" required></v-text-field>
                        <v-textarea v-model="content" label="내용" required></v-textarea>
                    </v-card-text>
                    <v-card-actions>
                        <v-btn color="primary" @click="submitPost">수정하기</v-btn>
                    </v-card-actions>
                </v-card>
            </v-col>
        </v-row>
    </v-container>
</template>
  
<script setup>
import axios from 'axios';

import { ref } from 'vue'
import { useAuthStore } from '~/stores/auth';

const auth = useAuthStore();
const user = auth.user;

const route = useRoute();
const idx = ref(route.params.idx)

const title = ref('');
const content = ref('');
const isLoading = ref(false);
const bbs = ref(null);
// idx 로 값 불러오기
const fetchData = async () => {
    try {
        const response = await axios.get(`http://localhost:8080/bbs/bbsInfo/${idx.value}`, { withCredentials: true })
        if (response.data.success === "Y") {
            bbs.value = response.data.result;
            if(user.uid !== bbs.value.UserID)
            {
                alert("수정 권한이 없습니다.");
                navigateTo('/board');
                return;
            }
            console.log(bbs.value);
            title.value = bbs.value.Title; // 불러온 제목 데이터를 title에 할당
            content.value = bbs.value.Content; // 불러온 내용 데이터를 content에 할당
        }
        console.log(response)
    } catch (error) {
        console.error("Error fetching data:", error)
    } finally {
        isLoading.value = false
    }
}

const submitPost = async () => {
    if (!title.value.trim()) {
        alert("제목을 입력해주세요.");
        return;
    }

    if (!content.value.trim()) {
        alert("내용을 입력해주세요.");
        return;
    }

    // 데이터를 서버로 전송
    try {
        isLoading.value = true; // 로딩 상태 시작
        const response = await axios.post('http://localhost:8080/bbs/edit', {
            Idx : idx.value,
            Title: title.value,
            Content: content.value,
        },
            { withCredentials: true }
        );
        if (response.data.Success === "Y") {
            alert("게시글이 수정되었습니다.");
            navigateTo('/board');
        } else {
            alert("게시글 작성에 실패하였습니다. 다시 시도해주세요.");
        }
    } catch (error) {
        console.error("게시글 작성 중 에러 발생:", error);
        alert("게시글 작성 중 문제가 발생하였습니다. 다시 시도해주세요.");
    } finally {
        isLoading.value = false; // 로딩 상태 종료
    }
}

fetchData()
</script>
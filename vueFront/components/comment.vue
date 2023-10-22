<template>
    <div>
        <!-- 댓글 목록 섹션 -->
        <v-row v-for="(comment, index) in comments" :key="index">
            <v-col>
                <v-card outlined class="my-2">
                    <v-card-subtitle class="d-flex justify-space-between align-center">
                        <div>{{ comment.UserName }}</div>
                        <div class="text-end">{{ comment.RegDate }}</div>
                    </v-card-subtitle>
                    <v-card-text class="d-flex justify-space-between align-center">
                        <div>{{ comment.Content }}</div>
                        <div>
                            <v-btn flat class="text-primary" @click="editComment(comment.UserID)"
                                v-if="isAuthor(comment.UserID)">수정</v-btn>
                            <v-btn flat class="text-red" @click="deleteComment(comment.UserID)"
                                v-if="isAuthor(comment.UserID)">삭제</v-btn>
                        </div>
                    </v-card-text>
                    <!-- 대댓글 추가 버튼 또는 다른 인터랙션 요소를 이곳에 배치 -->
                </v-card>
            </v-col>
        </v-row>
        <!-- 댓글 입력 섹션 -->
        <v-row class="my-3">
            <v-col>
                <v-textarea v-model="newCommentText" label="댓글 입력" outlined auto-grow rows="3"></v-textarea>
            </v-col>
        </v-row>
        <v-row class="mb-3">
            <v-col class="text-end">
                <v-btn color="primary" @click="postComment">등록</v-btn>
            </v-col>
        </v-row>
    </div>
</template>
  
<script setup>
import axios from 'axios'
import { onMounted, ref, defineProps } from 'vue';
import { useAuthStore } from '~/stores/auth';


const auth = useAuthStore();
const user = auth.user;

// props 정의
const props = defineProps({
    bbsId: String,
    required: true,
});

const isLoading = ref(false)
const newCommentText = ref(''); // 댓글 쓰기
const page = ref(1); //페이징
const comments = ref(null); //댓글

// 로그인 사용자 확인
const isAuthor = (UserID) => {
    return user.uid === UserID;
};

//댓글 불러오기
const fetchData = async () => {
    isLoading.value = true; // 로딩 상태 시작
    try {
        const response = await axios.get('http://localhost:8080/bbs/comment', {
            params: { page: page.value, bbsId: props.bbsId },
            withCredentials: true
        }
        );

        if (response.data.success === "Y") {
            console.log(response.data.result)
            comments.value = response.data.result;
        }
        console.log(response)
    } catch (error) {
        console.error("Error fetching data:", error)
    } finally {
        isLoading.value = false
    }
};
//댓글 수정하기
const editComment = async (UserID) => {
    console.log(UserID)
}
//댓글 삭제
const deleteComment = async (UserID) => {
    console.log(UserID)
}
// 댓글 게시 기능
const postComment = async () => {
    if (newCommentText.value.trim() === '') {
        alert('댓글을 입력해 주세요');
        return;
    }
    try {
        isLoading.value = true; // 로딩 상태 시작
        const response = await axios.post('http://localhost:8080/bbs/commentWrite', {
            BbsId: props.bbsId,
            Content: newCommentText.value,
            UserID: user.uid,
            UserName: user.name
        },
            { withCredentials: true }
        );

        if (response.data.success === "Y") {
            await fetchData(); // 작성 성공하면 다시 불러오기
        } else {
            alert("댓글 작성이 실패했습니다.")
        }
        console.log(response)
    } catch (error) {
        console.error("Error fetching data:", error)
    } finally {
        isLoading.value = false
    }
    // 댓글 입력란 초기화
    newCommentText.value = '';
};
onMounted(async () => {
    await fetchData();
});
</script>
  
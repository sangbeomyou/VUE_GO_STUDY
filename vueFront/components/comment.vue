<template>
    <div>
        <!-- 댓글 목록 섹션 -->
        <v-row v-for="(comment, index) in comments" :key="index">
            <v-col>
                <v-card outlined class="my-2">
                    <v-card-subtitle>
                        {{ comment.author }}
                    </v-card-subtitle>
                    <v-card-text>
                        {{ comment.content }}
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
                <v-btn color="primary" @click="postComment">Post Comment</v-btn>
            </v-col>
        </v-row>
    </div>
</template>
  
<script setup>
import { ref } from 'vue';

const newCommentText = ref('');

const comments = ref([
    {
        author: '테스트',
        content: '테스트11111',
    },
]);

// 댓글 게시 기능
const postComment = () => {
    if (newCommentText.value.trim() === '') {
        alert('댓글을 입력해 주세요');
        return;
    }

    // 서버에 새 댓글을 전송하는 로직을 이곳에 작성해야 함
    // 이 예제에서는 클라이언트 사이드에서만 처리

    comments.value.push({
        author: 'You', // 실제 사용자 이름으로 교체해야 함
        content: newCommentText.value.trim(),
    });

    // 댓글 입력란 초기화
    newCommentText.value = '';
};
</script>
  
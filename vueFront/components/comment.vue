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
                        <div v-if="isEdit(comment.Idx)">
                            {{ comment.Content }}
                        </div>
                        <div v-else class="w-100">
                            <v-textarea v-model="editContents[comment.Idx]" outlined auto-grow solo rows="3" class="w-100" ></v-textarea>
                            <v-btn @click="editComment(comment.Idx, editContents[comment.Idx])">저장</v-btn>
                        </div>
                        <div v-if="isAuthor(comment.UserID) && isEdit(comment.Idx)">
                            <v-btn flat class="text-primary" @click="startEditing(comment)">수정</v-btn>
                            <v-btn flat class="text-red" @click="deleteComment(comment.Idx)">삭제</v-btn>
                        </div>
                    </v-card-text>
                    <!-- 대댓글 추가 버튼 또는 다른 인터랙션 요소를 이곳에 배치 -->
                </v-card>
            </v-col>
        </v-row>

        <v-pagination v-model="page" :length="pages" @update:model-value="updatePage" :total-visible="5"
            rounded="circle"></v-pagination>

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
import { onMounted, reactive, ref, defineProps } from 'vue';
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
const pages = ref(null); //전체 페이지
const comments = ref(null); //댓글
const limit = ref(5); //보여주는 댓글 수
const total = ref(null); //전체 댓글 수
const editingIdx = ref(null); //수정할 댓글 아이디
const editContents = ref({}); // 수정 댓글

// 로그인 사용자 확인
const isAuthor = (UserID) => {
    return user.uid === UserID;
};

//댓글 불러오기
const fetchData = async () => {
    if (isLoading.value) return;

    isLoading.value = true; // 로딩 상태 시작
    try {
        const response = await axios.get('http://localhost:8080/bbs/comment', {
            params: { page: page.value, limit: limit.value, bbsId: props.bbsId },
            withCredentials: true
        }
        );

        if (response.data.success === "Y") {
            if (response.data.result.length) {
                comments.value = response.data.result;
                page.value = response.data.page; // 페이지 세팅
                pages.value = response.data.pages; // 전체 페이지 세팅
                total.value = response.data.total; // 페이지 세팅
            }
        }
        console.log(response)
    } catch (error) {
        console.error("Error fetching data:", error)
    } finally {
        isLoading.value = false
    }
};


// 댓글 페이지가 변경 
const updatePage = (newPage) => {
    console.log(newPage)
    page.value = newPage;
    fetchData(); //댓글 다시 불러오기
};

//댓글 수정하기
const editComment = async (Idx, Content) => {
    if (!Content.trim()) {
        alert("댓글을 입력해주세요.");
        return;
    }

    // 데이터를 서버로 전송
    try {
        isLoading.value = true; // 로딩 상태 시작
        const response = await axios.post('http://localhost:8080/bbs/commentEdit', {
            Idx : String(Idx),
            Content: Content,
        },
            { withCredentials: true }
        );
        console.log(response.data)
        if (response.data.success === "Y") {
            alert("댓글이 수정되었습니다.");
            editingIdx.value = null; // 댓글 수정 창 종료
            isLoading.value = false; // 로딩 상태 종료
            fetchData(); 
        } else {
            alert("댓글 수정에 실패하였습니다. 다시 시도해주세요.");
        }
    } catch (error) {
        console.error("댓글 작성 중 에러 발생:", error);
        alert("댓글 작성 중 문제가 발생하였습니다. 다시 시도해주세요.");
    } finally {
        isLoading.value = false; // 로딩 상태 종료
    }
}

const startEditing = (comment) => {
    editingIdx.value = comment.Idx; // 수정할 댓글 idx 세팅
    editContents.value[comment.Idx] = comment.Content; // 수정할 댓글 세팅
}
// 수정 아이디 확인
const isEdit = (Idx) => {
    return editingIdx.value !== Idx;
};

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
            isLoading.value = false
            // 배수로 딱 맞아떨어지면 다음 페이지를 보여줘야함
            if (total.value % limit.value === 0) {
                pages.value += 1;
            }
            updatePage(pages.value)
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
  
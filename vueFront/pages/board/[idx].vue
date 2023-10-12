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
              <v-col cols="8" class="headline">{{ bbs.Title }}</v-col>

              <v-col cols="4" class="text-end small-text">
                <span><strong>작성자:</strong> {{ bbs.UserName }}</span>
                <span class="ml-4"><strong>작성일:</strong> {{ bbs.RegDate }}</span>
              </v-col>
            </v-row>
          </v-card-title>

          <v-card-text>
            <v-textarea readonly v-model="bbs.Content"></v-textarea>
          </v-card-text>
          <!-- 내 아이디일 경우 수정버튼 활성 -->
          <v-card-actions v-if="bbs.UserID === user.uid">
            <v-btn color="primary" @click="navigateToWrite">수정하기</v-btn>
            <v-btn color="error" @click="deleteBbs">삭제하기</v-btn>
          </v-card-actions>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<style>
/* 작성자와 작성일 텍스트 크기를 줄이는 CSS */
.small-text {
  font-size: 0.8rem;
}
</style>


  
<script setup>
import axios from 'axios'
import { ref } from 'vue'
import { useRoute } from 'vue-router'
import { useAuthStore } from '~/stores/auth';

const auth = useAuthStore();
const user = auth.user;

const route = useRoute()
const idx = ref(route.params.idx)

const bbs = ref(null)
const isLoading = ref(true)

const fetchData = async () => {
  try {
    const response = await axios.get(`http://localhost:8080/bbs/bbsInfo/${idx.value}`, { withCredentials: true })
    if (response.data.success === "Y") {
      console.log(response.data.result)
      bbs.value = response.data.result;
    } else {
      bbs.value = [];
    }
    console.log(response)
  } catch (error) {
    console.error("Error fetching data:", error)
  } finally {
    isLoading.value = false
  }
}

const deleteBbs = async () => {
  if (window.confirm("게시글을 정말로 삭제하시겠습니까?")) {
    try {
      const response = await axios.post(`http://localhost:8080/bbs/delete`, {
        Idx: idx.value
      }, { withCredentials: true });
      if (response.data.Success === "Y") {
        alert("게시글이 성공적으로 삭제되었습니다.");
        navigateTo('/board');
      } else {
        console.log(response)
        alert("게시글 삭제에 실패하였습니다. 다시 시도해주세요.");
      }
    } catch (error) {
      console.error("게시글 삭제 중 에러 발생:", error);
      alert("게시글 삭제 중 문제가 발생하였습니다. 다시 시도해주세요.");
    }
  }
}
const navigateToWrite = () => {
    // idx 값과 함께 /bbs/write 페이지로 이동
    navigateTo(`/bbs/write${idx.value}`);
}


fetchData()
</script>
  
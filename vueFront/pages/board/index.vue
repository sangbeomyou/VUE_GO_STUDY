<template>
  <div>
    <div v-if="isLoading">
      <v-progress-circular indeterminate></v-progress-circular>
    </div>
    <div class="text-end"> <!-- text-end 클래스를 사용하여 내용을 오른쪽 정렬 -->
      <NuxtLink :to="{ name: 'board-write' }">
        <v-btn color="primary">글쓰기</v-btn>
      </NuxtLink>
    </div>
    <v-table>
      <thead>
        <tr>
          <th class="text-left">
            번호
          </th>
          <th class="text-left">
            제목
          </th>
          <th class="text-left">
            작성자
          </th>
          <th class="text-left">
            작성 시간
          </th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="bbs in bbsList" :key="bbs.Idx">
          <td>{{ bbs.Idx }}</td>
          <td>
            <NuxtLink :to="{ name: 'board-info-idx', params: { idx: bbs.Idx } }">{{ bbs.Title }}</NuxtLink>
          </td>
          <td>{{ bbs.UserName }}</td>
          <td>{{ bbs.RegDate }}</td>
        </tr>
      </tbody>
    </v-table>

    <div>
      <v-pagination
      v-model="page"
      :length="pages"
      @update:model-value ="updatePage"
      rounded="circle"
    ></v-pagination>

    </div>
  </div>
</template>
<script setup>
import axios from 'axios'
import { ref } from 'vue'

const isLoading = ref(false); // 데이터 로딩 상태
const bbsList = ref(null); //게시판 리스트
const page = ref(1); //현재 페이지
const pages = ref(null); //전체 페이지
const limit = ref(10); //최대 갯수

const fetchData = async () => {
  isLoading.value = true; // 데이터 로딩 시작
  try {
    // 서버로부터 데이터를 받아오는 부분에서 현재 페이지와 limit 값을 사용합니다.
    const response = await axios.get('http://localhost:8080/bbs/bbsList', {
      params: { page: page.value, limit: limit.value },
      withCredentials: true
    });
    if (response.data.success === "Y") {
      console.log(response.data);
      pages.value =response.data.pages;
      bbsList.value = response.data.result;
    } else {
      console.log(response);
      bbsList.value = []; // 실패한 경우 초기화
    }
  } catch (error) {
    console.error('API 호출 오류:', error); // 콘솔에 오류 메시지 출력
    bbsList.value = []; // 오류 발생 시 초기화
  } finally {
    isLoading.value = false; // 어떤 상황에서든 로딩 상태는 false로 설정
  }
}

// 페이지가 변경 
const updatePage = (newPage) => {
  console.log(newPage)
  page.value = newPage;
  navigateTo(`/board/${page.value}`);
  //fetchData();
};


fetchData();
</script>
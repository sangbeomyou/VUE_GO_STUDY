<template>
  <div>
    <div v-if="isLoading">
      <v-progress-circular indeterminate></v-progress-circular>
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
            <NuxtLink :to="{ name: 'board-idx', params: { idx: bbs.Idx } }">{{ bbs.Title }}</NuxtLink>
          </td>
          <td>{{ bbs.UserName }}</td>
          <td>{{ bbs.RegDate }}</td>
        </tr>
      </tbody>
    </v-table>
    <NuxtLink :to="{ name: 'board-write' }">
      <v-btn color="primary">글쓰기</v-btn>
    </NuxtLink>
    </div>

</template>
<script>
import axios from 'axios'

export default {
  data() {
    return {
      isLoading: false, // 데이터 로딩 상태
      bbsList: [],
      page: 1,       // 현재 페이지
      limit: 10,     // 페이지 당 게시물 수
    }
  },
  async created() {
    this.isLoading = true; // 데이터 로딩 시작
    try {
      // 서버로부터 데이터를 받아오는 부분에서 현재 페이지와 limit 값을 사용합니다.
      const response = await axios.get('http://localhost:8080/bbs/bbsList', { 
        params: { page: this.page, limit: this.limit },
        withCredentials: true 
      });

      if (response.data.success === "Y") {
        console.log(response.data);
        this.bbsList = response.data.result; 
      } else {
        console.log(response);
        this.bbsList = []; // 실패한 경우 초기화
      }
    } catch (error) {
      console.error('API 호출 오류:', error); // 콘솔에 오류 메시지 출력
      this.bbsList = []; // 오류 발생 시 초기화
    }  finally {
      this.isLoading = false; // 어떤 상황에서든 로딩 상태는 false로 설정
    }
  }
}

</script>
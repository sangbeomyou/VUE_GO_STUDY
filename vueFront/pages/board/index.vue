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
      isLoading: true, // 데이터 로딩 상태
      bbsList: []
    }
  },
  async created() {
    try {
      const response = await axios.get('http://localhost:8080/bbs/bbsList', { withCredentials: true })
      if (response.data.success === "Y") {
        console.log(response.data.result)
        this.bbsList = response.data.result
      } else {
        console.log(response)
        this.bbsList = []  // 실패한 경우 초기화
      }
      this.isLoading = false // 데이터 로딩 완료
    } catch (error) {
      console.error('API 호출 오류:', error)
      this.isLoading = false // 데이터 로딩 실패
      this.bbsList = []      // 오류 발생 시 초기화
    }
  }
}

</script>
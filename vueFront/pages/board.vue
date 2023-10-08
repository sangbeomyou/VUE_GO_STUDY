<template>
  <div>
    <div v-if="isLoading">로딩 중...</div>
    <v-list v-else two-line>
      <v-list-item v-for="bbs in bbsList" :key="bbs.id">
        <v-list-item-content>
          <v-list-item-title>{{ bbs.Title }} | {{ bbs.UserID }} | {{ bbs.RegDate }}</v-list-item-title>
          <v-list-item-subtitle>
            {{ bbs.Content }}
          </v-list-item-subtitle>
        </v-list-item-content>
        <v-list-item-action>
          <v-icon>mdi-pencil</v-icon>
        </v-list-item-action>
      </v-list-item>
    </v-list>
  </div>
</template>
<script>
import axios from 'axios'

export default {
  data() {
    return {
      isLoading: true, // 데이터 로딩 상태
      posts: []
    }
  },
  async created() {
    try {
      const response = await axios.get('http://localhost:8080/bbs/bbsList', {withCredentials: true})

      if (response.data.success === "Y") {
        console.log(response.data.result)
        this.bbsList = response.data.result
      }
      this.isLoading = false // 데이터 로딩 완료
    } catch (error) {
      console.error('API 호출 오류:', error)
      this.isLoading = false // 데이터 로딩 실패
    }
  }
}

</script>
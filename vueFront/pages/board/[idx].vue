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
  import { ref } from 'vue'
  import { useRoute } from 'vue-router'
  import axios from 'axios'
  
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
  
  fetchData()
  </script>
  
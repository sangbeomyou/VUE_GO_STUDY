<template>
    <div>
      <div v-if="isLoading">Loading...</div>
      <div v-else>{{ data }}</div>
    </div>
  </template>
  
  <script setup>
  import { ref } from 'vue'
  import { useRoute } from 'vue-router'
  import axios from 'axios'
  
  const route = useRoute()
  const idx = ref(route.params.idx)  
  
  const data = ref(null)
  const isLoading = ref(true)
  
  const fetchData = async () => {
    try {
      const response = await axios.get(`http://localhost:8080/bbs/bbsInfo/${idx.value}`, { withCredentials: true })
      data.value = response.data
    } catch (error) {
      console.error("Error fetching data:", error)
    } finally {
      isLoading.value = false
    }
  }
  
  fetchData()
  </script>
  
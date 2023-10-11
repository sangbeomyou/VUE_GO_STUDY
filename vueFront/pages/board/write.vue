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
                            <v-col cols="8" class="headline">게시글 작성</v-col>
                        </v-row>
                    </v-card-title>

                    <v-card-text>
                        <v-text-field v-model="title" label="제목" required></v-text-field>
                        <v-textarea v-model="content" label="내용" required></v-textarea>
                    </v-card-text>
                    <v-card-actions>
                        <v-btn color="primary" @click="submitPost">게시하기</v-btn>
                    </v-card-actions>
                </v-card>
            </v-col>
        </v-row>
    </v-container>
</template>
  
<script setup>
import { ref } from 'vue'
import { useAuthStore } from '~/stores/auth';

const auth = useAuthStore();
const user = auth.user;

const title = ref('');
const content = ref('');
const isLoading = ref(false);

    const submitPost = () => {
  if (!title.value.trim()) {
    alert("제목을 입력해주세요.");
    return;
  }
  
  if (!content.value.trim()) {
    alert("내용을 입력해주세요.");
    return;
  }

  console.log(title.value, content.value,  user.uid, user.name, user.email);
}


</script>
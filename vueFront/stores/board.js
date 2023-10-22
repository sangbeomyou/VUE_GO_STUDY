import { defineStore } from 'pinia';
import axios from 'axios'; 

export const useBoardStore = defineStore({
    id: 'board',
    state: () => ({
        bbsList: [],  // 게시물 리스트
        bbsInfo: null,  // 상세 페이지의 현재 게시물
        page : 1, // 현제 페이지
        pages : null, // 전체 페이지
        limit : 10, // 게시글 보이는 최대 갯수
        isLoading: false,  // 게시물 로딩 중 여부
    }),
    actions: {
        // 게시물 불러오기
        async fetchBoard() {
            this.isLoading = true;
            try {
                const response = await axios.get('http://localhost:8080/bbs/bbsList', {
                    params: { page: this.page, limit: this.limit },
                  withCredentials: true
                });
                if (response.data.success === "Y") {
                    this.pages = response.data.pages;
                    this.bbsList = response.data.result;
                } else {
                  console.log(response);
                  this.bbsList = []; // 데이터 로드 실패시 초기화
                }
              } catch (error) {
                console.error('API 호출 오류:', error); // 콘솔에 오류 메시지 출력
                this.bbsList = []; // 오류 발생시 리스트 초기화
              } finally {
                this.isLoading = false; // 로딩 상태 종료
              }
        },

        // 게시물 상세 정보 불러오기
        async fetchPostDetail(postId) {
            this.isLoading = true;
            try {
                const response = await axios.get(`/api/posts/${postId}`);  // 실제 API 엔드포인트로 변경 필요
                this.currentPost = response.data;
            } catch (error) {
                console.error('An error occurred while fetching the post detail:', error);
                // 오류 처리 로직 추가
            } finally {
                this.isLoading = false;
            }
        },

        // 게시물 추가, 수정, 삭제 등의 메서드도 여기에 정의합니다.
    },
});

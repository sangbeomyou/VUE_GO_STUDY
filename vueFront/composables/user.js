import { ref } from 'vue';

const user = ref(null);
const loginDate = ref(null);
const isAuthenticated = ref(false);
//로그인 정보 저장
const setUser = (newUser) => {
  user.value = newUser;
};
//로그인 일시 저장
const setLoginDate = () => {
  loginDate.value = new Date().toLocaleString();
};
//로그아웃 정보 지우기
const logout = () => {
  user.value = null;
  loginDate.value = null;
  isAuthenticated.value = false;
  navigateTo('/login')
};

const setAuthenticated = (authStatus) => {
  isAuthenticated.value = authStatus;
};

const login = (newUser) => {
  setUser(newUser);
  setLoginDate();
  setAuthenticated(true);
};

export {
  user,
  loginDate,
  isAuthenticated,
  login,
  logout
};

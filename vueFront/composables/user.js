import { ref } from 'vue';

const user = ref(null);
const loginDate = ref(null);
const isAuthenticated = ref(false);

const setUser = (newUser) => {
  user.value = newUser;
};

const setLoginDate = () => {
  loginDate.value = new Date().toLocaleString();
};

const logout = () => {
  user.value = null;
  loginDate.value = null;
  isAuthenticated.value = false;
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

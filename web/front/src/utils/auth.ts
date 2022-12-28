const TokenKey = 'piniaUser';
const TokenPrefix = 'Bearer ';
const isLogin = () => {
  return !!localStorage.getItem(TokenKey);
};
const getToken = () => {
  return localStorage.getItem(TokenKey);
};
const setToken = (token: string) => {
  localStorage.setItem(TokenKey, token);
};
const clearToken = () => {
  localStorage.removeItem(TokenKey);
};
const clearAll = () => {
  localStorage.clear();
};
export { TokenPrefix, isLogin, getToken, setToken, clearToken ,clearAll};

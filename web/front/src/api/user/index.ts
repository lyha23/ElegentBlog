// 权限问题后期增加
import { get, post } from '/@/utils/http/axios';
enum URL {
  login='/web/login',
  wxlogin = '/web/wxLogin',
  logout = '/user/logout',
  profile = '/user/profile',
}
interface LoginRes {
  token: string;
}   

export interface LoginData {
  code: string; 
}

const login =async()=>get({url:URL.login})
const wxlogin = async (params: LoginData) => get<any>({ url: URL.wxlogin, params });
const getUserProfile = async () => get({ url: URL.profile });
const logout = async () => post<LoginRes>({ url: URL.logout });
export { login, wxlogin,logout,getUserProfile };

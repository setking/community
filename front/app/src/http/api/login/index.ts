import type * as Users from './type'
import { request } from '@/http/axios.ts'

/** 获取当前登录用户详情 */
export function LoginApi (data: Users.UserResponseData) {
  return request<Users.UserResponseData>({
    url: '/user/login',
    method: 'post',
    data,
  })
}
export function RegisterApi (data: Users.UserResponseData) {
  return request<Users.UserResponseData>({
    url: '/user/signup',
    method: 'post',
    data,
  })
}

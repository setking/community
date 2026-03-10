import type * as Post from './type'
import { request } from '@/http/axios.ts'

/** 获取当前登录用户详情 */
export function PostListApi (params: Post.PostResponseData) {
  return request<Post.PostResponseData>({
    url: '/post/list',
    method: 'get',
    params,
  })
}

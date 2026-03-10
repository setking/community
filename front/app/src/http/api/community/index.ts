import type * as Community from './type'
import { request } from '@/http/axios.ts'

/** 获取当前登录用户详情 */
export function CommunityListApi (params: Community.CommunityResponseData) {
  return request<Community.CommunityResponseData>({
    url: '/community/list',
    method: 'get',
    params,
  })
}

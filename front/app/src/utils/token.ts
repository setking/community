import Cookies from 'js-cookie'
const SYSTEM_NAME = 'community'
/** 缓存数据时用到的 Key */
export class CacheKey {
  static readonly TOKEN = `${SYSTEM_NAME}-token-key`
}

export function getToken () {
  return Cookies.get(CacheKey.TOKEN)
}

export function setToken (token: string) {
  Cookies.set(CacheKey.TOKEN, token)
}

export function removeToken () {
  Cookies.remove(CacheKey.TOKEN)
}

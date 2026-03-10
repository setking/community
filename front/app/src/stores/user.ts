// Utilities
import { defineStore } from 'pinia'
import { pinia } from './index'

export const useUserStore = defineStore('user', () => {
  const token = ref<string>('')
  const logout = () => {
    token.value = ''
  }
})
/**
 * @description 在 SPA 应用中可用于在 pinia 实例被激活前使用 store
 * @description 在 SSR 应用中可用于在 setup 外使用 store
 */
export function useUserStoreOutside () {
  return useUserStore(pinia)
}

/**
 * router/index.ts
 *
 * Automatic routes for `./src/pages/*.vue`
 */

import type { RouteRecordRaw } from 'vue-router'
// Composables
import { setupLayouts } from 'virtual:generated-layouts'
import { createRouter, createWebHistory } from 'vue-router'
const routes: RouteRecordRaw[] = [
  {
    path: '/',
    redirect: '/home',
  },
  {
    path: '/',
    component: () => import('../layouts/index.vue'),
    children: [
      {
        path: 'home',
        name: 'home',
        component: () => import('../pages/home/index.vue'),
      },
    ],
  },
  {
    path: '/',
    component: () => import('../layouts/index.vue'),
    children: [
      {
        path: 'detail/:id',
        name: 'detail',
        component: () => import('../pages/detail/index.vue'),
      },
    ],
  },
]
const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: setupLayouts(routes),
})

// Workaround for https://github.com/vitejs/vite/issues/11804
router.onError((err, to) => {
  if (err?.message?.includes?.('Failed to fetch dynamically imported module')) {
    if (localStorage.getItem('vuetify:dynamic-reload')) {
      console.error('Dynamic import error, reloading page did not fix it', err)
    } else {
      console.log('Reloading page to fix dynamic import error')
      localStorage.setItem('vuetify:dynamic-reload', 'true')
      location.assign(to.fullPath)
    }
  } else {
    console.error(err)
  }
})

router.isReady().then(() => {
  localStorage.removeItem('vuetify:dynamic-reload')
})

export default router

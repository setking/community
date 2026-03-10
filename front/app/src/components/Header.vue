<template>
  <v-app-bar class="bar" elevation="4" rounded>
    <template #prepend>
      <v-btn icon>
        <v-img :height="30" src="@/assets/logo.png" :width="160" />
      </v-btn>
    </template>

    <template #append>
      <div class="search-wrapper">
        <v-btn icon="mdi-magnify" @click="expand = !expand" />
        <v-slide-x-reverse-transition>
          <v-text-field
            v-if="expand"
            v-model="query"
            autofocus
            class="search-input"
            clearable
            dense
            hide-details
            label="请输入搜索内容"
            @keyup.enter="handleSearch"
          />
        </v-slide-x-reverse-transition>
      </div>
      <v-btn v-if="!isLogin" variant="text" @click="register">
        注册
      </v-btn>
      <v-btn v-if="!isLogin" variant="text" @click="login">
        登录
      </v-btn>
      <v-menu v-if="isLogin" open-on-hover>
        <template #activator="{ props }">
          <v-btn color="primary" v-bind="props">
            用户
          </v-btn>
        </template>

        <v-list>
          <v-list-item v-for="(item, index) in items" :key="index" :value="index">
            <v-list-item-title @click="checkItem(item)">{{ item.title }}</v-list-item-title>
          </v-list-item>
        </v-list>
      </v-menu>

    </template>
  </v-app-bar>
  <v-dialog
    v-model="dialog"
    max-width="400"
    persistent
  >
    <v-card>
      <v-card-title class="d-flex justify-space-between align-center">
        <div class="text-h5 text-medium-emphasis ps-2">
          {{ isRegister ? '注册' : '登录' }}
        </div>

        <v-btn
          icon="mdi-close"
          variant="text"
          @click="dialog = false"
        />
      </v-card-title>
      <v-card-text>
        <Login :is-register="isRegister" @close-dialog="closeDialog" />
      </v-card-text>
    </v-card>
  </v-dialog>
</template>

<script setup lang="ts">
  import { fa } from 'vuetify/locale'
  import Login from '@/pages/login/index.vue'
  import { getToken, removeToken } from '@/utils/token'
  type Item = {
    title: string
    id: number
  }
  const query = ref('')
  const expand = ref<boolean>(false)
  const items = ref<Item[]>([])
  const dialog = ref(false)
  const isLogin = ref(false)
  const isRegister = ref(false)
  const router = useRouter()
  if (getToken()) {
    isLogin.value = true
  }
  items.value = [
    { title: '发帖', id: 1 },
    { title: '个人中心', id: 2 },
    { title: '退出', id: 3 },
  ]
  function checkItem (item: Item) {
    console.log(item.id)
    if (item.id === 3) {
      logout()
    }
  }
  function handleSearch () {
    if (!query.value) return
    console.log('搜索关键字:', query.value)
  // TODO: 替换成路由跳转或 API 调用
  }
  function closeDialog (data: any) {
    dialog.value = data.msg
    isLogin.value = true
    router.push('/')
  }
  function register () {
    dialog.value = true
    isRegister.value = true
  }
  function login () {
    dialog.value = true
    isRegister.value = false
  }
  function logout () {
    removeToken()
    isLogin.value = false
    router.push('/')
  }
</script>

<style scoped lang="sass">
.bar
  padding-left: 20px
  padding-right: 20px
.search-wrapper
  display: flex;
  align-items: center;
  gap: 4px;
  position: relative;

.search-input
  position: absolute;
  right: 48px; /* 搜索按钮宽度 + 间距 */
  width: 200px;
  max-width: 200px;
  z-index: 10; /* 保证覆盖右侧按钮 */

</style>

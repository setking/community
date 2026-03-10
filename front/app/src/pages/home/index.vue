<template class="overflow-y-auto" height="600">
  <v-infinite-scroll @load="load">
    <v-list lines="three">
      <v-list-item
        v-for="item in items"
        :key="item.id"
        align="start"
        class="post with-divider"
        min-height="160"
        @click="detail(item)"
      >
        <!-- 左侧头像 -->
        <template #prepend>
          <v-avatar size="60">
            <v-img src="https://cdn.vuetifyjs.com/images/lists/1.jpg" />
          </v-avatar>
        </template>

        <!-- 标题 -->
        <v-list-item-title>
          {{ item.Title }}
        </v-list-item-title>

        <!-- 副标题（支持 HTML） -->
        <v-list-item-subtitle>
          <div v-html="item.Content" />
        </v-list-item-subtitle>
      </v-list-item>
    </v-list>

    <template #empty>
      <v-alert type="warning">没有更多的帖子了</v-alert>
    </template>
  </v-infinite-scroll>
</template>

<script lang="ts" setup>
  import { I } from 'vue-router/dist/router-CWoNjPRp.mjs'
  import { id } from 'vuetify/locale'
  import { PostListApi } from '@/http/api/post'
  type Post = {
    prependAvatar: string
    title: string
    subtitle: string
  }
  const router = useRouter()
  const page = ref<{ p: number, n: number }>({ p: 1, n: 10 })
  const items = ref<any>([])
  const loading = ref(false)
  const noMore = ref(false)
  async function GetList () {
    const res = await PostListApi(page.value)
    return res.data.data as Post[]
  }
  async function load ({ done }: { done: (status: 'ok' | 'empty') => void }) {
    if (loading.value || noMore.value) {
      done('empty')
      return
    }
    loading.value = true

    const list = await GetList()

    if (list.length === 0) {
      noMore.value = true
      done('empty')
    } else {
      items.value.push(...list)
      page.value.p += 1
      done('ok')
    }

    loading.value = false
  }
  /** ***********  ✨ Windsurf Command ⭐  *************/
  /**
   * 点击帖子项，跳转到对应的帖子详情页面
   * @param {any} item 帖子的数据
   */
  /** *****  0f7dede5-46c8-4957-87bf-0ef809065d0f  *******/
  function detail (item: any) {
    router.push(`/detail/${item.PostID}`)
  }
</script>
<style lang="scss" scoped>
.post {
  cursor: pointer;
}
.with-divider {
  border-bottom: 1px solid rgba(0, 0, 0, 0.12);
}
</style>

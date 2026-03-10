<template>
  <div class="content">
    <div class="d-flex ga-2">
      <v-chip
        class="ma-2"
        label
        to="home"
      >
        全部
      </v-chip>
      <template v-for="value in communityList" :key="value.id">
        <v-chip
          class="ma-2"
          label
          to="home"
        >
          {{ value.CommunityName }}
        </v-chip>
      </template>
    </div>
  </div>

</template>

<script setup lang="ts">
  import type { Community } from '@/http/api/community/type'
  import { CommunityListApi } from '@/http/api/community'
  const communityList = ref<any>([])
  const page = ref<{ p: number, n: number }>({ p: 1, n: 10 })
  CommunityListApi(page).then(res => {
    communityList.value = res.Data
  })
</script>
<style scoped lang="sass">
</style>

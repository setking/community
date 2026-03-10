<template>
  <Header />
  <v-sheet
    class="sheet overflow-hidden flex-wrap text-center mx-auto px-4"
    elevation="1"
    height="100%"
    max-width="1200"
    rounded
    width="100%"
  >
    <Tags class="tags" />
    <AppMain />
  </v-sheet>
</template>

<script lang="ts" setup>

</script>
<style lang="sass" scoped>
  .text-center
    padding-top: 64px;
    position: relative
    .tags
      max-width: 1180px
      width: 100%
      position: fixed
      z-index: 999
      background-color: #fff
</style>

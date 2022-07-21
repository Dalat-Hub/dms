<!--  This route can be used as a general routing page for parent class routing -->
<!--  To customize the parent class routing page, please refer to @/view/superAdmin/index.vue -->
<template>
  <div>
    <router-view v-slot="{ Component }">
      <transition mode="out-in" name="el-fade-in-linear">
        <keep-alive :include="routerStore.keepAliveRouters">
          <component :is="Component" />
        </keep-alive>
      </transition>
    </router-view>
  </div>
</template>

<script>
export default {
  name: 'RouterHolder'
}
</script>

<script setup>
import { useRouterStore } from '@/pinia/modules/router'
const routerStore = useRouterStore()
</script>

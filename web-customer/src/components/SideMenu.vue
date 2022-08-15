<template>
  <el-card class="box-card card-info" style="margin-bottom: 1rem">
    <template #header>
      <div class="card-header">
        <h2 class="card-title">{{ title }}</h2>
      </div>
    </template>
    <div
      v-for="item in items"
      :key="item.ID"
      class="text item"
      style="margin-bottom: 0.5rem"
    >
      <router-link :to="this.linkTo(item)">
        {{ item.name }}
      </router-link>
    </div>
  </el-card>
</template>

<script>
export default {
  name: "SideMenu",
  props: ["title", "items", "param"],
  methods: {
    linkTo(item) {
      return `/van-ban?phan-loai=${this.param}&id=${item.ID}`;
    },
  },
};
</script>

<style scoped>
.card-info {
  text-align: left;
}
.card-title {
  text-align: left;
  font-size: 1rem;
}
</style>

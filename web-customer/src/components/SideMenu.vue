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
      class="text item dms-sidemenu-item"
      style="margin-bottom: 0.5rem"
    >
      <router-link :to="this.linkTo(item)">
        {{ item.name }}
      </router-link>
      <el-tag v-if="item.count && item.count > 0" class="ml-2" type="success">{{
        item.count
      }}</el-tag>
    </div>
  </el-card>
</template>

<script>
export default {
  name: "SideMenu",
  props: ["title", "items", "param"],
  methods: {
    linkTo(item) {
      return `/van-ban?${this.param}=${item.ID}`;
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
.ml-2 {
  margin-left: 0.3rem;
}
.dms-sidemenu-item:not(:last-child) {
  border-bottom: 1px solid #eee;
  padding-bottom: 7px;
  margin-bottom: 16px !important;
}
</style>

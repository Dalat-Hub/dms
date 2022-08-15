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
        {{ item.shortTitle || item.title }}

        <span v-if="this.displayCreatedAt" class="document-date">
          {{ this.getCreatedDate(item) }}
        </span>

        <el-tag v-if="this.displayCounter" class="ml-2" type="success">{{
          item.viewCount
        }}</el-tag>
      </router-link>
    </div>
  </el-card>
</template>

<script>
import { getDateFormatted } from "@/utils/date";

export default {
  name: "SideContent",
  props: ["title", "items", "displayCounter", "displayCreatedAt"],
  methods: {
    linkTo(item) {
      return `/van-ban/${item.ID}`;
    },
    getCreatedDate(item) {
      return getDateFormatted(item.CreatedAt, "dd-MM-yyyy hh:mm:ss");
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
.document-date {
  font-size: 14px;
  color: gray;
  display: block;
}
</style>

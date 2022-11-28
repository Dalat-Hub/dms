<template>
  <el-card  shadow="never" class="box-card card-info" style="margin-bottom: 1rem">
    <template #header>
      <div class="card-header">
        <h2 class="card-title">{{ title }}</h2>
      </div>
    </template>
    <div class="text item" style="margin-bottom: 0.5rem">
      <el-tree
        :data="tree"
        :props="defaultProps"
        @node-click="(node) => $emit('onNodeClick', node)"
      />
    </div>
  </el-card>
</template>

<script>
export default {
  name: "AgencyTree",
  props: ["title", "tree"],
  emits: ["onNodeClick"],
};
</script>

<script setup>
const defaultProps = {
  children: "children",
  label: "label",
  categoryId: "categoryId",
  agencyId: "agencyId",
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
</style>

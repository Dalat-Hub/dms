<template>
  <el-breadcrumb separator="/" style="margin-bottom: 1rem">
    <el-breadcrumb-item :to="{ path: '/' }">Trang chủ</el-breadcrumb-item>
    <el-breadcrumb-item>Văn bản</el-breadcrumb-item>
  </el-breadcrumb>
</template>

<script>
export default {
  name: "BreadCrumb",
};
</script>

<style scoped></style>

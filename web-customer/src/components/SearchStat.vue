<template>
  <div>
    <p>
      <span v-if="count > 0">Tìm thấy {{ count }} văn bản </span>
      <span v-else>Không tìm thấy văn bản </span>

      <span v-if="Object.keys(searchBy).length === 0">trên hệ thống</span
      ><span v-else>thuộc:</span>
    </p>

    <p v-if="searchBy.signText">
      Từ khóa:
      <el-tag class="ml-2" type="danger">{{ this.searchBy.signText }}</el-tag>
    </p>
    <p v-if="searchBy.agencyId">
      Cơ quan ban hành:
      <el-tag class="ml-2" type="success">{{ this.getAgencyTitle() }}</el-tag>
    </p>
    <p v-if="searchBy.categoryId">
      Thể loại:
      <el-tag class="ml-2" type="info">{{ this.getCategoryTitle() }}</el-tag>
    </p>
    <p v-if="searchBy.fieldId">
      Lĩnh vực:
      <el-tag class="ml-2" type="warning">{{ this.getFieldTitle() }}</el-tag>
    </p>
  </div>
</template>

<script>
export default {
  name: "SearchStat",
  props: ["count", "searchBy", "meta"],
  methods: {
    getAgencyTitle() {
      return (
        this.meta.agency.items.find((f) => f.ID == this.searchBy.agencyId)
          ?.name || "--"
      );
    },
    getCategoryTitle() {
      return (
        this.meta.category.items.find((f) => f.ID == this.searchBy.categoryId)
          ?.name || "--"
      );
    },
    getFieldTitle() {
      return (
        this.meta.field.items.find((f) => f.ID == this.searchBy.fieldId)
          ?.name || "--"
      );
    },
  },
};
</script>

<style scoped></style>

<template>
  <div>
    <p>
      <span v-if="count > 0"
        >Tìm thấy
        <el-tag class="ml-2" type="danger">{{ count }}</el-tag>
        văn bản trên hệ thống
      </span>
      <span v-else>Không tìm thấy văn bản trên hệ thống</span>
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

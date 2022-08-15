<template>
  <div style="margin-bottom: 2rem">
    <el-image src="banner.png" fit="cover" />
    <el-menu
      :default-active="activeIndex"
      class="el-menu-demo"
      mode="horizontal"
      @select="handleSelect"
    >
      <el-menu-item index="1">
        <router-link to="/">Trang chủ</router-link>
      </el-menu-item>
      <el-menu-item index="2">
        <router-link to="/van-ban">Văn bản</router-link>
      </el-menu-item>
      <el-sub-menu index="3">
        <template #title>Cơ quan ban hành</template>
        <el-menu-item
          v-for="(agency, index) in this.agency.items"
          :index="'3-' + (index + 1)"
          :key="agency.ID"
          ><router-link
            :to="`/van-ban?phan-loai=co-quan-ban-hanh&id=${agency.ID}`"
            >{{ agency.name }}</router-link
          ></el-menu-item
        >
      </el-sub-menu>
      <el-sub-menu index="4">
        <template #title>Thể loại</template>
        <el-menu-item
          v-for="(category, index) in this.category.items"
          :index="'4-' + (index + 1)"
          :key="category.ID"
          ><router-link :to="`/van-ban?phan-loai=the-loai&id=${category.ID}`">{{
            category.name
          }}</router-link></el-menu-item
        >
      </el-sub-menu>
      <el-sub-menu index="5">
        <template #title>Lĩnh vực</template>
        <el-menu-item
          v-for="(field, index) in this.field.items"
          :index="'5-' + (index + 1)"
          :key="field.ID"
          ><router-link :to="`/van-ban?phan-loai=linh-vuc&id=${field.ID}`">{{
            field.name
          }}</router-link></el-menu-item
        >
      </el-sub-menu>
      <el-menu-item index="6">Yêu cầu văn bản</el-menu-item>
      <el-menu-item index="7">Giới thiệu</el-menu-item>
      <el-menu-item index="8">Đăng nhâp</el-menu-item>
    </el-menu>
  </div>
</template>

<script>
import { getDocumentCategoryList } from "@/api/category";
import { getDocumentFieldList } from "@/api/field";
import { getDocumentAgencyList } from "@/api/agency";

export default {
  // eslint-disable-next-line vue/multi-word-component-names
  name: "Header",
  data() {
    return {
      agency: {
        title: "Cơ quan ban hành",
        items: [],
      },
      category: {
        title: "Thể loại",
        items: [],
      },
      field: {
        title: "Lĩnh vực",
        items: [],
      },
      activeIndex: "1",
    };
  },
  mounted() {
    this.getAgencies();
    this.getCategories();
    this.getFields();
  },
  methods: {
    async getAgencies() {
      const table = await getDocumentAgencyList({ page: 1, pageSize: 1000 });

      if (table.data.code === 0) {
        this.agency.items = table.data.data.list.map((item) => {
          return {
            ...item,
            link: "/",
          };
        });
      }
    },
    async getCategories() {
      const table = await getDocumentCategoryList({ page: 1, pageSize: 1000 });

      if (table.data.code === 0) {
        this.category.items = table.data.data.list.map((item) => {
          return {
            ...item,
            link: "/",
          };
        });
      }
    },
    async getFields() {
      const table = await getDocumentFieldList({ page: 1, pageSize: 1000 });

      if (table.data.code === 0) {
        this.field.items = table.data.data.list.map((item) => {
          return {
            ...item,
            link: "/",
          };
        });
      }
    },
    handleSelect() {},
  },
};
</script>

<style scoped></style>

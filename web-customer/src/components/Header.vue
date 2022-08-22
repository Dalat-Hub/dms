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
          v-for="(agency, index) in agencies"
          :index="'3-' + (index + 1)"
          :key="agency.ID"
          ><router-link :to="`/van-ban?co-quan-ban-hanh=${agency.ID}`">{{
            agency.name
          }}</router-link></el-menu-item
        >
      </el-sub-menu>
      <el-sub-menu index="4">
        <template #title>Thể loại</template>
        <el-menu-item
          v-for="(category, index) in categories"
          :index="'4-' + (index + 1)"
          :key="category.ID"
          ><router-link :to="`/van-ban?the-loai=${category.ID}`">{{
            category.name
          }}</router-link></el-menu-item
        >
      </el-sub-menu>
      <el-menu-item index="5">Yêu cầu văn bản</el-menu-item>
      <el-menu-item index="6">Giới thiệu</el-menu-item>
      <el-menu-item index="7" v-if="!userStore.userInfo.nickName">
        <router-link to="/dang-nhap"> Đăng nhập </router-link>
      </el-menu-item>

      <el-sub-menu v-else index="8">
        <template #title>Xin chào {{ userStore.userInfo.nickName }}</template>
        <el-menu-item index="8-1">
          <span>Quản trị</span>
        </el-menu-item>
        <el-menu-item index="8-2">
          <router-link to="/">Tài khoản</router-link>
        </el-menu-item>
        <el-menu-item index="8-3">
          <router-link to="/">Đăng xuất</router-link>
        </el-menu-item>
      </el-sub-menu>
    </el-menu>
  </div>
</template>

<script>
export default {
  // eslint-disable-next-line vue/multi-word-component-names
  name: "Header",
};
</script>

<script setup>
import { getDocumentCategoryList } from "@/api/category";
import { getDocumentAgencyList } from "@/api/agency";
import { useUserStore } from "@/pinia/modules/user";

import { ref } from "vue";

const userStore = useUserStore();

const agencies = ref([]);
const categories = ref([]);

const activeIndex = ref("1");

const getUserInfo = () => {
  userStore.GetUserInfo();
};

const openAdminPage = () => {
  window.open(
    `${process.env.VUE_APP_ADMIN_URL}/#/login?token=${userStore.token}`,
    "_blank"
  );
};

const getAgencies = async () => {
  const table = await getDocumentAgencyList({ page: 1, pageSize: 1000 });

  if (table.code === 0) {
    agencies.value = table.data.list.map((item) => {
      return {
        ...item,
        link: "/",
      };
    });
  }
};

const getCategories = async () => {
  const table = await getDocumentCategoryList({ page: 1, pageSize: 1000 });

  if (table.code === 0) {
    categories.value = table.data.list.map((item) => {
      return {
        ...item,
        link: "/",
      };
    });
  }
};

const handleSelect = (key, keyPath) => {
  if (key === "8-1") openAdminPage();
};

getAgencies();
getCategories();
getUserInfo();
</script>

<style scoped></style>

<template>
  <div class="header-contailer">
    <el-image src="banner.png" fit="cover" />
    <el-menu :default-active="activeIndex" class="el-menu-demo dlu-menu" background-color="#7AA227" text-color="#fff"
      active-text-color="#FC7616" mode="horizontal" @select="handleSelect">
      <el-menu-item index="1">
        <router-link class="router-title" to="/">Trang chủ</router-link>
      </el-menu-item>
      <el-menu-item index="2">
        <router-link class="router-title" to="/van-ban">Văn bản</router-link>
      </el-menu-item>
      <el-sub-menu index="3" class="router-title">
        <template  #title> Cơ quan ban hành </template>
        <el-menu-item class="router-title" text-color="#fff" v-for="(agency, index) in agencies" :index="'3-' + (index + 1)" :key="agency.ID">
          <router-link text-color="#fff"   :to="`/van-ban?co-quan-ban-hanh=${agency.ID}`">{{
              agency.name
          }}</router-link>
        </el-menu-item>
      </el-sub-menu>
      <el-sub-menu index="4" class="router-title">
        <template #title> Thể loại </template>
        <el-menu-item class="router-title" v-for="(category, index) in categories" :index="'4-' + (index + 1)"
          :key="category.ID">
          <router-link class="router-title" :to="`/van-ban?the-loai=${category.ID}`">{{
              category.name
          }}
          </router-link>
        </el-menu-item>
      </el-sub-menu>
      <!-- <el-menu-item index="5">Yêu cầu văn bản</el-menu-item>
      <el-menu-item index="6">Giới thiệu</el-menu-item> -->
      <el-menu-item class="menu-item" index="7" v-if="!userStore.userInfo.nickName">
        <router-link class="router-title" to="/dang-nhap">Đăng nhập</router-link>
      </el-menu-item>
      <el-sub-menu v-else index="8" class="router-title">
        <template #title>Xin chào {{ userStore.userInfo.nickName }}</template>
        <el-menu-item class="router-title" index="8-1">
          Quản trị
        </el-menu-item>
        <el-menu-item index="8-2">
          <router-link class="router-title" to="/">Đăng xuất </router-link>
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
import { useRoute } from "vue-router";

import { ref } from "vue";

const userStore = useUserStore();
const route = useRoute();

const agencies = ref([]);
const categories = ref([]);

const activeIndex = ref("1");

const getUserInfo = () => {
  if (route.path.includes("set-token")) return;

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

<style scoped>
.router-title {
  text-transform: uppercase;
  font-weight: 550;
  color: #fff;
}

a[class^="router-link"], .text.item a {
    text-decoration: none;
    color: #fff !important;
}

/* .sub-menu{
  background-color: #f1f1f3 !important;
} */
/* .header-contailer{
  margin-bottom: 2rem;
  
} */
/* .dlu-menu{
background-color: #54880A;
} */
/* .el-sub-menu .is-opened{
  background-color: #54880A;
} */
</style>

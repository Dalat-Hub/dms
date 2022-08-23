<template>
  <div style="margin: 15rem 0">
    <h3 v-if="loginSuccess === null">Tiến hành đăng nhập...</h3>

    <div v-if="loginSuccess === false">
      <h1>Có lỗi trong quá trình đăng nhập</h1>
    </div>

    <div v-if="loginSuccess === true">
      <h1>Đăng nhập thành công</h1>

      <div style="margin-top: 1rem; margin-bottom: 1rem">
        <el-button type="primary" plain @click="handleOnHomePageClick"
          >Trang chủ</el-button
        >
      </div>

      <div>
        <el-button type="success" plain @click="handleOnAdminPageClick"
          >Trang quản trị</el-button
        >
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "SetToken",
};
</script>

<script setup>
import { useRoute } from "vue-router";
import { useUserStore } from "@/pinia/modules/user";
import { useRouter } from "vue-router";
import { onMounted, ref } from "vue";

const route = useRoute();
const userStore = useUserStore();
const router = useRouter();

const loginSuccess = ref(null);

const checkLoginToken = async () => {
  const existingToken = userStore.token || null;
  if (existingToken) {
    loginSuccess.value = true;
    return;
  }

  const token = route.query.token || null;
  if (!token) {
    loginSuccess.value = false;
    return;
  }

  userStore.setToken(route.query.token);
  await userStore.GetUserInfo();

  loginSuccess.value = true;
};

const handleOnHomePageClick = () => {
  router.push({ name: "home" });
};

const handleOnAdminPageClick = () => {
  window.open(
    `${process.env.VUE_APP_ADMIN_URL}/#/login?token=${userStore.token}`,
    "_blank"
  );
};

onMounted(() => {
  checkLoginToken();
});
</script>

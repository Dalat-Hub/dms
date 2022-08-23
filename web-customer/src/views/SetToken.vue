<template>
  <h1>Tiến hành đăng nhập</h1>

  <div>
    <h1>Đăng nhập thành công, nhấn vào đây để quay về trang chủ</h1>
  </div>
</template>

<script setup>
import { useRoute } from "vue-router";
import { useUserStore } from "@/pinia/modules/user";
import { useRouter } from "vue-router";
import { onMounted } from "vue";

const route = useRoute();
const userStore = useUserStore();
const router = useRouter();

const checkLoginToken = async () => {
  const existingToken = userStore.token || null;
  if (existingToken) {
    router.push({ name: "home" });
    return;
  }

  const token = route.query.token || null;
  if (!token) {
    router.push({ name: "dms-login" });
    return;
  }

  userStore.setToken(route.query.token);
  await userStore.GetUserInfo();

  router.push({ name: "home" });
};

onMounted(() => {
  checkLoginToken();
});
</script>

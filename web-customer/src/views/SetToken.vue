<template>
  <div style="margin: 15rem 0">
    <h3>Tiến hành đăng nhập...</h3>
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
    router.push({ name: "home" });
    return;
  }

  userStore.setToken(route.query.token);
  await userStore.GetUserInfo();
};

onMounted(() => {
  checkLoginToken();
});
</script>

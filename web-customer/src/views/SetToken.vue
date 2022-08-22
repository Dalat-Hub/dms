<template>
  <h1>Hello world</h1>
</template>

<script setup>
import { useRoute } from "vue-router";
import { useUserStore } from "@/pinia/modules/user";
import { useRouter } from "vue-router";

const route = useRoute();
const userStore = useUserStore();
const router = useRouter();

const checkLoginToken = async () => {
  const query = route.query;
  let hasValue = false;

  if (query.redirect) {
    if (!query.redirect.includes("token=")) return;

    const parts = query.redirect.split("token=");
    const token = parts[1];

    if (!token) return;

    userStore.setToken(token);
    hasValue = true;
  }

  if (query.token) {
    userStore.setToken(query.token);
    hasValue = true;
  }

  if (hasValue) {
    const res = await userStore.GetUserInfo();

    if (res.code === 0) {
      router.push({ path: "/" });
    } else {
      router.push({ path: "/?error=1" });
    }
  }
};
checkLoginToken();
</script>

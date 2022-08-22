<template>
  <h1>Processing login ...</h1>
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

  const existingToken = window.localStorage.getItem("token") || null;
  if (existingToken) {
    router.push({
      name: "home",
    });

    return;
  }

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
    await userStore.GetUserInfo();
    window.location.reload();
  }
};
checkLoginToken();
</script>

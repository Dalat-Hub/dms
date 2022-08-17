import { login, getUserInfo, setSelfInfo } from "@/api/user";
import { ElLoading, ElMessage } from "element-plus";
import { defineStore } from "pinia";
import { ref, computed, watch } from "vue";

export const useUserStore = defineStore("user", () => {
  const loadingInstance = ref(null);

  const userInfo = ref({
    uuid: "",
    nickName: "",
    headerImg: "",
    authority: {},
    sideMode: "dark",
    activeColor: "#4D70FF",
    baseColor: "#fff",
  });
  const token = ref(window.localStorage.getItem("token") || "");
  const setUserInfo = (val) => {
    userInfo.value = val;
  };

  const setToken = (val) => {
    token.value = val;
  };

  const ResetUserInfo = (value = {}) => {
    userInfo.value = {
      ...userInfo.value,
      ...value,
    };
  };

  const GetUserInfo = async () => {
    const res = await getUserInfo();
    if (res.code === 0) {
      setUserInfo(res.data.userInfo);
    }
    return res;
  };

  const LoginIn = async (loginInfo) => {
    loadingInstance.value = ElLoading.service({
      fullscreen: true,
      text: "Đang đăng nhập hệ thống ...",
    });
    try {
      const res = await login(loginInfo);
      if (res.code === 0) {
        setUserInfo(res.data.user);
        setToken(res.data.token);

        loadingInstance.value.close();
        return true;
      }
    } catch (e) {
      loadingInstance.value.close();
    }
    loadingInstance.value.close();
  };

  const LoginOut = async () => {
    token.value = "";
    sessionStorage.clear();
    localStorage.clear();
    window.location.reload();
  };

  const ClearStorage = async () => {
    token.value = "";
    sessionStorage.clear();
    localStorage.clear();
  };

  const changeSideMode = async (data) => {
    const res = await setSelfInfo({ sideMode: data });
    if (res.code === 0) {
      userInfo.value.sideMode = data;
      ElMessage({
        type: "success",
        message: "Cập nhật thành công",
      });
    }
  };

  const mode = computed(() => userInfo.value.sideMode);
  const sideMode = computed(() => {
    if (userInfo.value.sideMode === "dark") {
      return "#191a23";
    } else if (userInfo.value.sideMode === "light") {
      return "#fff";
    } else {
      return userInfo.value.sideMode;
    }
  });
  const baseColor = computed(() => {
    if (userInfo.value.sideMode === "dark") {
      return "#fff";
    } else if (userInfo.value.sideMode === "light") {
      return "#191a23";
    } else {
      return userInfo.value.baseColor;
    }
  });
  const activeColor = computed(() => {
    if (
      userInfo.value.sideMode === "dark" ||
      userInfo.value.sideMode === "light"
    ) {
      return "#4D70FF";
    }
    return userInfo.value.activeColor;
  });

  watch(
    () => token.value,
    () => {
      window.localStorage.setItem("token", token.value);
    }
  );

  return {
    userInfo,
    token,
    ResetUserInfo,
    GetUserInfo,
    LoginIn,
    LoginOut,
    changeSideMode,
    mode,
    sideMode,
    setToken,
    baseColor,
    activeColor,
    loadingInstance,
    ClearStorage,
  };
});

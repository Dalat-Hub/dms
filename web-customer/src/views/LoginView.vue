<template>
  <div class="heading-intro">
    <div class="dlu-logo">
    <el-image style="width: 100px; height: 100px" src="dlu-logo.webp"/>
  </div>
    <h2>TRƯỜNG ĐẠI HỌC ĐÀ LẠT</h2>
    <h3>HỆ THỐNG QUẢN LÝ TÀI LIỆU</h3>
  </div>

  <el-row justify="center" class="login-form">
    <el-col :sm="12" :lg="10">
      <el-form :model="form" label-width="150px" @keyup.enter="onSubmit" label-position="top">
        <el-form-item label="Tên đăng nhập">
          <el-input v-model="form.username" placeholder="Vui lòng điền tên đăng nhập" autocomplete="off" />
        </el-form-item>
        <el-form-item label="Mật khẩu">
          <el-input v-model="form.password" placeholder="Vui lòng điền mật khẩu" autocomplete="off" type="password" />
        </el-form-item>
        <el-form-item label="Mã xác thực">
          <el-input v-model="form.captcha" placeholder="Vui lòng điền mã xác thực" style="width: 60%" />
          <img v-if="picPath" :src="picPath" alt="Vui lòng nhập mã bảo mật" @click="loginVerify()" style="width: 40%" />
        </el-form-item>
        <el-form-item>
          <el-col class="center">
            <el-button lager type="success" @click="onSubmit">Đăng nhập</el-button>
          </el-col>
        </el-form-item>
        <el-divider>Hoặc</el-divider>
        <el-form-item>
          <el-col class="center">
            <a :href="getGoogleUrl('/')" style="
            display: flex;
            align-items: center;
            margin-top: 1rem;
            text-decoration: none;
          ">
              <img src="@/assets/google.png" alt="Google" class="link-icon" style="width: 2rem; margin-right: 0.5rem" />
              <span style="color: #333">Đăng nhập với Google</span>
            </a>
          </el-col>
        </el-form-item>
      </el-form>
      <el-divider></el-divider>
    </el-col>
  </el-row>
</template>

<script>
export default {
  name: "LoginView",
};
</script>

<script setup>
import { reactive, ref } from "vue";
import { captcha } from "@/api/user";
import { useUserStore } from "@/pinia/modules/user";
import router from "@/router/index";
import { getGoogleUrl } from "@/utils/getGoogleUrl";

const userStore = useUserStore();

const picPath = ref("");
const form = reactive({
  username: "",
  password: "",
  captcha: "",
  captchaId: "",
});

const login = async () => {
  return await userStore.LoginIn(form);
};

const loginVerify = () => {
  captcha({}).then((ele) => {
    picPath.value = ele.data.picPath;
    form.captchaId = ele.data.captchaId;
  });
};
loginVerify();

const onSubmit = async () => {
  const flag = await login();
  if (!flag) {
    loginVerify();
  } else {
    userStore.GetUserInfo().then(() => {
      router.push({
        path: "/",
      });
    });
  }
};
</script>

<style scoped>
.login-form {
  padding-top: 3rem;
  padding-bottom: 20rem;
}

.heading-intro {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-content: center;
}

.center {
  display: flex;
  justify-content: center;
}
.dlu-logo{
  width: 100%;
  margin: 12px;
}
</style>

<template>
  <div id="userLayout">
    <div class="login_panel">
      <div class="login_panel_form">
        <div class="login_panel_form_title">
          <img
            class="login_panel_form_title_logo"
            :src="$GIN_VUE_ADMIN.appLogo"
            alt
          >
          <p class="login_panel_form_title_p">{{ $GIN_VUE_ADMIN.appName }}</p>
        </div>
        <el-form
          ref="loginForm"
          :model="loginFormData"
          @keyup.enter="submitForm"
        >
          <el-form-item prop="username">
            <el-input
              v-model="loginFormData.username"
              placeholder="Điền tên đăng nhập của bạn"
            >
              <template #suffix>
                <span class="input-icon">
                  <el-icon>
                    <user />
                  </el-icon>
                </span>
              </template>
            </el-input>
          </el-form-item>
          <el-form-item prop="password">
            <el-input
              v-model="loginFormData.password"
              :type="lock === 'lock' ? 'password' : 'text'"
              placeholder="Điền mật khẩu của bạn"
            >
              <template #suffix>
                <span class="input-icon">
                  <el-icon>
                    <component
                      :is="lock"
                      @click="changeLock"
                    />
                  </el-icon>
                </span>
              </template>
            </el-input>
          </el-form-item>
          <el-form-item prop="captcha">
            <div class="vPicBox">
              <el-input
                v-model="loginFormData.captcha"
                placeholder="Vui lòng nhập mã bảo mật"
                style="width: 60%"
              />
              <div class="vPic">
                <img
                  v-if="picPath"
                  :src="picPath"
                  alt="Vui lòng nhập mã bảo mật"
                  @click="loginVerify()"
                >
              </div>
            </div>
          </el-form-item>
          <el-form-item>
            <el-button
              v-if="notYetInit"
              type="primary"
              style="width: 46%; margin-right: 8%;"
              size="large"
              @click="checkInit"
            >Khởi tạo hệ thống</el-button>
            <el-button
              type="primary"
              size="large"
              style="width: 46%;"
              @click="submitForm"
            >Đăng nhập</el-button>
          </el-form-item>
        </el-form>

        <hr>

        <div>
          <a :href="getGoogleUrl('/')" style="display: flex; align-items: center; margin-top: 1rem;">
            <img src="@/assets/google.png" alt="Google" class="link-icon" style="width: 2rem; margin-right: 0.5rem;">
            <span style="color: #333;">Đăng nhập với Google</span>
          </a>
        </div>

      </div>
      <div class="login_panel_right" />
      <div class="login_panel_foot">
        <div class="links">
          <a href="http://doc.henrongyi.top/" target="_blank">
            <img src="@/assets/docs.png" class="link-icon">
          </a>
          <a href="https://support.qq.com/product/371961" target="_blank">
            <img src="@/assets/kefu.png" class="link-icon">
          </a>
          <a
            href="https://github.com/flipped-aurora/gin-vue-admin"
            target="_blank"
          >
            <img src="@/assets/github.png" class="link-icon">
          </a>
          <a href="https://space.bilibili.com/322210472" target="_blank">
            <img src="@/assets/video.png" class="link-icon">
          </a>
        </div>
        <div class="copyright">
          <BottomInfo />
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Login',
}
</script>

<script setup>
import { captcha } from '@/api/user'
import { checkDB } from '@/api/initdb'
import BottomInfo from '@/view/layout/bottomInfo/bottomInfo.vue'
import { reactive, ref } from 'vue'
import { ElMessage } from 'element-plus'
import { useRouter } from 'vue-router'
import { useRoute } from 'vue-router'
import { useUserStore } from '@/pinia/modules/user'
import { getGoogleUrl } from '@/utils/getGoogleUrl'

const userStore = useUserStore()
const router = useRouter()
const route = useRoute()

const notYetInit = ref(true)

const checkLoginToken = async() => {
  const query = route.query

  if (!query.redirect) return
  if (!query.redirect.includes('token=')) return

  const parts = query.redirect.split('token=')
  const token = parts[1]

  if (!token) return

  userStore.setToken(token)
  const res = await userStore.GetUserInfo()

  if (res.code === 0) {
    window.location.reload()
  } else {
    ElMessage({
      type: 'error',
      message: 'Thông tin đăng nhập không hợp lệ',
      showClose: true,
      duration: 5000
    })
  }
}
checkLoginToken()

const checkUsername = (rule, value, callback) => {
  if (value.length < 5) {
    return callback(new Error('Vui lòng điền tên đăng nhập'))
  } else {
    callback()
  }
}
const checkPassword = (rule, value, callback) => {
  if (value.length < 6) {
    return callback(new Error('Vui lòng điền mật khẩu'))
  } else {
    callback()
  }
}

const loginVerify = () => {
  captcha({}).then((ele) => {
    rules.captcha[1].max = ele.data.captchaLength
    rules.captcha[1].min = ele.data.captchaLength
    picPath.value = ele.data.picPath
    loginFormData.captchaId = ele.data.captchaId
  })
}
loginVerify()

const lock = ref('lock')
const changeLock = () => {
  lock.value = lock.value === 'lock' ? 'unlock' : 'lock'
}

const loginForm = ref(null)
const picPath = ref('')
const loginFormData = reactive({
  username: '',
  password: '',
  captcha: '',
  captchaId: '',
})
const rules = reactive({
  username: [{ validator: checkUsername, trigger: 'blur' }],
  password: [{ validator: checkPassword, trigger: 'blur' }],
  captcha: [
    { required: true, message: 'Vui lòng nhập mã bảo mật', trigger: 'blur' },
    {
      message: 'Mã xác thực không hợp lệ',
      trigger: 'blur',
    },
  ],
})

const login = async() => {
  return await userStore.LoginIn(loginFormData)
}
const submitForm = () => {
  loginForm.value.validate(async(v) => {
    if (v) {
      const flag = await login()
      if (!flag) {
        loginVerify()
      }
    } else {
      ElMessage({
        type: 'error',
        message: 'Vui lòng điền đủ thông tin đăng nhập',
        showClose: true,
      })
      loginVerify()
      return false
    }
  })
}

const checkInit = async() => {
  const res = await checkDB()
  if (res.code === 0) {
    if (res.data?.needInit) {
      userStore.NeedInit()
      router.push({ name: 'Init' })
    } else {
      ElMessage({
        type: 'info',
        message: 'Hệ thống đã được khởi tạo',
      })
    }
  }
}

const checkSystemInit = async() => {
  const res = await checkDB()

  if (res.code === 0) {
    if (!res.data?.needInit) {
      notYetInit.value = false
    }
  }
}
checkSystemInit()

</script>

<style lang="scss" scoped>
@import "@/style/newLogin.scss";
</style>

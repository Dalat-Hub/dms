import axios from 'axios' // 引入axios
import { ElMessage, ElMessageBox } from 'element-plus'
import { useUserStore } from '@/pinia/modules/user'
import { emitter } from '@/utils/bus.js'
import router from '@/router/index'

const service = axios.create({
  baseURL: import.meta.env.VITE_BASE_API,
  timeout: 99999
})
let acitveAxios = 0
let timer
const showLoading = () => {
  acitveAxios++
  if (timer) {
    clearTimeout(timer)
  }
  timer = setTimeout(() => {
    if (acitveAxios > 0) {
      emitter.emit('showLoading')
    }
  }, 400)
}

const closeLoading = () => {
  acitveAxios--
  if (acitveAxios <= 0) {
    clearTimeout(timer)
    emitter.emit('closeLoading')
  }
}

// http request interceptor
service.interceptors.request.use(
  config => {
    if (!config.donNotShowLoading) {
      showLoading()
    }
    const userStore = useUserStore()
    config.headers = {
      'Content-Type': 'application/json',
      'x-token': userStore.token,
      'x-user-id': userStore.userInfo.ID,
      ...config.headers
    }
    return config
  },
  error => {
    closeLoading()
    ElMessage({
      showClose: true,
      message: error,
      type: 'error'
    })
    return error
  }
)

// http response interceptor
service.interceptors.response.use(
  response => {
    const userStore = useUserStore()
    closeLoading()
    if (response.headers['new-token']) {
      userStore.setToken(response.headers['new-token'])
    }
    if (response.data.code === 0 || response.headers.success === 'true') {
      if (response.headers.msg) {
        response.data.msg = decodeURI(response.headers.msg)
      }
      return response.data
    } else {
      ElMessage({
        showClose: true,
        message: response.data.msg || decodeURI(response.headers.msg),
        type: 'error'
      })
      if (response.data.data && response.data.data.reload) {
        userStore.token = ''
        localStorage.clear()
        router.push({ name: 'Login', replace: true })
      }
      return response.data.msg ? response.data : response
    }
  },
  error => {
    closeLoading()

    if (!error.response) {
      ElMessageBox.confirm(`
        <p>Có lỗi xảy ra</p>
        <p>${error}</p>
        `, 'Lỗi', {
        dangerouslyUseHTMLString: true,
        distinguishCancelAndClose: true,
        confirmButtonText: 'Thử lại',
        cancelButtonText: 'Huỷ'
      })
      return
    }

    switch (error.response.status) {
      case 500:
        ElMessageBox.confirm(`
        <p>Lỗi nghiêm trọng: ${error}</p>
        <p>Mã lỗi: <span style="color:red"> 500 </span>：Thử thoát ra và đăng nhập lại hoặc liên hệ quản trị viên</p>
        `, 'Lỗi nghiêm trọng', {
          dangerouslyUseHTMLString: true,
          distinguishCancelAndClose: true,
          confirmButtonText: 'Xoá vùng nhớ tạm thời',
          cancelButtonText: 'Huỷ'
        })
          .then(() => {
            const userStore = useUserStore()
            userStore.token = ''
            localStorage.clear()
            router.push({ name: 'Login', replace: true })
          })
        break
      case 404:
        ElMessageBox.confirm(`
          <p>Lỗi: ${error}</p>
          <p>Mã lỗi: <span style="color:red"> 404 </span>：Không tìm thấy trang này</p>
          `, 'Lỗi', {
          dangerouslyUseHTMLString: true,
          distinguishCancelAndClose: true,
          confirmButtonText: 'Đã hiểu',
          cancelButtonText: 'Đóng'
        })
        break
    }

    return error
  }
)
export default service

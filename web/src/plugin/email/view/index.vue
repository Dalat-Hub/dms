<template>
  <div>
    <warning-bar title="The email configuration file needs to be configured in advance. In order to prevent unnecessary spam, the online experience function does not open this function experience." />
    <div class="gva-form-box">
      <el-form ref="emailForm" label-position="right" label-width="80px" :model="form">
        <el-form-item label="Đến">
          <el-input v-model="form.to" />
        </el-form-item>
        <el-form-item label="Tiêu đề">
          <el-input v-model="form.subject" />
        </el-form-item>
        <el-form-item label="Nội dung">
          <el-input v-model="form.body" type="textarea" />
        </el-form-item>
        <el-form-item>
          <el-button @click="sendTestEmail">Gửi email mẫu</el-button>
          <el-button @click="sendEmail">Gửi email</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>

</template>

<script>
export default {
  name: 'Email',
}
</script>

<script setup>
import WarningBar from '@/components/warningBar/warningBar.vue'
import { emailTest } from '@/plugin/email/api/email.js'
import { ElMessage } from 'element-plus'
import { reactive, ref } from 'vue'
const emailForm = ref(null)
const form = reactive({
  to: '',
  subject: '',
  body: '',
})
const sendTestEmail = async() => {
  const res = await emailTest()
  if (res.code === 0) {
    ElMessage.success('Gửi thành công')
  }
}

const sendEmail = async() => {
  const res = await emailTest()
  if (res.code === 0) {
    ElMessage.success('Gửi thành công')
  }
}
</script>

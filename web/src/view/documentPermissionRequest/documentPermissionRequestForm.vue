<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="acceptAt字段:" prop="acceptAt">
          <el-date-picker v-model="formData.acceptAt" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
        </el-form-item>
        <el-form-item label="acceptUserId字段:" prop="acceptUserId">
          <el-input v-model.number="formData.acceptUserId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="documentId字段:" prop="documentId">
          <el-input v-model.number="formData.documentId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="requestPermission字段:" prop="requestPermission">
          <el-input v-model="formData.requestPermission" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="requestUserId字段:" prop="requestUserId">
          <el-input v-model.number="formData.requestUserId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item>
          <el-button size="mini" type="primary" @click="save">保存</el-button>
          <el-button size="mini" type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
export default {
  name: 'DocumentPermissionRequest'
}
</script>

<script setup>
import {
  createDocumentPermissionRequest,
  updateDocumentPermissionRequest,
  findDocumentPermissionRequest
} from '@/api/documentPermissionRequest'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            acceptAt: new Date(),
            acceptUserId: 0,
            documentId: 0,
            requestPermission: '',
            requestUserId: 0,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findDocumentPermissionRequest({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.redocumentPermissionRequest
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createDocumentPermissionRequest(formData.value)
               break
             case 'update':
               res = await updateDocumentPermissionRequest(formData.value)
               break
             default:
               res = await createDocumentPermissionRequest(formData.value)
               break
           }
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>

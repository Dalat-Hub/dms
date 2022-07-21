<template>
  <div>
    <div class="gva-form-box">
      <el-form ref="elFormRef" :model="formData" label-position="right" :rules="rule" label-width="120px">
        <el-form-item label="documentId字段:" prop="documentId">
          <el-input v-model.number="formData.documentId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="userId字段:" prop="userId">
          <el-input v-model.number="formData.userId" :clearable="true" placeholder="请输入" />
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
  name: 'DocumentSignerReferences'
}
</script>

<script setup>
import {
  createDocumentSignerReferences,
  updateDocumentSignerReferences,
  findDocumentSignerReferences
} from '@/api/documentSignerReferences'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
  documentId: 0,
  userId: 0,
})
// Validation rules
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async() => {
  // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
  if (route.query.id) {
    const res = await findDocumentSignerReferences({ ID: route.query.id })
    if (res.code === 0) {
      formData.value = res.data.redocumentSignerReferences
      type.value = 'update'
    }
  } else {
    type.value = 'create'
  }
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate(async(valid) => {
        if (!valid) return
        let res
        switch (type.value) {
          case 'create':
            res = await createDocumentSignerReferences(formData.value)
            break
          case 'update':
            res = await updateDocumentSignerReferences(formData.value)
            break
          default:
            res = await createDocumentSignerReferences(formData.value)
            break
        }
        if (res.code === 0) {
          ElMessage({
            type: 'success',
            message: 'Thực hiện thao tác thành công'
          })
        }
      })
}

// Back button
const back = () => {
  router.go(-1)
}

</script>

<style>
</style>

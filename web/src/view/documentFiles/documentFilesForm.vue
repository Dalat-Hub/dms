<template>
  <div>
    <div class="gva-form-box">
      <el-form ref="elFormRef" :model="formData" label-position="right" :rules="rule" label-width="120px">
        <el-form-item label="name字段:" prop="name">
          <el-input v-model="formData.name" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="key字段:" prop="key">
          <el-input v-model="formData.key" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="url字段:" prop="url">
          <el-input v-model="formData.url" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="tag字段:" prop="tag">
          <el-input v-model="formData.tag" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="size字段:" prop="size">
          <el-input-number v-model="formData.size" :precision="2" :clearable="true" />
        </el-form-item>
        <el-form-item label="order字段:" prop="order">
          <el-input v-model.number="formData.order" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="documentId字段:" prop="documentId">
          <el-input v-model.number="formData.documentId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="path字段:" prop="path">
          <el-input v-model="formData.path" :clearable="true" placeholder="请输入" />
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
  name: 'DocumentFiles'
}
</script>

<script setup>
import {
  createDocumentFiles,
  updateDocumentFiles,
  findDocumentFiles
} from '@/api/documentFiles'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
  name: '',
  key: '',
  url: '',
  tag: '',
  size: 0,
  order: 0,
  documentId: 0,
  path: '',
})
// Validation rules
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async() => {
  // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
  if (route.query.id) {
    const res = await findDocumentFiles({ ID: route.query.id })
    if (res.code === 0) {
      formData.value = res.data.redocumentFiles
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
            res = await createDocumentFiles(formData.value)
            break
          case 'update':
            res = await updateDocumentFiles(formData.value)
            break
          default:
            res = await createDocumentFiles(formData.value)
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

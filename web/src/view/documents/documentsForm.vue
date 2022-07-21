<template>
  <div>
    <div class="gva-form-box">
      <el-form ref="elFormRef" :model="formData" label-position="right" :rules="rule" label-width="120px">
        <el-form-item label="title字段:" prop="title">
          <el-input v-model="formData.title" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="expert字段:" prop="expert">
          <el-input v-model="formData.expert" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="content字段:" prop="content">
          <el-input v-model="formData.content" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="dateIssued字段:" prop="dateIssued">
          <el-date-picker v-model="formData.dateIssued" type="date" placeholder="选择日期" :clearable="true" />
        </el-form-item>
        <el-form-item label="stillInEffect字段:" prop="stillInEffect">
          <el-switch v-model="formData.stillInEffect" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable />
        </el-form-item>
        <el-form-item label="effectDate字段:" prop="effectDate">
          <el-date-picker v-model="formData.effectDate" type="date" placeholder="选择日期" :clearable="true" />
        </el-form-item>
        <el-form-item label="expirationDate字段:" prop="expirationDate">
          <el-date-picker v-model="formData.expirationDate" type="date" placeholder="选择日期" :clearable="true" />
        </el-form-item>
        <el-form-item label="signNumber字段:" prop="signNumber">
          <el-input v-model.number="formData.signNumber" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="signYear字段:" prop="signYear">
          <el-input v-model.number="formData.signYear" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="signCategory字段:" prop="signCategory">
          <el-input v-model="formData.signCategory" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="signAgency字段:" prop="signAgency">
          <el-input v-model="formData.signAgency" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="signText字段:" prop="signText">
          <el-input v-model="formData.signText" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="categoryId字段:" prop="categoryId">
          <el-input v-model.number="formData.categoryId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="agencyId字段:" prop="agencyId">
          <el-input v-model.number="formData.agencyId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="createdBy字段:" prop="createdBy">
          <el-input v-model.number="formData.createdBy" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="beResponsibleBy字段:" prop="beResponsibleBy">
          <el-input v-model.number="formData.beResponsibleBy" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="viewCount字段:" prop="viewCount">
          <el-input v-model.number="formData.viewCount" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="downloadCount字段:" prop="downloadCount">
          <el-input v-model.number="formData.downloadCount" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="status字段:" prop="status">
          <el-input v-model="formData.status" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="type字段:" prop="type">
          <el-input v-model="formData.type" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="priorityId字段:" prop="priorityId">
          <el-input v-model.number="formData.priorityId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="parentId字段:" prop="parentId">
          <el-input v-model.number="formData.parentId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="currentId字段:" prop="currentId">
          <el-input v-model.number="formData.currentId" :clearable="true" placeholder="请输入" />
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
  name: 'Documents'
}
</script>

<script setup>
import {
  createDocuments,
  updateDocuments,
  findDocuments
} from '@/api/documents'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
  title: '',
  expert: '',
  content: '',
  dateIssued: new Date(),
  stillInEffect: false,
  effectDate: new Date(),
  expirationDate: new Date(),
  signNumber: 0,
  signYear: 0,
  signCategory: '',
  signAgency: '',
  signText: '',
  categoryId: 0,
  agencyId: 0,
  createdBy: 0,
  beResponsibleBy: 0,
  viewCount: 0,
  downloadCount: 0,
  status: '',
  type: '',
  priorityId: 0,
  parentId: 0,
  currentId: 0,
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
    const res = await findDocuments({ ID: route.query.id })
    if (res.code === 0) {
      formData.value = res.data.redocuments
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
            res = await createDocuments(formData.value)
            break
          case 'update':
            res = await updateDocuments(formData.value)
            break
          default:
            res = await createDocuments(formData.value)
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

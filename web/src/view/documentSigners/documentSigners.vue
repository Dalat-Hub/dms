<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item>
          <el-button size="small" type="primary" icon="search" @click="onSubmit">Tìm kiếm</el-button>
          <el-button size="small" icon="refresh" @click="onReset">Đặt lại bộ lọc</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button size="small" type="primary" icon="plus" @click="openDialog">Thêm mới</el-button>
        <el-popover v-model:visible="deleteVisible" placement="top" width="160">
          <p>Bạn có chắc muốn xoá những người này?</p>
          <div style="text-align: right; margin-top: 8px;">
            <el-button size="small" type="primary" link @click="deleteVisible = false">Đóng</el-button>
            <el-button size="small" type="primary" @click="onDelete">Xoá</el-button>
          </div>
          <template #reference>
            <el-button icon="delete" size="small" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="deleteVisible = true">Xoá</el-button>
          </template>
        </el-popover>
      </div>
      <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="Phòng ban" prop="agencyId" width="200">
          <template #default="scope">
            {{ scope.row.agency.name }}
          </template>
        </el-table-column>
        <el-table-column align="left" label="Chức danh" prop="title" width="200">
          <template #default="scope">
            {{ scope.row.signerTitle.ID !== 0 && scope.row.signerTitle.label }}
          </template>
        </el-table-column>
        <el-table-column align="left" label="Họ tên" prop="fullname" width="200" />
        <el-table-column align="left" label="Số văn bản" prop="count" width="200" />
        <el-table-column align="left" label="Ngày tạo" width="250">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="Hành động" width="150" fixed="right">
          <template #default="scope">
            <el-button type="primary" link icon="edit" size="small" class="table-button" @click="updateDocumentSignersFunc(scope.row)">Sửa</el-button>
            <el-button type="primary" link icon="delete" size="small" @click="deleteRow(scope.row)">Xoá</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination
          layout="total, sizes, prev, pager, next, jumper"
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :total="total"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="Thêm người kí mới">
      <el-form ref="elFormRef" :model="formData" label-position="right" :rules="rule" label-width="120px">
        <el-form-item label="Phòng ban:" prop="agencyId">
          <el-select
            v-model.number="formData.agencyId"
            :style="{ width: '100%' }"
            clearable
            filterable
            placeholder="Chọn phòng ban"
          >
            <el-option
              v-for="item in agencyOptions"
              :key="item.ID"
              :label="item.name"
              :value="item.ID"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="Chức danh:" prop="title">
          <el-select
            v-model.number="formData.title"
            :style="{ width: '100%' }"
            clearable
            filterable
            placeholder="Chức danh"
          >
            <el-option
              v-for="(item, index) in signerTitleOptions"
              :key="index"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="Họ và tên" prop="fullname">
          <el-input v-model="formData.fullname" :clearable="true" placeholder="Họ và tên đầy đủ" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialog">Đóng</el-button>
          <el-button size="small" type="primary" @click="enterDialog">Thêm</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: 'DocumentSigners'
}
</script>

<script setup>
import {
  createDocumentSigners,
  deleteDocumentSigners,
  deleteDocumentSignersByIds,
  updateDocumentSigners,
  findDocumentSigners,
  getDocumentSignersList
} from '@/api/documentSigners'

import { formatDate } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

import { getDocumentAgenciesList } from '../../api/documentAgencies'
import { getDict } from '../../utils/dictionary'

const formData = ref({
  agencyId: null,
  count: 0,
  fullname: '',
  title: null,
})

const rule = reactive({
})

const elFormRef = ref()

const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

const agencyOptions = ref([])
const signerTitleOptions = ref([])

const onReset = () => {
  searchInfo.value = {}
}

const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
  getTableData()
}

const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

const getTableData = async() => {
  const table = await getDocumentSignersList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

const loadAgencyOptions = async() => {
  const table = await getDocumentAgenciesList({ page: 1, pageSize: 1000 })
  if (table.code === 0) {
    agencyOptions.value = table.data.list
  }
}

const loadSignerTitleOptions = async() => {
  const data = await getDict('signerTitles')
  signerTitleOptions.value = data
}

getTableData()
loadAgencyOptions()
loadSignerTitleOptions()

const setOptions = async() => {
}

setOptions()

const multipleSelection = ref([])

const handleSelectionChange = (val) => {
  multipleSelection.value = val
}

const deleteRow = (row) => {
  ElMessageBox.confirm('Bạn có muốn xoá người kí này?', 'Cảnh bảo', {
    confirmButtonText: 'Xoá',
    cancelButtonText: 'Đóng',
    type: 'warning'
  }).then(() => {
    deleteDocumentSignersFunc(row)
  })
}

const deleteVisible = ref(false)

const onDelete = async() => {
  const ids = []
  if (multipleSelection.value.length === 0) {
    ElMessage({
      type: 'warning',
      message: 'Hãy chọn người cần xoá'
    })
    return
  }
  multipleSelection.value &&
        multipleSelection.value.map(item => {
          ids.push(item.ID)
        })
  const res = await deleteDocumentSignersByIds({ ids })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: 'Xoá bỏ thành công'
    })
    if (tableData.value.length === ids.length && page.value > 1) {
      page.value--
    }
    deleteVisible.value = false
    getTableData()
  }
}

const type = ref('')

const updateDocumentSignersFunc = async(row) => {
  const res = await findDocumentSigners({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.redocumentSigners
    dialogFormVisible.value = true
  }
}

const deleteDocumentSignersFunc = async(row) => {
  const res = await deleteDocumentSigners({ ID: row.ID })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: 'Xoá bỏ thành công'
    })
    if (tableData.value.length === 1 && page.value > 1) {
      page.value--
    }
    getTableData()
  }
}

const dialogFormVisible = ref(false)

const openDialog = () => {
  type.value = 'create'
  dialogFormVisible.value = true
}

const closeDialog = () => {
  dialogFormVisible.value = false
  formData.value = {
    agencyId: 0,
    count: 0,
    fullname: '',
    title: 0,
  }
}

const enterDialog = async() => {
     elFormRef.value?.validate(async(valid) => {
       if (!valid) return
       let res
       switch (type.value) {
         case 'create':
           res = await createDocumentSigners(formData.value)
           break
         case 'update':
           res = await updateDocumentSigners(formData.value)
           break
         default:
           res = await createDocumentSigners(formData.value)
           break
       }
       if (res.code === 0) {
         ElMessage({
           type: 'success',
           message: 'Thực hiện thao tác thành công'
         })
         closeDialog()
         getTableData()
       }
     })
}
</script>

<style>
</style>

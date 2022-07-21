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
      <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
      >
        <template #empty>Không có dữ liệu</template>
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="Tên" prop="name" width="250" />
        <el-table-column align="left" label="Đường dẫn" prop="url" width="300" />
        <el-table-column align="left" label="Loại" prop="tag" width="60" />
        <el-table-column align="left" label="Kích thước" prop="size" width="120" />
        <el-table-column align="left" label="Tài liệu" prop="documentId" width="300" />
        <el-table-column align="left" label="Ngày tạo" width="250">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="Hành động">
          <template #default="scope">
            <el-button type="primary" link icon="edit" size="small" class="table-button" @click="updateDocumentFilesFunc(scope.row)">Sửa</el-button>
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
  deleteDocumentFiles,
  deleteDocumentFilesByIds,
  updateDocumentFiles,
  findDocumentFiles,
  getDocumentFilesList
} from '@/api/documentFiles'

// Full introduction of formatting tools, please keep as needed
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

// Auto-generated dictionary (may be empty) and fields
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

// =========== Form Control Section ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

// Reset
const onReset = () => {
  searchInfo.value = {}
}

// Search
const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
  getTableData()
}

// Pagination
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// Modify page size
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// Get data
const getTableData = async() => {
  const table = await getDocumentFilesList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============= End of form control section ===============

// Get the required dictionary, may be empty, keep as needed
const setOptions = async() => {
}

// Get the required dictionary, may be empty, keep as needed
setOptions()

// Multiple selection data
const multipleSelection = ref([])
// Multiple choice
const handleSelectionChange = (val) => {
  multipleSelection.value = val
}

// Delete row
const deleteRow = (row) => {
  ElMessageBox.confirm('Bạn có chắc chắn muốn xoá?', 'Cảnh báo', {
    confirmButtonText: 'Xoá',
    cancelButtonText: 'Huỷ',
    type: 'warning'
  }).then(() => {
    deleteDocumentFilesFunc(row)
  })
}

// Batch delete control tag
const deleteVisible = ref(false)

// Bulk delete action
const onDelete = async() => {
  const ids = []
  if (multipleSelection.value.length === 0) {
    ElMessage({
      type: 'warning',
      message: 'Hãy chọn danh sách các mục cần xoá'
    })
    return
  }
  multipleSelection.value &&
        multipleSelection.value.map(item => {
          ids.push(item.ID)
        })
  const res = await deleteDocumentFilesByIds({ ids })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: 'Xoá thành công'
    })
    if (tableData.value.length === ids.length && page.value > 1) {
      page.value--
    }
    deleteVisible.value = false
    getTableData()
  }
}

// Behavior control mark (need to be added or changed inside the pop-up window)
const type = ref('')

// Update row
const updateDocumentFilesFunc = async(row) => {
  const res = await findDocumentFiles({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.redocumentFiles
    dialogFormVisible.value = true
  }
}

// Delete row
const deleteDocumentFilesFunc = async(row) => {
  const res = await deleteDocumentFiles({ ID: row.ID })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: 'Xoá thành công'
    })
    if (tableData.value.length === 1 && page.value > 1) {
      page.value--
    }
    getTableData()
  }
}

// Popup control marker
const dialogFormVisible = ref(false)

// Open the popup
const openDialog = () => {
  type.value = 'create'
  dialogFormVisible.value = true
}

// Close the popup
const closeDialog = () => {
  dialogFormVisible.value = false
  formData.value = {
    name: '',
    key: '',
    url: '',
    tag: '',
    size: 0,
    order: 0,
    documentId: 0,
    path: '',
  }
}
// Confirm the popup
const enterDialog = async() => {
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
         closeDialog()
         getTableData()
       }
     })
}
</script>

<style>
</style>

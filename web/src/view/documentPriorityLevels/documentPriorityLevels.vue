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
          <p>Xoá những mục đã chọn?</p>
          <div style="text-align: right; margin-top: 8px;">
            <el-button size="small" type="primary" link @click="deleteVisible = false">Huỷ bỏ</el-button>
            <el-button size="small" type="primary" @click="onDelete">Xác nhận</el-button>
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
        <el-table-column align="left" label="Mức cấp độ" prop="level" width="250" />
        <el-table-column align="left" label="Tên cấp độ" prop="name" width="250" />
        <el-table-column align="left" label="Ngày tạo" width="250">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="Hành động">
          <template #default="scope">
            <el-button type="primary" link icon="edit" size="small" class="table-button" @click="updateDocumentPriorityLevelsFunc(scope.row)">Sửa</el-button>
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
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="Hộp thoại cấp độ">
      <el-form ref="elFormRef" :model="formData" label-position="right" :rules="rule" label-width="120px">
        <el-form-item label="Mức cấp độ" prop="level">
          <el-input v-model.number="formData.level" :clearable="true" type="number" placeholder="Mức cấp độ" />
        </el-form-item>
        <el-form-item label="Tên cấp độ" prop="name">
          <el-input v-model="formData.name" :clearable="true" placeholder="Tên cấp độ" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialog">Đóng hộp thoại</el-button>
          <el-button size="small" type="primary" @click="enterDialog">Xác nhận</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: 'DocumentPriorityLevels'
}
</script>

<script setup>
import {
  createDocumentPriorityLevels,
  deleteDocumentPriorityLevels,
  deleteDocumentPriorityLevelsByIds,
  updateDocumentPriorityLevels,
  findDocumentPriorityLevels,
  getDocumentPriorityLevelsList
} from '@/api/documentPriorityLevels'

// Full introduction of formatting tools, please keep as needed
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

// Auto-generated dictionary (may be empty) and fields
const formData = ref({
  level: 0,
  name: '',
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
  const table = await getDocumentPriorityLevelsList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
    deleteDocumentPriorityLevelsFunc(row)
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
  const res = await deleteDocumentPriorityLevelsByIds({ ids })
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
const updateDocumentPriorityLevelsFunc = async(row) => {
  const res = await findDocumentPriorityLevels({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.redocumentPriorityLevels
    dialogFormVisible.value = true
  }
}

// Delete row
const deleteDocumentPriorityLevelsFunc = async(row) => {
  const res = await deleteDocumentPriorityLevels({ ID: row.ID })
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
    level: 0,
    name: '',
  }
}
// Confirm the popup
const enterDialog = async() => {
     elFormRef.value?.validate(async(valid) => {
       if (!valid) return
       let res
       switch (type.value) {
         case 'create':
           res = await createDocumentPriorityLevels(formData.value)
           break
         case 'update':
           res = await updateDocumentPriorityLevels(formData.value)
           break
         default:
           res = await createDocumentPriorityLevels(formData.value)
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

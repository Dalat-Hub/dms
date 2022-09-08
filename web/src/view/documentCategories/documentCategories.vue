<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="Tên">
          <el-input v-model="searchInfo.name" placeholder="Tên thể loại" />
        </el-form-item>
        <el-form-item label="Viết tắt">
          <el-input v-model="searchInfo.code" placeholder="Mã thể loại" />
        </el-form-item>
        <el-form-item label="Trạng thái" prop="hidden">
          <el-select v-model="searchInfo.hidden" clear placeholder="Trạng thái">
            <el-option key="false" label="Hiển thị" value="false" />
            <el-option key="true" label="Ẩn" value="true" />
          </el-select>
        </el-form-item>

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
          <p>Bạn có muốn xoá những thể loại đã chọn?</p>
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
        <el-table-column align="left" label="Thứ tự" prop="order" width="100" />
        <el-table-column align="left" label="Tên thể loại" prop="name" width="300" />
        <el-table-column align="left" label="Mã thể loại" prop="code" width="120" />
        <el-table-column align="left" label="Số văn bản" prop="count" width="120" />
        <el-table-column align="left" label="Trạng thái" width="100">
          <template #default="scope"><el-switch v-model="categoryStatus[scope.row.ID]" @change="(e) => hanleOnStatusChanged(e, scope.row.ID)" /></template>
        </el-table-column>
        <el-table-column align="left" label="Ngày tạo" width="250">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="Hành động">
          <template #default="scope">
            <el-button type="primary" link icon="edit" size="small" class="table-button" @click="updateDocumentCategoriesFunc(scope.row)">Sửa</el-button>
            <el-button type="primary" link icon="delete" size="small" @click="deleteRow(scope.row)">Xoá bỏ</el-button>
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
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="Cửa sổ thể loại">
      <el-form ref="elFormRef" :model="formData" label-position="right" :rules="rule" label-width="120px">
        <el-form-item label="Thứ tự:" prop="order">
          <el-input v-model.number="formData.order" :clearable="true" placeholder="Nhập thứ tự" />
        </el-form-item>
        <el-form-item label="Tên thể loại:" prop="name">
          <el-input v-model="formData.name" :clearable="true" placeholder="Nhập tên thể loại" />
        </el-form-item>
        <el-form-item label="Mã thể loại:" prop="code">
          <el-input v-model="formData.code" :clearable="true" placeholder="Nhập mã thể loại" />
        </el-form-item>
        <el-form-item label="Trạng thái:" prop="hidden">
          <el-select v-model="formData.hidden" placeholder="Trạng thái" :clearable="true">
            <el-option label="Hiển thị" :value="false" />
            <el-option label="Ẩn" :value="true" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-radio-group v-model="formData.validDocument">
            <el-radio :label="true" size="large">Thuộc thể loại văn bản hành chính</el-radio>
            <el-radio :label="false" size="large">Không thuộc thể loại văn bản hành chính</el-radio>
          </el-radio-group>
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
  name: 'DocumentCategories'
}
</script>

<script setup>
import {
  createDocumentCategories,
  deleteDocumentCategories,
  deleteDocumentCategoriesByIds,
  updateDocumentCategories,
  findDocumentCategories,
  getDocumentCategoriesList
} from '@/api/documentCategories'

// Full introduction of formatting tools, please keep as needed
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

// Auto-generated dictionary (may be empty) and fields
const formData = ref({
  name: '',
  code: '',
  count: 0,
  order: 0,
  hidden: false,
  validDocument: true
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
const categoryStatus = ref({})

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
  const table = await getDocumentCategoriesList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize

    const map = {}
    for (const e of table.data.list) {
      map[e.ID] = !e.hidden
    }

    categoryStatus.value = { ...map }
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
  ElMessageBox.confirm('Bạn có muốn xoá thể loại này?', 'Cảnh báo', {
    confirmButtonText: 'Xoá',
    cancelButtonText: 'Huỷ',
    type: 'warning'
  }).then(() => {
    deleteDocumentCategoriesFunc(row)
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
      message: 'Chọn những thể loại cần xoá'
    })
    return
  }
  multipleSelection.value &&
        multipleSelection.value.map(item => {
          ids.push(item.ID)
        })
  const res = await deleteDocumentCategoriesByIds({ ids })
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
const updateDocumentCategoriesFunc = async(row) => {
  const res = await findDocumentCategories({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.redocumentCategories
    dialogFormVisible.value = true
  }
}

// Delete row
const deleteDocumentCategoriesFunc = async(row) => {
  const res = await deleteDocumentCategories({ ID: row.ID })
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
    code: '',
    count: 0,
  }
}
// Confirm the popup
const enterDialog = async() => {
     elFormRef.value?.validate(async(valid) => {
       if (!valid) return
       let res
       switch (type.value) {
         case 'create':
           res = await createDocumentCategories(formData.value)
           break
         case 'update':
           res = await updateDocumentCategories(formData.value)
           break
         default:
           res = await createDocumentCategories(formData.value)
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

const hanleOnStatusChanged = async(newValue, categoryId) => {
  const category = tableData.value.find(s => s.ID === categoryId)
  if (!category) {
    alert('Thể loại không hợp lệ')
    return
  }

  const updatedData = {
    ...category,
    hidden: !newValue
  }

  const res = await updateDocumentCategories(updatedData)

  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: 'Thực hiện thao tác thành công'
    })
    getTableData()
  }
}
</script>

<style>
</style>

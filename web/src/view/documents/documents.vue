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
          <p>Xoá những tài liệu đã chọn?</p>
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
        <el-table-column align="left" label="Số hiệu" prop="signText" width="300" />
        <el-table-column align="left" label="Ngày ban hành" width="250">
          <template #default="scope">{{ formatDate(scope.row.dateIssued) }}</template>
        </el-table-column>
        <el-table-column align="left" label="Hiệu lực" prop="stillInEffect" width="120">
          <template #default="scope">{{ formatBoolean(scope.row.stillInEffect) }}</template>
        </el-table-column>
        <el-table-column align="left" label="Lượt xem" prop="viewCount" width="120" />
        <el-table-column align="left" label="Lượt tải" prop="downloadCount" width="120" />
        <el-table-column align="left" label="Trạng thái" prop="status" width="120" />
        <el-table-column align="left" label="Ngày tạo" width="250">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column fixed="right" label="Hành động" width="200">
          <template #default="scope">
            <el-button type="primary" link icon="edit" size="small" class="table-button" @click="showDocumentDetail(scope.row)">Xem</el-button>
            <el-button type="primary" link icon="edit" size="small" class="table-button" @click="updateDocumentsFunc(scope.row)">Sửa</el-button>
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
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="Hộp thoại văn bản">
      <el-form ref="elFormRef" :model="formData" label-position="right" :rules="rule" label-width="150px">
        <el-form-item label="Tiêu đề:" prop="title">
          <el-input v-model="formData.title" :clearable="true" placeholder="Tiêu đề văn bản" />
        </el-form-item>
        <el-form-item label="Mô tả ngắn" prop="expert">
          <el-input v-model="formData.expert" :clearable="true" placeholder="Mô tả ngắn của văn bản" />
        </el-form-item>
        <el-form-item label="Nội dung" prop="content">
          <el-input v-model="formData.content" :clearable="true" placeholder="Nội dung của văn bản" />
        </el-form-item>
        <el-form-item label="Ngày ban hành" prop="dateIssued">
          <el-date-picker v-model="formData.dateIssued" type="date" style="width:100%" placeholder="Ngày ban hành" :clearable="true" />
        </el-form-item>
        <el-form-item label="Ngày có hiệu lực" prop="effectDate">
          <el-date-picker v-model="formData.effectDate" type="date" style="width:100%" placeholder="Ngày có hiệu lực" :clearable="true" />
        </el-form-item>
        <el-form-item label="Ngày hết hiệu lực" prop="expirationDate">
          <el-date-picker v-model="formData.expirationDate" type="date" style="width:100%" placeholder="Ngày hết hiệu lực" :clearable="true" />
        </el-form-item>
        <el-form-item label="Số kí hiệu" prop="signNumber">
          <el-input v-model.number="formData.signNumber" :clearable="true" placeholder="Số kí hiệu" />
        </el-form-item>
        <el-form-item label="Năm ban hành" prop="signYear">
          <el-input v-model.number="formData.signYear" :clearable="true" placeholder="Năm ban hành" />
        </el-form-item>
        <el-form-item label="Thể loại" prop="categoryId">
          <el-input v-model.number="formData.categoryId" :clearable="true" placeholder="Thể loại" />
        </el-form-item>
        <el-form-item label="Cơ quan ban hành" prop="agencyId">
          <el-input v-model.number="formData.agencyId" :clearable="true" placeholder="Cơ quan ban hành" />
        </el-form-item>
        <el-form-item label="Tạo bởi" prop="createdBy">
          <el-input v-model.number="formData.createdBy" :clearable="true" placeholder="Tạo bởi" />
        </el-form-item>
        <el-form-item label="Chịu trách nhiệm bởi" prop="beResponsibleBy">
          <el-input v-model.number="formData.beResponsibleBy" :clearable="true" placeholder="Chịu trách nhiệm bởi" />
        </el-form-item>
        <el-form-item label="Trạng thái" prop="status">
          <el-input v-model="formData.status" :clearable="true" placeholder="Trạng thái" />
        </el-form-item>
        <el-form-item label="Độ ưu tiên" prop="priorityId">
          <el-input v-model.number="formData.priorityId" :clearable="true" placeholder="Độ ưu tiên" />
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
  name: 'Documents'
}
</script>

<script setup>
import {
  createDocuments,
  deleteDocuments,
  deleteDocumentsByIds,
  updateDocuments,
  findDocuments,
  getDocumentsList
} from '@/api/documents'

// Full introduction of formatting tools, please keep as needed
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()

// Auto-generated dictionary (may be empty) and fields
const formData = ref({
  title: '',
  expert: '',
  content: '',
  dateIssued: null,
  stillInEffect: false,
  effectDate: null,
  expirationDate: null,
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
  if (searchInfo.value.stillInEffect === '') {
    searchInfo.value.stillInEffect = null
  }
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
  const table = await getDocumentsList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
    deleteDocumentsFunc(row)
  })
}

const showDocumentDetail = (document) => {
  router.push({
    name: 'documents-detail',
    params: {
      document_id: document.ID
    }
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
  const res = await deleteDocumentsByIds({ ids })
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
const updateDocumentsFunc = async(row) => {
  const res = await findDocuments({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.redocuments
    dialogFormVisible.value = true
  }
}

// Delete row
const deleteDocumentsFunc = async(row) => {
  const res = await deleteDocuments({ ID: row.ID })
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
  }
}
// Confirm the popup
const enterDialog = async() => {
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
         closeDialog()
         getTableData()
       }
     })
}
</script>

<style>
</style>

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
        <el-table-column align="left" label="Số hiệu" width="300">
          <template #default="scope">
            <span v-if="scope.row.shortTitle"> {{ scope.row.shortTitle }}</span>
            <span v-else style="color: orangered;">N/A</span>
          </template>
        </el-table-column>
        <el-table-column align="left" label="Ngày ban hành" width="250">
          <template #default="scope">
            <span v-if="scope.row.dateIssued"> {{ formatDate(scope.row.dateIssued) }}</span>
            <span v-else style="color: orangered;">N/A</span>
          </template>
        </el-table-column>
        <el-table-column align="left" label="Hiệu lực" width="120">
          <template #default="scope">
            <el-tag v-if="scope.row.stillInEffect" class="ml-2" type="success">Còn</el-tag>
            <el-tag v-else class="ml-2" type="danger">Hết</el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="Lượt xem" prop="viewCount" width="120" />
        <el-table-column align="left" label="Lượt tải" prop="downloadCount" width="120" />
        <el-table-column align="left" label="Trạng thái" width="120">
          <template #default="scope"> {{ statusOptions[scope.row.status] || '--' }} </template>
        </el-table-column>
        <el-table-column align="left" label="Ngày tạo" width="250">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column fixed="right" label="Hành động" width="300">
          <template #default="scope">
            <el-button type="primary" link icon="edit" size="small" class="table-button" @click="showDocumentDetail(scope.row)">Xem</el-button>
            <el-button type="primary" link icon="edit" size="small" class="table-button" @click="showDocumentUpdate(scope.row)">Sửa</el-button>
            <el-button type="primary" link icon="edit" size="small" class="table-button" @click="duplicateDocument(scope.row)">Nhân bản</el-button>
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
  name: 'Documents'
}
</script>

<script setup>
import {
  deleteDocuments,
  deleteDocumentsByIds,
  getDocumentsList,
  makeDuplication
} from '@/api/documents'

// Full introduction of formatting tools, please keep as needed
import { formatDate, formatBoolean } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { getDict } from '../../utils/dictionary'

const router = useRouter()

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

const statusOptions = ref({ })

const loadStatusOptions = async() => {
  const data = await getDict('documentStatuses')

  statusOptions.value = data.reduce((acc, cur) => {
    return {
      ...acc,
      [cur.value]: cur.label
    }
  }, {})

  console.log(data)
}

getTableData()
loadStatusOptions()

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

const showDocumentUpdate = (document) => {
  router.push({
    name: 'documents-update',
    params: {
      id: document.ID
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

const duplicateDocument = async(document) => {
  const response = await makeDuplication({ id: document.ID })

  if (response.code === 0) {
    getTableData()

    ElMessage({
      type: 'success',
      message: 'Nhân bản tài liệu thành công'
    })
  }
}
</script>

<style>
</style>

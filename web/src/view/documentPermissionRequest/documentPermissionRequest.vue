<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item>
          <el-button
            size="small"
            type="primary"
            icon="search"
            @click="onSubmit"
          >Tìm kiếm</el-button>
          <el-button
            size="small"
            icon="refresh"
            @click="onReset"
          >Đặt lại bộ lọc</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button
          size="small"
          type="primary"
          icon="plus"
          @click="openDialog"
        >Cấp quyền cho người dùng</el-button>
        <el-popover v-model:visible="deleteVisible" placement="top" width="160">
          <p>Bạn có chắc muốn xoá những yêu cầu này?</p>
          <div style="text-align: right; margin-top: 8px">
            <el-button
              size="small"
              type="primary"
              link
              @click="deleteVisible = false"
            >Quay lại</el-button>
            <el-button
              size="small"
              type="primary"
              @click="onDelete"
            >Xoá</el-button>
          </div>
          <template #reference>
            <el-button
              icon="delete"
              size="small"
              style="margin-left: 10px"
              :disabled="!multipleSelection.length"
              @click="deleteVisible = true"
            >Xoá</el-button>
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
        <el-table-column
          align="left"
          label="Văn bản"
          prop="documentId"
          width="250"
        >
          <template #default="scope">
            {{ scope.row.document.shortTitle }}
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          label="Người yêu cầu"
          prop="requestUserId"
          width="200"
        >
          <template #default="scope">
            {{ scope.row.requestUser.nickName }}
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          label="Quyền"
          prop="requestPermission"
          width="100"
        />
        <el-table-column align="left" label="Yêu cầu lúc" width="180">
          <template #default="scope">{{
            formatDate(scope.row.CreatedAt)
          }}</template>
        </el-table-column>
        <el-table-column
          align="left"
          label="Người cấp"
          prop="acceptUserId"
          width="200"
        >
          <template #default="scope">
            <el-tag v-if="!scope.row.acceptUser" type="danger">N/A</el-tag>
            <span v-else>{{ scope.row.acceptUser.nickName }}</span>
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          label="Cấp lúc"
          prop="acceptAt"
          width="180"
        >
          <template #default="scope">
            <el-tag v-if="!scope.row.acceptAt" type="danger">N/A</el-tag>
            <span v-else>{{ formatDate(scope.row.acceptAt) }}</span>
          </template>
        </el-table-column>
        <el-table-column width="180" align="left" label="Hành động" fixed="right">
          <template #default="scope">
            <el-button
              type="primary"
              link
              icon="check"
              size="small"
              class="table-button"
              @click="approvePermission(scope.row)"
            >Phê duyệt</el-button>
            <el-button
              type="primary"
              link
              icon="delete"
              size="small"
              @click="deleteRow(scope.row)"
            >Xoá</el-button>
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
    <el-dialog
      v-model="dialogFormVisible"
      :before-close="closeDialog"
      title="Hộp thoại cấp quyền văn bản"
    >
      <el-form
        ref="elFormRef"
        :model="formData"
        label-position="right"
        :rules="rule"
        label-width="150px"
      >
        <el-form-item label="Người cấp quyền:" prop="acceptUserId">
          <el-input
            v-model.number="formData.acceptUserId"
            :clearable="true"
            placeholder="Người cấp quyền"
          />
        </el-form-item>
        <el-form-item label="Văn bản:" prop="documentId">
          <el-input
            v-model.number="formData.documentId"
            :clearable="true"
            placeholder="Văn bản"
          />
        </el-form-item>
        <el-form-item label="Người được cấp" prop="requestUserId">
          <el-input
            v-model.number="formData.requestUserId"
            :clearable="true"
            placeholder="Người được cấp"
          />
        </el-form-item>
        <el-form-item label="Quyền được cấp" prop="requestPermission">
          <el-input
            v-model="formData.requestPermission"
            :clearable="true"
            placeholder="Quyền được cấp"
          />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialog">Quay lại</el-button>
          <el-button
            size="small"
            type="primary"
            @click="enterDialog"
          >Xác nhận</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: 'DocumentPermissionRequest',
}
</script>

<script setup>
import {
  createDocumentPermissionRequest,
  deleteDocumentPermissionRequest,
  deleteDocumentPermissionRequestByIds,
  updateDocumentPermissionRequest,
  findDocumentPermissionRequest,
  getDocumentPermissionRequestList,
  approvePermissionRequest
} from '@/api/documentPermissionRequest'

import { formatDate } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

const formData = ref({
  acceptAt: new Date(),
  acceptUserId: 0,
  documentId: 0,
  requestPermission: '',
  requestUserId: 0,
})

const rule = reactive({})

const elFormRef = ref()

const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

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
  const table = await getDocumentPermissionRequestList({
    page: page.value,
    pageSize: pageSize.value,
    ...searchInfo.value,
  })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

const setOptions = async() => {}

setOptions()

const multipleSelection = ref([])
const handleSelectionChange = (val) => {
  multipleSelection.value = val
}

const deleteRow = (row) => {
  ElMessageBox.confirm('Bạn có muốn xoá những yêu cầu này?', 'Thông báo', {
    confirmButtonText: 'Xoá',
    cancelButtonText: 'Quay lại',
    type: 'warning',
  }).then(() => {
    deleteDocumentPermissionRequestFunc(row)
  })
}

const deleteVisible = ref(false)

const onDelete = async() => {
  const ids = []
  if (multipleSelection.value.length === 0) {
    ElMessage({
      type: 'warning',
      message: 'Chọn những yêu cầu cần xoá',
    })
    return
  }
  multipleSelection.value &&
    multipleSelection.value.map((item) => {
      ids.push(item.ID)
    })
  const res = await deleteDocumentPermissionRequestByIds({ ids })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: 'Thực hiện thao tác thành công',
    })
    if (tableData.value.length === ids.length && page.value > 1) {
      page.value--
    }
    deleteVisible.value = false
    getTableData()
  }
}

const type = ref('')

const updateDocumentPermissionRequestFunc = async(row) => {
  const res = await findDocumentPermissionRequest({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.redocumentPermissionRequest
    dialogFormVisible.value = true
  }
}

const deleteDocumentPermissionRequestFunc = async(row) => {
  const res = await deleteDocumentPermissionRequest({ ID: row.ID })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: 'Thực hiện thao tác thành công',
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
    acceptAt: new Date(),
    acceptUserId: 0,
    documentId: 0,
    requestPermission: '',
    requestUserId: 0,
  }
}

const enterDialog = async() => {
  elFormRef.value?.validate(async(valid) => {
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
        message: 'Thực hiện thao tác thành công',
      })
      closeDialog()
      getTableData()
    }
  })
}

const approvePermission = async(request) => {
  const response = await approvePermissionRequest(request)
  if (response.code === 0) {
    ElMessage({
      type: 'success',
      message: 'Phê duyệt yêu cầu thành công',
    })

    getTableData()
  }
}
</script>

<style></style>

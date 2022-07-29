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
        :row-class-name="tableRowClassName"
      >
        <el-table-column align="left" label="Văn bản" prop="shortTitle" width="300" />
        <el-table-column align="left" label="Sẽ được hoàn thiện bởi" prop="userId" width="300">
          <template #default="scope">
            <span v-if="scope.row.userId"> {{ scope.row?.user.nickName || 'N/A' }}</span>
            <span v-else style="color: orangered;">Chưa phân bổ</span>
          </template>
        </el-table-column>
        <el-table-column align="left" label="Xuất bản lúc" width="200">
          <template #default="scope">
            <span v-if="scope.row.doneAt"> {{ formatDate(scope.row.doneAt) }}</span>
            <span v-else style="color: orangered;">N/A</span>
          </template>
        </el-table-column>
        <el-table-column align="left" label="Tạo lúc" width="200">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="Thao tác" width="200" fixed="right">
          <template #default="scope">
            <el-button v-if="scope.row.userId != 0" type="primary" link icon="edit" size="small" class="table-button" @click="() => showDocumentUpdate(scope.row)">Cập nhật</el-button>
            <el-button v-else type="primary" link icon="edit" size="small" @click="() => openDistributionForm(scope.row)">Phân bổ</el-button>
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
      v-model="distributionFormVisible"
      :before-close="closeDistributionForm"
      title="Thêm nhanh văn bản"
    >
      <el-form :model="distributionFormData" label-position="right" label-width="150px">
        <el-form-item label="Tiêu đề">
          <el-input v-model="distributionFormData.title" readonly clearable />
        </el-form-item>
        <el-form-item label="Gắn thẻ người dùng">
          <el-select
            v-model="distributionFormData.relatedUsers"
            :style="{ width: '100%' }"
            clearable
            filterable
            multiple
            placeholder="Chọn người dùng liên quan"
          >
            <el-option
              v-for="item in usersOptions"
              :key="item.ID"
              :label="item.nickName"
              :value="item.ID"
            />
          </el-select>
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDistributionForm">Đóng</el-button>
          <el-button size="small" type="primary" @click="enterDistributionForm">Phân bổ</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: 'DocumentUsers'
}
</script>

<script setup>
import {
  getDocumentUsersList,
  distributeTasks
} from '@/api/documentUsers'
import { getUserList } from '../../api/user'

import { formatDate } from '@/utils/format'
import { ref } from 'vue'

import { ElMessage } from 'element-plus'
import { useRouter } from 'vue-router'

const router = useRouter()

const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})
const usersOptions = ref([])

const distributionFormData = ref({
  id: '',
  title: '',
  relatedUsers: []
})
const distributionFormVisible = ref(false)
const currentDocumentTask = ref(null)

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

const loadUserOptions = async() => {
  const table = await getUserList({ page: 1, pageSize: 1000 })
  if (table.code === 0) {
    usersOptions.value = table.data.list
  }
}

const getTableData = async() => {
  const table = await getDocumentUsersList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()
loadUserOptions()

const tableRowClassName = (data, rowIndex) => {
  if (data.row.userId === 0) { return 'danger-row' }

  if (data.row.doneAt) { return 'success-row' }

  return 'warning-row'
}

const closeDistributionForm = () => {
  distributionFormVisible.value = false
}

const openDistributionForm = (documentTask) => {
  distributionFormVisible.value = true
  currentDocumentTask.value = documentTask
  distributionFormData.value.title = documentTask.shortTitle
  distributionFormData.value.id = documentTask.ID
}

const enterDistributionForm = async() => {
  const distribution = {
    id: distributionFormData.value.id,
    documentId: 0,
    userIds: distributionFormData.value.relatedUsers
  }

  const response = await distributeTasks(distribution)

  if (response.code === 0) {
    ElMessage({
      type: 'success',
      message: 'Phân bổ người dùng thành công'
    })

    distributionFormData.value.id = null
    distributionFormData.value.relatedUsers = []
    distributionFormData.value.title = ''

    getTableData()
  }
}

const showDocumentUpdate = (task) => {
  router.push({
    name: 'documents-update',
    params: {
      id: task.documentId
    }
  })
}

</script>

<style>
.el-table .warning-row {
  background-color: #fdf6ec;
}
.el-table .danger-row {
  background-color: #ffeff1;
}
.el-table .success-row {
  background-color: #f0f9eb;
}
</style>

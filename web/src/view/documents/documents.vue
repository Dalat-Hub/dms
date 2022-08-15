<template>
  <div>
    <div class="gva-search-box">
      <el-form :model="searchInfo" class="demo-form-inline">
        <el-row>
          <el-col :lg="24">
            <el-form-item label="Số kí hiệu">
              <el-input v-model="searchInfo.signText" placeholder="11/2022-QĐ/ĐHĐL" />
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :lg="8">
            <el-form-item label="Cơ quan">
              <el-select v-model.number="searchInfo.agencyId" clear filterable placeholder="Chọn cơ quan ban hành" class="w-100">
                <el-option v-for="agency in documentAgencies" :key="agency.ID" :label="agency.name" :value="agency.ID" />
              </el-select>
            </el-form-item>
          </el-col>

          <el-col :lg="8">
            <el-form-item label="Thể loại">
              <el-select v-model.number="searchInfo.categoryId" clear filterable placeholder="Chọn thể loại" class="w-100">
                <el-option v-for="category in documentCategories" :key="category.ID" :label="category.name" :value="category.ID" />
              </el-select>
            </el-form-item>
          </el-col>

          <el-col :lg="8">
            <el-form-item label="Lĩnh vực">
              <el-select v-model.number="searchInfo.fieldId" clear filterable placeholder="Chọn lĩnh vực" class="w-100">
                <el-option v-for="field in documentFields" :key="field.ID" :label="field.name" :value="field.ID" />
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :lg="8">
            <el-form-item label="Người kí">
              <el-select v-model.number="searchInfo.signerId" clear filterable placeholder="Chọn người kí" class="w-100">
                <el-option v-for="signer in documentSigners" :key="signer.ID" :label="signer.fullname" :value="signer.ID" />
              </el-select>
            </el-form-item>
          </el-col>

          <el-col :lg="8">
            <el-form-item label="Người tạo">
              <el-select v-model.number="searchInfo.createdBy" clear filterable placeholder="Chọn người tạo" class="w-100">
                <el-option v-for="user in systemUsers" :key="user.ID" :label="user.nickName" :value="user.ID" />
              </el-select>
            </el-form-item>
          </el-col>

          <el-col :lg="8">
            <el-form-item label="Chịu trách nhiệm">
              <el-select v-model.number="searchInfo.beResponsibleBy" clear filterable placeholder="Chọn người chịu trách nhiệm" class="w-100">
                <el-option v-for="user in systemUsers" :key="user.ID" :label="user.nickName" :value="user.ID" />
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :lg="8">
            <el-form-item label="Hiệu lực">
              <el-select v-model="searchInfo.stillInEffect" clear filterable placeholder="Chọn cơ quan ban hành" class="w-100">
                <el-option key="true" label="Còn" value="true" />
                <el-option key="false" label="Hết" value="false" />
              </el-select>
            </el-form-item>
          </el-col>

          <el-col :lg="8">
            <el-form-item label="Mức độ">
              <el-select v-model.number="searchInfo.priority" clear filterable placeholder="Chọn mức độ ưu tiên" class="w-100">
                <el-option v-for="priority in priorityLevelOptions" :key="priority.value" :label="priority.label" :value="priority.value" />
              </el-select>
            </el-form-item>
          </el-col>

          <el-col :lg="8">
            <el-form-item label="Trạng thái">
              <el-select v-model.number="searchInfo.status" clear filterable placeholder="Chọn cơ quan ban hành" class="w-100">
                <el-option v-for="status in statusOptionsList" :key="status.value" :label="status.label" :value="status.value" />
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-form-item>
          <el-button size="small" type="primary" icon="search" @click="onSubmit">Tìm kiếm</el-button>
          <el-button size="small" type="primary" icon="search" @click="toggleAdvancedSearchBox">Nâng cao</el-button>
          <el-button size="small" icon="refresh" @click="onReset">Đặt lại bộ lọc</el-button>
        </el-form-item>
      </el-form>
    </div>

    <div v-if="showAdvancedSearchBox" class="gva-search-box">
      <el-form :model="advancedSearchInfo" class="demo-form-inline">

        <el-row :gutter="20">
          <el-col :lg="12">
            <el-form-item label="Từ ngày">
              <el-date-picker
                v-model="advancedSearchInfo.fromDate"
                type="date"
                placeholder="Chọn ngày bắt đầu"
                class="w-100"
                style="width: 100%"
              />
            </el-form-item>
          </el-col>

          <el-col :lg="12">
            <el-form-item label="Đến ngày">
              <el-date-picker
                v-model="advancedSearchInfo.toDate"
                type="date"
                placeholder="Chọn ngày kết thúc"
                class="w-100"
                style="width: 100%"
              />
            </el-form-item>
          </el-col>
        </el-row>

        <el-form-item>
          <el-button size="small" type="primary" icon="search" @click="onSubmit">Tìm kiếm nâng cao</el-button>
          <el-button size="small" icon="refresh" @click="onReset">Đặt lại bộ lọc</el-button>
          <el-button size="small" icon="close" @click="toggleAdvancedSearchBox">Đóng</el-button>
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
import { getDocumentAgenciesList } from '../../api/documentAgencies'
import { getDocumentCategoriesList } from '../../api/documentCategories'
import { getDocumentFieldsList } from '../../api/documentFields'
import { getDocumentSignersList } from '../../api/documentSigners'
import { getUserList } from '../../api/user'
import { getDict } from '../../utils/dictionary'

const router = useRouter()

// =========== Form Control Section ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({
  status: 2
})
const advancedSearchInfo = ref({})
const showAdvancedSearchBox = ref(false)

const documentAgencies = ref([])
const documentCategories = ref([])
const documentFields = ref([])
const documentSigners = ref([])
const systemUsers = ref([])

const priorityLevelOptions = ref([])
const statusOptionsList = ref([])

// Toggle advanced search box
const toggleAdvancedSearchBox = () => {
  showAdvancedSearchBox.value = !showAdvancedSearchBox.value
}

// Reset
const onReset = () => {
  searchInfo.value = {
    status: 2
  }
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
  if (searchInfo.value.stillInEffect === '') {
    searchInfo.value.stillInEffect = null
  }

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

  statusOptionsList.value = data
  statusOptions.value = data.reduce((acc, cur) => {
    return {
      ...acc,
      [cur.value]: cur.label
    }
  }, {})
}

getTableData()
loadStatusOptions()

// init search data

const loadDocumentAgenciesList = async() => {
  const response = await getDocumentAgenciesList()
  if (response.code === 0) {
    documentAgencies.value = response.data.list
  }
}

const loadDocumentCategoriesList = async() => {
  const response = await getDocumentCategoriesList()
  if (response.code === 0) {
    documentCategories.value = response.data.list
  }
}

const loadDocumentFieldsList = async() => {
  const response = await getDocumentFieldsList()
  if (response.code === 0) {
    documentFields.value = response.data.list
  }
}

const loadDocumentSigners = async() => {
  const response = await getDocumentSignersList()
  if (response.code === 0) {
    documentSigners.value = response.data.list
  }
}

const loadSystemUsers = async() => {
  const response = await getUserList({ page: 1, pageSize: 100 })
  if (response.code === 0) {
    systemUsers.value = response.data.list
  }
}

const loadPriorityOptions = async() => {
  const data = await getDict('documentPriorities')
  priorityLevelOptions.value = data
}

loadDocumentAgenciesList()
loadDocumentCategoriesList()
loadDocumentFieldsList()
loadDocumentSigners()
loadSystemUsers()
loadPriorityOptions()

// end of init search data section

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

<style scoped>
.w-100 {
  width: 100%;
}
</style>

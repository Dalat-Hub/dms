<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo">
        <el-form-item label="Label">
          <el-input v-model="searchInfo.label" placeholder="Search condition" />
        </el-form-item>
        <el-form-item label="Valuee">
          <el-input v-model="searchInfo.value" placeholder="Search condition" />
        </el-form-item>
        <el-form-item label="Status" prop="status">
          <el-select v-model="searchInfo.status" placeholder="Please choose">
            <el-option key="true" label="Yes" value="true" />
            <el-option key="false" label="No" value="false" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary" icon="search" @click="onSubmit">Search</el-button>
          <el-button size="small" icon="refresh" @click="onReset">Reset</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button size="small" type="primary" icon="plus" @click="openDialog">Add entry</el-button>
      </div>
      <el-table
        ref="multipleTable"
        :data="tableData"
        style="width: 100%"
        tooltip-effect="dark"
        row-key="ID"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="Date" width="180">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>

        <el-table-column align="left" label="Label" prop="label" width="120" />

        <el-table-column align="left" label="Value" prop="value" width="120" />

        <el-table-column align="left" label="Status" prop="status" width="120">
          <template #default="scope">{{ formatBoolean(scope.row.status) }}</template>
        </el-table-column>

        <el-table-column align="left" label="Order" prop="sort" width="120" />

        <el-table-column align="left" label="Operations">
          <template #default="scope">
            <el-button size="small" type="primary" link icon="edit" @click="updateSysDictionaryDetailFunc(scope.row)">Edit</el-button>
            <el-popover v-model="scope.row.visible" placement="top" width="160">
              <p>Are you sure?</p>
              <div style="text-align: right; margin-top: 8px;">
                <el-button size="small" type="primary" link @click="scope.row.visible = false">Cancel</el-button>
                <el-button type="primary" size="small" @click="deleteSysDictionaryDetailFunc(scope.row)">Delete</el-button>
              </div>
              <template #reference>
                <el-button type="primary" link icon="delete" size="small" @click="scope.row.visible = true">Delete</el-button>
              </template>
            </el-popover>
          </template>
        </el-table-column>
      </el-table>

      <div class="gva-pagination">
        <el-pagination
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :total="total"
          layout="total, sizes, prev, pager, next, jumper"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>

    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="Popup operations">
      <el-form ref="dialogForm" :model="formData" :rules="rules" size="medium" label-width="110px">
        <el-form-item label="Label" prop="label">
          <el-input
            v-model="formData.label"
            placeholder="Enter label"
            clearable
            :style="{width: '100%'}"
          />
        </el-form-item>
        <el-form-item label="Value" prop="value">
          <el-input-number
            v-model.number="formData.value"
            step-strictly
            :step="1"
            placeholder="Enter value"
            clearable
            :style="{width: '100%'}"
          />
        </el-form-item>
        <el-form-item label="Status" prop="status" required>
          <el-switch v-model="formData.status" active-text="Enable" inactive-text="Disable" />
        </el-form-item>
        <el-form-item label="Order" prop="sort">
          <el-input-number v-model.number="formData.sort" placeholder="Sort order" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialog">Cancel</el-button>
          <el-button size="small" type="primary" @click="enterDialog">Submit</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: 'SysDictionaryDetail'
}
</script>

<script setup>
import {
  createSysDictionaryDetail,
  deleteSysDictionaryDetail,
  updateSysDictionaryDetail,
  findSysDictionaryDetail,
  getSysDictionaryDetailList
} from '@/api/sysDictionaryDetail'
import { ref, watch } from 'vue'
import { useRoute } from 'vue-router'
import { ElMessage } from 'element-plus'
import { formatBoolean, formatDate } from '@/utils/format'
const route = useRoute()

watch(() => route.params.id, (id) => {
  searchInfo.value.sysDictionaryID = Number(id)
  getTableData()
})

const formData = ref({
  label: null,
  value: null,
  status: true,
  sort: null
})
const rules = ref({
  label: [
    {
      required: true,
      message: 'Please enter label',
      trigger: 'blur'
    }
  ],
  value: [
    {
      required: true,
      message: 'Please enter value',
      trigger: 'blur'
    }
  ],
  sort: [
    {
      required: true,
      message: 'Please enter sort order',
      trigger: 'blur'
    }
  ]
})

const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({ sysDictionaryID: Number(route.params.id) })
const onReset = () => {
  searchInfo.value = { sysDictionaryID: Number(route.params.id) }
}

const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
  if (searchInfo.value.status === '') {
    searchInfo.value.status = null
  }
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
  const table = await getSysDictionaryDetailList({
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

const type = ref('')
const dialogFormVisible = ref(false)
const updateSysDictionaryDetailFunc = async(row) => {
  const res = await findSysDictionaryDetail({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.reSysDictionaryDetail
    dialogFormVisible.value = true
  }
}

const closeDialog = () => {
  dialogFormVisible.value = false
  formData.value = {
    label: null,
    value: null,
    status: true,
    sort: null,
    sysDictionaryID: ''
  }
}
const deleteSysDictionaryDetailFunc = async(row) => {
  row.visible = false
  const res = await deleteSysDictionaryDetail({ ID: row.ID })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: 'Delete successfully'
    })
    if (tableData.value.length === 1 && page.value > 1) {
      page.value--
    }
    getTableData()
  }
}

const dialogForm = ref(null)
const enterDialog = async() => {
  formData.value.sysDictionaryID = Number(route.params.id)
  dialogForm.value.validate(async valid => {
    if (!valid) return
    let res
    switch (type.value) {
      case 'create':
        res = await createSysDictionaryDetail(formData.value)
        break
      case 'update':
        res = await updateSysDictionaryDetail(formData.value)
        break
      default:
        res = await createSysDictionaryDetail(formData.value)
        break
    }
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: 'Operate successfully'
      })
      closeDialog()
      getTableData()
    }
  })
}
const openDialog = () => {
  type.value = 'create'
  dialogFormVisible.value = true
}

</script>

<style>
</style>

<template>
  <div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button size="small" type="primary" icon="plus" @click="addMenu('0')">Thêm Menu Gốc</el-button>
      </div>

      <!-- Since the menu here corresponds to the list on the left, there is no need for paging. The default pageSize is 999 -->
      <el-table :data="tableData" row-key="ID">
        <el-table-column align="left" label="ID" min-width="100" prop="ID" />
        <el-table-column align="left" label="Display Name" min-width="120" prop="authorityName">
          <template #default="scope">
            <span>{{ scope.row.meta.title }}</span>
          </template>
        </el-table-column>
        <el-table-column align="left" label="Icon" min-width="140" prop="authorityName">
          <template #default="scope">
            <div v-if="scope.row.meta.icon" class="icon-column">
              <el-icon>
                <component :is="scope.row.meta.icon" />
              </el-icon>
              <span>{{ scope.row.meta.icon }}</span>
            </div>
          </template>
        </el-table-column>
        <el-table-column align="left" label="Route Name" show-overflow-tooltip min-width="160" prop="name" />
        <el-table-column align="left" label="Route Path" show-overflow-tooltip min-width="160" prop="path" />
        <el-table-column align="left" label="Hide Menu" min-width="100" prop="hidden">
          <template #default="scope">
            <span>{{ scope.row.hidden?"Hide":"Visible" }}</span>
          </template>
        </el-table-column>
        <el-table-column align="left" label="Parent" min-width="90" prop="parentId" />
        <el-table-column align="left" label="Sort" min-width="70" prop="sort" />
        <el-table-column align="left" label="File Path" min-width="360" prop="component" />
        <el-table-column align="left" fixed="right" label="Operation" width="300">
          <template #default="scope">
            <el-button
              size="small"
              type="primary"
              link
              icon="plus"
              @click="addMenu(scope.row.ID)"
            >Add SubMenu</el-button>
            <el-button
              size="small"
              type="primary"
              link
              icon="edit"
              @click="editMenu(scope.row.ID)"
            >Edit</el-button>
            <el-button
              size="small"
              type="primary"
              link
              icon="delete"
              @click="deleteMenu(scope.row.ID)"
            >Delete</el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>
    <el-dialog v-model="dialogFormVisible" :before-close="handleClose" :title="dialogTitle">
      <warning-bar title="New menu, you need to configure permissions in the Role Management to use" />
      <el-form
        v-if="dialogFormVisible"
        ref="menuForm"
        :inline="true"
        :model="form"
        :rules="rules"
        label-position="top"
        label-width="85px"
      >
        <el-form-item label="Route Name" prop="path" style="width:30%">
          <el-input
            v-model="form.name"
            autocomplete="off"
            placeholder="Unique English String"
            @change="changeName"
          />
        </el-form-item>
        <el-form-item prop="path" style="width:30%">
          <template #label>
            <div style="display:inline-flex">
              Route Path
              <el-checkbox v-model="checkFlag" style="float:right;margin-left:20px;">Add Params</el-checkbox>
            </div>
          </template>

          <el-input
            v-model="form.path"
            :disabled="!checkFlag"
            autocomplete="off"
            placeholder="It is recommended to splice parameters only in the rear"
          />
        </el-form-item>
        <el-form-item label="Whether to hide" style="width:30%">
          <el-select v-model="form.hidden" placeholder="Whether to hide the menu">
            <el-option :value="false" label="No" />
            <el-option :value="true" label="Yes" />
          </el-select>
        </el-form-item>
        <el-form-item label="Parent Menu" style="width:30%">
          <el-cascader
            v-model="form.parentId"
            style="width:100%"
            :disabled="!isEdit"
            :options="menuOption"
            :props="{ checkStrictly: true,label:'title',value:'ID',disabled:'disabled',emitPath:false}"
            :show-all-levels="false"
            filterable
          />
        </el-form-item>
        <el-form-item label="File Path" prop="component" style="width:60%">
          <el-input v-model="form.component" autocomplete="off" placeholder="Page: view/xxx/xx.vue Plugin: plugin/xx/xx.vue" @blur="fmtComponent" />
          <span style="font-size:12px;margin-right:12px;">If the menu contains submenus, create a router-view secondary routing page or</span><el-button style="margin-top:4px" size="small" @click="form.component = 'view/routerHolder.vue'">Click me to set</el-button>
        </el-form-item>
        <el-form-item label="Display Name" prop="meta.title" style="width:30%">
          <el-input v-model="form.meta.title" autocomplete="off" />
        </el-form-item>
        <el-form-item label="Icon" prop="meta.icon" style="width:30%">
          <icon :meta="form.meta" style="width:100%" />
        </el-form-item>
        <el-form-item label="Sort Order" prop="sort" style="width:30%">
          <el-input v-model.number="form.sort" autocomplete="off" />
        </el-form-item>
        <el-form-item label="Keep caches" prop="meta.keepAlive" style="width:30%">
          <el-select v-model="form.meta.keepAlive" style="width:100%" placeholder="Whether to keep keepAlive cache pages">
            <el-option :value="false" label="No" />
            <el-option :value="true" label="Yes" />
          </el-select>
        </el-form-item>
        <el-form-item label="Close tab automatically" prop="meta.closeTab" style="width:30%">
          <el-select v-model="form.meta.closeTab" style="width:100%" placeholder="Whether to close the tab automatically">
            <el-option :value="false" label="No" />
            <el-option :value="true" label="Yes" />
          </el-select>
        </el-form-item>
      </el-form>
      <div>
        <el-button
          size="small"
          type="primary"
          icon="edit"
          @click="addParameter(form)"
        >Add Menu Params</el-button>
        <el-table :data="form.parameters" style="width: 100%">
          <template #empty>No data</template>
          <el-table-column align="left" prop="type" label="Type" width="180">
            <template #default="scope">
              <el-select v-model="scope.row.type" placeholder="Please choose">
                <el-option key="query" value="query" label="query" />
                <el-option key="params" value="params" label="params" />
              </el-select>
            </template>
          </el-table-column>
          <el-table-column align="left" prop="key" label="Key" width="180">
            <template #default="scope">
              <div>
                <el-input v-model="scope.row.key" />
              </div>
            </template>
          </el-table-column>
          <el-table-column align="left" prop="value" label="Value">
            <template #default="scope">
              <div>
                <el-input v-model="scope.row.value" />
              </div>
            </template>
          </el-table-column>
          <el-table-column align="left">
            <template #default="scope">
              <div>
                <el-button
                  type="danger"
                  size="small"
                  icon="delete"
                  @click="deleteParameter(form.parameters,scope.$index)"
                >Delete</el-button>
              </div>
            </template>
          </el-table-column>
        </el-table>

        <el-button
          style="margin-top:12px"
          size="small"
          type="primary"
          icon="edit"
          @click="addBtn(form)"
        >Added controllable buttons</el-button>
        <el-table :data="form.menuBtn" style="width: 100%">
          <template #empty>No data</template>
          <el-table-column align="left" prop="name" label="Button Name" width="180">
            <template #default="scope">
              <div>
                <el-input v-model="scope.row.name" />
              </div>
            </template>
          </el-table-column>
          <el-table-column align="left" prop="name" label="Description" width="180">
            <template #default="scope">
              <div>
                <el-input v-model="scope.row.desc" />
              </div>
            </template>
          </el-table-column>
          <el-table-column align="left">
            <template #default="scope">
              <div>
                <el-button
                  type="danger"
                  size="small"
                  icon="delete"
                  @click="deleteBtn(form.menuBtn,scope.$index)"
                >Delete</el-button>
              </div>
            </template>
          </el-table-column>
        </el-table>
      </div>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialog">Cancel</el-button>
          <el-button size="small" type="primary" @click="enterDialog">Submit</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import {
  updateBaseMenu,
  getMenuList,
  addBaseMenu,
  deleteBaseMenu,
  getBaseMenuById
} from '@/api/menu'
import icon from '@/view/superAdmin/menu/icon.vue'
import WarningBar from '@/components/warningBar/warningBar.vue'
import { canRemoveAuthorityBtnApi } from '@/api/authorityBtn'
import { reactive, ref } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'

const rules = reactive({
  path: [{ required: true, message: 'Please enter menu name', trigger: 'blur' }],
  component: [
    { required: true, message: 'Please enter file path', trigger: 'blur' }
  ],
  'meta.title': [
    { required: true, message: 'Please enter menu display name', trigger: 'blur' }
  ]
})

const page = ref(1)
const total = ref(0)
const pageSize = ref(999)
const tableData = ref([])
const searchInfo = ref({})

// Get data
const getTableData = async() => {
  const table = await getMenuList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// Add params
const addParameter = (form) => {
  if (!form.parameters) {
    form.parameters = []
  }
  form.parameters.push({
    type: 'query',
    key: '',
    value: ''
  })
}

const fmtComponent = () => {
  form.value.component = form.value.component.replace(/\\/g, '/')
}

// Delete parameter
const deleteParameter = (parameters, index) => {
  parameters.splice(index, 1)
}

// Add controllable buttons
const addBtn = (form) => {
  if (!form.menuBtn) {
    form.menuBtn = []
  }
  form.menuBtn.push({
    name: '',
    desc: '',
  })
}

// Remove controllable buttons
const deleteBtn = async(btns, index) => {
  const btn = btns[index]
  if (btn.ID === 0) {
    btns.splice(index, 1)
    return
  }
  const res = await canRemoveAuthorityBtnApi({ id: btn.ID })
  if (res.code === 0) {
    btns.splice(index, 1)
    return
  }
}

const form = ref({
  ID: 0,
  path: '',
  name: '',
  hidden: '',
  parentId: '',
  component: '',
  meta: {
    title: '',
    icon: '',
    defaultMenu: false,
    closeTab: false,
    keepAlive: false
  },
  parameters: [],
  menuBtn: []
})
const changeName = () => {
  form.value.path = form.value.name
}

const handleClose = (done) => {
  initForm()
  done()
}

// Delete menu
const deleteMenu = (ID) => {
  ElMessageBox.confirm('This action will permanently delete all roles under this menu, do you want to continue?', 'Hint', {
    confirmButtonText: 'Yes',
    cancelButtonText: 'Cancel',
    type: 'warning'
  })
    .then(async() => {
      const res = await deleteBaseMenu({ ID })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: 'Successfully delete'
        })
        if (tableData.value.length === 1 && page.value > 1) {
          page.value--
        }
        getTableData()
      }
    })
    .catch(() => {
      ElMessage({
        type: 'info',
        message: 'Fail to delete'
      })
    })
}

// Initialize the form method in the pop-up window
const menuForm = ref(null)
const checkFlag = ref(false)
const initForm = () => {
  checkFlag.value = false
  menuForm.value.resetFields()
  form.value = {
    ID: 0,
    path: '',
    name: '',
    hidden: '',
    parentId: '',
    component: '',
    meta: {
      title: '',
      icon: '',
      defaultMenu: false,
      keepAlive: ''
    }
  }
}

// Close popup
const dialogFormVisible = ref(false)
const closeDialog = () => {
  initForm()
  dialogFormVisible.value = false
}

// Add new menu
const enterDialog = async() => {
  menuForm.value.validate(async valid => {
    if (valid) {
      let res
      if (isEdit.value) {
        res = await updateBaseMenu(form.value)
      } else {
        res = await addBaseMenu(form.value)
      }
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: isEdit.value ? 'Edited successfully' : 'Added successfully!'
        })
        getTableData()
      }
      initForm()
      dialogFormVisible.value = false
    }
  })
}

const menuOption = ref([
  {
    ID: '0',
    title: 'Root menu'
  }
])
const setOptions = () => {
  menuOption.value = [
    {
      ID: '0',
      title: 'Root directory'
    }
  ]
  setMenuOptions(tableData.value, menuOption.value, false)
}
const setMenuOptions = (menuData, optionsData, disabled) => {
  menuData &&
        menuData.forEach(item => {
          if (item.children && item.children.length) {
            const option = {
              title: item.meta.title,
              ID: String(item.ID),
              disabled: disabled || item.ID === form.value.ID,
              children: []
            }
            setMenuOptions(
              item.children,
              option.children,
              disabled || item.ID === form.value.ID
            )
            optionsData.push(option)
          } else {
            const option = {
              title: item.meta.title,
              ID: String(item.ID),
              disabled: disabled || item.ID === form.value.ID
            }
            optionsData.push(option)
          }
        })
}

// Add menu method, id is 0 to add root menu
const isEdit = ref(false)
const dialogTitle = ref('Add New Menu')
const addMenu = (id) => {
  dialogTitle.value = 'Add New Menu'
  form.value.parentId = String(id)
  isEdit.value = false
  setOptions()
  dialogFormVisible.value = true
}

// Modify menu method
const editMenu = async(id) => {
  dialogTitle.value = 'Edit Menu'
  const res = await getBaseMenuById({ id })
  form.value = res.data.menu
  isEdit.value = true
  setOptions()
  dialogFormVisible.value = true
}

</script>

<script>
export default {
  name: 'Menus',
}
</script>

<style scoped lang="scss">
.warning {
  color: #dc143c;
}
.icon-column{
  display: flex;
  align-items: center;
  .el-icon{
    margin-right: 8px;
  }
}
</style>

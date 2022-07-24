<template>
  <div>
    <el-form ref="elForm" :model="formData" label-position="top" label-width="150px">
      <el-row justify="center">
        <el-col :lg="22">
          <el-row :gutter="10">
            <el-col :lg="17" :sm="24">
              <div class="gva-card-box">
                <div class="el-card is-always-shadow gva-card quick-entrance">
                  <div class="el-card__body">
                    <el-form-item label="Tiêu đề văn bản">
                      <el-input
                        v-model="formData.title"
                        :style="{ width: '100%' }"
                        clearable
                        placeholder="Nhập tiêu đề văn bản"
                      />
                    </el-form-item>
                    <el-row align="bottom" justify="space-between" type="flex">
                      <el-col class="el-col-18 el-col-lg-20">
                        <el-form-item label="Cơ quan ban hành văn bản">
                          <el-select
                            v-model="formData.agency"
                            :style="{ width: '100%' }"
                            clearable
                            filterable
                            placeholder="Chọn cơ quan ban hành văn bản"
                            @change="handleOnAgencyChange"
                          >
                            <el-option
                              v-for="item in agencyOptions"
                              :key="item.ID"
                              :label="item.name"
                              :value="item.ID"
                            />
                          </el-select>
                        </el-form-item>
                      </el-col>
                      <el-col class="el-col-6 el-col-lg-4">
                        <el-form-item label="">
                          <el-button size="large" style="width: 100%" type="success" @click="openAgencyDialog"> Thêm Nhanh</el-button>
                        </el-form-item>
                      </el-col>
                    </el-row>
                    <el-row align="bottom" justify="space-between" type="flex">
                      <el-col class="el-col-18 el-col-lg-20">
                        <el-form-item label="Thể loại văn bản">
                          <el-select
                            v-model="formData.category"
                            :style="{ width: '100%' }"
                            clearable
                            filterable
                            placeholder="Chọn thể loại văn bản"
                            @change="handleOnCategoryChange"
                          >
                            <el-option
                              v-for="item in categoryOptions"
                              :key="item.ID"
                              :disabled="item.disabled"
                              :label="item.name"
                              :value="item.ID"
                            />
                          </el-select>
                        </el-form-item>
                      </el-col>
                      <el-col class="el-col-6 el-col-lg-4">
                        <el-form-item label="">
                          <el-button size="large" style="width: 100%" type="success"> Thêm Nhanh</el-button>
                        </el-form-item>
                      </el-col>
                    </el-row>
                    <el-row align="bottom" justify="space-between" type="flex">
                      <el-form-item class="el-col el-col-lg-8" label="Ngày ban hành">
                        <el-date-picker
                          v-model="formData.date_issued"
                          :style="{ width: '100%' }"
                          clearable
                          format="DD/MM/YYYY"
                          type="date"
                          @change="handleOnIssuedDateChange"
                        />
                      </el-form-item>
                      <el-form-item class="el-col el-col-lg-8" label="Ngày có hiệu lực">
                        <el-date-picker
                          v-model="formData.date_effected"
                          :style="{ width: '100%' }"
                          clearable
                          format="DD/MM/YYYY"
                          type="date"
                        />
                      </el-form-item>
                      <el-form-item class="el-col el-col-lg-8" label="Ngày hết hiệu lực">
                        <el-date-picker
                          v-model="formData.date_expiration"
                          :style="{ width: '100%' }"
                          clearable
                          format="DD/MM/YYYY"
                          type="date"
                        />
                      </el-form-item>
                    </el-row>
                    <el-row align="bottom" justify="space-between" type="flex">
                      <el-form-item class="el-col el-col-sm-12 el-col-lg-6" label="Số">
                        <el-input
                          v-model="formData.signNumber"
                          :style="{ width: '100%' }"
                          clearable
                          placeholder="Số văn bản"
                        />
                      </el-form-item>
                      <el-form-item class="el-col el-col-sm-12 el-col-lg-6" label="Năm">
                        <el-input
                          v-model="formData.signYear"
                          :style="{ width: '100%' }"
                          clearable
                          placeholder="Năm phát hành"
                        />
                      </el-form-item>
                      <el-form-item class="el-col el-col-sm-12 el-col-lg-6" label="Thể loại">
                        <el-input
                          v-model="formData.categoryReadonly"
                          :style="{ width: '100%' }"
                          placeholder="Chọn phía trên"
                          readonly
                        />
                      </el-form-item>
                      <el-form-item class="el-col el-col-sm-12 el-col-lg-6" label="Cơ quan">
                        <el-input
                          v-model="formData.agencyReadonly"
                          :style="{ width: '100%' }"
                          placeholder="Chọn phía trên"
                          readonly
                        />
                      </el-form-item>
                    </el-row>
                    <el-row align="bottom" justify="space-between" type="flex">
                      <el-col class="el-col-18 el-col-lg-20">
                        <el-form-item label="Lĩnh vực văn bản">
                          <el-select
                            v-model="formData.fields"
                            :style="{ width: '100%' }"
                            clearable
                            filterable
                            multiple
                            placeholder="Chọn 1 hoặc nhiều lĩnh vực"
                          >
                            <el-option
                              v-for="item in fieldsOptions"
                              :key="item.ID"
                              :label="item.name"
                              :value="item.ID"
                            />
                          </el-select>
                        </el-form-item>
                      </el-col>
                      <el-col class="el-col-6 el-col-lg-4">
                        <el-form-item label="">
                          <el-button size="large" style="width: 100%" type="success"> Thêm Nhanh</el-button>
                        </el-form-item>
                      </el-col>
                    </el-row>

                    <el-form-item label="Mô tả ngắn">
                      <el-input
                        v-model="formData.expert"
                        :style="{ width: '100%' }"
                        clearable
                        placeholder="Nhập mô tả ngắn nếu có"
                        rows="5"
                        type="textarea"
                      />
                    </el-form-item>
                    <el-form-item label="Nội dung">
                      <el-input
                        v-model="formData.content"
                        :style="{ width: '100%' }"
                        clearable
                        placeholder="Nhập nội dung văn bản nếu có"
                        rows="10"
                        type="textarea"
                      />
                    </el-form-item>
                  </div>
                </div>
              </div>
              <div class="gva-card-box" style="margin: 10px 0">
                <div class="el-card is-always-shadow gva-card quick-entrance">
                  <div class="el-card__body">
                    <el-row align="bottom" justify="space-between" type="flex">
                      <el-col class="el-col-18 el-col-lg-20">
                        <el-form-item label="Văn bản căn cứ">
                          <el-select
                            v-model="formData.baseDocuments"
                            :style="{ width: '100%' }"
                            clearable
                            filterable
                            multiple
                            placeholder="Chọn 1 hoặc nhiều văn bản căn cứ"
                          >
                            <el-option
                              v-for="item in documentsOptions"
                              :key="item.ID"
                              :label="item.title"
                              :value="item.ID"
                            />
                          </el-select>
                        </el-form-item>
                      </el-col>
                      <el-col class="el-col-6 el-col-lg-4">
                        <el-form-item label="">
                          <el-button
                            size="large"
                            style="width: 100%"
                            type="success"
                            @click="openQuickDocumentDialogForm"
                          > Thêm Nhanh
                          </el-button>
                        </el-form-item>
                      </el-col>
                    </el-row>

                    <el-row align="bottom" justify="space-between" type="flex">
                      <el-col class="el-col-18 el-col-lg-20">
                        <el-form-item label="Văn bản liên quan">
                          <el-select
                            v-model="formData.relatedDocuments"
                            :style="{ width: '100%' }"
                            clearable
                            filterable
                            multiple
                            placeholder="Chọn 1 hoặc nhiều văn bản liên quan"
                          >
                            <el-option
                              v-for="item in documentsOptions"
                              :key="item.ID"
                              :label="item.title"
                              :value="item.ID"
                            />
                          </el-select>
                        </el-form-item>
                      </el-col>
                      <el-col class="el-col-6 el-col-lg-4">
                        <el-form-item label="">
                          <el-button
                            size="large"
                            style="width: 100%;"
                            type="success"
                            @click="openQuickDocumentDialogForm"
                          > Thêm Nhanh
                          </el-button>
                        </el-form-item>
                      </el-col>
                    </el-row>

                    <el-form-item label="Phòng ban liên quan">
                      <el-select
                        v-model="formData.relatedAgencies"
                        :style="{ width: '100%' }"
                        clearable
                        filterable
                        multiple
                        placeholder="Chọn phòng ban liên quan"
                      >
                        <el-option
                          v-for="item in agencyOptions"
                          :key="item.ID"
                          :label="item.name"
                          :value="item.ID"
                        />
                      </el-select>
                    </el-form-item>

                    <el-form-item label="Người dùng liên quan">
                      <el-select
                        v-model="formData.relatedUsers"
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
                  </div>
                </div>
              </div>
            </el-col>
            <el-col :lg="7" :sm="24">
              <div class="gva-card-box">
                <div class="el-card is-always-shadow gva-card quick-entrance">
                  <div class="el-card__body">
                    <el-form-item label="Tải lên danh sách các tập tin đính kèm">
                      <el-upload
                        ref="documentFileUpload"
                        :action="`${path}/fileUploadAndDownload/upload`"
                        :auto-upload="false"
                        :before-upload="onBeforeUpload"
                        :file-list="documentFileList"
                        :headers="{ 'x-token': userStore.token }"
                        :limit="1"
                        :on-change="onSelectNewFile"
                        :on-error="onUploadError"
                        :on-remove="onRemoveFile"
                        :on-success="onUploadSuccess"
                        :show-file-list="true"
                      >
                        <el-button icon="el-icon-upload" size="small" type="primary">Chọn danh sách tập tin đính kèm
                        </el-button>
                      </el-upload>
                    </el-form-item>
                    <el-form-item label="Văn bản được tạo bởi">
                      <el-select
                        v-model="formData.createdBy"
                        :style="{ width: '100%' }"
                        clearable
                        filterable
                        placeholder="Chọn cá nhân tạo văn bản"
                      >
                        <el-option
                          v-for="item in usersOptions"
                          :key="item.ID"
                          :label="item.nickName"
                          :value="item.ID"
                        />
                      </el-select>
                    </el-form-item>
                    <el-form-item label="Được chịu trách nhiệm bởi">
                      <el-select
                        v-model="formData.beResponsibleBy"
                        :style="{ width: '100%' }"
                        clearable
                        filterable
                        placeholder="Chọn cá nhân chịu trách nhiệm"
                      >
                        <el-option
                          v-for="item in usersOptions"
                          :key="item.ID"
                          :label="item.nickName"
                          :value="item.ID"
                        />
                      </el-select>
                    </el-form-item>
                    <el-form-item label="Danh sách cá nhân kí văn bản">
                      <el-select
                        v-model="formData.signers"
                        :style="{ width: '100%' }"
                        clearable
                        filterable
                        multiple
                        placeholder="Chọn 1 hoặc nhiều cá nhân kí văn bản"
                      >
                        <el-option
                          v-for="item in usersOptions"
                          :key="item.ID"
                          :label="item.nickName"
                          :value="item.ID"
                        />
                      </el-select>
                    </el-form-item>
                    <el-form-item label="Độ ưu tiên">
                      <el-select
                        v-model="formData.priorityLevel"
                        :style="{ width: '100%' }"
                        clearable
                        placeholder="Mức độ ưu tiên của văn bản"
                      >
                        <el-option
                          v-for="item in priorityLevelOptions"
                          :key="item.value"
                          :label="item.label"
                          :value="item.value"
                        />
                      </el-select>
                    </el-form-item>
                    <el-form-item label="Trạng thái">
                      <el-select
                        v-model="formData.status"
                        :style="{ width: '100%' }"
                        clearable
                        placeholder="Trạng thái của văn bản"
                      >
                        <el-option
                          v-for="(item, index) in statusOptions"
                          :key="index"
                          :label="item.label"
                          :value="item.value"
                        />
                      </el-select>
                    </el-form-item>

                    <el-form-item size="large">
                      <el-button type="primary" @click="submitForm">Tạo văn bản mới</el-button>
                    </el-form-item>

                  </div>
                </div>
              </div>
            </el-col>
          </el-row>
        </el-col>
      </el-row>

    </el-form>

    <el-dialog v-model="agencyDialogFormVisible" :before-close="closeAgencyDialog" title="Thêm phòng ban nhanh">
      <el-form ref="elFormRef" :model="agencyFormData" label-position="right" label-width="120px">
        <el-form-item label="Tên phòng ban:">
          <el-input v-model="agencyFormData.name" :clearable="true" placeholder="Tên phòng ban" />
        </el-form-item>
        <el-form-item label="Mã phòng ban:">
          <el-input v-model="agencyFormData.code" :clearable="true" placeholder="Mã phòng ban" />
        </el-form-item>
        <el-form-item label="Cấp độ:">
          <el-select v-model="agencyFormData.level" placeholder="Chọn cấp độ" :clearable="true">
            <el-option v-for="item in agencyLevelOptions" :key="item.value" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeAgencyDialog">Đóng</el-button>
          <el-button size="small" type="primary" @click="enterAgencyDialog">Xác nhận</el-button>
        </div>
      </template>
    </el-dialog>

    <el-dialog
      v-model="quickDocumentDialogFormVisible"
      :before-close="closeQuickDocumentDialogForm"
      title="Thêm nhanh văn bản"
    >
      <el-form :model="quickDocumentFormData" label-position="right" label-width="80px">
        <el-form-item label="Tiêu đề">
          <el-input v-model="quickDocumentFormData.title" clearable placeholder="Tiêu đề văn bản" />
        </el-form-item>
        <el-form-item label="Liên quan" prop="relatedUsers">
          <el-select
            v-model="quickDocumentFormData.relatedUsers"
            :style="{ width: '100%' }"
            clearable
            filterable
            multiple
            placeholder="Chọn người dùng liên quan"
          >
            <el-option
              v-for="item in usersOptions"
              :key="item.ID"
              :disabled="item.disabled"
              :label="item.nickName"
              :value="item.ID"
            />
          </el-select>
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeQuickDocumentDialogForm">Đóng</el-button>
          <el-button size="small" type="primary" @click="createNewDraftDocument">Thêm nhanh</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: 'DocumentCreate'
}
</script>

<script setup>
import { ref } from 'vue'
import { useUserStore } from '@/pinia/modules/user'
import { getDict } from '../../utils/dictionary'
import { getDocumentAgenciesList } from '../../api/documentAgencies'
import { getDocumentCategoriesList } from '../../api/documentCategories'
import { getDocumentFieldsList } from '../../api/documentFields'
import { getDocumentsList } from '../../api/documents'
import { getUserList } from '../../api/user'

const userStore = useUserStore()

const formData = ref({
  title: '',
  agency: null,
  category: null,
  date_issued: null,
  date_effected: null,
  date_expiration: null,
  signNumber: null,
  signYear: null,
  categoryReadonly: '',
  agencyReadonly: '',
  fields: [],
  createdBy: userStore.userInfo.ID || 0,
  beResponsibleBy: userStore.userInfo.ID || 0,
  expert: '',
  content: '',
  priorityLevel: null,
  status: null,
  baseDocuments: [],
  relatedDocuments: [],
  relatedAgencies: [],
  relatedUsers: [],
  signers: [],
})

const agencyDialogFormVisible = ref(false)
const agencyFormData = ref({
  name: '',
  code: '',
  level: null
})

const quickDocumentFormData = ref({
  title: '',
  relatedUsers: []
})

const statusOptions = ref([])
const agencyOptions = ref([])
const priorityLevelOptions = ref([])
const usersOptions = ref([])
const documentsOptions = ref([])
const fieldsOptions = ref([])
const categoryOptions = ref([])
const agencyLevelOptions = ref([])
const documentFileList = ref([])
const quickDocumentDialogFormVisible = ref(false)
const path = ref([])

// ================= Prepare data section =================

const loadStatusOptions = async() => {
  const data = await getDict('documentStatuses')
  statusOptions.value = data
}

const loadPriorityOptions = async() => {
  const data = await getDict('documentPriorities')
  priorityLevelOptions.value = data
}

const loadAgencyOptions = async() => {
  const table = await getDocumentAgenciesList({ page: 1, pageSize: 1000 })
  if (table.code === 0) {
    agencyOptions.value = table.data.list
  }
}

const loadAgencyLevelOptions = async() => {
  const data = await getDict('agencyLevels')
  agencyLevelOptions.value = data
}

const loadCategoryOptions = async() => {
  const table = await getDocumentCategoriesList({ page: 1, pageSize: 1000 })
  if (table.code === 0) {
    categoryOptions.value = table.data.list
  }
}

const loadFieldOptions = async() => {
  const table = await getDocumentFieldsList({ page: 1, pageSize: 1000 })
  if (table.code === 0) {
    fieldsOptions.value = table.data.list
  }
}

const loadDocumentOptions = async() => {
  const table = await getDocumentsList({ page: 1, pageSize: 1000 })
  if (table.code === 0) {
    documentsOptions.value = table.data.list
  }
}

const loadUserOptions = async() => {
  const table = await getUserList({ page: 1, pageSize: 1000 })
  if (table.code === 0) {
    usersOptions.value = table.data.list
  }
}

loadStatusOptions()
loadAgencyOptions()
loadAgencyLevelOptions()
loadCategoryOptions()
loadFieldOptions()
loadPriorityOptions()
loadDocumentOptions()
loadUserOptions()

// ================= End of preparing data section =================

// ================= Reactive section =================

const handleOnAgencyChange = (selected) => {
  if (selected) {
    formData.value.agencyReadonly = agencyOptions.value.find(f => f.ID === selected)?.code
  }
}

const handleOnCategoryChange = (selected) => {
  if (selected) {
    formData.value.categoryReadonly = categoryOptions.value.find(f => f.ID === selected)?.code
  }
}

const handleOnIssuedDateChange = (date) => {
  const dateObj = new Date(date)

  const dd = dateObj.getDate()
  const mm = dateObj.getMonth()
  const yy = dateObj.getFullYear()

  formData.value.date_effected = new Date(yy, mm, dd)
}

const openAgencyDialog = () => {
  agencyDialogFormVisible.value = true
}

const closeAgencyDialog = () => {
  agencyDialogFormVisible.value = false
}

// ================= End of reactive section =================

const enterAgencyDialog = () => {
  console.log(agencyFormData.value)
}

const openQuickDocumentDialogForm = () => {

}

const onBeforeUpload = () => {

}

const onSelectNewFile = () => {

}

const onUploadError = () => {

}

const onRemoveFile = () => {

}

const onUploadSuccess = () => {

}

const submitForm = () => {

}

const closeQuickDocumentDialogForm = () => {

}

const createNewDraftDocument = () => {

}
</script>

<template>
  <div>
    <el-form ref="elForm" :model="formData" label-position="top" label-width="150px">
      <el-row justify="center">
        <el-col v-if="existingDocuments.length > 0" :span="22">
          <el-card class="box-card">
            <template #header>
              <div class="card-header">
                <span>Tìm thấy văn bản với số hiệu "{{ existingDocuments[0].signText }}" trên hệ thống</span>
              </div>
            </template>
            <div
              v-for="document in existingDocuments"
              :key="document.ID"
              class="text item"
              style="cursor: pointer;"
              @click="() => handleOnExistingDocumentClicked(document)"
            >{{ document.title }}</div>
          </el-card>
        </el-col>

        <el-col :lg="22">
          <el-row :gutter="10">

            <el-col :lg="17" :sm="24">
              <div class="gva-card-box">
                <div class="el-card is-always-shadow gva-card quick-entrance">
                  <div class="el-card__body">
                    <el-row align="bottom" justify="space-between" type="flex">
                      <el-col class="el-col-18 el-col-lg-20">
                        <el-form-item label="Số hiệu (Nhập tự động chỉ hoạt động với văn bản hành chính)">
                          <el-input
                            v-model="formData.signText"
                            :style="{ width: '100%' }"
                            clearable
                            placeholder="Số hiệu định danh văn bản (để trống nếu văn bản không có số hiệu)"
                            @blur="handleOnSignTextChanged"
                          />
                        </el-form-item>
                      </el-col>
                      <el-col class="el-col-6 el-col-lg-4">
                        <el-form-item label="">
                          <el-button size="large" style="width: 100%" type="success" @click="handleOnAutofillClick">Nhập tự động</el-button>
                        </el-form-item>
                      </el-col>
                    </el-row>
                  </div>
                </div>
              </div>
            </el-col>

            <el-col :lg="17" :sm="24" style="margin-top: 1rem;">
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
                            v-model="formData.agencies"
                            :style="{ width: '100%' }"
                            clearable
                            multiple
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
                          <el-button size="large" style="width: 100%" type="success" @click="openCategoryDialog"> Thêm Nhanh</el-button>
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
                          v-model.number="formData.signNumber"
                          :style="{ width: '100%' }"
                          clearable
                          placeholder="Số văn bản"
                        />
                      </el-form-item>
                      <el-form-item class="el-col el-col-sm-12 el-col-lg-6" label="Năm">
                        <el-input
                          v-model.number="formData.signYear"
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
                          <el-button size="large" style="width: 100%" type="success" @click="openFieldDialog"> Thêm Nhanh</el-button>
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
                              :label="item.shortTitle"
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
                            @click="() => openDocumentDialogForm('base')"
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
                              :label="item.shortTitle"
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
                            @click="() => openDocumentDialogForm('relation')"
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
                    <el-form-item label="Xem văn bản">
                      <el-radio-group v-model="documentRule.view">
                        <el-radio label="all" size="large">Tất cả mọi người</el-radio>
                        <el-radio label="limit" size="large">Hạn chế</el-radio>
                      </el-radio-group>
                    </el-form-item>
                    <el-form-item v-if="documentRule.view === 'limit'" label="Cá nhân xem văn bản">
                      <el-select
                        v-model="documentRule.viewUsers"
                        :style="{ width: '100%' }"
                        clearable
                        filterable
                        multiple
                        placeholder="Chọn 1 hoặc nhiều cá nhân xem văn bản"
                      >
                        <el-option
                          v-for="item in usersOptions"
                          :key="item.ID"
                          :label="item.nickName"
                          :value="item.ID"
                        />
                      </el-select>
                    </el-form-item>
                    <el-form-item v-if="documentRule.view === 'limit'" label="Nhóm xem văn bản">
                      <el-select
                        v-model="documentRule.viewRoles"
                        :style="{ width: '100%' }"
                        clearable
                        filterable
                        multiple
                        placeholder="Chọn 1 hoặc nhiều nhóm xem văn bản"
                      >
                        <el-option
                          v-for="item in roleOptions"
                          :key="item.authorityId"
                          :label="item.authorityName"
                          :value="item.authorityId"
                        />
                      </el-select>
                    </el-form-item>
                    <el-form-item label="Tải tập tin đính kèm">
                      <el-radio-group v-model="documentRule.download">
                        <el-radio label="all" size="large">Tất cả mọi người</el-radio>
                        <el-radio label="limit" size="large">Hạn chế</el-radio>
                      </el-radio-group>
                    </el-form-item>
                    <el-form-item v-if="documentRule.download === 'limit'" label="Cá nhân tải tập tin">
                      <el-select
                        v-model="documentRule.downloadUsers"
                        :style="{ width: '100%' }"
                        clearable
                        filterable
                        multiple
                        placeholder="Chọn 1 hoặc nhiều cá nhân tải tập tin"
                      >
                        <el-option
                          v-for="item in usersOptions"
                          :key="item.ID"
                          :label="item.nickName"
                          :value="item.ID"
                        />
                      </el-select>
                    </el-form-item>
                    <el-form-item v-if="documentRule.download === 'limit'" label="Nhóm tải tập tin">
                      <el-select
                        v-model="documentRule.downloadRoles"
                        :style="{ width: '100%' }"
                        clearable
                        filterable
                        multiple
                        placeholder="Chọn 1 hoặc nhiều nhóm tải tập tin"
                      >
                        <el-option
                          v-for="item in roleOptions"
                          :key="item.authorityId"
                          :label="item.authorityName"
                          :value="item.authorityId"
                        />
                      </el-select>
                    </el-form-item>
                    <el-form-item label="Chỉnh sủa">
                      <el-radio-group v-model="documentRule.edit">
                        <el-radio label="only" size="large">Chỉ tôi</el-radio>
                        <el-radio label="limit" size="large">Hạn chế</el-radio>
                      </el-radio-group>
                    </el-form-item>
                    <el-form-item v-if="documentRule.edit === 'limit'" label="Cá nhân sửa tập tin">
                      <el-select
                        v-model="documentRule.editUsers"
                        :style="{ width: '100%' }"
                        clearable
                        filterable
                        multiple
                        placeholder="Chọn 1 hoặc nhiều cá nhân sửa tập tin"
                      >
                        <el-option
                          v-for="item in usersOptions"
                          :key="item.ID"
                          :label="item.nickName"
                          :value="item.ID"
                        />
                      </el-select>
                    </el-form-item>
                    <el-form-item v-if="documentRule.edit === 'limit'" label="Nhóm sửa tập tin">
                      <el-select
                        v-model="documentRule.editRoles"
                        :style="{ width: '100%' }"
                        clearable
                        filterable
                        multiple
                        placeholder="Chọn 1 hoặc nhiều nhóm sửa tập tin"
                      >
                        <el-option
                          v-for="item in roleOptions"
                          :key="item.authorityId"
                          :label="item.authorityName"
                          :value="item.authorityId"
                        />
                      </el-select>
                    </el-form-item>
                    <el-form-item label="Sở hữu">
                      <el-radio-group v-model="documentRule.owner">
                        <el-radio label="only" size="large">Chỉ tôi</el-radio>
                        <el-radio label="limit" size="large">Hạn chế</el-radio>
                      </el-radio-group>
                    </el-form-item>
                    <el-form-item v-if="documentRule.owner === 'limit'" label="Cá nhân sở hữu tập tin">
                      <el-select
                        v-model="documentRule.ownerUsers"
                        :style="{ width: '100%' }"
                        clearable
                        filterable
                        multiple
                        placeholder="Chọn 1 hoặc nhiều cá nhân sở hữu tập tin"
                      >
                        <el-option
                          v-for="item in usersOptions"
                          :key="item.ID"
                          :label="item.nickName"
                          :value="item.ID"
                        />
                      </el-select>
                    </el-form-item>
                    <el-form-item v-if="documentRule.owner === 'limit'" label="Nhóm sở hữu tập tin">
                      <el-select
                        v-model="documentRule.ownerRoles"
                        :style="{ width: '100%' }"
                        clearable
                        filterable
                        multiple
                        placeholder="Chọn 1 hoặc nhiều nhóm sở hữu tập tin"
                      >
                        <el-option
                          v-for="item in roleOptions"
                          :key="item.authorityId"
                          :label="item.authorityName"
                          :value="item.authorityId"
                        />
                      </el-select>
                    </el-form-item>
                  </div>
                </div>
              </div>

              <div class="gva-card-box" style="margin-top: 1rem;">
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

                    <el-row align="bottom" justify="space-between" type="flex">
                      <el-col class="el-col-20 el-col-lg-21">
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
                              v-for="item in signerOptions"
                              :key="item.ID"
                              :label="`${signerTitleMap[item.title] || ''} ${item.fullname}`"
                              :value="item.ID"
                            />
                          </el-select>
                        </el-form-item>
                      </el-col>
                      <el-col class="el-col-4 el-col-lg-3">
                        <el-form-item>
                          <el-button size="large" type="primary" @click="openSignerDialog">+</el-button>
                        </el-form-item>
                      </el-col>
                    </el-row>
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
                      <el-button type="primary" :loading="!enableSubmitButton" @click="submitForm">
                        {{ enableSubmitButton ? 'Tạo văn bản mới' : 'Đang tạo văn bản mới ...' }}
                      </el-button>
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

    <el-dialog v-model="categoryDialogFormVisible" :before-close="closeCategoryDialog" title="Thêm thể loại nhanh">
      <el-form ref="elFormRef" :model="categoryFormData" label-position="right" label-width="120px">
        <el-form-item label="Tên thể loại:" prop="name">
          <el-input v-model="categoryFormData.name" :clearable="true" placeholder="Nhập tên thể loại" />
        </el-form-item>
        <el-form-item label="Mã thể loại:" prop="code">
          <el-input v-model="categoryFormData.code" :clearable="true" placeholder="Nhập mã thể loại" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeCategoryDialog">Đóng hộp thoại</el-button>
          <el-button size="small" type="primary" @click="enterCategoryDialog">Xác nhận</el-button>
        </div>
      </template>
    </el-dialog>

    <el-dialog v-model="fieldDialogFormVisible" :before-close="closeFieldDialog" title="Thêm nhanh lĩnh vực">
      <el-form ref="elFormRef" :model="fieldFormData" label-position="right" label-width="120px">
        <el-form-item label="Tên lĩnh vực:">
          <el-input v-model="fieldFormData.name" :clearable="true" placeholder="Nhập tên lĩnh vực" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeFieldDialog">Đóng hộp thoại</el-button>
          <el-button size="small" type="primary" @click="enterFieldDialog">Xác nhận</el-button>
        </div>
      </template>
    </el-dialog>

    <el-dialog
      v-model="documentDialogFormVisible"
      :before-close="closeDocumentDialog"
      title="Thêm nhanh văn bản"
    >
      <el-form :model="documentFormData" label-position="right" label-width="150px">
        <el-form-item label="Tiêu đề">
          <el-input v-model="documentFormData.title" clearable placeholder="Tiêu đề văn bản" />
          <p>Ví dụ: Luật Giáo dục ngày 14 tháng 6 năm 2019</p>
          <p>Hoặc: Nghị định số 69/2017/NĐ-CP ngày 25 tháng 5 năm 2017</p>
        </el-form-item>
        <el-form-item label="Số hiệu văn bản">
          <el-input v-model="documentFormData.signText" clearable placeholder="Số hiệu văn bản (nếu có)" />
          <p>Ví dụ: 69/2017/NĐ-CP</p>
          <p>Nếu VB do nhiều đơn vị ban hành, phân cách mỗi đơn vị bởi dấu @, <span style="display: block">VD: 69/2017/TTLT-BGDĐT@BCA</span></p>
        </el-form-item>
        <el-form-item label="Gắn thẻ người dùng">
          <el-select
            v-model="documentFormData.relatedUsers"
            :style="{ width: '100%' }"
            clearable
            filterable
            multiple
            placeholder="Chọn người dùng liên quan"
          >
            <el-option :key="0" label="Phân bổ sau" :value="0" />
            <el-option
              v-for="item in usersOptions"
              :key="item.ID"
              :label="item.nickName"
              :value="item.ID"
            />
          </el-select>
          <p>Trong trường hợp chưa xác định người dùng hoàn thiện văn bản này, vui lòng chọn "Phân bổ sau"</p>
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDocumentDialog">Đóng</el-button>
          <el-button size="small" type="primary" @click="enterDocumentDialog">Thêm nhanh</el-button>
        </div>
      </template>
    </el-dialog>

    <el-dialog v-model="signerDialogFormVisible" :before-close="closeSignerDialog" title="Thêm người kí mới">
      <el-form ref="elFormRef" :model="signerFormData" label-position="right" label-width="120px">
        <el-form-item label="Phòng ban:">
          <el-select
            v-model.number="signerFormData.agencyId"
            :style="{ width: '100%' }"
            clearable
            filterable
            placeholder="Chọn phòng ban"
          >
            <el-option
              v-for="item in agencyOptions"
              :key="item.ID"
              :label="item.name"
              :value="item.ID"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="Chức danh:">
          <el-select
            v-model.number="signerFormData.title"
            :style="{ width: '100%' }"
            clearable
            filterable
            placeholder="Chức danh"
          >
            <el-option
              v-for="(item, index) in signerTitleOptions"
              :key="index"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="Họ và tên">
          <el-input v-model="signerFormData.fullname" :clearable="true" placeholder="Họ và tên đầy đủ" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeSignerDialog">Đóng</el-button>
          <el-button size="small" type="primary" @click="enterSignerDialog">Thêm</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: 'DocumentCreate',
}
</script>

<script setup>
import { ref, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import { useRouter } from 'vue-router'

import { useUserStore } from '@/pinia/modules/user'
import { getDict } from '../../utils/dictionary'
import { createDocumentAgencies, getDocumentAgenciesList } from '../../api/documentAgencies'
import { createDocumentCategories, getDocumentCategoriesList } from '../../api/documentCategories'
import { createDocumentFields, getDocumentFieldsList } from '../../api/documentFields'
import { createDraftDocument, createFullDocument, getDocumentsList } from '../../api/documents'
import { getUserList } from '../../api/user'
import { getAuthorityInfo } from '../../api/authority'
import {
  createDocumentSigners,
  getDocumentSignersList
} from '@/api/documentSigners'

const router = useRouter()
const userStore = useUserStore()

const formData = ref({
  signText: '',
  notSupported: false,
  title: '',
  category: null,
  date_issued: null,
  date_effected: null,
  date_expiration: null,
  signNumber: null,
  signYear: null,
  categoryReadonly: '',
  agencyReadonly: '',
  fields: [],
  agencies: [],
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

const categoryDialogFormVisible = ref(false)
const categoryFormData = ref({
  name: '',
  code: ''
})

const fieldDialogFormVisible = ref(false)
const fieldFormData = ref({
  name: ''
})

const documentDialogFormVisible = ref(false)
const documentFormType = ref('')
const documentFormData = ref({
  title: '',
  signText: '',
  relatedUsers: []
})

const signerDialogFormVisible = ref(false)
const signerFormData = ref({
  agencyId: null,
  count: 0,
  fullname: '',
  title: null,
})
const signerTitleMap = ref({})

const statusOptions = ref([])
const agencyOptions = ref([])
const priorityLevelOptions = ref([])
const usersOptions = ref([])
const documentsOptions = ref([])
const fieldsOptions = ref([])
const categoryOptions = ref([])
const agencyLevelOptions = ref([])
const roleOptions = ref([])
const signerOptions = ref([])
const signerTitleOptions = ref([])
const documentFileList = ref([])
const path = import.meta.env.VITE_BASE_API
const enableSubmitButton = ref(true)
const hasFileAttached = ref(false)
const documentFileUpload = ref(null)

const documentRule = ref({
  view: 'all',
  download: 'all',
  edit: 'only',
  owner: 'only',
  viewUsers: [],
  viewRoles: [],
  downloadUsers: [],
  downloadRoles: [],
  editUsers: [],
  editRoles: [],
  ownerUsers: [],
  ownerRoles: []
})

const existingDocuments = ref([])

// ================= Init ref =============================

onMounted(() => {
  // console.log(documentFileUpload.value)
})

// ================= End of init ref ======================

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

const loadRoleOptions = async() => {
  const table = await getAuthorityInfo({ authorityId: 10 })

  if (table.code === 0) {
    roleOptions.value = table.data.authority.children
  }
}

const loadSignerOptions = async() => {
  const table = await getDocumentSignersList({ page: 1, pageSize: 1000 })
  if (table.code === 0) {
    signerOptions.value = table.data.list
  }
}

const loadSignerTitleOptions = async() => {
  const data = await getDict('signerTitles')
  signerTitleOptions.value = data

  signerTitleMap.value = data.reduce((acc, cur) => {
    return {
      ...acc,
      [cur.value]: cur.label
    }
  }, {})
}

loadStatusOptions()
loadAgencyOptions()
loadAgencyLevelOptions()
loadCategoryOptions()
loadFieldOptions()
loadPriorityOptions()
loadDocumentOptions()
loadUserOptions()
loadRoleOptions()
loadSignerOptions()
loadSignerTitleOptions()

// ================= End of preparing data section =================

// ================= Reactive section =================

const handleOnExistingDocumentClicked = (document) => {
  router.push({
    name: 'documents-detail',
    params: {
      document_id: document.ID
    }
  })
}

const handleOnSignTextChanged = async() => {
  if (!formData.value.signText) return

  formData.value.signText = formData.value.signText.toUpperCase()
}

const handleOnAutofillClick = async() => {
  if (!formData.value.signText) {
    ElMessage({
      type: 'error',
      message: 'Vui lòng nhập số hiệu văn bản để tự động điền'
    })
    return
  }

  existingDocuments.value = []

  formData.value.signText = formData.value.signText.toUpperCase()

  const text = formData.value.signText
  const parts = text.split('/')

  if (parts.length !== 3) {
    ElMessage({
      type: 'error',
      message: 'Số kí hiệu không đúng định dạng (số/năm/thể loại - đơn vị)'
    })
    return
  }

  const [signNumber, signYear, another] = parts
  const anotherParts = another.split('-')

  if (anotherParts.length < 2) {
    ElMessage({
      type: 'error',
      message: 'Số kí hiệu không đúng định dạng (số/năm/thể loại - đơn vị)'
    })
    return
  }

  const [categoryText, ...agenciesParts] = anotherParts
  let agencyText = agenciesParts.join('-')

  const agencyIDs = []
  if (agencyText.includes('@')) {
    const parts = agencyText.split('@')

    for (const p of parts) {
      const agency = agencyOptions.value.find(s => s.code.toUpperCase() === p)?.ID || null

      if (!agency) {
        ElMessage({
          type: 'error',
          message: `Đơn vị ${p} không tồn tại trên hệ thống`
        })
        return
      }

      agencyIDs.push(agency)
    }
  } else {
    const agencyId = agencyOptions.value.find(s => s.code.toUpperCase() === agencyText)?.ID || null
    if (!agencyId) {
      ElMessage({
        type: 'error',
        message: `Đơn vị ${agencyText} không tồn tại trên hệ thống`
      })
      return
    }

    agencyIDs.push(agencyId)
  }

  const categoryId = categoryOptions.value.find(s => s.code.toUpperCase() === categoryText)?.ID || null

  if (!categoryId) {
    ElMessage({
      type: 'error',
      message: `Thể loại văn bản ${categoryText} không tồn tại trên hệ thống`
    })
    return
  }

  agencyText = agencyText.replace(/@/g, '-')

  formData.value.signNumber = parseInt(signNumber)
  formData.value.signYear = parseInt(signYear)
  formData.value.categoryReadonly = categoryText
  formData.value.agencyReadonly = agencyText

  formData.value.agencies = [...agencyIDs]
  formData.value.category = categoryId

  // check if the document with the same signText exists on the DB
  const signNumberText = (signNumber * 1) > 9 ? `${signNumber}` : `0${parseInt(signNumber)}`
  const keyword = `${signNumberText}/${signYear}/${categoryText}-${agencyText}`
  const res = await getDocumentsList({
    keyword: keyword,
    page: 1,
    pageSize: 10
  })

  if (res.code === 0) {
    if (Array.isArray(res.data?.list) && res.data.list.length > 0) {
      existingDocuments.value = res.data.list.filter(s => s.signText === keyword)
    }
  }
}

const handleOnAgencyChange = (choices) => {
  if (choices) {
    const agencies = agencyOptions.value.filter(f => choices.includes(f.ID))
    const agencyText = agencies.map(e => e.code).join('-')

    formData.value.agencyReadonly = agencyText
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
  formData.value.signYear = yy
}

const openAgencyDialog = () => {
  agencyDialogFormVisible.value = true
}

const closeAgencyDialog = () => {
  agencyDialogFormVisible.value = false
}

const openCategoryDialog = () => {
  categoryDialogFormVisible.value = true
}

const closeCategoryDialog = () => {
  categoryDialogFormVisible.value = false
}

const openFieldDialog = () => {
  fieldDialogFormVisible.value = true
}

const closeFieldDialog = () => {
  fieldDialogFormVisible.value = false
}

const openDocumentDialogForm = (type) => {
  documentFormType.value = type
  documentDialogFormVisible.value = true
}

const closeDocumentDialog = () => {
  documentDialogFormVisible.value = false
}

const openSignerDialog = () => {
  signerDialogFormVisible.value = true
}

const closeSignerDialog = () => {
  signerDialogFormVisible.value = false
}

// ================= End of reactive section =================

// ================= Business section =================

const enterAgencyDialog = async() => {
  const res = await createDocumentAgencies(agencyFormData.value)

  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: 'Tạo phòng ban thành công'
    })

    agencyFormData.value.code = ''
    agencyFormData.value.name = ''
    agencyFormData.value.level = null

    agencyOptions.value = [...agencyOptions.value, res.data.agency]
    formData.value.agency = res.data.agency.ID
    formData.value.agencyReadonly = res.data.agency.code

    closeAgencyDialog()
  } else {
    ElMessage({
      type: 'error',
      message: 'Có lỗi diễn ra: ' + res.message
    })
  }
}

const enterCategoryDialog = async() => {
  const res = await createDocumentCategories(categoryFormData.value)

  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: 'Tạo thể loại thành công'
    })

    categoryFormData.value.code = ''
    categoryFormData.value.name = ''

    categoryOptions.value = [...categoryOptions.value, res.data.category]
    formData.value.category = res.data.category.ID
    formData.value.categoryReadonly = res.data.category.code

    closeCategoryDialog()
  } else {
    ElMessage({
      type: 'error',
      message: 'Có lỗi diễn ra: ' + res.message
    })
  }
}

const enterFieldDialog = async() => {
  const res = await createDocumentFields(fieldFormData.value)

  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: 'Tạo lĩnh vực thành công'
    })

    fieldFormData.value.name = ''

    fieldsOptions.value = [...fieldsOptions.value, res.data.field]
    formData.value.fields = [...formData.value.fields, res.data.field.ID]

    closeFieldDialog()
  } else {
    ElMessage({
      type: 'error',
      message: 'Có lỗi diễn ra: ' + res.message
    })
  }
}

const enterDocumentDialog = async() => {
  let response

  if (documentFormData.value.signText) {
    const text = documentFormData.value.signText.toUpperCase()
    const parts = text.split('/')

    if (parts.length !== 3) {
      alert('Số kí hiệu không đúng định dạng')
      return
    }

    const [signNumber, signYear, another] = parts
    const anotherParts = another.split('-')

    if (anotherParts.length < 2) {
      alert('Số kí hiệu không đúng định dạng')
      return
    }

    const [categoryText, ...agenciesParts] = anotherParts
    let agencyText = agenciesParts.join('-')

    const agencyIDs = []
    if (agencyText.includes('@')) {
      const parts = agencyText.split('@')

      for (const p of parts) {
        const agency = agencyOptions.value.find(s => s.code.toUpperCase() === p)?.ID || null

        if (!agency) {
          alert('Phòng ban ' + p + ' không tồn tại')
          return
        }

        agencyIDs.push(agency)
      }
    } else {
      const agencyId = agencyOptions.value.find(s => s.code.toUpperCase() === agencyText)?.ID || null
      if (!agencyId) {
        alert('Cơ quan ban hành văn bản không tồn tại trên hệ thống')
        return
      }

      agencyIDs.push(agencyId)
    }

    const categoryId = categoryOptions.value.find(s => s.code.toUpperCase() === categoryText)?.ID || null

    if (!categoryId) {
      alert('Thể loại văn bản không tồn tại trên hệ thống')
      return
    }

    const signNumberAsText = (signNumber * 1) > 9 ? `${signNumber}` : `0${signNumber}`

    agencyText = agencyText.replace(/@/g, '-')
    const signText = `${signNumberAsText}/${signYear}/${categoryText}-${agencyText}`

    const draftData = {
      ...documentFormData.value,
      signText: signText,
      category: categoryId,
      agencies: agencyIDs,
      signNumber: parseInt(signNumber),
      signYear: parseInt(signYear),
      categoryText: categoryText,
      agencyText: agencyText
    }

    response = await createDraftDocument(draftData)
  } else {
    response = await createDraftDocument({ ...documentFormData.value })
  }

  if (response.code === 0) {
    ElMessage({
      type: 'success',
      message: 'Thêm nhanh văn bản thành công'
    })

    documentFormData.value.title = ''
    documentFormData.value.signText = ''
    documentFormData.value.relatedUsers = []

    documentsOptions.value = [...documentsOptions.value, response.data.document]
    if (documentFormType.value === 'relation') {
      formData.value.relatedDocuments = [...formData.value.relatedDocuments, response.data.document.ID]
    } else if (documentFormType.value === 'base') {
      formData.value.baseDocuments = [...formData.value.baseDocuments, response.data.document.ID]
    }

    closeDocumentDialog()
  }
}

const enterSignerDialog = async() => {
  const response = await createDocumentSigners(signerFormData.value)

  if (response.code === 0) {
    ElMessage({
      type: 'success',
      message: 'Thêm nhanh người kí thành công'
    })

    signerFormData.value.agencyId = null
    signerFormData.value.count = 0
    signerFormData.value.fullname = ''
    signerFormData.value.title = null

    signerOptions.value = [...signerOptions.value, response.data.signer]
    formData.value.signers = [...formData.value.signers, response.data.signer.ID]

    closeSignerDialog()
  }
}

const onBeforeUpload = () => { }

const onSelectNewFile = () => {
  hasFileAttached.value = true
}

const onUploadError = (error) => {
  hasFileAttached.value = false

  ElMessage({
    type: 'error',
    message: 'Lỗi xảy ra trong quá trình tải lên tập tin! ' + error,
  })
}

const onRemoveFile = () => {
  enableSubmitButton.value = true
}

const onUploadSuccess = async(response, uploadFile, uploadFiles) => {
  const data = response.data

  await createNewDocument({
    name: data.file.name || '',
    key: data.file.key || '',
    url: data.file.url || '',
    tag: data.file.tag || '',
    size: uploadFile.size || '',
    order: 0,
    path: data.file.url || '',
  })
}

const submitForm = () => {
  enableSubmitButton.value = false

  if (hasFileAttached.value) {
    documentFileUpload.value.submit()
  } else {
    createNewDocument(null)
  }
}

const createNewDocument = async(fileInfo) => {
  const signNumber = formData.value.signNumber < 10 ? `0${formData.value.signNumber}` : formData.value.signNumber.toString()
  const selectedSigners = signerOptions.value.filter(s => formData.value.signers.includes(s.ID))

  const signerText = selectedSigners.map(s => {
    return `${signerTitleMap.value[s.title] || ''} ${s.fullname}`
  }).join(',')

  const documentData = {
    title: formData.value.title,
    expert: formData.value.expert,
    content: formData.value.content,
    dateIssued: formData.value.date_issued,
    effectDate: formData.value.date_effected,
    stillInEffect: true,
    expirationDate: formData.value.date_expiration,
    signNumber: parseInt(formData.value.signNumber),
    signYear: parseInt(formData.value.signYear),
    signCategory: formData.value.categoryReadonly,
    signAgency: formData.value.agencyReadonly,
    signText: `${signNumber}/${formData.value.signYear}/${formData.value.categoryReadonly}-${formData.value.agencyReadonly}`,
    categoryId: formData.value.category,
    agencies: [],
    createdBy: userStore.userInfo.ID || formData.value.createdBy,
    beResponsibleBy: formData.value.beResponsibleBy,
    status: formData.value.status,
    priorityId: formData.value.priorityLevel,
    fields: [],
    signers: [],
    documentBaseOns: [],
    documentReferences: [],
    documentUsersRelated: [],
    documentAgenciesRelated: [],
    fileInfo: null,
    ruleInfo: documentRule.value,
    signerText: signerText
  }

  documentData.agencies = formData.value.agencies
  documentData.fields = formData.value.fields
  documentData.signers = formData.value.signers
  documentData.documentBaseOns = formData.value.baseDocuments
  documentData.documentReferences = formData.value.relatedDocuments
  documentData.documentUsersRelated = formData.value.relatedUsers
  documentData.documentAgenciesRelated = formData.value.relatedAgencies

  if (fileInfo) {
    documentData.fileInfo = {
      name: fileInfo?.name || '',
      key: fileInfo?.key || '',
      url: fileInfo?.url || '',
      tag: fileInfo?.tag || '',
      size: fileInfo?.size || 0,
      order: 0,
      path: fileInfo?.url || '',
    }
  }

  const documentResponse = await createFullDocument(documentData)
  enableSubmitButton.value = true

  if (documentResponse.code === 0) {
    ElMessage({
      type: 'success',
      message: 'Tạo văn bản thành công',
    })

    setTimeout(() => {
      router.push({
        name: 'documents-all'
      })
    }, 3000)
  }
}

// ================= End of business section =================

</script>

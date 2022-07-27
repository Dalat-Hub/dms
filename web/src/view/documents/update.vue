<template>
  <div>
    <el-row justify="center">
      <el-col :lg="22">
        <el-row :gutter="10">
          <el-col :lg="17" :sm="24">
            <div class="gva-card-box">
              <div class="el-card is-always-shadow gva-card quick-entrance">
                <div class="el-card__body">
                  <div class="el-card__heading">
                    <h3>Thông tin sở hữu</h3>
                  </div>
                  <el-form ref="elForm" :model="formOwnerData" label-position="top" label-width="150px">
                    <el-form-item label="Văn bản được tạo bởi">
                      <el-col :span="11">
                        <el-input
                          v-model="createdUserName"
                          :style="{ width: '100%' }"
                          readonly
                        />
                      </el-col>
                      <el-col class="text-center" :span="2" style="width: 100%; text-align: center;">-</el-col>

                      <el-col :span="11">
                        <el-input
                          v-model="createdTime"
                          :style="{ width: '100%' }"
                          readonly
                        />
                      </el-col>

                    </el-form-item>
                    <el-form-item label="Văn bản được cập nhật bởi">
                      <el-select
                        v-model="formOwnerData.updatedBy"
                        :style="{ width: '100%' }"
                        clearable
                        filterable
                        placeholder="Chọn cá nhân cập nhật văn bản"
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
                        v-model="formOwnerData.beResponsibleBy"
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

                  </el-form>
                </div>
              </div>
            </div>

            <div class="gva-card-box" style="margin: 1rem 0">
              <div class="el-card is-always-shadow gva-card quick-entrance">
                <div class="el-card__body">
                  <div class="el-card__heading">
                    <h3>Cập nhật thông tin cơ bản</h3>
                  </div>
                  <el-form ref="elForm" :model="formBasicData" label-position="top" label-width="150px">
                    <el-form-item label="Tiêu đề văn bản">
                      <el-input
                        v-model="formBasicData.title"
                        :style="{ width: '100%' }"
                        clearable
                        placeholder="Nhập tiêu đề văn bản"
                      />
                    </el-form-item>
                    <el-row align="bottom" justify="space-between" type="flex">
                      <el-col class="el-col-18 el-col-lg-20">
                        <el-form-item label="Cơ quan ban hành văn bản">
                          <el-select
                            v-model="formBasicData.agency"
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
                            v-model="formBasicData.category"
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
                          v-model="formBasicData.date_issued"
                          :style="{ width: '100%' }"
                          clearable
                          format="DD/MM/YYYY"
                          type="date"
                          @change="handleOnIssuedDateChange"
                        />
                      </el-form-item>
                      <el-form-item class="el-col el-col-lg-8" label="Ngày có hiệu lực">
                        <el-date-picker
                          v-model="formBasicData.date_effected"
                          :style="{ width: '100%' }"
                          clearable
                          format="DD/MM/YYYY"
                          type="date"
                        />
                      </el-form-item>
                      <el-form-item class="el-col el-col-lg-8" label="Ngày hết hiệu lực">
                        <el-date-picker
                          v-model="formBasicData.date_expiration"
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
                          v-model="formBasicData.signNumber"
                          :style="{ width: '100%' }"
                          clearable
                          placeholder="Số văn bản"
                        />
                      </el-form-item>
                      <el-form-item class="el-col el-col-sm-12 el-col-lg-6" label="Năm">
                        <el-input
                          v-model="formBasicData.signYear"
                          :style="{ width: '100%' }"
                          clearable
                          placeholder="Năm phát hành"
                        />
                      </el-form-item>
                      <el-form-item class="el-col el-col-sm-12 el-col-lg-6" label="Thể loại">
                        <el-input
                          v-model="formBasicData.categoryReadonly"
                          :style="{ width: '100%' }"
                          placeholder="Chọn phía trên"
                          readonly
                        />
                      </el-form-item>
                      <el-form-item class="el-col el-col-sm-12 el-col-lg-6" label="Cơ quan">
                        <el-input
                          v-model="formBasicData.agencyReadonly"
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
                            v-model="formBasicData.fields"
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
                        v-model="formBasicData.expert"
                        :style="{ width: '100%' }"
                        clearable
                        placeholder="Nhập mô tả ngắn nếu có"
                        rows="5"
                        type="textarea"
                      />
                    </el-form-item>
                    <el-form-item label="Nội dung">
                      <el-input
                        v-model="formBasicData.content"
                        :style="{ width: '100%' }"
                        clearable
                        placeholder="Nhập nội dung văn bản nếu có"
                        rows="10"
                        type="textarea"
                      />
                    </el-form-item>
                    <el-form-item label="Danh sách cá nhân kí văn bản">
                      <el-select
                        v-model="formBasicData.signers"
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
                        v-model="formBasicData.priorityLevel"
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
                        v-model="formBasicData.status"
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
                      <el-button type="primary" @click="updateBasicDocument">Cập nhật thông tin cơ bản</el-button>
                    </el-form-item>
                  </el-form>
                </div>
              </div>
            </div>

            <div class="gva-card-box" style="margin: 1rem 0">
              <div class="el-card is-always-shadow gva-card quick-entrance">
                <div class="el-card__body">
                  <div class="el-card__heading">
                    <h3>Cập nhật thông tin liên quan</h3>
                  </div>
                  <el-form ref="elForm" :model="formRelatedData" label-position="top" label-width="150px">
                    <el-row align="bottom" justify="space-between" type="flex">
                      <el-col class="el-col-18 el-col-lg-20">
                        <el-form-item label="Văn bản căn cứ">
                          <el-select
                            v-model="formRelatedData.baseDocuments"
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
                            v-model="formRelatedData.relatedDocuments"
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
                            @click="() => openDocumentDialogForm('relation')"
                          > Thêm Nhanh
                          </el-button>
                        </el-form-item>
                      </el-col>
                    </el-row>
                    <el-form-item label="Phòng ban liên quan">
                      <el-select
                        v-model="formRelatedData.relatedAgencies"
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
                        v-model="formRelatedData.relatedUsers"
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

                    <el-form-item size="large">
                      <el-button type="primary" @click="updateDocumentRelation">Cập nhật thông tin liên quan</el-button>
                    </el-form-item>
                  </el-form>
                </div>
              </div>
            </div>

            <div v-if="formAuthorityData.hasValue" class="gva-card-box" style="margin: 1rem 0">
              <div class="el-card is-always-shadow gva-card quick-entrance">
                <div class="el-card__body">
                  <div class="el-card__heading">
                    <h3>Thông tin phân quyền</h3>
                  </div>
                  <el-form ref="elForm" :model="formAuthorityData" label-position="top" label-width="150px">
                    <el-form-item label="Xem văn bản">
                      <el-radio-group v-model="formAuthorityData.view">
                        <el-radio label="all" size="large">Công khai</el-radio>
                        <el-radio label="admin" size="large">Chỉ người dùng quản trị</el-radio>
                        <el-radio label="limit" size="large">Hạn chế</el-radio>
                      </el-radio-group>
                    </el-form-item>
                    <el-form-item v-if="formAuthorityData.view === 'limit'" label="Cá nhân xem văn bản">
                      <el-select
                        v-model="formAuthorityData.viewUsers"
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
                    <el-form-item v-if="formAuthorityData.view === 'limit'" label="Nhóm xem văn bản">
                      <el-select
                        v-model="formAuthorityData.viewRoles"
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
                      <el-radio-group v-model="formAuthorityData.download">
                        <el-radio label="all" size="large">Công khai</el-radio>
                        <el-radio label="admin" size="large">Chỉ người dùng quản trị</el-radio>
                        <el-radio label="limit" size="large">Hạn chế</el-radio>
                      </el-radio-group>
                    </el-form-item>
                    <el-form-item v-if="formAuthorityData.download === 'limit'" label="Cá nhân tải tập tin">
                      <el-select
                        v-model="formAuthorityData.downloadUsers"
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
                    <el-form-item v-if="formAuthorityData.download === 'limit'" label="Nhóm tải tập tin">
                      <el-select
                        v-model="formAuthorityData.downloadRoles"
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
                      <el-radio-group v-model="formAuthorityData.edit">
                        <el-radio label="only" size="large">Chỉ tôi</el-radio>
                        <el-radio label="limit" size="large">Hạn chế</el-radio>
                      </el-radio-group>
                    </el-form-item>
                    <el-form-item v-if="formAuthorityData.edit === 'limit'" label="Cá nhân sửa tập tin">
                      <el-select
                        v-model="formAuthorityData.editUsers"
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
                    <el-form-item v-if="formAuthorityData.edit === 'limit'" label="Nhóm sửa tập tin">
                      <el-select
                        v-model="formAuthorityData.editRoles"
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
                      <el-radio-group v-model="formAuthorityData.owner">
                        <el-radio label="only" size="large">Chỉ tôi</el-radio>
                        <el-radio label="limit" size="large">Hạn chế</el-radio>
                      </el-radio-group>
                    </el-form-item>
                    <el-form-item v-if="formAuthorityData.owner === 'limit'" label="Cá nhân sở hữu tập tin">
                      <el-select
                        v-model="formAuthorityData.ownerUsers"
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
                    <el-form-item v-if="formAuthorityData.edit === 'limit'" label="Nhóm sở hữu tập tin">
                      <el-select
                        v-model="formAuthorityData.ownerRoles"
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

                    <el-form-item size="large">
                      <el-button type="info" @click="updateDocumentAuthority">Khôi phục như lúc đầu</el-button>
                      <el-button type="primary" @click="updateDocumentAuthority">Cập nhật thông tin phân quyền</el-button>
                    </el-form-item>

                  </el-form>
                </div>
              </div>
            </div>

          </el-col>
          <el-col :lg="7" :sm="24">
            <div class="gva-card-box">2</div>

            <div class="gva-card-box" style="margin-top: 1rem;">3</div>
          </el-col>
        </el-row>
      </el-col>
    </el-row>

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
      <el-form :model="documentFormData" label-position="right" label-width="120px">
        <el-form-item label="Tiêu đề">
          <el-input v-model="documentFormData.title" clearable placeholder="Tiêu đề văn bản" />
        </el-form-item>
        <el-form-item label="Liên quan">
          <el-select
            v-model="documentFormData.relatedUsers"
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
          <el-button size="small" @click="closeDocumentDialog">Đóng</el-button>
          <el-button size="small" type="primary" @click="enterDocumentDialog">Thêm nhanh</el-button>
        </div>
      </template>
    </el-dialog>

  </div>
</template>

<script>
export default {
  name: 'DocumentDetail',
}
</script>

<script setup>
import { ref, watch } from 'vue'
import { useRoute } from 'vue-router'

import { useUserStore } from '@/pinia/modules/user'
import { getDict } from '../../utils/dictionary'
import { createDocumentAgencies, getDocumentAgenciesList } from '../../api/documentAgencies'
import { createDocumentCategories, getDocumentCategoriesList } from '../../api/documentCategories'
import { createDocumentFields, getDocumentFieldsList } from '../../api/documentFields'
import { findDocuments, createDraftDocument, createFullDocument, getDocumentsList } from '../../api/documents'
import { getUserList } from '../../api/user'
import { getAuthorityInfo } from '../../api/authority'
import { formatDate } from '../../utils/format'

const route = useRoute()
const userStore = useUserStore()

// ====================== init ref ========================

const formBasicData = ref({
  title: '',
  agency: null,
  category: null,
  date_issued: null,
  date_effected: null,
  date_expiration: null,
  signNumber: null,
  signYear: null,
  categoryReadonly: null,
  agencyReadonly: null,
  fields: [],
  expert: null,
  content: null,
  signers: [],
  status: null,
  priorityLevel: null
})

const formRelatedData = ref({
  baseDocuments: [],
  relatedDocuments: [],
  relatedAgencies: [],
  relatedUsers: [],
})

const formOwnerData = ref({
  updatedBy: userStore.userInfo.ID,
  beResponsibleBy: userStore.userInfo.ID
})

const formAuthorityData = ref({
  hasValue: false,
  view: 'admin',
  download: 'admin',
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

const document = ref(null)
const createdUserName = ref('')
const createdTime = ref('')

const statusOptions = ref([])
const agencyOptions = ref([])
const priorityLevelOptions = ref([])
const usersOptions = ref([])
const documentsOptions = ref([])
const fieldsOptions = ref([])
const categoryOptions = ref([])
const agencyLevelOptions = ref([])
const roleOptions = ref([])

const searchInfo = ref({
  ID: Number(route.params.id)
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
  relatedUsers: []
})

// ====================== end of init ref section =========

// ================= reactive to router params changed ====

watch(() => route.params.id, (id) => {
  searchInfo.value.ID = Number(id)
  getDocument()
})

// ========== end of reactive to router params changed ====

// ========== prepare data section ========================

const getDocument = async() => {
  if (!Number.isInteger(searchInfo.value.ID)) { return }

  const body = await findDocuments({
    ...searchInfo.value,
    preloadFields: 1,
    preloadSigners: 1,
    preloadBasedDocs: 1,
    preloadRelatedDocs: 1,
    preloadRelatedUsers: 1,
    preloadRelatedAgencies: 1,
    preloadCreatedBy: 1,
    preloadAuthority: 1
  })

  if (body.code === 0) {
    document.value = body.data.document
    const d = body.data.document

    formBasicData.value = {
      ...formBasicData.value,
      title: d.title,
      agency: d.agencyId,
      category: d.categoryId,
      date_issued: d.dateIssued && new Date(d.dateIssued),
      date_effected: d.effectDate && new Date(d.effectDate),
      date_expiration: d.dateExpiration && new Date(d.dateExpiration),
      signNumber: d.signNumber,
      signYear: d.signYear,
      categoryReadonly: d.signCategory,
      agencyReadonly: d.signAgency,
      fields: [...d.fields.map(f => f.ID)],
      signers: [...d.signers.map(f => f.ID)],
      expert: d.expert,
      content: d.content,
      status: d.status,
      priorityLevel: d.priority
    }

    formRelatedData.value = {
      ...formRelatedData.value,
      baseDocuments: d.basedDocuments.map(f => f.ID),
      relatedDocuments: d.relatedDocuments.map(f => f.ID),
      relatedUsers: d.relatedUsers.map(f => f.ID),
      relatedAgencies: d.relatedAgencies.map(f => f.ID)
    }

    createdUserName.value = d.createdUser.nickName
    createdTime.value = formatDate(d.CreatedAt)

    // authority
    if (d.authority) {
      formAuthorityData.value = {
        ...formAuthorityData.value,
        hasValue: true,
        viewUsers: d.authority.viewLimitUsers.map(f => f.ID),
        viewRoles: d.authority.viewLimitRoles.map(f => f.authorityId),
        downloadUsers: d.authority.downloadLimitUsers.map(f => f.ID),
        downloadRoles: d.authority.downloadLimitRoles.map(f => f.authorityId),
        editUsers: d.authority.updateLimitUsers.map(f => f.ID),
        editRoles: d.authority.updateLimitRoles.map(f => f.authorityId),
        ownerUsers: d.authority.ownerLimitUsers.map(f => f.ID),
        ownerRoles: d.authority.ownerLimitRoles.map(f => f.authorityId)
      }

      if (d.authority.publicToView) {
        formAuthorityData.value.view = 'all'
      } else {
        if (d.authority.onlyAdminCanView) {
          formAuthorityData.value.view = 'admin'
        } else {
          formAuthorityData.value.view = 'limit'
        }
      }

      if (d.authority.publicToDownload) {
        formAuthorityData.value.download = 'all'
      } else {
        if (d.authority.onlyAdminCanDownload) {
          formAuthorityData.value.download = 'admin'
        } else {
          formAuthorityData.value.download = 'limit'
        }
      }
    }
  }
}

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
  const table = await getAuthorityInfo({ authorityId: 100 })

  if (table.code === 0) {
    roleOptions.value = table.data.authority.children
  }
}

getDocument()
loadAgencyOptions()
loadAgencyLevelOptions()
loadCategoryOptions()
loadFieldOptions()
loadDocumentOptions()
loadUserOptions()
loadRoleOptions()
loadPriorityOptions()
loadStatusOptions()

// ========== end of prepare data section ==================

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

// ================= End of reactive section =================

// ================= Business section =================

const updateBasicDocument = () => {

}

const updateDocumentRelation = () => {

}

const updateDocumentAuthority = () => {

}

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
  const response = await createDraftDocument(documentFormData.value)

  if (response.code === 0) {
    ElMessage({
      type: 'success',
      message: 'Thêm nhanh văn bản thành công'
    })

    documentFormData.value.title = ''
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

// ================= End of business section =================
</script>

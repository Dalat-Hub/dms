<template>
  <div id="search" style="margin-bottom: 2rem">
    <el-row :gutter="10" style="margin-bottom: 1rem">
      <el-col :sm="18">
        <el-row>
          <el-col :span="20">
            <el-input
              class="w-full"
              placeholder="Nhập từ khóa tìm kiếm"
              clearable
              v-model="this.searchTerm"
            />
          </el-col>
          <el-col :span="4">
            <el-button
              type="success"
              class="w-full"
              @click="$emit('onSimpleSearchSubmit', this.searchTerm)"
              >Tìm</el-button
            >
          </el-col>
        </el-row>
      </el-col>

      <el-col :sm="6">
        <el-button
          type="info"
          class="w-full"
          @click="
            () => (this.openAdvancedSearchBox = !this.openAdvancedSearchBox)
          "
          >Nâng cao</el-button
        >
      </el-col>
    </el-row>

    <el-card v-if="this.openAdvancedSearchBox">
      <el-form label-width="180px" :model="this.form">
        <el-form-item label="Chọn cơ quan ban hành">
          <el-select
            placeholder="cơ quan ban hành"
            filterable
            clearable
            class="w-full"
            v-model="this.form.agency"
          >
            <el-option
              v-for="agency in agencies"
              :label="agency.name"
              :value="agency.ID"
              :key="agency.ID"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="Chọn loại văn bản">
          <el-select
            placeholder="thể loại văn bản"
            filterable
            clearable
            class="w-full"
            v-model="this.form.category"
          >
            <el-option
              v-for="category in categories"
              :label="category.name"
              :value="category.ID"
              :key="category.ID"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="Chọn lĩnh vực văn bản">
          <el-select
            placeholder="lĩnh vực văn bản"
            filterable
            clearable
            class="w-full"
            v-model="this.form.field"
          >
            <el-option
              v-for="field in fields"
              :label="field.name"
              :value="field.ID"
              :key="field.ID"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="Ngày ban hành">
          <el-col :span="11">
            <el-date-picker
              type="date"
              placeholder="Từ ngày"
              class="w-full"
              v-model="form.fromDate"
              format="DD-MM-YYYY"
              @change="this.handleOnFromDateChange"
            />
          </el-col>
          <el-col :span="2" class="text-center">
            <span class="text-gray-500">-</span>
          </el-col>
          <el-col :span="11">
            <el-date-picker
              type="date"
              placeholder="Đến ngày"
              class="w-full"
              format="DD-MM-YYYY"
              v-model="form.toDate"
            />
          </el-col>
        </el-form-item>
        <el-form-item label="Tình trạng hiệu lực">
          <el-radio-group v-model="form.stillValid">
            <el-radio :label="true">Còn hiệu lực</el-radio>
            <el-radio :label="false">Hết hiệu lực</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item>
          <el-button
            type="success"
            @click="$emit('onAdvancedSearchSubmit', this.form)"
            >Tìm kiếm nâng cao</el-button
          >
          <el-button @click="this.form = {}">Đặt lại bộ lọc</el-button>
          <el-button @click="this.openAdvancedSearchBox = false"
            >Đóng</el-button
          >
        </el-form-item>
      </el-form>
    </el-card>
  </div>
</template>

<script>
export default {
  name: "SearchBox",
  props: ["agencies", "categories", "fields"],
  emits: ["onSimpleSearchSubmit", "onAdvancedSearchSubmit"],
  data() {
    return {
      openAdvancedSearchBox: false,
      form: {},
      searchTerm: "",
    };
  },
  methods: {
    handleOnFromDateChange(date) {
      const dateObj = new Date(date);

      const dd = dateObj.getDate();
      const mm = dateObj.getMonth();
      const yy = dateObj.getFullYear();

      this.form.toDate = new Date(yy, mm, dd);
    },
  },
};
</script>

<style scoped>
.w-full {
  width: 100%;
}
</style>

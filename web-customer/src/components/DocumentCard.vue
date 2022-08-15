<template>
  <el-card class="box-card" style="margin-bottom: 1rem">
    <template #header>
      <div class="card-header">
        <h3
          @click="$emit('onViewDetailClick', document)"
          class="cursor-pointer"
        >
          {{ this.getDocumentTitle(document) }}
        </h3>
      </div>
    </template>
    <div>
      <p v-html="document?.expert || '--'"></p>
      <p>
        Ngày ban hành:
        {{ this.getIssuedDate(document?.dateIssued) }}
      </p>
      <p>
        Tình trạng hiệu lực:
        <el-tag v-if="document.stillInEffect" class="ml-2" type="success"
          >Còn hiệu lực</el-tag
        >
        <el-tag v-else class="ml-2" type="danger">Hết hiệu lực</el-tag>
      </p>
      <el-popover placement="bottom" :width="500" trigger="click">
        <template #reference>
          <el-button type="info">Thuộc tính</el-button>
        </template>
        <el-table :data="this.getDocumentProps(document)">
          <el-table-column width="150" property="prop" label="Thuộc tính" />
          <el-table-column width="350" property="value" label="Giá trị" />
        </el-table>
      </el-popover>
      <el-button type="success" @click="$emit('onViewDetailClick', document)"
        >Xem chi tiết</el-button
      >
    </div>
  </el-card>
</template>

<script>
import { getDateFormatted } from "@/utils/date";

export default {
  name: "DocumentCard",
  props: ["document"],
  emits: ["onViewDetailClick"],
  methods: {
    getDocumentTitle(document) {
      return `${document.category.name} ${document.signText}`;
    },
    getIssuedDate(date) {
      return getDateFormatted(date);
    },
    getDocumentProps(document) {
      function getSignersAsString() {
        if (Array.isArray(document.signers) && document.signers.length > 0) {
          return document.signers.map((s) => s.nickName).join(", ");
        }

        return "--";
      }

      const template = [
        {
          prop: "Số ký hiệu",
          value: document?.signText || "--",
        },
        {
          prop: "Loại văn bản",
          value: document?.category?.name || "--",
        },
        {
          prop: "Cơ quan ban hành",
          value: document?.agency?.name || "--",
        },
        {
          prop: "Người kí",
          value: getSignersAsString(),
        },
        {
          prop: "Ngày ban hành",
          value: getDateFormatted(document?.dateIssued),
        },
        {
          prop: "Ngày có hiệu lực",
          value: getDateFormatted(document?.effectDate),
        },
        {
          prop: "Ngày hết hiệu lực",
          value: getDateFormatted(document?.expirationDate),
        },
        {
          prop: "Tình trạng",
          value: document?.stillInEffect ? "Còn hiệu lực" : "Hết hiệu lực",
        },
      ];

      return template;
    },
  },
};
</script>

<style scoped>
.cursor-pointer {
  cursor: pointer;
}
</style>

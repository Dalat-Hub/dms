<template>
  <div class="document-view">
    <BreadCrumb />

    <el-row :gutter="10">
      <el-col :sm="8" :lg="6">
        <SideMenu
          :title="this.agency.title"
          :items="this.agency.items"
          param="co-quan-ban-hanh"
        />
        <SideMenu
          :title="this.category.title"
          :items="this.category.items"
          param="the-loai"
        />
        <SideMenu
          :title="this.field.title"
          :items="this.field.items"
          param="linh-vuc"
        />
      </el-col>
      <el-col :sm="16" :lg="18">
        <el-card class="box-card margin-bottom-1-rem card-info">
          <h2 class="card-title">Tiêu đề văn bản</h2>
          <br />
          <p>{{ this.document?.title || "--" }}</p>
        </el-card>
        <el-card class="box-card margin-bottom-1-rem card-info">
          <h2 class="card-title">Mô tả ngắn</h2>
          <br />
          <p>{{ this.document?.expert || "--" }}</p>
        </el-card>
        <div class="grid-content bg-purple margin-bottom-1-rem">
          <el-card class="box-card margin-bottom-1-rem">
            <el-descriptions
              :column="2"
              border
              direction="horizontal"
              title="Thuộc tính của văn bản"
            >
              <el-descriptions-item label="Số ký hiệu">{{
                document?.signText || "--"
              }}</el-descriptions-item>
              <el-descriptions-item label="Ngày ban hành">{{
                getDateFormat(document?.dateIssued)
              }}</el-descriptions-item>

              <el-descriptions-item label="Loại văn bản">{{
                document?.category?.name || "--"
              }}</el-descriptions-item>
              <el-descriptions-item label="Lĩnh vực văn bản">
                {{ document?.fields.map((f) => f.name).join(", ") }}
              </el-descriptions-item>

              <el-descriptions-item label="Cơ quan ban hành">{{
                document?.agency?.name || "--"
              }}</el-descriptions-item>
              <el-descriptions-item label="Người ký">
                {{ document?.signers.map((s) => s.nickName).join(", ") }}
              </el-descriptions-item>

              <el-descriptions-item label="Ngày có hiệu lực">{{
                getDateFormat(document?.effectDate)
              }}</el-descriptions-item>
              <el-descriptions-item label="Ngày hết hiệu lực">{{
                getDateFormat(document?.expirationDate)
              }}</el-descriptions-item>

              <el-descriptions-item label="Tình trạng hiệu lực" span="2">
                <el-tag
                  v-if="document?.stillInEffect"
                  class="ml-2"
                  type="success"
                  >Còn hiệu lực</el-tag
                >
                <el-tag v-else class="ml-2" type="danger">Hết hiệu lực</el-tag>
              </el-descriptions-item>

              <el-descriptions-item label="Tài liệu căn cứ" span="2">
                {{ relations?.base_on?.map((d) => d.title).join(", ") || "--" }}
              </el-descriptions-item>

              <el-descriptions-item label="Tài liệu liên quan" span="2">
                {{
                  relations?.related_docs?.map((d) => d.title).join(", ") ||
                  "--"
                }}
              </el-descriptions-item>
            </el-descriptions>
          </el-card>
          <el-card class="box-card margin-bottom-1-rem card-info">
            <h2 class="card-title">Nội dung</h2>
            <br />
            <p>{{ document?.content || "--" }}</p>
          </el-card>

          <el-card
            v-if="document?.file?.url"
            class="box-card margin-bottom-1-rem card-info"
          >
            <h2 class="card-title">Tập tin đính kèm</h2>
            <br />
            <iframe
              :src="this.getFileURL()"
              style="width: 100%; height: 700px"
            />
          </el-card>

          <el-card
            class="box-card margin-bottom-1-rem card-info"
            style="margin-bottom: 1rem"
            v-if="this.baseDocuments.items.length > 0"
          >
            <template #header>
              <div class="card-header">
                <h2 class="card-title">Văn bản căn cứ</h2>
              </div>
            </template>
            <div
              v-for="doc in this.baseDocuments.items"
              :key="doc.ID"
              class="text item"
            >
              <p>
                <router-link :to="`/van-ban/${doc.ID}`">{{
                  doc.title
                }}</router-link>
              </p>
            </div>
          </el-card>
          <el-card
            class="box-card margin-bottom-1-rem card-info"
            style="margin-bottom: 1rem"
            v-if="this.relatedDocuments.items.length > 0"
          >
            <template #header>
              <div class="card-header">
                <h2 class="card-title">Văn bản liên quan</h2>
              </div>
            </template>
            <div
              v-for="doc in this.relatedDocuments.items"
              :key="doc.ID"
              class="text item"
            >
              <p>
                <router-link :to="`/van-ban/${doc.ID}`">{{
                  doc.title
                }}</router-link>
              </p>
            </div>
          </el-card>
        </div>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import BreadCrumb from "@/components/BreadCrumb";
import SideMenu from "@/components/SideMenu";

import { getDateFormatted } from "@/utils/date";

import { getDocumentCategoryList } from "@/api/category";
import { getDocumentFieldList } from "@/api/field";
import { getDocumentAgencyList } from "@/api/agency";
import { findDocument, getDocumentRelations } from "@/api/document";

export default {
  name: "DocumentView",
  components: {
    BreadCrumb,
    SideMenu,
  },
  data() {
    return {
      agency: {
        title: "Cơ quan ban hành",
        items: [],
      },
      category: {
        title: "Thể loại",
        items: [],
      },
      field: {
        title: "Lĩnh vực",
        items: [],
      },
      document: null,
      baseDocuments: {
        title: "Văn bản căn cứ",
        items: [],
      },
      relatedDocuments: {
        title: "Văn bản liên quan",
        items: [],
      },
    };
  },
  created() {
    this.$watch(
      () => this.$route.params,
      () => {
        this.getDocument();
        this.getDocumentRelation();
      }
    );
  },
  mounted() {
    this.getCategories();
    this.getFields();
    this.getAgencies();
    this.getDocument();
    this.getDocumentRelation();
  },
  methods: {
    getFileURL() {
      return (
        process.env.VUE_APP_BASE_URL +
        "/" +
        this.document.file.url +
        "#toolbar=0"
      );
    },
    getDateFormat(date) {
      return getDateFormatted(date);
    },
    async getAgencies() {
      const table = await getDocumentAgencyList({ page: 1, pageSize: 1000 });

      if (table.data.code === 0) {
        this.agency.items = table.data.data.list.map((item) => {
          return {
            ...item,
            link: "/",
          };
        });
      }
    },
    async getCategories() {
      const table = await getDocumentCategoryList({ page: 1, pageSize: 1000 });

      if (table.data.code === 0) {
        this.category.items = table.data.data.list.map((item) => {
          return {
            ...item,
            link: "/",
          };
        });
      }
    },
    async getFields() {
      const table = await getDocumentFieldList({ page: 1, pageSize: 1000 });

      if (table.data.code === 0) {
        this.field.items = table.data.data.list.map((item) => {
          return {
            ...item,
            link: "/",
          };
        });
      }
    },
    async getDocument() {
      const res = await findDocument({ ID: this.$route.params.id });

      if (res.data.code === 0) {
        this.document = res.data.data.redocuments;
      }
    },
    async getDocumentRelation() {
      const res = await getDocumentRelations({ ID: this.$route.params.id });

      if (res.data.code === 0) {
        this.baseDocuments.items = res.data.data.list.base_on || [];
        this.relatedDocuments.items = res.data.data.list.related_docs || [];
      }
    },
  },
};
</script>

<style scoped>
.box-card {
  width: 100%;
}
.w-full {
  width: 100%;
}
.margin-bottom-1-rem {
  margin-bottom: 1rem;
}
.card-info {
  text-align: left;
}
.card-title {
  text-align: left;
  font-size: 1rem;
}
</style>

<template>
  <div class="document-view">
    <el-divider></el-divider>
    <BreadCrumb />
    <el-divider>
    </el-divider>
    <el-row :gutter="30">
      <el-col :sm="14" :lg="16">
        <div class="document-container">
          <h3 class="document-title">{{ this.document?.title || "--" }}</h3>
        </div>
        <el-card class="margin-bottom-1-rem card-info" shadow="never">
          <h2 class="card-title">Mô tả</h2>
          <div class="document-expert">
            <p>{{ this.document?.expert || "--" }}</p>
          </div>
        </el-card>
        <el-tabs type="border-card" class="margin-bottom-1-rem">
          <el-tab-pane label="Nội dung văn bản">
            <div class="document-content">
              <p>{{ document?.content || "--" }}</p>
            </div>
          </el-tab-pane>
          <el-tab-pane label="Tập tin văn bản">

            <el-card shadow="never" class="margin-bottom-1-rem card-info" style="margin-bottom: 1rem"
              v-if="this.file.files.length > 0">
              <template #header>
                <div class="card-header">
                  <h2 class="card-title">{{ this.file.title }}</h2>
                  <el-button type="success" @click="this.handleFileDownload">Tải tập tin</el-button>
                </div>
              </template>
              <div>
                <iframe :src="this.file.src" frameborder="0" style="width: 100%; height: 700px"></iframe>
              </div>
            </el-card>
          </el-tab-pane>
        </el-tabs>
        <div class="grid-content bg-purple margin-bottom-1-rem">
          <el-card shadow="never" v-if="document?.file?.url" class="box-card margin-bottom-1-rem card-info">
            <h2 class="card-title">Tập tin đính kèm</h2>
            <br />
            <iframe :src="this.getFileURL()" style="width: 100%; height: 700px" />
          </el-card>
          <el-card shadow="never" class="box-card margin-bottom-1-rem card-info" style="margin-bottom: 1rem"
            v-if="this.baseDocuments.items.length > 0">
            <template #header>
              <div class="card-header">
                <h2 class="card-title">Văn bản căn cứ</h2>
              </div>
            </template>
            <div v-for="(doc, index) in this.baseDocuments.items" :key="doc.ID" class="text item">
              <p>
                <router-link :to="`/van-ban/${doc.ID}`">{{ index + 1 }}. {{ doc.title }}</router-link>
              </p>
            </div>
          </el-card>
          <el-card shadow="never" class="box-card margin-bottom-1-rem card-info" style="margin-bottom: 1rem"
            v-if="this.relatedDocuments.items.length > 0">
            <template #header>
              <div class="card-header">
                <h2 class="card-title">Văn bản liên quan</h2>
              </div>
            </template>
            <div v-for="doc in this.relatedDocuments.items" :key="doc.ID" class="text item">
              <p>
                <router-link :to="`/van-ban/${doc.ID}`">{{
                    doc.title
                }}</router-link>
              </p>
            </div>
          </el-card>

        </div>
      </el-col>
      <el-col :sm="10" :lg="8">
        <el-card shadow="never" class="box-card margin-bottom-1-rem">
          <el-descriptions :column="1" border direction="horizontal" title="Thuộc tính của văn bản">
            <el-descriptions-item label="Tình trạng hiệu lực" span="2">
              <el-tag v-if="document?.stillInEffect" class="ml-2" type="success">Còn hiệu lực</el-tag>
              <el-tag v-else class="ml-2" type="danger">Hết hiệu lực</el-tag>
            </el-descriptions-item>
            <el-descriptions-item label="Số ký hiệu">{{ document?.signText || "--" }}</el-descriptions-item>
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
                document?.agencies?.map((e) => e.name).join(", ") || "--"
            }}</el-descriptions-item>
            <el-descriptions-item label="Người ký">
              {{ document?.signerText.split(",").join(", ") }}
            </el-descriptions-item>

            <el-descriptions-item label="Ngày có hiệu lực">{{
                getDateFormat(document?.effectDate)
            }}</el-descriptions-item>
            <el-descriptions-item label="Ngày hết hiệu lực">{{
                getDateFormat(document?.expirationDate)
            }}</el-descriptions-item>

          </el-descriptions>
        </el-card>
        <AgencyTree title="Đơn vị & Thể loại" :tree="this.agencyTree" @onNodeClick="handleOnTreeNodeClick" />
        <SideMenu :title="this.agency.title" :items="this.agency.items" param="co-quan-ban-hanh" />
        <SideMenu :title="this.category.title" :items="this.category.items" param="the-loai" />
        <SideMenu :title="this.field.title" :items="this.field.items" param="linh-vuc" />
      </el-col>
    </el-row>
  </div>
</template>

<script>
import BreadCrumb from "@/components/BreadCrumb";
import SideMenu from "@/components/SideMenu";
import AgencyTree from "@/components/AgencyTree";

import { getDateFormatted } from "@/utils/date";

import { getDocumentCategoryList } from "@/api/category";
import { getDocumentFieldList } from "@/api/field";
import { getDocumentAgencyList } from "@/api/agency";
import { findDocument, getAttachedFiles } from "@/api/document";

import { getAgenciesTree } from "@/api/agency";

const path = process.env.VUE_APP_BASE_URL;

export default {
  name: "DocumentView",
  components: {
    BreadCrumb,
    SideMenu,
    AgencyTree,
  },
  data() {
    return {
      agency: {
        title: "Đơn vị ban hành",
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
      file: {
        title: "Tập tin đính kèm",
        canDownload: false,
        files: [],
        src: "",
      },
      agencyTree: [],
    };
  },
  created() {
    this.$watch(
      () => this.$route.params,
      () => {
        const reloadData = async () => {
          const flag = await this.getDocument();
          if (flag) {
            this.getFiles();
          }
        };

        reloadData();
      }
    );
  },
  mounted() {
    this.getCategories();
    this.getFields();
    this.getAgencies();
    this.getDocument();
    this.getFiles();
    this.getAgencyTree();
  },
  methods: {
    handleOnTreeNodeClick(node) {
      const agencyId = node.agencyId || null;
      const categoryId = node.categoryId || null;

      if (!agencyId || !categoryId) return;

      this.$router.push({
        path: `/`,
        query: {
          "co-quan-ban-hanh": agencyId,
          "the-loai": categoryId,
        },
        replace: true,
      });
    },
    handleOnAgencyFieldNodeClick(node) {
      const agencyId = node.agencyId || null;
      const fieldId = node.fieldId || null;

      if (!agencyId || !fieldId) return;

      this.$router.push({
        path: `/`,
        query: {
          "co-quan-ban-hanh": agencyId,
          "linh-vuc": fieldId,
        },
        replace: true,
      });
    },
    async getAgencyTree() {
      const response = await getAgenciesTree();
      if (response.code === 0) {
        const nodes = response.data.nodes.reduce((acc, cur) => {
          if (!acc[cur.agencyId]) {
            return {
              ...acc,
              [cur.agencyId]: {
                label: cur.agency.name,
                children: [
                  {
                    label: `${cur.category.name} - (${cur.count})`,
                    agencyId: cur.agency.ID,
                    categoryId: cur.category.ID,
                  },
                ],
              },
            };
          }

          return {
            ...acc,
            [cur.agencyId]: {
              ...acc[cur.agencyId],
              children: [
                ...acc[cur.agencyId].children,
                {
                  label: `${cur.category.name} - (${cur.count})`,
                  agencyId: cur.agency.ID,
                  categoryId: cur.category.ID,
                },
              ],
            },
          };
        }, {});

        const tree = Object.keys(nodes).reduce((acc, cur) => {
          return [...acc, nodes[cur]];
        }, []);

        this.agencyTree = tree;
      }
    },
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

      if (table.code === 0) {
        this.agency.items = table.data.list.map((item) => {
          return {
            ...item,
            link: "/",
          };
        });
      }
    },
    async getCategories() {
      const table = await getDocumentCategoryList({ page: 1, pageSize: 1000 });

      if (table.code === 0) {
        this.category.items = table.data.list.map((item) => {
          return {
            ...item,
            link: "/",
          };
        });
      }
    },
    async getFields() {
      const table = await getDocumentFieldList({ page: 1, pageSize: 1000 });

      if (table.code === 0) {
        this.field.items = table.data.list.map((item) => {
          return {
            ...item,
            link: "/",
          };
        });
      }
    },
    async getDocument() {
      const id = this.$route.params.id || 0;

      if (id <= 0) return;

      const res = await findDocument({
        ID: id,
        preloadCategory: 1,
        preloadAgency: 1,
        preloadFields: 1,
        preloadSigners: 1,
        preloadBasedDocs: 1,
        preloadRelatedDocs: 1,
        preloadRelatedUsers: 1,
        preloadRelatedAgencies: 1,
        preloadCreatedBy: 1,
      });

      if (res.code === 0) {
        this.document = res.data.document;
        this.baseDocuments = {
          ...this.baseDocuments,
          items: res.data.document?.basedDocuments || [],
        };
        this.relatedDocuments = {
          ...this.relatedDocuments,
          items: res.data.document?.relatedDocuments || [],
        };

        return true;
      }

      return false;
    },
    async getFiles() {
      const id = this.$route.params.id || 0;

      if (id <= 0) return;

      const res = await getAttachedFiles({
        id: this.$route.params.id,
      });

      if (res.code === 0) {
        this.file.files = res.data.files;
        this.file.canDownload = res.data.canDownload;

        if (res.data.files.length > 0) {
          if (res.data.files[0].url.includes(".pdf")) {
            this.file.src = `${path}/${res.data.files[0].url}#toolbar=0`;
          }
        }
      }
    },
    handleFileDownload() {
      const url = `${path}/${this.file.files[0].url}`;
      window.open(url);
    },
  },
};
</script>

<style>
.document-view {
  margin-left: 5%;
  margin-right: 5%;
}


.document-view .box-card {
  width: 95%;
}

.document-view .w-full {
  width: 100%;
}

.document-view .margin-bottom-1-rem {
  margin-bottom: 1rem;
}

.document-view .card-info {
  text-align: left;
}

.document-view .card-title {
  text-align: left;
  font-size: 1rem;
}

/* .document-container{

} */

.document-title {
  text-align: left;
  text-transform: uppercase;
  line-height: 1.9em;
}

.document-content {
  text-align: left;
  font-size: 16px;
  line-height: 1.6em;
}
</style>

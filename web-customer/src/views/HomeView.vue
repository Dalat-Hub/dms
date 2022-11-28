<template>
  <div class="home-view">
    <el-divider></el-divider>
    <BreadCrumb />
    <el-divider></el-divider>
    <el-row :gutter="16" class="main-col">
      <el-col :lg="5" class="aside">
        <AgencyTree title="Đơn vị & Thể loại" :tree="this.agencyTree" @onNodeClick="handleOnTreeNodeClick" />
        <SideMenu :title="this.agency.title" :items="this.agency.items" param="co-quan-ban-hanh" />
        <SideMenu :title="this.category.title" :items="this.category.items" param="the-loai" />
        <SideMenu :title="this.field.title" :items="this.field.items" param="linh-vuc" />
        <SideContent title="Văn bản mới" :items="latestDocuments" :displayCreatedAt="true"
          className="card-latest-documents"></SideContent>

        <SideContent title="Văn bản xem nhiều" :items="this.mostViewDocuments" :displayCounter="true"
          className="card-most-view-documents"></SideContent>
      </el-col>
      <el-col :lg="19">
        <div class="left-panel">
          <SearchBox :agencies="this.agency.items" :categories="this.category.items" :fields="this.field.items"
            @onSimpleSearchSubmit="this.handleOnSimpleSearchSubmit"
            @onAdvancedSearchSubmit="this.handleOnAdvancedSearchSubmit" />

          <SearchStat :count="this.searchStats.count" :searchBy="this.searchStats.searchBy" :meta="{
            agency: this.agency,
            category: this.category,
            field: this.field,
            signText: this.searchStats.signText,
          }" />
        </div>
        <div id="main">

          <!-- <DocumentCard v-for="document in this.document.items" :key="document.ID" :document="document"
            @onViewDetailClick="handleOnDocumentViewDetailClick" />
          <div style="display: flex; justify-content: center; margin-top: 2rem">
            <el-pagination v-if="parseInt(this.searchStats.count / this.pageSize) > 0" background
              layout="prev, pager, next" :total="this.searchStats.count" :page-size="this.pageSize"
              @current-change="handlePageChanged" />
          </div> -->
          <div>
            <el-auto-resizer>

              <el-table @rowClick="handleOnDocumentViewDetailClick" :data="this.document.items" :default-sort="{ prop: 'date', order: 'descending' }"
                style="width: 100%">
                <el-table-column prop="title" label="Văn bản" sortable width="300" />
                <el-table-column prop="signText" label="Số hiệu" sortable />
                <el-table-column prop="stillInEffect" label="Tình trạng" sortable>
                  <template #default="scope">
                    <el-tag v-if="scope.row.stillInEffect" class="ml-2" type="success">Còn hiệu lực</el-tag>
                    <el-tag v-else class="ml-2" type="danger">Hết hiệu lực</el-tag>
                  </template>
                </el-table-column>
                <el-table-column prop="category.name" label="Loại văn bản" sortable />
                <el-table-column prop="dateIssued" label="Ngày ban hành" sortable>
                  <template #default="scope">
                    {{ getDateFormat(scope.row.dateIssued) }}
                  </template>
                </el-table-column>
                <el-table-column prop="effectDate" label="Ngày hiệu lực" sortable>
                  <template #default="scope">
                    {{ getDateFormat(scope.row.effectDate) }}
                  </template>
                </el-table-column>
                <el-table-column prop="signers[0].fullname" label="Người ký" />
              </el-table>
              <div class="pagination-con">
                <el-pagination v-if="parseInt(this.searchStats.count / this.pageSize) > 0" background
                  layout="prev, pager, next" :total="this.searchStats.count" :page-size="this.pageSize"
                  @current-change="handlePageChanged"></el-pagination>
              </div>
            </el-auto-resizer>
          </div>
        </div>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import BreadCrumb from "@/components/BreadCrumb";
import SearchBox from "@/components/SearchBox";
// import DocumentCard from "@/components/DocumentCard";
import SideMenu from "@/components/SideMenu";
import SideContent from "@/components/SideContent";
import SearchStat from "@/components/SearchStat";
import AgencyTree from "@/components/AgencyTree";
import { ElMessage } from "element-plus";

import { getDocumentCategoryList } from "@/api/category";
import { getDocumentFieldList } from "@/api/field";
import { getAgenciesTree, getDocumentAgencyList } from "@/api/agency";
import { getDocumentList } from "@/api/document";

import { getDateFormatted } from "@/utils/date";

import { TableV2FixedDir, TableV2SortOrder } from 'element-plus'


export default {
  name: "HomeView",
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
      document: {
        title: "Văn bản",
        items: [],
      },
      latestDocuments: [],
      mostViewDocuments: [],
      pageSize: 15,
      searchStats: {
        count: 0,
        searchBy: {},
        signText: "",
      },
      agencyTree: [],
    };
  },
  components: {
    BreadCrumb,
    SearchBox,
    // DocumentCard,
    SideMenu,
    SearchStat,
    SideContent,
    AgencyTree,
  },
  created() {
    this.$watch(
      () => this.$route.params,
      () => {
        const agencyId = this.$route.query["co-quan-ban-hanh"] || null;
        const fieldId = this.$route.query["linh-vuc"] || null;
        const categoryId = this.$route.query["the-loai"] || null;

        const query = {};

        if (agencyId) query.agency = agencyId;
        if (fieldId) query.field = fieldId;
        if (categoryId) query.category = categoryId;

        if (Object.keys(query).length > 0) {
          this.getDocuments(query);
        }
      }
    );
  },
  mounted() {
    this.getCategories();
    this.getFields();
    this.getAgencies();
    this.getDocuments();
    this.getDocumentByDay();
    this.getNewestDocuments();
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
    async getNewestDocuments() {
      const response = await getDocumentList({
        page: 1,
        pageSize: 10,
        preloadCategory: 1,
        preloadAgency: 1,
        preloadFields: 1,
        preloadSigners: 1,
      });

      if (response.code === 0) {
        this.latestDocuments = response.data.list;
      }
    },
    async getDocumentByDay() {
      const response = await getDocumentList({
        page: 1,
        pageSize: 10,
        preloadCategory: 1,
        preloadAgency: 1,
        preloadFields: 1,
        preloadSigners: 1,
        attachAuthority: 1,
        orderBy: "view_count",
      });

      if (response.code === 0) {
        this.mostViewDocuments = response.data.list;
      }
    },
    async getDocumentByFilter(type, id) {
      const validMap = {
        "co-quan-ban-hanh": "agency",
        "the-loai": "category",
        "linh-vuc": "field",
      };

      const filterBy = validMap[type] || null;
      if (!filterBy) return;

      this.getDocuments({
        [filterBy]: id,
      });
    },
    async getAgencies() {
      const table = await getDocumentAgencyList({ page: 1, pageSize: 1000 });

      if (table.code === 0) {
        const agencies = table.data.list.map((item) => {
          return {
            ...item,
            link: "/",
          };
        });

        this.agency.items = agencies;
      }
    },
    async getCategories() {
      const table = await getDocumentCategoryList({ page: 1, pageSize: 1000 });

      if (table.code === 0) {
        const categories = table.data.list.map((item) => {
          return {
            ...item,
            link: "/",
          };
        });

        this.category.items = categories;
      }
    },
    async getFields() {
      const table = await getDocumentFieldList({ page: 1, pageSize: 1000 });

      if (table.code === 0) {
        const fields = table.data.list.map((item) => {
          return {
            ...item,
            link: "/",
          };
        });

        fields.sort((a, b) => b.count - a.count);

        this.field.items = fields;
      }
    },
    async getDocuments(filter = {}) {
      const agencyId = this.$route.query["co-quan-ban-hanh"] || null;
      const fieldId = this.$route.query["linh-vuc"] || null;
      const categoryId = this.$route.query["the-loai"] || null;

      const query = {};

      if (agencyId) query.agency = agencyId;
      if (fieldId) query.field = fieldId;
      if (categoryId) query.category = categoryId;

      const table = await getDocumentList({
        ...query,
        page: 1,
        pageSize: this.pageSize,
        ...filter,
        preloadCategory: 1,
        preloadAgency: 1,
        preloadFields: 1,
        preloadSigners: 1,
        attachAuthority: 1,
      });

      if (table.code === 0) {
        this.document.items = table.data.list;
        console.log(this.document)
        this.searchStats.count = table.data.total;
        this.searchStats.searchBy = filter;
        delete this.searchStats.searchBy.page;

        ElMessage({
          type: "success",
          message: "Lấy danh sách tài liệu thành công!!!",
        });
      }
    },
    handleOnDocumentViewDetailClick(document) {
      // console.log("handleOnDocumentViewDetailClick")
      // console.log(document)
      this.$router.push("/van-ban/" + document.ID);
    },
    async handleOnSimpleSearchSubmit(keyword) {
      this.searchStats.signText = keyword;
      await this.getDocuments({ keyword: keyword });
      this.searchStats.signText = "";
    },
    handleOnAdvancedSearchSubmit(form) {
      let fromDate = "";
      let toDate = "";

      if (form.fromDate instanceof Date) {
        fromDate = form.fromDate.setDate(form.fromDate.getDate() + 1);
        fromDate = form.fromDate.toISOString();
        fromDate = `"${fromDate}"`;
      }

      if (form.toDate instanceof Date) {
        toDate = form.toDate.setDate(form.toDate.getDate() + 1);
        toDate = form.toDate.toISOString();
        toDate = `"${toDate}"`;
      }

      if (fromDate && toDate) this.getDocuments({ ...form, fromDate, toDate });
      else this.getDocuments({ ...form });
    },
    handlePageChanged(page) {
      this.getDocuments({
        ...this.searchStats.searchBy,
        page,
        pageSize: this.pageSize,
      });
    },
    getDateFormat(date) {
      return getDateFormatted(date);
    },
  },
};
</script>

<style scoped>
.home-view {
  margin-left: 32px;
  margin-right: 32px;
}

.box-card {
  width: 100%;
}

.w-full {
  width: 95%;
}

.el-pagination {
  margin-top: 32px;
    width: 100%;
    overflow: auto;
    display: flex;
    justify-content: center;
  }

@media screen and (max-width: 600px) {
  .w-full {
    width: 100%;
  }

  .el-pagination {
    margin-top: 32px !important;
  }

  .home-view {
    margin: 0
  }

  .el-pagination {
    width: 100%;
    overflow: auto;
  }

  .main-col {
    display: flex;
    flex-wrap: wrap-reverse;
  }


  .aside {
    margin-top: 7em;
  }


}
</style>

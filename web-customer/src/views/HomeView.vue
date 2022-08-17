<template>
  <div class="home-view">
    <BreadCrumb />

    <el-row :gutter="16">
      <el-col :lg="6">
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
      <el-col :lg="12">
        <SearchBox
          :agencies="this.agency.items"
          :categories="this.category.items"
          :fields="this.field.items"
          @onSimpleSearchSubmit="this.handleOnSimpleSearchSubmit"
          @onAdvancedSearchSubmit="this.handleOnAdvancedSearchSubmit"
        />

        <SearchStat
          :count="this.searchStats.count"
          :searchBy="this.searchStats.searchBy"
          :meta="{
            agency: this.agency,
            category: this.category,
            field: this.field,
            signText: this.searchStats.signText,
          }"
        />

        <div id="main">
          <DocumentCard
            v-for="document in this.document.items"
            :key="document.ID"
            :document="document"
            @onViewDetailClick="handleOnDocumentViewDetailClick"
          />

          <div style="display: flex; justify-content: center; margin-top: 2rem">
            <el-pagination
              v-if="parseInt(this.searchStats.count / this.pageSize) > 0"
              background
              layout="prev, pager, next"
              :total="this.searchStats.count"
              :page-size="this.pageSize"
              @current-change="handlePageChanged"
            />
          </div>
        </div>
      </el-col>
      <el-col :lg="6">
        <SideContent
          title="Văn bản mới"
          :items="latestDocuments"
          :displayCreatedAt="true"
        ></SideContent>

        <SideContent
          title="Văn bản xem nhiều"
          :items="this.mostViewDocuments"
          :displayCounter="true"
        ></SideContent>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import BreadCrumb from "@/components/BreadCrumb";
import SearchBox from "@/components/SearchBox";
import DocumentCard from "@/components/DocumentCard";
import SideMenu from "@/components/SideMenu";
import SideContent from "@/components/SideContent";
import SearchStat from "@/components/SearchStat";

import { getDocumentCategoryList } from "@/api/category";
import { getDocumentFieldList } from "@/api/field";
import { getDocumentAgencyList } from "@/api/agency";
import { getDocumentList } from "@/api/document";

export default {
  name: "HomeView",
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
      document: {
        title: "Văn bản",
        items: [],
      },
      latestDocuments: [],
      mostViewDocuments: [],
      pageSize: 10,
      searchStats: {
        count: 0,
        searchBy: {},
        signText: "",
      },
    };
  },
  components: {
    BreadCrumb,
    SearchBox,
    DocumentCard,
    SideMenu,
    SearchStat,
    SideContent,
  },
  created() {
    this.$watch(
      () => this.$route.params,
      () => {
        const type = this.$route.query["phan-loai"] || null;
        const id = this.$route.query.id || null;

        if (!type || !id) {
          this.getDocuments();
        } else {
          this.getDocumentByFilter(type, id);
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
  },
  methods: {
    async getNewestDocuments() {
      const response = await getDocumentList({
        page: 1,
        pageSize: 20,
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
        pageSize: 20,
        preloadCategory: 1,
        preloadAgency: 1,
        preloadFields: 1,
        preloadSigners: 1,
        orderBy: "view_count",
      });

      if (response.code === 0) {
        this.mostViewDocuments = response.data.list;
      }
    },
    async getDocumentByFilter(type, id) {
      const validMap = {
        "co-quan-ban-hanh": "agencyId",
        "the-loai": "categoryId",
        "linh-vuc": "fieldId",
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
    async getDocuments(filter = {}) {
      const table = await getDocumentList({
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
        this.searchStats.count = table.data.total;
        this.searchStats.searchBy = filter;
        delete this.searchStats.searchBy.page;
      }
    },
    handleOnDocumentViewDetailClick(document) {
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
</style>

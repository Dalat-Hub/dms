<template>
  <div class="home-view">
    <BreadCrumb />

    <el-row :gutter="16">
      <el-col :sm="7" :lg="6">
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
      <el-col :sm="10" :lg="12">
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
              v-if="this.searchStats.count > 0"
              background
              layout="prev, pager, next"
              :total="this.searchStats.count"
              :page-size="10"
              @current-change="handlePageChanged"
            />
          </div>
        </div>
      </el-col>
      <el-col :sm="7" :lg="6">
        <SideContent
          title="Văn bản mới"
          :items="document.items"
          :displayCreatedAt="true"
        ></SideContent>

        <SideContent
          title="Xem nhiều trong ngày"
          :items="document.items"
          :displayCounter="true"
        ></SideContent>

        <SideContent
          title="Xem nhiều trong tuần"
          :items="document.items"
          :displayCounter="true"
        ></SideContent>

        <SideContent
          title="Xem nhiều trong tháng"
          :items="document.items"
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
  },
  methods: {
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
    async getDocuments(filter = {}) {
      const table = await getDocumentList({
        page: 1,
        pageSize: 10,
        ...filter,
        preloadCategory: 1,
        preloadAgency: 1,
        preloadFields: 1,
        preloadSigners: 1,
      });

      if (table.data.code === 0) {
        this.document.items = table.data.data.list;
        this.searchStats.count = table.data.data.total;
        this.searchStats.searchBy = filter;
        delete this.searchStats.searchBy.page;
      }
    },
    handleOnDocumentViewDetailClick(document) {
      this.$router.push("/van-ban/" + document.ID);
    },
    async handleOnSimpleSearchSubmit(keyword) {
      this.searchStats.signText = keyword;
      await this.getDocuments({ signText: keyword });
      this.searchStats.signText = "";
    },
    handleOnAdvancedSearchSubmit(form) {
      this.getDocuments({ ...form });
    },
    handlePageChanged(page) {
      this.getDocuments({ ...this.searchStats.searchBy, page });
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

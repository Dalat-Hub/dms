import service from "@/utils/request";

export const getDocumentCategoryList = (params) => {
  return service({
    url: "public/documentCategories/getDocumentCategoryList",
    method: "get",
    params,
  });
};

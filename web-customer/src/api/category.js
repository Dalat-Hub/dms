import service from "@/utils/request";

export const getDocumentCategoryList = (params) => {
  return service({
    url: "api/v1/categories",
    method: "get",
    params,
  });
};

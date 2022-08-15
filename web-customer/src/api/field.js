import service from "@/utils/request";

export const getDocumentFieldList = (params) => {
  return service({
    url: "api/v1/fields",
    method: "get",
    params,
  });
};

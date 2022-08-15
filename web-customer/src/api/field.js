import service from "@/utils/request";

export const getDocumentFieldList = (params) => {
  return service({
    url: "public/documentFields/getFieldList",
    method: "get",
    params,
  });
};

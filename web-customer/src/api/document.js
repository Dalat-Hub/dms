import service from "@/utils/request";

export const getDocumentList = (params) => {
  return service({
    url: "api/v1/documents",
    method: "get",
    params,
  });
};

export const findDocument = (params) => {
  return service({
    url: "api/v1/documents/find",
    method: "get",
    params,
  });
};

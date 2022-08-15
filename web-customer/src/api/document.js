import service from "@/utils/request";

export const getDocumentList = (params) => {
  return service({
    url: "public/documents/getDocumentList",
    method: "get",
    params,
  });
};

export const findDocument = (params) => {
  return service({
    url: "public/documents/findDocument",
    method: "get",
    params,
  });
};

export const getDocumentRelations = (params) => {
  return service({
    url: "public/documents/relations",
    method: "get",
    params,
  });
};

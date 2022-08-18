import service from "@/utils/request";

export const getDocumentAgencyList = (params) => {
  return service({
    url: "api/v1/agencies",
    method: "get",
    params,
  });
};

export const getAgenciesTree = (params) => {
  return service({
    url: "api/v1/agencies/tree",
    method: "get",
    params,
  });
};

export const getAgenciesTreeForField = (params) => {
  return service({
    url: "api/v1/agencies/tree/fields",
    method: "get",
    params,
  });
};

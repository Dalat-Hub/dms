import service from "@/utils/request";

export const getDocumentAgencyList = (params) => {
  return service({
    url: "api/v1/agencies",
    method: "get",
    params,
  });
};

import service from "@/utils/request";

export const getDocumentAgencyList = (params) => {
  return service({
    url: "public/documentAgencies/getDocumentAgencyList",
    method: "get",
    params,
  });
};

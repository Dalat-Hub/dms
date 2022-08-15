import service from "@/utils/request";

export const createDocumentRequest = (data) => {
  return service({
    url: "public/documentRequest/createDocumentRequest",
    method: "post",
    data,
  });
};

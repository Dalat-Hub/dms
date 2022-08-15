import service from "@/utils/request";

export const createDocumentContact = (data) => {
  return service({
    url: "public/documentContact/createDocumentContact",
    method: "post",
    data,
  });
};

import service from "@/utils/request";

export const createDocumentPermissionRequest = (data) => {
  return service({
    url: "/api/v1/permission-requests",
    method: "post",
    data,
  });
};

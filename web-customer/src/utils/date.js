import format from "date-fns/format";

import { DATE_FORMAT } from "@/utils/constant";

export const getDateFormatted = (dateString, dateFormat = DATE_FORMAT) => {
  try {
    if (!dateString) return "--";

    const dateObject = new Date(dateString);
    return format(dateObject, dateFormat);
  } catch {
    return "--";
  }
};

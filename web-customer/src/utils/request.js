import axios from "axios";

const service = axios.create({
  baseURL: process.env.VUE_APP_BASE_URL,
  timeout: 99999,
});

export default service;

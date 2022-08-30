import { createApp } from "vue";
import ElementPlus from "element-plus";
import "element-plus/dist/index.css";
import viVN from "element-plus/es/locale/lang/vi";
import App from "./App.vue";
import { store } from "@/pinia";
import router from "./router";

const app = createApp(App);

app.use(store);
app.use(ElementPlus, { locale: viVN });
app.use(router);

app.mount("#app");

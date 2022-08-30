<template>
  <Header />
  <router-view />
  <Footer />
  <el-backtop :right="50" :bottom="50" />
</template>

<script>
import Header from "@/components/Header";
import Footer from "@/components/Footer";

export default {
  name: "App",
  components: {
    Header,
    Footer,
  },
};
</script>

<style>
#app {
  font-family: Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

nav {
  padding: 30px;
}

nav a {
  font-weight: bold;
  color: #2c3e50;
}

nav a.router-link-exact-active {
  color: #42b983;
}

a[class^="router-link"],
.text.item a {
  text-decoration: none;
  color: rgba(44, 62, 80, 0.8);
}
a[class^="router-link"]:visited,
.text.item a:visited {
  color: rgba(44, 62, 80, 0.8);
}
a[class^="router-link"]:hover,
.text.item a:hover {
  color: rgba(44, 62, 80, 1);
}
.el-menu-item a {
  text-decoration: none;
}
</style>

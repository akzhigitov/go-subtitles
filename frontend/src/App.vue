<template>
  <v-app>
    <v-main>
      <subtitle-page></subtitle-page>
    </v-main>
  </v-app>
</template>

<script>
import SubtitlePage from "./pages/subtitle-page.vue";

export default {
  name: "App",

  components: {
    SubtitlePage,
  },

  data: () => ({
    //
  }),
};
</script>

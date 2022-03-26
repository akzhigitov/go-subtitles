<template>
  <app>
    <v-file-input
      :rules="rules"
      accept="image/png, image/jpeg, image/bmp"
      placeholder="Pick an avatar"
      prepend-icon="mdi-camera"
      label="Avatar"
    ></v-file-input>

    <words-list></words-list>
  </app>
</template>

<script>
import App from "../App.vue";
import WordsList from "../components/WordsList.vue";
export default {
  components: { WordsList, App },
  methods: {},
};
</script>

<style>
</style>
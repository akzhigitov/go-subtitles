<template>
  <v-container>
    <v-file-input
      accept=".srt"
      placeholder="Pick a subtitle"
      prepend-icon="mdi-subtitles"
      label="Subtitle"
      @change="fileInputChange"
    ></v-file-input>
    <v-tabs>
      <v-tab>
        <v-icon color="red" left> mdi-book-open </v-icon>
        Unknown {{unknownWords.length}}
      </v-tab>
      <v-tab>
        <v-icon color="green" left> mdi-book-open </v-icon>
          Known {{knownWords.length}}
      </v-tab>
      <v-tab>
        <v-icon left> mdi-book-open </v-icon>
       Broken {{brokenWords.length}}
      </v-tab>

      <v-tab-item>
        <words-list
          :words="unknownWords"
        ></words-list>
      </v-tab-item>
            <v-tab-item>
        <words-list
          :words="knownWords"
        ></words-list>
      </v-tab-item>
            <v-tab-item>
        <words-list
          :words="brokenWords"
        ></words-list>
      </v-tab-item>
    </v-tabs>

    
  </v-container>
</template>

<script>
import axios from "@/plugins/axios";
import WordsList from "../components/WordsList.vue";
export default {
  components: { WordsList },
  data() {
    return {
      unknownWords: [],
      knownWords: [],
      brokenWords: [],
    };
  },
  methods: {
    async fileInputChange(file) {
      console.log(file);
      let formData = new FormData();
      formData.append("file", file);
      const response = await axios.post(
        "http://localhost:1323/upload",
        formData,
        {
          headers: {
            "Content-Type": "multipart/form-data",
          },
        }
      );

      this.unknownWords = response.data.unknownWords;
      this.knownWords = response.data.knownWords;
      this.brokenWords = response.data.brokenWords;
    },
  },
};
</script>

<style>
</style>
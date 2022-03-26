<template>
  <v-app>
    <v-main>
     <words-list/>
    </v-main>
  </v-app>
</template>

<script>
import WordsList from './components/WordsList.vue';

export default {
  name: 'App',

  components: {
    WordsList,
  },

  data: () => ({
    //
  }),
};
</script>

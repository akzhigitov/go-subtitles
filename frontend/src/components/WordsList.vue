<template>
  <v-data-table :headers="headers" :items="words" class="elevation-1" :sort-by="['freq']" :items-per-page=50>
    <template v-slot:[`item.freq`]="{ item }">
      <v-chip :color="getColor(item.freq)"  dark>
        {{ shortNum(item.freq) }}
      </v-chip>
    </template>
      <template v-slot:[`item.value`]="{ item }">
     <div class="balloon-row">  {{ item.value }} </div>
    </template>
      <template v-slot:[`item.phrase`]="{ item }">
     <div class="balloon-row">  {{ item.phrase }} </div>
    </template>
  </v-data-table>
</template>

<script>
const shortNumber = require('number-shortener')
export default {
  data() {
    return {
      headers: [
        {
          text: "Word",
          value: "value",
        },
        { text: "Frequency", value: "freq" },
         { text: "Phrase", value: "phrase", sortable:false },
      ],
    };
  },
  methods: {
    shortNum(num){
        return shortNumber(num)
    },
    getColor(frequency) {
      if (frequency < 200000) return "red";
      else if (frequency < 1000000) return "orange";
      else return "green";
    },
  },
  props: {
    words: Array,
  },
  mount:{

  }
};
</script>

<style>
</style>
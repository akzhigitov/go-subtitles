<template>
  <v-list subheader>
      <v-subheader>Words</v-subheader>

      <v-list-item
        v-for="word in words"
        :key="word.value"
      >

        <v-list-item-content>
          <v-list-item-title v-text="word.value"></v-list-item-title>
        </v-list-item-content>

      </v-list-item>
    </v-list>

</template>

<script>
export default {
    data: () =>({
        words:[
            {
                value: "missing"
            },
            {
                value: "corner"
            }
        ]
    })
}
</script>

<style>

</style>
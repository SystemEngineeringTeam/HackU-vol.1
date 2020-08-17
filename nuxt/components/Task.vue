<template>
  <v-card>
    <v-card-title class="justify-center">
      {{ task.name }}
    </v-card-title>
    <v-card-subtitle>
      {{cardStr}}
    </v-card-subtitle>
    <v-card-text>
      {{task.description}}
    </v-card-text>
    <v-card-actions class="justify-center">
      <v-btn @click="test">更新</v-btn>
      <v-btn @click="test">削除</v-btn>
    </v-card-actions>
  </v-card>
</template>

<script>
export default {
  name: "Task",

  props: ["task"],

  data: () => ({
    weight: ["ぬるい","ふつう","えぐい"]
  }),

  methods: {
    test: function(){

    }
  },

  computed: {
    cardStr(){
      return this.task.deadlineDate + " " + this.task.deadlineTime + " / " + this.weight[this.task.weight];
    }
  }
};
</script>

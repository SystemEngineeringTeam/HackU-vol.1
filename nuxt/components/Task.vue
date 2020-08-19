<template>
  <v-card>
    <v-card-title class="justify-center">
      {{ task.title }}
    </v-card-title>
    <v-card-subtitle>
      {{cardStr}}
    </v-card-subtitle>
    <v-card-text>
      {{task.description}}
    </v-card-text>
    <v-card-actions class="justify-center">
      <v-btn @click="success">達成</v-btn>
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
    success: function(){
      this.$store.dispatch("tasks/successTask",this.id);
    }
  },

  computed: {
    cardStr(){
      return this.task.deadlineDate + " " + this.task.deadlineTime + " / " + this.weight[this.task.weight];
    }
  }
};
</script>

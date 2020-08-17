<template>
  <v-row class="text-center">
    <v-col cols="12" v-for="(task,i) in tasks" :key="i">
      <div class="text-center">
        <Task :task="task" />
      </div>
    </v-col>
  </v-row>
</template>

<script>
import Task from "../components/Task";

export default {
  components: {
    Task
  },

  data: () => ({
    tasks: [
      {
        id: 0,
        name: "ごはん　たべる",
        deadlineDate: "2020-08-21",
        deadlineTime: "12:00:00",
        description: "0",
        weight: "0",
      },
      {
        id: 1,
        name: "うああああ",
        deadlineDate: "2020-8-22",
        deadlineTime: "09:12:00",
        description: "0",
        weight: "1"
      }
    ]
  }),
}
</script>

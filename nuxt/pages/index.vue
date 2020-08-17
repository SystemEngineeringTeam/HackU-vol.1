<template>
  <v-row class="text-center">
    <v-col cols="12" v-for="(task,i) in $store.state.tasks.tasks" :key="i">
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

  data: () => ({}),
}
</script>

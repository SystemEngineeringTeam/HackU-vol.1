<template>
  <v-row class="text-center">
    <v-col cols="12" v-for="(task, i) in $store.state.tasks.tasks" :key="i">
      <div class="text-center">
        <Task :task="task" />
      </div>
    </v-col>
    <postDialog />
  </v-row>
</template>

<script>
import Task from '../components/Task'
import postDialog from '../components/postDialog'

export default {
  components: {
    Task,
    postDialog,
  },

  data: () => ({}),

  methods: {},

  created() {
    this.$store.dispatch('tasks/setTasks')
    //this.$store.dispatch("tasks/postTask");
  },
}
</script>

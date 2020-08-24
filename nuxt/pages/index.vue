<template>
  <v-row class="text-center">
    <v-col cols="12">
      <Logbox />
    </v-col>
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
import Logbox from '../components/Logbox'

export default {
  components: {
    Task,
    postDialog,
    Logbox,
  },

  data: () => ({
    intervalID: null,
  }),

  methods: {
    secondFunc: function () {
      this.$store.dispatch('game/writeDamageLog')
    },
  },

  created() {
    this.$store.dispatch('tasks/setTasks')
    this.$store.dispatch('game/gameInit')
    this.intervalID = setInterval(this.secondFunc, 1000)
  },

  destroyed() {
    clearInterval(this.intervalID)
  },
}
</script>

<style lang="scss">
.v-application {
  font-family: 'PixelMplus';
}
</style>

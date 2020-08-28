<template>
  <v-row class="text-center">
    <v-col cols="12">
      <MonsterView />
    </v-col>
    <v-col cols="12">
      <Logbox />
    </v-col>
    <v-col cols="12" v-for="(task, i) in $store.state.tasks.tasks" :key="i">
      <div class="text-center">
        <Task :task="task" />
      </div>
    </v-col>
    <postDialog />
    <DieDialog />
  </v-row>
</template>

<script>
import Task from '../components/Task'
import postDialog from '../components/postDialog'
import Logbox from '../components/Logbox'
import MonsterView from '../components/Monster/MonsterView'
import DieDialog from '../components/DieDialog'

export default {
  components: {
    Task,
    postDialog,
    Logbox,
    MonsterView,
    DieDialog
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
    if (this.$store.state.user.token === '') {
      this.$router.push('/login')
      return
    }
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

html,body {
    width: 100%;
    margin: 0px;
    padding: 0px;
    overflow-x: hidden;
}
</style>
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
    lowerHP: function () {
      let hp = this.$store.state.user.HP
      hp = Math.max(0, hp - this.$store.state.tasks.tasks.length)
      this.$store.commit('user/setHP', hp)
    },

    writeLog: function () {
      let log = this.$store.state.user.log
      this.$store.state.tasks.tasks.forEach((element) => {
        log =
          element.title +
          'の攻撃！' +
          this.$store.state.user.name +
          'は' +
          1 +
          'のダメージを受けた！\n' +
          log
      })
      this.$store.commit('user/setLog', log)
    },

    secondFunc: function () {
      this.lowerHP()
      this.writeLog()
    },
  },

  created() {
    this.$store.dispatch('tasks/setTasks')
    this.$store.dispatch('user/getHP')
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

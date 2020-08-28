<template>
  <v-row v-resize="setSlicedTasks">
    <v-col cols="3" sm="2" v-for="(task, i) in slicedTasks" :key="i">
      <Monster :task="task" />
    </v-col>
  </v-row>
</template>

<script>
import Monster from './Monster'
export default {
  components: {
    Monster,
  },
  name: 'MonsterView',

  data: () => ({
    tasksLen: 0,
    slicedTasks: [],
  }),

  methods: {
    // 3~6匹までモンスターを表示する用
    setSlicedTasks: function () {
      let monsterNum = 6
      if (window.innerWidth < 600) {
        monsterNum = 4
      }

      if (this.tasks.length <= monsterNum) {
        this.slicedTasks = JSON.parse(JSON.stringify(this.tasks))
      } else {
        this.slicedTasks = JSON.parse(
          JSON.stringify(this.tasks.slice(0, monsterNum))
        )
      }
    },
  },

  computed: {
    tasks: {
      get() {
        return this.$store.state.tasks.tasks
      },
    },
  },

  watch: {
    tasks: function () {
      if (this.tasks.length < this.tasksLen) {
        // アニメを表示させる時間を稼ぐため1秒待たせる

        setTimeout(() => {
          this.setSlicedTasks()
        }, 1000)
      } else {
        this.setSlicedTasks()
      }
      this.tasksLen = this.tasks.length
    },
  },

  created() {
    this.setSlicedTasks()
  },
}
</script>

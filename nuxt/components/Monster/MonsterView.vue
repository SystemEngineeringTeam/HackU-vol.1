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
        this.slicedTasks = this.tasks
      } else {
        this.slicedTasks = this.tasks.slice(0, monsterNum)
      }
    }
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
      this.setSlicedTasks()
    },
  },

  created() {
    this.setSlicedTasks()
  },
}
</script>

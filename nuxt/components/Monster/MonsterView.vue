<template>
  <v-row>
    <v-col cols="2" v-for="(task, i) in tasks" :key="i">
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

  data: () => ({}),

  computed: {
    // 6匹までモンスターを表示する用
    tasks() {
      if (this.$store.state.tasks.tasks.length <= 6) {
        return this.$store.state.tasks.tasks
      } else {
        return this.$store.state.tasks.tasks.slice(0, 6)
      }
    },
  },
}
</script>

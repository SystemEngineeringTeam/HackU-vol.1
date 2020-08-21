<template>
  <v-card :outlined="true">
    <v-card-title class="justify-center">
      ステータス情報
    </v-card-title>
    <v-textarea solo readonly v-model="log"></v-textarea>
    <HPGauge />
  </v-card>
</template>

<script>
import HPGauge from '../components/HPGauge'

export default {
  name: 'Logbox',

  components: {
    HPGauge,
  },

  data: () => ({
    log: '',
    writeLogid: null,
  }),

  methods: {
    writeLog: function () {
      this.tasks.forEach((element) => {
        this.log =
          element.title +
          'の攻撃！' +
          this.userName +
          'は' +
          1 +
          'のダメージを受けた！\n' +
          this.log
      })
    },
  },

  computed: {
    tasks: {
      get() {
        return this.$store.state.tasks.tasks
      },
    },
    userName: {
      get() {
        return this.$store.state.user.name
      },
    },
  },

  created() {
    this.writeLogid = setInterval(this.writeLog, 1000)
  },

  destroyed() {
    setInterval(this.writeLogid)
  },
}
</script>
